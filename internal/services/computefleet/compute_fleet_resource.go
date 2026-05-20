// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package computefleet

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurefleet/2024-11-01/fleets"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

var _ sdk.Resource = ComputeFleetResource{}

type ComputeFleetResource struct{}

type ComputeFleetResourceModel struct {
	Name                   string                        `tfschema:"name"`
	ResourceGroupName      string                        `tfschema:"resource_group_name"`
	Location               string                        `tfschema:"location"`
	Zones                  []string                      `tfschema:"zones"`
	ComputeProfile         []ComputeProfileModel         `tfschema:"compute_profile"`
	VmSizesProfile         []VmSizeProfileModel          `tfschema:"vm_sizes_profile"`
	SpotPriorityProfile    []SpotPriorityProfileModel    `tfschema:"spot_priority_profile"`
	RegularPriorityProfile []RegularPriorityProfileModel `tfschema:"regular_priority_profile"`
	Tags                   map[string]string             `tfschema:"tags"`
}

type ComputeProfileModel struct {
	BaseVirtualMachineProfile []BaseVirtualMachineProfileModel `tfschema:"base_virtual_machine_profile"`
}

type BaseVirtualMachineProfileModel struct {
	OsProfile       []OsProfileModel       `tfschema:"os_profile"`
	StorageProfile  []StorageProfileModel  `tfschema:"storage_profile"`
	NetworkProfile  []NetworkProfileModel  `tfschema:"network_profile"`
	SecurityProfile []SecurityProfileModel `tfschema:"security_profile"`
	LicenseType     string                 `tfschema:"license_type"`
}

type OsProfileModel struct {
	AdminUsername      string                   `tfschema:"admin_username"`
	AdminPassword      string                   `tfschema:"admin_password"`
	LinuxConfiguration []LinuxConfigurationModel `tfschema:"linux_configuration"`
}

type LinuxConfigurationModel struct {
	DisablePasswordAuthentication bool               `tfschema:"disable_password_authentication"`
	Ssh                           []SshConfigModel   `tfschema:"ssh"`
}

type SshConfigModel struct {
	PublicKeys []SshPublicKeyModel `tfschema:"public_key"`
}

type SshPublicKeyModel struct {
	KeyData string `tfschema:"key_data"`
}

type StorageProfileModel struct {
	ImageReference []ImageReferenceModel `tfschema:"image_reference"`
}

type ImageReferenceModel struct {
	Publisher string `tfschema:"publisher"`
	Offer     string `tfschema:"offer"`
	Sku       string `tfschema:"sku"`
	Version   string `tfschema:"version"`
}

type NetworkProfileModel struct {
	NetworkInterfaceConfigurations []NetworkInterfaceConfigurationModel `tfschema:"network_interface_configuration"`
}

type NetworkInterfaceConfigurationModel struct {
	Name                     string                  `tfschema:"name"`
	NetworkSecurityGroupId   string                  `tfschema:"network_security_group_id"`
	EnableAcceleratedNetworking bool                 `tfschema:"enable_accelerated_networking"`
	IpConfigurations         []IpConfigurationModel  `tfschema:"ip_configuration"`
}

type IpConfigurationModel struct {
	Name                             string                          `tfschema:"name"`
	SubnetId                         string                          `tfschema:"subnet_id"`
	PublicIpAddressConfiguration     []PublicIpAddressConfigModel    `tfschema:"public_ip_address_configuration"`
	LoadBalancerBackendAddressPoolIds []string                       `tfschema:"load_balancer_backend_address_pool_ids"`
	LoadBalancerInboundNatPoolIds    []string                        `tfschema:"load_balancer_inbound_nat_pool_ids"`
}

type PublicIpAddressConfigModel struct {
	Name string `tfschema:"name"`
}

type SecurityProfileModel struct {
	SecurityType string             `tfschema:"security_type"`
	UefiSettings []UefiSettingsModel `tfschema:"uefi_settings"`
}

type UefiSettingsModel struct {
	SecureBootEnabled bool `tfschema:"secure_boot_enabled"`
	VTpmEnabled       bool `tfschema:"v_tpm_enabled"`
}

type VmSizeProfileModel struct {
	Name string `tfschema:"name"`
}

type SpotPriorityProfileModel struct {
	Capacity           int     `tfschema:"capacity"`
	MinCapacity        int     `tfschema:"min_capacity"`
	Maintain           bool    `tfschema:"maintain"`
	EvictionPolicy     string  `tfschema:"eviction_policy"`
	AllocationStrategy string  `tfschema:"allocation_strategy"`
	MaxPricePerVM      float64 `tfschema:"max_price_per_vm"`
}

type RegularPriorityProfileModel struct {
	Capacity           int    `tfschema:"capacity"`
	MinCapacity        int    `tfschema:"min_capacity"`
	AllocationStrategy string `tfschema:"allocation_strategy"`
}

func (r ComputeFleetResource) ModelObject() interface{} {
	return &ComputeFleetResourceModel{}
}

func (r ComputeFleetResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return fleets.ValidateFleetID
}

func (r ComputeFleetResource) ResourceType() string {
	return "azurerm_compute_fleet"
}

