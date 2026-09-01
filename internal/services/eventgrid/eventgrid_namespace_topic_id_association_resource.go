package eventgrid

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2025-02-15/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2025-02-15/namespacetopics"
	"github.com/hashicorp/terraform-provider-azurerm/internal/locks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

//go:generate go run ../../tools/generator-tests resourceidentity -parent-id "eventgrid_namespace_id"

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
			Type:     pluginsdk.TypeString,
			Required: true,
			// `ForceNew` as `eventgrid_namespace_topic_id` can only be assigned to `eventgrid_namespace_id` of same namespace and `eventgrid_namespace_topic_id` is unique for each namespace
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

func (EventgridNamespaceTopicIdAssociationResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return namespaces.ValidateNamespaceID
}

func (EventgridNamespaceTopicIdAssociationResource) Identity() resourceids.ResourceId {
	return &namespaces.NamespaceId{}
}

func (EventgridNamespaceTopicIdAssociationResource) IdentityType() pluginsdk.ResourceTypeForIdentity {
	return pluginsdk.ResourceTypeForIdentityVirtual
}

func (EventgridNamespaceTopicIdAssociationResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.EventGrid.NamespacesClient_v2025_02_15
			var config EventgridNamespaceTopicIdAssociationModel
			if err := metadata.Decode(&config); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			id, err := namespaces.ParseNamespaceID(config.EventgridNamespaceId)
			if err != nil {
				return err
			}

			locks.ByID(config.EventgridNamespaceId)
			defer locks.UnlockByID(config.EventgridNamespaceId)

			existing, err := client.Get(ctx, *id)
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}

			if existing.Model == nil {
				return fmt.Errorf("retrieving %s: `model` was nil", *id)
			}
			if existing.Model.Properties == nil {
				return fmt.Errorf("retrieving %s: `properties` was nil", *id)
			}

			if existing.Model.Properties.TopicSpacesConfiguration == nil {
				existing.Model.Properties.TopicSpacesConfiguration = &namespaces.TopicSpacesConfiguration{}
			}

			existing.Model.Properties.TopicSpacesConfiguration.RouteTopicResourceId = pointer.To(config.EventgridNamespaceTopicId)

			if err := client.CreateOrUpdateCallbackThenPoll(ctx, *id, *existing.Model, metadata.SetIDAndIdentityCallback(id)); err != nil {
				return fmt.Errorf("updating Namespace Topic ID for %s: %+v", *id, err)
			}

			metadata.SetID(id)
			return pluginsdk.SetResourceIdentityData(metadata.ResourceData, id, pluginsdk.ResourceTypeForIdentityVirtual)
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

			if err := pluginsdk.SetResourceIdentityData(metadata.ResourceData, id, pluginsdk.ResourceTypeForIdentityVirtual); err != nil {
				return err
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

			// If parent resource is not found or `TopicSpacesConfiguration` property cannot be accessed, assume that `RouteTopicResourceId` is already unset and continue to remove this resource from state
			resp, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return nil
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			if resp.Model == nil {
				return nil
			}
			if resp.Model.Properties == nil {
				return nil
			}

			if resp.Model.Properties.TopicSpacesConfiguration == nil {
				return nil
			}

			resp.Model.Properties.TopicSpacesConfiguration.RouteTopicResourceId = nil
			if err := client.CreateOrUpdateThenPoll(ctx, *id, *resp.Model); err != nil {
				return fmt.Errorf("removing Namespace Topic ID from %s: %+v", *id, err)
			}

			return nil
		},
	}
}
