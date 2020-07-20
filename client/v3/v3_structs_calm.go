package v3

// This is due to timestamp difference between Prism and Calm API
// CalmMetadata Metadata The kind metadata
type CalmMetadata struct {
	Kind             	*string           		`json:"kind" mapstructure:"kind"`
	UUID             	*string           		`json:"uuid,omitempty" mapstructure:"uuid,omitempty"`
	SpecVersion      	*int64            		`json:"spec_version,omitempty" mapstructure:"spec_version,omitempty"`
	Categories       	map[string]string 		`json:"categories,omitempty" mapstructure:"categories,omitempty"`
	Name             	*string           		`json:"name,omitempty" mapstructure:"name,omitempty"`
	ProjectReference  	ProjectReferenceUUID 	`json:"project_reference,omitempty"`
}

// ProjectReference ...
type ProjectReferenceUUID struct {
	UUID 	string	`json:"uuid,omitempty"`
	Kind 	string	`json:"kind,omitempty"`
	Name 	string	`json:"name,omitempty"`
}


// MarketItemStatus ...
type MarketItemStatus struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AppState    string `json:"app_state,omitempty"`
	Version     string `json:"version,omitempty"`
}

// MarketItem Response object
type MarketItem struct {
	Status     *MarketItemStatus `json:"status,omitempty"`
	Spec       *MarketItemSpec   `json:"spec,omitempty"`
	APIVersion string            `json:"api_version,omitempty"`
	Metadata   *CalmMetadata     `json:"metadata,omitempty"`
}

// MarketItemListResponse Response object
type MarketItemListResponse struct {
	APIVersion string              `json:"api_version,omitempty"`
	Entities   []*MarketItem       `json:"entities,omitempty"`
	Metadata   *ListMetadataOutput `json:"metadata,omitempty"`
}

// MarketItemSpecResponse
type MarketItemSpecResponse struct {
	Metadata	*CalmMetadata	`json:"metadata,omitempty"`
	Spec 		*MarketItemSpec `json:"spec,omitempty"`
}

// MarketItemSpec
type MarketItemSpec struct {
	Description string 						`json:"description,omitempty"`
	Name        string 						`json:"name,omitempty"`
	Resources   *BlueprintResourceTemplate 	`json:"resources,omitempty"`
}

// BlueprintResource ...
type BlueprintResourceTemplate struct {
	AppBlueprintTemplate struct {
		Spec struct {
			Resources	*BlueprintResource		`json:"resources,omitempty"`
		} `json:"spec,omitempty"`
	} `json:"app_blueprint_template,omitempty"`
	ProjectReferenceList []struct {
		Kind string `json:"kind,omitempty"`
		Name string `json:"name,omitempty"`
		UUID string `json:"uuid,omitempty"`
	} `json:"project_reference_list,omitempty"`
	Version string `json:"version,omitempty"`
}

// BlueprintCloneRequest ...
type BlueprintCloneRequest struct {
	Spec		BlueprintSpec	`json:"spec,omitempty"`
	ApiVersion 	string			`json:"api_version,omitempty"`
	Metadata	CalmMetadata	`json:"metadata,omitempty"`
}

// BlueprintSpec ...
type BlueprintSpec struct {
	Description 				string					`json:"description,omitempty"`
	SourceMarketplaceName		string					`json:"source_marketplace_name,omitempty"`
	SourceMarketplaceVersion	string					`json:"source_marketplace_version,omitempty"`
	AppBlueprintName			string					`json:"app_blueprint_name,omitempty"`
	EnvironmentUUID				string					`json:"environment_uuid"`
	Resources					*BlueprintResource		`json:"resources,omitempty"`
}

