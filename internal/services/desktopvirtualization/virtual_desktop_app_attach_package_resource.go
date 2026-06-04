// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package desktopvirtualization

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/location"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/hostpool"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2025-10-10/appattachpackage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2025-10-10/msiximage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/desktopvirtualization/method"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

//go:generate go run ../../tools/generator-tests resourceidentity -resource-name virtual_desktop_app_attach_package -service-package-name desktopvirtualization -properties "name,resource_group_name" -known-values "subscription_id:data.Subscriptions.Primary"

type VirtualDesktopAppAttachPackageResource struct{}

var (
	_ sdk.Resource             = VirtualDesktopAppAttachPackageResource{}
	_ sdk.ResourceWithUpdate   = VirtualDesktopAppAttachPackageResource{}
	_ sdk.ResourceWithIdentity = VirtualDesktopAppAttachPackageResource{}
)

func (r VirtualDesktopAppAttachPackageResource) Identity() resourceids.ResourceId {
	return &appattachpackage.AppAttachPackageId{}
}

type VirtualDesktopAppAttachPackageModel struct {
	Name                            string                        `tfschema:"name"`
	ResourceGroupName               string                        `tfschema:"resource_group_name"`
	Location                        string                        `tfschema:"location"`
	HostPoolReferences              []string                      `tfschema:"host_pool_references"`
	FailHealthCheckOnStagingFailure string                        `tfschema:"fail_health_check_on_staging_failure"`
	ImageUri                        string                        `tfschema:"image_uri"`
	PackageFullName                 string                        `tfschema:"package_full_name"`
	DisplayName                     string                        `tfschema:"display_name"`
	IsRegularRegistration           bool                          `tfschema:"is_regular_registration"`
	IsActive                        bool                          `tfschema:"is_active"`
	LastUpdated                     string                        `tfschema:"last_updated"`
	PackageFamilyName               string                        `tfschema:"package_family_name"`
	PackageName                     string                        `tfschema:"package_name"`
	PackageRelativePath             string                        `tfschema:"package_relative_path"`
	Version                         string                        `tfschema:"version"`
	PackageApplications             []MsixPackageApplicationModel `tfschema:"package_applications"`
	Tags                            map[string]string             `tfschema:"tags"`
}

type MsixPackageApplicationModel struct {
	AppId          string `tfschema:"app_id"`
	AppUserModelID string `tfschema:"app_user_model_id"`
	Description    string `tfschema:"description"`
	FriendlyName   string `tfschema:"friendly_name"`
	IconImageName  string `tfschema:"icon_image_name"`
	RawIcon        string `tfschema:"raw_icon"`
	RawPng         string `tfschema:"raw_png"`
}

func (r VirtualDesktopAppAttachPackageResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringDoesNotContainAny("\\/+?&"),
		},

		"resource_group_name": commonschema.ResourceGroupName(),

		"location": commonschema.Location(),

		"host_pool_references": {
			Type:     pluginsdk.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: hostpool.ValidateHostPoolID,
			},
		},

		"fail_health_check_on_staging_failure": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			Default:      appattachpackage.FailHealthCheckOnStagingFailureNeedsAssistance,
			ValidateFunc: validation.StringInSlice(appattachpackage.PossibleValuesForFailHealthCheckOnStagingFailure(), false),
		},

		"image_uri": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.IsURLWithHTTPorHTTPS,
		},

		"package_full_name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"display_name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"is_regular_registration": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  true,
		},

		"is_active": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  false,
		},

		"tags": commonschema.Tags(),
	}
}

func (r VirtualDesktopAppAttachPackageResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"last_updated": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"package_family_name": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"package_name": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"package_relative_path": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"version": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"package_applications": {
			Type:     pluginsdk.TypeList,
			Computed: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"app_id": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"app_user_model_id": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"description": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"friendly_name": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"icon_image_name": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"raw_icon": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"raw_png": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},
				},
			},
		},
	}
}

func (r VirtualDesktopAppAttachPackageResource) ModelObject() interface{} {
	return &VirtualDesktopAppAttachPackageModel{}
}

