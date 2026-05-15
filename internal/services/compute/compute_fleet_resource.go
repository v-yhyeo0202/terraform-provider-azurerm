// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package compute

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/location"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/tags"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurefleet/2024-11-01/fleets"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type ComputeFleetResource struct{}

var (
	_ sdk.ResourceWithCustomizeDiff = ComputeFleetResource{}
	_ sdk.ResourceWithUpdate        = ComputeFleetResource{}
)

type ComputeFleetResourceModel struct {
	Name                           string                           `tfschema:"name"`
	ResourceGroupName              string                           `tfschema:"resource_group_name"`
	Location                       string                           `tfschema:"location"`
	VmSizesProfile                 []ComputeFleetVmSizeProfile      `tfschema:"vm_sizes_profile"`
	CreateOption                   string                           `tfschema:"create_option"`
	OsType                         string                           `tfschema:"os_type"`
	ImageReference                 []ComputeFleetImageReference     `tfschema:"image_reference"`
	AdminUsername                  string                           `tfschema:"admin_username"`
	AdminPassword                  string                           `tfschema:"admin_password"`
	ComputerNamePrefix             string                           `tfschema:"computer_name_prefix"`
	NetworkInterfaceConfigurations []ComputeFleetNicConfig          `tfschema:"network_interface_configurations"`
	NetworkApiVersion              string                           `tfschema:"network_api_version"`
	Tags                           map[string]interface{}           `tfschema:"tags"`
	GalleryApplications            []ComputeFleetGalleryApplication `tfschema:"gallery_applications"`
	Extensions                     []ComputeFleetExtension          `tfschema:"extensions"`
	CustomData                     string                           `tfschema:"custom_data"`
	// linux_configuration (MaxItems:1) flattened per rule 10
	EnableVMAgentPlatformUpdates             bool     `tfschema:"enable_vm_agent_platform_updates"`
	PatchMode                                string   `tfschema:"patch_mode"`
	BypassPlatformSafetyChecksOnUserSchedule bool     `tfschema:"bypass_platform_safety_checks_on_user_schedule"`
	RebootSetting                            string   `tfschema:"reboot_setting"`
	ProvisionVMAgent                         bool     `tfschema:"provision_vm_agent"`
	KeyData                                  []string `tfschema:"key_data"`
	// windows_configuration (MaxItems:1) flattened per rule 10; conflicting names prefixed
	EnableAutomaticUpdates                                       bool                                    `tfschema:"enable_automatic_updates"`
	WindowsConfigurationEnableVMAgentPlatformUpdates             bool                                    `tfschema:"windows_configuration_enable_vm_agent_platform_updates"`
	WindowsConfigurationPatchMode                                string                                  `tfschema:"windows_configuration_patch_mode"`
	WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule bool                                    `tfschema:"windows_configuration_bypass_platform_safety_checks_on_user_schedule"`
	WindowsConfigurationRebootSetting                            string                                  `tfschema:"windows_configuration_reboot_setting"`
	EnableHotpatching                                            bool                                    `tfschema:"enable_hotpatching"`
	WindowsConfigurationProvisionVMAgent                         bool                                    `tfschema:"windows_configuration_provision_vm_agent"`
	TimeZone                                                     string                                  `tfschema:"time_zone"`
	AdditionalUnattendContent                                    []ComputeFleetAdditionalUnattendContent `tfschema:"additional_unattend_content"`
	Listeners                                                    []ComputeFleetWinRMListener             `tfschema:"listeners"`
	Secrets                                                      []ComputeFleetSecret                    `tfschema:"secrets"`
	DataDisks                                                    []ComputeFleetDataDisk                  `tfschema:"data_disks"`
	OsDiskCaching                                                string                                  `tfschema:"os_disk_caching"`
	OsDiskDeleteOption                                           string                                  `tfschema:"os_disk_delete_option"`
	// diff_disk_settings (MaxItems:1) flattened per rule 10
	Option                        string                                     `tfschema:"option"`
	Placement                     string                                     `tfschema:"placement"`
	OsDiskSizeGB                  int64                                      `tfschema:"os_disk_size_gb"`
	OsDiskDiskEncryptionSetId     string                                     `tfschema:"os_disk_disk_encryption_set_id"`
	OsDiskSecurityEncryptionType  string                                     `tfschema:"os_disk_security_encryption_type"`
	OsDiskStorageAccountType      string                                     `tfschema:"os_disk_storage_account_type"`
	OsDiskWriteAcceleratorEnabled bool                                       `tfschema:"os_disk_write_accelerator_enabled"`
	Plan                          []ComputeFleetPlan                         `tfschema:"plan"`
	RegularPriorityProfile        []ComputeFleetRegularPriorityProfile       `tfschema:"regular_priority_profile"`
	SpotPriorityProfile           []ComputeFleetSpotPriorityProfile          `tfschema:"spot_priority_profile"`
	Identity                      []identity.ModelSystemAssignedUserAssigned `tfschema:"identity"`
	// Fleet top-level
	Zones []string `tfschema:"zones"`
	// computeProfile top-level
	ComputeApiVersion        string `tfschema:"compute_api_version"`
	PlatformFaultDomainCount int64  `tfschema:"platform_fault_domain_count"`
	// additionalVirtualMachineCapabilities (MaxItems:1) flattened per rule 10
	UltraSSDEnabled    bool `tfschema:"ultra_ssd_enabled"`
	HibernationEnabled bool `tfschema:"hibernation_enabled"`
	// diagnosticsProfile.bootDiagnostics (MaxItems:1 × 2) flattened per rule 10
	BootDiagnosticsEnabled bool   `tfschema:"enabled"`
	StorageUri             string `tfschema:"storage_uri"`
	// capacityReservation (MaxItems:1) → capacityReservationGroup (1 child) rule 9 → prefixed
	CapacityReservationGroupId string `tfschema:"capacity_reservation_group_id"`
	// baseVirtualMachineProfile direct
	LicenseType string `tfschema:"license_type"`
	UserData    string `tfschema:"user_data"`
	// extensionProfile (MaxItems:1) flattened per rule 10
	ExtensionsTimeBudget string `tfschema:"extensions_time_budget"`
	// osProfile (MaxItems:1) flattened per rule 10
	AllowExtensionOperations bool `tfschema:"allow_extension_operations"`
	// securityProfile (MaxItems:1) flattened per rule 10
	EncryptionAtHost bool `tfschema:"encryption_at_host"`
	// uefiSettings (MaxItems:1) under securityProfile, flattened per rule 10
	SecureBootEnabled bool `tfschema:"secure_boot_enabled"`
	VTpmEnabled       bool `tfschema:"v_tpm_enabled"`
	// scheduledEventsProfile → osImageNotificationProfile (MaxItems:1) flattened per rule 10
	NotBeforeTimeout string `tfschema:"not_before_timeout"`
	// scheduledEventsProfile → terminateNotificationProfile (MaxItems:1) flattened; conflict → prefixed
	TerminateNotificationProfileNotBeforeTimeout string `tfschema:"terminate_notification_profile_not_before_timeout"`
	// imageReference additional fields (MaxItems:1 already flattened)
	ImageReferenceId        string `tfschema:"image_reference_id"`
	SharedGalleryImageId    string `tfschema:"shared_gallery_image_id"`
	CommunityGalleryImageId string `tfschema:"community_gallery_image_id"`
}

type ComputeFleetImageReference struct {
	Publisher string `tfschema:"publisher"`
	Offer     string `tfschema:"offer"`
	Sku       string `tfschema:"sku"`
	Version   string `tfschema:"version"`
}

type ComputeFleetPlan struct {
	Name          string `tfschema:"name"`
	Publisher     string `tfschema:"publisher"`
	Product       string `tfschema:"product"`
	PromotionCode string `tfschema:"promotion_code"`
}

type ComputeFleetRegularPriorityProfile struct {
	Capacity           int64  `tfschema:"capacity"`
	MinCapacity        int64  `tfschema:"min_capacity"`
	AllocationStrategy string `tfschema:"allocation_strategy"`
}

type ComputeFleetSpotPriorityProfile struct {
	Capacity           int64   `tfschema:"capacity"`
	MinCapacity        int64   `tfschema:"min_capacity"`
	MaxPricePerVM      float64 `tfschema:"max_price_per_vm"`
	EvictionPolicy     string  `tfschema:"eviction_policy"`
	AllocationStrategy string  `tfschema:"allocation_strategy"`
	Maintain           bool    `tfschema:"maintain"`
}

type ComputeFleetVmSizeProfile struct {
	Name string `tfschema:"name"`
	Rank int64  `tfschema:"rank"`
}

type ComputeFleetGalleryApplication struct {
	PackageReferenceId              string `tfschema:"package_reference_id"`
	EnableAutomaticUpgrade          bool   `tfschema:"enable_automatic_upgrade"`
	ConfigurationReference          string `tfschema:"configuration_reference"`
	Order                           int64  `tfschema:"order"`
	GalleryApplicationTags          string `tfschema:"gallery_application_tags"`
	TreatFailureAsDeploymentFailure bool   `tfschema:"treat_failure_as_deployment_failure"`
}

type ComputeFleetExtension struct {
	Name                    string `tfschema:"name"`
	Publisher               string `tfschema:"publisher"`
	Type                    string `tfschema:"type"`
	TypeHandlerVersion      string `tfschema:"type_handler_version"`
	AutoUpgradeMinorVersion bool   `tfschema:"auto_upgrade_minor_version"`
	EnableAutomaticUpgrade  bool   `tfschema:"enable_automatic_upgrade"`
	ForceUpdateTag          string `tfschema:"force_update_tag"`
	ProtectedSettings       string `tfschema:"protected_settings"`
	// protected_settings_from_key_vault (MaxItems:1) flattened per rule 10
	SecretUrl                string   `tfschema:"secret_url"`
	SourceVaultId            string   `tfschema:"source_vault_id"`
	ProvisionAfterExtensions []string `tfschema:"provision_after_extensions"`
	SuppressFailures         bool     `tfschema:"suppress_failures"`
	Settings                 string   `tfschema:"settings"`
}

