// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package desktopvirtualization

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/hostpool"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2025-10-10/appattachpackage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

type VirtualDesktopAppAttachPackageResource struct{}

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
			Optional: true,
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

func (r VirtualDesktopAppAttachPackageResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return appattachpackage.ValidateAppAttachPackageID
}
