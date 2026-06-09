// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package desktopvirtualization_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

func TestAccVirtualDesktopMsixPackage_vf_display_name_emojiSpecialChar(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_msix_package", "test")
	r := VirtualDesktopMsixPackageResource{}

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

func TestAccVirtualDesktopMsixPackage_vf_host_pool_name_emojiSpecialChar(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_msix_package", "test")
	r := VirtualDesktopMsixPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.vfHostPoolNameEmojiSpecialChar(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopMsixPackage_vf_image_uri_emojiSpecialChar(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_msix_package", "test")
	r := VirtualDesktopMsixPackageResource{}

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

func TestAccVirtualDesktopMsixPackage_vf_name_emojiSpecialChar(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_msix_package", "test")
	r := VirtualDesktopMsixPackageResource{}

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

func TestAccVirtualDesktopMsixPackage_vf_package_full_name_emojiSpecialChar(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_msix_package", "test")
	r := VirtualDesktopMsixPackageResource{}

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

func (r VirtualDesktopMsixPackageResource) vfPackageFullNameEmojiSpecialChar(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_msix_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  display_name        = "XmlNotepad"
  host_pool_name      = azurerm_virtual_desktop_host_pool.test.name
  image_uri           = azurerm_storage_share_file.test6.id
  package_full_name   = "🙂\\/\"[]:|<>+=;,?*@&"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopMsixPackageResource) vfNameEmojiSpecialChar(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_msix_package" "test" {
  name                = "🙂\\/\"[]:|<>+=;,?*@&"
  resource_group_name = azurerm_resource_group.test.name
  display_name        = "XmlNotepad"
  host_pool_name      = azurerm_virtual_desktop_host_pool.test.name
  image_uri           = azurerm_storage_share_file.test6.id
  package_full_name   = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data))
}

func (r VirtualDesktopMsixPackageResource) vfHostPoolNameEmojiSpecialChar(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_msix_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  display_name        = "XmlNotepad"
  host_pool_name      = "🙂\\/\"[]:|<>+=;,?*@&"
  image_uri           = azurerm_storage_share_file.test6.id
  package_full_name   = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopMsixPackageResource) vfDisplayNameEmojiSpecialChar(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_msix_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  display_name        = "🙂\\/\"[]:|<>+=;,?*@&"
  host_pool_name      = azurerm_virtual_desktop_host_pool.test.name
  image_uri           = azurerm_storage_share_file.test6.id
  package_full_name   = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopMsixPackageResource) vfImageUriEmojiSpecialChar(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_msix_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  display_name        = "XmlNotepad"
  host_pool_name      = azurerm_virtual_desktop_host_pool.test.name
  image_uri           = "🙂\\/\"[]:|<>+=;,?*@&"
  package_full_name   = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}