type ComputeFleetNicConfig struct {
	Name                        string                 `tfschema:"name"`
	IPConfigurations            []ComputeFleetIPConfig `tfschema:"ip_configurations"`
	EnableAcceleratedNetworking bool                   `tfschema:"enable_accelerated_networking"`
	AuxiliaryMode               string                 `tfschema:"auxiliary_mode"`
	AuxiliarySku                string                 `tfschema:"auxiliary_sku"`
	DeleteOption                string                 `tfschema:"delete_option"`
	DnsServers                  []string               `tfschema:"dns_servers"`
	EnableIPForwarding          bool                   `tfschema:"enable_ip_forwarding"`
	NetworkSecurityGroupId      string                 `tfschema:"network_security_group_id"`
	Primary                     bool                   `tfschema:"primary"`
}

type ComputeFleetIPConfig struct {
	Name                    string `tfschema:"name"`
	SubnetId                string `tfschema:"subnet_id"`
	Primary                 bool   `tfschema:"primary"`
	PrivateIPAddressVersion string `tfschema:"private_ip_address_version"`
	// public_ip_address_configuration (MaxItems:1) flattened per rule 10; name conflict → prefixed
	PublicIPAddressConfigurationName string `tfschema:"public_ip_address_configuration_name"`
	DeleteOption                     string `tfschema:"delete_option"`
	DomainNameLabel                  string `tfschema:"domain_name_label"`
	DomainNameLabelScope             string `tfschema:"domain_name_label_scope"`
	IdleTimeoutInMinutes             int64  `tfschema:"idle_timeout_in_minutes"`
	PublicIPAddressVersion           string `tfschema:"public_ip_address_version"`
	SkuName                          string `tfschema:"sku_name"`
	SkuTier                          string `tfschema:"sku_tier"`
}

type ComputeFleetAdditionalUnattendContent struct {
	SettingName string `tfschema:"setting_name"`
	Content     string `tfschema:"content"`
}

type ComputeFleetWinRMListener struct {
	CertificateUrl string `tfschema:"certificate_url"`
	Protocol       string `tfschema:"protocol"`
}

type ComputeFleetSecret struct {
	SourceVaultId     string                         `tfschema:"source_vault_id"`
	VaultCertificates []ComputeFleetVaultCertificate `tfschema:"vault_certificates"`
}

type ComputeFleetVaultCertificate struct {
	CertificateUrl   string `tfschema:"certificate_url"`
	CertificateStore string `tfschema:"certificate_store"`
}

type ComputeFleetDataDisk struct {
	Caching                 string `tfschema:"caching"`
	CreateOption            string `tfschema:"create_option"`
	DeleteOption            string `tfschema:"delete_option"`
	DiskSizeGB              int64  `tfschema:"disk_size_gb"`
	DiskEncryptionSetId     string `tfschema:"disk_encryption_set_id"`
	StorageAccountType      string `tfschema:"storage_account_type"`
	Lun                     int64  `tfschema:"lun"`
	WriteAcceleratorEnabled bool   `tfschema:"write_accelerator_enabled"`
}

func (r ComputeFleetResource) ResourceType() string {
	return "azurerm_compute_fleet"
}

func (r ComputeFleetResource) ModelObject() interface{} {
	return &ComputeFleetResourceModel{}
}

func (r ComputeFleetResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return fleets.ValidateFleetID
}

func (r ComputeFleetResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"resource_group_name": commonschema.ResourceGroupName(),

		"location": commonschema.Location(),

		"tags": commonschema.Tags(),

		"vm_sizes_profile": {
			Type:     pluginsdk.TypeList,
			Required: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"name": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"rank": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},
				},
			},
		},

		"create_option": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"os_type": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"image_reference": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			ForceNew: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"publisher": {
						Type:     pluginsdk.TypeString,
						Required: true,
						ForceNew: true,
					},
					"offer": {
						Type:     pluginsdk.TypeString,
						Required: true,
						ForceNew: true,
					},
					"sku": {
						Type:     pluginsdk.TypeString,
						Required: true,
						ForceNew: true,
					},
					"version": {
						Type:     pluginsdk.TypeString,
						Required: true,
						ForceNew: true,
					},
				},
			},
			ExactlyOneOf: []string{
				"image_reference",
				"image_reference_id",
				"shared_gallery_image_id",
				"community_gallery_image_id",
			},
		},

		"admin_username": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"admin_password": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"computer_name_prefix": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"network_interface_configurations": {
			Type:     pluginsdk.TypeList,
			Required: true,
			ForceNew: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"name": {
						Type:     pluginsdk.TypeString,
						Required: true,
						ForceNew: true,
					},

					"ip_configurations": {
						Type:     pluginsdk.TypeList,
						Required: true,
						ForceNew: true,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"name": {
									Type:     pluginsdk.TypeString,
									Required: true,
									ForceNew: true,
								},

								"subnet_id": {
									Type:     pluginsdk.TypeString,
									Required: true,
									ForceNew: true,
								},

								"primary": {
									Type:     pluginsdk.TypeBool,
									Optional: true,
									ForceNew: true,
									Default:  false,
								},

								"private_ip_address_version": {
									Type:     pluginsdk.TypeString,
									Optional: true,
									ForceNew: true,
								},

								"public_ip_address_configuration_name": {
									Type:     pluginsdk.TypeString,
									Optional: true,
									ForceNew: true,
								},

								"delete_option": {
									Type:     pluginsdk.TypeString,
									Optional: true,
									ForceNew: true,
								},

								"domain_name_label": {
									Type:     pluginsdk.TypeString,
									Optional: true,
									ForceNew: true,
								},

								"domain_name_label_scope": {
									Type:     pluginsdk.TypeString,
									Optional: true,
									ForceNew: true,
								},

								"idle_timeout_in_minutes": {
									Type:     pluginsdk.TypeInt,
									Optional: true,
									ForceNew: true,
								},

								"public_ip_address_version": {
									Type:     pluginsdk.TypeString,
									Optional: true,
									ForceNew: true,
								},

								"sku_name": {
									Type:     pluginsdk.TypeString,
									Optional: true,
									ForceNew: true,
								},

								"sku_tier": {
									Type:     pluginsdk.TypeString,
									Optional: true,
									ForceNew: true,
								},
							},
						},
					},

					"enable_accelerated_networking": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						ForceNew: true,
						Default:  false,
					},

					"auxiliary_mode": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},

					"auxiliary_sku": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},

					"delete_option": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},

					"dns_servers": {
						Type:     pluginsdk.TypeList,
						Optional: true,
						ForceNew: true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},

					"enable_ip_forwarding": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						ForceNew: true,
						Default:  false,
					},

					"network_security_group_id": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},

					"primary": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						ForceNew: true,
						Default:  false,
					},
				},
			},
		},

		"network_api_version": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"gallery_applications": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"package_reference_id": {
						Type:     pluginsdk.TypeString,
						Required: true,
						ForceNew: true,
					},
					"enable_automatic_upgrade": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						ForceNew: true,
						Default:  false,
					},
					"configuration_reference": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"order": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
						ForceNew: true,
					},
					"gallery_application_tags": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"treat_failure_as_deployment_failure": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						ForceNew: true,
					},
				},
			},
		},

		"extensions": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"name": {
						Type:     pluginsdk.TypeString,
						Required: true,
						ForceNew: true,
					},
					"publisher": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"type": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"type_handler_version": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"auto_upgrade_minor_version": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						ForceNew: true,
					},
					"enable_automatic_upgrade": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						ForceNew: true,
					},
					"force_update_tag": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"protected_settings": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"secret_url": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"source_vault_id": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"provision_after_extensions": {
						Type:     pluginsdk.TypeList,
						Optional: true,
						ForceNew: true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},
					"suppress_failures": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						ForceNew: true,
					},
					"settings": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
				},
			},
		},

		"custom_data": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		// linux_configuration (MaxItems:1) flattened per rule 10
		"enable_vm_agent_platform_updates": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		"patch_mode": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"bypass_platform_safety_checks_on_user_schedule": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		"reboot_setting": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"provision_vm_agent": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		// key_data: ssh.publicKeys is TypeList<{keyData}> (1 child) → rule 9 → TypeList<string>;
		// ssh has 1 child publicKeys → rule 9 again → promoted. linux_configuration MaxItems:1 → rule 10.
		"key_data": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
		},

		// windows_configuration (MaxItems:1) flattened per rule 10; conflicting names prefixed
		"enable_automatic_updates": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		"windows_configuration_enable_vm_agent_platform_updates": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		"windows_configuration_patch_mode": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"windows_configuration_bypass_platform_safety_checks_on_user_schedule": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		"windows_configuration_reboot_setting": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"enable_hotpatching": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		"windows_configuration_provision_vm_agent": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		"time_zone": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"additional_unattend_content": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"setting_name": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"content": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
				},
			},
		},

		// winRM (1 child listeners) flattened per rule 9 → "listeners"; windows_configuration MaxItems:1 → rule 10 → top level
		"listeners": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"certificate_url": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"protocol": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
				},
			},
		},

		"secrets": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"source_vault_id": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"vault_certificates": {
						Type:     pluginsdk.TypeList,
						Optional: true,
						ForceNew: true,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"certificate_url": {
									Type:     pluginsdk.TypeString,
									Required: true,
									ForceNew: true,
								},
								"certificate_store": {
									Type:     pluginsdk.TypeString,
									Optional: true,
									ForceNew: true,
								},
							},
						},
					},
				},
			},
		},

		"data_disks": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"caching": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"create_option": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"delete_option": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"disk_size_gb": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
						ForceNew: true,
					},
					"disk_encryption_set_id": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"storage_account_type": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"lun": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
						ForceNew: true,
					},
					"write_accelerator_enabled": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						ForceNew: true,
					},
				},
			},
		},

		"os_disk_caching": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"os_disk_delete_option": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		// diff_disk_settings (MaxItems:1) flattened per rule 10
		"option": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"placement": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"os_disk_size_gb": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"os_disk_disk_encryption_set_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"os_disk_security_encryption_type": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"os_disk_storage_account_type": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"os_disk_write_accelerator_enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		"plan": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"name": {
						Type:     pluginsdk.TypeString,
						Required: true,
						ForceNew: true,
					},
					"publisher": {
						Type:     pluginsdk.TypeString,
						Required: true,
						ForceNew: true,
					},
					"product": {
						Type:     pluginsdk.TypeString,
						Required: true,
						ForceNew: true,
					},
					"promotion_code": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ForceNew: true,
					},
				},
			},
		},

		"regular_priority_profile": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			AtLeastOneOf: []string{
				"regular_priority_profile",
				"spot_priority_profile",
			},
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"capacity": {
						Type:     pluginsdk.TypeInt,
						Required: true,
					},

					"min_capacity": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},

					"allocation_strategy": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						Default:  string(fleets.RegularPriorityAllocationStrategyLowestPrice),
					},
				},
			},
		},

		"spot_priority_profile": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			AtLeastOneOf: []string{
				"regular_priority_profile",
				"spot_priority_profile",
			},
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"capacity": {
						Type:     pluginsdk.TypeInt,
						Required: true,
					},

					"min_capacity": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},

					"max_price_per_vm": {
						Type:     pluginsdk.TypeFloat,
						Optional: true,
					},

					"eviction_policy": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						Default:  string(fleets.EvictionPolicyDelete),
					},

					"allocation_strategy": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						Default:  string(fleets.SpotAllocationStrategyPriceCapacityOptimized),
					},

					"maintain": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						Default:  true,
					},
				},
			},
		},

		"identity": commonschema.SystemAssignedUserAssignedIdentityOptional(),

		"zones": commonschema.ZonesMultipleOptionalForceNew(),

		// computeProfile top-level
		"compute_api_version": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},

		"platform_fault_domain_count": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
			ForceNew: true,
			Default:  1,
		},

		// additionalVirtualMachineCapabilities (MaxItems:1) flattened per rule 10
		"ultra_ssd_enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		"hibernation_enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		// diagnosticsProfile.bootDiagnostics (MaxItems:1 × 2) flattened per rule 10
		"enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		"storage_uri": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		// capacityReservation (MaxItems:1) → capacityReservationGroup (1 child id) rule 9 → prefixed
		"capacity_reservation_group_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		// baseVirtualMachineProfile direct
		"license_type": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"user_data": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		// extensionProfile (MaxItems:1) flattened per rule 10
		"extensions_time_budget": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		// osProfile (MaxItems:1) flattened per rule 10
		"allow_extension_operations": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		// securityProfile (MaxItems:1) flattened per rule 10
		"encryption_at_host": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		// uefiSettings (MaxItems:1) under securityProfile, flattened per rule 10
		"secure_boot_enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		"v_tpm_enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},

		// scheduledEventsProfile → osImageNotificationProfile (MaxItems:1) flattened per rule 10
		"not_before_timeout": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		// scheduledEventsProfile → terminateNotificationProfile (MaxItems:1) flattened; conflict → prefixed
		"terminate_notification_profile_not_before_timeout": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
		},

		// imageReference additional fields (MaxItems:1 already flattened)
		"image_reference_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
			ExactlyOneOf: []string{
				"image_reference",
				"image_reference_id",
				"shared_gallery_image_id",
				"community_gallery_image_id",
			},
		},

		"shared_gallery_image_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
			ExactlyOneOf: []string{
				"image_reference",
				"image_reference_id",
				"shared_gallery_image_id",
				"community_gallery_image_id",
			},
		},

		"community_gallery_image_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
			ExactlyOneOf: []string{
				"image_reference",
				"image_reference_id",
				"shared_gallery_image_id",
				"community_gallery_image_id",
			},
		},
	}
}

