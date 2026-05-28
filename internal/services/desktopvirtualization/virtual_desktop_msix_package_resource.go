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
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2025-10-10/msixpackage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

//go:generate go run ../../tools/generator-tests resourceidentity -resource-name virtual_desktop_msix_package -service-package-name desktopvirtualization -properties "name,resource_group_name,host_pool_name" -known-values "subscription_id:data.Subscriptions.Primary"

var (
	_ sdk.ResourceWithUpdate   = VirtualDesktopMsixPackageResource{}
	_ sdk.ResourceWithIdentity = VirtualDesktopMsixPackageResource{}
)

type VirtualDesktopMsixPackageResource struct{}

func (r VirtualDesktopMsixPackageResource) Identity() resourceids.ResourceId {
	return &msixpackage.MsixPackageId{}
}

type VirtualDesktopMsixPackageModel struct {
	ResourceGroupName     string `tfschema:"resource_group_name"`
	HostPoolName          string `tfschema:"host_pool_name"`
	Name                  string `tfschema:"name"`
	ImagePath             string `tfschema:"image_path"`
	DisplayName           string `tfschema:"display_name"`
	IsRegularRegistration bool   `tfschema:"is_regular_registration"`
	IsActive              bool   `tfschema:"is_active"`
	LastUpdated           string `tfschema:"last_updated"`
	PackageFamilyName     string `tfschema:"package_family_name"`
	PackageName           string `tfschema:"package_name"`
	PackageRelativePath   string `tfschema:"package_relative_path"`
	Version               string `tfschema:"version"`
	AppId                 string `tfschema:"app_id"`
	AppUserModelID        string `tfschema:"app_user_model_id"`
	Description           string `tfschema:"description"`
	FriendlyName          string `tfschema:"friendly_name"`
	IconImageName         string `tfschema:"icon_image_name"`
	RawIcon               string `tfschema:"raw_icon"`
	RawPng                string `tfschema:"raw_png"`
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

		"image_path": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
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

		"last_updated": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Computed: true,
		},

		"package_family_name": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"package_name": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"package_relative_path": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"version": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"app_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"app_user_model_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"description": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"friendly_name": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"icon_image_name": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"raw_icon": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"raw_png": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},
	}
}

func (r VirtualDesktopMsixPackageResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r VirtualDesktopMsixPackageResource) ModelObject() interface{} {
	return &VirtualDesktopMsixPackageModel{}
}

