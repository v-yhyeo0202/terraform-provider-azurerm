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
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2025-10-10/msiximage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2025-10-10/msixpackage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/desktopvirtualization/method"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

//go:generate go run ../../tools/generator-tests resourceidentity -resource-name virtual_desktop_msix_package -service-package-name desktopvirtualization -properties "name" -compare-values "subscription_id:host_pool_id,resource_group_name:host_pool_id,host_pool_name:host_pool_id"

// Use preflight?

var (
	_ sdk.Resource             = VirtualDesktopMsixPackageResource{}
	_ sdk.ResourceWithUpdate   = VirtualDesktopMsixPackageResource{}
	_ sdk.ResourceWithIdentity = VirtualDesktopMsixPackageResource{}
)

type VirtualDesktopMsixPackageResource struct{}

func (r VirtualDesktopMsixPackageResource) Identity() resourceids.ResourceId {
	return &msixpackage.MsixPackageId{}
}

type VirtualDesktopMsixPackageModel struct {
	ResourceGroupName     string                        `tfschema:"resource_group_name"`
	HostPoolName          string                        `tfschema:"host_pool_name"`
	Name                  string                        `tfschema:"name"`
	PackageFullName       string                        `tfschema:"package_full_name"`
	ImageUri              string                        `tfschema:"image_uri"`
	DisplayName           string                        `tfschema:"display_name"`
	IsRegularRegistration bool                          `tfschema:"is_regular_registration"`
	IsActive              bool                          `tfschema:"is_active"`
	LastUpdated           string                        `tfschema:"last_updated"`
	PackageFamilyName     string                        `tfschema:"package_family_name"`
	PackageName           string                        `tfschema:"package_name"`
	PackageRelativePath   string                        `tfschema:"package_relative_path"`
	Version               string                        `tfschema:"version"`
	PackageApplications   []MsixPackageApplicationModel `tfschema:"package_applications"`
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

func (r VirtualDesktopMsixPackageResource) ModelObject() interface{} {
	return &VirtualDesktopMsixPackageModel{}
}

func (r VirtualDesktopMsixPackageResource) ResourceType() string {
	return "azurerm_virtual_desktop_msix_package"
}

func (r VirtualDesktopMsixPackageResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"resource_group_name": commonschema.ResourceGroupName(),

		"host_pool_name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"package_full_name": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"image_uri": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"display_name": {
			Type:     pluginsdk.TypeString,
			Required: true,
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
	}
}

func (r VirtualDesktopMsixPackageResource) Attributes() map[string]*pluginsdk.Schema {
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

func (r VirtualDesktopMsixPackageResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DesktopVirtualization.MsixPackagesClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			var model VirtualDesktopMsixPackageModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			id := msixpackage.NewMsixPackageID(subscriptionId, model.ResourceGroupName, model.HostPoolName, model.Name)

			existing, err := client.Get(ctx, id)
			if err != nil {
				if !response.WasNotFound(existing.HttpResponse) {
					return fmt.Errorf("checking for the presence of an existing %s: %+v", id, err)
				}
			}

			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			hostPoolId := msiximage.NewHostPoolID(subscriptionId, model.ResourceGroupName, model.HostPoolName)
			msixImageProperties, err := getMsixImageProperties(ctx, metadata, hostPoolId, model.ImageUri, model.PackageFullName, "")
			if err != nil {
				return fmt.Errorf("retrieving MSIX image properties: %+v", err)
			}

			imagePath := strings.ReplaceAll(model.ImageUri, "/", "\\")
			imagePath = strings.TrimLeft(imagePath, "https:")

			param := msixpackage.MSIXPackage{
				Properties: msixpackage.MSIXPackageProperties{
					DisplayName:           pointer.To(model.DisplayName),
					ImagePath:             pointer.To(imagePath),
					IsActive:              pointer.To(model.IsActive),
					IsRegularRegistration: pointer.To(model.IsRegularRegistration),
					LastUpdated:           msixImageProperties.LastUpdated,
					PackageApplications:   r.expandPackageApplications(msixImageProperties.PackageApplications),
					PackageFamilyName:     msixImageProperties.PackageFamilyName,
					PackageName:           msixImageProperties.PackageName,
					PackageRelativePath:   msixImageProperties.PackageRelativePath,
					Version:               msixImageProperties.Version,
				},
			}

			if _, err := client.CreateOrUpdate(ctx, id, param); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			if err := pluginsdk.SetResourceIdentityData(metadata.ResourceData, &id); err != nil {
				return err
			}
			panic("debug0")
			return nil
		},
	}
}

func (r VirtualDesktopMsixPackageResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DesktopVirtualization.MsixPackagesClient

			id, err := msixpackage.ParseMsixPackageID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model VirtualDesktopMsixPackageModel
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

			if metadata.ResourceData.HasChange("package_full_name") {
				hostPoolId := msiximage.NewHostPoolID(metadata.Client.Account.SubscriptionId, model.ResourceGroupName, model.HostPoolName)
				msixImageProperties, err := getMsixImageProperties(ctx, metadata, hostPoolId, model.ImageUri, model.PackageFullName, "")
				if err != nil {
					return fmt.Errorf("retrieving MSIX image properties: %+v", err)
				}

				param.Properties.LastUpdated = msixImageProperties.LastUpdated
				param.Properties.PackageApplications = r.expandPackageApplications(msixImageProperties.PackageApplications)
				param.Properties.PackageFamilyName = msixImageProperties.PackageFamilyName
				param.Properties.PackageName = msixImageProperties.PackageName
				param.Properties.PackageRelativePath = msixImageProperties.PackageRelativePath
				param.Properties.Version = msixImageProperties.Version
			}

			if metadata.ResourceData.HasChange("display_name") {
				param.Properties.DisplayName = pointer.To(model.DisplayName)
			}

			if metadata.ResourceData.HasChange("is_regular_registration") {
				param.Properties.IsRegularRegistration = pointer.To(model.IsRegularRegistration)
			}

			if metadata.ResourceData.HasChange("is_active") {
				param.Properties.IsActive = pointer.To(model.IsActive)
			}

			if _, err := client.CreateOrUpdate(ctx, *id, param); err != nil {
				return fmt.Errorf("updating %s: %+v", *id, err)
			}

			return nil
		},
	}
}

func (r VirtualDesktopMsixPackageResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DesktopVirtualization.MsixPackagesClient

			id, err := msixpackage.ParseMsixPackageID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(*id)
				}
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}

			return r.flatten(ctx, metadata, id, resp.Model)
		},
	}
}

func (r VirtualDesktopMsixPackageResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DesktopVirtualization.MsixPackagesClient

			id, err := msixpackage.ParseMsixPackageID(metadata.ResourceData.Id())
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

func (r VirtualDesktopMsixPackageResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return msixpackage.ValidateMsixPackageID
}

func (r VirtualDesktopMsixPackageResource) flatten(ctx context.Context, metadata sdk.ResourceMetaData, id *msixpackage.MsixPackageId, input *msixpackage.MSIXPackage) error {
	state := VirtualDesktopMsixPackageModel{
		ResourceGroupName: id.ResourceGroupName,
		HostPoolName:      id.HostPoolName,
		Name:              id.MsixPackageName,
	}

	if input != nil {
		props := input.Properties

		if imagePath := props.ImagePath; imagePath != nil {
			imageUri := strings.ReplaceAll(pointer.From(imagePath), "\\", "/")
			imageUri = fmt.Sprintf("https:%s", imageUri)
			state.ImageUri = imageUri

			hostPoolId := msiximage.NewHostPoolID(metadata.Client.Account.SubscriptionId, id.ResourceGroupName, id.HostPoolName)
			msixImageProperties, err := getMsixImageProperties(ctx, metadata, hostPoolId, imageUri, "", pointer.From(props.PackageRelativePath))
			if err != nil {
				return fmt.Errorf("retrieving MSIX image properties: %+v", err)
			}

			state.PackageFullName = pointer.From(msixImageProperties.PackageFullName)
		}

		state.DisplayName = pointer.From(props.DisplayName)

		if isRegularRegistration := props.IsRegularRegistration; isRegularRegistration != nil {
			state.IsRegularRegistration = pointer.From(isRegularRegistration)
		}

		if isActive := props.IsActive; isActive != nil {
			state.IsActive = pointer.From(isActive)
		}

		state.LastUpdated = pointer.From(props.LastUpdated)
		state.PackageFamilyName = pointer.From(props.PackageFamilyName)
		state.PackageName = pointer.From(props.PackageName)
		state.PackageRelativePath = pointer.From(props.PackageRelativePath)
		state.Version = pointer.From(props.Version)
		state.PackageApplications = r.flattenPackageApplications(props.PackageApplications)
	}

	if err := pluginsdk.SetResourceIdentityData(metadata.ResourceData, id); err != nil {
		return err
	}
	return metadata.Encode(&state)
}

func (r VirtualDesktopMsixPackageResource) expandPackageApplications(inputs *[]msiximage.MsixPackageApplications) *[]msixpackage.MsixPackageApplications {
	outputs := make([]msixpackage.MsixPackageApplications, 0)
	if inputs == nil {
		return pointer.To(outputs)
	}

	for _, input := range *inputs {
		outputs = append(outputs, msixpackage.MsixPackageApplications{
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

func (r VirtualDesktopMsixPackageResource) flattenPackageApplications(inputs *[]msixpackage.MsixPackageApplications) []MsixPackageApplicationModel {
	outputs := make([]MsixPackageApplicationModel, 0)
	if inputs == nil {
		return outputs
	}

	for _, input := range *inputs {
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

func getMsixImageProperties(ctx context.Context, metadata sdk.ResourceMetaData, hostPoolId msiximage.HostPoolId, imageUri string, packageFullName string, packageRelativePath string) (*msiximage.ExpandMsixImageProperties, error) {
	msixImageUri := msiximage.MSIXImageURI{
		Uri: pointer.To(imageUri),
	}
	result, err := method.ExpandCompleteMsixImage(ctx, metadata, hostPoolId, msixImageUri)
	if err != nil {
		return nil, fmt.Errorf("expanding MSIX image of host pool %s: %+v", hostPoolId, err)
	}

	msixImages := result.Items
	var msixImageProperties *msiximage.ExpandMsixImageProperties = nil
	for _, msixImage := range msixImages {
		if properties := msixImage.Properties; properties != nil {
			packageFullNameMatched := packageFullName != "" && properties.PackageFullName != nil && strings.EqualFold(pointer.From(properties.PackageFullName), packageFullName)

			packageRelativePathMatched := packageRelativePath != "" && properties.PackageRelativePath != nil && strings.EqualFold(pointer.From(properties.PackageRelativePath), packageRelativePath)

			if packageFullNameMatched || packageRelativePathMatched {
				msixImageProperties = properties
				break
			}
		}
	}

	if msixImageProperties == nil {
		if packageFullName != "" {
			availablePackageFullNames := make([]string, 0)
			for _, msixImage := range msixImages {
				if properties := msixImage.Properties; properties != nil && properties.PackageFullName != nil {
					availablePackageFullNames = append(availablePackageFullNames, pointer.From(properties.PackageFullName))
				}
			}

			return nil, fmt.Errorf("no matched MSIX image with package full name %s was found. The available package full names are: %v", packageFullName, availablePackageFullNames)
		}

		availablePackageRelativePaths := make([]string, 0)
		for _, msixImage := range msixImages {
			if properties := msixImage.Properties; properties != nil && properties.PackageRelativePath != nil {
				availablePackageRelativePaths = append(availablePackageRelativePaths, pointer.From(properties.PackageRelativePath))
			}
		}

		return nil, fmt.Errorf("no matched MSIX image with package relative path %s was found. The available package relative paths are: %v", packageRelativePath, availablePackageRelativePaths)
	}

	return msixImageProperties, nil
}