func (r ComputeFleetResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ComputeFleetResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Compute.FleetsClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			var config ComputeFleetResourceModel
			if err := metadata.Decode(&config); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			id := fleets.NewFleetID(subscriptionId, config.ResourceGroupName, config.Name)

			existing, err := client.Get(ctx, id)
			if err != nil {
				if !response.WasNotFound(existing.HttpResponse) {
					return fmt.Errorf("checking for the presence of an existing %s: %+v", id, err)
				}
			}
			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			parameters := fleets.Fleet{
				Location: location.Normalize(config.Location),
				Tags:     tags.Expand(config.Tags),
				Properties: &fleets.FleetProperties{
					VMSizesProfile: expandComputeFleetVmSizesProfile(config.VmSizesProfile, metadata),
					ComputeProfile: expandComputeFleetComputeProfile(config, metadata),
				},
			}

			if len(config.Zones) > 0 {
				z := zones.Schema(config.Zones)
				parameters.Zones = &z
			}

			expandedIdentity, err := identity.ExpandLegacySystemAndUserAssignedMapFromModel(config.Identity)
			if err != nil {
				return fmt.Errorf("expanding `identity`: %+v", err)
			}
			parameters.Identity = expandedIdentity

			if rpp := expandComputeFleetRegularPriorityProfile(config, metadata); rpp != nil {
				parameters.Properties.RegularPriorityProfile = rpp
			}

			if spp := expandComputeFleetSpotPriorityProfile(config, metadata); spp != nil {
				parameters.Properties.SpotPriorityProfile = spp
			}

			if plan := expandComputeFleetPlan(config.Plan); plan != nil {
				parameters.Plan = plan
			}
			if err = client.CreateOrUpdateThenPoll(ctx, id, parameters); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ComputeFleetResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Compute.FleetsClient

			id, err := fleets.ParseFleetID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(*id)
				}
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}

			schema := ComputeFleetResourceModel{
				Name:              id.FleetName,
				ResourceGroupName: id.ResourceGroupName,
			}

			if model := resp.Model; model != nil {
				schema.Location = location.Normalize(model.Location)
				schema.Tags = tags.Flatten(model.Tags)
				schema.Plan = flattenComputeFleetPlan(model.Plan)
				flattenedIdentity, err := identity.FlattenLegacySystemAndUserAssignedMapToModel(model.Identity)
				if err != nil {
					return fmt.Errorf("flattening `identity`: %+v", err)
				}
				schema.Identity = flattenedIdentity
				if model.Zones != nil {
					schema.Zones = zones.Flatten(model.Zones)
				}

				if props := model.Properties; props != nil {
					schema.VmSizesProfile = flattenComputeFleetVmSizesProfile(props.VMSizesProfile)
					flattenComputeFleetRegularPriorityProfile(props.RegularPriorityProfile, &schema)
					flattenComputeFleetSpotPriorityProfile(props.SpotPriorityProfile, &schema)

					vmProfile := props.ComputeProfile.BaseVirtualMachineProfile

					if sp := vmProfile.StorageProfile; sp != nil {
						if sp.ImageReference != nil {
							imgRef := ComputeFleetImageReference{}
							if sp.ImageReference.Publisher != nil {
								imgRef.Publisher = *sp.ImageReference.Publisher
							}
							if sp.ImageReference.Offer != nil {
								imgRef.Offer = *sp.ImageReference.Offer
							}
							if sp.ImageReference.Sku != nil {
								imgRef.Sku = *sp.ImageReference.Sku
							}
							if sp.ImageReference.Version != nil {
								imgRef.Version = *sp.ImageReference.Version
							}
							if imgRef.Publisher != "" || imgRef.Offer != "" || imgRef.Sku != "" || imgRef.Version != "" {
								schema.ImageReference = []ComputeFleetImageReference{imgRef}
							}
							if sp.ImageReference.Id != nil {
								schema.ImageReferenceId = *sp.ImageReference.Id
							}
							if sp.ImageReference.SharedGalleryImageId != nil {
								schema.SharedGalleryImageId = *sp.ImageReference.SharedGalleryImageId
							}
							if sp.ImageReference.CommunityGalleryImageId != nil {
								schema.CommunityGalleryImageId = *sp.ImageReference.CommunityGalleryImageId
							}
						}
						if osDisk := sp.OsDisk; osDisk != nil {
							schema.CreateOption = string(osDisk.CreateOption)
							if osDisk.OsType != nil {
								schema.OsType = string(*osDisk.OsType)
							}
							if osDisk.ManagedDisk != nil {
								if osDisk.ManagedDisk.DiskEncryptionSet != nil && osDisk.ManagedDisk.DiskEncryptionSet.Id != nil {
									schema.OsDiskDiskEncryptionSetId = *osDisk.ManagedDisk.DiskEncryptionSet.Id
								}
								if osDisk.ManagedDisk.StorageAccountType != nil {
									schema.OsDiskStorageAccountType = string(*osDisk.ManagedDisk.StorageAccountType)
								}
								if osDisk.ManagedDisk.SecurityProfile != nil && osDisk.ManagedDisk.SecurityProfile.SecurityEncryptionType != nil {
									schema.OsDiskSecurityEncryptionType = string(*osDisk.ManagedDisk.SecurityProfile.SecurityEncryptionType)
								}
							}
							if osDisk.Caching != nil {
								schema.OsDiskCaching = string(*osDisk.Caching)
							}
							if osDisk.DeleteOption != nil {
								schema.OsDiskDeleteOption = string(*osDisk.DeleteOption)
							}
							if osDisk.DiskSizeGB != nil {
								schema.OsDiskSizeGB = *osDisk.DiskSizeGB
							}
							if osDisk.WriteAcceleratorEnabled != nil {
								schema.OsDiskWriteAcceleratorEnabled = *osDisk.WriteAcceleratorEnabled
							}
							if osDisk.DiffDiskSettings != nil {
								if osDisk.DiffDiskSettings.Option != nil {
									schema.Option = string(*osDisk.DiffDiskSettings.Option)
								}
								if osDisk.DiffDiskSettings.Placement != nil {
									schema.Placement = string(*osDisk.DiffDiskSettings.Placement)
								}
							}
						}
						schema.DataDisks = flattenComputeFleetDataDisks(sp.DataDisks)
					}

					if np := vmProfile.NetworkProfile; np != nil {
						schema.NetworkInterfaceConfigurations = flattenComputeFleetNicConfigs(np.NetworkInterfaceConfigurations)
						if np.NetworkApiVersion != nil {
							schema.NetworkApiVersion = string(*np.NetworkApiVersion)
						}
					}

					if op := vmProfile.OsProfile; op != nil {
						if op.AdminUsername != nil {
							schema.AdminUsername = *op.AdminUsername
						}
						schema.AdminPassword = metadata.ResourceData.Get("admin_password").(string)
						if op.ComputerNamePrefix != nil {
							schema.ComputerNamePrefix = *op.ComputerNamePrefix
						}
						if op.CustomData != nil {
							schema.CustomData = metadata.ResourceData.Get("custom_data").(string)
						}
						if op.AllowExtensionOperations != nil {
							schema.AllowExtensionOperations = *op.AllowExtensionOperations
						}
						flattenComputeFleetLinuxConfiguration(op.LinuxConfiguration, &schema)
						flattenComputeFleetWindowsConfiguration(op.WindowsConfiguration, &schema)
						schema.Secrets = flattenComputeFleetSecrets(op.Secrets)
					}

					if ap := vmProfile.ApplicationProfile; ap != nil {
						if ap.GalleryApplications != nil {
							schema.GalleryApplications = flattenComputeFleetGalleryApplications(*ap.GalleryApplications)
						}
					}

					if ep := vmProfile.ExtensionProfile; ep != nil {
						schema.Extensions = flattenComputeFleetExtensions(ep.Extensions)
						if ep.ExtensionsTimeBudget != nil {
							schema.ExtensionsTimeBudget = *ep.ExtensionsTimeBudget
						}
					}

					if vmProfile.LicenseType != nil {
						schema.LicenseType = *vmProfile.LicenseType
					}
					if vmProfile.UserData != nil {
						schema.UserData = *vmProfile.UserData
					}

					if dp := vmProfile.DiagnosticsProfile; dp != nil && dp.BootDiagnostics != nil {
						if dp.BootDiagnostics.Enabled != nil {
							schema.BootDiagnosticsEnabled = *dp.BootDiagnostics.Enabled
						}
						if dp.BootDiagnostics.StorageUri != nil {
							schema.StorageUri = *dp.BootDiagnostics.StorageUri
						}
					}

					if cr := vmProfile.CapacityReservation; cr != nil && cr.CapacityReservationGroup != nil && cr.CapacityReservationGroup.Id != nil {
						schema.CapacityReservationGroupId = *cr.CapacityReservationGroup.Id
					}

					if sp2 := vmProfile.SecurityProfile; sp2 != nil {
						if sp2.EncryptionAtHost != nil {
							schema.EncryptionAtHost = *sp2.EncryptionAtHost
						}
						if sp2.UefiSettings != nil {
							if sp2.UefiSettings.SecureBootEnabled != nil {
								schema.SecureBootEnabled = *sp2.UefiSettings.SecureBootEnabled
							}
							if sp2.UefiSettings.VTpmEnabled != nil {
								schema.VTpmEnabled = *sp2.UefiSettings.VTpmEnabled
							}
						}
					}

					if sep := vmProfile.ScheduledEventsProfile; sep != nil {
						if sep.OsImageNotificationProfile != nil && sep.OsImageNotificationProfile.NotBeforeTimeout != nil {
							schema.NotBeforeTimeout = *sep.OsImageNotificationProfile.NotBeforeTimeout
						}
						if sep.TerminateNotificationProfile != nil && sep.TerminateNotificationProfile.NotBeforeTimeout != nil {
							schema.TerminateNotificationProfileNotBeforeTimeout = *sep.TerminateNotificationProfile.NotBeforeTimeout
						}
					}

					if props.ComputeProfile.ComputeApiVersion != nil {
						schema.ComputeApiVersion = *props.ComputeProfile.ComputeApiVersion
					}
					if props.ComputeProfile.PlatformFaultDomainCount != nil {
						schema.PlatformFaultDomainCount = *props.ComputeProfile.PlatformFaultDomainCount
					}
					if ac := props.ComputeProfile.AdditionalVirtualMachineCapabilities; ac != nil {
						if ac.UltraSSDEnabled != nil {
							schema.UltraSSDEnabled = *ac.UltraSSDEnabled
						}
						if ac.HibernationEnabled != nil {
							schema.HibernationEnabled = *ac.HibernationEnabled
						}
					}
				}
			}

			return metadata.Encode(&schema)
		},
	}
}

