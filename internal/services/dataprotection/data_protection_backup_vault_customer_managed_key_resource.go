// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotection

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/backupvaults"
	"github.com/hashicorp/terraform-provider-azurerm/internal/locks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	keyVaultParse "github.com/hashicorp/terraform-provider-azurerm/internal/services/keyvault/parse"
	keyVaultValidate "github.com/hashicorp/terraform-provider-azurerm/internal/services/keyvault/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type DataProtectionBackupVaultCustomerManagedKeyResource struct{}

type DataProtectionBackupVaultCustomerManagedKeyModel struct {
	DataProtectionBackupVaultID     string `tfschema:"data_protection_backup_vault_id"`
	KeyVaultKeyID                   string `tfschema:"key_vault_key_id"`
	InfrastructureEncryptionEnabled bool   `tfscheme:"infrastructure_encryption_enabled"`
}

var _ sdk.ResourceWithUpdate = DataProtectionBackupVaultCustomerManagedKeyResource{}

// var _ sdk.ResourceWithCustomizeDiff = DataProtectionBackupVaultCustomerManagedKeyResource{}

func (r DataProtectionBackupVaultCustomerManagedKeyResource) ModelObject() interface{} {
	return &DataProtectionBackupVaultCustomerManagedKeyResource{}
}

func (r DataProtectionBackupVaultCustomerManagedKeyResource) ResourceType() string {
	return "azurerm_data_protection_backup_vault_customer_managed_key"
}

func (r DataProtectionBackupVaultCustomerManagedKeyResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return backupvaults.ValidateBackupVaultID
}

func (r DataProtectionBackupVaultCustomerManagedKeyResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"data_protection_backup_vault_id": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: backupvaults.ValidateBackupVaultID,
		},

		"key_vault_key_id": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: keyVaultValidate.NestedItemIdWithOptionalVersion,
		},

		"infrastructure_encryption_enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  false,
			ForceNew: true,
		},
	}
}

func (r DataProtectionBackupVaultCustomerManagedKeyResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r DataProtectionBackupVaultCustomerManagedKeyResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DataProtection.BackupVaultClient
			var cmk DataProtectionBackupVaultCustomerManagedKeyModel

			if err := metadata.Decode(&cmk); err != nil {
				return err
			}
			fmt.Println("debug11 ", cmk.InfrastructureEncryptionEnabled)
			id, err := backupvaults.ParseBackupVaultID(cmk.DataProtectionBackupVaultID)
			if err != nil {
				return err
			}
			fmt.Println("debug12 ", cmk.InfrastructureEncryptionEnabled)
			locks.ByID(id.ID())
			defer locks.UnlockByID(id.ID())
			fmt.Println("debug13 ", cmk.InfrastructureEncryptionEnabled)
			resp, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("%s was not found", *id)
				}
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}
			fmt.Println("debug14 ", cmk.InfrastructureEncryptionEnabled)
			if resp.Model == nil {
				return fmt.Errorf("retrieving %s: `model` is nil", *id)
			}
			fmt.Println("debug15 ", cmk.InfrastructureEncryptionEnabled)
			if resp.Model.Properties.SecuritySettings != nil && resp.Model.Properties.SecuritySettings.EncryptionSettings != nil {
				return metadata.ResourceRequiresImport(r.ResourceType(), *id)
			}
			fmt.Println("debug16 ", cmk.InfrastructureEncryptionEnabled)
			payload := resp.Model

			keyId, err := keyVaultParse.ParseOptionallyVersionedNestedItemID(cmk.KeyVaultKeyID)
			if err != nil {
				return err
			}
			fmt.Println("debug17 ", cmk.InfrastructureEncryptionEnabled)
			payload.Properties.SecuritySettings.EncryptionSettings = &backupvaults.EncryptionSettings{
				State: pointer.To(backupvaults.EncryptionStateEnabled),
				KeyVaultProperties: &backupvaults.CmkKeyVaultProperties{
					KeyUri: pointer.To(keyId.ID()),
				},
			}

			fmt.Println("debug5 ", cmk.InfrastructureEncryptionEnabled)

			if cmk.InfrastructureEncryptionEnabled {
				fmt.Println("debug6 ", backupvaults.InfrastructureEncryptionStateEnabled)
				payload.Properties.SecuritySettings.EncryptionSettings.InfrastructureEncryption = pointer.To(backupvaults.InfrastructureEncryptionStateEnabled)
			} else {
				fmt.Println("debug7 ", backupvaults.InfrastructureEncryptionStateDisabled)
				payload.Properties.SecuritySettings.EncryptionSettings.InfrastructureEncryption = pointer.To(backupvaults.InfrastructureEncryptionStateDisabled)
			}

			fmt.Println("debug8 ", pointer.From(payload.Properties.SecuritySettings.EncryptionSettings.InfrastructureEncryption))

			payload.Properties.SecuritySettings.EncryptionSettings.KekIdentity = &backupvaults.CmkKekIdentity{
				IdentityType: pointer.To(backupvaults.IdentityTypeSystemAssigned),
			}

			err = client.CreateOrUpdateThenPoll(ctx, *id, *payload, backupvaults.DefaultCreateOrUpdateOperationOptions())
			if err != nil {
				return fmt.Errorf("creating Customer Managed Key for %s: %+v", *id, err)
			}

			metadata.SetID(id)

			return nil
		},
	}
}