func (r VirtualDesktopAppAttachPackageResource) ResourceType() string {
	return "azurerm_virtual_desktop_app_attach_package"
}

func (r VirtualDesktopAppAttachPackageResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DesktopVirtualization.AppAttachPackagesClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			var model VirtualDesktopAppAttachPackageModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			id := appattachpackage.NewAppAttachPackageID(subscriptionId, model.ResourceGroupName, model.Name)

			if !metadata.Client.Features.SkipImportCheckOnCreateAndAllowOverwritingExistingResources {
				existing, err := client.Get(ctx, id)
				if err != nil && !response.WasNotFound(existing.HttpResponse) {
					return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
				}
				if !response.WasNotFound(existing.HttpResponse) {
					return metadata.ResourceRequiresImport(r.ResourceType(), id)
				}
			}

			msixImageProperties, err := getMsixImageProperties(ctx, metadata, model.HostPoolReferences, model.ImageUri, model.PackageFullName)
			if err != nil {
				return fmt.Errorf("retrieving MSIX image properties: %+v", err)
			}

			param := appattachpackage.AppAttachPackage{
				Location: location.Normalize(model.Location),
				Properties: appattachpackage.AppAttachPackageProperties{
					FailHealthCheckOnStagingFailure: pointer.ToEnum[appattachpackage.FailHealthCheckOnStagingFailure](model.FailHealthCheckOnStagingFailure),
					HostPoolReferences:              pointer.To(model.HostPoolReferences),
					Image:                           r.expandVirtualDesktopAppAttachPackageImage(model, msixImageProperties),
				},
			}

			if len(model.Tags) > 0 {
				param.Tags = pointer.To(model.Tags)
			}

			if _, err := client.CreateOrUpdate(ctx, id, param); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			if err := pluginsdk.SetResourceIdentityData(metadata.ResourceData, &id); err != nil {
				return err
			}

			return nil
		},
	}
}

func (r VirtualDesktopAppAttachPackageResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DesktopVirtualization.AppAttachPackagesClient

			id, err := appattachpackage.ParseAppAttachPackageID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model VirtualDesktopAppAttachPackageModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			existing, err := client.Get(ctx, *id)
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}
			if existing.Model == nil {
				return fmt.Errorf("retrieving %s: `model` was nil", *id)
			}

			param := *existing.Model

			if metadata.ResourceData.HasChange("fail_health_check_on_staging_failure") {
				param.Properties.FailHealthCheckOnStagingFailure = pointer.ToEnum[appattachpackage.FailHealthCheckOnStagingFailure](model.FailHealthCheckOnStagingFailure)
			}

			if metadata.ResourceData.HasChange("host_pool_references") {
				param.Properties.HostPoolReferences = pointer.To(model.HostPoolReferences)
			}

			if metadata.ResourceData.HasChanges("image_uri", "package_full_name", "display_name", "is_regular_registration", "is_active") {
				msixImageProperties, err := getMsixImageProperties(ctx, metadata, model.HostPoolReferences, model.ImageUri, model.PackageFullName)
				if err != nil {
					return fmt.Errorf("retrieving MSIX image properties: %+v", err)
				}

				param.Properties.Image = r.expandVirtualDesktopAppAttachPackageImage(model, msixImageProperties)
			}

			if metadata.ResourceData.HasChange("tags") {
				param.Tags = pointer.To(model.Tags)
			}

			if _, err := client.CreateOrUpdate(ctx, *id, param); err != nil {
				return fmt.Errorf("updating %s: %+v", *id, err)
			}

			return nil
		},
	}
}

func (r VirtualDesktopAppAttachPackageResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DesktopVirtualization.AppAttachPackagesClient

			id, err := appattachpackage.ParseAppAttachPackageID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			existing, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(existing.HttpResponse) {
					return metadata.MarkAsGone(*id)
				}
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}

			return r.flatten(metadata, id, existing.Model)
		},
	}
}

func (r VirtualDesktopAppAttachPackageResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DesktopVirtualization.AppAttachPackagesClient

			id, err := appattachpackage.ParseAppAttachPackageID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			if _, err := client.Delete(ctx, *id); err != nil {
				return fmt.Errorf("deleting %s: %+v", *id, err)
			}

			return nil
		},
	}
}

func (r VirtualDesktopAppAttachPackageResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return appattachpackage.ValidateAppAttachPackageID
}

func (r VirtualDesktopAppAttachPackageResource) flatten(metadata sdk.ResourceMetaData, id *appattachpackage.AppAttachPackageId, model *appattachpackage.AppAttachPackage) error {
	state := VirtualDesktopAppAttachPackageModel{
		Name:              id.AppAttachPackageName,
		ResourceGroupName: id.ResourceGroupName,
	}

	if model != nil {
		state.Location = location.Normalize(model.Location)
		state.Tags = pointer.From(model.Tags)

		props := model.Properties
		if props.FailHealthCheckOnStagingFailure != nil {
			state.FailHealthCheckOnStagingFailure = pointer.FromEnum(props.FailHealthCheckOnStagingFailure)
		}

		state.HostPoolReferences = pointer.From(props.HostPoolReferences)
		for i, hostPoolReference := range state.HostPoolReferences {
			hostPoolReference = strings.ReplaceAll(hostPoolReference, "resourcegroups", "resourceGroups")
			state.HostPoolReferences[i] = strings.ReplaceAll(hostPoolReference, "hostpools", "hostPools")
		}

		r.flattenVirtualDesktopAppAttachPackageImage(&state, props.Image)
	}

	if err := pluginsdk.SetResourceIdentityData(metadata.ResourceData, id); err != nil {
		return err
	}
	return metadata.Encode(&state)
}

func (r VirtualDesktopAppAttachPackageResource) expandVirtualDesktopAppAttachPackageImage(model VirtualDesktopAppAttachPackageModel, msixImageProperties *msiximage.ExpandMsixImageProperties) *appattachpackage.AppAttachPackageInfoProperties {
	imagePath := strings.ReplaceAll(model.ImageUri, "/", "\\")
	imagePath = strings.TrimLeft(imagePath, "https:")

	return &appattachpackage.AppAttachPackageInfoProperties{
		DisplayName:           pointer.To(model.DisplayName),
		ImagePath:             pointer.To(imagePath),
		IsActive:              pointer.To(model.IsActive),
		IsRegularRegistration: pointer.To(model.IsRegularRegistration),
		LastUpdated:           msixImageProperties.LastUpdated,
		PackageApplications:   r.expandVirtualDesktopAppAttachPackageApplications(msixImageProperties.PackageApplications),
		PackageFamilyName:     msixImageProperties.PackageFamilyName,
		PackageFullName:       pointer.To(model.PackageFullName),
		PackageName:           msixImageProperties.PackageName,
		PackageRelativePath:   msixImageProperties.PackageRelativePath,
		Version:               msixImageProperties.Version,
	}
}

func (r VirtualDesktopAppAttachPackageResource) expandVirtualDesktopAppAttachPackageApplications(inputs *[]msiximage.MsixPackageApplications) *[]appattachpackage.MsixPackageApplications {
	outputs := make([]appattachpackage.MsixPackageApplications, 0)
	if inputs == nil {
		return pointer.To(outputs)
	}

	for _, input := range *inputs {
		outputs = append(outputs, appattachpackage.MsixPackageApplications{
			AppId:          input.AppId,
			AppUserModelID: input.AppUserModelID,
			Description:    input.Description,
			FriendlyName:   input.FriendlyName,
			IconImageName:  input.IconImageName,
			RawIcon:        input.RawIcon,
			RawPng:         input.RawPng,
		})
	}

	return pointer.To(outputs)
}