func (r ComputeFleetResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Compute.FleetsClient

			id, err := fleets.ParseFleetID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			existing, err := client.Get(ctx, *id)
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			if existing.Model == nil {
				return fmt.Errorf("retrieving %s: `model` was nil", id)
			}

			payload := *existing.Model

			var config ComputeFleetResourceModel
			if err := metadata.Decode(&config); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			baseVirtualMachineProfile := payload.Properties.ComputeProfile.BaseVirtualMachineProfile
			if osProfile := baseVirtualMachineProfile.OsProfile; osProfile != nil && config.AdminPassword != "" {
				osProfile.AdminPassword = pointer.To(config.AdminPassword)
			}

			if metadata.ResourceData.HasChange("tags") {
				payload.Tags = tags.Expand(config.Tags)
			}

			if metadata.ResourceData.HasChange("identity") {
				expandedIdentity, err := identity.ExpandLegacySystemAndUserAssignedMapFromModel(config.Identity)
				if err != nil {
					return fmt.Errorf("expanding `identity`: %+v", err)
				}
				payload.Identity = expandedIdentity
			}

			if metadata.ResourceData.HasChange("zones") {
				if len(config.Zones) > 0 {
					z := zones.Schema(config.Zones)
					payload.Zones = &z
				} else {
					payload.Zones = nil
				}
			}

			if metadata.ResourceData.HasChange("vm_sizes_profile") ||
				metadata.ResourceData.HasChange("regular_priority_profile") ||
				metadata.ResourceData.HasChange("spot_priority_profile") {
				if payload.Properties == nil {
					payload.Properties = &fleets.FleetProperties{}
				}
				if metadata.ResourceData.HasChange("vm_sizes_profile") {
					payload.Properties.VMSizesProfile = expandComputeFleetVmSizesProfile(config.VmSizesProfile, metadata)
				}
				if metadata.ResourceData.HasChange("regular_priority_profile") {
					payload.Properties.RegularPriorityProfile = expandComputeFleetRegularPriorityProfile(config, metadata)
				}
				if metadata.ResourceData.HasChange("spot_priority_profile") {
					payload.Properties.SpotPriorityProfile = expandComputeFleetSpotPriorityProfile(config, metadata)
				}
			}

			if err = client.CreateOrUpdateThenPoll(ctx, *id, payload); err != nil {
				return fmt.Errorf("updating %s: %+v", *id, err)
			}

			return nil
		},
	}
}

func (r ComputeFleetResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Compute.FleetsClient

			id, err := fleets.ParseFleetID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			if err = client.DeleteThenPoll(ctx, *id); err != nil {
				return fmt.Errorf("deleting %s: %+v", *id, err)
			}

			return nil
		},
	}
}

func expandComputeFleetVmSizesProfile(input []ComputeFleetVmSizeProfile, metadata sdk.ResourceMetaData) []fleets.VMSizeProfile {
	result := make([]fleets.VMSizeProfile, 0, len(input))
	for i, v := range input {
		item := fleets.VMSizeProfile{
			Name: v.Name,
		}

		if !metadata.ResourceData.GetRawConfig().AsValueMap()["vm_sizes_profile"].AsValueSlice()[i].AsValueMap()["rank"].IsNull() {
			rank := v.Rank
			item.Rank = &rank
		}

		result = append(result, item)
	}
	return result
}

func (r ComputeFleetResource) CustomizeDiff() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			var config ComputeFleetResourceModel
			if err := metadata.DecodeDiff(&config); err != nil {
				return fmt.Errorf("DecodeDiff: %+v", err)
			}

			forceNewRegularPriorityProfileProperties := []string{
				"min_capacity",
				"allocation_strategy",
			}
			oldRegularPriorityProfileProperty, _ := metadata.ResourceDiff.GetChange("regular_priority_profile")
			if len(oldRegularPriorityProfileProperty.([]interface{})) > 0 {
				if len(config.RegularPriorityProfile) == 0 {
					metadata.ResourceDiff.ForceNew("regular_priority_profile")
				}

				for _, regularPriorityProfileProperty := range forceNewRegularPriorityProfileProperties {
					regularPriorityProfileProperty = fmt.Sprintf("regular_priority_profile.%s", regularPriorityProfileProperty)
					if metadata.ResourceDiff.HasChange(regularPriorityProfileProperty) {
						metadata.ResourceDiff.ForceNew(regularPriorityProfileProperty)
					}
				}
			}

			forceNewSpotPriorityProfileProperties := []string{
				"min_capacity",
				"max_price_per_vm",
				"eviction_policy",
				"allocation_strategy",
				"maintain",
			}
			oldSpotPriorityProfileProperty, _ := metadata.ResourceDiff.GetChange("spot_priority_profile")
			if len(oldSpotPriorityProfileProperty.([]interface{})) > 0 {
				if len(config.SpotPriorityProfile) == 0 {
					metadata.ResourceDiff.ForceNew("spot_priority_profile")
				}

				for _, spotPriorityProfileProperty := range forceNewSpotPriorityProfileProperties {
					spotPriorityProfileProperty = fmt.Sprintf("spot_priority_profile.%s", spotPriorityProfileProperty)
					if metadata.ResourceDiff.HasChange(spotPriorityProfileProperty) {
						metadata.ResourceDiff.ForceNew(spotPriorityProfileProperty)
					}
				}
			}

			return nil
		},
	}
}

func flattenComputeFleetVmSizesProfile(input []fleets.VMSizeProfile) []ComputeFleetVmSizeProfile {
	result := make([]ComputeFleetVmSizeProfile, 0, len(input))
	for _, v := range input {
		item := ComputeFleetVmSizeProfile{Name: v.Name}
		if v.Rank != nil {
			item.Rank = *v.Rank
		}
		result = append(result, item)
	}
	return result
}

