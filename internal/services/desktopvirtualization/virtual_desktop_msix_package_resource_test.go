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
  name     = "acctestRG-vdesktop-%[1]d"
  location = "%[2]s"
}

resource "azurerm_virtual_desktop_host_pool" "test" {
  name                = "acctest-vdpool-%[1]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  type                = "Pooled"
  load_balancer_type  = "BreadthFirst"
}

resource "azurerm_virtual_desktop_msix_package" "test" {
  name                = "acctest-msix-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  host_pool_name      = azurerm_virtual_desktop_host_pool.test.name
}
`, data.RandomInteger, data.Locations.Secondary)
}
