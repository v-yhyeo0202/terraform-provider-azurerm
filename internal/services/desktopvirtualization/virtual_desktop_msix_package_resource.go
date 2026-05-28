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
	ResourceGroupName     string               `tfschema:"resource_group_name"`
	HostPoolName          string               `tfschema:"host_pool_name"`
	Name                  string               `tfschema:"name"`
	ImagePath             string               `tfschema:"image_path"`
	DisplayName           string               `tfschema:"display_name"`
	IsRegularRegistration bool                 `tfschema:"is_regular_registration"`
	IsActive              bool                 `tfschema:"is_active"`
	LastUpdated           string               `tfschema:"last_updated"`
	PackageFamilyName     string               `tfschema:"package_family_name"`
	PackageName           string               `tfschema:"package_name"`
	PackageRelativePath   string               `tfschema:"package_relative_path"`
	Version               string               `tfschema:"version"`
	PackageApplications   []PackageApplication `tfschema:"package_application"`
}

type PackageApplication struct {
	AppId          string `tfschema:"app_id"`
	AppUserModelID string `tfschema:"app_user_model_id"`
	Description    string `tfschema:"description"`
	FriendlyName   string `tfschema:"friendly_name"`
	IconImageName  string `tfschema:"icon_image_name"`
	RawIcon        string `tfschema:"raw_icon"`
	RawPng         string `tfschema:"raw_png"`
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

		"package_application": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
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
				},
			},
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

			if len(model.PackageApplications) > 0 {
				apps := make([]msixpackage.MsixPackageApplications, 0, len(model.PackageApplications))
				for _, a := range model.PackageApplications {
					app := msixpackage.MsixPackageApplications{}
					if a.AppId != "" {
						app.AppId = pointer.To(a.AppId)
					}
					if a.AppUserModelID != "" {
						app.AppUserModelID = pointer.To(a.AppUserModelID)
					}
					if a.Description != "" {
						app.Description = pointer.To(a.Description)
					}
					if a.FriendlyName != "" {
						app.FriendlyName = pointer.To(a.FriendlyName)
					}
					if a.IconImageName != "" {
						app.IconImageName = pointer.To(a.IconImageName)
					}
					if a.RawIcon != "" {
						app.RawIcon = pointer.To(a.RawIcon)
					}
					if a.RawPng != "" {
						app.RawPng = pointer.To(a.RawPng)
					}
					apps = append(apps, app)
				}
				params.Properties.PackageApplications = &apps
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

			if metadata.ResourceData.HasChange("package_application") {
				apps := make([]msixpackage.MsixPackageApplications, 0, len(model.PackageApplications))
				for _, a := range model.PackageApplications {
					app := msixpackage.MsixPackageApplications{}
					if a.AppId != "" {
						app.AppId = pointer.To(a.AppId)
					}
					if a.AppUserModelID != "" {
						app.AppUserModelID = pointer.To(a.AppUserModelID)
					}
					if a.Description != "" {
						app.Description = pointer.To(a.Description)
					}
					if a.FriendlyName != "" {
						app.FriendlyName = pointer.To(a.FriendlyName)
					}
					if a.IconImageName != "" {
						app.IconImageName = pointer.To(a.IconImageName)
					}
					if a.RawIcon != "" {
						app.RawIcon = pointer.To(a.RawIcon)
					}
					if a.RawPng != "" {
						app.RawPng = pointer.To(a.RawPng)
					}
					apps = append(apps, app)
				}
				existing.Properties.PackageApplications = &apps
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
		if model.Properties.PackageApplications != nil {
			apps := make([]PackageApplication, 0, len(*model.Properties.PackageApplications))
			for _, a := range *model.Properties.PackageApplications {
				apps = append(apps, PackageApplication{
					AppId:          pointer.From(a.AppId),
					AppUserModelID: pointer.From(a.AppUserModelID),
					Description:    pointer.From(a.Description),
					FriendlyName:   pointer.From(a.FriendlyName),
					IconImageName:  pointer.From(a.IconImageName),
					RawIcon:        pointer.From(a.RawIcon),
					RawPng:         pointer.From(a.RawPng),
				})
			}
			state.PackageApplications = apps
		}
	}

	if err := pluginsdk.SetResourceIdentityData(metadata.ResourceData, id); err != nil {
		return err
	}

	return metadata.Encode(&state)
}