func expandComputeFleetComputeProfile(config ComputeFleetResourceModel, metadata sdk.ResourceMetaData) fleets.ComputeProfile {
	osDisk := &fleets.VirtualMachineScaleSetOSDisk{
		CreateOption: fleets.DiskCreateOptionTypes(config.CreateOption),
	}
	if config.OsType != "" {
		osTypeVal := fleets.OperatingSystemTypes(config.OsType)
		osDisk.OsType = &osTypeVal
	}
	if config.OsDiskCaching != "" {
		c := fleets.CachingTypes(config.OsDiskCaching)
		osDisk.Caching = &c
	}
	if config.OsDiskDeleteOption != "" {
		d := fleets.DiskDeleteOptionTypes(config.OsDiskDeleteOption)
		osDisk.DeleteOption = &d
	}
	if !metadata.ResourceData.GetRawConfig().AsValueMap()["os_disk_size_gb"].IsNull() {
		sz := config.OsDiskSizeGB
		osDisk.DiskSizeGB = &sz
	}
	if config.OsDiskWriteAcceleratorEnabled {
		osDisk.WriteAcceleratorEnabled = &config.OsDiskWriteAcceleratorEnabled
	}
	if config.Option != "" || config.Placement != "" {
		dds := &fleets.DiffDiskSettings{}
		if config.Option != "" {
			o := fleets.DiffDiskOptions(config.Option)
			dds.Option = &o
		}
		if config.Placement != "" {
			p := fleets.DiffDiskPlacement(config.Placement)
			dds.Placement = &p
		}
		osDisk.DiffDiskSettings = dds
	}
	if osDisk.ManagedDisk == nil {
		osDisk.ManagedDisk = &fleets.VirtualMachineScaleSetManagedDiskParameters{}
	}
	if config.OsDiskDiskEncryptionSetId != "" || config.OsDiskStorageAccountType != "" || config.OsDiskSecurityEncryptionType != "" {
		if osDisk.ManagedDisk == nil {
			osDisk.ManagedDisk = &fleets.VirtualMachineScaleSetManagedDiskParameters{}
		}
		if config.OsDiskDiskEncryptionSetId != "" {
			osDisk.ManagedDisk.DiskEncryptionSet = &fleets.DiskEncryptionSetParameters{Id: &config.OsDiskDiskEncryptionSetId}
		}
		if config.OsDiskStorageAccountType != "" {
			sat := fleets.StorageAccountTypes(config.OsDiskStorageAccountType)
			osDisk.ManagedDisk.StorageAccountType = &sat
		}
		if config.OsDiskSecurityEncryptionType != "" {
			sp := &fleets.VMDiskSecurityProfile{}
			set := fleets.SecurityEncryptionTypes(config.OsDiskSecurityEncryptionType)
			sp.SecurityEncryptionType = &set
			osDisk.ManagedDisk.SecurityProfile = sp
		}
	}

	var imageRef *fleets.ImageReference
	if len(config.ImageReference) > 0 || config.ImageReferenceId != "" || config.SharedGalleryImageId != "" || config.CommunityGalleryImageId != "" {
		imageRef = &fleets.ImageReference{}
		if len(config.ImageReference) > 0 {
			if config.ImageReference[0].Publisher != "" {
				imageRef.Publisher = &config.ImageReference[0].Publisher
			}
			if config.ImageReference[0].Offer != "" {
				imageRef.Offer = &config.ImageReference[0].Offer
			}
			if config.ImageReference[0].Sku != "" {
				imageRef.Sku = &config.ImageReference[0].Sku
			}
			if config.ImageReference[0].Version != "" {
				imageRef.Version = &config.ImageReference[0].Version
			}
		}
		if config.ImageReferenceId != "" {
			imageRef.Id = &config.ImageReferenceId
		}
		if config.SharedGalleryImageId != "" {
			imageRef.SharedGalleryImageId = &config.SharedGalleryImageId
		}
		if config.CommunityGalleryImageId != "" {
			imageRef.CommunityGalleryImageId = &config.CommunityGalleryImageId
		}
	}

	np := expandComputeFleetNetworkProfile(config.NetworkInterfaceConfigurations, config.NetworkApiVersion, metadata)

	osProfile := &fleets.VirtualMachineScaleSetOSProfile{}
	if config.AdminUsername != "" {
		osProfile.AdminUsername = &config.AdminUsername
	}
	if config.AdminPassword != "" {
		osProfile.AdminPassword = &config.AdminPassword
	}
	if config.ComputerNamePrefix != "" {
		osProfile.ComputerNamePrefix = &config.ComputerNamePrefix
	}
	if config.CustomData != "" {
		osProfile.CustomData = &config.CustomData
	}
	if config.AllowExtensionOperations {
		osProfile.AllowExtensionOperations = &config.AllowExtensionOperations
	}
	if config.EnableVMAgentPlatformUpdates || config.PatchMode != "" || config.BypassPlatformSafetyChecksOnUserSchedule || config.RebootSetting != "" || config.ProvisionVMAgent || len(config.KeyData) > 0 {
		osProfile.LinuxConfiguration = expandComputeFleetLinuxConfiguration(config)
	}
	if config.EnableAutomaticUpdates || config.WindowsConfigurationEnableVMAgentPlatformUpdates || config.WindowsConfigurationPatchMode != "" || config.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule || config.WindowsConfigurationRebootSetting != "" || config.EnableHotpatching || config.WindowsConfigurationProvisionVMAgent || config.TimeZone != "" || len(config.AdditionalUnattendContent) > 0 || len(config.Listeners) > 0 {
		osProfile.WindowsConfiguration = expandComputeFleetWindowsConfiguration(config)
	}
	if len(config.Secrets) > 0 {
		osProfile.Secrets = expandComputeFleetSecrets(config.Secrets)
	}

	sp := &fleets.VirtualMachineScaleSetStorageProfile{
		OsDisk:         osDisk,
		ImageReference: imageRef,
	}
	if len(config.DataDisks) > 0 {
		sp.DataDisks = expandComputeFleetDataDisks(config.DataDisks, metadata)
	}

	vmProfile := fleets.BaseVirtualMachineProfile{
		StorageProfile: sp,
		NetworkProfile: np,
		OsProfile:      osProfile,
	}

	if len(config.GalleryApplications) > 0 {
		vmProfile.ApplicationProfile = &fleets.ApplicationProfile{
			GalleryApplications: func() *[]fleets.VMGalleryApplication {
				v := expandComputeFleetGalleryApplications(config.GalleryApplications, metadata)
				return &v
			}(),
		}
	}

	if len(config.Extensions) > 0 || config.ExtensionsTimeBudget != "" {
		vmProfile.ExtensionProfile = &fleets.VirtualMachineScaleSetExtensionProfile{}
		if len(config.Extensions) > 0 {
			vmProfile.ExtensionProfile.Extensions = expandComputeFleetExtensions(config.Extensions)
		}
		if config.ExtensionsTimeBudget != "" {
			vmProfile.ExtensionProfile.ExtensionsTimeBudget = &config.ExtensionsTimeBudget
		}
	}

	if config.LicenseType != "" {
		vmProfile.LicenseType = &config.LicenseType
	}
	if config.UserData != "" {
		vmProfile.UserData = &config.UserData
	}

	if config.BootDiagnosticsEnabled || config.StorageUri != "" {
		vmProfile.DiagnosticsProfile = &fleets.DiagnosticsProfile{
			BootDiagnostics: &fleets.BootDiagnostics{
				Enabled:    &config.BootDiagnosticsEnabled,
				StorageUri: &config.StorageUri,
			},
		}
	}

	if config.CapacityReservationGroupId != "" {
		vmProfile.CapacityReservation = &fleets.CapacityReservationProfile{
			CapacityReservationGroup: &fleets.SubResource{Id: &config.CapacityReservationGroupId},
		}
	}

	if config.EncryptionAtHost || config.SecureBootEnabled || config.VTpmEnabled {
		vmProfile.SecurityProfile = &fleets.SecurityProfile{}
		if config.EncryptionAtHost {
			vmProfile.SecurityProfile.EncryptionAtHost = &config.EncryptionAtHost
		}
		if config.SecureBootEnabled || config.VTpmEnabled {
			vmProfile.SecurityProfile.UefiSettings = &fleets.UefiSettings{}
			if config.SecureBootEnabled {
				vmProfile.SecurityProfile.UefiSettings.SecureBootEnabled = &config.SecureBootEnabled
			}
			if config.VTpmEnabled {
				vmProfile.SecurityProfile.UefiSettings.VTpmEnabled = &config.VTpmEnabled
			}
		}
	}

	if config.NotBeforeTimeout != "" || config.TerminateNotificationProfileNotBeforeTimeout != "" {
		vmProfile.ScheduledEventsProfile = &fleets.ScheduledEventsProfile{}
		if config.NotBeforeTimeout != "" {
			vmProfile.ScheduledEventsProfile.OsImageNotificationProfile = &fleets.OSImageNotificationProfile{
				NotBeforeTimeout: &config.NotBeforeTimeout,
			}
		}
		if config.TerminateNotificationProfileNotBeforeTimeout != "" {
			vmProfile.ScheduledEventsProfile.TerminateNotificationProfile = &fleets.TerminateNotificationProfile{
				NotBeforeTimeout: &config.TerminateNotificationProfileNotBeforeTimeout,
			}
		}
	}

	computeProfile := fleets.ComputeProfile{
		BaseVirtualMachineProfile: vmProfile,
	}
	if config.ComputeApiVersion != "" {
		computeProfile.ComputeApiVersion = &config.ComputeApiVersion
	}
	if !metadata.ResourceData.GetRawConfig().AsValueMap()["platform_fault_domain_count"].IsNull() {
		computeProfile.PlatformFaultDomainCount = &config.PlatformFaultDomainCount
	}
	if config.UltraSSDEnabled || config.HibernationEnabled {
		computeProfile.AdditionalVirtualMachineCapabilities = &fleets.AdditionalCapabilities{}
		if config.UltraSSDEnabled {
			computeProfile.AdditionalVirtualMachineCapabilities.UltraSSDEnabled = &config.UltraSSDEnabled
		}
		if config.HibernationEnabled {
			computeProfile.AdditionalVirtualMachineCapabilities.HibernationEnabled = &config.HibernationEnabled
		}
	}
	return computeProfile
}

func expandComputeFleetGalleryApplications(input []ComputeFleetGalleryApplication, metadata sdk.ResourceMetaData) []fleets.VMGalleryApplication {
	result := make([]fleets.VMGalleryApplication, 0, len(input))
	for i, v := range input {
		item := fleets.VMGalleryApplication{
			PackageReferenceId: v.PackageReferenceId,
		}
		if v.EnableAutomaticUpgrade {
			item.EnableAutomaticUpgrade = &v.EnableAutomaticUpgrade
		}
		if v.ConfigurationReference != "" {
			item.ConfigurationReference = &v.ConfigurationReference
		}
		if !metadata.ResourceData.GetRawConfig().AsValueMap()["gallery_applications"].AsValueSlice()[i].AsValueMap()["order"].IsNull() {
			order := v.Order
			item.Order = &order
		}
		if v.GalleryApplicationTags != "" {
			item.Tags = &v.GalleryApplicationTags
		}
		if v.TreatFailureAsDeploymentFailure {
			item.TreatFailureAsDeploymentFailure = &v.TreatFailureAsDeploymentFailure
		}
		result = append(result, item)
	}
	return result
}

func flattenComputeFleetGalleryApplications(input []fleets.VMGalleryApplication) []ComputeFleetGalleryApplication {
	result := make([]ComputeFleetGalleryApplication, 0, len(input))
	for _, v := range input {
		item := ComputeFleetGalleryApplication{
			PackageReferenceId: v.PackageReferenceId,
		}
		if v.EnableAutomaticUpgrade != nil {
			item.EnableAutomaticUpgrade = *v.EnableAutomaticUpgrade
		}
		if v.ConfigurationReference != nil {
			item.ConfigurationReference = *v.ConfigurationReference
		}
		if v.Order != nil {
			item.Order = int64(*v.Order)
		}
		if v.Tags != nil {
			item.GalleryApplicationTags = *v.Tags
		}
		if v.TreatFailureAsDeploymentFailure != nil {
			item.TreatFailureAsDeploymentFailure = *v.TreatFailureAsDeploymentFailure
		}
		result = append(result, item)
	}
	return result
}

