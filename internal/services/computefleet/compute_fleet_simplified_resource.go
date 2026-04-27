// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package computefleet

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/location"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurefleet/2024-11-01/fleets"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/tf"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type ComputeFleetResource struct{}

var _ sdk.ResourceWithUpdate = ComputeFleetResource{}

// VMSizeProfileModel has 2 children – stays as a block (rule 9/10 don't apply).
type VMSizeProfileModel struct {
	Name string `tfschema:"name"`
	Rank int64  `tfschema:"rank"`
}

// GalleryApplicationModel has 6 children – stays as a block.
type GalleryApplicationModel struct {
	PackageReferenceId              string `tfschema:"package_reference_id"`
	EnableAutomaticUpgrade          bool   `tfschema:"enable_automatic_upgrade"`
	ConfigurationReference          string `tfschema:"configuration_reference"`
	Order                           int64  `tfschema:"order"`
	Tags                            string `tfschema:"tags"`
	TreatFailureAsDeploymentFailure bool   `tfschema:"treat_failure_as_deployment_failure"`
}

// ExtensionModel: protected_settings_from_key_vault (MaxItems:1) flattened here by Rule 10.
// secret_url and source_vault_id are now direct fields (no clash with existing extension fields).
type ExtensionModel struct {
	Name                     string   `tfschema:"name"`
	Publisher                string   `tfschema:"publisher"`
	Type                     string   `tfschema:"type"`
	TypeHandlerVersion       string   `tfschema:"type_handler_version"`
	AutoUpgradeMinorVersion  bool     `tfschema:"auto_upgrade_minor_version"`
	EnableAutomaticUpgrade   bool     `tfschema:"enable_automatic_upgrade"`
	ForceUpdateTag           string   `tfschema:"force_update_tag"`
	ProtectedSettings        string   `tfschema:"protected_settings"`
	SecretUrl                string   `tfschema:"secret_url"`
	SourceVaultId            string   `tfschema:"source_vault_id"`
	ProvisionAfterExtensions []string `tfschema:"provision_after_extensions"`
	SuppressFailures         bool     `tfschema:"suppress_failures"`
	Settings                 string   `tfschema:"settings"`
}

// SecretModel: vault_certificate has 1 child (certificate_url) → Rule 9 flattens it.
// certificate_url becomes a TypeList of TypeString directly in secret.
type SecretModel struct {
	SourceVaultId  string   `tfschema:"source_vault_id"`
	CertificateUrl []string `tfschema:"certificate_url"`
}

// DataDiskModel: managed_disk (MaxItems:1) flattened by Rule 10.
// disk_encryption_set_id and storage_account_type are now direct fields.
type DataDiskModel struct {
	Lun                     int64  `tfschema:"lun"`
	CreateOption            string `tfschema:"create_option"`
	Caching                 string `tfschema:"caching"`
	DeleteOption            string `tfschema:"delete_option"`
	DiskSizeGB              int64  `tfschema:"disk_size_gb"`
	WriteAcceleratorEnabled bool   `tfschema:"write_accelerator_enabled"`
	DiskEncryptionSetId     string `tfschema:"disk_encryption_set_id"`
	StorageAccountType      string `tfschema:"storage_account_type"`
}

// AdditionalUnattendContentModel has 2 children – stays as a block.
type AdditionalUnattendContentModel struct {
	SettingName string `tfschema:"setting_name"`
	Content     string `tfschema:"content"`
}

// WinRMListenerModel has 2 children – stays as a block.
type WinRMListenerModel struct {
	CertificateUrl string `tfschema:"certificate_url"`
	Protocol       string `tfschema:"protocol"`
}

// IpConfigurationModel: public_ip_address_configuration (MaxItems:1) flattened by Rule 10.
// name clashes with ip_configuration.name → public_ip_address_configuration_name.
// All other public_ip children have no clash and use direct names.
type IpConfigurationModel struct {
	Name                                    string   `tfschema:"name"`
	Subnet                                  string   `tfschema:"subnet"`
	ApplicationGatewayBackendAddressPoolsId []string `tfschema:"application_gateway_backend_address_pools_id"`
	ApplicationSecurityGroupsId             []string `tfschema:"application_security_groups_id"`
	LoadBalancerBackendAddressPoolsId       []string `tfschema:"load_balancer_backend_address_pools_id"`
	Primary                                 bool     `tfschema:"primary"`
	PrivateIPAddressVersion                 string   `tfschema:"private_ip_address_version"`
	PublicIPAddressConfigurationName        string   `tfschema:"public_ip_address_configuration_name"`
	DeleteOption                            string   `tfschema:"delete_option"`
	DomainNameLabel                         string   `tfschema:"domain_name_label"`
	DomainNameLabelScope                    string   `tfschema:"domain_name_label_scope"`
	IdleTimeoutInMinutes                    int64    `tfschema:"idle_timeout_in_minutes"`
	PublicIPAddressVersion                  string   `tfschema:"public_ip_address_version"`
	SkuName                                 string   `tfschema:"sku_name"`
	SkuTier                                 string   `tfschema:"sku_tier"`
}

// NetworkInterfaceConfigurationModel: stays as a block (many children, no MaxItems:1).
type NetworkInterfaceConfigurationModel struct {
	Name                        string                 `tfschema:"name"`
	IpConfiguration             []IpConfigurationModel `tfschema:"ip_configuration"`
	EnableAcceleratedNetworking bool                   `tfschema:"enable_accelerated_networking"`
	AuxiliaryMode               string                 `tfschema:"auxiliary_mode"`
	AuxiliarySku                string                 `tfschema:"auxiliary_sku"`
	DeleteOption                string                 `tfschema:"delete_option"`
	DnsServers                  []string               `tfschema:"dns_servers"`
	EnableIPForwarding          bool                   `tfschema:"enable_ip_forwarding"`
	NetworkSecurityGroupId      string                 `tfschema:"network_security_group_id"`
	Primary                     bool                   `tfschema:"primary"`
}