type BlueprintResource struct {
	ServiceDefinitionList []struct {
		Singleton  bool   `json:"singleton"`
		Name       string `json:"name"`
		ActionList []struct {
			Critical    bool   `json:"critical"`
			Type        string `json:"type"`
			Description string `json:"description"`
			Runbook     struct {
				VariableList       []interface{} `json:"variable_list"`
				TaskDefinitionList []struct {
					TargetAnyLocalReference struct {
						Kind string `json:"kind"`
						Name string `json:"name"`
					} `json:"target_any_local_reference"`
					Retries                      string        `json:"retries"`
					Description                  string        `json:"description"`
					ChildTasksLocalReferenceList []interface{} `json:"child_tasks_local_reference_list"`
					Attrs                        struct {
						Edges []interface{} `json:"edges"`
						Type  string        `json:"type"`
					} `json:"attrs"`
					TimeoutSecs  string        `json:"timeout_secs"`
					Type         string        `json:"type"`
					VariableList []interface{} `json:"variable_list"`
					Name         string        `json:"name"`
				} `json:"task_definition_list"`
				Name                   string `json:"name"`
				MainTaskLocalReference struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
				} `json:"main_task_local_reference"`
				Description string `json:"description"`
			} `json:"runbook"`
			Name string `json:"name"`
		} `json:"action_list"`
		Description   string        `json:"description"`
		PortList      []interface{} `json:"port_list"`
		Tier          string        `json:"tier"`
		VariableList  []interface{} `json:"variable_list"`
		DependsOnList []interface{} `json:"depends_on_list"`
	} `json:"service_definition_list"`
	SubstrateDefinitionList []struct {
		Description    string        `json:"description"`
		ActionList     []interface{} `json:"action_list"`
		ReadinessProbe struct {
			ConnectionType        string `json:"connection_type"`
			Retries               string `json:"retries"`
			ConnectionProtocol    string `json:"connection_protocol"`
			ConnectionPort        int    `json:"connection_port"`
			Address               string `json:"address"`
			DelaySecs             string `json:"delay_secs"`
			DisableReadinessProbe bool   `json:"disable_readiness_probe"`
		} `json:"readiness_probe"`
		Editables struct {
			CreateSpec struct {
				Resources struct {
					GuestCustomization bool `json:"guest_customization"`
					NumSockets         bool `json:"num_sockets"`
					MemorySizeMib      bool `json:"memory_size_mib"`
					SerialPortList     struct {
					} `json:"serial_port_list"`
					NicList struct {
						Num0 struct {
							SubnetReference bool `json:"subnet_reference"`
						} `json:"0"`
					} `json:"nic_list"`
				} `json:"resources"`
			} `json:"create_spec"`
		} `json:"editables"`
		OsType     string `json:"os_type"`
		Type       string `json:"type"`
		CreateSpec struct {
			Name      string `json:"name"`
			Resources struct {
				NicList []struct {
					NicType                       string        `json:"nic_type"`
					IPEndpointList                []interface{} `json:"ip_endpoint_list"`
					NetworkFunctionChainReference interface{}   `json:"network_function_chain_reference"`
					NetworkFunctionNicType        string        `json:"network_function_nic_type"`
					MacAddress                    string        `json:"mac_address"`
					SubnetReference               struct {
						Kind string `json:"kind"`
						Type string `json:"type"`
						Name string `json:"name"`
						UUID string `json:"uuid"`
					} `json:"subnet_reference"`
					Type string `json:"type"`
				} `json:"nic_list"`
				SerialPortList        []interface{} `json:"serial_port_list"`
				GuestTools            interface{}   `json:"guest_tools"`
				NumVcpusPerSocket     int           `json:"num_vcpus_per_socket"`
				NumSockets            int           `json:"num_sockets"`
				GpuList               []interface{} `json:"gpu_list"`
				MemorySizeMib         int           `json:"memory_size_mib"`
				ParentReference       interface{}   `json:"parent_reference"`
				HardwareClockTimezone string        `json:"hardware_clock_timezone"`
				GuestCustomization    struct {
					CloudInit struct {
						MetaData string `json:"meta_data"`
						Type     string `json:"type"`
						UserData string `json:"user_data"`
					} `json:"cloud_init"`
					Type    string      `json:"type"`
					Sysprep interface{} `json:"sysprep"`
				} `json:"guest_customization"`
				PowerState  string `json:"power_state"`
				Type        string `json:"type"`
				AccountUUID string `json:"account_uuid"`
				BootConfig  struct {
					BootDevice struct {
						Type        string `json:"type"`
						DiskAddress struct {
							Type        string `json:"type"`
							DeviceIndex int    `json:"device_index"`
							AdapterType string `json:"adapter_type"`
						} `json:"disk_address"`
					} `json:"boot_device"`
					Type       string `json:"type"`
					BootType   string `json:"boot_type"`
					MacAddress string `json:"mac_address"`
				} `json:"boot_config"`
				DiskList []struct {
					DataSourceReference struct {
						Kind string `json:"kind"`
						Type string `json:"type"`
						Name string `json:"name"`
						UUID string `json:"uuid"`
					} `json:"data_source_reference"`
					Type                 string      `json:"type"`
					DiskSizeMib          int         `json:"disk_size_mib"`
					VolumeGroupReference interface{} `json:"volume_group_reference"`
					DeviceProperties     struct {
						Type        string `json:"type"`
						DiskAddress struct {
							Type        string `json:"type"`
							DeviceIndex int    `json:"device_index"`
							AdapterType string `json:"adapter_type"`
						} `json:"disk_address"`
						DeviceType string `json:"device_type"`
					} `json:"device_properties"`
				} `json:"disk_list"`
			} `json:"resources"`
			AvailabilityZoneReference interface{} `json:"availability_zone_reference"`
			BackupPolicy              interface{} `json:"backup_policy"`
			Type                      string      `json:"type"`
			ClusterReference          interface{} `json:"cluster_reference"`
			Categories                string      `json:"categories"`
		} `json:"create_spec"`
		VariableList []interface{} `json:"variable_list"`
		Name         string        `json:"name"`
	} `json:"substrate_definition_list"`
	EndpointDefinitionList   []interface{} `json:"endpoint_definition_list"`
	CredentialDefinitionList []interface{} `json:"credential_definition_list"`
	PackageDefinitionList    []struct {
		Description               string        `json:"description"`
		ActionList                []interface{} `json:"action_list"`
		ServiceLocalReferenceList []struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
		} `json:"service_local_reference_list"`
		Version string `json:"version"`
		Type    string `json:"type"`
		Options struct {
			InstallRunbook struct {
				MainTaskLocalReference struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
				} `json:"main_task_local_reference"`
				TaskDefinitionList []struct {
					TargetAnyLocalReference struct {
						Kind string `json:"kind"`
						Name string `json:"name"`
					} `json:"target_any_local_reference"`
					Retries                      string        `json:"retries"`
					Name                         string        `json:"name"`
					ChildTasksLocalReferenceList []interface{} `json:"child_tasks_local_reference_list"`
					Attrs                        struct {
						Edges []interface{} `json:"edges"`
						Type  string        `json:"type"`
					} `json:"attrs"`
					TimeoutSecs  string        `json:"timeout_secs"`
					Type         string        `json:"type"`
					VariableList []interface{} `json:"variable_list"`
					Description  string        `json:"description"`
				} `json:"task_definition_list"`
				Description  string        `json:"description"`
				VariableList []interface{} `json:"variable_list"`
				Name         string        `json:"name"`
			} `json:"install_runbook"`
			Type             string `json:"type"`
			UninstallRunbook struct {
				MainTaskLocalReference struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
				} `json:"main_task_local_reference"`
				TaskDefinitionList []struct {
					TargetAnyLocalReference struct {
						Kind string `json:"kind"`
						Name string `json:"name"`
					} `json:"target_any_local_reference"`
					Retries                      string        `json:"retries"`
					Name                         string        `json:"name"`
					ChildTasksLocalReferenceList []interface{} `json:"child_tasks_local_reference_list"`
					Attrs                        struct {
						Edges []interface{} `json:"edges"`
						Type  string        `json:"type"`
					} `json:"attrs"`
					TimeoutSecs  string        `json:"timeout_secs"`
					Type         string        `json:"type"`
					VariableList []interface{} `json:"variable_list"`
					Description  string        `json:"description"`
				} `json:"task_definition_list"`
				Description  string        `json:"description"`
				VariableList []interface{} `json:"variable_list"`
				Name         string        `json:"name"`
			} `json:"uninstall_runbook"`
		} `json:"options"`
		VariableList []interface{} `json:"variable_list"`
		Name         string        `json:"name"`
	} `json:"package_definition_list"`
	AppProfileList []struct {
		ActionList           []interface{} `json:"action_list"`
		DeploymentCreateList []struct {
			Name                      string        `json:"name"`
			ActionList                []interface{} `json:"action_list"`
			PackageLocalReferenceList []struct {
				Kind string `json:"kind"`
				Name string `json:"name"`
			} `json:"package_local_reference_list"`
			Description             string        `json:"description"`
			DefaultReplicas         string        `json:"default_replicas"`
			DependsOnList           []interface{} `json:"depends_on_list"`
			MaxReplicas             string        `json:"max_replicas"`
			Type                    string        `json:"type"`
			SubstrateLocalReference struct {
				Kind string `json:"kind"`
				Name string `json:"name"`
			} `json:"substrate_local_reference"`
			MinReplicas                        string        `json:"min_replicas"`
			VariableList                       []interface{} `json:"variable_list"`
			PublishedServiceLocalReferenceList []interface{} `json:"published_service_local_reference_list"`
		} `json:"deployment_create_list"`
		Name         string        `json:"name"`
		UUID		 string			`json:"uuid,omitempty"`
		VariableList []interface{} `json:"variable_list"`
		Description  string        `json:"description"`
	} `json:"app_profile_list"`
	PublishedServiceDefinitionList []interface{} `json:"published_service_definition_list"`
	Type                           string        `json:"type"`
}

