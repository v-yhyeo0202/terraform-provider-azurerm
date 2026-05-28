// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package desktopvirtualization_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2025-10-10/msixpackage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type VirtualDesktopMsixPackageResource struct{}

func TestAccVirtualDesktopMsixPackage_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_msix_package", "test")
	r := VirtualDesktopMsixPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (VirtualDesktopMsixPackageResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := msixpackage.ParseMsixPackageID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.DesktopVirtualization.MsixPackagesClient.Get(ctx, *id)
	if err != nil {
		return nil, fmt.Errorf("retrieving %s: %+v", *id, err)
	}

	return pointer.To(resp.Model != nil), nil
}

func (VirtualDesktopMsixPackageResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-vdesktop-%d"
  location = "%s"
}

resource "azurerm_virtual_desktop_host_pool" "test" {
  name                 = "acctestHP%s"
  location             = azurerm_resource_group.test.location
  resource_group_name  = azurerm_resource_group.test.name
  type                 = "Pooled"
  validate_environment = true
  load_balancer_type   = "DepthFirst"
}

resource "azurerm_virtual_desktop_msix_package" "test" {
  name                = "acctestMSIX%s_1.0.0.0_neutral__abcd1234efgh5"
  resource_group_name = azurerm_resource_group.test.name
  host_pool_name      = azurerm_virtual_desktop_host_pool.test.name
  image_path          = "\\\\testserver\\msixshare\\testpackage.vhd"
  package_family_name   = "MsixPackage_FamilyName"
  package_name          = "MsixPackage_Name"
  package_relative_path = "MsixPackage_RelativePath"
  version               = "1.0.0.0"
  last_updated          = "2024-01-01T00:00:00"

  package_application {
    app_id            = "MsixPackage_AppId"
    app_user_model_id = "MsixPackage_AppUserModelID"
    description       = "MsixPackage_Description"
    friendly_name     = "MsixPackage_FriendlyName"
    icon_image_name   = "MsixPackage_IconImageName"
    raw_icon          = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNgYAAAAAMAASsJTYQAAAAASUVORK5CYII="
    raw_png           = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNgYAAAAAMAASsJTYQAAAAASUVORK5CYII="
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString, data.RandomString)
}
