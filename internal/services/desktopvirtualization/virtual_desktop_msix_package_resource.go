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
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2025-10-10/msixpackage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

var _ sdk.ResourceWithUpdate = VirtualDesktopMsixPackageResource{}

type VirtualDesktopMsixPackageResource struct{}

type VirtualDesktopMsixPackageModel struct {
	ResourceGroupName     string `tfschema:"resource_group_name"`
	HostPoolName          string `tfschema:"host_pool_name"`
	Name                  string `tfschema:"name"`
	ImagePath             string `tfschema:"image_path"`
	DisplayName           string `tfschema:"display_name"`
	IsRegularRegistration bool   `tfschema:"is_regular_registration"`
	IsActive              bool   `tfschema:"is_active"`
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

			if _, err := client.CreateOrUpdate(ctx, id, params); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
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

			state := VirtualDesktopMsixPackageModel{
				ResourceGroupName: id.ResourceGroupName,
				HostPoolName:      id.HostPoolName,
				Name:              id.MsixPackageName,
			}

			if model := resp.Model; model != nil {
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
			}

			return metadata.Encode(&state)
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
