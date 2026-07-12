// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package orbitalplanetarycomputer_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/orbitalplanetarycomputer/2026-04-15/geocatalogs"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type GeoCatalogResource struct{}

func TestAccGeoCatalog_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_orbital_geo_catalog", "test")
	r := GeoCatalogResource{}

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

func (r GeoCatalogResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := geocatalogs.ParseGeoCatalogID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.OrbitalPlanetaryComputer.GeoCatalogsClient.Get(ctx, *id)
	if err != nil {
		return nil, fmt.Errorf("retrieving %s: %+v", *id, err)
	}

	return pointer.To(resp.Model != nil), nil
}

func (r GeoCatalogResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_orbital_geo_catalog" "test" {
  name                = "acctest-geocatalog-%s"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}
`, r.template(data), data.RandomString)
}

func (r GeoCatalogResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctest-planetarycomputer-%d"
  location = "%s"
}
`, data.RandomInteger, data.Locations.Primary)
}
