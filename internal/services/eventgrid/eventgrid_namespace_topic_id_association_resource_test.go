package eventgrid_test

import (
	"context"
	"fmt"
	"regexp"
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

func TestAccEventgridNamespaceTopicIdAssociation_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventgrid_namespace_topic_id_association", "test")
	r := EventgridNamespaceTopicIdAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccEventgridNamespaceTopicIdAssociation_update(t *testing.T) {
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
		{
			Config: r.updateToTopicFromOtherNamespace(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccEventgridNamespaceTopicIdAssociation_topicIdSetTwiceError(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_eventgrid_namespace_topic_id_association", "test")
	r := EventgridNamespaceTopicIdAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.topicIdSetTwiceError(data),
			ExpectError: regexp.MustCompile("updating .*: topic ID of `azurerm_eventgrid_namespace` resource is found to be set in multiple resources, ensure that only 1 `azurerm_eventgrid_namespace_topic_id_association` resource is set without setting `topic_spaces_configuration.0.route_topic_id` property of `azurerm_eventgrid_namespace`, or only `topic_spaces_configuration.0.route_topic_id` property of `azurerm_eventgrid_namespace` is set without creating `azurerm_eventgrid_namespace_topic_id_association`, run `terraform import` on `azurerm_eventgrid_namespace_topic_id_association` if this message keeps showing up due to existing topic ID"),
		},
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

func (r EventgridNamespaceTopicIdAssociationResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_eventgrid_namespace_topic_id_association" "import" {
  eventgrid_namespace_id = azurerm_eventgrid_namespace_topic_id_association.eventgrid_namespace_id
  eventgrid_namespace_topic_id = azurerm_eventgrid_namespace_topic_id_association.eventgrid_namespace_topic_id
}
`, r.basic(data))
}

func (r EventgridNamespaceTopicIdAssociationResource) updateToTopicFromOtherNamespace(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_eventgrid_namespace" "test2" {
  name                = "acctest-evgns2-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_eventgrid_namespace_topic" "test2" {
  name = "acctest-evgnst2-%[2]d"
  eventgrid_namespace_id = azurerm_eventgrid_namespace.test2.id
}

resource "azurerm_eventgrid_namespace_topic_id_association" test {
  eventgrid_namespace_id = azurerm_eventgrid_namespace.test.id
  eventgrid_namespace_topic_id = azurerm_eventgrid_namespace_topic.test2.id
}
`, r.template(data), data.RandomInteger)
}

func (EventgridNamespaceTopicIdAssociationResource) topicIdSetTwiceError(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name = "acctest-evg-%[1]d"
  location = "%[2]s"
}

resource "azurerm_eventgrid_topic" "test" {
  name                = "acctest-evgt-%[1]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_eventgrid_namespace" "test" {
  name                = "acctest-evgns-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  topic_spaces_configuration {
    route_topic_id = azurerm_eventgrid_topic.test.id
  }
}

resource "azurerm_eventgrid_namespace_topic" "test" {
  name = "acctest-evgnst-%[1]d"
  eventgrid_namespace_id = azurerm_eventgrid_namespace.test.id
}

resource "azurerm_eventgrid_namespace_topic_id_association" test {
  eventgrid_namespace_id = azurerm_eventgrid_namespace.test.id
  eventgrid_namespace_topic_id = azurerm_eventgrid_namespace_topic.test.id
}
`, data.RandomInteger, data.Locations.Primary)
}

func (r EventgridNamespaceTopicIdAssociationResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name = "acctest-evg-%[1]d"
  location = "%[2]s"
}

resource "azurerm_eventgrid_namespace" "test" {
  name                = "acctest-evgns-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_eventgrid_namespace_topic" "test" {
  name = "acctest-evgnst-%[1]d"
  eventgrid_namespace_id = azurerm_eventgrid_namespace.test.id
}
`, data.RandomInteger, data.Locations.Primary)
}
