// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package desktopvirtualization_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

func TestAccVirtualDesktopAppAttachPackage_update_host_pool_references(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.updateHostPoolReferences(data, `[
    azurerm_virtual_desktop_host_pool.test.id
  ]`),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateHostPoolReferences(data, `[
    azurerm_virtual_desktop_host_pool.test2.id
  ]`),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateHostPoolReferences(data, `[
    azurerm_virtual_desktop_host_pool.test.id
  ]`),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateHostPoolReferences(data, `[
    azurerm_virtual_desktop_host_pool.test.id,
    azurerm_virtual_desktop_host_pool.test2.id
  ]`),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateHostPoolReferences(data, `[
    azurerm_virtual_desktop_host_pool.test.id
  ]`),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopAppAttachPackage_two_host_pool_references(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.updateHostPoolReferences(data, `[
    azurerm_virtual_desktop_host_pool.test.id,
    azurerm_virtual_desktop_host_pool.test2.id
  ]`),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateHostPoolReferences(data, `[
    azurerm_virtual_desktop_host_pool.test.id
  ]`),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccVirtualDesktopAppAttachPackage_update_fail_health_check_on_staging_failure(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.updateFailHealthCheckOnStagingFailure(data, ""),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateFailHealthCheckOnStagingFailure(data, "DoNotFail"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateFailHealthCheckOnStagingFailure(data, "Unhealthy"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateFailHealthCheckOnStagingFailure(data, "DoNotFail"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateFailHealthCheckOnStagingFailure(data, ""),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopAppAttachPackage_update_image_uri(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.updateImageUri(data, "azurerm_storage_share_file.test6.id"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateImageUri(data, "azurerm_storage_share_file.test_image2.id"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateImageUri(data, "azurerm_storage_share_file.test6.id"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r VirtualDesktopAppAttachPackageResource) updateImageUri(data acceptance.TestData, imageUri string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_storage_share_file" "test_image2" {
  name              = "xmlNotepad2.cim"
  storage_share_url = azurerm_storage_share.test.url
  source            = "${path.module}/testdata/xmlNotepad.cim"
}

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name      = "XmlNotepad"
  image_uri         = %[3]s
  package_full_name = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger, imageUri)
}

func (r VirtualDesktopAppAttachPackageResource) updateFailHealthCheckOnStagingFailure(data acceptance.TestData, failHealthCheckOnStagingFailure string) string {
	failHealthCheckOnStagingFailureConfig := ""
	if failHealthCheckOnStagingFailure != "" {
		failHealthCheckOnStagingFailureConfig = fmt.Sprintf(`fail_health_check_on_staging_failure = %q`, failHealthCheckOnStagingFailure)
	}

	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name      = "XmlNotepad"
  image_uri         = azurerm_storage_share_file.test6.id
  package_full_name = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"
  %[3]s

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger, failHealthCheckOnStagingFailureConfig)
}

func (r VirtualDesktopAppAttachPackageResource) updateHostPoolReferences(data acceptance.TestData, hostPoolReferences string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_host_pool" "test2" {
  name                = "acctest-vdpool2-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  type                = "Pooled"
  load_balancer_type  = "BreadthFirst"
}

resource "azurerm_virtual_desktop_host_pool_registration_info" "test2" {
  hostpool_id     = azurerm_virtual_desktop_host_pool.test2.id
  expiration_date = azurerm_virtual_desktop_host_pool_registration_info.test.expiration_date
}

resource "azurerm_network_interface" "test2" {
  name                = "acctest-nic2-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  ip_configuration {
    name                          = "acctest-ipconfig2-%[2]d"
    private_ip_address_allocation = "Dynamic"
    subnet_id                     = azurerm_subnet.test0.id
  }
}

resource "azurerm_windows_virtual_machine" "test2" {
  name                = "vm2-%[3]s"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  size                = "Standard_F1als_v7"
  admin_password      = "Password1234"
  admin_username      = "adminuser"
  network_interface_ids = [
    azurerm_network_interface.test2.id
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

resource "azurerm_virtual_machine_extension" "test3" {
  name                 = "acctest-vmext-3-%[2]d"
  virtual_machine_id   = azurerm_windows_virtual_machine.test2.id
  publisher            = "Microsoft.Azure.Security.WindowsAttestation"
  type                 = "GuestAttestation"
  type_handler_version = "1.0"

  depends_on = [
    azurerm_nat_gateway.test
  ]
}

resource "azurerm_virtual_machine_extension" "test4" {
  name                 = "acctest-vmext-4-%[2]d"
  virtual_machine_id   = azurerm_windows_virtual_machine.test2.id
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
      hostPoolName = azurerm_virtual_desktop_host_pool.test2.name
      aadJoin      = true
    }
  })

  depends_on = [
    azurerm_nat_gateway.test
  ]
}

resource "azurerm_virtual_machine_extension" "test5" {
  name                 = "acctest-vmext-5-%[2]d"
  virtual_machine_id   = azurerm_windows_virtual_machine.test2.id
  publisher            = "Microsoft.Azure.ActiveDirectory"
  type                 = "AADLoginForWindows"
  type_handler_version = "2.2"

  depends_on = [
    azurerm_nat_gateway.test
  ]
}

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                 = "acctest-msix-%[2]d"
  resource_group_name  = azurerm_resource_group.test.name
  location             = azurerm_resource_group.test.location
  host_pool_references = %[4]s

  display_name      = "XmlNotepad"
  image_uri         = azurerm_storage_share_file.test6.id
  package_full_name = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2,
    azurerm_virtual_machine_extension.test3,
    azurerm_virtual_machine_extension.test4,
    azurerm_virtual_machine_extension.test5
  ]
}
`, r.template(data), data.RandomInteger, data.RandomString, hostPoolReferences)
}
