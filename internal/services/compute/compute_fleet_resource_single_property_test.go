// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package compute_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

// --- spot_priority_profile.capacity ---

func TestAccComputeFleet_spotPriorityProfileCapacity(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.spotPriorityProfileCapacity(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) spotPriorityProfileCapacity(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

resource "azurerm_virtual_network" "test" {
  name                = "acctest-vnet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  address_space       = ["10.0.0.0/16"]
}

resource "azurerm_subnet" "test" {
  name                 = "acctest-subnet-%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefixes     = ["10.0.1.0/24"]
}

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  vm_sizes_profile {
    name = "Standard_F1als_v7"
  }

  create_option = "FromImage"

  os_type = "Linux"

  image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = "22_04-lts-gen2"
    version   = "latest"
  }

  admin_username = "adminuser"
  admin_password = "P@ssw0rd1234!"

  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  spot_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- gallery_applications.package_reference_id ---

func TestAccComputeFleet_galleryApplicationsPackageReferenceId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.galleryApplicationsPackageReferenceId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) galleryApplicationsPackageReferenceId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

resource "azurerm_virtual_network" "test" {
  name                = "acctest-vnet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  address_space       = ["10.0.0.0/16"]
}

resource "azurerm_subnet" "test" {
  name                 = "acctest-subnet-%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefixes     = ["10.0.1.0/24"]
}

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  vm_sizes_profile {
    name = "Standard_F1als_v7"
  }

  create_option = "FromImage"

  os_type = "Linux"

  image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = "22_04-lts-gen2"
    version   = "latest"
  }

  admin_username = "adminuser"
  admin_password = "P@ssw0rd1234!"

  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  gallery_applications {
    package_reference_id = azurerm_gallery_application_version.test.id
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.galleryApplicationsTemplate(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- gallery_applications.enable_automatic_upgrade ---

func TestAccComputeFleet_galleryApplicationsEnableAutomaticUpgrade(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.galleryApplicationsEnableAutomaticUpgrade(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) galleryApplicationsEnableAutomaticUpgrade(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

resource "azurerm_virtual_network" "test" {
  name                = "acctest-vnet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  address_space       = ["10.0.0.0/16"]
}

resource "azurerm_subnet" "test" {
  name                 = "acctest-subnet-%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefixes     = ["10.0.1.0/24"]
}

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  vm_sizes_profile {
    name = "Standard_F1als_v7"
  }

  create_option = "FromImage"

  os_type = "Linux"

  image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = "22_04-lts-gen2"
    version   = "latest"
  }

  admin_username = "adminuser"
  admin_password = "P@ssw0rd1234!"

  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  gallery_applications {
    package_reference_id     = azurerm_gallery_application_version.test.id
    enable_automatic_upgrade = true
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.galleryApplicationsTemplate(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- gallery_applications.configuration_reference ---

func TestAccComputeFleet_galleryApplicationsConfigurationReference(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.galleryApplicationsConfigurationReference(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) galleryApplicationsConfigurationReference(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

resource "azurerm_storage_blob" "config" {
  name                   = "config"
  storage_account_name   = azurerm_storage_account.test.name
  storage_container_name = azurerm_storage_container.test.name
  type                   = "Page"
  size                   = 512
}

resource "azurerm_virtual_network" "test" {
  name                = "acctest-vnet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  address_space       = ["10.0.0.0/16"]
}

resource "azurerm_subnet" "test" {
  name                 = "acctest-subnet-%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefixes     = ["10.0.1.0/24"]
}

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  vm_sizes_profile {
    name = "Standard_F1als_v7"
  }

  create_option = "FromImage"

  os_type = "Linux"

  image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = "22_04-lts-gen2"
    version   = "latest"
  }

  admin_username = "adminuser"
  admin_password = "P@ssw0rd1234!"

  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  gallery_applications {
    package_reference_id    = azurerm_gallery_application_version.test.id
    configuration_reference = azurerm_storage_blob.config.id
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.galleryApplicationsTemplate(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- gallery_applications.order ---

func TestAccComputeFleet_galleryApplicationsOrder(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.galleryApplicationsOrder(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) galleryApplicationsOrder(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

resource "azurerm_virtual_network" "test" {
  name                = "acctest-vnet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  address_space       = ["10.0.0.0/16"]
}

resource "azurerm_subnet" "test" {
  name                 = "acctest-subnet-%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefixes     = ["10.0.1.0/24"]
}

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  vm_sizes_profile {
    name = "Standard_F1als_v7"
  }

  create_option = "FromImage"

  os_type = "Linux"

  image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = "22_04-lts-gen2"
    version   = "latest"
  }

  admin_username = "adminuser"
  admin_password = "P@ssw0rd1234!"

  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  gallery_applications {
    package_reference_id = azurerm_gallery_application_version.test.id
    order                = 1
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.galleryApplicationsTemplate(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) galleryApplicationsTemplate(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_storage_account" "test" {
  name                     = "accteststr%[3]s"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_storage_container" "test" {
  name                  = "test"
  storage_account_name  = azurerm_storage_account.test.name
  container_access_type = "blob"
}

resource "azurerm_storage_blob" "test" {
  name                   = "script"
  storage_account_name   = azurerm_storage_account.test.name
  storage_container_name = azurerm_storage_container.test.name
  type                   = "Page"
  size                   = 512
}

resource "azurerm_shared_image_gallery" "test" {
  name                = "acctestsig%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_gallery_application" "test" {
  name              = "acctest-app-%[2]d"
  gallery_id        = azurerm_shared_image_gallery.test.id
  location          = azurerm_shared_image_gallery.test.location
  supported_os_type = "Linux"
}

resource "azurerm_gallery_application_version" "test" {
  name                   = "0.0.1"
  gallery_application_id = azurerm_gallery_application.test.id
  location               = azurerm_gallery_application.test.location

  source {
    media_link = azurerm_storage_blob.test.id
  }

  manage_action {
    install = "[install command]"
    remove  = "[remove command]"
  }

  target_region {
    name                   = azurerm_gallery_application.test.location
    regional_replica_count = 1
    storage_account_type   = "Premium_LRS"
  }
}
`, r.template(data), data.RandomInteger, data.RandomString)
}
