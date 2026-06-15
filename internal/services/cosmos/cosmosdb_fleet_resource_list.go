// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package cosmos

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-10-15/fleets"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type CosmosdbFleetListResource struct{}

var _ sdk.FrameworkListWrappedResource = new(CosmosdbFleetListResource)

func (CosmosdbFleetListResource) ResourceFunc() *pluginsdk.Resource {
	return sdk.WrappedResource(CosmosdbFleetResource{})
}

func (CosmosdbFleetListResource) Metadata(_ context.Context, _ resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = "azurerm_cosmosdb_fleet"
}

func (CosmosdbFleetListResource) List(ctx context.Context, request list.ListRequest, stream *list.ListResultsStream, metadata sdk.ResourceMetadata) {
	client := metadata.Client.Cosmos.CosmosdbFleetsClient

	var data sdk.DefaultListModel
	diags := request.Config.Get(ctx, &data)
	if diags.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
	}

	results := make([]fleets.FleetResource, 0)

	subscriptionID := metadata.SubscriptionId
	if !data.SubscriptionId.IsNull() {
		subscriptionID = data.SubscriptionId.ValueString()
	}

	switch {
	case !data.ResourceGroupName.IsNull():
		resp, err := client.FleetListByResourceGroupComplete(ctx, commonids.NewResourceGroupID(subscriptionID, data.ResourceGroupName.ValueString()))
		if err != nil {
			sdk.SetResponseErrorDiagnostic(stream, fmt.Sprintf("listing `%s`", "azurerm_cosmosdb_fleet"), err)
			return
		}

		results = resp.Items
	default:
		resp, err := client.FleetListComplete(ctx, commonids.NewSubscriptionID(subscriptionID))
		if err != nil {
			sdk.SetResponseErrorDiagnostic(stream, fmt.Sprintf("listing `%s`", "azurerm_cosmosdb_fleet"), err)
			return
		}

		results = resp.Items
	}

	stream.Results = func(push func(list.ListResult) bool) {
		for _, CosmosdbFleetResult := range results {
			result := request.NewListResult(ctx)
			result.DisplayName = pointer.From(CosmosdbFleetResult.Name)

			id, err := fleets.ParseFleetID(pointer.From(CosmosdbFleetResult.Id))
			if err != nil {
				sdk.SetErrorDiagnosticAndPushListResult(result, push, "parsing Cosmosdb Fleet ID", err)
				return
			}

			rmd := sdk.NewResourceMetaData(metadata.Client, CosmosdbFleetResource{})
			rmd.ResourceData.SetId(id.ID())

			resp := fleets.FleetGetOperationResponse{
				Model: &CosmosdbFleetResult,
			}
			fleetResource := CosmosdbFleetResource{}
			if err := fleetResource.flatten(id, resp, rmd); err != nil {
				sdk.SetErrorDiagnosticAndPushListResult(result, push, fmt.Sprintf("encoding `%s` resource data", "azurerm_cosmosdb_fleet"), err)
				return
			}

			sdk.EncodeListResult(ctx, rmd.ResourceData, &result)
			if result.Diagnostics.HasError() {
				push(result)
				return
			}

			if !push(result) {
				return
			}
		}
	}
}
