package nutanix

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	v3 "terraform-provider-nutanix/client/v3"
)

func dataSourceNutanixMarketplaceItems() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNutanixMarketplaceItemsRead,
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"app_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"last_update_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"creation_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"spec_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceNutanixMarketplaceItemsRead(d *schema.ResourceData, meta interface{}) error {
	//Get client connection
	conn := meta.(*Client).API

	resp, err := conn.V3.ListAllMarketItem()
	if err != nil {
		return err
	}

	if err := d.Set("api_version", resp.APIVersion); err != nil {
		return err
	}
	if err := d.Set("entities", flattenMarketplaceItemsEntities(resp.Entities)); err != nil {
		return err
	}

	d.SetId(resource.UniqueId())
	return nil
}

func flattenMarketplaceItemsEntities(items []*v3.MarketItem) []map[string]interface{} {
	entities := make([]map[string]interface{}, len(items))

	for i, item := range items {
		metadata, _ := setCalmEntityMetadata(item.Metadata)

		entities[i] = map[string]interface{}{
			"name":						item.Status.Name,
			"description":              item.Status.Description,
			"app_state":                item.Status.AppState,
			"version":					item.Status.Version,
			"metadata":                 metadata,
		}
	}
	return entities
}