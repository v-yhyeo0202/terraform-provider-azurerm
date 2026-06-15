// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package cosmos

import (
	"context"
	"time"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-10-15/fleets"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type CosmosDbFleetModel struct {
	Name              string            `tfschema:"name"`
	ResourceGroupName string            `tfschema:"resource_group_name"`
	Location          string            `tfschema:"location"`
	Tags              map[string]string `tfschema:"tags"`
}

type CosmosDbFleetResource struct{}

var _ sdk.Resource = CosmosDbFleetResource{}
var _ sdk.ResourceWithUpdate = CosmosDbFleetResource{}

func (r CosmosDbFleetResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"resource_group_name": commonschema.ResourceGroupName(),

		"location": commonschema.Location(),

		"tags": commonschema.Tags(),
	}
}

func (r CosmosDbFleetResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r CosmosDbFleetResource) ModelObject() interface{} {
	return &CosmosDbFleetModel{}
}

func (r CosmosDbFleetResource) ResourceType() string {
	return "azurerm_cosmosdb_fleet"
}

func (r CosmosDbFleetResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return fleets.ValidateFleetID
}

func (r CosmosDbFleetResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return nil
		},
	}
}

func (r CosmosDbFleetResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return nil
		},
	}
}

func (r CosmosDbFleetResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return nil
		},
	}
}

func (r CosmosDbFleetResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return nil
		},
	}
}
