// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package compute_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

func TestAccComputeFleet_updateTags(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.tagsAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.tagsFirst(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.tagsSecond(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.tagsFirst(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.tagsAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) tagsAbsent(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) tagsFirst(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }

  tags = {
    environment = "test"
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) tagsSecond(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }

  tags = {
    environment = "production"
    owner       = "team"
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func TestAccComputeFleet_updateVmSizesProfile(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.vmSizesProfileOneElement(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.vmSizesProfileTwoElements(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.vmSizesProfileOneElement(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func TestAccComputeFleet_updateSpotPriorityProfile(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.spotPriorityProfileAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.spotPriorityProfilePresent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func TestAccComputeFleet_updateIdentity(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.identityAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.identitySystemAssigned(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.identityUserAssigned(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.identitySystemAssigned(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.identityAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) identityAbsent(data acceptance.TestData) string {
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

resource "azurerm_user_assigned_identity" "test2" {
  name                = "acctest-uai2-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) identitySystemAssigned(data acceptance.TestData) string {
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

resource "azurerm_user_assigned_identity" "test2" {
  name                = "acctest-uai2-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }

  identity {
    type = "UserAssigned"
    identity_ids = [
      azurerm_user_assigned_identity.test.id
    ]
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) identityUserAssigned(data acceptance.TestData) string {
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

resource "azurerm_user_assigned_identity" "test2" {
  name                = "acctest-uai2-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }

  identity {
    type         = "UserAssigned"
    identity_ids = [
      azurerm_user_assigned_identity.test.id,
      azurerm_user_assigned_identity.test2.id
    ]
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) spotPriorityProfileAbsent(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) spotPriorityProfilePresent(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }

  spot_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) vmSizesProfileOneElement(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func TestAccComputeFleet_updateRegularPriorityProfile(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.regularPriorityProfileAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.regularPriorityProfilePresent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) regularPriorityProfileAbsent(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

func (r ComputeFleetTestResource) regularPriorityProfilePresent(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }

  spot_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func TestAccComputeFleet_updateVmSizesProfileName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.vmSizesProfileNameFirst(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.vmSizesProfileNameSecond(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.vmSizesProfileNameFirst(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func TestAccComputeFleet_updateVmSizesProfileRank(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.vmSizesProfileRankAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.vmSizesProfileRankFirst(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.vmSizesProfileRankSecond(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.vmSizesProfileRankFirst(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.vmSizesProfileRankAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func TestAccComputeFleet_updateRegularPriorityProfileCapacity(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.regularPriorityProfileCapacityFirst(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.regularPriorityProfileCapacitySecond(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.regularPriorityProfileCapacityFirst(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func TestAccComputeFleet_updateSpotPriorityProfileCapacity(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.spotPriorityProfileCapacityFirst(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.spotPriorityProfileCapacitySecond(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.spotPriorityProfileCapacityFirst(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) spotPriorityProfileCapacityFirst(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

func (r ComputeFleetTestResource) spotPriorityProfileCapacitySecond(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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
    capacity = 2
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) regularPriorityProfileCapacityFirst(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) regularPriorityProfileCapacitySecond(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 2
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) vmSizesProfileRankAbsent(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
    allocation_strategy = "Prioritized"
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) vmSizesProfileRankFirst(data acceptance.TestData) string {
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
    rank = 0
  }

  create_option = "FromImage"

  os_type = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
    allocation_strategy = "Prioritized"
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) vmSizesProfileRankSecond(data acceptance.TestData) string {
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
    rank = 1
  }

  vm_sizes_profile {
    name = "Standard_D4s_v3"
    rank = 0
  }

  create_option = "FromImage"

  os_type = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
    allocation_strategy = "Prioritized"
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) vmSizesProfileNameFirst(data acceptance.TestData) string {
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

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) vmSizesProfileNameSecond(data acceptance.TestData) string {
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
    name = "Standard_D4s_v3"
  }

  create_option = "FromImage"

  os_type = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r ComputeFleetTestResource) vmSizesProfileTwoElements(data acceptance.TestData) string {
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

  vm_sizes_profile {
    name = "Standard_D4s_v3"
  }

  create_option = "FromImage"

  os_type = "Linux"

  publisher = "Canonical"
  offer     = "0001-com-ubuntu-server-jammy"
  sku       = "22_04-lts-gen2"
  version   = "latest"

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

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}
