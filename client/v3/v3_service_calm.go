package v3

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/terraform-providers/terraform-provider-nutanix/utils"
)

/*ListAllMarketItem gets a list of marketplace items published for the user
 * This operation gets a list of Marketplace items, allowing for sorting and pagination.
 * Note: Entities that have not been created successfully are not listed.
 * @return *MarketItemListResponse
 */
func (op Operations) ListAllMarketItem() (*MarketItemListResponse, error) {
	entities := make([]*MarketItem, 0)

	resp, err := op.ListMarketItem(&DSMetadata{
		Kind:   utils.StringPtr("marketplace_item"),
		Filter: utils.StringPtr("app_state==PUBLISHED"),
		Length: utils.Int64Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int64Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int64Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListMarketItem(&DSMetadata{
				Kind:   utils.StringPtr("marketplace_item"),
				Filter: utils.StringPtr("app_state==PUBLISHED"),
				Length: utils.Int64Ptr(itemsPerPage),
				Offset: utils.Int64Ptr(offset),
			})

			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

/*ListMarketItem gets a list of Markletplace Items.
 *
 * @param metadata allows create filters to get specific data - *DSMetadata.
 * @return *MarketItemListResponse
 */
func (op Operations) ListMarketItem(getEntitiesRequest *DSMetadata) (*MarketItemListResponse, error) {
	ctx := context.TODO()
	path := "/calm_marketplace_items/list"

	marketItemList := new(MarketItemListResponse)

	req, err := op.client.NewRequest(ctx, http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	return marketItemList, op.client.Do(ctx, req, marketItemList)
}

func (op Operations) GetMarketItem(uuid string) (*MarketItemSpecResponse, error) {
	ctx := context.TODO()
	path := fmt.Sprintf("/calm_marketplace_items/%s", uuid)

	req, err := op.client.NewRequest(ctx, http.MethodGet, path, nil)

	marketItemResponse := new(MarketItemSpecResponse)

	if err != nil {
		return nil, err
	}

	return marketItemResponse, op.client.Do(ctx, req, marketItemResponse)
}

func (op Operations) CloneBlueprint(blueprint *BlueprintCloneRequest) (*BlueprintCloneResponse, error) {
	ctx := context.TODO()

	req, err := op.client.NewRequest(ctx, http.MethodPost, "/blueprints/marketplace_launch", blueprint)
	cloneResponse := new(BlueprintCloneResponse)
	if err != nil {
		return nil, err
	}

	return cloneResponse, op.client.Do(ctx, req, cloneResponse)
}

func (op Operations) GetBlueprint(uuid string) (*BlueprintGetResponse, error) {
	ctx := context.TODO()
	path := fmt.Sprintf("/blueprints/%s", uuid)

	req, err := op.client.NewRequest(ctx, http.MethodGet, path, nil)

	blueprintResponse := new(BlueprintGetResponse)

	if err != nil {
		return nil, err
	}

	return blueprintResponse, op.client.Do(ctx, req, blueprintResponse)
}

func (op Operations) LaunchBlueprint(blueprint *BlueprintLaunchRequest) (*BlueprintLaunchResponse, error) {
	ctx := context.TODO()

	path := fmt.Sprintf("/blueprints/%s/launch", *blueprint.Metadata.UUID)
	req, err := op.client.NewRequest(ctx, http.MethodPost, path, blueprint)
	launchResponse := new(BlueprintLaunchResponse)
	if err != nil {
		return nil, err
	}

	return launchResponse, op.client.Do(ctx, req, launchResponse)
}

func (op Operations) ListApplications(getEntitiesRequest *DSMetadata) (*ApplicationListResponse, error) {
	ctx := context.TODO()
	path := "/apps/list"

	applicationList := new(ApplicationListResponse)

	req, err := op.client.NewRequest(ctx, http.MethodPost, path, getEntitiesRequest)
	if err != nil {
		return nil, err
	}

	return applicationList, op.client.Do(ctx, req, applicationList)
}

func (op Operations) ListAllApplications () (*ApplicationListResponse, error) {
	entities := make([]ApplicationEntity, 0)

	resp, err := op.ListApplications(&DSMetadata{
		Kind:   utils.StringPtr("app"),
		Length: utils.Int64Ptr(itemsPerPage),
	})
	if err != nil {
		return nil, err
	}

	totalEntities := utils.Int64Value(resp.Metadata.TotalMatches)
	remaining := totalEntities
	offset := utils.Int64Value(resp.Metadata.Offset)

	if totalEntities > itemsPerPage {
		for hasNext(&remaining) {
			resp, err = op.ListApplications(&DSMetadata{
				Kind:   utils.StringPtr("app"),
				Length: utils.Int64Ptr(itemsPerPage),
				Offset: utils.Int64Ptr(offset),
			})

			if err != nil {
				return nil, err
			}

			entities = append(entities, resp.Entities...)

			offset += itemsPerPage
			log.Printf("[Debug] app_list total=%d, remaining=%d, offset=%d len(entities)=%d\n", totalEntities, remaining, offset, len(entities))
		}

		resp.Entities = entities
	}

	return resp, nil
}

func (op Operations) GetApplication(uuid string) (*ApplicationGetResponse, error) {
	ctx := context.TODO()
	path := fmt.Sprintf("/apps/%s", uuid)

	req, err := op.client.NewRequest(ctx, http.MethodGet, path, nil)

	applicationResponse := new(ApplicationGetResponse)

	if err != nil {
		return nil, err
	}

	return applicationResponse, op.client.Do(ctx, req, applicationResponse)
}

func (op Operations) DeleteApplication(uuid string) (*ApplicationDeleteResponse, error) {
	ctx := context.TODO()

	path := fmt.Sprintf("/apps/%s", uuid)

	req, err := op.client.NewRequest(ctx, http.MethodDelete, path, nil)
	deleteResponse := new(ApplicationDeleteResponse)

	if err != nil {
		return nil, err
	}

	return deleteResponse, op.client.Do(ctx, req, deleteResponse)
}