func (r VirtualDesktopMsixPackageResource) ResourceType() string {
	return "azurerm_virtual_desktop_msix_package"
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

			params := msixpackage.MSIXPackage{
				Properties: msixpackage.MSIXPackageProperties{
					IsRegularRegistration: pointer.To(model.IsRegularRegistration),
					IsActive:              pointer.To(model.IsActive),
				},
			}

			if model.ImagePath != "" {
				params.Properties.ImagePath = pointer.To(model.ImagePath)
			}

			if model.DisplayName != "" {
				params.Properties.DisplayName = pointer.To(model.DisplayName)
			}

			if model.PackageFamilyName != "" {
				params.Properties.PackageFamilyName = pointer.To(model.PackageFamilyName)
			}

			if model.PackageName != "" {
				params.Properties.PackageName = pointer.To(model.PackageName)
			}

			if model.PackageRelativePath != "" {
				params.Properties.PackageRelativePath = pointer.To(model.PackageRelativePath)
			}

			if model.Version != "" {
				params.Properties.Version = pointer.To(model.Version)
			}

			if model.LastUpdated != "" {
				params.Properties.LastUpdated = pointer.To(model.LastUpdated)
			}

			if model.AppId != "" || model.AppUserModelID != "" || model.Description != "" || model.FriendlyName != "" || model.IconImageName != "" || model.RawIcon != "" || model.RawPng != "" {
				app := msixpackage.MsixPackageApplications{}
				if model.AppId != "" {
					app.AppId = pointer.To(model.AppId)
				}
				if model.AppUserModelID != "" {
					app.AppUserModelID = pointer.To(model.AppUserModelID)
				}
				if model.Description != "" {
					app.Description = pointer.To(model.Description)
				}
				if model.FriendlyName != "" {
					app.FriendlyName = pointer.To(model.FriendlyName)
				}
				if model.IconImageName != "" {
					app.IconImageName = pointer.To(model.IconImageName)
				}
				if model.RawIcon != "" {
					app.RawIcon = pointer.To(model.RawIcon)
				}
				if model.RawPng != "" {
					app.RawPng = pointer.To(model.RawPng)
				}
				params.Properties.PackageApplications = &[]msixpackage.MsixPackageApplications{app}
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

			resp, err := client.Get(ctx, *id)
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}
			if resp.Model == nil {
				return fmt.Errorf("retrieving %s: `model` was nil", *id)
			}
			existing := resp.Model

			if metadata.ResourceData.HasChange("display_name") {
				existing.Properties.DisplayName = pointer.To(model.DisplayName)
			}

			if metadata.ResourceData.HasChange("is_regular_registration") {
				existing.Properties.IsRegularRegistration = pointer.To(model.IsRegularRegistration)
			}

			if metadata.ResourceData.HasChange("is_active") {
				existing.Properties.IsActive = pointer.To(model.IsActive)
			}

			if metadata.ResourceData.HasChange("package_family_name") {
				existing.Properties.PackageFamilyName = pointer.To(model.PackageFamilyName)
			}

			if metadata.ResourceData.HasChange("package_name") {
				existing.Properties.PackageName = pointer.To(model.PackageName)
			}

			if metadata.ResourceData.HasChange("package_relative_path") {
				existing.Properties.PackageRelativePath = pointer.To(model.PackageRelativePath)
			}

			if metadata.ResourceData.HasChange("version") {
				existing.Properties.Version = pointer.To(model.Version)
			}

			if metadata.ResourceData.HasChange("last_updated") {
				existing.Properties.LastUpdated = pointer.To(model.LastUpdated)
			}

			if metadata.ResourceData.HasChange("app_id") || metadata.ResourceData.HasChange("app_user_model_id") || metadata.ResourceData.HasChange("description") || metadata.ResourceData.HasChange("friendly_name") || metadata.ResourceData.HasChange("icon_image_name") || metadata.ResourceData.HasChange("raw_icon") || metadata.ResourceData.HasChange("raw_png") {
				app := msixpackage.MsixPackageApplications{}
				if model.AppId != "" {
					app.AppId = pointer.To(model.AppId)
				}
				if model.AppUserModelID != "" {
					app.AppUserModelID = pointer.To(model.AppUserModelID)
				}
				if model.Description != "" {
					app.Description = pointer.To(model.Description)
				}
				if model.FriendlyName != "" {
					app.FriendlyName = pointer.To(model.FriendlyName)
				}
				if model.IconImageName != "" {
					app.IconImageName = pointer.To(model.IconImageName)
				}
				if model.RawIcon != "" {
					app.RawIcon = pointer.To(model.RawIcon)
				}
				if model.RawPng != "" {
					app.RawPng = pointer.To(model.RawPng)
				}
				existing.Properties.PackageApplications = &[]msixpackage.MsixPackageApplications{app}
			}

			if _, err := client.CreateOrUpdate(ctx, *id, *existing); err != nil {
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

			return r.flatten(metadata, id, resp.Model)
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

func (r VirtualDesktopMsixPackageResource) flatten(metadata sdk.ResourceMetaData, id *msixpackage.MsixPackageId, model *msixpackage.MSIXPackage) error {
	state := VirtualDesktopMsixPackageModel{
		ResourceGroupName: id.ResourceGroupName,
		HostPoolName:      id.HostPoolName,
		Name:              id.MsixPackageName,
	}

	if model != nil {
		if model.Properties.DisplayName != nil {
			state.DisplayName = pointer.From(model.Properties.DisplayName)
		}
		if model.Properties.ImagePath != nil {
			state.ImagePath = pointer.From(model.Properties.ImagePath)
		}
		if model.Properties.IsRegularRegistration != nil {
			state.IsRegularRegistration = pointer.From(model.Properties.IsRegularRegistration)
		}
		if model.Properties.IsActive != nil {
			state.IsActive = pointer.From(model.Properties.IsActive)
		}
		if model.Properties.PackageFamilyName != nil {
			state.PackageFamilyName = pointer.From(model.Properties.PackageFamilyName)
		}
		if model.Properties.PackageName != nil {
			state.PackageName = pointer.From(model.Properties.PackageName)
		}
		if model.Properties.PackageRelativePath != nil {
			state.PackageRelativePath = pointer.From(model.Properties.PackageRelativePath)
		}
		if model.Properties.Version != nil {
			state.Version = pointer.From(model.Properties.Version)
		}
		if model.Properties.LastUpdated != nil {
			state.LastUpdated = pointer.From(model.Properties.LastUpdated)
		}
		if model.Properties.PackageApplications != nil && len(*model.Properties.PackageApplications) > 0 {
			a := (*model.Properties.PackageApplications)[0]
			state.AppId = pointer.From(a.AppId)
			state.AppUserModelID = pointer.From(a.AppUserModelID)
			state.Description = pointer.From(a.Description)
			state.FriendlyName = pointer.From(a.FriendlyName)
			state.IconImageName = pointer.From(a.IconImageName)
			state.RawIcon = pointer.From(a.RawIcon)
			state.RawPng = pointer.From(a.RawPng)
		}
	}

	if err := pluginsdk.SetResourceIdentityData(metadata.ResourceData, id); err != nil {
		return err
	}

	return metadata.Encode(&state)
}
