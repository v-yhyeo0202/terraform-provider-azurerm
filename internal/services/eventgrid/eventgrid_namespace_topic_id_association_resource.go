package eventgrid

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2025-02-15/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2025-02-15/namespacetopics"
	"github.com/hashicorp/terraform-provider-azurerm/internal/locks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type EventgridNamespaceTopicIdAssociationResource struct{}

var (
	_ sdk.ResourceWithIdentity = EventgridNamespaceTopicIdAssociationResource{}
)

type EventgridNamespaceTopicIdAssociationModel struct {
	EventgridNamespaceId      string `tfschema:"eventgrid_namespace_id"`
	EventgridNamespaceTopicId string `tfschema:"eventgrid_namespace_topic_id"`
}

func (EventgridNamespaceTopicIdAssociationResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"eventgrid_namespace_id": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: namespaces.ValidateNamespaceID,
		},

		"eventgrid_namespace_topic_id": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: namespacetopics.ValidateNamespaceTopicID,
		},
	}
}

func (EventgridNamespaceTopicIdAssociationResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (EventgridNamespaceTopicIdAssociationResource) ModelObject() interface{} {
	return &EventgridNamespaceTopicIdAssociationModel{}
}

func (EventgridNamespaceTopicIdAssociationResource) ResourceType() string {
	return "azurerm_eventgrid_namespace_topic_id_association"
}

func (r EventgridNamespaceTopicIdAssociationResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.EventGrid.NamespacesClient_v2025_02_15
			var config EventgridNamespaceTopicIdAssociationModel
			if err := metadata.Decode(&config); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			eventgridNamespaceId, err := namespaces.ParseNamespaceID(config.EventgridNamespaceId)
			if err != nil {
				return err
			}

			locks.ByID(config.EventgridNamespaceId)
			defer locks.UnlockByID(config.EventgridNamespaceId)

			existing, err := client.Get(ctx, *eventgridNamespaceId)
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", *eventgridNamespaceId, err)
			}

			if existing.Model == nil {
				return fmt.Errorf("retrieving %s: `model` was nil", *eventgridNamespaceId)
			}
			if existing.Model.Properties == nil {
				return fmt.Errorf("retrieving %s: `properties` was nil", *eventgridNamespaceId)
			}

			if existing.Model.Properties.TopicsConfiguration == nil {
				existing.Model.Properties.TopicSpacesConfiguration = &namespaces.TopicSpacesConfiguration{}
			}

			if !metadata.ResourceData.IsNewResource() {
				if existing.Model.Properties.TopicSpacesConfiguration.RouteTopicResourceId != nil {
					return fmt.Errorf("updating %s: topic ID of `azurerm_eventgrid_namespace` resource is found to be set in multiple resources, ensure that only 1 `azurerm_eventgrid_namespace_topic_id_association` resource is set, or only `topic_spaces_configuration.0.route_topic_id` property of `azurerm_eventgrid_namespace` is set without creating `azurerm_eventgrid_namespace_topic_id_association`, run `terraform import` on `azurerm_eventgrid_namespace_topic_id_association` if this message keeps showing up due to existing topic ID", *eventgridNamespaceId)
				}
			}

			existing.Model.Properties.TopicSpacesConfiguration.RouteTopicResourceId = pointer.To(config.EventgridNamespaceTopicId)

			if err := client.CreateOrUpdateCallbackThenPoll(ctx, *eventgridNamespaceId, *existing.Model, metadata.SetIDCallback(eventgridNamespaceId)); err != nil {
				return fmt.Errorf("updating Namespace Topic ID for %s: %+v", *eventgridNamespaceId, err)
			}

			metadata.SetID(eventgridNamespaceId)
			return nil
		},
	}
}

func (EventgridNamespaceTopicIdAssociationResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.EventGrid.NamespacesClient_v2025_02_15
			var config EventgridNamespaceTopicIdAssociationModel
			if err := metadata.Decode(&config); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			id, err := namespaces.ParseNamespaceID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			state := EventgridNamespaceTopicIdAssociationModel{}

			if model := resp.Model; model != nil {
				state.EventgridNamespaceId = pointer.From(model.Id)

				if props := model.Properties; props != nil {
					if topicSpacesConfiguration := props.TopicSpacesConfiguration; topicSpacesConfiguration != nil {
						state.EventgridNamespaceTopicId = pointer.From(topicSpacesConfiguration.RouteTopicResourceId)
					}
				}
			}
			return metadata.Encode(&state)
		},
	}
}

func (EventgridNamespaceTopicIdAssociationResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.EventGrid.NamespacesClient_v2025_02_15
			id, err := namespaces.ParseNamespaceID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.Get(ctx, *id)
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			if resp.Model == nil {
				return fmt.Errorf("retrieving %s: `model` was nil", *id)
			}
			if resp.Model.Properties == nil {
				return fmt.Errorf("retrieving %s: `properties` was nil", *id)
			}

			if resp.Model.Properties.TopicSpacesConfiguration == nil {
				return fmt.Errorf("retrieving %s: `topicSpacesConfiguration` was nil", *id)
			}

			resp.Model.Properties.TopicSpacesConfiguration.RouteTopicResourceId = nil
			if err := client.CreateOrUpdateThenPoll(ctx, *id, *resp.Model); err != nil {
				return fmt.Errorf("removing Namespace Topic ID from %s: %+v", *id, err)
			}

			return nil
		},
	}
}