func (r DataProtectionBackupVaultCustomerManagedKeyResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DataProtection.BackupVaultClient

			id, err := backupvaults.ParseBackupVaultID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			existing, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(existing.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}

			var state DataProtectionBackupVaultCustomerManagedKeyModel
			state.DataProtectionBackupVaultID = id.ID()

			fmt.Println("debug2")

			if model := existing.Model; model != nil {
				fmt.Println("debug3")
				props := model.Properties
				if props.SecuritySettings != nil && props.SecuritySettings.EncryptionSettings != nil {
					fmt.Println("debug4")
					if props.SecuritySettings.EncryptionSettings.KeyVaultProperties != nil {
						state.KeyVaultKeyID = pointer.From(props.SecuritySettings.EncryptionSettings.KeyVaultProperties.KeyUri)
					}

					fmt.Println("debug1 ", string(pointer.From(props.SecuritySettings.EncryptionSettings.InfrastructureEncryption)))

					if props.SecuritySettings.EncryptionSettings.InfrastructureEncryption != nil {
						state.InfrastructureEncryptionEnabled = strings.EqualFold(string(pointer.From(props.SecuritySettings.EncryptionSettings.InfrastructureEncryption)), string(backupvaults.InfrastructureEncryptionStateEnabled))
					}
				}
			}
			return metadata.Encode(&state)
		},
	}
}

func (r DataProtectionBackupVaultCustomerManagedKeyResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			log.Printf(`[INFO] Customer Managed Keys cannot be removed from Data Protection Backup Vaults once added. To remove the Customer Managed Key, delete and recreate the parent Data Protection Backup Vault`)
			return nil
		},
	}
}

func (r DataProtectionBackupVaultCustomerManagedKeyResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DataProtection.BackupVaultClient
			var cmk DataProtectionBackupVaultCustomerManagedKeyModel

			if err := metadata.Decode(&cmk); err != nil {
				return err
			}
			fmt.Println("debug18 ", cmk.InfrastructureEncryptionEnabled)
			id, err := backupvaults.ParseBackupVaultID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}
			fmt.Println("debug19 ", cmk.InfrastructureEncryptionEnabled)
			locks.ByID(id.ID())
			defer locks.UnlockByID(id.ID())
			fmt.Println("debug20 ", cmk.InfrastructureEncryptionEnabled)
			resp, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("%s was not found", *id)
				}
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}
			fmt.Println("debug21 ", cmk.InfrastructureEncryptionEnabled)
			if resp.Model == nil {
				return fmt.Errorf("retrieving %s: `model` is nil", *id)
			}
			fmt.Println("debug22 ", cmk.InfrastructureEncryptionEnabled)
			if resp.Model.Properties.SecuritySettings == nil || resp.Model.Properties.SecuritySettings.EncryptionSettings == nil {
				return fmt.Errorf("retrieving %s: Customer Managed Key was not found", *id)
			}
			fmt.Println("debug23 ", cmk.InfrastructureEncryptionEnabled)
			payload := resp.Model
			fmt.Println("debug24 ", cmk.InfrastructureEncryptionEnabled)
			if metadata.ResourceData.HasChange("key_vault_key_id") {
				keyId, err := keyVaultParse.ParseOptionallyVersionedNestedItemID(cmk.KeyVaultKeyID)
				if err != nil {
					return err
				}
				payload.Properties.SecuritySettings.EncryptionSettings.KeyVaultProperties = &backupvaults.CmkKeyVaultProperties{
					KeyUri: pointer.To(keyId.ID()),
				}
			}
			fmt.Println("debug24 ", cmk.InfrastructureEncryptionEnabled)
			if metadata.ResourceData.HasChange("infrastructure_encryption_enabled") {
				fmt.Println("debug25 ", cmk.InfrastructureEncryptionEnabled)
				if cmk.InfrastructureEncryptionEnabled {
					fmt.Println("debug26")
					payload.Properties.SecuritySettings.EncryptionSettings.InfrastructureEncryption = pointer.To(backupvaults.InfrastructureEncryptionStateEnabled)
				} else {
					fmt.Println("debug27")
					payload.Properties.SecuritySettings.EncryptionSettings.InfrastructureEncryption = pointer.To(backupvaults.InfrastructureEncryptionStateDisabled)
				}
			}

			fmt.Println("debug28, ", pointer.From(payload.Properties.SecuritySettings.EncryptionSettings.InfrastructureEncryption))

			err = client.CreateOrUpdateThenPoll(ctx, *id, *payload, backupvaults.DefaultCreateOrUpdateOperationOptions())
			if err != nil {
				return fmt.Errorf("updating Customer Managed Key for %s: %+v", *id, err)
			}

			return nil
		},
	}
}

/*
func (r DataProtectionBackupVaultCustomerManagedKeyResource) CustomizeDiff() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			if metadata.ResourceDiff.HasChange("infrastructure_encryption_enabled") {

			}

			return nil
		},
	}
}
*/
