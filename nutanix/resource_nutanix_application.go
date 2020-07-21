package nutanix

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"math/rand"
	"strconv"
	"strings"
	v3 "terraform-provider-nutanix/client/v3"
	"time"
)

func resourceNutanixApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceNutanixApplicationCreate,
		Read:   resourceNutanixApplicationRead,
		Update: resourceNutanixApplicationUpdate,
		Delete: resourceNutanixApplicationDelete,
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spec_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"project_reference": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"api_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nic_list_status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nic_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"floating_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_function_nic_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_endpoint_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"network_function_chain_reference": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kind": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"subnet_uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_connected": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// RESOURCES ARGUMENTS
			"marketitem_uuid" : {
				Type: schema.TypeString,
				Required: true,
			},
			"project_uuid": {
				Type: schema.TypeString,
				Required: true,
			},
			"environment_uuid": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nic_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nic_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"model": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"network_function_nic_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_address": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_endpoint_list": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"network_function_chain_reference": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kind": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"uuid": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"subnet_uuid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"subnet_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_connected": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "true",
						},
					},
				},
			},
			"guest_os_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"power_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"num_vcpus_per_socket": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_sockets": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memory_size_mib": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"boot_device_order_list": {
				Type: schema.TypeList,
				// remove MaxItems when the issue #28 is fixed
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"boot_device_disk_address": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_index": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"adapter_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"power_state_mechanism": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disk_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"disk_size_bytes": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"disk_size_mib": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"storage_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"flash_mode": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"storage_container_reference": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"kind": {
													Type:     schema.TypeString,
													Optional: true,
													Default:  "storage_container",
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"uuid": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"device_properties": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"disk_address": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"device_index": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"adapter_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"data_source_reference": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kind": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"uuid": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"volume_group_reference": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,

							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kind": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"uuid": {
										Type:     schema.TypeString,
										Optional: true,
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

func resourceNutanixApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	// Get client connection
	conn := meta.(*Client).API

	// Read arguments for blueprint cloning
	itemUUID, noUUID := d.GetOk("marketitem_uuid")
	projectUUID, noProject := d.GetOk("project_uuid")
	environmentUUID, noEnironment := d.GetOk("environment_uuid")
	if !noUUID || !noProject || !noEnironment {
		return fmt.Errorf("please provide market uuid to launch")
	}
	appDescription := d.Get("description")
	appName, noName := d.GetOk("name")
	if !noName {
		return fmt.Errorf("please provide name for the application")
	}


	resp, err := conn.V3.GetMarketItem(itemUUID.(string))
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Calm get marketplace item: #%v\n", resp.Metadata)

	blueprintCloneRequest := v3.BlueprintCloneRequest{}
	blueprintCloneRequest.Spec.Description = resp.Spec.Description
	blueprintCloneRequest.Spec.SourceMarketplaceName = resp.Spec.Name
	blueprintCloneRequest.Spec.SourceMarketplaceVersion = resp.Spec.Resources.Version
	rand.Seed(time.Now().UnixNano())
	blueprintCloneRequest.Spec.AppBlueprintName = resp.Spec.Name + strconv.Itoa(rand.Intn(8999999999) + 1000000000)
	blueprintCloneRequest.Spec.Resources = resp.Spec.Resources.AppBlueprintTemplate.Spec.Resources
	blueprintCloneRequest.ApiVersion = "3.0"
	itemKind := "blueprint"
	blueprintCloneRequest.Metadata.Kind = &itemKind
	blueprintCloneRequest.Metadata.ProjectReference.Kind = "project"
	blueprintCloneRequest.Metadata.ProjectReference.UUID = projectUUID.(string)
	blueprintCloneRequest.Spec.EnvironmentUUID = environmentUUID.(string)
	//log.Printf("blueprint: %s", blueprintCloneRequest.Spec.Resources.ServiceDefinitionList[0].Name)


	// clone Maketplace item to a temp blueprint
	cloneResp, err := conn.V3.CloneBlueprint(&blueprintCloneRequest)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Calm Cloned marketitem to blueprint new uuid: %s", *cloneResp.Metadata.UUID)

	// get the cloned blueprint with the new uuid
	clonedBlueprint, err := conn.V3.GetBlueprint(*cloneResp.Metadata.UUID)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Calm Get of new blueprint uuid:%s, spec version: %d", *clonedBlueprint.Metadata.UUID, *clonedBlueprint.Metadata.SpecVersion)
	log.Printf("[DEBUG] Calm new blueprint environment uuid:%s", clonedBlueprint.Spec.Resources.AppProfileList[0].UUID)

	// prepare payload for launching the new blueprint
	blueprintLaunchRequest := new(v3.BlueprintLaunchRequest)

	blueprintLaunchRequest.Spec.Description = appDescription.(string)
	blueprintLaunchRequest.Spec.ApplicationName = appName.(string)
	blueprintLaunchRequest.Spec.AppProfileReference.Kind = "app_profile"
	blueprintLaunchRequest.Spec.AppProfileReference.UUID = clonedBlueprint.Spec.Resources.AppProfileList[0].UUID
	blueprintLaunchRequest.Spec.Resources = clonedBlueprint.Spec.Resources
	blueprintLaunchRequest.ApiVersion = "3.0"
	blueprintLaunchRequest.Metadata = clonedBlueprint.Metadata

	newBlueprintResponse, err := conn.V3.LaunchBlueprint(blueprintLaunchRequest)
	if err != nil {
		return err
	}
	refBlueprintUUID := newBlueprintResponse.Metadata.UUID
	log.Printf("[DEBUG] Reference Blueprint UUID: %s", refBlueprintUUID)

	// wait for the new blueprint to be cloned
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"QUEUED", "RUNNING"},
		Target:     []string{"SUCCEEDED"},
		Refresh:    taskStateRefreshFunc(conn, newBlueprintResponse.Status.RequestID),
		Timeout:    vmTimeout,
		Delay:      vmDelay,
		MinTimeout: vmMinTimeout,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"error waiting for blueprint (%s) to launch: %s", newBlueprintResponse.Metadata.UUID, err)
	}

	fullApplicationList, err := conn.V3.ListAllApplications()
	if err != nil {
		return err
	}

	var newApp = v3.ApplicationEntity{}
	for _, app := range fullApplicationList.Entities {
		if app.Status.ConfigReference == refBlueprintUUID {
			newApp = app
			log.Printf("[DEBUG] Calm found the new app uuid: %s", *newApp.Metadata.UUID)
		}
	}

	d.SetId(*newApp.Metadata.UUID)
	return resourceNutanixApplicationRead(d, meta)
}

