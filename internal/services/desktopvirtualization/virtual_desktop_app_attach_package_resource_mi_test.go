// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package desktopvirtualization_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

func TestAccVirtualDesktopAppAttachPackage_mi_host_pool_references(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.miHostPoolReferences(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r VirtualDesktopAppAttachPackageResource) templateMultipleHostPools(data acceptance.TestData) string {
	cimFileNames := []string{
		"objectid_c57e6597-3a9f-4723-8865-3272302f8c12_0",
		"objectid_c57e6597-3a9f-4723-8865-3272302f8c12_1",
		"objectid_c57e6597-3a9f-4723-8865-3272302f8c12_2",
		"region_c57e6597-3a9f-4723-8865-3272302f8c12_0",
		"region_c57e6597-3a9f-4723-8865-3272302f8c12_1",
		"region_c57e6597-3a9f-4723-8865-3272302f8c12_2",
		"xmlNotepad.cim",
		"icon.png",
	}

	fileShareConfig := ""
	for i, cimFileName := range cimFileNames {
		fileShareConfig += fmt.Sprintf(`
resource "azurerm_storage_share_file" "test%[1]d" {
  name              = "%[2]s"
  storage_share_url = azurerm_storage_share.test.url
  source            = "${path.module}/testdata/%[2]s"
}
`, i, cimFileName)
	}

	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-vdesktop-%[1]d"
  location = "%[2]s"
}

resource "azurerm_storage_account" "test" {
  name                     = "acctestst%[3]s"
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

%[4]s

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

resource "azurerm_virtual_desktop_host_pool" "test" {
  count               = 128
  name                = "acctest-vdpool-%[1]d-${count.index}"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  type                = "Pooled"
  load_balancer_type  = "BreadthFirst"
}

resource "azurerm_virtual_desktop_host_pool_registration_info" "test" {
  hostpool_id     = azurerm_virtual_desktop_host_pool.test[0].id
  expiration_date = "%[5]s"
}

resource "azurerm_windows_virtual_machine" "test" {
  name                = "vm-%[3]s"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  size                = "Standard_F1als_v7"
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
      hostPoolName = azurerm_virtual_desktop_host_pool.test[0].name
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
`, data.RandomInteger, data.Locations.Secondary, data.RandomString, fileShareConfig, time.Now().UTC().AddDate(0, 0, 1).Format(time.RFC3339))
}

func (r VirtualDesktopAppAttachPackageResource) miHostPoolReferences(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                 = "acctest-msix-%[2]d"
  resource_group_name  = azurerm_resource_group.test.name
  location             = azurerm_resource_group.test.location
  host_pool_references = azurerm_virtual_desktop_host_pool.test[*].id

  display_name      = "XmlNotepad"
  image_uri         = azurerm_storage_share_file.test6.id
  package_full_name = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.templateMultipleHostPools(data), data.RandomInteger)
}
