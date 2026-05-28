// Copyright (c) HashiCorp, Inc.
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
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/hostpool"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2025-10-10/appattachpackage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

//go:generate go run ../../tools/generator-tests resourceidentity -resource-name virtual_desktop_app_attach_package -service-package-name desktopvirtualization -properties "name,resource_group_name" -known-values "subscription_id:data.Subscriptions.Primary"

type VirtualDesktopAppAttachPackageResource struct{}

var (
	_ sdk.ResourceWithUpdate   = VirtualDesktopAppAttachPackageResource{}
	_ sdk.ResourceWithIdentity = VirtualDesktopAppAttachPackageResource{}
)

func (r VirtualDesktopAppAttachPackageResource) Identity() resourceids.ResourceId {
	return &appattachpackage.AppAttachPackageId{}
}

type VirtualDesktopAppAttachPackageImageModel struct {
	ImagePath             string `tfschema:"image_path"`
	PackageFullName       string `tfschema:"package_full_name"`
	DisplayName           string `tfschema:"display_name"`
	IsRegularRegistration bool   `tfschema:"is_regular_registration"`
	IsActive              bool   `tfschema:"is_active"`
}

type VirtualDesktopAppAttachPackageModel struct {
	ResourceGroupName               string                                     `tfschema:"resource_group_name"`
	Location                        string                                     `tfschema:"location"`
	Name                            string                                     `tfschema:"name"`
	HostPoolReferences              []string                                   `tfschema:"host_pool_references"`
	Image                           []VirtualDesktopAppAttachPackageImageModel `tfschema:"image"`
	FailHealthCheckOnStagingFailure string                                     `tfschema:"fail_health_check_on_staging_failure"`
	Tags                            map[string]string                          `tfschema:"tags"`
}

func (r VirtualDesktopAppAttachPackageResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"resource_group_name": commonschema.ResourceGroupName(),

		"location": commonschema.Location(),

		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"host_pool_references": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: hostpool.ValidateHostPoolID,
			},
		},

		"image": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"image_path": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"package_full_name": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"display_name": {
						Type:     pluginsdk.TypeString,
						Optional: true,
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

		"fail_health_check_on_staging_failure": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			Default:      appattachpackage.FailHealthCheckOnStagingFailureNeedsAssistance,
			ValidateFunc: validation.StringInSlice(appattachpackage.PossibleValuesForFailHealthCheckOnStagingFailure(), false),
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
			if err != nil {
				if !response.WasNotFound(existing.HttpResponse) {
					return fmt.Errorf("checking for the presence of an existing %s: %+v", id, err)
				}
			}
			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			params := appattachpackage.AppAttachPackage{
				Location:   location.Normalize(model.Location),
				Properties: appattachpackage.AppAttachPackageProperties{},
				Tags:       pointer.To(model.Tags),
			}

			fmt.Println("debug0", model.FailHealthCheckOnStagingFailure)
			if model.FailHealthCheckOnStagingFailure != "" {
				failHealthCheck := appattachpackage.FailHealthCheckOnStagingFailure(model.FailHealthCheckOnStagingFailure)
				params.Properties.FailHealthCheckOnStagingFailure = &failHealthCheck
			}

			if len(model.HostPoolReferences) > 0 {
				params.Properties.HostPoolReferences = &model.HostPoolReferences
			}

			if len(model.Image) > 0 {
				params.Properties.Image = expandAppAttachPackageImage(model.Image)
			}

			if _, err := client.CreateOrUpdate(ctx, id, params); err != nil {
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

			resp, err := client.Get(ctx, *id)
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}
			if resp.Model == nil {
				return fmt.Errorf("retrieving %s: `model` was nil", *id)
			}
			existing := resp.Model

			if metadata.ResourceData.HasChange("host_pool_references") {
				existing.Properties.HostPoolReferences = &model.HostPoolReferences
			}

			if metadata.ResourceData.HasChange("image") {
				if len(model.Image) > 0 {
					existing.Properties.Image = expandAppAttachPackageImage(model.Image)
				} else {
					existing.Properties.Image = nil
				}
			}

			if metadata.ResourceData.HasChange("fail_health_check_on_staging_failure") {
				existing.Properties.FailHealthCheckOnStagingFailure = pointer.ToEnum[appattachpackage.FailHealthCheckOnStagingFailure](model.FailHealthCheckOnStagingFailure)
			}

			if metadata.ResourceData.HasChange("tags") {
				existing.Tags = pointer.To(model.Tags)
			}

			if _, err := client.CreateOrUpdate(ctx, *id, *existing); err != nil {
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

			resp, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(*id)
				}
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}

			model := resp.Model
			if model == nil {
				return fmt.Errorf("retrieving %s: `model` was nil", *id)
			}

			return r.flatten(metadata, id, model)
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
		ResourceGroupName: id.ResourceGroupName,
		Location:          location.Normalize(model.Location),
		Name:              id.AppAttachPackageName,
	}

	if model.Tags != nil {
		state.Tags = pointer.From(model.Tags)
	}

	props := model.Properties

	if props.FailHealthCheckOnStagingFailure != nil {
		state.FailHealthCheckOnStagingFailure = pointer.FromEnum[appattachpackage.FailHealthCheckOnStagingFailure](props.FailHealthCheckOnStagingFailure)
	}

	if props.HostPoolReferences != nil {
		state.HostPoolReferences = pointer.From(props.HostPoolReferences)
	}

	if props.Image != nil {
		state.Image = flattenAppAttachPackageImage(props.Image)
	}

	if err := pluginsdk.SetResourceIdentityData(metadata.ResourceData, id); err != nil {
		return err
	}

	return metadata.Encode(&state)
}

func expandAppAttachPackageImage(input []VirtualDesktopAppAttachPackageImageModel) *appattachpackage.AppAttachPackageInfoProperties {
	if len(input) == 0 {
		return nil
	}
	item := input[0]
	result := appattachpackage.AppAttachPackageInfoProperties{
		IsRegularRegistration: pointer.To(item.IsRegularRegistration),
		IsActive:              pointer.To(item.IsActive),
	}
	if item.ImagePath != "" {
		result.ImagePath = pointer.To(item.ImagePath)
	}
	if item.PackageFullName != "" {
		result.PackageFullName = pointer.To(item.PackageFullName)
	}
	if item.DisplayName != "" {
		result.DisplayName = pointer.To(item.DisplayName)
	}
	return &result
}

func flattenAppAttachPackageImage(input *appattachpackage.AppAttachPackageInfoProperties) []VirtualDesktopAppAttachPackageImageModel {
	if input == nil {
		return []VirtualDesktopAppAttachPackageImageModel{}
	}
	result := VirtualDesktopAppAttachPackageImageModel{}
	if input.ImagePath != nil {
		result.ImagePath = pointer.From(input.ImagePath)
	}
	if input.PackageFullName != nil {
		result.PackageFullName = pointer.From(input.PackageFullName)
	}
	if input.DisplayName != nil {
		result.DisplayName = pointer.From(input.DisplayName)
	}
	if input.IsRegularRegistration != nil {
		result.IsRegularRegistration = pointer.From(input.IsRegularRegistration)
	}
	if input.IsActive != nil {
		result.IsActive = pointer.From(input.IsActive)
	}
	return []VirtualDesktopAppAttachPackageImageModel{result}
}