// ComputeFleetResourceModel:
//
// linux_configuration (MaxItems:1) flattened to root by Rule 10.
//
//	patch_settings (MaxItems:1) and automatic_by_platform_settings (MaxItems:1) recursively
//	flattened through linux_configuration before reaching root. All linux fields have no clash.
//
// windows_configuration (MaxItems:1) flattened to root by Rule 10.
//
//	patch_settings (MaxItems:1) and automatic_by_platform_settings (MaxItems:1) recursively
//	flattened. Fields that clash with linux equivalents get windows_configuration_ prefix.
//
// os_disk (MaxItems:1) flattened to root.
//
//	diff_disk_settings (MaxItems:1) and managed_disk (MaxItems:1) recursively flattened
//	into os_disk before it reaches root. No clashes with root.
//
// plan (MaxItems:1) flattened to root: name→plan_name, publisher→plan_publisher.
//
// regular_priority_profile (MaxItems:1) flattened to root: capacity→regular_priority_profile_capacity.
type ComputeFleetResourceModel struct {
	Name              string `tfschema:"name"`
	ResourceGroupName string `tfschema:"resource_group_name"`
	Location          string `tfschema:"location"`
	ComputeApiVersion string `tfschema:"compute_api_version"`

	// Image reference
	Offer     string `tfschema:"offer"`
	Publisher string `tfschema:"publisher"`
	Sku       string `tfschema:"sku"`
	Version   string `tfschema:"version"`

	// OS profile basics
	AdminPassword      string `tfschema:"admin_password"`
	AdminUsername      string `tfschema:"admin_username"`
	ComputerNamePrefix string `tfschema:"computer_name_prefix"`
	CustomData         string `tfschema:"custom_data"`

	// Linux configuration (flattened from MaxItems:1 block + nested MaxItems:1 blocks)
	DisablePasswordAuthentication            bool     `tfschema:"disable_password_authentication"`
	EnableVMAgentPlatformUpdates             bool     `tfschema:"enable_vm_agent_platform_updates"`
	ProvisionVMAgent                         bool     `tfschema:"provision_vm_agent"`
	KeyData                                  []string `tfschema:"key_data"`
	PatchMode                                string   `tfschema:"patch_mode"`
	BypassPlatformSafetyChecksOnUserSchedule bool     `tfschema:"bypass_platform_safety_checks_on_user_schedule"`
	RebootSetting                            string   `tfschema:"reboot_setting"`

	// Windows configuration (flattened; clash fields prefixed with windows_configuration_)
	EnableAutomaticUpdates                                       bool                             `tfschema:"enable_automatic_updates"`
	WindowsConfigurationEnableVMAgentPlatformUpdates             bool                             `tfschema:"windows_configuration_enable_vm_agent_platform_updates"`
	WindowsConfigurationProvisionVMAgent                         bool                             `tfschema:"windows_configuration_provision_vm_agent"`
	TimeZone                                                     string                           `tfschema:"time_zone"`
	WindowsConfigurationPatchMode                                string                           `tfschema:"windows_configuration_patch_mode"`
	EnableHotpatching                                            bool                             `tfschema:"enable_hotpatching"`
	WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule bool                             `tfschema:"windows_configuration_bypass_platform_safety_checks_on_user_schedule"`
	WindowsConfigurationRebootSetting                            string                           `tfschema:"windows_configuration_reboot_setting"`
	AdditionalUnattendContent                                    []AdditionalUnattendContentModel `tfschema:"additional_unattend_content"`
	WinRMListeners                                               []WinRMListenerModel             `tfschema:"winrm_listener"`

	// OS profile secrets (vault_certificate rule 9 flattened)
	Secrets []SecretModel `tfschema:"secret"`

	// Application profile
	GalleryApplications []GalleryApplicationModel `tfschema:"gallery_application"`

	// Extension profile (protected_settings_from_key_vault rule 10 flattened)
	Extensions []ExtensionModel `tfschema:"extension"`

	// Data disks (managed_disk MaxItems:1 flattened into data_disk)
	DataDisks []DataDiskModel `tfschema:"data_disk"`

	// OS disk (MaxItems:1 flattened to root; nested MaxItems:1 blocks also flattened first)
	Caching                 string `tfschema:"caching"`
	DeleteOption            string `tfschema:"delete_option"`
	DiskSizeGB              int64  `tfschema:"disk_size_gb"`
	WriteAcceleratorEnabled bool   `tfschema:"write_accelerator_enabled"`
	Option                  string `tfschema:"option"`
	Placement               string `tfschema:"placement"`
	DiskEncryptionSetId     string `tfschema:"disk_encryption_set_id"`
	SecurityEncryptionType  string `tfschema:"security_encryption_type"`
	StorageAccountType      string `tfschema:"storage_account_type"`

	// Plan (MaxItems:1 flattened; name→plan_name, publisher→plan_publisher due to clashes)
	PlanName      string `tfschema:"plan_name"`
	PlanPublisher string `tfschema:"plan_publisher"`
	Product       string `tfschema:"product"`
	PromotionCode string `tfschema:"promotion_code"`

	// Spot priority profile
	Capacity               int64   `tfschema:"capacity"`
	SpotMinCapacity        int64   `tfschema:"spot_min_capacity"`
	MaxPricePerVM          float64 `tfschema:"max_price_per_vm"`
	EvictionPolicy         string  `tfschema:"eviction_policy"`
	SpotAllocationStrategy string  `tfschema:"spot_allocation_strategy"`
	Maintain               bool    `tfschema:"maintain"`

	// Regular priority profile (MaxItems:1 flattened; capacity→regular_priority_profile_capacity due to clash)
	RegularPriorityProfileCapacity int64  `tfschema:"regular_priority_profile_capacity"`
	MinCapacity                    int64  `tfschema:"min_capacity"`
	AllocationStrategy             string `tfschema:"allocation_strategy"`

	// VM size profiles
	VmSizeProfiles []VMSizeProfileModel `tfschema:"vm_size_profile"`

	// Network profile
	NetworkApiVersion             string                               `tfschema:"network_api_version"`
	NetworkInterfaceConfiguration []NetworkInterfaceConfigurationModel `tfschema:"network_interface_configuration"`
}

func (r ComputeFleetResource) ModelObject() interface{} {
	return &ComputeFleetResourceModel{}
}

func (r ComputeFleetResource) ResourceType() string {
	return "azurerm_compute_fleet"
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

		"compute_api_version": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Computed: true,
		},

		"capacity": {
			Type:     pluginsdk.TypeInt,
			Required: true,
		},

		"offer": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"publisher": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"sku": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"version": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"admin_username": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"admin_password": {
			Type:      pluginsdk.TypeString,
			Required:  true,
			Sensitive: true,
		},

		"computer_name_prefix": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"custom_data": {
			Type:      pluginsdk.TypeString,
			Optional:  true,
			ForceNew:  true,
			Sensitive: true,
		},

		"vm_size_profile": {
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

		// ── Linux configuration (Rule 10: linux_configuration MaxItems:1 flattened) ──────
		// patch_settings (MaxItems:1) and automatic_by_platform_settings (MaxItems:1) also
		// flattened recursively before reaching root.

		"disable_password_authentication": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"enable_vm_agent_platform_updates": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"provision_vm_agent": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"key_data": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
		},

		"patch_mode": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"bypass_platform_safety_checks_on_user_schedule": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"reboot_setting": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		// ── Windows configuration (Rule 10: windows_configuration MaxItems:1 flattened) ──
		// Fields clashing with linux equivalents prefixed with windows_configuration_.

		"enable_automatic_updates": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"windows_configuration_enable_vm_agent_platform_updates": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"windows_configuration_provision_vm_agent": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"time_zone": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"windows_configuration_patch_mode": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"enable_hotpatching": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"windows_configuration_bypass_platform_safety_checks_on_user_schedule": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"windows_configuration_reboot_setting": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		// additional_unattend_content has 2 children → stays as a block (rule 9 n/a).
		"additional_unattend_content": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"setting_name": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},

					"content": {
						Type:      pluginsdk.TypeString,
						Required:  true,
						Sensitive: true,
					},
				},
			},
		},

		// winrm_listener has 2 children → stays as a block (rule 9 n/a).
		"winrm_listener": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"certificate_url": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"protocol": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},
				},
			},
		},

		// ── Secrets (Rule 9: vault_certificate has 1 child, certificate_url flattened) ──
		"secret": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"source_vault_id": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},

					// Rule 9: vault_certificate had 1 child (certificate_url).
					// Child is flattened → certificate_url is now a TypeList of TypeString.
					"certificate_url": {
						Type:     pluginsdk.TypeList,
						Required: true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},
				},
			},
		},

		// ── Gallery applications ──────────────────────────────────────────────────────────
		"gallery_application": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"package_reference_id": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},

					"enable_automatic_upgrade": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},

					"configuration_reference": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"order": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},

					"tags": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"treat_failure_as_deployment_failure": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
				},
			},
		},

		// ── Extensions (Rule 10: protected_settings_from_key_vault MaxItems:1 flattened) ──
		"extension": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"name": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},

					"publisher": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},

					"type": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},

					"type_handler_version": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"auto_upgrade_minor_version": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},

					"enable_automatic_upgrade": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},

					"force_update_tag": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"protected_settings": {
						Type:      pluginsdk.TypeString,
						Optional:  true,
						Sensitive: true,
					},

					// Rule 10: protected_settings_from_key_vault (MaxItems:1) flattened.
					// secret_url and source_vault_id have no clash within extension.
					"secret_url": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"source_vault_id": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"provision_after_extensions": {
						Type:     pluginsdk.TypeList,
						Optional: true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},

					"suppress_failures": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},

					"settings": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},
				},
			},
		},

		// ── Data disks (Rule 10: managed_disk MaxItems:1 flattened into data_disk) ────────
		"data_disk": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"lun": {
						Type:     pluginsdk.TypeInt,
						Required: true,
					},

					"create_option": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},

					"caching": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"delete_option": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"disk_size_gb": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},

					"write_accelerator_enabled": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},

					// Rule 10: managed_disk (MaxItems:1) flattened.
					"disk_encryption_set_id": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"storage_account_type": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},
				},
			},
		},

		// ── OS disk (Rule 10: os_disk MaxItems:1 flattened to root) ──────────────────────
		// diff_disk_settings (MaxItems:1) and managed_disk (MaxItems:1) first flattened
		// into os_disk, then os_disk flattened to root. No field name clashes at root.

		"caching": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"delete_option": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"disk_size_gb": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
		},

		"write_accelerator_enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		// From diff_disk_settings (MaxItems:1 flattened into os_disk, then to root)
		"option": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"placement": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		// From os_disk managed_disk (MaxItems:1 flattened into os_disk, then to root)
		"disk_encryption_set_id": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"security_encryption_type": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"storage_account_type": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		// ── Plan (Rule 10: plan MaxItems:1 flattened to root) ────────────────────────────
		// name clashes with root name → plan_name
		// publisher clashes with root publisher → plan_publisher

		"plan_name": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"plan_publisher": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"product": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"promotion_code": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		// ── Regular priority profile (Rule 10: MaxItems:1 flattened to root) ─────────────
		// capacity clashes with spot capacity → regular_priority_profile_capacity

		"regular_priority_profile_capacity": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
		},

		"min_capacity": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
		},

		"allocation_strategy": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		// ── Spot priority profile fields (already at root) ───────────────────────────────
		"spot_min_capacity": {
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
		},

		"spot_allocation_strategy": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"maintain": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		// ── Network profile ───────────────────────────────────────────────────────────────
		"network_api_version": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		// network_interface_configuration has many children – stays as a block.
		"network_interface_configuration": {
			Type:     pluginsdk.TypeList,
			Required: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"name": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},

					// ip_configuration has public_ip_address_configuration (MaxItems:1) flattened.
					"ip_configuration": {
						Type:     pluginsdk.TypeList,
						Required: true,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"name": {
									Type:     pluginsdk.TypeString,
									Required: true,
								},

								"subnet": {
									Type:     pluginsdk.TypeString,
									Required: true,
								},

								"application_gateway_backend_address_pools_id": {
									Type:     pluginsdk.TypeList,
									Optional: true,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},

								"application_security_groups_id": {
									Type:     pluginsdk.TypeList,
									Optional: true,
									MaxItems: 10,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},

								"load_balancer_backend_address_pools_id": {
									Type:     pluginsdk.TypeList,
									Optional: true,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},

								"primary": {
									Type:     pluginsdk.TypeBool,
									Optional: true,
								},

								"private_ip_address_version": {
									Type:     pluginsdk.TypeString,
									Optional: true,
								},

								// Rule 10: public_ip_address_configuration (MaxItems:1) flattened.
								// name clashes with ip_configuration.name → public_ip_address_configuration_name.
								"public_ip_address_configuration_name": {
									Type:     pluginsdk.TypeString,
									Optional: true,
								},

								"delete_option": {
									Type:     pluginsdk.TypeString,
									Optional: true,
								},

								"domain_name_label": {
									Type:     pluginsdk.TypeString,
									Optional: true,
								},

								"domain_name_label_scope": {
									Type:     pluginsdk.TypeString,
									Optional: true,
								},

								"idle_timeout_in_minutes": {
									Type:     pluginsdk.TypeInt,
									Optional: true,
								},

								"public_ip_address_version": {
									Type:     pluginsdk.TypeString,
									Optional: true,
								},

								"sku_name": {
									Type:     pluginsdk.TypeString,
									Optional: true,
								},

								"sku_tier": {
									Type:     pluginsdk.TypeString,
									Optional: true,
								},
							},
						},
					},

					"enable_accelerated_networking": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},

					"auxiliary_mode": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"auxiliary_sku": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"delete_option": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"dns_servers": {
						Type:     pluginsdk.TypeList,
						Optional: true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},

					"enable_ip_forwarding": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},

					"network_security_group_id": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"primary": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
				},
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
			client := metadata.Client.ComputeFleet.FleetsClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			var model ComputeFleetResourceModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			id := fleets.NewFleetID(subscriptionId, model.ResourceGroupName, model.Name)

			existing, err := client.Get(ctx, id)
			if err != nil {
				if !response.WasNotFound(existing.HttpResponse) {
					return fmt.Errorf("checking for the presence of an existing %s: %+v", id, err)
				}
			}
			if !response.WasNotFound(existing.HttpResponse) {
				return tf.ImportAsExistsError("azurerm_compute_fleet", id.ID())
			}

			payload := fleets.Fleet{
				Location:   location.Normalize(model.Location),
				Properties: expandComputeFleetProperties(model),
			}

			if plan := expandPlan(model); plan != nil {
				payload.Plan = plan
			}

			if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
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
			client := metadata.Client.ComputeFleet.FleetsClient

			id, err := fleets.ParseFleetID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}

			state := ComputeFleetResourceModel{
				Name:              id.FleetName,
				ResourceGroupName: id.ResourceGroupName,
			}

			if model := resp.Model; model != nil {
				state.Location = location.Normalize(model.Location)
				if props := model.Properties; props != nil {
					flattenComputeFleetProperties(&state, props)
				}
				flattenPlan(&state, model.Plan)
			}

			state.AdminPassword = metadata.ResourceData.Get("admin_password").(string)
			state.CustomData = metadata.ResourceData.Get("custom_data").(string)

			return metadata.Encode(&state)
		},
	}
}

