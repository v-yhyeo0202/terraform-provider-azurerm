// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package compute_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

func TestAccComputeFleet_validateFleetName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateFleetName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateResourceGroupName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateResourceGroupName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateLocation(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateLocation(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateOffer(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateOffer(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validatePublisher(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validatePublisher(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSku(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSku(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateVersion(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateVersion(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateAdminPassword(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateAdminPassword(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateAdminUsername(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateAdminUsername(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateVmSizeProfileName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateVmSizeProfileName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateNICName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateNICName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSubnet(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSubnet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validatePackageReferenceId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validatePackageReferenceId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateConfigurationReference(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateConfigurationReference(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSpotPriorityProfileCapacityNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSpotPriorityProfileCapacity(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSpotPriorityProfileCapacityZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSpotPriorityProfileCapacity(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSpotPriorityProfileCapacityLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSpotPriorityProfileCapacity(data, 4294967296),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateRegularPriorityProfileCapacityNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateRegularPriorityProfileCapacity(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateRegularPriorityProfileCapacityZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateRegularPriorityProfileCapacity(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateRegularPriorityProfileCapacityLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateRegularPriorityProfileCapacity(data, 4294967296),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateGalleryApplicationOrderNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateGalleryApplicationOrder(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateGalleryApplicationOrderZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateGalleryApplicationOrder(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateGalleryApplicationOrderLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateGalleryApplicationOrder(data, 4294967296),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r ComputeFleetTestResource) validateFleetName(data acceptance.TestData) string {
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
  name                = "🙂"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  vm_sizes_profile {
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateResourceGroupName(data acceptance.TestData) string {
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
  resource_group_name = "🙂"
  location            = azurerm_resource_group.test.location

  vm_sizes_profile {
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateLocation(data acceptance.TestData) string {
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
  location            = "🙂"

  vm_sizes_profile {
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateOffer(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "🙂"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validatePublisher(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "🙂"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateSku(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "🙂"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateVersion(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "🙂"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateAdminPassword(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "🙂"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateAdminUsername(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "🙂"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateVmSizeProfileName(data acceptance.TestData) string {
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
    name = "🙂"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateNICName(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "🙂"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateIPConfigName(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "🙂"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateSubnet(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  vm_sizes_profile {
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = "🙂"
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validatePackageReferenceId(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
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
    package_reference_id = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateConfigurationReference(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
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
    configuration_reference = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateSpotPriorityProfileCapacity(data acceptance.TestData, capacity int) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
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
    capacity = %d
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, capacity)
}

func (r ComputeFleetTestResource) validateRegularPriorityProfileCapacity(data acceptance.TestData, capacity int) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = %d
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, capacity)
}

func (r ComputeFleetTestResource) validateGalleryApplicationOrder(data acceptance.TestData, order int) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
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
    order = %d
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, order)
}

func TestAccComputeFleet_validateGalleryApplicationTags(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateGalleryApplicationTags(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateExtensionName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateExtensionPublisher(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionPublisher(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateExtensionType(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionType(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateExtensionTypeHandlerVersion(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionTypeHandlerVersion(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateExtensionForceUpdateTag(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionForceUpdateTag(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateExtensionProtectedSettings(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionProtectedSettings(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateExtensionSecretUrl(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionSecretUrl(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateExtensionSourceVaultId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionSourceVaultId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateExtensionProvisionAfterExtensions(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionProvisionAfterExtensions(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateExtensionProvisionAfterExtensions64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionProvisionAfterExtensions64Elements(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateExtensionSettings(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionSettings(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigPrivateIPAddressVersion(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigPrivateIPAddressVersion(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r ComputeFleetTestResource) validateGalleryApplicationTags(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
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
    gallery_application_tags = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateExtensionName(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  extensions {
    name = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateExtensionPublisher(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  extensions {
    publisher = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateExtensionType(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  extensions {
    type = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateExtensionTypeHandlerVersion(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  extensions {
    type_handler_version = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateExtensionForceUpdateTag(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  extensions {
    force_update_tag = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateExtensionProtectedSettings(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  extensions {
    protected_settings = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateExtensionSecretUrl(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  extensions {
    secret_url = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateExtensionSourceVaultId(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  extensions {
    source_vault_id = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateExtensionProvisionAfterExtensions(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  extensions {
    provision_after_extensions = ["🙂"]
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateExtensionProvisionAfterExtensions64Elements(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

locals {
  provision_after = [for i in range(64) : "ext-${i}"]
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  extensions {
    provision_after_extensions = local.provision_after
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateExtensionSettings(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  extensions {
    settings = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateIPConfigPrivateIPAddressVersion(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name                     = "acctest-ip-%d"
      subnet_id                = azurerm_subnet.test.id
      private_ip_address_version = "🙂"
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func TestAccComputeFleet_validateNICAuxiliaryMode(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateNICAuxiliaryMode(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateNICAuxiliarySku(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateNICAuxiliarySku(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateNICDeleteOption(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateNICDeleteOption(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateNICDnsServers(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateNICDnsServers(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateNICDnsServers64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateNICDnsServers64Elements(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateNICNetworkSecurityGroupId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateNICNetworkSecurityGroupId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigDeleteOption(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigDeleteOption(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigPublicIPAddressConfigurationName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigPublicIPAddressConfigurationName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigDomainNameLabel(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigDomainNameLabel(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigDomainNameLabelScope(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigDomainNameLabelScope(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigIdleTimeoutInMinutesNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigIdleTimeoutInMinutes(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigIdleTimeoutInMinutesZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigIdleTimeoutInMinutes(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigIdleTimeoutInMinutesLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigIdleTimeoutInMinutes(data, 4294967296),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigPublicIPAddressVersion(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigPublicIPAddressVersion(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigSkuName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigSkuName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIPConfigSkuTier(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIPConfigSkuTier(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateCustomData(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateCustomData(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateComputerNamePrefix(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateComputerNamePrefix(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r ComputeFleetTestResource) validateNICAuxiliaryMode(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name           = "acctest-nic-%d"
    auxiliary_mode = "🙂"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateNICAuxiliarySku(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name          = "acctest-nic-%d"
    auxiliary_sku = "🙂"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateNICDeleteOption(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name          = "acctest-nic-%d"
    delete_option = "🙂"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateNICDnsServers(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name       = "acctest-nic-%d"
    dns_servers = ["🙂"]
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateNICDnsServers64Elements(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

locals {
  dns_servers = [for i in range(64) : "10.0.${i}.1"]
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name        = "acctest-nic-%d"
    dns_servers = local.dns_servers
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateNICNetworkSecurityGroupId(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name                    = "acctest-nic-%d"
    network_security_group_id = "🙂"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateIPConfigDeleteOption(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name          = "acctest-ip-%d"
      subnet_id     = azurerm_subnet.test.id
      delete_option = "🙂"
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateIPConfigPublicIPAddressConfigurationName(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name                                 = "acctest-ip-%d"
      subnet_id                            = azurerm_subnet.test.id
      public_ip_address_configuration_name = "🙂"
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateIPConfigDomainNameLabel(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name              = "acctest-ip-%d"
      subnet_id         = azurerm_subnet.test.id
      domain_name_label = "🙂"
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateIPConfigDomainNameLabelScope(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name                   = "acctest-ip-%d"
      subnet_id              = azurerm_subnet.test.id
      domain_name_label_scope = "🙂"
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateIPConfigIdleTimeoutInMinutes(data acceptance.TestData, timeout int) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name                   = "acctest-ip-%d"
      subnet_id              = azurerm_subnet.test.id
      idle_timeout_in_minutes = %d
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, timeout)
}

func (r ComputeFleetTestResource) validateIPConfigPublicIPAddressVersion(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name                    = "acctest-ip-%d"
      subnet_id               = azurerm_subnet.test.id
      public_ip_address_version = "🙂"
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateIPConfigSkuName(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
      sku_name  = "🙂"
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateIPConfigSkuTier(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
      sku_tier  = "🙂"
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateCustomData(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"
  custom_data          = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateComputerNamePrefix(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func TestAccComputeFleet_validateLinuxPatchMode(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateLinuxPatchMode(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateLinuxRebootSetting(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateLinuxRebootSetting(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateKeyData(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateKeyData(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateKeyData64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateKeyData64Elements(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSecretsSourceVaultId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSecretsSourceVaultId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSecretsCertificateUrl(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSecretsCertificateUrl(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSecretsCertificateUrl64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSecretsCertificateUrl64Elements(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateWindowsPatchMode(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateWindowsPatchMode(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateWindowsRebootSetting(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateWindowsRebootSetting(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateTimeZone(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateTimeZone(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateAdditionalUnattendContentSettingName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateAdditionalUnattendContentSettingName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateAdditionalUnattendContentContent(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateAdditionalUnattendContentContent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateListenersCertificateUrl(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateListenersCertificateUrl(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateListenersProtocol(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateListenersProtocol(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDataDisksCaching(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataDisksCaching(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r ComputeFleetTestResource) validateLinuxPatchMode(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"
  patch_mode           = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateLinuxRebootSetting(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"
  reboot_setting       = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateKeyData(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"
  key_data             = ["🙂"]

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateKeyData64Elements(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

locals {
  key_data = [for i in range(64) : "ssh-rsa key-${i}"]
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"
  key_data             = local.key_data

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateSecretsSourceVaultId(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  secrets {
    source_vault_id = "🙂"
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateSecretsCertificateUrl(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  secrets {
    certificate_url = ["🙂"]
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateSecretsCertificateUrl64Elements(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

locals {
  certificate_urls = [for i in range(64) : "https://vault.example.com/cert-${i}"]
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  secrets {
    certificate_url = local.certificate_urls
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateWindowsPatchMode(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Windows"

  publisher = "MicrosoftWindowsServer"
  offer     = "WindowsServer"
  sku       = "2019-Datacenter"
  version   = "latest"

  admin_username                   = "adminuser"
  admin_password                   = "P@ssw0rd1234!"
  computer_name_prefix             = "acctest"
  windows_configuration_patch_mode = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateWindowsRebootSetting(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Windows"

  publisher = "MicrosoftWindowsServer"
  offer     = "WindowsServer"
  sku       = "2019-Datacenter"
  version   = "latest"

  admin_username                       = "adminuser"
  admin_password                       = "P@ssw0rd1234!"
  computer_name_prefix                 = "acctest"
  windows_configuration_reboot_setting = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateTimeZone(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Windows"

  publisher = "MicrosoftWindowsServer"
  offer     = "WindowsServer"
  sku       = "2019-Datacenter"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"
  time_zone            = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateAdditionalUnattendContentSettingName(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Windows"

  publisher = "MicrosoftWindowsServer"
  offer     = "WindowsServer"
  sku       = "2019-Datacenter"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  additional_unattend_content {
    setting_name = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateAdditionalUnattendContentContent(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Windows"

  publisher = "MicrosoftWindowsServer"
  offer     = "WindowsServer"
  sku       = "2019-Datacenter"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  additional_unattend_content {
    content = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateListenersCertificateUrl(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Windows"

  publisher = "MicrosoftWindowsServer"
  offer     = "WindowsServer"
  sku       = "2019-Datacenter"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  listeners {
    certificate_url = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateListenersProtocol(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Windows"

  publisher = "MicrosoftWindowsServer"
  offer     = "WindowsServer"
  sku       = "2019-Datacenter"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  listeners {
    protocol = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateDataDisksCaching(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  data_disks {
    caching = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func TestAccComputeFleet_validateDataDisksCreateOption(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataDisksCreateOption(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDataDisksDeleteOption(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataDisksDeleteOption(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDataDisksDiskSizeGBNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataDisksDiskSizeGB(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDataDisksDiskSizeGBZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataDisksDiskSizeGB(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDataDisksDiskSizeGBLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataDisksDiskSizeGB(data, 4294967296),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDataDisksDiskEncryptionSetId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataDisksDiskEncryptionSetId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDataDisksStorageAccountType(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataDisksStorageAccountType(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDataDisksLunNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataDisksLun(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDataDisksLunZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataDisksLun(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDataDisksLunLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataDisksLun(data, 4294967296),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateOsDiskCaching(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateOsDiskCaching(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateOsDiskDeleteOption(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateOsDiskDeleteOption(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDiffDiskSettingsOption(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDiffDiskSettingsOption(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateDiffDiskSettingsPlacement(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDiffDiskSettingsPlacement(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateOsDiskSizeGBNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateOsDiskSizeGB(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateOsDiskSizeGBZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateOsDiskSizeGB(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateOsDiskSizeGBLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateOsDiskSizeGB(data, 4294967296),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateOsDiskDiskEncryptionSetId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateOsDiskDiskEncryptionSetId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateOsDiskSecurityEncryptionType(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateOsDiskSecurityEncryptionType(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateOsDiskStorageAccountType(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateOsDiskStorageAccountType(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validatePlanName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validatePlanName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validatePlanPublisher(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validatePlanPublisher(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validatePlanProduct(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validatePlanProduct(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validatePlanPromotionCode(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validatePlanPromotionCode(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r ComputeFleetTestResource) validateDataDisksCreateOption(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  data_disks {
    create_option = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateDataDisksDeleteOption(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  data_disks {
    delete_option = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateDataDisksDiskSizeGB(data acceptance.TestData, diskSizeGB int) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  data_disks {
    disk_size_gb = %d
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, diskSizeGB, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateDataDisksDiskEncryptionSetId(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  data_disks {
    disk_encryption_set_id = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateDataDisksStorageAccountType(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  data_disks {
    storage_account_type = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateDataDisksLun(data acceptance.TestData, lun int) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  data_disks {
    lun = %d
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, lun, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateOsDiskCaching(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"
  os_disk_caching      = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateOsDiskDeleteOption(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username         = "adminuser"
  admin_password         = "P@ssw0rd1234!"
  computer_name_prefix   = "acctest"
  os_disk_delete_option  = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateDiffDiskSettingsOption(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"
  option               = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateDiffDiskSettingsPlacement(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"
  placement            = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateOsDiskSizeGB(data acceptance.TestData, diskSizeGB int) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"
  os_disk_size_gb      = %d

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, diskSizeGB, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateOsDiskDiskEncryptionSetId(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username                  = "adminuser"
  admin_password                  = "P@ssw0rd1234!"
  computer_name_prefix            = "acctest"
  os_disk_disk_encryption_set_id  = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateOsDiskSecurityEncryptionType(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username                    = "adminuser"
  admin_password                    = "P@ssw0rd1234!"
  computer_name_prefix              = "acctest"
  os_disk_security_encryption_type  = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateOsDiskStorageAccountType(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username               = "adminuser"
  admin_password               = "P@ssw0rd1234!"
  computer_name_prefix         = "acctest"
  os_disk_storage_account_type = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validatePlanName(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  plan {
    name      = "🙂"
    publisher = "acctest-publisher"
    product   = "acctest-product"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validatePlanPublisher(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  plan {
    name      = "acctest-plan"
    publisher = "🙂"
    product   = "acctest-product"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validatePlanProduct(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  plan {
    name      = "acctest-plan"
    publisher = "acctest-publisher"
    product   = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validatePlanPromotionCode(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  plan {
    name           = "acctest-plan"
    publisher      = "acctest-publisher"
    product        = "acctest-product"
    promotion_code = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func TestAccComputeFleet_validateRegularPriorityProfileMinCapacityNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateRegularPriorityProfileMinCapacity(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateRegularPriorityProfileMinCapacityZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateRegularPriorityProfileMinCapacity(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateRegularPriorityProfileMinCapacityLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateRegularPriorityProfileMinCapacity(data, 4294967296),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateRegularPriorityProfileAllocationStrategy(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateRegularPriorityProfileAllocationStrategy(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSpotPriorityProfileMinCapacityNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSpotPriorityProfileMinCapacity(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSpotPriorityProfileMinCapacityZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSpotPriorityProfileMinCapacity(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSpotPriorityProfileMinCapacityLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSpotPriorityProfileMinCapacity(data, 4294967296),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSpotPriorityProfileMaxPricePerVMNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSpotPriorityProfileMaxPricePerVM(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSpotPriorityProfileMaxPricePerVMZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSpotPriorityProfileMaxPricePerVM(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSpotPriorityProfileMaxPricePerVMLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSpotPriorityProfileMaxPricePerVM(data, 999999),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSpotPriorityProfileEvictionPolicy(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSpotPriorityProfileEvictionPolicy(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSpotPriorityProfileAllocationStrategy(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSpotPriorityProfileAllocationStrategy(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateVmSizeProfileRankNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateVmSizeProfileRank(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateVmSizeProfileRankZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateVmSizeProfileRank(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateVmSizeProfileRankLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateVmSizeProfileRank(data, 4294967296),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateNetworkApiVersion(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateNetworkApiVersion(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateIdentityType(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateIdentityType(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateUserAssignedIdentities(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateUserAssignedIdentities(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validatePlatformFaultDomainCountNegativeOne(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validatePlatformFaultDomainCount(data, -1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validatePlatformFaultDomainCountZero(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validatePlatformFaultDomainCount(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validatePlatformFaultDomainCountLarge(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validatePlatformFaultDomainCount(data, 4294967296),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateComputeApiVersion(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateComputeApiVersion(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateStorageUri(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateStorageUri(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateCapacityReservationGroupId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateCapacityReservationGroupId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r ComputeFleetTestResource) validateRegularPriorityProfileMinCapacity(data acceptance.TestData, minCapacity int) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity     = 1
    min_capacity = %d
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, minCapacity)
}

func (r ComputeFleetTestResource) validateRegularPriorityProfileAllocationStrategy(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity            = 1
    allocation_strategy = "🙂"
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateSpotPriorityProfileMinCapacity(data acceptance.TestData, minCapacity int) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
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
    capacity     = 1
    min_capacity = %d
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, minCapacity)
}

func (r ComputeFleetTestResource) validateSpotPriorityProfileMaxPricePerVM(data acceptance.TestData, maxPrice float64) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
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
    capacity         = 1
    max_price_per_vm = %g
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, maxPrice)
}

func (r ComputeFleetTestResource) validateSpotPriorityProfileEvictionPolicy(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
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
    capacity        = 1
    eviction_policy = "🙂"
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateSpotPriorityProfileAllocationStrategy(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
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
    capacity            = 1
    allocation_strategy = "🙂"
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateVmSizeProfileRank(data acceptance.TestData, rank int) string {
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
    name = "Standard_D2s_v3"
    rank = %d
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, rank, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateNetworkApiVersion(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "🙂"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateIdentityType(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  identity {
    type = "🙂"
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateUserAssignedIdentities(data acceptance.TestData) string {
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

resource "azurerm_user_assigned_identity" "test" {
  name                = "acctest-uai-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  vm_sizes_profile {
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  identity {
    type         = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.test.id]
  }

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validatePlatformFaultDomainCount(data acceptance.TestData, count int) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username              = "adminuser"
  admin_password              = "P@ssw0rd1234!"
  computer_name_prefix        = "acctest"
  platform_fault_domain_count = %d

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, count, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateComputeApiVersion(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option       = "FromImage"
  os_type             = "Linux"
  compute_api_version = "🙂"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateStorageUri(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"
  storage_uri          = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateCapacityReservationGroupId(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username                = "adminuser"
  admin_password                = "P@ssw0rd1234!"
  computer_name_prefix          = "acctest"
  capacity_reservation_group_id = "🙂"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func TestAccComputeFleet_validateExtensionsTimeBudget(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateExtensionsTimeBudget(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateLicenseType(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateLicenseType(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateOsImageNotificationProfileNotBeforeTimeout(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateOsImageNotificationProfileNotBeforeTimeout(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateTerminateNotificationProfileNotBeforeTimeout(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateTerminateNotificationProfileNotBeforeTimeout(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateImageReferenceId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateImageReferenceId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateSharedGalleryImageId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateSharedGalleryImageId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateCommunityGalleryImageId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateCommunityGalleryImageId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateUserData(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateUserData(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateZones(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateZones(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccComputeFleet_validateZones64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateZones64Elements(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r ComputeFleetTestResource) validateExtensionsTimeBudget(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option         = "FromImage"
  os_type               = "Linux"
  extensions_time_budget = "🙂"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateLicenseType(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"
  license_type  = "🙂"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateOsImageNotificationProfileNotBeforeTimeout(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option       = "FromImage"
  os_type             = "Linux"
  not_before_timeout  = "🙂"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateTerminateNotificationProfileNotBeforeTimeout(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"
  terminate_notification_profile_not_before_timeout = "🙂"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateImageReferenceId(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option      = "FromImage"
  os_type            = "Linux"
  image_reference_id = "🙂"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateSharedGalleryImageId(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option           = "FromImage"
  os_type                 = "Linux"
  shared_gallery_image_id = "🙂"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateCommunityGalleryImageId(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option              = "FromImage"
  os_type                    = "Linux"
  community_gallery_image_id = "🙂"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateUserData(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"
  user_data     = "🙂"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateZones(data acceptance.TestData) string {
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"
  zones         = ["🙂"]

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) validateZones64Elements(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

locals {
  zones = [for i in range(64) : "zone-${i}"]
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
    name = "Standard_D2s_v3"
  }

  create_option = "FromImage"
  os_type       = "Linux"
  zones         = local.zones

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

  admin_username       = "adminuser"
  admin_password       = "P@ssw0rd1234!"
  computer_name_prefix = "acctest"

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}
