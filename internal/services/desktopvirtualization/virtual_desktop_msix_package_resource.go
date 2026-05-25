// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package desktopvirtualization

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2025-10-10/msixpackage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

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
			Type: pluginsdk.TypeString,
		},

		"name": {
			Type: pluginsdk.TypeString,
		},

		"image_path": {
			Type: pluginsdk.TypeString,
		},

		"display_name": {
			Type: pluginsdk.TypeString,
		},

		"is_regular_registration": {
			Type: pluginsdk.TypeBool,
		},

		"is_active": {
			Type: pluginsdk.TypeBool,
		},
	}
}

func (r VirtualDesktopMsixPackageResource) ModelObject() interface{} {
	return &VirtualDesktopMsixPackageModel{}
}

func (r VirtualDesktopMsixPackageResource) ResourceType() string {
	return "azurerm_virtual_desktop_msix_package"
}

func (r VirtualDesktopMsixPackageResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return msixpackage.ValidateMsixPackageID
}

func (r VirtualDesktopMsixPackageResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r VirtualDesktopMsixPackageResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{}
}

func (r VirtualDesktopMsixPackageResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{}
}

func (r VirtualDesktopMsixPackageResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{}
}