func (r ComputeFleetResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ComputeFleet.FleetsClient

			id, err := fleets.ParseFleetID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ComputeFleetResourceModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			update := fleets.FleetUpdate{
				Properties: expandComputeFleetProperties(model),
			}

			if err := client.UpdateThenPoll(ctx, *id, update); err != nil {
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
			client := metadata.Client.ComputeFleet.FleetsClient

			id, err := fleets.ParseFleetID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			if err := client.DeleteThenPoll(ctx, *id); err != nil {
				return fmt.Errorf("deleting %s: %+v", *id, err)
			}

			return nil
		},
	}
}

func expandComputeFleetProperties(model ComputeFleetResourceModel) *fleets.FleetProperties {
	props := &fleets.FleetProperties{
		ComputeProfile: fleets.ComputeProfile{
			BaseVirtualMachineProfile: fleets.BaseVirtualMachineProfile{},
		},
		VMSizesProfile: expandVmSizeProfiles(model.VmSizeProfiles),
	}

	if model.ComputeApiVersion != "" {
		props.ComputeProfile.ComputeApiVersion = pointer.To(model.ComputeApiVersion)
	}

	spotProfile := &fleets.SpotPriorityProfile{}
	hasSpot := false
	if model.Capacity != 0 {
		spotProfile.Capacity = pointer.To(model.Capacity)
		hasSpot = true
	}
	if model.SpotMinCapacity != 0 {
		spotProfile.MinCapacity = pointer.To(model.SpotMinCapacity)
		hasSpot = true
	}
	if model.MaxPricePerVM != 0 {
		spotProfile.MaxPricePerVM = pointer.To(model.MaxPricePerVM)
		hasSpot = true
	}
	if model.EvictionPolicy != "" {
		ep := fleets.EvictionPolicy(model.EvictionPolicy)
		spotProfile.EvictionPolicy = &ep
		hasSpot = true
	}
	if model.SpotAllocationStrategy != "" {
		s := fleets.SpotAllocationStrategy(model.SpotAllocationStrategy)
		spotProfile.AllocationStrategy = &s
		hasSpot = true
	}
	if model.Maintain {
		spotProfile.Maintain = pointer.To(model.Maintain)
		hasSpot = true
	}
	if hasSpot {
		props.SpotPriorityProfile = spotProfile
	}

	if model.RegularPriorityProfileCapacity != 0 || model.MinCapacity != 0 || model.AllocationStrategy != "" {
		rp := &fleets.RegularPriorityProfile{}
		if model.RegularPriorityProfileCapacity != 0 {
			rp.Capacity = pointer.To(model.RegularPriorityProfileCapacity)
		}
		if model.MinCapacity != 0 {
			rp.MinCapacity = pointer.To(model.MinCapacity)
		}
		if model.AllocationStrategy != "" {
			s := fleets.RegularPriorityAllocationStrategy(model.AllocationStrategy)
			rp.AllocationStrategy = &s
		}
		props.RegularPriorityProfile = rp
	}

	osProfile := expandOsProfile(model)
	if osProfile != nil {
		props.ComputeProfile.BaseVirtualMachineProfile.OsProfile = osProfile
	}

	storageProfile := expandStorageProfile(model)
	if storageProfile != nil {
		props.ComputeProfile.BaseVirtualMachineProfile.StorageProfile = storageProfile
	}

	if len(model.NetworkInterfaceConfiguration) > 0 || model.NetworkApiVersion != "" {
		networkProfile := &fleets.VirtualMachineScaleSetNetworkProfile{}
		if len(model.NetworkInterfaceConfiguration) > 0 {
			networkProfile.NetworkInterfaceConfigurations = pointer.To(expandNetworkInterfaceConfigurations(model.NetworkInterfaceConfiguration))
		}
		if model.NetworkApiVersion != "" {
			nav := fleets.NetworkApiVersion(model.NetworkApiVersion)
			networkProfile.NetworkApiVersion = &nav
		}
		props.ComputeProfile.BaseVirtualMachineProfile.NetworkProfile = networkProfile
	}

	if len(model.GalleryApplications) > 0 {
		props.ComputeProfile.BaseVirtualMachineProfile.ApplicationProfile = &fleets.ApplicationProfile{
			GalleryApplications: pointer.To(expandGalleryApplications(model.GalleryApplications)),
		}
	}

	if len(model.Extensions) > 0 {
		props.ComputeProfile.BaseVirtualMachineProfile.ExtensionProfile = &fleets.VirtualMachineScaleSetExtensionProfile{
			Extensions: pointer.To(expandExtensions(model.Extensions)),
		}
	}

	return props
}

