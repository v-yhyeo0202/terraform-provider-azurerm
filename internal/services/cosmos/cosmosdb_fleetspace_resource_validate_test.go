// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package cosmos_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
)

// Batch 1

func TestAccCosmosDbFleetspace_validateNameEmoji(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_cosmosdb_fleetspace", "test")
	r := CosmosDbFleetspaceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateNameConfig(data, "🙂"),
		},
	})
}

func TestAccCosmosDbFleetspace_validateFleetNameEmoji(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_cosmosdb_fleetspace", "test")
	r := CosmosDbFleetspaceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateFleetNameConfig(data, "🙂"),
		},
	})
}

func TestAccCosmosDbFleetspace_validateResourceGroupNameEmoji(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_cosmosdb_fleetspace", "test")
	r := CosmosDbFleetspaceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateResourceGroupNameConfig(data, "🙂"),
		},
	})
}

func TestAccCosmosDbFleetspace_validateDataRegionEmoji(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_cosmosdb_fleetspace", "test")
	r := CosmosDbFleetspaceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataRegionConfig(data, `["🙂"]`),
		},
	})
}

// Batch 2

func TestAccCosmosDbFleetspace_validateDataRegion64Elements(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_cosmosdb_fleetspace", "test")
	r := CosmosDbFleetspaceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateDataRegion64ElementsConfig(data),
		},
	})
}

func TestAccCosmosDbFleetspace_validateServiceTierEmoji(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_cosmosdb_fleetspace", "test")
	r := CosmosDbFleetspaceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.validateServiceTierConfig(data, "🙂"),
		},
	})
}

func (r CosmosDbFleetspaceResource) validateNameConfig(data acceptance.TestData, name string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctest-cosmos-%d"
  location = "%s"
}

resource "azurerm_cosmosdb_fleet" "test" {
  name                = "acctest-cosfleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_cosmosdb_fleetspace" "test" {
  name                = "%s"
  fleet_name          = azurerm_cosmosdb_fleet.test.name
  resource_group_name = azurerm_resource_group.test.name
  data_regions         = ["%s"]
  service_tier        = "GeneralPurpose"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, name, data.Locations.Primary)
}

func (r CosmosDbFleetspaceResource) validateFleetNameConfig(data acceptance.TestData, fleetName string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctest-cosmos-%d"
  location = "%s"
}

resource "azurerm_cosmosdb_fleetspace" "test" {
  name                = "acctest-cosfs-%d"
  fleet_name          = "%s"
  resource_group_name = azurerm_resource_group.test.name
  data_regions         = ["%s"]
  service_tier        = "GeneralPurpose"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, fleetName, data.Locations.Primary)
}

func (r CosmosDbFleetspaceResource) validateResourceGroupNameConfig(data acceptance.TestData, resourceGroupName string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_cosmosdb_fleetspace" "test" {
  name                = "acctest-cosfs-%d"
  fleet_name          = "acctest-cosfleet-%d"
  resource_group_name = "%s"
  data_regions         = ["%s"]
  service_tier        = "GeneralPurpose"
}
`, data.RandomInteger, data.RandomInteger, resourceGroupName, data.Locations.Primary)
}

func (r CosmosDbFleetspaceResource) validateDataRegionConfig(data acceptance.TestData, dataRegion string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctest-cosmos-%d"
  location = "%s"
}

resource "azurerm_cosmosdb_fleet" "test" {
  name                = "acctest-cosfleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_cosmosdb_fleetspace" "test" {
  name                = "acctest-cosfs-%d"
  fleet_name          = azurerm_cosmosdb_fleet.test.name
  resource_group_name = azurerm_resource_group.test.name
  data_regions         = %s
  service_tier        = "GeneralPurpose"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, dataRegion)
}

func (r CosmosDbFleetspaceResource) validateDataRegion64ElementsConfig(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctest-cosmos-%d"
  location = "%s"
}

resource "azurerm_cosmosdb_fleet" "test" {
  name                = "acctest-cosfleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_cosmosdb_fleetspace" "test" {
  name                = "acctest-cosfs-%d"
  fleet_name          = azurerm_cosmosdb_fleet.test.name
  resource_group_name = azurerm_resource_group.test.name
  data_regions         = [
    "australiacentral",
    "australiacentral2",
    "australiaeast",
    "australiasoutheast",
    "austriaeast",
    "belgiumcentral",
    "brazilsouth",
    "brazilsoutheast",
    "canadacentral",
    "canadaeast",
    "centralindia",
    "centralus",
    "centraluseuap",
    "chilecentral",
    "denmarkeast",
    "eastasia",
    "eastus",
    "eastus2",
    "eastus2euap",
    "francecentral",
    "francesouth",
    "germanynorth",
    "germanywestcentral",
    "indonesiacentral",
    "israelcentral",
    "israelnorthwest",
    "italynorth",
    "japaneast",
    "japanwest",
    "jioindiacentral",
    "jioindiawest",
    "koreacentral",
    "koreasouth",
    "malaysiasouth",
    "malaysiawest",
    "mexicocentral",
    "newzealandnorth",
    "northcentralus",
    "northeurope",
    "norwayeast",
    "norwaywest",
    "polandcentral",
    "qatarcentral",
    "southafricanorth",
    "southafricawest",
    "southcentralus",
    "southcentralus2",
    "southeastasia",
    "southeastus",
    "southindia",
    "spaincentral",
    "swedencentral",
    "swedensouth",
    "switzerlandnorth",
    "switzerlandwest",
    "uaecentral",
    "uaenorth",
    "uksouth",
    "ukwest",
    "westcentralus",
    "westeurope",
    "westindia",
    "westus",
    "westus2",
    "westus3",
    "eastus3",
    "eastusslv",
    "indiasouthcentral",
    "northeastus5",
    "saudiarabiaeast",
    "southeastasia3",
    "southeastus3",
    "southeastus5",
    "southwestus",
    "westcentralusfre"
  ]
  service_tier        = "GeneralPurpose"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger)
}

func (r CosmosDbFleetspaceResource) validateServiceTierConfig(data acceptance.TestData, serviceTier string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctest-cosmos-%d"
  location = "%s"
}

resource "azurerm_cosmosdb_fleet" "test" {
  name                = "acctest-cosfleet-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_cosmosdb_fleetspace" "test" {
  name                = "acctest-cosfs-%d"
  fleet_name          = azurerm_cosmosdb_fleet.test.name
  resource_group_name = azurerm_resource_group.test.name
  data_regions         = ["%s"]
  service_tier        = "%s"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.Locations.Primary, serviceTier)
}
