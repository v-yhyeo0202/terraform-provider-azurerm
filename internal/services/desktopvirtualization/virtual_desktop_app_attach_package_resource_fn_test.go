// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package desktopvirtualization_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

func TestAccVirtualDesktopAppAttachPackage_fn_host_pool_references(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.fnHostPoolReferencesFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnHostPoolReferencesSecondValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnHostPoolReferencesFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnHostPoolReferencesBothValues(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnHostPoolReferencesFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopAppAttachPackage_fn_fail_health_check_on_staging_failure(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.fnFailHealthCheckOnStagingFailureNotSet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnFailHealthCheckOnStagingFailureSet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnFailHealthCheckOnStagingFailureNotSet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopAppAttachPackage_fn_image_uri(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.fnImageUriFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnImageUriSecondValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnImageUriFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopAppAttachPackage_fn_package_full_name(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.fnPackageFullNameFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnPackageFullNameSecondValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnPackageFullNameFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccVirtualDesktopAppAttachPackage_fn_display_name(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.fnDisplayNameFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnDisplayNameSecondValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnDisplayNameFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r VirtualDesktopAppAttachPackageResource) fnDisplayName(data acceptance.TestData, displayName string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name      = "%[3]s"
  image_uri         = azurerm_storage_share_file.test6.id
  package_full_name = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger, displayName)
}

func (r VirtualDesktopAppAttachPackageResource) fnDisplayNameFirstValue(data acceptance.TestData) string {
	return r.fnDisplayName(data, "XmlNotepad")
}

func (r VirtualDesktopAppAttachPackageResource) fnDisplayNameSecondValue(data acceptance.TestData) string {
	return r.fnDisplayName(data, "XmlNotepadUpdated")
}

func TestAccVirtualDesktopAppAttachPackage_fn_is_regular_registration(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.fnIsRegularRegistrationNotSet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnIsRegularRegistrationSet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnIsRegularRegistrationNotSet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r VirtualDesktopAppAttachPackageResource) fnIsRegularRegistrationNotSet(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name                         = "XmlNotepad"
  image_uri                            = azurerm_storage_share_file.test6.id
  package_full_name                    = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"
  fail_health_check_on_staging_failure = "DoNotFail"
  is_active                            = true

  tags = {
    Environment = "Production"
    Foo         = "Bar"
  }

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopAppAttachPackageResource) fnIsRegularRegistrationSet(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name                         = "XmlNotepad"
  image_uri                            = azurerm_storage_share_file.test6.id
  package_full_name                    = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"
  fail_health_check_on_staging_failure = "DoNotFail"
  is_regular_registration              = false
  is_active                            = true

  tags = {
    Environment = "Production"
    Foo         = "Bar"
  }

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func TestAccVirtualDesktopAppAttachPackage_fn_is_active(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.fnIsActiveNotSet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnIsActiveSet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnIsActiveNotSet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r VirtualDesktopAppAttachPackageResource) fnIsActiveNotSet(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name                         = "XmlNotepad"
  image_uri                            = azurerm_storage_share_file.test6.id
  package_full_name                    = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"
  fail_health_check_on_staging_failure = "DoNotFail"
  is_regular_registration              = false

  tags = {
    Environment = "Production"
    Foo         = "Bar"
  }

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopAppAttachPackageResource) fnIsActiveSet(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name                         = "XmlNotepad"
  image_uri                            = azurerm_storage_share_file.test6.id
  package_full_name                    = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"
  fail_health_check_on_staging_failure = "DoNotFail"
  is_regular_registration              = false
  is_active                            = true

  tags = {
    Environment = "Production"
    Foo         = "Bar"
  }

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopAppAttachPackageResource) fnFailHealthCheckOnStagingFailureNotSet(data acceptance.TestData) string {
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
  package_full_name = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopAppAttachPackageResource) fnFailHealthCheckOnStagingFailureSet(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name                         = "XmlNotepad"
  image_uri                            = azurerm_storage_share_file.test6.id
  package_full_name                    = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"
  fail_health_check_on_staging_failure = "DoNotFail"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopAppAttachPackageResource) fnPackageFullName(data acceptance.TestData, packageFullName string) string {
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
  package_full_name = "%[3]s"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger, packageFullName)
}

func (r VirtualDesktopAppAttachPackageResource) fnPackageFullNameFirstValue(data acceptance.TestData) string {
	return r.fnPackageFullName(data, "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj")
}

func (r VirtualDesktopAppAttachPackageResource) fnPackageFullNameSecondValue(data acceptance.TestData) string {
	return r.fnPackageFullName(data, "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral_split.scale-100_hndwmj480pefj")
}

func (r VirtualDesktopAppAttachPackageResource) fnTemplate(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_host_pool" "test2" {
  name                = "acctest-vdpool2-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  type                = "Pooled"
  load_balancer_type  = "BreadthFirst"
}
`, r.template(data), data.RandomInteger)
}

func (r VirtualDesktopAppAttachPackageResource) fnHostPoolReferences(data acceptance.TestData, hostPoolReferences string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    %[3]s
  ]

  display_name      = "XmlNotepad"
  image_uri         = azurerm_storage_share_file.test6.id
  package_full_name = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.fnTemplate(data), data.RandomInteger, hostPoolReferences)
}

func (r VirtualDesktopAppAttachPackageResource) fnHostPoolReferencesFirstValue(data acceptance.TestData) string {
	return r.fnHostPoolReferences(data, `azurerm_virtual_desktop_host_pool.test.id`)
}

func (r VirtualDesktopAppAttachPackageResource) fnHostPoolReferencesSecondValue(data acceptance.TestData) string {
	return r.fnHostPoolReferences(data, `azurerm_virtual_desktop_host_pool.test2.id`)
}

func (r VirtualDesktopAppAttachPackageResource) fnHostPoolReferencesBothValues(data acceptance.TestData) string {
	return r.fnHostPoolReferences(data, `azurerm_virtual_desktop_host_pool.test.id,
    azurerm_virtual_desktop_host_pool.test2.id`)
}

func (r VirtualDesktopAppAttachPackageResource) fnImageUriTemplate(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_storage_share_file" "test_image_uri" {
  name              = "xmlNotepad2.cim"
  storage_share_url = azurerm_storage_share.test.url
  source            = "${path.module}/testdata/xmlNotepad.cim"
}
`, r.template(data))
}

func (r VirtualDesktopAppAttachPackageResource) fnImageUri(data acceptance.TestData, imageUri string) string {
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
  image_uri         = %[3]s
  package_full_name = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.fnImageUriTemplate(data), data.RandomInteger, imageUri)
}

func (r VirtualDesktopAppAttachPackageResource) fnImageUriFirstValue(data acceptance.TestData) string {
	return r.fnImageUri(data, `azurerm_storage_share_file.test6.id`)
}

func (r VirtualDesktopAppAttachPackageResource) fnImageUriSecondValue(data acceptance.TestData) string {
	return r.fnImageUri(data, `azurerm_storage_share_file.test_image_uri.id`)
}

func TestAccVirtualDesktopAppAttachPackage_fn_tags(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_desktop_app_attach_package", "test")
	r := VirtualDesktopAppAttachPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.fnTagsNotSet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnTagsOneElementFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnTagsOneElementSecondValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnTagsOneElementFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnTagsTwoElements(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnTagsOneElementFirstValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fnTagsNotSet(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r VirtualDesktopAppAttachPackageResource) fnTags(data acceptance.TestData, tags string) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_virtual_desktop_app_attach_package" "test" {
  name                = "acctest-msix-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  host_pool_references = [
    azurerm_virtual_desktop_host_pool.test.id
  ]

  display_name                         = "XmlNotepad"
  image_uri                            = azurerm_storage_share_file.test6.id
  package_full_name                    = "43906ChrisLovett.XmlNotepad_2.9.0.21_neutral__hndwmj480pefj"
  fail_health_check_on_staging_failure = "DoNotFail"
  is_regular_registration              = false
  is_active                            = true

%[3]s

  depends_on = [
    azurerm_virtual_machine_extension.test0,
    azurerm_virtual_machine_extension.test1,
    azurerm_virtual_machine_extension.test2
  ]
}
`, r.template(data), data.RandomInteger, tags)
}

func (r VirtualDesktopAppAttachPackageResource) fnTagsNotSet(data acceptance.TestData) string {
	return r.fnTags(data, "")
}

func (r VirtualDesktopAppAttachPackageResource) fnTagsOneElementFirstValue(data acceptance.TestData) string {
	return r.fnTags(data, `  tags = {
    Environment = "Production"
  }`)
}

func (r VirtualDesktopAppAttachPackageResource) fnTagsOneElementSecondValue(data acceptance.TestData) string {
	return r.fnTags(data, `  tags = {
    Environment = "Staging"
  }`)
}

func (r VirtualDesktopAppAttachPackageResource) fnTagsTwoElements(data acceptance.TestData) string {
	return r.fnTags(data, `  tags = {
    Environment = "Production"
    Foo         = "Bar"
  }`)
}