func (r VirtualDesktopAppAttachPackageResource) flattenVirtualDesktopAppAttachPackageImage(state *VirtualDesktopAppAttachPackageModel, input *appattachpackage.AppAttachPackageInfoProperties) {
	if input == nil {
		return
	}

	state.PackageFullName = pointer.From(input.PackageFullName)
	state.DisplayName = pointer.From(input.DisplayName)
	state.PackageApplications = r.flattenVirtualDesktopAppAttachPackageApplications(input.PackageApplications)
	state.LastUpdated = pointer.From(input.LastUpdated)
	state.PackageFamilyName = pointer.From(input.PackageFamilyName)
	state.PackageName = pointer.From(input.PackageName)
	state.PackageRelativePath = pointer.From(input.PackageRelativePath)
	state.Version = pointer.From(input.Version)

	if input.ImagePath != nil {
		imageUri := strings.ReplaceAll(pointer.From(input.ImagePath), "\\", "/")
		imageUri = fmt.Sprintf("https:%s", imageUri)
		state.ImageUri = imageUri
	}

	if input.IsRegularRegistration != nil {
		state.IsRegularRegistration = pointer.From(input.IsRegularRegistration)
	}

	if input.IsActive != nil {
		state.IsActive = pointer.From(input.IsActive)
	}
}

func (r VirtualDesktopAppAttachPackageResource) flattenVirtualDesktopAppAttachPackageApplications(inputs *[]appattachpackage.MsixPackageApplications) []MsixPackageApplicationModel {
	outputs := make([]MsixPackageApplicationModel, 0)
	if inputs == nil {
		return outputs
	}

	for _, input := range pointer.From(inputs) {
		outputs = append(outputs, MsixPackageApplicationModel{
			AppId:          pointer.From(input.AppId),
			AppUserModelID: pointer.From(input.AppUserModelID),
			Description:    pointer.From(input.Description),
			FriendlyName:   pointer.From(input.FriendlyName),
			IconImageName:  pointer.From(input.IconImageName),
			RawIcon:        pointer.From(input.RawIcon),
			RawPng:         pointer.From(input.RawPng),
		})
	}

	return outputs
}

func getMsixImageProperties(ctx context.Context, metadata sdk.ResourceMetaData, hostPoolReferences []string, imageUri string, packageFullName string) (*msiximage.ExpandMsixImageProperties, error) {
	msixImageUri := msiximage.MSIXImageURI{
		Uri: pointer.To(imageUri),
	}

	var msixImageProperties *msiximage.ExpandMsixImageProperties = nil
	availablePackageFullNames := make(map[string][]string)

	for _, hostPoolReference := range hostPoolReferences {
		availablePackageFullNames[hostPoolReference] = make([]string, 0)
		hostPoolId, err := hostpool.ParseHostPoolIDInsensitively(hostPoolReference)
		if err != nil {
			return nil, fmt.Errorf("parsing host pool reference %s: %+v", hostPoolReference, err)
		}

		msixImageHostPoolId := msiximage.NewHostPoolID(hostPoolId.SubscriptionId, hostPoolId.ResourceGroupName, hostPoolId.HostPoolName)
		result, err := method.ExpandCompleteMsixImage(ctx, metadata, msixImageHostPoolId, msixImageUri)
		if err != nil {
			return nil, fmt.Errorf("expanding MSIX image for host pool %s: %+v", hostPoolReference, err)
		}

		msixImages := result.Items
		for _, msixImage := range msixImages {
			if properties := msixImage.Properties; properties != nil {
				if properties.PackageFullName != nil && strings.EqualFold(pointer.From(properties.PackageFullName), packageFullName) {
					msixImageProperties = properties
					break
				}
			}
		}

		if msixImageProperties == nil {
			for _, msixImage := range msixImages {
				if properties := msixImage.Properties; properties != nil && properties.PackageFullName != nil {
					availablePackageFullNames[hostPoolReference] = append(availablePackageFullNames[hostPoolReference], pointer.From(properties.PackageFullName))
				}
			}
		} else {
			break
		}
	}

	if msixImageProperties == nil {
		concatenatedAvailablePackageFullNames := ""
		for hostPoolReference, packageFullNames := range availablePackageFullNames {
			concatenatedAvailablePackageFullNames += fmt.Sprintf("%v from %s, ", packageFullNames, hostPoolReference)
		}

		concatenatedAvailablePackageFullNames = strings.TrimSuffix(concatenatedAvailablePackageFullNames, ", ")

		return nil, fmt.Errorf("no matched MSIX image with package full name %s was found. The available package full names are %s", packageFullName, concatenatedAvailablePackageFullNames)
	}

	return msixImageProperties, nil
}
