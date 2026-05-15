// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package compute_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
)

// --- ip_configurations: 64 elements ---

func TestAccComputeFleet_validateIPConfigurations64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{Config: r.validateIPConfigurations64ElementsConfig(data)},
	})
}

func (r ComputeFleetTestResource) validateIPConfigurations64ElementsConfig(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

resource "azurerm_virtual_network" "test" {
  name                = "acctest-vnet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  address_space       = ["10.0.0.0/8"]
}

resource "azurerm_subnet" "test" {
  name                 = "acctest-subnet-%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefixes     = ["10.0.0.0/16"]
}

resource "azurerm_compute_fleet" "test" {
  name                = "acctest-fleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  vm_sizes_profile {
    name = "Standard_F1als_v7"
  }

  create_option = "FromImage"
  os_type       = "Linux"

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
      name      = "acctest-ip"
      subnet_id = azurerm_subnet.test.id
      primary   = true
    }

    dynamic "ip_configurations" {
      for_each = range(127)
      content {
        name      = "acctest-ip-${ip_configurations.key}"
        subnet_id = azurerm_subnet.test.id
      }
    }

    dynamic "ip_configurations" {
      for_each = range(128)
      content {
        name                      = "acctest-ipv6-${ip_configurations.key}"
        subnet_id                 = azurerm_subnet.test.id
        public_ip_address_version = "IPv6"
      }
    }
  }

  network_api_version = "2020-11-01"

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- provision_after_extensions: 64 elements ---

func TestAccComputeFleet_validateProvisionAfterExtensions64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{Config: r.validateProvisionAfterExtensions64ElementsConfig(data)},
	})
}

func (r ComputeFleetTestResource) validateProvisionAfterExtensions64ElementsConfig(data acceptance.TestData) string {
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
  os_type       = "Linux"

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

  extensions {
    name                       = "HealthExtension"
    publisher                  = "Microsoft.ManagedServices"
    type                       = "ApplicationHealthWindows"
    type_handler_version       = "1.0"
    auto_upgrade_minor_version = true
    provision_after_extensions = [for i in range(64) : "ext-${i}"]

    settings = jsonencode({
      "protocol"    = "http"
      "port"        = 80
      "requestPath" = "/healthEndpoint"
    })
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- additional_unattend_content: 64 elements ---

func TestAccComputeFleet_validateAdditionalUnattendContent64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{Config: r.validateAdditionalUnattendContent64ElementsConfig(data)},
	})
}

func (r ComputeFleetTestResource) validateAdditionalUnattendContent64ElementsConfig(data acceptance.TestData) string {
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
  os_type       = "Windows"

  image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2022-datacenter-azure-edition-core"
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

  dynamic "additional_unattend_content" {
    for_each = range(64)
    content {
      setting_name = "FirstLogonCommands-${additional_unattend_content.key}"
      content      = "<FirstLogonCommands><SynchronousCommand><CommandLine>echo ${additional_unattend_content.key}</CommandLine><Description>cmd-${additional_unattend_content.key}</Description><Order>${additional_unattend_content.key + 1}</Order></SynchronousCommand></FirstLogonCommands>"
    }
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- listeners: 64 elements ---

func TestAccComputeFleet_validateListeners64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{Config: r.validateListeners64ElementsConfig(data)},
	})
}

func (r ComputeFleetTestResource) validateListeners64ElementsConfig(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {
    key_vault {
      purge_soft_delete_on_destroy               = false
      purge_soft_deleted_certificates_on_destroy = false
    }
  }
}

data "azurerm_client_config" "current" {}

resource "azurerm_key_vault" "test" {
  name                = "acctestkeyvault%s"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  tenant_id           = data.azurerm_client_config.current.tenant_id

  sku_name                        = "standard"
  enabled_for_template_deployment = true
  enabled_for_deployment          = true

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    certificate_permissions = [
      "Create",
      "Delete",
      "Get",
      "Update",
    ]

    key_permissions = ["Get", "Create", "Delete", "List", "Restore", "Recover", "UnwrapKey", "WrapKey", "Purge", "Encrypt", "Decrypt", "Sign", "Verify", "GetRotationPolicy"]

    secret_permissions = [
      "Set",
    ]

    storage_permissions = [
      "Set",
    ]
  }
}

