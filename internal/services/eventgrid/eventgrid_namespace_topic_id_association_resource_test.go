// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package eventgrid_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2025-02-15/namespaces"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type EventgridNamespaceTopicIdAssociationResource struct{}

func TestAccEventgridNamespaceTopicIdAssociation_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventgrid_namespace_topic_id_association", "test")
	r := EventgridNamespaceTopicIdAssociationResource{}

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

func (EventgridNamespaceTopicIdAssociationResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := namespaces.ParseNamespaceID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.EventGrid.NamespacesClient_v2025_02_15.Get(ctx, *id)
	if err != nil {
		return nil, fmt.Errorf("retrieving %s: %+v", *id, err)
	}

	return pointer.To(resp.Model != nil), nil
}

func (r EventgridNamespaceTopicIdAssociationResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_eventgrid_namespace_topic_id_association" test {
  eventgrid_namespace_id = azurerm_eventgrid_namespace.test.id
  eventgrid_namespace_topic_id = azurerm_eventgrid_namespace_topic.test.id
}
`, r.template(data))
}

func (r EventgridNamespaceTopicIdAssociationResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctest-evg-%[1]d"
  location = "%[2]s"
}

resource "azurerm_eventgrid_namespace" "test" {
  name                = "acctest-evgns-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_eventgrid_namespace_topic" "test" {
  name                   = "acctest-evgnst-%[1]d"
  eventgrid_namespace_id = azurerm_eventgrid_namespace.test.id
}
`, data.RandomInteger, data.Locations.Primary)
}
