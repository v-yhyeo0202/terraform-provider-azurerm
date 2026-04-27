// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computefleet_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

func TestAccComputeFleet_updateComputeApiVersion(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.computeApiVersionAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withComputeApiVersion(data, "2024-07-01"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withComputeApiVersion(data, "2023-03-01"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withComputeApiVersion(data, "2024-07-01"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.computeApiVersionAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func TestAccComputeFleet_updateCapacity(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withCapacity(data, 1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withCapacity(data, 0),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withCapacity(data, 1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func TestAccComputeFleet_test(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withCapacity(data, 1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r ComputeFleetTestResource) withCapacity(data acceptance.TestData, capacity int) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = %[3]d
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, capacity)
}

func (r ComputeFleetTestResource) computeApiVersionAbsent(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withComputeApiVersion(data acceptance.TestData, computeApiVersion string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  compute_api_version  = %[3]q
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, computeApiVersion)
}

func TestAccComputeFleet_updateOffer(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withOffer(data, "ubuntu-24_04-lts"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withOffer(data, "ubuntu-22_04-lts"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withOffer(data, "ubuntu-24_04-lts"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) withOffer(data acceptance.TestData, offer string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = %[3]q
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, offer)
}

func TestAccComputeFleet_updatePublisher(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withPublisher(data, "canonical"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withPublisher(data, "Canonical"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withPublisher(data, "canonical"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) withPublisher(data acceptance.TestData, publisher string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = %[3]q
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, publisher)
}

func TestAccComputeFleet_updateSku(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withSku(data, "server"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withSku(data, "server-gen1"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withSku(data, "server"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) withSku(data acceptance.TestData, sku string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = %[3]q
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, sku)
}

func TestAccComputeFleet_updateVersion(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withVersion(data, "latest"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withVersion(data, "1.0.0"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withVersion(data, "latest"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) withVersion(data acceptance.TestData, version string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = %[3]q
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, version)
}

func TestAccComputeFleet_updateAdminUsername(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withAdminUsername(data, "v-yyeo"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withAdminUsername(data, "adminuser"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withAdminUsername(data, "v-yyeo"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) withAdminUsername(data acceptance.TestData, adminUsername string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = %[3]q
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, adminUsername)
}

func TestAccComputeFleet_updateAdminPassword(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.adminPasswordAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withAdminPassword(data, "Password1234!"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withAdminPassword(data, "Password5678!"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withAdminPassword(data, "Password1234!"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.adminPasswordAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) adminPasswordAbsent(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withAdminPassword(data acceptance.TestData, adminPassword string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = %[3]q
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, adminPassword)
}

func TestAccComputeFleet_updateComputerNamePrefix(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withComputerNamePrefix(data, "default-"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withComputerNamePrefix(data, "updated-"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withComputerNamePrefix(data, "default-"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) withComputerNamePrefix(data acceptance.TestData, computerNamePrefix string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = %[3]q
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, computerNamePrefix)
}

func TestAccComputeFleet_updateDisablePasswordAuthentication(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.disablePasswordAuthenticationAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withDisablePasswordAuthentication(data, true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withDisablePasswordAuthentication(data, false),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withDisablePasswordAuthentication(data, true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.disablePasswordAuthenticationAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) disablePasswordAuthenticationAbsent(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withDisablePasswordAuthentication(data acceptance.TestData, disabled bool) string {
	adminPasswordLine := `admin_password = "Password1234!"`
	if disabled {
		adminPasswordLine = ""
	}
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                            = "acctest-fleet-%[2]d"
  resource_group_name             = azurerm_resource_group.test.name
  location                        = azurerm_resource_group.test.location
  capacity                         = 1
  offer                            = "ubuntu-24_04-lts"
  publisher                        = "canonical"
  sku                              = "server"
  version                          = "latest"
  admin_username                   = "v-yyeo"
  %[3]s
  computer_name_prefix             = "default-"
  disable_password_authentication  = %[4]t
  network_api_version              = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, adminPasswordLine, disabled)
}

func TestAccComputeFleet_updateVmSizeProfileName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withVmSizeProfileNames(data, []string{"Standard_F1als_v7"}),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withVmSizeProfileNames(data, []string{"Standard_F1als_v7", "Standard_D2s_v3"}),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withVmSizeProfileNames(data, []string{"Standard_F1als_v7"}),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) withVmSizeProfileNames(data acceptance.TestData, names []string) string {
	blocksFormatted := ""
	for _, name := range names {
		blocksFormatted += fmt.Sprintf("  vm_size_profile {\n    name = %q\n  }\n", name)
	}
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

%[3]s}`, r.template(data), data.RandomInteger, blocksFormatted)
}

func TestAccComputeFleet_updateNetworkInterfaceConfiguration(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withOneNetworkInterfaceConfiguration(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withTwoNetworkInterfaceConfigurations(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withOneNetworkInterfaceConfiguration(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) withOneNetworkInterfaceConfiguration(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withTwoNetworkInterfaceConfigurations(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_subnet" "test2" {
  name                 = "acctsub2-%[2]d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefixes     = ["10.0.3.0/24"]
}

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"
  }

  network_interface_configuration {
    name = "acctest-nic2-%[2]d"
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger)
}

func TestAccComputeFleet_updateNetworkApiVersion(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withNetworkApiVersion(data, "2020-11-01"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withNetworkApiVersion(data, "2022-04-01"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withNetworkApiVersion(data, "2020-11-01"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) withNetworkApiVersion(data acceptance.TestData, networkApiVersion string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = %[3]q

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, networkApiVersion)
}

func TestAccComputeFleet_updateIpConfiguration(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.ipConfigurationAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withOneIpConfiguration(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withTwoIpConfigurations(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withOneIpConfiguration(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.ipConfigurationAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) ipConfigurationAbsent(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withOneIpConfiguration(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withTwoIpConfigurations(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }

    ip_configuration {
      name   = "acctest-ipconfig2-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger)
}

func TestAccComputeFleet_updateIpConfigurationPrimary(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.ipConfigurationPrimaryAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withIpConfigurationPrimary(data, true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withIpConfigurationPrimary(data, false),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withIpConfigurationPrimary(data, true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.ipConfigurationPrimaryAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) ipConfigurationPrimaryAbsent(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withIpConfigurationPrimary(data acceptance.TestData, primary bool) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name    = "acctest-ipconfig-%[2]d"
      subnet  = azurerm_subnet.test.id
      primary = %[3]t
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, primary)
}

func TestAccComputeFleet_updateIpConfigurationPrivateIPAddressVersion(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.ipConfigurationPrivateIPAddressVersionAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withIpConfigurationPrivateIPAddressVersion(data, "IPv4"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withIpConfigurationPrivateIPAddressVersion(data, "IPv6"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withIpConfigurationPrivateIPAddressVersion(data, "IPv4"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.ipConfigurationPrivateIPAddressVersionAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) ipConfigurationPrivateIPAddressVersionAbsent(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withIpConfigurationPrivateIPAddressVersion(data acceptance.TestData, ipVersion string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name                    = "acctest-ipconfig-%[2]d"
      subnet                  = azurerm_subnet.test.id
      private_ip_address_version = %[3]q
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.template(data), data.RandomInteger, ipVersion)
}

func TestAccComputeFleet_updateApplicationSecurityGroupsId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.applicationSecurityGroupsIdAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withOneApplicationSecurityGroupId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withTwoApplicationSecurityGroupIds(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withOneApplicationSecurityGroupId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.applicationSecurityGroupsIdAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) templateWithAsgs(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_application_security_group" "test" {
  name                = "acctest-asg-%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_application_security_group" "test2" {
  name                = "acctest-asg2-%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}
`, r.template(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) applicationSecurityGroupsIdAbsent(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.templateWithAsgs(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withOneApplicationSecurityGroupId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name                         = "acctest-ipconfig-%[2]d"
      subnet                       = azurerm_subnet.test.id
      application_security_groups_id = [azurerm_application_security_group.test.id]
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.templateWithAsgs(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withTwoApplicationSecurityGroupIds(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name                         = "acctest-ipconfig-%[2]d"
      subnet                       = azurerm_subnet.test.id
      application_security_groups_id = [
        azurerm_application_security_group.test.id,
        azurerm_application_security_group.test2.id,
      ]
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.templateWithAsgs(data), data.RandomInteger)
}

func TestAccComputeFleet_updateLoadBalancerBackendAddressPoolsId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.loadBalancerBackendAddressPoolsIdAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withOneLoadBalancerBackendAddressPoolId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withTwoLoadBalancerBackendAddressPoolIds(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withOneLoadBalancerBackendAddressPoolId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.loadBalancerBackendAddressPoolsIdAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) templateWithLoadBalancers(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_public_ip" "test" {
  name                = "acctest-pip-%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  allocation_method   = "Static"
  sku                 = "Standard"
}

resource "azurerm_lb" "test" {
  name                = "acctest-lb-%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  sku                 = "Standard"

  frontend_ip_configuration {
    name                 = "frontend"
    public_ip_address_id = azurerm_public_ip.test.id
  }
}

resource "azurerm_lb_backend_address_pool" "test" {
  loadbalancer_id = azurerm_lb.test.id
  name            = "acctest-lbbap-%[2]d"
}

resource "azurerm_lb_backend_address_pool" "test2" {
  loadbalancer_id = azurerm_lb.test.id
  name            = "acctest-lbbap2-%[2]d"
}
`, r.template(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) loadBalancerBackendAddressPoolsIdAbsent(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.templateWithLoadBalancers(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withOneLoadBalancerBackendAddressPoolId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name                               = "acctest-ipconfig-%[2]d"
      subnet                             = azurerm_subnet.test.id
      load_balancer_backend_address_pools_id = [azurerm_lb_backend_address_pool.test.id]
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.templateWithLoadBalancers(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withTwoLoadBalancerBackendAddressPoolIds(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name                               = "acctest-ipconfig-%[2]d"
      subnet                             = azurerm_subnet.test.id
      load_balancer_backend_address_pools_id = [
        azurerm_lb_backend_address_pool.test.id,
        azurerm_lb_backend_address_pool.test2.id,
      ]
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.templateWithLoadBalancers(data), data.RandomInteger)
}

func TestAccComputeFleet_updateApplicationGatewayBackendAddressPoolsId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.applicationGatewayBackendAddressPoolsIdAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withOneApplicationGatewayBackendAddressPoolId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withTwoApplicationGatewayBackendAddressPoolIds(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.withOneApplicationGatewayBackendAddressPoolId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
		{
			Config: r.applicationGatewayBackendAddressPoolsIdAbsent(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("admin_password"),
	})
}

func (r ComputeFleetTestResource) templateWithApplicationGateway(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_public_ip" "appgw" {
  name                = "acctest-pip-appgw-%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  allocation_method   = "Static"
  sku                 = "Standard"
}

resource "azurerm_application_gateway" "test" {
  name                = "acctest-appgw-%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name

  sku {
    name     = "Standard_v2"
    tier     = "Standard_v2"
    capacity = 1
  }

  gateway_ip_configuration {
    name      = "gateway-ip-config"
    subnet_id = azurerm_subnet.test.id
  }

  frontend_port {
    name = "frontend-port"
    port = 80
  }

  frontend_ip_configuration {
    name                 = "frontend-ip-config"
    public_ip_address_id = azurerm_public_ip.appgw.id
  }

  backend_address_pool {
    name = "acctest-appgw-bap-%[2]d"
  }

  backend_address_pool {
    name = "acctest-appgw-bap2-%[2]d"
  }

  backend_http_settings {
    name                  = "backend-http-settings"
    cookie_based_affinity = "Disabled"
    port                  = 80
    protocol              = "Http"
    request_timeout       = 1
  }

  http_listener {
    name                           = "http-listener"
    frontend_ip_configuration_name = "frontend-ip-config"
    frontend_port_name             = "frontend-port"
    protocol                       = "Http"
  }

  request_routing_rule {
    name                       = "routing-rule"
    rule_type                  = "Basic"
    http_listener_name         = "http-listener"
    backend_address_pool_name  = "acctest-appgw-bap-%[2]d"
    backend_http_settings_name = "backend-http-settings"
    priority                   = 1
  }
}
`, r.template(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) applicationGatewayBackendAddressPoolsIdAbsent(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name   = "acctest-ipconfig-%[2]d"
      subnet = azurerm_subnet.test.id
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.templateWithApplicationGateway(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withOneApplicationGatewayBackendAddressPoolId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name                                        = "acctest-ipconfig-%[2]d"
      subnet                                      = azurerm_subnet.test.id
      application_gateway_backend_address_pools_id = [azurerm_application_gateway.test.backend_address_pool[0].id]
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.templateWithApplicationGateway(data), data.RandomInteger)
}

func (r ComputeFleetTestResource) withTwoApplicationGatewayBackendAddressPoolIds(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  capacity             = 1
  offer                = "ubuntu-24_04-lts"
  publisher            = "canonical"
  sku                  = "server"
  version              = "latest"
  admin_username       = "v-yyeo"
  admin_password       = "Password1234!"
  computer_name_prefix = "default-"
  network_api_version  = "2020-11-01"

  network_interface_configuration {
    name = "acctest-nic-%[2]d"

    ip_configuration {
      name                                        = "acctest-ipconfig-%[2]d"
      subnet                                      = azurerm_subnet.test.id
      application_gateway_backend_address_pools_id = [
        azurerm_application_gateway.test.backend_address_pool[0].id,
        azurerm_application_gateway.test.backend_address_pool[1].id,
      ]
    }
  }

  vm_size_profile {
    name = "Standard_F1als_v7"
  }
}
`, r.templateWithApplicationGateway(data), data.RandomInteger)
}