resource "azurerm_key_vault_certificate" "first" {
  count        = 2
  name         = "first-${count.index}"
  key_vault_id = azurerm_key_vault.test.id

  certificate_policy {
    issuer_parameters {
      name = "Self"
    }

    key_properties {
      exportable = true
      key_size   = 2048
      key_type   = "RSA"
      reuse_key  = true
    }

    lifetime_action {
      action {
        action_type = "AutoRenew"
      }

      trigger {
        days_before_expiry = 30
      }
    }

    secret_properties {
      content_type = "application/x-pkcs12"
    }

    x509_certificate_properties {
      key_usage = [
        "cRLSign",
        "dataEncipherment",
        "digitalSignature",
        "keyAgreement",
        "keyCertSign",
        "keyEncipherment",
      ]

      subject            = "CN=hello-world-first"
      validity_in_months = 12
    }
  }
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
  os_type       = "Windows"

  image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2022-datacenter-azure-edition-core"
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
  enable_hotpatching  = true

  secrets {
    source_vault_id = azurerm_key_vault.test.id

    dynamic "vault_certificates" {
      for_each = azurerm_key_vault_certificate.first
      content {
        certificate_url = vault_certificates.value.secret_id
      }
    }
  }

  dynamic "listeners" {
    for_each = range(2)
    content {
      protocol        = "Https"
      certificate_url = azurerm_key_vault_certificate.first[listeners.key].secret_id
    }
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomStringOfLength(6), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- secrets: 64 elements ---

func TestAccComputeFleet_validateSecrets64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{Config: r.validateSecrets64ElementsConfig(data)},
	})
}

func (r ComputeFleetTestResource) validateSecrets64ElementsConfig(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {
    key_vault {
      purge_soft_delete_on_destroy               = false
      purge_soft_deleted_certificates_on_destroy = false
    }
  }
}

data "azurerm_client_config" "current" {}

resource "azurerm_key_vault" "test" {
  name                = "acctestkeyvault%s"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  tenant_id           = data.azurerm_client_config.current.tenant_id

  sku_name                        = "standard"
  enabled_for_template_deployment = true
  enabled_for_deployment          = true

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    certificate_permissions = [
      "Create",
      "Delete",
      "Get",
      "Update",
    ]

    key_permissions = ["Get", "Create", "Delete", "List", "Restore", "Recover", "UnwrapKey", "WrapKey", "Purge", "Encrypt", "Decrypt", "Sign", "Verify", "GetRotationPolicy"]

    secret_permissions = [
      "Set",
    ]

    storage_permissions = [
      "Set",
    ]
  }
}

resource "azurerm_key_vault_certificate" "first" {
  count        = 64
  name         = "first-${count.index}"
  key_vault_id = azurerm_key_vault.test.id

  certificate_policy {
    issuer_parameters {
      name = "Self"
    }

    key_properties {
      exportable = true
      key_size   = 2048
      key_type   = "RSA"
      reuse_key  = true
    }

    lifetime_action {
      action {
        action_type = "AutoRenew"
      }

      trigger {
        days_before_expiry = 30
      }
    }

    secret_properties {
      content_type = "application/x-pkcs12"
    }

    x509_certificate_properties {
      key_usage = [
        "cRLSign",
        "dataEncipherment",
        "digitalSignature",
        "keyAgreement",
        "keyCertSign",
        "keyEncipherment",
      ]

      subject            = "CN=hello-world-first"
      validity_in_months = 12
    }
  }
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
  os_type       = "Windows"

  image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2022-datacenter-azure-edition-core"
    version   = "latest"
  }

  admin_username = "adminuser"
  admin_password = "P@ssw0rd1234!"

  computer_name_prefix = "acctest"
  enable_hotpatching   = true

  network_interface_configurations {
    name = "acctest-nic-%d"
    ip_configurations {
      name      = "acctest-ip-%d"
      subnet_id = azurerm_subnet.test.id
    }
  }

  network_api_version = "2020-11-01"

  dynamic "secrets" {
    for_each = range(64)
    content {
      source_vault_id = azurerm_key_vault.test.id

      vault_certificates {
        certificate_url = azurerm_key_vault_certificate.first[secrets.key].secret_id
      }
    }
  }

  extensions {
    name                       = "HealthExtension"
    publisher                  = "Microsoft.ManagedServices"
    type                       = "ApplicationHealthWindows"
    type_handler_version       = "1.0"
    auto_upgrade_minor_version = true

    settings = jsonencode({
      "protocol"    = "http"
      "port"        = 80
      "requestPath" = "/healthEndpoint"
    })
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomStringOfLength(6), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- vault_certificates: 64 elements ---

func TestAccComputeFleet_validateVaultCertificates64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{Config: r.validateVaultCertificates64ElementsConfig(data)},
	})
}

