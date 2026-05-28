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

locals {
  cim_file_paths = [
    "objectid_c57e6597-3a9f-4723-8865-3272302f8c12_0",
    "objectid_c57e6597-3a9f-4723-8865-3272302f8c12_1",
    "objectid_c57e6597-3a9f-4723-8865-3272302f8c12_2",
    "region_c57e6597-3a9f-4723-8865-3272302f8c12_0",
    "region_c57e6597-3a9f-4723-8865-3272302f8c12_1",
    "region_c57e6597-3a9f-4723-8865-3272302f8c12_2",
    "xmlNotepad.cim"
  ]
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-vdesktop-%[1]d"
  location = "%[2]s"
}

resource "azurerm_storage_account" "test" {
  name                     = "acctestst%[3]d"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_storage_share" "test" {
  name               = "acctest-share-%[1]d"
  storage_account_id = azurerm_storage_account.test.id
  quota              = 16
}

resource "azurerm_storage_share_file" "test" {
  for_each = toset(local.cim_file_paths)

  name              = each.value
  storage_share_url = azurerm_storage_share.test.url
  source            = "${path.module}/testdata/${each.value}"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctest-vnet-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  address_space       = ["10.0.0.0/24"]
}

resource "azurerm_subnet" "test0" {
  name                 = "acctest-snet-%[1]d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefixes     = ["10.0.0.0/28"]
}

resource "azurerm_subnet" "test1" {
  name                 = "GatewaySubnet"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefixes     = ["10.0.0.16/28"]
}

resource "azurerm_network_interface" "test" {
  name                = "acctest-nic-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  ip_configuration {
    name                          = "acctest-ipconfig-%[1]d"
    private_ip_address_allocation = "Dynamic"
    subnet_id                     = azurerm_subnet.test0.id
  }
}

resource "azurerm_nat_gateway" "test" {
  name                = "acctest-ng-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_windows_virtual_machine" "test" {
  name                = "acctest-vm-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  size                = "Standard_F2as_v7"
  admin_password      = "Password1234"
  admin_username      = "adminuser"
  network_interface_ids = [
    azurerm_network_interface.test.id
  ]

  secure_boot_enabled = true
  vtpm_enabled        = true

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "microsoftwindowsdesktop"
    offer     = "office-365"
    sku       = "win11-24h2-avd-m365"
    version   = "latest"
  }

  identity {
    type = "SystemAssigned"
  }
}

resource "azurerm_virtual_machine_extension" "test0" {
  name                 = "acctest-vmext-0-%[1]d"
  virtual_machine_id   = azurerm_windows_virtual_machine.test.id
  publisher            = "Microsoft.Azure.Security.WindowsAttestation"
  type                 = "GuestAttestation"
  type_handler_version = "1.0"

  depends_on = [
    azurerm_nat_gateway.test
  ]
}

resource "azurerm_virtual_machine_extension" "test1" {
  name                 = "acctest-vmext-1-%[1]d"
  virtual_machine_id   = azurerm_windows_virtual_machine.test.id
  publisher            = "Microsoft.Powershell"
  type                 = "DSC"
  type_handler_version = "2.83"

  protected_settings = jsonencode({
    properties = {
      registrationInfoToken = azurerm_virtual_desktop_host_pool_registration_info.test.token
    }
  })

  settings = jsonencode({
    modulesUrl            = "https://wvdportalstorageblob.blob.core.windows.net/galleryartifacts/Configuration_01-20-2022.zip"
    configurationFunction = "Configuration.ps1\\AddSessionHost"
    properties = {
      hostPoolName = azurerm_virtual_desktop_host_pool.test.name
      aadJoin      = true
    }
  })

  depends_on = [
    azurerm_nat_gateway.test
  ]
}

resource "azurerm_virtual_machine_extension" "test2" {
  name                 = "acctest-vmext-2-%[1]d"
  virtual_machine_id   = azurerm_windows_virtual_machine.test.id
  publisher            = "Microsoft.Azure.ActiveDirectory"
  type                 = "AADLoginForWindows"
  type_handler_version = "2.2"

  depends_on = [
    azurerm_nat_gateway.test
  ]
}

resource "azurerm_virtual_desktop_host_pool" "test" {
  name                = "acctest-vdpool-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  type                = "Pooled"
  load_balancer_type  = "BreadthFirst"
}

resource "azurerm_virtual_desktop_host_pool_registration_info" "test" {
  hostpool_id     = azurerm_virtual_desktop_host_pool.test.id
  expiration_date = "2026-05-30T00:00:00Z"
}

resource "azurerm_virtual_desktop_msix_package" "test" {
  name                = "acctest-msix-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  display_name        = "XmlNotepad"
  host_pool_name      = azurerm_virtual_desktop_host_pool.test.name
  image_uri           = azurerm_storage_share_file.test["xmlNotepad.cim"].id
  package_full_name   = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
	azurerm_virtual_machine_extension.test1,
	azurerm_virtual_machine_extension.test2
  ]
}
`, data.RandomInteger, data.Locations.Secondary, data.RandomIntOfLength(14))
}
