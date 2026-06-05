// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package desktopvirtualization_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

func TestAccVirtualDesktopAppAttachPackage_vf_name_emojiSpecialChar(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.vfNameEmojiSpecialChar(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopAppAttachPackage_vf_name_maxLength(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.vfNameMaxLength(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopAppAttachPackage_vf_image_uri_emojiSpecialChar(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.vfImageUriEmojiSpecialChar(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopAppAttachPackage_vf_package_full_name_emojiSpecialChar(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.vfPackageFullNameEmojiSpecialChar(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopAppAttachPackage_vf_display_name_emojiSpecialChar(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.vfDisplayNameEmojiSpecialChar(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopAppAttachPackage_vf_display_name_maxLength(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.vfDisplayNameMaxLength(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r VirtualDesktopAppAttachPackageResource) vfDisplayNameMaxLength(data acceptance.TestData) string {

	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_ids = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name          = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  storage_share_file_id = azurerm_storage_share_file.test6.id
  msix_package_name     = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopAppAttachPackageResource) vfDisplayNameEmojiSpecialChar(data acceptance.TestData) string {

	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_ids = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name          = "🙂\\/\"[]:|<>+=;,?*@&"
  storage_share_file_id = azurerm_storage_share_file.test6.id
  msix_package_name     = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopAppAttachPackageResource) vfImageUriEmojiSpecialChar(data acceptance.TestData) string {

	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name      = "XmlNotepad"
  image_uri         = "🙂\\/\"[]:|<>+=;,?*@&"
  package_full_name = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopAppAttachPackageResource) vfNameEmojiSpecialChar(data acceptance.TestData) string {

	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "🙂\\/\"[]:|<>+=;,?*@&"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_ids = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name          = "XmlNotepad"
  storage_share_file_id = azurerm_storage_share_file.test6.id
  msix_package_name     = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data))
}

func (r VirtualDesktopAppAttachPackageResource) vfNameMaxLength(data acceptance.TestData) string {

	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_ids = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name          = "XmlNotepad"
  storage_share_file_id = azurerm_storage_share_file.test6.id
  msix_package_name     = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data))
}

func (r VirtualDesktopAppAttachPackageResource) vfPackageFullNameEmojiSpecialChar(data acceptance.TestData) string {

	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name      = "XmlNotepad"
  image_uri         = azurerm_storage_share_file.test6.id
  package_full_name = "🙂\\/\"[]:|<>+=;,?*@&"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}