func resourceNutanixApplicationRead(d *schema.ResourceData, meta interface{}) error {
	// Get client connection
	conn := meta.(*Client).API
	setVMTimeout(meta)
	// Make request to the API
	resp, err := conn.V3.GetApplication(d.Id())

	if err != nil {
		if strings.Contains(fmt.Sprint(err), "ENTITY_NOT_FOUND") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error reading Virtual Machine %s: %s", d.Id(), err)
	}

	// Added check for deletion. Re-running TF right after VM deletion, can cause an error because the ID is still present in API.
	// Check if state is deleting
	if resp.Status.State == "deleting"{
		d.SetId("")
		return nil
	}

	d.SetId(*resp.Metadata.UUID)
	return nil
}

func resourceNutanixApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceNutanixApplicationRead(d, meta)
}

func resourceNutanixApplicationDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*Client).API
	setVMTimeout(meta)
	log.Printf("[DEBUG] Deleting application: %s, %s", d.Get("name").(string), d.Id())
	resp, err := conn.V3.DeleteApplication(d.Id())
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "ENTITY_NOT_FOUND") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error while deleting application UUID(%s): %s", d.Id(), err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"QUEUED", "RUNNING"},
		Target:     []string{"SUCCEEDED"},
		Refresh:    taskStateRefreshFunc(conn, resp.Status.ErgonTaskUUID),
		Timeout:    vmTimeout,
		Delay:      vmDelay,
		MinTimeout: vmMinTimeout,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"error waiting for app (%s) to delete: %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}

