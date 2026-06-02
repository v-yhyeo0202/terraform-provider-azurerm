// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package desktopvirtualization

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/location"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/hostpool"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2025-10-10/appattachpackage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

type VirtualDesktopAppAttachPackageResource struct{}

var (
	_ sdk.Resource           = VirtualDesktopAppAttachPackageResource{}
	_ sdk.ResourceWithUpdate = VirtualDesktopAppAttachPackageResource{}
)

type VirtualDesktopAppAttachPackageModel struct {
	Name                            string                                     `tfschema:"name"`
	ResourceGroupName               string                                     `tfschema:"resource_group_name"`
	Location                        string                                     `tfschema:"location"`
	HostPoolReferences              []string                                   `tfschema:"host_pool_references"`
	FailHealthCheckOnStagingFailure string                                     `tfschema:"fail_health_check_on_staging_failure"`
	Image                           []VirtualDesktopAppAttachPackageImageModel `tfschema:"image"`
	Tags                            map[string]string                          `tfschema:"tags"`
}

type VirtualDesktopAppAttachPackageImageModel struct {
	ImagePath             string `tfschema:"image_path"`
	PackageFullName       string `tfschema:"package_full_name"`
	DisplayName           string `tfschema:"display_name"`
	IsRegularRegistration bool   `tfschema:"is_regular_registration"`
	IsActive              bool   `tfschema:"is_active"`
}

func (r VirtualDesktopAppAttachPackageResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"resource_group_name": commonschema.ResourceGroupName(),

		"location": commonschema.Location(),

		"host_pool_references": {
			Type:     pluginsdk.TypeList,
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

		"image": {
			Type:     pluginsdk.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"image_path": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},

					"package_full_name": {
						Type:     pluginsdk.TypeString,
						Required: true,
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
				},
			},
		},

		"tags": commonschema.Tags(),
	}
}

func (r VirtualDesktopAppAttachPackageResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
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

			existing, err := client.Get(ctx, id)
			if err != nil && !response.WasNotFound(existing.HttpResponse) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}
			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			param := appattachpackage.AppAttachPackage{
				Location: location.Normalize(model.Location),
				Properties: appattachpackage.AppAttachPackageProperties{
					FailHealthCheckOnStagingFailure: pointer.ToEnum[appattachpackage.FailHealthCheckOnStagingFailure](model.FailHealthCheckOnStagingFailure),
					HostPoolReferences:              pointer.To(model.HostPoolReferences),
					Image:                           expandVirtualDesktopAppAttachPackageImage(model.Image),
				},
			}

			if len(model.Tags) > 0 {
				param.Tags = pointer.To(model.Tags)
			}

			if _, err := client.CreateOrUpdate(ctx, id, param); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)

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

			if metadata.ResourceData.HasChange("image") {
				param.Properties.Image = expandVirtualDesktopAppAttachPackageImage(model.Image)
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
			return nil
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

func expandVirtualDesktopAppAttachPackageImage(input []VirtualDesktopAppAttachPackageImageModel) *appattachpackage.AppAttachPackageInfoProperties {
	image := input[0]

	return &appattachpackage.AppAttachPackageInfoProperties{
		ImagePath:             pointer.To(image.ImagePath),
		PackageFullName:       pointer.To(image.PackageFullName),
		DisplayName:           pointer.To(image.DisplayName),
		IsRegularRegistration: pointer.To(image.IsRegularRegistration),
		IsActive:              pointer.To(image.IsActive),
	}
}