// BlueprintGetSpec ...
type BlueprintGetSpec struct {
	Description 				string					`json:"description,omitempty"`
	SourceMarketplaceName		string					`json:"source_marketplace_name,omitempty"`
	SourceMarketplaceVersion	string					`json:"source_marketplace_version,omitempty"`
	AppBlueprintName			string					`json:"app_blueprint_name,omitempty"`
	EnvironmentUUID				string					`json:"environment_uuid"`
	Resources					*BlueprintGetResource		`json:"resources,omitempty"`
}

// BlueprintCloneResponse
type BlueprintCloneResponse struct {
	Metadata	*CalmMetadata	`json:"metadata,omitempty"`
}

// BlueprintGetResponse
type BlueprintGetResponse struct {
	Spec 		*BlueprintGetSpec 	`json:"spec,omitempty"`
	ApiVersion 	string			`json:"api_version,omitempty"`
	Metadata	*CalmMetadata 	`json:"metadata,omitempty"`
}

// AppBlueprintProfileRef
type AppBlueprintProfileRef struct {
	Kind 	string	`json:"kind,omitempty"`
	UUID 	string 	`json:"uuid,omitempty"`
	Name 	string	`json:"name,omitempty"`
}

// BlueprintLaunchSpec
type BlueprintLaunchSpec struct {
	Description 			string					`json:"description,omitemoty"`
	ApplicationName 		string					`json:"application_name,omitempty"`
	AppProfileReference		AppBlueprintProfileRef 	`json:"app_profile_reference"`
	Resources 				*BlueprintGetResource		`json:"resources,omitempty"`
}

// BlueprintLaunchRequest
type BlueprintLaunchRequest struct {
	Spec 		BlueprintLaunchSpec		`json:"spec,omitempty"`
	ApiVersion 	string					`json:"api_version,omitempty"`
	Metadata	*CalmMetadata 			`json:"metadata,omitempty"`
}