func expandVmSizeProfiles(input []VMSizeProfileModel) []fleets.VMSizeProfile {
	result := make([]fleets.VMSizeProfile, 0, len(input))
	for _, v := range input {
		p := fleets.VMSizeProfile{Name: v.Name}
		if v.Rank != 0 {
			p.Rank = pointer.To(v.Rank)
		}
		result = append(result, p)
	}
	return result
}

func expandGalleryApplications(input []GalleryApplicationModel) []fleets.VMGalleryApplication {
	result := make([]fleets.VMGalleryApplication, 0, len(input))
	for _, v := range input {
		app := fleets.VMGalleryApplication{PackageReferenceId: v.PackageReferenceId}
		if v.EnableAutomaticUpgrade {
			app.EnableAutomaticUpgrade = pointer.To(v.EnableAutomaticUpgrade)
		}
		if v.ConfigurationReference != "" {
			app.ConfigurationReference = pointer.To(v.ConfigurationReference)
		}
		if v.Order != 0 {
			app.Order = pointer.To(v.Order)
		}
		if v.Tags != "" {
			app.Tags = pointer.To(v.Tags)
		}
		if v.TreatFailureAsDeploymentFailure {
			app.TreatFailureAsDeploymentFailure = pointer.To(v.TreatFailureAsDeploymentFailure)
		}
		result = append(result, app)
	}
	return result
}

func expandExtensions(input []ExtensionModel) []fleets.VirtualMachineScaleSetExtension {
	result := make([]fleets.VirtualMachineScaleSetExtension, 0, len(input))
	for _, v := range input {
		ext := fleets.VirtualMachineScaleSetExtension{Name: pointer.To(v.Name)}
		props := &fleets.VirtualMachineScaleSetExtensionProperties{
			Publisher: pointer.To(v.Publisher),
			Type:      pointer.To(v.Type),
		}
		if v.TypeHandlerVersion != "" {
			props.TypeHandlerVersion = pointer.To(v.TypeHandlerVersion)
		}
		if v.AutoUpgradeMinorVersion {
			props.AutoUpgradeMinorVersion = pointer.To(v.AutoUpgradeMinorVersion)
		}
		if v.EnableAutomaticUpgrade {
			props.EnableAutomaticUpgrade = pointer.To(v.EnableAutomaticUpgrade)
		}
		if v.ForceUpdateTag != "" {
			props.ForceUpdateTag = pointer.To(v.ForceUpdateTag)
		}
		if v.ProtectedSettings != "" {
			var ps map[string]interface{}
			if err := json.Unmarshal([]byte(v.ProtectedSettings), &ps); err == nil {
				props.ProtectedSettings = &ps
			}
		}
		// Flattened protected_settings_from_key_vault fields
		if v.SecretUrl != "" || v.SourceVaultId != "" {
			props.ProtectedSettingsFromKeyVault = &fleets.KeyVaultSecretReference{
				SecretURL:   v.SecretUrl,
				SourceVault: fleets.SubResource{Id: pointer.To(v.SourceVaultId)},
			}
		}
		if len(v.ProvisionAfterExtensions) > 0 {
			props.ProvisionAfterExtensions = pointer.To(v.ProvisionAfterExtensions)
		}
		if v.SuppressFailures {
			props.SuppressFailures = pointer.To(v.SuppressFailures)
		}
		if v.Settings != "" {
			var s map[string]interface{}
			if err := json.Unmarshal([]byte(v.Settings), &s); err == nil {
				props.Settings = &s
			}
		}
		ext.Properties = props
		result = append(result, ext)
	}
	return result
}

func expandOsProfile(model ComputeFleetResourceModel) *fleets.VirtualMachineScaleSetOSProfile {
	hasLinux := model.DisablePasswordAuthentication || model.EnableVMAgentPlatformUpdates ||
		model.ProvisionVMAgent || len(model.KeyData) > 0 || model.PatchMode != "" ||
		model.BypassPlatformSafetyChecksOnUserSchedule || model.RebootSetting != ""

	hasWindows := model.EnableAutomaticUpdates || model.WindowsConfigurationEnableVMAgentPlatformUpdates ||
		model.WindowsConfigurationProvisionVMAgent || model.TimeZone != "" ||
		model.WindowsConfigurationPatchMode != "" || model.EnableHotpatching ||
		model.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule ||
		model.WindowsConfigurationRebootSetting != "" ||
		len(model.AdditionalUnattendContent) > 0 || len(model.WinRMListeners) > 0

	hasContent := model.AdminUsername != "" || model.AdminPassword != "" ||
		model.ComputerNamePrefix != "" || model.CustomData != "" ||
		len(model.Secrets) > 0 || hasLinux || hasWindows

	if !hasContent {
		return nil
	}

	osProfile := &fleets.VirtualMachineScaleSetOSProfile{}
	if model.AdminUsername != "" {
		osProfile.AdminUsername = pointer.To(model.AdminUsername)
	}
	if model.AdminPassword != "" {
		osProfile.AdminPassword = pointer.To(model.AdminPassword)
	}
	if model.ComputerNamePrefix != "" {
		osProfile.ComputerNamePrefix = pointer.To(model.ComputerNamePrefix)
	}
	if model.CustomData != "" {
		osProfile.CustomData = pointer.To(model.CustomData)
	}
	if len(model.Secrets) > 0 {
		osProfile.Secrets = pointer.To(expandSecrets(model.Secrets))
	}

	if hasLinux {
		linuxConfig := &fleets.LinuxConfiguration{}
		if model.DisablePasswordAuthentication {
			linuxConfig.DisablePasswordAuthentication = pointer.To(model.DisablePasswordAuthentication)
		}
		if model.EnableVMAgentPlatformUpdates {
			linuxConfig.EnableVMAgentPlatformUpdates = pointer.To(model.EnableVMAgentPlatformUpdates)
		}
		if model.ProvisionVMAgent {
			linuxConfig.ProvisionVMAgent = pointer.To(model.ProvisionVMAgent)
		}
		if len(model.KeyData) > 0 {
			keys := make([]fleets.SshPublicKey, 0, len(model.KeyData))
			for _, k := range model.KeyData {
				k := k
				keys = append(keys, fleets.SshPublicKey{KeyData: pointer.To(k)})
			}
			linuxConfig.Ssh = &fleets.SshConfiguration{PublicKeys: pointer.To(keys)}
		}
		if model.PatchMode != "" || model.BypassPlatformSafetyChecksOnUserSchedule || model.RebootSetting != "" {
			ps := &fleets.LinuxPatchSettings{}
			if model.PatchMode != "" {
				pm := fleets.LinuxVMGuestPatchMode(model.PatchMode)
				ps.PatchMode = &pm
			}
			if model.BypassPlatformSafetyChecksOnUserSchedule || model.RebootSetting != "" {
				auto := &fleets.LinuxVMGuestPatchAutomaticByPlatformSettings{}
				if model.BypassPlatformSafetyChecksOnUserSchedule {
					auto.BypassPlatformSafetyChecksOnUserSchedule = pointer.To(model.BypassPlatformSafetyChecksOnUserSchedule)
				}
				if model.RebootSetting != "" {
					rs := fleets.LinuxVMGuestPatchAutomaticByPlatformRebootSetting(model.RebootSetting)
					auto.RebootSetting = &rs
				}
				ps.AutomaticByPlatformSettings = auto
			}
			linuxConfig.PatchSettings = ps
		}
		osProfile.LinuxConfiguration = linuxConfig
	}

	if hasWindows {
		winConfig := &fleets.WindowsConfiguration{}
		if model.EnableAutomaticUpdates {
			winConfig.EnableAutomaticUpdates = pointer.To(model.EnableAutomaticUpdates)
		}
		if model.WindowsConfigurationEnableVMAgentPlatformUpdates {
			winConfig.EnableVMAgentPlatformUpdates = pointer.To(model.WindowsConfigurationEnableVMAgentPlatformUpdates)
		}
		if model.WindowsConfigurationProvisionVMAgent {
			winConfig.ProvisionVMAgent = pointer.To(model.WindowsConfigurationProvisionVMAgent)
		}
		if model.TimeZone != "" {
			winConfig.TimeZone = pointer.To(model.TimeZone)
		}
		if model.WindowsConfigurationPatchMode != "" || model.EnableHotpatching ||
			model.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule ||
			model.WindowsConfigurationRebootSetting != "" {
			ps := &fleets.PatchSettings{}
			if model.WindowsConfigurationPatchMode != "" {
				pm := fleets.WindowsVMGuestPatchMode(model.WindowsConfigurationPatchMode)
				ps.PatchMode = &pm
			}
			if model.EnableHotpatching {
				ps.EnableHotpatching = pointer.To(model.EnableHotpatching)
			}
			if model.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule || model.WindowsConfigurationRebootSetting != "" {
				auto := &fleets.WindowsVMGuestPatchAutomaticByPlatformSettings{}
				if model.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule {
					auto.BypassPlatformSafetyChecksOnUserSchedule = pointer.To(model.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule)
				}
				if model.WindowsConfigurationRebootSetting != "" {
					rs := fleets.WindowsVMGuestPatchAutomaticByPlatformRebootSetting(model.WindowsConfigurationRebootSetting)
					auto.RebootSetting = &rs
				}
				ps.AutomaticByPlatformSettings = auto
			}
			winConfig.PatchSettings = ps
		}
		if len(model.AdditionalUnattendContent) > 0 {
			contents := make([]fleets.AdditionalUnattendContent, 0, len(model.AdditionalUnattendContent))
			for _, c := range model.AdditionalUnattendContent {
				sn := fleets.SettingNames(c.SettingName)
				contents = append(contents, fleets.AdditionalUnattendContent{
					SettingName: &sn,
					Content:     pointer.To(c.Content),
				})
			}
			winConfig.AdditionalUnattendContent = pointer.To(contents)
		}
		if len(model.WinRMListeners) > 0 {
			listeners := make([]fleets.WinRMListener, 0, len(model.WinRMListeners))
			for _, l := range model.WinRMListeners {
				listener := fleets.WinRMListener{}
				if l.CertificateUrl != "" {
					listener.CertificateURL = pointer.To(l.CertificateUrl)
				}
				if l.Protocol != "" {
					pt := fleets.ProtocolTypes(l.Protocol)
					listener.Protocol = &pt
				}
				listeners = append(listeners, listener)
			}
			winConfig.WinRM = &fleets.WinRMConfiguration{Listeners: pointer.To(listeners)}
		}
		osProfile.WindowsConfiguration = winConfig
	}

	return osProfile
}

