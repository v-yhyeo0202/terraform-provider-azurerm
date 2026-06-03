// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package computefleet

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/location"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurefleet/2024-11-01/fleets"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

var (
	_ sdk.Resource           = ComputeFleetResource{}
	_ sdk.ResourceWithUpdate = ComputeFleetResource{}
)

type ComputeFleetResource struct{}

type ComputeFleetResourceModel struct {
	Name                      string                           `tfschema:"name"`
	ResourceGroupName         string                           `tfschema:"resource_group_name"`
	Location                  string                           `tfschema:"location"`
	Zones                     []string                         `tfschema:"zones"`
	BaseVirtualMachineProfile []BaseVirtualMachineProfileModel `tfschema:"base_virtual_machine_profile"`
	VmSizesProfile            []string                         `tfschema:"vm_sizes_profile"`
	SpotPriorityProfile       []SpotPriorityProfileModel       `tfschema:"spot_priority_profile"`
	RegularPriorityProfile    []RegularPriorityProfileModel    `tfschema:"regular_priority_profile"`
	Tags                      map[string]string                `tfschema:"tags"`
}

type BaseVirtualMachineProfileModel struct {
	OsProfile                      []OsProfileModel                     `tfschema:"os_profile"`
	ImageReference                 []ImageReferenceModel                 `tfschema:"image_reference"`
	NetworkInterfaceConfigurations []NetworkInterfaceConfigurationModel  `tfschema:"network_interface_configuration"`
	SecurityProfile                []SecurityProfileModel                `tfschema:"security_profile"`
	LicenseType                    string                                `tfschema:"license_type"`
}

type OsProfileModel struct {
	AdminUsername                 string   `tfschema:"admin_username"`
	AdminPassword                 string   `tfschema:"admin_password"`
	DisablePasswordAuthentication bool     `tfschema:"disable_password_authentication"`
	PublicKey                     []string `tfschema:"public_key"`
}

type ImageReferenceModel struct {
	Publisher string `tfschema:"publisher"`
	Offer     string `tfschema:"offer"`
	Sku       string `tfschema:"sku"`
	Version   string `tfschema:"version"`
}

type NetworkInterfaceConfigurationModel struct {
	Name                        string                 `tfschema:"name"`
	NetworkSecurityGroupId      string                 `tfschema:"network_security_group_id"`
	EnableAcceleratedNetworking bool                   `tfschema:"enable_accelerated_networking"`
	IpConfigurations            []IpConfigurationModel `tfschema:"ip_configuration"`
}

type IpConfigurationModel struct {
	Name                             string   `tfschema:"name"`
	SubnetId                         string   `tfschema:"subnet_id"`
	PublicIpAddressConfigurationName string   `tfschema:"public_ip_address_configuration_name"`
	LoadBalancerBackendAddressPoolIds []string `tfschema:"load_balancer_backend_address_pool_ids"`
	LoadBalancerInboundNatPoolIds    []string `tfschema:"load_balancer_inbound_nat_pool_ids"`
}