// BlueprintLaunchResponse
type BlueprintLaunchResponse struct {
	Status struct {
		IsCloned  bool   `json:"is_cloned"`
		RequestID string `json:"request_id"`
	} `json:"status"`
	Spec struct {
		ApplicationName     string `json:"application_name"`
		AppProfileReference struct {
			Kind string `json:"kind"`
			UUID string `json:"uuid"`
		} `json:"app_profile_reference"`
		Description string `json:"description"`
		Resources   struct {
			ServiceDefinitionList []struct {
				Singleton  bool   `json:"singleton"`
				UUID       string `json:"uuid"`
				ActionList []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
					Critical    bool   `json:"critical"`
					Runbook     struct {
						TaskDefinitionList []struct {
							TargetAnyLocalReference struct {
								Kind string `json:"kind"`
								UUID string `json:"uuid"`
								Name string `json:"name"`
							} `json:"target_any_local_reference"`
							Retries                      string        `json:"retries"`
							UUID                         string        `json:"uuid"`
							ChildTasksLocalReferenceList []interface{} `json:"child_tasks_local_reference_list"`
							Name                         string        `json:"name"`
							Attrs                        struct {
								Edges []interface{} `json:"edges"`
								Type  string        `json:"type"`
							} `json:"attrs"`
							TimeoutSecs  string        `json:"timeout_secs"`
							Type         string        `json:"type"`
							VariableList []interface{} `json:"variable_list"`
							Description  string        `json:"description"`
						} `json:"task_definition_list"`
						Description            string        `json:"description"`
						UUID                   string        `json:"uuid"`
						VariableList           []interface{} `json:"variable_list"`
						MainTaskLocalReference struct {
							Kind string `json:"kind"`
							Name string `json:"name"`
							UUID string `json:"uuid"`
						} `json:"main_task_local_reference"`
						Name string `json:"name"`
					} `json:"runbook"`
					Type string `json:"type"`
					UUID string `json:"uuid"`
				} `json:"action_list"`
				DependsOnList []interface{} `json:"depends_on_list"`
				Name          string        `json:"name"`
				PortList      []interface{} `json:"port_list"`
				Tier          string        `json:"tier"`
				VariableList  []interface{} `json:"variable_list"`
				Description   string        `json:"description"`
			} `json:"service_definition_list"`
			SubstrateDefinitionList []struct {
				Name           string        `json:"name"`
				ActionList     []interface{} `json:"action_list"`
				UUID           string        `json:"uuid"`
				ReadinessProbe struct {
					ConnectionType        string `json:"connection_type"`
					Retries               string `json:"retries"`
					ConnectionProtocol    string `json:"connection_protocol"`
					DisableReadinessProbe bool   `json:"disable_readiness_probe"`
					TimeoutSecs           string `json:"timeout_secs"`
					Address               string `json:"address"`
					DelaySecs             string `json:"delay_secs"`
					ConnectionPort        int    `json:"connection_port"`
				} `json:"readiness_probe"`
				Editables struct {
					CreateSpec struct {
						Resources struct {
							NicList struct {
								Num0 struct {
									SubnetReference bool `json:"subnet_reference"`
								} `json:"0"`
							} `json:"nic_list"`
							SerialPortList struct {
							} `json:"serial_port_list"`
							NumSockets    bool `json:"num_sockets"`
							MemorySizeMib bool `json:"memory_size_mib"`
							BootConfig    struct {
								BootDevice struct {
								} `json:"boot_device"`
							} `json:"boot_config"`
							GuestCustomization bool `json:"guest_customization"`
							ParentReference    bool `json:"parent_reference"`
							DiskList           struct {
								Num0 struct {
									DeviceProperties struct {
									} `json:"device_properties"`
									VolumeGroupReference bool `json:"volume_group_reference"`
								} `json:"0"`
							} `json:"disk_list"`
						} `json:"resources"`
					} `json:"create_spec"`
				} `json:"editables"`
				OsType     string `json:"os_type"`
				Type       string `json:"type"`
				CreateSpec struct {
					Name                      string      `json:"name"`
					AvailabilityZoneReference interface{} `json:"availability_zone_reference"`
					BackupPolicy              interface{} `json:"backup_policy"`
					Type                      string      `json:"type"`
					ClusterReference          interface{} `json:"cluster_reference"`
					Resources                 struct {
						NicList []struct {
							NicType                       string        `json:"nic_type"`
							IPEndpointList                []interface{} `json:"ip_endpoint_list"`
							NetworkFunctionChainReference interface{}   `json:"network_function_chain_reference"`
							NetworkFunctionNicType        string        `json:"network_function_nic_type"`
							MacAddress                    string        `json:"mac_address"`
							SubnetReference               struct {
								Kind string `json:"kind"`
								Type string `json:"type"`
								Name string `json:"name"`
								UUID string `json:"uuid"`
							} `json:"subnet_reference"`
							Type string `json:"type"`
						} `json:"nic_list"`
						SerialPortList        []interface{} `json:"serial_port_list"`
						GuestTools            interface{}   `json:"guest_tools"`
						HardwareClockTimezone string        `json:"hardware_clock_timezone"`
						NumVcpusPerSocket     int           `json:"num_vcpus_per_socket"`
						NumSockets            int           `json:"num_sockets"`
						GpuList               []interface{} `json:"gpu_list"`
						MemorySizeMib         int           `json:"memory_size_mib"`
						ParentReference       interface{}   `json:"parent_reference"`
						ClusterUUID           string        `json:"cluster_uuid"`
						GuestCustomization    struct {
							CloudInit struct {
								MetaData string `json:"meta_data"`
								Type     string `json:"type"`
								UserData string `json:"user_data"`
							} `json:"cloud_init"`
							Type    string      `json:"type"`
							Sysprep interface{} `json:"sysprep"`
						} `json:"guest_customization"`
						PowerState  string `json:"power_state"`
						Type        string `json:"type"`
						AccountUUID string `json:"account_uuid"`
						BootConfig  struct {
							BootDevice struct {
								Type        string `json:"type"`
								DiskAddress struct {
									Type        string `json:"type"`
									DeviceIndex int    `json:"device_index"`
									AdapterType string `json:"adapter_type"`
								} `json:"disk_address"`
							} `json:"boot_device"`
							Type       string `json:"type"`
							BootType   string `json:"boot_type"`
							MacAddress string `json:"mac_address"`
						} `json:"boot_config"`
						DiskList []struct {
							DataSourceReference struct {
								Kind string `json:"kind"`
								Type string `json:"type"`
								Name string `json:"name"`
								UUID string `json:"uuid"`
							} `json:"data_source_reference"`
							Type                 string      `json:"type"`
							DiskSizeMib          int         `json:"disk_size_mib"`
							VolumeGroupReference interface{} `json:"volume_group_reference"`
							DeviceProperties     struct {
								Type        string `json:"type"`
								DiskAddress struct {
									Type        string `json:"type"`
									DeviceIndex int    `json:"device_index"`
									AdapterType string `json:"adapter_type"`
								} `json:"disk_address"`
								DeviceType string `json:"device_type"`
							} `json:"device_properties"`
						} `json:"disk_list"`
					} `json:"resources"`
				} `json:"create_spec"`
				VariableList []interface{} `json:"variable_list"`
				Description  string        `json:"description"`
			} `json:"substrate_definition_list"`
			CredentialDefinitionList []interface{} `json:"credential_definition_list"`
			Type                     string        `json:"type"`
			AppProfileList           []struct {
				DeploymentCreateList []struct {
					UUID                      string        `json:"uuid"`
					ActionList                []interface{} `json:"action_list"`
					PackageLocalReferenceList []struct {
						Kind string `json:"kind"`
						Name string `json:"name"`
						UUID string `json:"uuid"`
					} `json:"package_local_reference_list"`
					DefaultReplicas                    string        `json:"default_replicas"`
					Name                               string        `json:"name"`
					MinReplicas                        string        `json:"min_replicas"`
					DependsOnList                      []interface{} `json:"depends_on_list"`
					PublishedServiceLocalReferenceList []interface{} `json:"published_service_local_reference_list"`
					MaxReplicas                        string        `json:"max_replicas"`
					SubstrateLocalReference            struct {
						Kind string `json:"kind"`
						Name string `json:"name"`
						UUID string `json:"uuid"`
					} `json:"substrate_local_reference"`
					Type         string        `json:"type"`
					VariableList []interface{} `json:"variable_list"`
					Description  string        `json:"description"`
				} `json:"deployment_create_list"`
				Description  string        `json:"description"`
				ActionList   []interface{} `json:"action_list"`
				Name         string        `json:"name"`
				VariableList []interface{} `json:"variable_list"`
				UUID         string        `json:"uuid"`
			} `json:"app_profile_list"`
			PublishedServiceDefinitionList []interface{} `json:"published_service_definition_list"`
			PackageDefinitionList          []struct {
				UUID                      string        `json:"uuid"`
				ActionList                []interface{} `json:"action_list"`
				ServiceLocalReferenceList []struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
					UUID string `json:"uuid"`
				} `json:"service_local_reference_list"`
				Name    string `json:"name"`
				Version string `json:"version"`
				Type    string `json:"type"`
				Options struct {
					InstallRunbook struct {
						TaskDefinitionList []struct {
							TargetAnyLocalReference struct {
								Kind string `json:"kind"`
								Name string `json:"name"`
								UUID string `json:"uuid"`
							} `json:"target_any_local_reference"`
							Retries                      string        `json:"retries"`
							Name                         string        `json:"name"`
							Description                  string        `json:"description"`
							ChildTasksLocalReferenceList []interface{} `json:"child_tasks_local_reference_list"`
							Attrs                        struct {
								Edges []interface{} `json:"edges"`
								Type  string        `json:"type"`
							} `json:"attrs"`
							TimeoutSecs  string        `json:"timeout_secs"`
							Type         string        `json:"type"`
							VariableList []interface{} `json:"variable_list"`
							UUID         string        `json:"uuid"`
						} `json:"task_definition_list"`
						Description            string `json:"description"`
						Name                   string `json:"name"`
						MainTaskLocalReference struct {
							Kind string `json:"kind"`
							Name string `json:"name"`
							UUID string `json:"uuid"`
						} `json:"main_task_local_reference"`
						VariableList []interface{} `json:"variable_list"`
						UUID         string        `json:"uuid"`
					} `json:"install_runbook"`
					Type             string `json:"type"`
					UninstallRunbook struct {
						TaskDefinitionList []struct {
							TargetAnyLocalReference struct {
								Kind string `json:"kind"`
								Name string `json:"name"`
								UUID string `json:"uuid"`
							} `json:"target_any_local_reference"`
							Retries                      string        `json:"retries"`
							Name                         string        `json:"name"`
							Description                  string        `json:"description"`
							ChildTasksLocalReferenceList []interface{} `json:"child_tasks_local_reference_list"`
							Attrs                        struct {
								Edges []interface{} `json:"edges"`
								Type  string        `json:"type"`
							} `json:"attrs"`
							TimeoutSecs  string        `json:"timeout_secs"`
							Type         string        `json:"type"`
							VariableList []interface{} `json:"variable_list"`
							UUID         string        `json:"uuid"`
						} `json:"task_definition_list"`
						Description            string `json:"description"`
						Name                   string `json:"name"`
						MainTaskLocalReference struct {
							Kind string `json:"kind"`
							Name string `json:"name"`
							UUID string `json:"uuid"`
						} `json:"main_task_local_reference"`
						VariableList []interface{} `json:"variable_list"`
						UUID         string        `json:"uuid"`
					} `json:"uninstall_runbook"`
				} `json:"options"`
				VariableList []interface{} `json:"variable_list"`
				Description  string        `json:"description"`
			} `json:"package_definition_list"`
		} `json:"resources"`
	} `json:"spec"`
	APIVersion string `json:"api_version"`
	Metadata   struct {
		LastUpdateTime       string `json:"last_update_time"`
		UseCategoriesMapping bool   `json:"use_categories_mapping"`
		Kind                 string `json:"kind"`
		Name                 string `json:"name"`
		ProjectReference     struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"project_reference"`
		SpecVersion    int    `json:"spec_version"`
		CreationTime   string `json:"creation_time"`
		OwnerReference struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"owner_reference"`
		Categories struct {
			Project string `json:"Project"`
		} `json:"categories"`
		UUID string `json:"uuid"`
	} `json:"metadata"`
}