func expandSecrets(input []SecretModel) []fleets.VaultSecretGroup {
	result := make([]fleets.VaultSecretGroup, 0, len(input))
	for _, v := range input {
		group := fleets.VaultSecretGroup{
			SourceVault: &fleets.SubResource{Id: pointer.To(v.SourceVaultId)},
		}
		if len(v.CertificateUrl) > 0 {
			certs := make([]fleets.VaultCertificate, 0, len(v.CertificateUrl))
			for _, url := range v.CertificateUrl {
				url := url
				certs = append(certs, fleets.VaultCertificate{CertificateURL: pointer.To(url)})
			}
			group.VaultCertificates = pointer.To(certs)
		}
		result = append(result, group)
	}
	return result
}

func expandStorageProfile(model ComputeFleetResourceModel) *fleets.VirtualMachineScaleSetStorageProfile {
	hasImageRef := model.Offer != "" || model.Publisher != "" || model.Sku != "" || model.Version != ""
	hasDataDisks := len(model.DataDisks) > 0
	hasOsDisk := model.Caching != "" || model.DeleteOption != "" || model.DiskSizeGB != 0 ||
		model.WriteAcceleratorEnabled || model.Option != "" || model.Placement != "" ||
		model.DiskEncryptionSetId != "" || model.SecurityEncryptionType != "" || model.StorageAccountType != ""

	if !hasImageRef && !hasDataDisks && !hasOsDisk {
		return nil
	}

	profile := &fleets.VirtualMachineScaleSetStorageProfile{}

	if hasImageRef {
		imageRef := &fleets.ImageReference{}
		if model.Offer != "" {
			imageRef.Offer = pointer.To(model.Offer)
		}
		if model.Publisher != "" {
			imageRef.Publisher = pointer.To(model.Publisher)
		}
		if model.Sku != "" {
			imageRef.Sku = pointer.To(model.Sku)
		}
		if model.Version != "" {
			imageRef.Version = pointer.To(model.Version)
		}
		profile.ImageReference = imageRef
	}

	if hasDataDisks {
		profile.DataDisks = pointer.To(expandDataDisks(model.DataDisks))
	}

	if hasOsDisk {
		osDisk := &fleets.VirtualMachineScaleSetOSDisk{
			CreateOption: fleets.DiskCreateOptionTypesFromImage,
		}
		if model.Caching != "" {
			ct := fleets.CachingTypes(model.Caching)
			osDisk.Caching = &ct
		}
		if model.DeleteOption != "" {
			do := fleets.DiskDeleteOptionTypes(model.DeleteOption)
			osDisk.DeleteOption = &do
		}
		if model.DiskSizeGB != 0 {
			osDisk.DiskSizeGB = pointer.To(model.DiskSizeGB)
		}
		if model.WriteAcceleratorEnabled {
			osDisk.WriteAcceleratorEnabled = pointer.To(model.WriteAcceleratorEnabled)
		}
		if model.Option != "" || model.Placement != "" {
			diff := &fleets.DiffDiskSettings{}
			if model.Option != "" {
				opt := fleets.DiffDiskOptions(model.Option)
				diff.Option = &opt
			}
			if model.Placement != "" {
				pl := fleets.DiffDiskPlacement(model.Placement)
				diff.Placement = &pl
			}
			osDisk.DiffDiskSettings = diff
		}
		if model.DiskEncryptionSetId != "" || model.StorageAccountType != "" || model.SecurityEncryptionType != "" {
			md := &fleets.VirtualMachineScaleSetManagedDiskParameters{}
			if model.DiskEncryptionSetId != "" {
				md.DiskEncryptionSet = &fleets.DiskEncryptionSetParameters{Id: pointer.To(model.DiskEncryptionSetId)}
			}
			if model.StorageAccountType != "" {
				sat := fleets.StorageAccountTypes(model.StorageAccountType)
				md.StorageAccountType = &sat
			}
			if model.SecurityEncryptionType != "" {
				set := fleets.SecurityEncryptionTypes(model.SecurityEncryptionType)
				md.SecurityProfile = &fleets.VMDiskSecurityProfile{SecurityEncryptionType: &set}
			}
			osDisk.ManagedDisk = md
		}
		profile.OsDisk = osDisk
	}

	return profile
}

func expandDataDisks(input []DataDiskModel) []fleets.VirtualMachineScaleSetDataDisk {
	result := make([]fleets.VirtualMachineScaleSetDataDisk, 0, len(input))
	for _, v := range input {
		disk := fleets.VirtualMachineScaleSetDataDisk{
			Lun:          v.Lun,
			CreateOption: fleets.DiskCreateOptionTypes(v.CreateOption),
		}
		if v.Caching != "" {
			ct := fleets.CachingTypes(v.Caching)
			disk.Caching = &ct
		}
		if v.DeleteOption != "" {
			do := fleets.DiskDeleteOptionTypes(v.DeleteOption)
			disk.DeleteOption = &do
		}
		if v.DiskSizeGB != 0 {
			disk.DiskSizeGB = pointer.To(v.DiskSizeGB)
		}
		if v.WriteAcceleratorEnabled {
			disk.WriteAcceleratorEnabled = pointer.To(v.WriteAcceleratorEnabled)
		}
		if v.DiskEncryptionSetId != "" || v.StorageAccountType != "" {
			md := &fleets.VirtualMachineScaleSetManagedDiskParameters{}
			if v.DiskEncryptionSetId != "" {
				md.DiskEncryptionSet = &fleets.DiskEncryptionSetParameters{Id: pointer.To(v.DiskEncryptionSetId)}
			}
			if v.StorageAccountType != "" {
				sat := fleets.StorageAccountTypes(v.StorageAccountType)
				md.StorageAccountType = &sat
			}
			disk.ManagedDisk = md
		}
		result = append(result, disk)
	}
	return result
}