type SecurityProfileModel struct {
	SecurityType      string `tfschema:"security_type"`
	SecureBootEnabled bool   `tfschema:"secure_boot_enabled"`
	VTpmEnabled       bool   `tfschema:"v_tpm_enabled"`
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
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
		},

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

								"disable_password_authentication": {
									Type:     pluginsdk.TypeBool,
									Optional: true,
								},

								"public_key": {
									Type:     pluginsdk.TypeList,
									Optional: true,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},
							},
						},
					},

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

											"public_ip_address_configuration_name": {
												Type:     pluginsdk.TypeString,
												Optional: true,
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

					"license_type": {
						Type:     pluginsdk.TypeString,
						Optional: true,
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
			client := metadata.Client.ComputeFleet.FleetsClient

			var model ComputeFleetResourceModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			subscriptionId := metadata.Client.Account.SubscriptionId
			id := fleets.NewFleetID(subscriptionId, model.ResourceGroupName, model.Name)

			existing, err := client.Get(ctx, id)
			if err != nil {
				if !response.WasNotFound(existing.HttpResponse) {
					return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
				}
			}
			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			rawConfig, err := metadata.GetRawConfig()
			if err != nil {
				return fmt.Errorf("getting raw config: %+v", err)
			}

			vmSizesProfile := make([]fleets.VMSizeProfile, 0)
			for _, v := range model.VmSizesProfile {
				vmSizesProfile = append(vmSizesProfile, fleets.VMSizeProfile{Name: v})
			}

			properties := fleets.FleetProperties{
				ComputeProfile: fleets.ComputeProfile{
					BaseVirtualMachineProfile: expandBaseVirtualMachineProfile(model.BaseVirtualMachineProfile[0]),
				},
				VMSizesProfile: vmSizesProfile,
			}

			if len(model.SpotPriorityProfile) > 0 {
				rawSpotSlice := rawConfig.AsValueMap()["spot_priority_profile"].AsValueSlice()
				if len(rawSpotSlice) > 0 {
					properties.SpotPriorityProfile = expandSpotPriorityProfile(model.SpotPriorityProfile[0], rawSpotSlice[0].AsValueMap())
				}
			}

			if len(model.RegularPriorityProfile) > 0 {
				rawRegularSlice := rawConfig.AsValueMap()["regular_priority_profile"].AsValueSlice()
				if len(rawRegularSlice) > 0 {
					properties.RegularPriorityProfile = expandRegularPriorityProfile(model.RegularPriorityProfile[0], rawRegularSlice[0].AsValueMap())
				}
			}

			payload := fleets.Fleet{
				Location:   location.Normalize(model.Location),
				Properties: &properties,
				Tags:       pointer.To(model.Tags),
			}

			if len(model.Zones) > 0 {
				z := zones.Expand(model.Zones)
				payload.Zones = &z
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
					return metadata.MarkAsGone(*id)
				}
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}

			state := ComputeFleetResourceModel{
				Name:              id.FleetName,
				ResourceGroupName: id.ResourceGroupName,
			}

			if model := resp.Model; model != nil {
				state.Location = location.Normalize(model.Location)
				state.Tags = pointer.From(model.Tags)
				state.Zones = zones.Flatten(model.Zones)

				if props := model.Properties; props != nil {
					vmSizes := make([]string, 0)
					for _, v := range props.VMSizesProfile {
						vmSizes = append(vmSizes, v.Name)
					}
					state.VmSizesProfile = vmSizes

					state.BaseVirtualMachineProfile = flattenBaseVirtualMachineProfile(props.ComputeProfile.BaseVirtualMachineProfile)
					state.SpotPriorityProfile = flattenSpotPriorityProfile(props.SpotPriorityProfile)
					state.RegularPriorityProfile = flattenRegularPriorityProfile(props.RegularPriorityProfile)
				}
			}

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

			rawConfig, err := metadata.GetRawConfig()
			if err != nil {
				return fmt.Errorf("getting raw config: %+v", err)
			}

			vmSizesProfile := make([]fleets.VMSizeProfile, 0)
			for _, v := range model.VmSizesProfile {
				vmSizesProfile = append(vmSizesProfile, fleets.VMSizeProfile{Name: v})
			}

			properties := fleets.FleetProperties{
				ComputeProfile: fleets.ComputeProfile{
					BaseVirtualMachineProfile: expandBaseVirtualMachineProfile(model.BaseVirtualMachineProfile[0]),
				},
				VMSizesProfile: vmSizesProfile,
			}

			if len(model.SpotPriorityProfile) > 0 {
				rawSpotSlice := rawConfig.AsValueMap()["spot_priority_profile"].AsValueSlice()
				if len(rawSpotSlice) > 0 {
					properties.SpotPriorityProfile = expandSpotPriorityProfile(model.SpotPriorityProfile[0], rawSpotSlice[0].AsValueMap())
				}
			}

			if len(model.RegularPriorityProfile) > 0 {
				rawRegularSlice := rawConfig.AsValueMap()["regular_priority_profile"].AsValueSlice()
				if len(rawRegularSlice) > 0 {
					properties.RegularPriorityProfile = expandRegularPriorityProfile(model.RegularPriorityProfile[0], rawRegularSlice[0].AsValueMap())
				}
			}

			payload := fleets.FleetUpdate{
				Properties: &properties,
				Tags:       pointer.To(model.Tags),
			}

			if err := client.UpdateThenPoll(ctx, *id, payload); err != nil {
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

func expandBaseVirtualMachineProfile(input BaseVirtualMachineProfileModel) fleets.BaseVirtualMachineProfile {
	result := fleets.BaseVirtualMachineProfile{}

	if len(input.OsProfile) > 0 {
		result.OsProfile = expandOsProfile(input.OsProfile[0])
	}

	if len(input.ImageReference) > 0 {
		result.StorageProfile = &fleets.VirtualMachineScaleSetStorageProfile{
			ImageReference: expandImageReference(input.ImageReference[0]),
		}
	}

	if len(input.NetworkInterfaceConfigurations) > 0 {
		result.NetworkProfile = &fleets.VirtualMachineScaleSetNetworkProfile{
			NetworkInterfaceConfigurations: expandNetworkInterfaceConfigurations(input.NetworkInterfaceConfigurations),
		}
	}

	if len(input.SecurityProfile) > 0 {
		result.SecurityProfile = expandSecurityProfile(input.SecurityProfile[0])
	}

	if input.LicenseType != "" {
		result.LicenseType = pointer.To(input.LicenseType)
	}

	return result
}

func expandOsProfile(input OsProfileModel) *fleets.VirtualMachineScaleSetOSProfile {
	result := &fleets.VirtualMachineScaleSetOSProfile{}

	if input.AdminUsername != "" {
		result.AdminUsername = pointer.To(input.AdminUsername)
	}

	if input.AdminPassword != "" {
		result.AdminPassword = pointer.To(input.AdminPassword)
	}

	if input.DisablePasswordAuthentication || len(input.PublicKey) > 0 {
		linuxConfig := &fleets.LinuxConfiguration{}

		if input.DisablePasswordAuthentication {
			linuxConfig.DisablePasswordAuthentication = pointer.To(input.DisablePasswordAuthentication)
		}

		if len(input.PublicKey) > 0 {
			publicKeys := make([]fleets.SshPublicKey, 0, len(input.PublicKey))
			for _, key := range input.PublicKey {
				publicKeys = append(publicKeys, fleets.SshPublicKey{
					KeyData: pointer.To(key),
				})
			}
			linuxConfig.Ssh = &fleets.SshConfiguration{
				PublicKeys: &publicKeys,
			}
		}

		result.LinuxConfiguration = linuxConfig
	}

	return result
}

func expandImageReference(input ImageReferenceModel) *fleets.ImageReference {
	result := &fleets.ImageReference{}

	if input.Publisher != "" {
		result.Publisher = pointer.To(input.Publisher)
	}

	if input.Offer != "" {
		result.Offer = pointer.To(input.Offer)
	}

	if input.Sku != "" {
		result.Sku = pointer.To(input.Sku)
	}

	if input.Version != "" {
		result.Version = pointer.To(input.Version)
	}

	return result
}

func expandNetworkInterfaceConfigurations(input []NetworkInterfaceConfigurationModel) *[]fleets.VirtualMachineScaleSetNetworkConfiguration {
	results := make([]fleets.VirtualMachineScaleSetNetworkConfiguration, 0, len(input))

	for _, v := range input {
		config := fleets.VirtualMachineScaleSetNetworkConfiguration{
			Name:       v.Name,
			Properties: expandNetworkInterfaceConfigurationProperties(v),
		}
		results = append(results, config)
	}

	return &results
}

func expandNetworkInterfaceConfigurationProperties(input NetworkInterfaceConfigurationModel) *fleets.VirtualMachineScaleSetNetworkConfigurationProperties {
	props := &fleets.VirtualMachineScaleSetNetworkConfigurationProperties{
		IPConfigurations: expandIpConfigurations(input.IpConfigurations),
	}

	if input.NetworkSecurityGroupId != "" {
		props.NetworkSecurityGroup = &fleets.SubResource{
			Id: pointer.To(input.NetworkSecurityGroupId),
		}
	}

	if input.EnableAcceleratedNetworking {
		props.EnableAcceleratedNetworking = pointer.To(input.EnableAcceleratedNetworking)
	}

	return props
}

func expandIpConfigurations(input []IpConfigurationModel) []fleets.VirtualMachineScaleSetIPConfiguration {
	results := make([]fleets.VirtualMachineScaleSetIPConfiguration, 0, len(input))

	for _, v := range input {
		config := fleets.VirtualMachineScaleSetIPConfiguration{
			Name:       v.Name,
			Properties: expandIpConfigurationProperties(v),
		}
		results = append(results, config)
	}

	return results
}

func expandIpConfigurationProperties(input IpConfigurationModel) *fleets.VirtualMachineScaleSetIPConfigurationProperties {
	props := &fleets.VirtualMachineScaleSetIPConfigurationProperties{}

	if input.SubnetId != "" {
		props.Subnet = &fleets.ApiEntityReference{
			Id: pointer.To(input.SubnetId),
		}
	}

	if input.PublicIpAddressConfigurationName != "" {
		props.PublicIPAddressConfiguration = &fleets.VirtualMachineScaleSetPublicIPAddressConfiguration{
			Name: input.PublicIpAddressConfigurationName,
		}
	}

	if len(input.LoadBalancerBackendAddressPoolIds) > 0 {
		pools := make([]fleets.SubResource, 0, len(input.LoadBalancerBackendAddressPoolIds))
		for _, id := range input.LoadBalancerBackendAddressPoolIds {
			pools = append(pools, fleets.SubResource{Id: pointer.To(id)})
		}
		props.LoadBalancerBackendAddressPools = &pools
	}

	if len(input.LoadBalancerInboundNatPoolIds) > 0 {
		natPools := make([]fleets.SubResource, 0, len(input.LoadBalancerInboundNatPoolIds))
		for _, id := range input.LoadBalancerInboundNatPoolIds {
			natPools = append(natPools, fleets.SubResource{Id: pointer.To(id)})
		}
		props.LoadBalancerInboundNatPools = &natPools
	}

	return props
}

func expandSecurityProfile(input SecurityProfileModel) *fleets.SecurityProfile {
	result := &fleets.SecurityProfile{}

	if input.SecurityType != "" {
		result.SecurityType = pointer.To(fleets.SecurityTypes(input.SecurityType))
	}

	if input.SecureBootEnabled || input.VTpmEnabled {
		uefiSettings := &fleets.UefiSettings{}

		if input.SecureBootEnabled {
			uefiSettings.SecureBootEnabled = pointer.To(input.SecureBootEnabled)
		}

		if input.VTpmEnabled {
			uefiSettings.VTpmEnabled = pointer.To(input.VTpmEnabled)
		}

		result.UefiSettings = uefiSettings
	}

	return result
}

func expandSpotPriorityProfile(input SpotPriorityProfileModel, rawConfig map[string]cty.Value) *fleets.SpotPriorityProfile {
	result := &fleets.SpotPriorityProfile{}

	if v, ok := rawConfig["capacity"]; ok && !v.IsNull() {
		result.Capacity = pointer.To(int64(input.Capacity))
	}

	if v, ok := rawConfig["min_capacity"]; ok && !v.IsNull() {
		result.MinCapacity = pointer.To(int64(input.MinCapacity))
	}

	if input.Maintain {
		result.Maintain = pointer.To(input.Maintain)
	}

	if input.EvictionPolicy != "" {
		result.EvictionPolicy = pointer.To(fleets.EvictionPolicy(input.EvictionPolicy))
	}

	if input.AllocationStrategy != "" {
		result.AllocationStrategy = pointer.To(fleets.SpotAllocationStrategy(input.AllocationStrategy))
	}

	if input.MaxPricePerVM != 0 {
		result.MaxPricePerVM = pointer.To(input.MaxPricePerVM)
	}

	return result
}

func expandRegularPriorityProfile(input RegularPriorityProfileModel, rawConfig map[string]cty.Value) *fleets.RegularPriorityProfile {
	result := &fleets.RegularPriorityProfile{}

	if v, ok := rawConfig["capacity"]; ok && !v.IsNull() {
		result.Capacity = pointer.To(int64(input.Capacity))
	}

	if v, ok := rawConfig["min_capacity"]; ok && !v.IsNull() {
		result.MinCapacity = pointer.To(int64(input.MinCapacity))
	}

	if input.AllocationStrategy != "" {
		result.AllocationStrategy = pointer.To(fleets.RegularPriorityAllocationStrategy(input.AllocationStrategy))
	}

	return result
}

func flattenBaseVirtualMachineProfile(input fleets.BaseVirtualMachineProfile) []BaseVirtualMachineProfileModel {
	result := BaseVirtualMachineProfileModel{
		OsProfile:                      flattenOsProfile(input.OsProfile),
		NetworkInterfaceConfigurations: flattenNetworkInterfaceConfigurations(input.NetworkProfile),
		SecurityProfile:                flattenSecurityProfile(input.SecurityProfile),
		LicenseType:                    pointer.From(input.LicenseType),
	}

	if input.StorageProfile != nil {
		result.ImageReference = flattenImageReference(input.StorageProfile.ImageReference)
	}

	return []BaseVirtualMachineProfileModel{result}
}

func flattenOsProfile(input *fleets.VirtualMachineScaleSetOSProfile) []OsProfileModel {
	if input == nil {
		return []OsProfileModel{}
	}

	result := OsProfileModel{
		AdminUsername: pointer.From(input.AdminUsername),
		AdminPassword: pointer.From(input.AdminPassword),
	}

	if linuxConfig := input.LinuxConfiguration; linuxConfig != nil {
		result.DisablePasswordAuthentication = pointer.From(linuxConfig.DisablePasswordAuthentication)

		if ssh := linuxConfig.Ssh; ssh != nil && ssh.PublicKeys != nil {
			keys := make([]string, 0, len(*ssh.PublicKeys))
			for _, k := range *ssh.PublicKeys {
				keys = append(keys, pointer.From(k.KeyData))
			}
			result.PublicKey = keys
		}
	}

	return []OsProfileModel{result}
}

func flattenImageReference(input *fleets.ImageReference) []ImageReferenceModel {
	if input == nil {
		return []ImageReferenceModel{}
	}

	return []ImageReferenceModel{{
		Publisher: pointer.From(input.Publisher),
		Offer:     pointer.From(input.Offer),
		Sku:       pointer.From(input.Sku),
		Version:   pointer.From(input.Version),
	}}
}

func flattenNetworkInterfaceConfigurations(input *fleets.VirtualMachineScaleSetNetworkProfile) []NetworkInterfaceConfigurationModel {
	if input == nil || input.NetworkInterfaceConfigurations == nil {
		return []NetworkInterfaceConfigurationModel{}
	}

	results := make([]NetworkInterfaceConfigurationModel, 0, len(*input.NetworkInterfaceConfigurations))
	for _, v := range *input.NetworkInterfaceConfigurations {
		config := NetworkInterfaceConfigurationModel{
			Name:             v.Name,
			IpConfigurations: flattenIpConfigurations(v.Properties),
		}

		if props := v.Properties; props != nil {
			if props.NetworkSecurityGroup != nil {
				config.NetworkSecurityGroupId = pointer.From(props.NetworkSecurityGroup.Id)
			}
			config.EnableAcceleratedNetworking = pointer.From(props.EnableAcceleratedNetworking)
		}

		results = append(results, config)
	}

	return results
}

func flattenIpConfigurations(input *fleets.VirtualMachineScaleSetNetworkConfigurationProperties) []IpConfigurationModel {
	if input == nil {
		return []IpConfigurationModel{}
	}

	results := make([]IpConfigurationModel, 0, len(input.IPConfigurations))
	for _, v := range input.IPConfigurations {
		config := IpConfigurationModel{
			Name: v.Name,
		}

		if props := v.Properties; props != nil {
			if props.Subnet != nil {
				config.SubnetId = pointer.From(props.Subnet.Id)
			}

			if props.PublicIPAddressConfiguration != nil {
				config.PublicIpAddressConfigurationName = props.PublicIPAddressConfiguration.Name
			}

			if props.LoadBalancerBackendAddressPools != nil {
				poolIds := make([]string, 0, len(*props.LoadBalancerBackendAddressPools))
				for _, p := range *props.LoadBalancerBackendAddressPools {
					poolIds = append(poolIds, pointer.From(p.Id))
				}
				config.LoadBalancerBackendAddressPoolIds = poolIds
			}

			if props.LoadBalancerInboundNatPools != nil {
				natPoolIds := make([]string, 0, len(*props.LoadBalancerInboundNatPools))
				for _, p := range *props.LoadBalancerInboundNatPools {
					natPoolIds = append(natPoolIds, pointer.From(p.Id))
				}
				config.LoadBalancerInboundNatPoolIds = natPoolIds
			}
		}

		results = append(results, config)
	}

	return results
}

func flattenSecurityProfile(input *fleets.SecurityProfile) []SecurityProfileModel {
	if input == nil {
		return []SecurityProfileModel{}
	}

	result := SecurityProfileModel{
		SecurityType: string(pointer.From(input.SecurityType)),
	}

	if input.UefiSettings != nil {
		result.SecureBootEnabled = pointer.From(input.UefiSettings.SecureBootEnabled)
		result.VTpmEnabled = pointer.From(input.UefiSettings.VTpmEnabled)
	}

	return []SecurityProfileModel{result}
}

func flattenSpotPriorityProfile(input *fleets.SpotPriorityProfile) []SpotPriorityProfileModel {
	if input == nil {
		return []SpotPriorityProfileModel{}
	}

	return []SpotPriorityProfileModel{{
		Capacity:           int(pointer.From(input.Capacity)),
		MinCapacity:        int(pointer.From(input.MinCapacity)),
		Maintain:           pointer.From(input.Maintain),
		EvictionPolicy:     string(pointer.From(input.EvictionPolicy)),
		AllocationStrategy: string(pointer.From(input.AllocationStrategy)),
		MaxPricePerVM:      pointer.From(input.MaxPricePerVM),
	}}
}

func flattenRegularPriorityProfile(input *fleets.RegularPriorityProfile) []RegularPriorityProfileModel {
	if input == nil {
		return []RegularPriorityProfileModel{}
	}

	return []RegularPriorityProfileModel{{
		Capacity:           int(pointer.From(input.Capacity)),
		MinCapacity:        int(pointer.From(input.MinCapacity)),
		AllocationStrategy: string(pointer.From(input.AllocationStrategy)),
	}}
}