func (r ComputeFleetResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"resource_group_name": commonschema.ResourceGroupName(),

		"location": commonschema.Location(),

		"vm_sizes_profile": {
			Type:     pluginsdk.TypeList,
			Required: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"name": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
				},
			},
		},

		"compute_profile": {
			Type:     pluginsdk.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"base_virtual_machine_profile": {
						Type:     pluginsdk.TypeList,
						Required: true,
						MaxItems: 1,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"os_profile": {
									Type:     pluginsdk.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"admin_username": {
												Type:     pluginsdk.TypeString,
												Optional: true,
											},

											"admin_password": {
												Type:     pluginsdk.TypeString,
												Optional: true,
											},

											"linux_configuration": {
												Type:     pluginsdk.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &pluginsdk.Resource{
													Schema: map[string]*pluginsdk.Schema{
														"disable_password_authentication": {
															Type:     pluginsdk.TypeBool,
															Optional: true,
														},

														"ssh": {
															Type:     pluginsdk.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &pluginsdk.Resource{
																Schema: map[string]*pluginsdk.Schema{
																	"public_key": {
																		Type:     pluginsdk.TypeList,
																		Optional: true,
																		Elem: &pluginsdk.Resource{
																			Schema: map[string]*pluginsdk.Schema{
																				"key_data": {
																					Type:     pluginsdk.TypeString,
																					Optional: true,
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},

								"storage_profile": {
									Type:     pluginsdk.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"image_reference": {
												Type:     pluginsdk.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &pluginsdk.Resource{
													Schema: map[string]*pluginsdk.Schema{
														"publisher": {
															Type:     pluginsdk.TypeString,
															Optional: true,
														},

														"offer": {
															Type:     pluginsdk.TypeString,
															Optional: true,
														},

														"sku": {
															Type:     pluginsdk.TypeString,
															Optional: true,
														},

														"version": {
															Type:     pluginsdk.TypeString,
															Optional: true,
														},
													},
												},
											},
										},
									},
								},

								"network_profile": {
									Type:     pluginsdk.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"network_interface_configuration": {
												Type:     pluginsdk.TypeList,
												Optional: true,
												Elem: &pluginsdk.Resource{
													Schema: map[string]*pluginsdk.Schema{
														"name": {
															Type:     pluginsdk.TypeString,
															Required: true,
														},

														"network_security_group_id": {
															Type:     pluginsdk.TypeString,
															Optional: true,
														},

														"enable_accelerated_networking": {
															Type:     pluginsdk.TypeBool,
															Optional: true,
														},

														"ip_configuration": {
															Type:     pluginsdk.TypeList,
															Required: true,
															Elem: &pluginsdk.Resource{
																Schema: map[string]*pluginsdk.Schema{
																	"name": {
																		Type:     pluginsdk.TypeString,
																		Required: true,
																	},

																	"subnet_id": {
																		Type:     pluginsdk.TypeString,
																		Optional: true,
																	},

																	"public_ip_address_configuration": {
																		Type:     pluginsdk.TypeList,
																		Optional: true,
																		MaxItems: 1,
																		Elem: &pluginsdk.Resource{
																			Schema: map[string]*pluginsdk.Schema{
																				"name": {
																					Type:     pluginsdk.TypeString,
																					Required: true,
																				},
																			},
																		},
																	},

																	"load_balancer_backend_address_pool_ids": {
																		Type:     pluginsdk.TypeList,
																		Optional: true,
																		Elem: &pluginsdk.Schema{
																			Type: pluginsdk.TypeString,
																		},
																	},

																	"load_balancer_inbound_nat_pool_ids": {
																		Type:     pluginsdk.TypeList,
																		Optional: true,
																		Elem: &pluginsdk.Schema{
																			Type: pluginsdk.TypeString,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},

								"security_profile": {
									Type:     pluginsdk.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"security_type": {
												Type:     pluginsdk.TypeString,
												Optional: true,
											},

											"uefi_settings": {
												Type:     pluginsdk.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &pluginsdk.Resource{
													Schema: map[string]*pluginsdk.Schema{
														"secure_boot_enabled": {
															Type:     pluginsdk.TypeBool,
															Optional: true,
														},

														"v_tpm_enabled": {
															Type:     pluginsdk.TypeBool,
															Optional: true,
														},
													},
												},
											},
										},
									},
								},

								"license_type": {
									Type:     pluginsdk.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},

		"spot_priority_profile": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"capacity": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},

					"min_capacity": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},

					"maintain": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},

					"eviction_policy": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"allocation_strategy": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"max_price_per_vm": {
						Type:     pluginsdk.TypeFloat,
						Optional: true,
					},
				},
			},
		},

		"regular_priority_profile": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"capacity": {
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
				},
			},
		},

		"zones": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
		},

		"tags": commonschema.Tags(),
	}
}

func (r ComputeFleetResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ComputeFleetResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return fmt.Errorf("not implemented")
		},
	}
}

func (r ComputeFleetResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return fmt.Errorf("not implemented")
		},
	}
}

func (r ComputeFleetResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return fmt.Errorf("not implemented")
		},
	}
}

func (r ComputeFleetResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return fmt.Errorf("not implemented")
		},
	}
}