func expandPlan(model ComputeFleetResourceModel) *fleets.Plan {
	if model.PlanName == "" && model.PlanPublisher == "" && model.Product == "" {
		return nil
	}
	plan := &fleets.Plan{
		Name:      model.PlanName,
		Publisher: model.PlanPublisher,
		Product:   model.Product,
	}
	if model.PromotionCode != "" {
		plan.PromotionCode = pointer.To(model.PromotionCode)
	}
	return plan
}

func expandNetworkInterfaceConfigurations(input []NetworkInterfaceConfigurationModel) []fleets.VirtualMachineScaleSetNetworkConfiguration {
	result := make([]fleets.VirtualMachineScaleSetNetworkConfiguration, 0, len(input))
	for _, v := range input {
		nic := fleets.VirtualMachineScaleSetNetworkConfiguration{Name: v.Name}
		props := &fleets.VirtualMachineScaleSetNetworkConfigurationProperties{
			IPConfigurations: expandIpConfigurations(v.IpConfiguration),
		}
		if v.EnableAcceleratedNetworking {
			props.EnableAcceleratedNetworking = pointer.To(v.EnableAcceleratedNetworking)
		}
		if v.AuxiliaryMode != "" {
			am := fleets.NetworkInterfaceAuxiliaryMode(v.AuxiliaryMode)
			props.AuxiliaryMode = &am
		}
		if v.AuxiliarySku != "" {
			as := fleets.NetworkInterfaceAuxiliarySku(v.AuxiliarySku)
			props.AuxiliarySku = &as
		}
		if v.DeleteOption != "" {
			do := fleets.DeleteOptions(v.DeleteOption)
			props.DeleteOption = &do
		}
		if len(v.DnsServers) > 0 {
			props.DnsSettings = &fleets.VirtualMachineScaleSetNetworkConfigurationDnsSettings{
				DnsServers: pointer.To(v.DnsServers),
			}
		}
		if v.EnableIPForwarding {
			props.EnableIPForwarding = pointer.To(v.EnableIPForwarding)
		}
		if v.NetworkSecurityGroupId != "" {
			props.NetworkSecurityGroup = &fleets.SubResource{Id: pointer.To(v.NetworkSecurityGroupId)}
		}
		if v.Primary {
			props.Primary = pointer.To(v.Primary)
		}
		nic.Properties = props
		result = append(result, nic)
	}
	return result
}

func expandIpConfigurations(input []IpConfigurationModel) []fleets.VirtualMachineScaleSetIPConfiguration {
	result := make([]fleets.VirtualMachineScaleSetIPConfiguration, 0, len(input))
	for _, v := range input {
		ipConfig := fleets.VirtualMachineScaleSetIPConfiguration{Name: v.Name}
		props := &fleets.VirtualMachineScaleSetIPConfigurationProperties{}

		if v.Subnet != "" {
			props.Subnet = &fleets.ApiEntityReference{Id: pointer.To(v.Subnet)}
		}
		if len(v.ApplicationGatewayBackendAddressPoolsId) > 0 {
			pools := make([]fleets.SubResource, 0, len(v.ApplicationGatewayBackendAddressPoolsId))
			for _, id := range v.ApplicationGatewayBackendAddressPoolsId {
				id := id
				pools = append(pools, fleets.SubResource{Id: pointer.To(id)})
			}
			props.ApplicationGatewayBackendAddressPools = &pools
		}
		if len(v.ApplicationSecurityGroupsId) > 0 {
			groups := make([]fleets.SubResource, 0, len(v.ApplicationSecurityGroupsId))
			for _, id := range v.ApplicationSecurityGroupsId {
				id := id
				groups = append(groups, fleets.SubResource{Id: pointer.To(id)})
			}
			props.ApplicationSecurityGroups = &groups
		}
		if len(v.LoadBalancerBackendAddressPoolsId) > 0 {
			pools := make([]fleets.SubResource, 0, len(v.LoadBalancerBackendAddressPoolsId))
			for _, id := range v.LoadBalancerBackendAddressPoolsId {
				id := id
				pools = append(pools, fleets.SubResource{Id: pointer.To(id)})
			}
			props.LoadBalancerBackendAddressPools = &pools
		}
		if v.Primary {
			props.Primary = pointer.To(v.Primary)
		}
		if v.PrivateIPAddressVersion != "" {
			ipv := fleets.IPVersion(v.PrivateIPAddressVersion)
			props.PrivateIPAddressVersion = &ipv
		}
		// Reconstruct public_ip_address_configuration from flattened fields
		if v.PublicIPAddressConfigurationName != "" {
			pubIP := &fleets.VirtualMachineScaleSetPublicIPAddressConfiguration{
				Name: v.PublicIPAddressConfigurationName,
			}
			pubProps := &fleets.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{}
			hasPubProps := false
			if v.DeleteOption != "" {
				do := fleets.DeleteOptions(v.DeleteOption)
				pubProps.DeleteOption = &do
				hasPubProps = true
			}
			if v.DomainNameLabel != "" {
				dns := &fleets.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings{
					DomainNameLabel: v.DomainNameLabel,
				}
				if v.DomainNameLabelScope != "" {
					scope := fleets.DomainNameLabelScopeTypes(v.DomainNameLabelScope)
					dns.DomainNameLabelScope = &scope
				}
				pubProps.DnsSettings = dns
				hasPubProps = true
			}
			if v.IdleTimeoutInMinutes != 0 {
				pubProps.IdleTimeoutInMinutes = pointer.To(v.IdleTimeoutInMinutes)
				hasPubProps = true
			}
			if v.PublicIPAddressVersion != "" {
				ipv := fleets.IPVersion(v.PublicIPAddressVersion)
				pubProps.PublicIPAddressVersion = &ipv
				hasPubProps = true
			}
			if hasPubProps {
				pubIP.Properties = pubProps
			}
			if v.SkuName != "" || v.SkuTier != "" {
				sku := &fleets.PublicIPAddressSku{}
				if v.SkuName != "" {
					sn := fleets.PublicIPAddressSkuName(v.SkuName)
					sku.Name = &sn
				}
				if v.SkuTier != "" {
					st := fleets.PublicIPAddressSkuTier(v.SkuTier)
					sku.Tier = &st
				}
				pubIP.Sku = sku
			}
			props.PublicIPAddressConfiguration = pubIP
		}

		ipConfig.Properties = props
		result = append(result, ipConfig)
	}
	return result
}