func expandComputeFleetExtensions(input []ComputeFleetExtension) *[]fleets.VirtualMachineScaleSetExtension {
	result := make([]fleets.VirtualMachineScaleSetExtension, 0, len(input))
	for _, v := range input {
		ext := fleets.VirtualMachineScaleSetExtension{}
		if v.Name != "" {
			ext.Name = &v.Name
		}
		props := &fleets.VirtualMachineScaleSetExtensionProperties{}
		if v.Publisher != "" {
			props.Publisher = &v.Publisher
		}
		if v.Type != "" {
			props.Type = &v.Type
		}
		if v.TypeHandlerVersion != "" {
			props.TypeHandlerVersion = &v.TypeHandlerVersion
		}
		if v.AutoUpgradeMinorVersion {
			props.AutoUpgradeMinorVersion = &v.AutoUpgradeMinorVersion
		}
		if v.EnableAutomaticUpgrade {
			props.EnableAutomaticUpgrade = &v.EnableAutomaticUpgrade
		}
		if v.ForceUpdateTag != "" {
			props.ForceUpdateTag = &v.ForceUpdateTag
		}
		if v.SuppressFailures {
			props.SuppressFailures = &v.SuppressFailures
		}
		if len(v.ProvisionAfterExtensions) > 0 {
			props.ProvisionAfterExtensions = &v.ProvisionAfterExtensions
		}
		if v.SecretUrl != "" || v.SourceVaultId != "" {
			props.ProtectedSettingsFromKeyVault = &fleets.KeyVaultSecretReference{
				SecretURL:   v.SecretUrl,
				SourceVault: fleets.SubResource{Id: &v.SourceVaultId},
			}
		}
		ext.Properties = props
		result = append(result, ext)
	}
	return &result
}

func flattenComputeFleetExtensions(input *[]fleets.VirtualMachineScaleSetExtension) []ComputeFleetExtension {
	if input == nil {
		return nil
	}
	result := make([]ComputeFleetExtension, 0, len(*input))
	for _, v := range *input {
		item := ComputeFleetExtension{}
		if v.Name != nil {
			item.Name = *v.Name
		}
		if p := v.Properties; p != nil {
			if p.Publisher != nil {
				item.Publisher = *p.Publisher
			}
			if p.Type != nil {
				item.Type = *p.Type
			}
			if p.TypeHandlerVersion != nil {
				item.TypeHandlerVersion = *p.TypeHandlerVersion
			}
			if p.AutoUpgradeMinorVersion != nil {
				item.AutoUpgradeMinorVersion = *p.AutoUpgradeMinorVersion
			}
			if p.EnableAutomaticUpgrade != nil {
				item.EnableAutomaticUpgrade = *p.EnableAutomaticUpgrade
			}
			if p.ForceUpdateTag != nil {
				item.ForceUpdateTag = *p.ForceUpdateTag
			}
			if p.SuppressFailures != nil {
				item.SuppressFailures = *p.SuppressFailures
			}
			if p.ProvisionAfterExtensions != nil {
				item.ProvisionAfterExtensions = *p.ProvisionAfterExtensions
			}
			if p.ProtectedSettingsFromKeyVault != nil {
				item.SecretUrl = p.ProtectedSettingsFromKeyVault.SecretURL
				if p.ProtectedSettingsFromKeyVault.SourceVault.Id != nil {
					item.SourceVaultId = *p.ProtectedSettingsFromKeyVault.SourceVault.Id
				}
			}
		}
		result = append(result, item)
	}
	return result
}

func expandComputeFleetNetworkProfile(nics []ComputeFleetNicConfig, networkApiVersion string, metadata sdk.ResourceMetaData) *fleets.VirtualMachineScaleSetNetworkProfile {
	if len(nics) == 0 {
		return nil
	}
	configs := make([]fleets.VirtualMachineScaleSetNetworkConfiguration, 0, len(nics))
	for nicIdx, nic := range nics {
		ipConfigs := make([]fleets.VirtualMachineScaleSetIPConfiguration, 0, len(nic.IPConfigurations))
		for ipIdx, ip := range nic.IPConfigurations {
			ipc := fleets.VirtualMachineScaleSetIPConfiguration{
				Name: ip.Name,
			}
			ipProps := &fleets.VirtualMachineScaleSetIPConfigurationProperties{}
			if ip.SubnetId != "" {
				ipProps.Subnet = &fleets.ApiEntityReference{Id: &ip.SubnetId}
			}
			if ip.Primary {
				ipProps.Primary = &ip.Primary
			}
			if ip.PrivateIPAddressVersion != "" {
				v := fleets.IPVersion(ip.PrivateIPAddressVersion)
				ipProps.PrivateIPAddressVersion = &v
			}
			if ip.PublicIPAddressConfigurationName != "" || ip.DeleteOption != "" || ip.DomainNameLabel != "" || ip.IdleTimeoutInMinutes != 0 || ip.SkuName != "" {
				ipProps.PublicIPAddressConfiguration = expandComputeFleetPublicIPConfig(ip, metadata, nicIdx, ipIdx)
			}
			ipc.Properties = ipProps
			ipConfigs = append(ipConfigs, ipc)
		}
		nicProps := &fleets.VirtualMachineScaleSetNetworkConfigurationProperties{
			IPConfigurations: ipConfigs,
		}
		if nic.Primary {
			nicProps.Primary = &nic.Primary
		}
		if nic.EnableAcceleratedNetworking {
			nicProps.EnableAcceleratedNetworking = &nic.EnableAcceleratedNetworking
		}
		if nic.EnableIPForwarding {
			nicProps.EnableIPForwarding = &nic.EnableIPForwarding
		}
		if nic.NetworkSecurityGroupId != "" {
			nicProps.NetworkSecurityGroup = &fleets.SubResource{Id: &nic.NetworkSecurityGroupId}
		}
		if nic.DeleteOption != "" {
			d := fleets.DeleteOptions(nic.DeleteOption)
			nicProps.DeleteOption = &d
		}
		if nic.AuxiliaryMode != "" {
			am := fleets.NetworkInterfaceAuxiliaryMode(nic.AuxiliaryMode)
			nicProps.AuxiliaryMode = &am
		}
		if nic.AuxiliarySku != "" {
			as := fleets.NetworkInterfaceAuxiliarySku(nic.AuxiliarySku)
			nicProps.AuxiliarySku = &as
		}
		if len(nic.DnsServers) > 0 {
			nicProps.DnsSettings = &fleets.VirtualMachineScaleSetNetworkConfigurationDnsSettings{
				DnsServers: &nic.DnsServers,
			}
		}
		configs = append(configs, fleets.VirtualMachineScaleSetNetworkConfiguration{
			Name:       nic.Name,
			Properties: nicProps,
		})
	}
	np := &fleets.VirtualMachineScaleSetNetworkProfile{
		NetworkInterfaceConfigurations: &configs,
	}
	if networkApiVersion != "" {
		nav := fleets.NetworkApiVersion(networkApiVersion)
		np.NetworkApiVersion = &nav
	}
	return np
}

func expandComputeFleetPublicIPConfig(input ComputeFleetIPConfig, metadata sdk.ResourceMetaData, nicIdx int, ipIdx int) *fleets.VirtualMachineScaleSetPublicIPAddressConfiguration {
	cfg := &fleets.VirtualMachineScaleSetPublicIPAddressConfiguration{
		Name: input.PublicIPAddressConfigurationName,
	}
	props := &fleets.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{}
	if input.DeleteOption != "" {
		d := fleets.DeleteOptions(input.DeleteOption)
		props.DeleteOption = &d
	}
	if !metadata.ResourceData.GetRawConfig().AsValueMap()["network_interface_configurations"].AsValueSlice()[nicIdx].AsValueMap()["ip_configurations"].AsValueSlice()[ipIdx].AsValueMap()["idle_timeout_in_minutes"].IsNull() {
		t := input.IdleTimeoutInMinutes
		props.IdleTimeoutInMinutes = &t
	}
	if input.PublicIPAddressVersion != "" {
		v := fleets.IPVersion(input.PublicIPAddressVersion)
		props.PublicIPAddressVersion = &v
	}
	if input.DomainNameLabel != "" || input.DomainNameLabelScope != "" {
		dns := &fleets.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings{
			DomainNameLabel: input.DomainNameLabel,
		}
		if input.DomainNameLabelScope != "" {
			scope := fleets.DomainNameLabelScopeTypes(input.DomainNameLabelScope)
			dns.DomainNameLabelScope = &scope
		}
		props.DnsSettings = dns
	}
	cfg.Properties = props
	if input.SkuName != "" || input.SkuTier != "" {
		sku := &fleets.PublicIPAddressSku{}
		if input.SkuName != "" {
			n := fleets.PublicIPAddressSkuName(input.SkuName)
			sku.Name = &n
		}
		if input.SkuTier != "" {
			tier := fleets.PublicIPAddressSkuTier(input.SkuTier)
			sku.Tier = &tier
		}
		cfg.Sku = sku
	}
	return cfg
}

func flattenComputeFleetPublicIPConfig(input *fleets.VirtualMachineScaleSetPublicIPAddressConfiguration, target *ComputeFleetIPConfig) {
	if input == nil {
		return
	}
	target.PublicIPAddressConfigurationName = input.Name
	if p := input.Properties; p != nil {
		if p.DeleteOption != nil {
			target.DeleteOption = string(*p.DeleteOption)
		}
		if p.IdleTimeoutInMinutes != nil {
			target.IdleTimeoutInMinutes = *p.IdleTimeoutInMinutes
		}
		if p.PublicIPAddressVersion != nil {
			target.PublicIPAddressVersion = string(*p.PublicIPAddressVersion)
		}
		if p.DnsSettings != nil {
			target.DomainNameLabel = p.DnsSettings.DomainNameLabel
			if p.DnsSettings.DomainNameLabelScope != nil {
				target.DomainNameLabelScope = string(*p.DnsSettings.DomainNameLabelScope)
			}
		}
	}
	if input.Sku != nil {
		if input.Sku.Name != nil {
			target.SkuName = string(*input.Sku.Name)
		}
		if input.Sku.Tier != nil {
			target.SkuTier = string(*input.Sku.Tier)
		}
	}
}

