package nutanix

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	v3 "github.com/terraform-providers/terraform-provider-nutanix/client/v3"
)

func dataSourceNutanixProjects() *schema.Resource {
	return &schema.Resource{
		Read:          dataSourceNutanixProjectsRead,
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
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default": {
							Type:     schema.TypeBool,
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

func dataSourceNutanixProjectsRead(d *schema.ResourceData, meta interface{}) error {
	//Get client connection
	conn := meta.(*Client).API

	resp, err := conn.V3.ListAllProject()
	if err != nil {
		return err
	}

	if err := d.Set("api_version", resp.APIVersion); err != nil {
		return err
	}
	//log.Printf("[DEBUG] Husain - Entities: %+v\n\n", resp.Metadata)
	if err := d.Set("entities", flattenProjectEntities(resp.Entities)); err != nil {
		return err
	}

	d.SetId(resource.UniqueId())
	return nil
}

func flattenProjectEntities(projects []*v3.Project) []map[string]interface{} {
	entities := make([]map[string]interface{}, len(projects))
	//log.Printf("[DEBUG] Husain - Projects count: %d ", len(projects))

	for i, project := range projects {
		metadata, _ := setRSEntityMetadata(project.Metadata)
		//log.Printf("Husain %d - Name: %s", i, project.Status.Name)
		//log.Printf("Husain - description: %s", project.Status.Description)
		//log.Printf("Husain - State: %s", project.Status.State)
		//log.Printf("Husain - Metadata: %+v\n", metadata)

		entities[i] = map[string]interface{}{
			"name":						project.Status.Name,
			"description":              project.Status.Descripion,
			"state":                    project.Status.State,
			"is_default":               project.Status.Resources.IsDefault,
			"metadata":                 metadata,
		}
	}
	return entities
}