func (r ComputeFleetTestResource) validateVaultCertificates64ElementsConfig(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {
    key_vault {
      purge_soft_delete_on_destroy               = false
      purge_soft_deleted_certificates_on_destroy = false
    }
  }
}

data "azurerm_client_config" "current" {}

resource "azurerm_key_vault" "test" {
  name                = "acctestkeyvault%s"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  tenant_id           = data.azurerm_client_config.current.tenant_id

  sku_name                        = "standard"
  enabled_for_template_deployment = true
  enabled_for_deployment          = true

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    certificate_permissions = [
      "Create",
      "Delete",
      "Get",
      "Update",
    ]

    key_permissions = ["Get", "Create", "Delete", "List", "Restore", "Recover", "UnwrapKey", "WrapKey", "Purge", "Encrypt", "Decrypt", "Sign", "Verify", "GetRotationPolicy"]

    secret_permissions = [
      "Set",
    ]

    storage_permissions = [
      "Set",
    ]
  }
}

resource "azurerm_key_vault_certificate" "first" {
  count        = 512
  name         = "first-${count.index}"
  key_vault_id = azurerm_key_vault.test.id

  certificate_policy {
    issuer_parameters {
      name = "Self"
    }

    key_properties {
      exportable = true
      key_size   = 2048
      key_type   = "RSA"
      reuse_key  = true
    }

    lifetime_action {
      action {
        action_type = "AutoRenew"
      }

      trigger {
        days_before_expiry = 30
      }
    }

    secret_properties {
      content_type = "application/x-pkcs12"
    }

    x509_certificate_properties {
      key_usage = [
        "cRLSign",
        "dataEncipherment",
        "digitalSignature",
        "keyAgreement",
        "keyCertSign",
        "keyEncipherment",
      ]

      subject            = "CN=hello-world-first"
      validity_in_months = 12
    }
  }
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
  os_type       = "Windows"

  image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2022-datacenter-azure-edition-core"
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
  enable_hotpatching  = true

  secrets {
    source_vault_id = azurerm_key_vault.test.id

    dynamic "vault_certificates" {
      for_each = azurerm_key_vault_certificate.first
      content {
        certificate_store = "store-${vault_certificates.key}"
        certificate_url   = vault_certificates.value.secret_id
      }
    }
  }

  extensions {
    name                       = "HealthExtension"
    publisher                  = "Microsoft.ManagedServices"
    type                       = "ApplicationHealthWindows"
    type_handler_version       = "1.0"
    auto_upgrade_minor_version = true

    settings = jsonencode({
      "protocol"    = "http"
      "port"        = 80
      "requestPath" = "/healthEndpoint"
    })
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomStringOfLength(6), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- extensions: 64 elements ---

func TestAccComputeFleet_validateExtensions64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{Config: r.validateExtensions64ElementsConfig(data)},
	})
}

func (r ComputeFleetTestResource) validateExtensions64ElementsConfig(data acceptance.TestData) string {
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
  os_type       = "Linux"

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

  dynamic "extensions" {
    for_each = range(64)
    content {
      name = "acctest-extension-${extensions.key}"
    }
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- key_data: 64 elements ---

func TestAccComputeFleet_validateKeyData64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{Config: r.validateKeyData64ElementsConfig(data)},
	})
}

func (r ComputeFleetTestResource) validateKeyData64ElementsConfig(data acceptance.TestData) string {
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
  os_type       = "Linux"

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

  key_data = [for i in range(64) : "ssh-key-${i}"]

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

// --- data_disks: 64 elements ---

func TestAccComputeFleet_validateDataDisks64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_compute_fleet", "test")
	r := ComputeFleetTestResource{}
	data.ResourceTest(t, r, []acceptance.TestStep{
		{Config: r.validateDataDisks64ElementsConfig(data)},
	})
}

func (r ComputeFleetTestResource) validateDataDisks64ElementsConfig(data acceptance.TestData) string {
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
  os_type       = "Linux"

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

  dynamic "data_disks" {
    for_each = range(64)
    content {
      create_option = "Empty"
      lun           = data_disks.key
      disk_size_gb  = 10
    }
  }

  regular_priority_profile {
    capacity = 1
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}