func flattenComputeFleetNicConfigs(input *[]fleets.VirtualMachineScaleSetNetworkConfiguration) []ComputeFleetNicConfig {
	if input == nil {
		return []ComputeFleetNicConfig{}
	}
	result := make([]ComputeFleetNicConfig, 0, len(*input))
	for _, nic := range *input {
		cfg := ComputeFleetNicConfig{Name: nic.Name}
		if p := nic.Properties; p != nil {
			if p.Primary != nil {
				cfg.Primary = *p.Primary
			}
			if p.EnableAcceleratedNetworking != nil {
				cfg.EnableAcceleratedNetworking = *p.EnableAcceleratedNetworking
			}
			if p.EnableIPForwarding != nil {
				cfg.EnableIPForwarding = *p.EnableIPForwarding
			}
			if p.NetworkSecurityGroup != nil && p.NetworkSecurityGroup.Id != nil {
				cfg.NetworkSecurityGroupId = *p.NetworkSecurityGroup.Id
			}
			if p.DeleteOption != nil {
				cfg.DeleteOption = string(*p.DeleteOption)
			}
			if p.AuxiliaryMode != nil {
				cfg.AuxiliaryMode = string(*p.AuxiliaryMode)
			}
			if p.AuxiliarySku != nil {
				cfg.AuxiliarySku = string(*p.AuxiliarySku)
			}
			if p.DnsSettings != nil && p.DnsSettings.DnsServers != nil {
				cfg.DnsServers = *p.DnsSettings.DnsServers
			}
			for _, ip := range p.IPConfigurations {
				ipCfg := ComputeFleetIPConfig{Name: ip.Name}
				if ip.Properties != nil {
					if ip.Properties.Subnet != nil && ip.Properties.Subnet.Id != nil {
						ipCfg.SubnetId = *ip.Properties.Subnet.Id
					}
					if ip.Properties.Primary != nil {
						ipCfg.Primary = *ip.Properties.Primary
					}
					if ip.Properties.PrivateIPAddressVersion != nil {
						ipCfg.PrivateIPAddressVersion = string(*ip.Properties.PrivateIPAddressVersion)
					}
					if ip.Properties.PublicIPAddressConfiguration != nil {
						flattenComputeFleetPublicIPConfig(ip.Properties.PublicIPAddressConfiguration, &ipCfg)
					}
				}
				cfg.IPConfigurations = append(cfg.IPConfigurations, ipCfg)
			}
		}
		result = append(result, cfg)
	}
	return result
}

func expandComputeFleetLinuxConfiguration(config ComputeFleetResourceModel) *fleets.LinuxConfiguration {
	cfg := &fleets.LinuxConfiguration{}
	if config.EnableVMAgentPlatformUpdates {
		cfg.EnableVMAgentPlatformUpdates = &config.EnableVMAgentPlatformUpdates
	}
	if config.ProvisionVMAgent {
		cfg.ProvisionVMAgent = &config.ProvisionVMAgent
	}
	if config.PatchMode != "" || config.BypassPlatformSafetyChecksOnUserSchedule || config.RebootSetting != "" {
		ps := &fleets.LinuxPatchSettings{}
		if config.PatchMode != "" {
			pm := fleets.LinuxVMGuestPatchMode(config.PatchMode)
			ps.PatchMode = &pm
		}
		if config.BypassPlatformSafetyChecksOnUserSchedule || config.RebootSetting != "" {
			abp := &fleets.LinuxVMGuestPatchAutomaticByPlatformSettings{}
			if config.BypassPlatformSafetyChecksOnUserSchedule {
				abp.BypassPlatformSafetyChecksOnUserSchedule = &config.BypassPlatformSafetyChecksOnUserSchedule
			}
			if config.RebootSetting != "" {
				rs := fleets.LinuxVMGuestPatchAutomaticByPlatformRebootSetting(config.RebootSetting)
				abp.RebootSetting = &rs
			}
			ps.AutomaticByPlatformSettings = abp
		}
		cfg.PatchSettings = ps
	}
	if len(config.KeyData) > 0 {
		keys := make([]fleets.SshPublicKey, 0, len(config.KeyData))
		for _, k := range config.KeyData {
			kd := k
			keys = append(keys, fleets.SshPublicKey{
				KeyData: &kd,
				Path:    pointer.To(fmt.Sprintf("/home/%s/.ssh/authorized_keys", config.AdminUsername)),
			})
		}
		cfg.Ssh = &fleets.SshConfiguration{PublicKeys: &keys}
	}
	return cfg
}

func flattenComputeFleetLinuxConfiguration(input *fleets.LinuxConfiguration, schema *ComputeFleetResourceModel) {
	if input == nil {
		return
	}
	if input.EnableVMAgentPlatformUpdates != nil {
		schema.EnableVMAgentPlatformUpdates = *input.EnableVMAgentPlatformUpdates
	}
	if input.ProvisionVMAgent != nil {
		schema.ProvisionVMAgent = *input.ProvisionVMAgent
	}
	if ps := input.PatchSettings; ps != nil {
		if ps.PatchMode != nil {
			schema.PatchMode = string(*ps.PatchMode)
		}
		if abp := ps.AutomaticByPlatformSettings; abp != nil {
			if abp.BypassPlatformSafetyChecksOnUserSchedule != nil {
				schema.BypassPlatformSafetyChecksOnUserSchedule = *abp.BypassPlatformSafetyChecksOnUserSchedule
			}
			if abp.RebootSetting != nil {
				schema.RebootSetting = string(*abp.RebootSetting)
			}
		}
	}
	if input.Ssh != nil && input.Ssh.PublicKeys != nil {
		for _, k := range *input.Ssh.PublicKeys {
			if k.KeyData != nil {
				schema.KeyData = append(schema.KeyData, *k.KeyData)
			}
		}
	}
}

func expandComputeFleetWindowsConfiguration(config ComputeFleetResourceModel) *fleets.WindowsConfiguration {
	cfg := &fleets.WindowsConfiguration{}
	if config.EnableAutomaticUpdates {
		cfg.EnableAutomaticUpdates = &config.EnableAutomaticUpdates
	}
	if config.WindowsConfigurationEnableVMAgentPlatformUpdates {
		cfg.EnableVMAgentPlatformUpdates = &config.WindowsConfigurationEnableVMAgentPlatformUpdates
	}
	if config.WindowsConfigurationProvisionVMAgent {
		cfg.ProvisionVMAgent = &config.WindowsConfigurationProvisionVMAgent
	}
	if config.TimeZone != "" {
		cfg.TimeZone = &config.TimeZone
	}
	if config.WindowsConfigurationPatchMode != "" || config.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule || config.WindowsConfigurationRebootSetting != "" || config.EnableHotpatching {
		ps := &fleets.PatchSettings{}
		if config.WindowsConfigurationPatchMode != "" {
			pm := fleets.WindowsVMGuestPatchMode(config.WindowsConfigurationPatchMode)
			ps.PatchMode = &pm
		}
		if config.EnableHotpatching {
			ps.EnableHotpatching = &config.EnableHotpatching
		}
		if config.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule || config.WindowsConfigurationRebootSetting != "" {
			abp := &fleets.WindowsVMGuestPatchAutomaticByPlatformSettings{}
			if config.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule {
				abp.BypassPlatformSafetyChecksOnUserSchedule = &config.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule
			}
			if config.WindowsConfigurationRebootSetting != "" {
				rs := fleets.WindowsVMGuestPatchAutomaticByPlatformRebootSetting(config.WindowsConfigurationRebootSetting)
				abp.RebootSetting = &rs
			}
			ps.AutomaticByPlatformSettings = abp
		}
		cfg.PatchSettings = ps
	}
	if len(config.AdditionalUnattendContent) > 0 {
		items := make([]fleets.AdditionalUnattendContent, 0, len(config.AdditionalUnattendContent))
		for _, a := range config.AdditionalUnattendContent {
			item := fleets.AdditionalUnattendContent{}
			if a.SettingName != "" {
				sn := fleets.SettingNames(a.SettingName)
				item.SettingName = &sn
			}
			if a.Content != "" {
				item.Content = &a.Content
			}
			items = append(items, item)
		}
		cfg.AdditionalUnattendContent = &items
	}
	if len(config.Listeners) > 0 {
		listeners := make([]fleets.WinRMListener, 0, len(config.Listeners))
		for _, l := range config.Listeners {
			wl := fleets.WinRMListener{}
			if l.Protocol != "" {
				p := fleets.ProtocolTypes(l.Protocol)
				wl.Protocol = &p
			}
			if l.CertificateUrl != "" {
				wl.CertificateURL = &l.CertificateUrl
			}
			listeners = append(listeners, wl)
		}
		cfg.WinRM = &fleets.WinRMConfiguration{Listeners: &listeners}
	}
	return cfg
}

func flattenComputeFleetWindowsConfiguration(input *fleets.WindowsConfiguration, schema *ComputeFleetResourceModel) {
	if input == nil {
		return
	}
	if input.EnableAutomaticUpdates != nil {
		schema.EnableAutomaticUpdates = *input.EnableAutomaticUpdates
	}
	if input.EnableVMAgentPlatformUpdates != nil {
		schema.WindowsConfigurationEnableVMAgentPlatformUpdates = *input.EnableVMAgentPlatformUpdates
	}
	if input.ProvisionVMAgent != nil {
		schema.WindowsConfigurationProvisionVMAgent = *input.ProvisionVMAgent
	}
	if input.TimeZone != nil {
		schema.TimeZone = *input.TimeZone
	}
	if ps := input.PatchSettings; ps != nil {
		if ps.PatchMode != nil {
			schema.WindowsConfigurationPatchMode = string(*ps.PatchMode)
		}
		if ps.EnableHotpatching != nil {
			schema.EnableHotpatching = *ps.EnableHotpatching
		}
		if abp := ps.AutomaticByPlatformSettings; abp != nil {
			if abp.BypassPlatformSafetyChecksOnUserSchedule != nil {
				schema.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule = *abp.BypassPlatformSafetyChecksOnUserSchedule
			}
			if abp.RebootSetting != nil {
				schema.WindowsConfigurationRebootSetting = string(*abp.RebootSetting)
			}
		}
	}
	if input.AdditionalUnattendContent != nil {
		for _, a := range *input.AdditionalUnattendContent {
			auc := ComputeFleetAdditionalUnattendContent{}
			if a.SettingName != nil {
				auc.SettingName = string(*a.SettingName)
			}
			if a.Content != nil {
				auc.Content = *a.Content
			}
			schema.AdditionalUnattendContent = append(schema.AdditionalUnattendContent, auc)
		}
	}
	if input.WinRM != nil && input.WinRM.Listeners != nil {
		for _, l := range *input.WinRM.Listeners {
			wl := ComputeFleetWinRMListener{}
			if l.Protocol != nil {
				wl.Protocol = string(*l.Protocol)
			}
			if l.CertificateURL != nil {
				wl.CertificateUrl = *l.CertificateURL
			}
			schema.Listeners = append(schema.Listeners, wl)
		}
	}
}

