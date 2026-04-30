// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package cosmos_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

func TestAccCosmosDbFleetspace_updateMinThroughput(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_cosmosdb_fleetspace", "test")
	r := CosmosDbFleetspaceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateMinThroughputConfig(data, 100000),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateMinThroughputConfig(data, 101000),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateMinThroughputConfig(data, 100000),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccCosmosDbFleetspace_updateMaxThroughput(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_cosmosdb_fleetspace", "test")
	r := CosmosDbFleetspaceResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateMaxThroughputConfig(data, 101000),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateMaxThroughputConfig(data, 102000),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateMaxThroughputConfig(data, 101000),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r CosmosDbFleetspaceResource) updateMinThroughputConfig(data acceptance.TestData, minThroughput int) string {
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
  data_region         = ["%s"]
  service_tier        = "GeneralPurpose"
  min_throughput      = %d
  max_throughput = 101000
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.Locations.Primary, minThroughput)
}

func (r CosmosDbFleetspaceResource) updateMaxThroughputConfig(data acceptance.TestData, maxThroughput int) string {
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
  data_region         = ["%s"]
  service_tier        = "GeneralPurpose"
  min_throughput = 100000
  max_throughput      = %d
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.Locations.Primary, maxThroughput)
}