// BlueprintGetResource
type BlueprintGetResource struct {
	ServiceDefinitionList []struct {
		Singleton  bool   `json:"singleton"`
		UUID       string `json:"uuid"`
		ActionList []struct {
			UUID     string `json:"uuid"`
			Name     string `json:"name"`
			Critical bool   `json:"critical"`
			Runbook  struct {
				TaskDefinitionList []struct {
					TargetAnyLocalReference struct {
						Kind string `json:"kind"`
						Name string `json:"name"`
						UUID string `json:"uuid"`
					} `json:"target_any_local_reference"`
					Retries                      string        `json:"retries"`
					Description                  string        `json:"description"`
					UUID                         string        `json:"uuid"`
					ChildTasksLocalReferenceList []interface{} `json:"child_tasks_local_reference_list"`
					Attrs                        struct {
						Edges []interface{} `json:"edges"`
						Type  string        `json:"type"`
					} `json:"attrs"`
					TimeoutSecs  string        `json:"timeout_secs"`
					Type         string        `json:"type"`
					VariableList []interface{} `json:"variable_list"`
					Name         string        `json:"name"`
				} `json:"task_definition_list"`
				UUID                   string        `json:"uuid"`
				Description            string        `json:"description"`
				VariableList           []interface{} `json:"variable_list"`
				MainTaskLocalReference struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
					UUID string `json:"uuid"`
				} `json:"main_task_local_reference"`
				Name string `json:"name"`
			} `json:"runbook"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"action_list"`
		Description   string        `json:"description"`
		DependsOnList []interface{} `json:"depends_on_list"`
		PortList      []interface{} `json:"port_list"`
		Tier          string        `json:"tier"`
		VariableList  []interface{} `json:"variable_list"`
		Name          string        `json:"name"`
	} `json:"service_definition_list"`
	SubstrateDefinitionList []struct {
		Description    string        `json:"description"`
		ActionList     []interface{} `json:"action_list"`
		UUID           string        `json:"uuid"`
		ReadinessProbe struct {
			ConnectionType        string `json:"connection_type"`
			Retries               string `json:"retries"`
			ConnectionProtocol    string `json:"connection_protocol"`
			DisableReadinessProbe bool   `json:"disable_readiness_probe"`
			TimeoutSecs           string `json:"timeout_secs"`
			Address               string `json:"address"`
			DelaySecs             string `json:"delay_secs"`
			ConnectionPort        int    `json:"connection_port"`
		} `json:"readiness_probe"`
		Editables struct {
			CreateSpec struct {
				Resources struct {
					NicList struct {
						Num0 struct {
							SubnetReference bool `json:"subnet_reference"`
						} `json:"0"`
					} `json:"nic_list"`
					SerialPortList struct {
					} `json:"serial_port_list"`
					NumSockets      bool `json:"num_sockets"`
					ParentReference bool `json:"parent_reference"`
					MemorySizeMib   bool `json:"memory_size_mib"`
					BootConfig      struct {
						BootDevice struct {
						} `json:"boot_device"`
					} `json:"boot_config"`
					GuestCustomization bool `json:"guest_customization"`
					DiskList           struct {
						Num0 struct {
							DeviceProperties struct {
							} `json:"device_properties"`
							VolumeGroupReference bool `json:"volume_group_reference"`
						} `json:"0"`
					} `json:"disk_list"`
				} `json:"resources"`
			} `json:"create_spec"`
		} `json:"editables"`
		OsType     string `json:"os_type"`
		Type       string `json:"type"`
		CreateSpec struct {
			Name                      string      `json:"name"`
			AvailabilityZoneReference interface{} `json:"availability_zone_reference"`
			BackupPolicy              interface{} `json:"backup_policy"`
			Type                      string      `json:"type"`
			ClusterReference          interface{} `json:"cluster_reference"`
			Resources                 struct {
				NicList []struct {
					NicType                       string        `json:"nic_type"`
					IPEndpointList                []interface{} `json:"ip_endpoint_list"`
					NetworkFunctionChainReference interface{}   `json:"network_function_chain_reference"`
					NetworkFunctionNicType        string        `json:"network_function_nic_type"`
					MacAddress                    string        `json:"mac_address"`
					SubnetReference               struct {
						Kind string `json:"kind"`
						Type string `json:"type"`
						Name string `json:"name"`
						UUID string `json:"uuid"`
					} `json:"subnet_reference"`
					Type string `json:"type"`
				} `json:"nic_list"`
				SerialPortList        []interface{} `json:"serial_port_list"`
				GuestTools            interface{}   `json:"guest_tools"`
				NumVcpusPerSocket     int           `json:"num_vcpus_per_socket"`
				NumSockets            int           `json:"num_sockets"`
				ClusterUUID           string        `json:"cluster_uuid"`
				GpuList               []interface{} `json:"gpu_list"`
				MemorySizeMib         int           `json:"memory_size_mib"`
				ParentReference       interface{}   `json:"parent_reference"`
				HardwareClockTimezone string        `json:"hardware_clock_timezone"`
				GuestCustomization    struct {
					CloudInit struct {
						MetaData string `json:"meta_data"`
						Type     string `json:"type"`
						UserData string `json:"user_data"`
					} `json:"cloud_init"`
					Type    string      `json:"type"`
					Sysprep interface{} `json:"sysprep"`
				} `json:"guest_customization"`
				PowerState  string `json:"power_state"`
				Type        string `json:"type"`
				AccountUUID string `json:"account_uuid"`
				BootConfig  struct {
					BootDevice struct {
						Type        string `json:"type"`
						DiskAddress struct {
							AdapterType string `json:"adapter_type"`
							DeviceIndex int    `json:"device_index"`
							Type        string `json:"type"`
						} `json:"disk_address"`
					} `json:"boot_device"`
					Type       string `json:"type"`
					BootType   string `json:"boot_type"`
					MacAddress string `json:"mac_address"`
				} `json:"boot_config"`
				DiskList []struct {
					DataSourceReference struct {
						Kind string `json:"kind"`
						Type string `json:"type"`
						Name string `json:"name"`
						UUID string `json:"uuid"`
					} `json:"data_source_reference"`
					Type                 string      `json:"type"`
					DiskSizeMib          int         `json:"disk_size_mib"`
					VolumeGroupReference interface{} `json:"volume_group_reference"`
					DeviceProperties     struct {
						Type        string `json:"type"`
						DeviceType  string `json:"device_type"`
						DiskAddress struct {
							AdapterType string `json:"adapter_type"`
							DeviceIndex int    `json:"device_index"`
							Type        string `json:"type"`
						} `json:"disk_address"`
					} `json:"device_properties"`
				} `json:"disk_list"`
			} `json:"resources"`
		} `json:"create_spec"`
		VariableList []interface{} `json:"variable_list"`
		Name         string        `json:"name"`
	} `json:"substrate_definition_list"`
	CredentialDefinitionList []interface{} `json:"credential_definition_list"`
	Type                     string        `json:"type"`
	AppProfileList           []struct {
		DeploymentCreateList []struct {
			UUID                      string        `json:"uuid"`
			ActionList                []interface{} `json:"action_list"`
			PackageLocalReferenceList []struct {
				Kind string `json:"kind"`
				Name string `json:"name"`
				UUID string `json:"uuid"`
			} `json:"package_local_reference_list"`
			DefaultReplicas                    string        `json:"default_replicas"`
			DependsOnList                      []interface{} `json:"depends_on_list"`
			MinReplicas                        string        `json:"min_replicas"`
			PublishedServiceLocalReferenceList []interface{} `json:"published_service_local_reference_list"`
			MaxReplicas                        string        `json:"max_replicas"`
			SubstrateLocalReference            struct {
				Kind string `json:"kind"`
				Name string `json:"name"`
				UUID string `json:"uuid"`
			} `json:"substrate_local_reference"`
			Type         string        `json:"type"`
			Name         string        `json:"name"`
			VariableList []interface{} `json:"variable_list"`
			Description  string        `json:"description"`
		} `json:"deployment_create_list"`
		UUID         string        `json:"uuid"`
		ActionList   []interface{} `json:"action_list"`
		Description  string        `json:"description"`
		VariableList []interface{} `json:"variable_list"`
		Name         string        `json:"name"`
	} `json:"app_profile_list"`
	PublishedServiceDefinitionList []interface{} `json:"published_service_definition_list"`
	PackageDefinitionList          []struct {
		Description               string        `json:"description"`
		ActionList                []interface{} `json:"action_list"`
		ServiceLocalReferenceList []struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"service_local_reference_list"`
		UUID    string `json:"uuid"`
		Version string `json:"version"`
		Type    string `json:"type"`
		Options struct {
			InstallRunbook struct {
				TaskDefinitionList []struct {
					TargetAnyLocalReference struct {
						Kind string `json:"kind"`
						Name string `json:"name"`
						UUID string `json:"uuid"`
					} `json:"target_any_local_reference"`
					Retries                      string        `json:"retries"`
					Name                         string        `json:"name"`
					UUID                         string        `json:"uuid"`
					ChildTasksLocalReferenceList []interface{} `json:"child_tasks_local_reference_list"`
					Attrs                        struct {
						Edges []interface{} `json:"edges"`
						Type  string        `json:"type"`
					} `json:"attrs"`
					TimeoutSecs  string        `json:"timeout_secs"`
					Type         string        `json:"type"`
					VariableList []interface{} `json:"variable_list"`
					Description  string        `json:"description"`
				} `json:"task_definition_list"`
				UUID                   string `json:"uuid"`
				Name                   string `json:"name"`
				MainTaskLocalReference struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
					UUID string `json:"uuid"`
				} `json:"main_task_local_reference"`
				VariableList []interface{} `json:"variable_list"`
				Description  string        `json:"description"`
			} `json:"install_runbook"`
			Type             string `json:"type"`
			UninstallRunbook struct {
				TaskDefinitionList []struct {
					TargetAnyLocalReference struct {
						Kind string `json:"kind"`
						Name string `json:"name"`
						UUID string `json:"uuid"`
					} `json:"target_any_local_reference"`
					Retries                      string        `json:"retries"`
					Name                         string        `json:"name"`
					UUID                         string        `json:"uuid"`
					ChildTasksLocalReferenceList []interface{} `json:"child_tasks_local_reference_list"`
					Attrs                        struct {
						Edges []interface{} `json:"edges"`
						Type  string        `json:"type"`
					} `json:"attrs"`
					TimeoutSecs  string        `json:"timeout_secs"`
					Type         string        `json:"type"`
					VariableList []interface{} `json:"variable_list"`
					Description  string        `json:"description"`
				} `json:"task_definition_list"`
				UUID                   string `json:"uuid"`
				Name                   string `json:"name"`
				MainTaskLocalReference struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
					UUID string `json:"uuid"`
				} `json:"main_task_local_reference"`
				VariableList []interface{} `json:"variable_list"`
				Description  string        `json:"description"`
			} `json:"uninstall_runbook"`
		} `json:"options"`
		VariableList []interface{} `json:"variable_list"`
		Name         string        `json:"name"`
	} `json:"package_definition_list"`
}

// ApplicationListResponse
type ApplicationListResponse struct {
	ApiVersion 	string	`json:"api_version,omitempty"`
	Metadata 	ListMetadataOutput `json:"metadata,omitempty"`
	Entities    []ApplicationEntity `json:"entities,omitempty"`
}

// ApplicationEntity
type ApplicationEntity struct {
	ApiVersion string `json:"api_version,omitempty"`
	Metadata	CalmMetadata `json:"metadata,omitempty"`
	Status		ApplicationStatus	`json:"status,omitempty"`
}

// ApplicationStatus
type ApplicationStatus struct {
	LastUpdateTime  int64  `json:"last_update_time"`
	Description     string `json:"description"`
	Deleted         bool   `json:"deleted"`
	CreationTime    int64  `json:"creation_time"`
	ConfigReference string `json:"config_reference"`
	UUID            string `json:"uuid"`
	Name            string `json:"name"`
	DeletionTime    int    `json:"deletion_time"`
	SpecVersion     int    `json:"spec_version"`
	SchemaVersion   string `json:"schema_version"`
	State           string `json:"state"`
	TenantUUID      string `json:"tenant_uuid"`
	Resources       struct {
		FailoverStatus   interface{} `json:"failover_status,omitempty"`
		ProtectionStatus struct {
			LastUpdateTime                   int64         `json:"last_update_time"`
			Description                      string        `json:"description"`
			UnprotectedSubstrateElementUuids []string      `json:"unprotected_substrate_element_uuids"`
			LastReplicationTime              int           `json:"last_replication_time"`
			UUID                             string        `json:"uuid"`
			State                            string        `json:"state"`
			MessageList                      []interface{} `json:"message_list"`
			RecoveryPointObjectiveSecs       int           `json:"recovery_point_objective_secs"`
			Type                             string        `json:"type"`
			Name                             string        `json:"name"`
		} `json:"protection_status"`
		SourceMarketplaceName     string `json:"source_marketplace_name"`
		AppProfileConfigReference struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"app_profile_config_reference"`
		AppBlueprintReference struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"app_blueprint_reference"`
		AppBlueprintConfigReference struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"app_blueprint_config_reference"`
		SourceMarketplaceVersion string `json:"source_marketplace_version"`
	} `json:"resources"`
}

// ApplicationEntity
type ApplicationGetResponse struct {
	ApiVersion string `json:"api_version,omitempty"`
	Metadata	CalmMetadata `json:"metadata,omitempty"`
	Status		ApplicationStatus	`json:"status,omitempty"`
}

// ApplicationDeleteResponse
type ApplicationDeleteResponse struct {
	Status struct {
		RunlogUUID    string `json:"runlog_uuid"`
		ErgonTaskUUID string `json:"ergon_task_uuid"`
	} `json:"status"`
	APIVersion string `json:"api_version"`
}