func expandComputeFleetSecrets(input []ComputeFleetSecret) *[]fleets.VaultSecretGroup {
	result := make([]fleets.VaultSecretGroup, 0, len(input))
	for _, s := range input {
		vsg := fleets.VaultSecretGroup{}
		if s.SourceVaultId != "" {
			vsg.SourceVault = &fleets.SubResource{Id: &s.SourceVaultId}
		}
		if len(s.VaultCertificates) > 0 {
			certs := make([]fleets.VaultCertificate, 0, len(s.VaultCertificates))
			for _, c := range s.VaultCertificates {
				vc := fleets.VaultCertificate{}
				if c.CertificateUrl != "" {
					vc.CertificateURL = &c.CertificateUrl
				}
				if c.CertificateStore != "" {
					vc.CertificateStore = &c.CertificateStore
				}
				certs = append(certs, vc)
			}
			vsg.VaultCertificates = &certs
		}
		result = append(result, vsg)
	}
	return &result
}

func flattenComputeFleetSecrets(input *[]fleets.VaultSecretGroup) []ComputeFleetSecret {
	if input == nil {
		return nil
	}
	result := make([]ComputeFleetSecret, 0, len(*input))
	for _, s := range *input {
		item := ComputeFleetSecret{}
		if s.SourceVault != nil && s.SourceVault.Id != nil {
			item.SourceVaultId = *s.SourceVault.Id
		}
		if s.VaultCertificates != nil {
			vcerts := make([]ComputeFleetVaultCertificate, 0, len(*s.VaultCertificates))
			for _, c := range *s.VaultCertificates {
				vc := ComputeFleetVaultCertificate{}
				if c.CertificateURL != nil {
					vc.CertificateUrl = *c.CertificateURL
				}
				if c.CertificateStore != nil {
					vc.CertificateStore = *c.CertificateStore
				}
				vcerts = append(vcerts, vc)
			}
			item.VaultCertificates = vcerts
		}
		result = append(result, item)
	}
	return result
}

func expandComputeFleetDataDisks(input []ComputeFleetDataDisk, metadata sdk.ResourceMetaData) *[]fleets.VirtualMachineScaleSetDataDisk {
	result := make([]fleets.VirtualMachineScaleSetDataDisk, 0, len(input))
	for i, d := range input {
		disk := fleets.VirtualMachineScaleSetDataDisk{
			Lun:          d.Lun,
			CreateOption: fleets.DiskCreateOptionTypes(d.CreateOption),
		}
		if d.Caching != "" {
			c := fleets.CachingTypes(d.Caching)
			disk.Caching = &c
		}
		if d.DeleteOption != "" {
			do := fleets.DiskDeleteOptionTypes(d.DeleteOption)
			disk.DeleteOption = &do
		}
		if !metadata.ResourceData.GetRawConfig().AsValueMap()["data_disks"].AsValueSlice()[i].AsValueMap()["disk_size_gb"].IsNull() {
			sz := d.DiskSizeGB
			disk.DiskSizeGB = &sz
		}
		if d.WriteAcceleratorEnabled {
			disk.WriteAcceleratorEnabled = &d.WriteAcceleratorEnabled
		}
		if d.DiskEncryptionSetId != "" || d.StorageAccountType != "" {
			md := &fleets.VirtualMachineScaleSetManagedDiskParameters{}
			if d.DiskEncryptionSetId != "" {
				md.DiskEncryptionSet = &fleets.DiskEncryptionSetParameters{Id: &d.DiskEncryptionSetId}
			}
			if d.StorageAccountType != "" {
				sat := fleets.StorageAccountTypes(d.StorageAccountType)
				md.StorageAccountType = &sat
			}
			disk.ManagedDisk = md
		}
		result = append(result, disk)
	}
	return &result
}

func flattenComputeFleetDataDisks(input *[]fleets.VirtualMachineScaleSetDataDisk) []ComputeFleetDataDisk {
	if input == nil {
		return nil
	}
	result := make([]ComputeFleetDataDisk, 0, len(*input))
	for _, d := range *input {
		item := ComputeFleetDataDisk{
			Lun:          int64(d.Lun),
			CreateOption: string(d.CreateOption),
		}
		if d.Caching != nil {
			item.Caching = string(*d.Caching)
		}
		if d.DeleteOption != nil {
			item.DeleteOption = string(*d.DeleteOption)
		}
		if d.DiskSizeGB != nil {
			item.DiskSizeGB = int64(*d.DiskSizeGB)
		}
		if d.WriteAcceleratorEnabled != nil {
			item.WriteAcceleratorEnabled = *d.WriteAcceleratorEnabled
		}
		if d.ManagedDisk != nil {
			if d.ManagedDisk.DiskEncryptionSet != nil && d.ManagedDisk.DiskEncryptionSet.Id != nil {
				item.DiskEncryptionSetId = *d.ManagedDisk.DiskEncryptionSet.Id
			}
			if d.ManagedDisk.StorageAccountType != nil {
				item.StorageAccountType = string(*d.ManagedDisk.StorageAccountType)
			}
		}
		result = append(result, item)
	}
	return result
}

func expandComputeFleetRegularPriorityProfile(config ComputeFleetResourceModel, metadata sdk.ResourceMetaData) *fleets.RegularPriorityProfile {
	if len(config.RegularPriorityProfile) == 0 {
		return nil
	}
	p := config.RegularPriorityProfile[0]
	rpp := &fleets.RegularPriorityProfile{}
	rpp.Capacity = &p.Capacity

	if !metadata.ResourceData.GetRawConfig().AsValueMap()["regular_priority_profile"].AsValueSlice()[0].AsValueMap()["min_capacity"].IsNull() {
		rpp.MinCapacity = &p.MinCapacity
	}
	if p.AllocationStrategy != "" {
		as := fleets.RegularPriorityAllocationStrategy(p.AllocationStrategy)
		rpp.AllocationStrategy = &as
	}
	return rpp
}

func flattenComputeFleetRegularPriorityProfile(input *fleets.RegularPriorityProfile, schema *ComputeFleetResourceModel) {
	if input == nil {
		schema.RegularPriorityProfile = []ComputeFleetRegularPriorityProfile{}
		return
	}
	p := ComputeFleetRegularPriorityProfile{}
	if input.Capacity != nil {
		p.Capacity = *input.Capacity
	}
	if input.MinCapacity != nil {
		p.MinCapacity = *input.MinCapacity
	}
	if input.AllocationStrategy != nil {
		p.AllocationStrategy = string(*input.AllocationStrategy)
	}
	schema.RegularPriorityProfile = []ComputeFleetRegularPriorityProfile{p}
}

func expandComputeFleetSpotPriorityProfile(config ComputeFleetResourceModel, metadata sdk.ResourceMetaData) *fleets.SpotPriorityProfile {
	if len(config.SpotPriorityProfile) == 0 {
		return nil
	}
	p := config.SpotPriorityProfile[0]
	spp := &fleets.SpotPriorityProfile{}
	spp.Capacity = &p.Capacity

	if !metadata.ResourceData.GetRawConfig().AsValueMap()["spot_priority_profile"].AsValueSlice()[0].AsValueMap()["min_capacity"].IsNull() {
		spp.MinCapacity = &p.MinCapacity
	}
	if !metadata.ResourceData.GetRawConfig().AsValueMap()["spot_priority_profile"].AsValueSlice()[0].AsValueMap()["max_price_per_vm"].IsNull() {
		spp.MaxPricePerVM = &p.MaxPricePerVM
	}
	if p.EvictionPolicy != "" {
		ep := fleets.EvictionPolicy(p.EvictionPolicy)
		spp.EvictionPolicy = &ep
	}
	if p.AllocationStrategy != "" {
		as := fleets.SpotAllocationStrategy(p.AllocationStrategy)
		spp.AllocationStrategy = &as
	}
	if p.Maintain {
		spp.Maintain = &p.Maintain
	}
	return spp
}

func flattenComputeFleetSpotPriorityProfile(input *fleets.SpotPriorityProfile, schema *ComputeFleetResourceModel) {
	if input == nil {
		schema.SpotPriorityProfile = []ComputeFleetSpotPriorityProfile{}
		return
	}
	p := ComputeFleetSpotPriorityProfile{}
	if input.Capacity != nil {
		p.Capacity = *input.Capacity
	}
	if input.MinCapacity != nil {
		p.MinCapacity = *input.MinCapacity
	}
	if input.MaxPricePerVM != nil {
		p.MaxPricePerVM = *input.MaxPricePerVM
	}
	if input.EvictionPolicy != nil {
		p.EvictionPolicy = string(*input.EvictionPolicy)
	}
	if input.AllocationStrategy != nil {
		p.AllocationStrategy = string(*input.AllocationStrategy)
	}
	if input.Maintain != nil {
		p.Maintain = *input.Maintain
	}
	schema.SpotPriorityProfile = []ComputeFleetSpotPriorityProfile{p}
}

func expandComputeFleetPlan(input []ComputeFleetPlan) *fleets.Plan {
	if len(input) == 0 {
		return nil
	}
	p := input[0]
	return &fleets.Plan{
		Name:          p.Name,
		Publisher:     p.Publisher,
		Product:       p.Product,
		PromotionCode: pointer.To(p.PromotionCode),
	}
}

func flattenComputeFleetPlan(input *fleets.Plan) []ComputeFleetPlan {
	if input == nil {
		return []ComputeFleetPlan{}
	}
	p := ComputeFleetPlan{
		Name:      input.Name,
		Publisher: input.Publisher,
		Product:   input.Product,
	}
	if input.PromotionCode != nil {
		p.PromotionCode = *input.PromotionCode
	}
	return []ComputeFleetPlan{p}
}