func flattenComputeFleetProperties(model *ComputeFleetResourceModel, props *fleets.FleetProperties) {
	model.VmSizeProfiles = flattenVmSizeProfiles(props.VMSizesProfile)

	if props.ComputeProfile.ComputeApiVersion != nil {
		model.ComputeApiVersion = *props.ComputeProfile.ComputeApiVersion
	}

	if props.SpotPriorityProfile != nil {
		sp := props.SpotPriorityProfile
		if sp.Capacity != nil {
			model.Capacity = *sp.Capacity
		}
		if sp.MinCapacity != nil {
			model.SpotMinCapacity = *sp.MinCapacity
		}
		if sp.MaxPricePerVM != nil {
			model.MaxPricePerVM = *sp.MaxPricePerVM
		}
		if sp.EvictionPolicy != nil {
			model.EvictionPolicy = string(*sp.EvictionPolicy)
		}
		if sp.AllocationStrategy != nil {
			model.SpotAllocationStrategy = string(*sp.AllocationStrategy)
		}
		if sp.Maintain != nil {
			model.Maintain = *sp.Maintain
		}
	}

	if props.RegularPriorityProfile != nil {
		rp := props.RegularPriorityProfile
		if rp.Capacity != nil {
			model.RegularPriorityProfileCapacity = *rp.Capacity
		}
		if rp.MinCapacity != nil {
			model.MinCapacity = *rp.MinCapacity
		}
		if rp.AllocationStrategy != nil {
			model.AllocationStrategy = string(*rp.AllocationStrategy)
		}
	}

	bvmp := props.ComputeProfile.BaseVirtualMachineProfile

	if bvmp.OsProfile != nil {
		op := bvmp.OsProfile
		if op.AdminUsername != nil {
			model.AdminUsername = *op.AdminUsername
		}
		if op.ComputerNamePrefix != nil {
			model.ComputerNamePrefix = *op.ComputerNamePrefix
		}
		if op.Secrets != nil {
			model.Secrets = flattenSecrets(*op.Secrets)
		}

		if lc := op.LinuxConfiguration; lc != nil {
			if lc.DisablePasswordAuthentication != nil {
				model.DisablePasswordAuthentication = *lc.DisablePasswordAuthentication
			}
			if lc.EnableVMAgentPlatformUpdates != nil {
				model.EnableVMAgentPlatformUpdates = *lc.EnableVMAgentPlatformUpdates
			}
			if lc.ProvisionVMAgent != nil {
				model.ProvisionVMAgent = *lc.ProvisionVMAgent
			}
			if lc.Ssh != nil && lc.Ssh.PublicKeys != nil {
				keys := make([]string, 0, len(*lc.Ssh.PublicKeys))
				for _, k := range *lc.Ssh.PublicKeys {
					if k.KeyData != nil {
						keys = append(keys, *k.KeyData)
					}
				}
				model.KeyData = keys
			}
			if lc.PatchSettings != nil {
				ps := lc.PatchSettings
				if ps.PatchMode != nil {
					model.PatchMode = string(*ps.PatchMode)
				}
				if ps.AutomaticByPlatformSettings != nil {
					abp := ps.AutomaticByPlatformSettings
					if abp.BypassPlatformSafetyChecksOnUserSchedule != nil {
						model.BypassPlatformSafetyChecksOnUserSchedule = *abp.BypassPlatformSafetyChecksOnUserSchedule
					}
					if abp.RebootSetting != nil {
						model.RebootSetting = string(*abp.RebootSetting)
					}
				}
			}
		}

		if wc := op.WindowsConfiguration; wc != nil {
			if wc.EnableAutomaticUpdates != nil {
				model.EnableAutomaticUpdates = *wc.EnableAutomaticUpdates
			}
			if wc.EnableVMAgentPlatformUpdates != nil {
				model.WindowsConfigurationEnableVMAgentPlatformUpdates = *wc.EnableVMAgentPlatformUpdates
			}
			if wc.ProvisionVMAgent != nil {
				model.WindowsConfigurationProvisionVMAgent = *wc.ProvisionVMAgent
			}
			if wc.TimeZone != nil {
				model.TimeZone = *wc.TimeZone
			}
			if wc.PatchSettings != nil {
				ps := wc.PatchSettings
				if ps.PatchMode != nil {
					model.WindowsConfigurationPatchMode = string(*ps.PatchMode)
				}
				if ps.EnableHotpatching != nil {
					model.EnableHotpatching = *ps.EnableHotpatching
				}
				if ps.AutomaticByPlatformSettings != nil {
					abp := ps.AutomaticByPlatformSettings
					if abp.BypassPlatformSafetyChecksOnUserSchedule != nil {
						model.WindowsConfigurationBypassPlatformSafetyChecksOnUserSchedule = *abp.BypassPlatformSafetyChecksOnUserSchedule
					}
					if abp.RebootSetting != nil {
						model.WindowsConfigurationRebootSetting = string(*abp.RebootSetting)
					}
				}
			}
			if wc.AdditionalUnattendContent != nil {
				contents := make([]AdditionalUnattendContentModel, 0, len(*wc.AdditionalUnattendContent))
				for _, c := range *wc.AdditionalUnattendContent {
					cm := AdditionalUnattendContentModel{}
					if c.SettingName != nil {
						cm.SettingName = string(*c.SettingName)
					}
					if c.Content != nil {
						cm.Content = *c.Content
					}
					contents = append(contents, cm)
				}
				model.AdditionalUnattendContent = contents
			}
			if wc.WinRM != nil && wc.WinRM.Listeners != nil {
				listeners := make([]WinRMListenerModel, 0, len(*wc.WinRM.Listeners))
				for _, l := range *wc.WinRM.Listeners {
					lm := WinRMListenerModel{}
					if l.CertificateURL != nil {
						lm.CertificateUrl = *l.CertificateURL
					}
					if l.Protocol != nil {
						lm.Protocol = string(*l.Protocol)
					}
					listeners = append(listeners, lm)
				}
				model.WinRMListeners = listeners
			}
		}
	}

	if bvmp.StorageProfile != nil {
		sp := bvmp.StorageProfile
		if sp.ImageReference != nil {
			ir := sp.ImageReference
			if ir.Offer != nil {
				model.Offer = *ir.Offer
			}
			if ir.Publisher != nil {
				model.Publisher = *ir.Publisher
			}
			if ir.Sku != nil {
				model.Sku = *ir.Sku
			}
			if ir.Version != nil {
				model.Version = *ir.Version
			}
		}
		if sp.DataDisks != nil {
			model.DataDisks = flattenDataDisks(*sp.DataDisks)
		}
		if sp.OsDisk != nil {
			flattenOsDisk(model, sp.OsDisk)
		}
	}

	if bvmp.NetworkProfile != nil {
		np := bvmp.NetworkProfile
		if np.NetworkInterfaceConfigurations != nil {
			model.NetworkInterfaceConfiguration = flattenNetworkInterfaceConfigurations(*np.NetworkInterfaceConfigurations)
		}
		if np.NetworkApiVersion != nil {
			model.NetworkApiVersion = string(*np.NetworkApiVersion)
		}
	}

	if bvmp.ApplicationProfile != nil && bvmp.ApplicationProfile.GalleryApplications != nil {
		model.GalleryApplications = flattenGalleryApplications(*bvmp.ApplicationProfile.GalleryApplications)
	}

	if bvmp.ExtensionProfile != nil && bvmp.ExtensionProfile.Extensions != nil {
		model.Extensions = flattenExtensions(*bvmp.ExtensionProfile.Extensions)
	}
}

func flattenVmSizeProfiles(input []fleets.VMSizeProfile) []VMSizeProfileModel {
	result := make([]VMSizeProfileModel, 0, len(input))
	for _, v := range input {
		p := VMSizeProfileModel{Name: v.Name}
		if v.Rank != nil {
			p.Rank = *v.Rank
		}
		result = append(result, p)
	}
	return result
}

func flattenSecrets(input []fleets.VaultSecretGroup) []SecretModel {
	result := make([]SecretModel, 0, len(input))
	for _, v := range input {
		s := SecretModel{}
		if v.SourceVault != nil && v.SourceVault.Id != nil {
			s.SourceVaultId = *v.SourceVault.Id
		}
		if v.VaultCertificates != nil {
			urls := make([]string, 0, len(*v.VaultCertificates))
			for _, c := range *v.VaultCertificates {
				if c.CertificateURL != nil {
					urls = append(urls, *c.CertificateURL)
				}
			}
			s.CertificateUrl = urls
		}
		result = append(result, s)
	}
	return result
}

func flattenGalleryApplications(input []fleets.VMGalleryApplication) []GalleryApplicationModel {
	result := make([]GalleryApplicationModel, 0, len(input))
	for _, v := range input {
		app := GalleryApplicationModel{PackageReferenceId: v.PackageReferenceId}
		if v.EnableAutomaticUpgrade != nil {
			app.EnableAutomaticUpgrade = *v.EnableAutomaticUpgrade
		}
		if v.ConfigurationReference != nil {
			app.ConfigurationReference = *v.ConfigurationReference
		}
		if v.Order != nil {
			app.Order = *v.Order
		}
		if v.Tags != nil {
			app.Tags = *v.Tags
		}
		if v.TreatFailureAsDeploymentFailure != nil {
			app.TreatFailureAsDeploymentFailure = *v.TreatFailureAsDeploymentFailure
		}
		result = append(result, app)
	}
	return result
}

func flattenExtensions(input []fleets.VirtualMachineScaleSetExtension) []ExtensionModel {
	result := make([]ExtensionModel, 0, len(input))
	for _, v := range input {
		ext := ExtensionModel{}
		if v.Name != nil {
			ext.Name = *v.Name
		}
		if v.Properties != nil {
			p := v.Properties
			if p.Publisher != nil {
				ext.Publisher = *p.Publisher
			}
			if p.Type != nil {
				ext.Type = *p.Type
			}
			if p.TypeHandlerVersion != nil {
				ext.TypeHandlerVersion = *p.TypeHandlerVersion
			}
			if p.AutoUpgradeMinorVersion != nil {
				ext.AutoUpgradeMinorVersion = *p.AutoUpgradeMinorVersion
			}
			if p.EnableAutomaticUpgrade != nil {
				ext.EnableAutomaticUpgrade = *p.EnableAutomaticUpgrade
			}
			if p.ForceUpdateTag != nil {
				ext.ForceUpdateTag = *p.ForceUpdateTag
			}
			if p.SuppressFailures != nil {
				ext.SuppressFailures = *p.SuppressFailures
			}
			if p.ProvisionAfterExtensions != nil {
				ext.ProvisionAfterExtensions = *p.ProvisionAfterExtensions
			}
			if p.ProtectedSettingsFromKeyVault != nil {
				kv := p.ProtectedSettingsFromKeyVault
				ext.SecretUrl = kv.SecretURL
				if kv.SourceVault.Id != nil {
					ext.SourceVaultId = *kv.SourceVault.Id
				}
			}
			if p.Settings != nil {
				if b, err := json.Marshal(*p.Settings); err == nil {
					ext.Settings = string(b)
				}
			}
		}
		result = append(result, ext)
	}
	return result
}

