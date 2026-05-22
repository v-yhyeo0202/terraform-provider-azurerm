// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package desktopvirtualization

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/msixpackage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

var (
	_ sdk.Resource           = VirtualDesktopMsixPackageResource{}
	_ sdk.ResourceWithUpdate = VirtualDesktopMsixPackageResource{}
)

type VirtualDesktopMsixPackageResource struct{}

type VirtualDesktopMsixPackageModel struct {
	Name                  string `tfschema:"name"`
	HostPoolId            string `tfschema:"host_pool_id"`
	ImagePath             string `tfschema:"image_path"`
	PackageName           string `tfschema:"package_name"`
	DisplayName           string `tfschema:"display_name"`
	IsRegularRegistration bool   `tfschema:"is_regular_registration"`
	IsActive              bool   `tfschema:"is_active"`
}

func (r VirtualDesktopMsixPackageResource) ModelObject() interface{} {
	return &VirtualDesktopMsixPackageModel{}
}

func (r VirtualDesktopMsixPackageResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return msixpackage.ValidateMsixPackageID
}

func (r VirtualDesktopMsixPackageResource) ResourceType() string {
	return "azurerm_virtual_desktop_msix_package"
}

func (r VirtualDesktopMsixPackageResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"host_pool_id": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"image_path": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"package_name": {
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
		},

		"is_active": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},
	}
}

func (r VirtualDesktopMsixPackageResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r VirtualDesktopMsixPackageResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return fmt.Errorf("not implemented")
		},
	}
}

func (r VirtualDesktopMsixPackageResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return fmt.Errorf("not implemented")
		},
	}
}

func (r VirtualDesktopMsixPackageResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return fmt.Errorf("not implemented")
		},
	}
}

func (r VirtualDesktopMsixPackageResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return fmt.Errorf("not implemented")
		},
	}
}