func flattenDataDisks(input []fleets.VirtualMachineScaleSetDataDisk) []DataDiskModel {
	result := make([]DataDiskModel, 0, len(input))
	for _, v := range input {
		disk := DataDiskModel{
			Lun:          v.Lun,
			CreateOption: string(v.CreateOption),
		}
		if v.Caching != nil {
			disk.Caching = string(*v.Caching)
		}
		if v.DeleteOption != nil {
			disk.DeleteOption = string(*v.DeleteOption)
		}
		if v.DiskSizeGB != nil {
			disk.DiskSizeGB = *v.DiskSizeGB
		}
		if v.WriteAcceleratorEnabled != nil {
			disk.WriteAcceleratorEnabled = *v.WriteAcceleratorEnabled
		}
		if v.ManagedDisk != nil {
			if v.ManagedDisk.DiskEncryptionSet != nil && v.ManagedDisk.DiskEncryptionSet.Id != nil {
				disk.DiskEncryptionSetId = *v.ManagedDisk.DiskEncryptionSet.Id
			}
			if v.ManagedDisk.StorageAccountType != nil {
				disk.StorageAccountType = string(*v.ManagedDisk.StorageAccountType)
			}
		}
		result = append(result, disk)
	}
	return result
}

func flattenOsDisk(model *ComputeFleetResourceModel, input *fleets.VirtualMachineScaleSetOSDisk) {
	if input == nil {
		return
	}
	if input.Caching != nil {
		model.Caching = string(*input.Caching)
	}
	if input.DeleteOption != nil {
		model.DeleteOption = string(*input.DeleteOption)
	}
	if input.DiskSizeGB != nil {
		model.DiskSizeGB = *input.DiskSizeGB
	}
	if input.WriteAcceleratorEnabled != nil {
		model.WriteAcceleratorEnabled = *input.WriteAcceleratorEnabled
	}
	if input.DiffDiskSettings != nil {
		if input.DiffDiskSettings.Option != nil {
			model.Option = string(*input.DiffDiskSettings.Option)
		}
		if input.DiffDiskSettings.Placement != nil {
			model.Placement = string(*input.DiffDiskSettings.Placement)
		}
	}
	if input.ManagedDisk != nil {
		if input.ManagedDisk.DiskEncryptionSet != nil && input.ManagedDisk.DiskEncryptionSet.Id != nil {
			model.DiskEncryptionSetId = *input.ManagedDisk.DiskEncryptionSet.Id
		}
		if input.ManagedDisk.StorageAccountType != nil {
			model.StorageAccountType = string(*input.ManagedDisk.StorageAccountType)
		}
		if input.ManagedDisk.SecurityProfile != nil && input.ManagedDisk.SecurityProfile.SecurityEncryptionType != nil {
			model.SecurityEncryptionType = string(*input.ManagedDisk.SecurityProfile.SecurityEncryptionType)
		}
	}
}

func flattenPlan(model *ComputeFleetResourceModel, input *fleets.Plan) {
	if input == nil {
		return
	}
	model.PlanName = input.Name
	model.PlanPublisher = input.Publisher
	model.Product = input.Product
	if input.PromotionCode != nil {
		model.PromotionCode = *input.PromotionCode
	}
}

func flattenNetworkInterfaceConfigurations(input []fleets.VirtualMachineScaleSetNetworkConfiguration) []NetworkInterfaceConfigurationModel {
	result := make([]NetworkInterfaceConfigurationModel, 0, len(input))
	for _, v := range input {
		nic := NetworkInterfaceConfigurationModel{Name: v.Name}
		if v.Properties != nil {
			p := v.Properties
			nic.IpConfiguration = flattenIpConfigurations(p.IPConfigurations)
			if p.EnableAcceleratedNetworking != nil {
				nic.EnableAcceleratedNetworking = *p.EnableAcceleratedNetworking
			}
			if p.AuxiliaryMode != nil {
				nic.AuxiliaryMode = string(*p.AuxiliaryMode)
			}
			if p.AuxiliarySku != nil {
				nic.AuxiliarySku = string(*p.AuxiliarySku)
			}
			if p.DeleteOption != nil {
				nic.DeleteOption = string(*p.DeleteOption)
			}
			if p.DnsSettings != nil && p.DnsSettings.DnsServers != nil {
				nic.DnsServers = *p.DnsSettings.DnsServers
			}
			if p.EnableIPForwarding != nil {
				nic.EnableIPForwarding = *p.EnableIPForwarding
			}
			if p.NetworkSecurityGroup != nil && p.NetworkSecurityGroup.Id != nil {
				nic.NetworkSecurityGroupId = *p.NetworkSecurityGroup.Id
			}
			if p.Primary != nil {
				nic.Primary = *p.Primary
			}
		}
		result = append(result, nic)
	}
	return result
}

func flattenIpConfigurations(input []fleets.VirtualMachineScaleSetIPConfiguration) []IpConfigurationModel {
	result := make([]IpConfigurationModel, 0, len(input))
	for _, v := range input {
		ipConfig := IpConfigurationModel{
			Name:                                    v.Name,
			ApplicationGatewayBackendAddressPoolsId: make([]string, 0),
			ApplicationSecurityGroupsId:             make([]string, 0),
			LoadBalancerBackendAddressPoolsId:       make([]string, 0),
		}
		if v.Properties != nil {
			p := v.Properties
			if p.Subnet != nil && p.Subnet.Id != nil {
				ipConfig.Subnet = *p.Subnet.Id
			}
			if p.ApplicationGatewayBackendAddressPools != nil {
				ids := make([]string, 0)
				for _, r := range *p.ApplicationGatewayBackendAddressPools {
					if r.Id != nil {
						ids = append(ids, *r.Id)
					}
				}
				ipConfig.ApplicationGatewayBackendAddressPoolsId = ids
			}
			if p.ApplicationSecurityGroups != nil {
				ids := make([]string, 0, len(*p.ApplicationSecurityGroups))
				for _, r := range *p.ApplicationSecurityGroups {
					if r.Id != nil {
						ids = append(ids, *r.Id)
					}
				}
				ipConfig.ApplicationSecurityGroupsId = ids
			}
			if p.LoadBalancerBackendAddressPools != nil {
				ids := make([]string, 0, len(*p.LoadBalancerBackendAddressPools))
				for _, r := range *p.LoadBalancerBackendAddressPools {
					if r.Id != nil {
						ids = append(ids, *r.Id)
					}
				}
				ipConfig.LoadBalancerBackendAddressPoolsId = ids
			}
			if p.Primary != nil {
				ipConfig.Primary = *p.Primary
			}
			if p.PrivateIPAddressVersion != nil {
				ipConfig.PrivateIPAddressVersion = string(*p.PrivateIPAddressVersion)
			}
			// Flatten public_ip_address_configuration fields
			if p.PublicIPAddressConfiguration != nil {
				pub := p.PublicIPAddressConfiguration
				ipConfig.PublicIPAddressConfigurationName = pub.Name
				if pub.Properties != nil {
					pp := pub.Properties
					if pp.DeleteOption != nil {
						ipConfig.DeleteOption = string(*pp.DeleteOption)
					}
					if pp.DnsSettings != nil {
						ipConfig.DomainNameLabel = pp.DnsSettings.DomainNameLabel
						if pp.DnsSettings.DomainNameLabelScope != nil {
							ipConfig.DomainNameLabelScope = string(*pp.DnsSettings.DomainNameLabelScope)
						}
					}
					if pp.IdleTimeoutInMinutes != nil {
						ipConfig.IdleTimeoutInMinutes = *pp.IdleTimeoutInMinutes
					}
					if pp.PublicIPAddressVersion != nil {
						ipConfig.PublicIPAddressVersion = string(*pp.PublicIPAddressVersion)
					}
				}
				if pub.Sku != nil {
					if pub.Sku.Name != nil {
						ipConfig.SkuName = string(*pub.Sku.Name)
					}
					if pub.Sku.Tier != nil {
						ipConfig.SkuTier = string(*pub.Sku.Tier)
					}
				}
			}
		}
		result = append(result, ipConfig)
	}
	return result
}
