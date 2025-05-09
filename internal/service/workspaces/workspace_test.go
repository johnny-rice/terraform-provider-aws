// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaces_test

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfworkspaces "github.com/hashicorp/terraform-provider-aws/internal/service/workspaces"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func testAccWorkspace_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.Workspace
	rName := sdkacctest.RandString(8)
	domain := acctest.RandomDomainName()

	resourceName := "aws_workspaces_workspace.test"
	directoryResourceName := "aws_workspaces_directory.test"
	bundleDataSourceName := "data.aws_workspaces_bundle.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheckDirectory(ctx, t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(ctx, t)
			acctest.PreCheckHasIAMRole(ctx, t, "workspaces_DefaultRole")
		},
		ErrorCheck:               acctest.ErrorCheck(t, strings.ToLower(workspaces.ServiceID)),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWorkspaceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Destroy: false,
				Config:  testAccWorkspaceConfig_basic(rName, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "directory_id", directoryResourceName, names.AttrID),
					resource.TestCheckResourceAttrPair(resourceName, "bundle_id", bundleDataSourceName, names.AttrID),
					resource.TestMatchResourceAttr(resourceName, names.AttrIPAddress, regexache.MustCompile(`\d+\.\d+\.\d+\.\d+`)),
					resource.TestCheckResourceAttr(resourceName, names.AttrState, string(types.WorkspaceStateAvailable)),
					resource.TestCheckResourceAttr(resourceName, "root_volume_encryption_enabled", acctest.CtFalse),
					resource.TestCheckResourceAttr(resourceName, names.AttrUserName, "Administrator"),
					resource.TestCheckResourceAttr(resourceName, "volume_encryption_key", ""),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.compute_type_name", string(types.ComputeValue)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.root_volume_size_gib", "80"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode", string(types.RunningModeAlwaysOn)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode_auto_stop_timeout_in_minutes", "0"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.user_volume_size_gib", "10"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", fmt.Sprintf("tf-testacc-workspaces-workspace-%[1]s", rName)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccWorkspace_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v1, v2, v3 types.Workspace
	rName := sdkacctest.RandString(8)
	domain := acctest.RandomDomainName()

	resourceName := "aws_workspaces_workspace.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheckDirectory(ctx, t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(ctx, t)
			acctest.PreCheckHasIAMRole(ctx, t, "workspaces_DefaultRole")
		},
		ErrorCheck:               acctest.ErrorCheck(t, strings.ToLower(workspaces.ServiceID)),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWorkspaceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkspaceConfig_tagsA(rName, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", fmt.Sprintf("tf-testacc-workspaces-workspace-%[1]s", rName)),
					resource.TestCheckResourceAttr(resourceName, "tags.Alpha", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccWorkspaceConfig_tagsB(rName, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", fmt.Sprintf("tf-testacc-workspaces-workspace-%[1]s", rName)),
					resource.TestCheckResourceAttr(resourceName, "tags.Beta", "2"),
				),
			},
			{
				Config: testAccWorkspaceConfig_tagsC(rName, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v3),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", fmt.Sprintf("tf-testacc-workspaces-workspace-%[1]s", rName)),
				),
			},
		},
	})
}

func testAccWorkspace_workspaceProperties(t *testing.T) {
	ctx := acctest.Context(t)
	var v1, v2, v3 types.Workspace
	rName := sdkacctest.RandString(8)
	domain := acctest.RandomDomainName()

	resourceName := "aws_workspaces_workspace.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheckDirectory(ctx, t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(ctx, t)
			acctest.PreCheckHasIAMRole(ctx, t, "workspaces_DefaultRole")
		},
		ErrorCheck:               acctest.ErrorCheck(t, strings.ToLower(workspaces.ServiceID)),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWorkspaceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Destroy: false,
				Config:  testAccWorkspaceConfig_propertiesA(rName, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.compute_type_name", string(types.ComputeValue)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.root_volume_size_gib", "80"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode", string(types.RunningModeAutoStop)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode_auto_stop_timeout_in_minutes", "120"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.user_volume_size_gib", "10"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccWorkspaceConfig_propertiesB(rName, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.compute_type_name", string(types.ComputeValue)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.root_volume_size_gib", "80"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode", string(types.RunningModeAlwaysOn)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode_auto_stop_timeout_in_minutes", "0"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.user_volume_size_gib", "10"),
				),
			},
			{
				Config: testAccWorkspaceConfig_propertiesC(rName, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v3),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.compute_type_name", string(types.ComputeValue)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.root_volume_size_gib", "80"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode", string(types.RunningModeAlwaysOn)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode_auto_stop_timeout_in_minutes", "0"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.user_volume_size_gib", "10"),
				),
			},
		},
	})
}

func testAccWorkspace_workspaceProperties_runningModeAutoStopTimeoutInMinutes(t *testing.T) {
	ctx := acctest.Context(t)
	var v1 types.Workspace
	rName := sdkacctest.RandString(8)
	domain := acctest.RandomDomainName()

	resourceName := "aws_workspaces_workspace.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheckDirectory(ctx, t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(ctx, t)
			acctest.PreCheckHasIAMRole(ctx, t, "workspaces_DefaultRole")
		},
		ErrorCheck:               acctest.ErrorCheck(t, strings.ToLower(workspaces.ServiceID)),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWorkspaceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Destroy: false,
				Config:  testAccWorkspaceConfig_propertiesAutoStopTimeoutInMinutes(rName, domain, 120),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.compute_type_name", string(types.ComputeValue)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.root_volume_size_gib", "80"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode", string(types.RunningModeAutoStop)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode_auto_stop_timeout_in_minutes", "120"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.user_volume_size_gib", "10"),
				),
			},
			{
				Config: testAccWorkspaceConfig_propertiesAutoStopTimeoutInMinutes(rName, domain, 180),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.compute_type_name", string(types.ComputeValue)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.root_volume_size_gib", "80"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode", string(types.RunningModeAutoStop)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode_auto_stop_timeout_in_minutes", "180"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.user_volume_size_gib", "10"),
				),
			},
		},
	})
}

// testAccWorkspace_workspaceProperties_runningModeAlwaysOn
// validates workspace resource creation/import when workspace_properties.running_mode is set to ALWAYS_ON
// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/13558
func testAccWorkspace_workspaceProperties_runningModeAlwaysOn(t *testing.T) {
	ctx := acctest.Context(t)
	var v1 types.Workspace
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_workspaces_workspace.test"
	domain := acctest.RandomDomainName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheckDirectory(ctx, t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(ctx, t)
			acctest.PreCheckHasIAMRole(ctx, t, "workspaces_DefaultRole")
		},
		ErrorCheck:               acctest.ErrorCheck(t, strings.ToLower(workspaces.ServiceID)),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWorkspaceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkspaceConfig_propertiesB(rName, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.compute_type_name", string(types.ComputeValue)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.root_volume_size_gib", "80"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode", string(types.RunningModeAlwaysOn)),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.running_mode_auto_stop_timeout_in_minutes", "0"),
					resource.TestCheckResourceAttr(resourceName, "workspace_properties.0.user_volume_size_gib", "10"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccWorkspace_validateRootVolumeSize(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandString(8)
	domain := acctest.RandomDomainName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, strings.ToLower(workspaces.ServiceID)),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWorkspaceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config:      testAccWorkspaceConfig_validateRootVolumeSize(rName, domain),
				ExpectError: regexache.MustCompile(regexp.QuoteMeta("expected workspace_properties.0.root_volume_size_gib to be one of [80], got 90")),
			},
		},
	})
}

func testAccWorkspace_validateUserVolumeSize(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandString(8)
	domain := acctest.RandomDomainName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, strings.ToLower(workspaces.ServiceID)),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWorkspaceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config:      testAccWorkspaceConfig_validateUserVolumeSize(rName, domain),
				ExpectError: regexache.MustCompile(regexp.QuoteMeta("workspace_properties.0.user_volume_size_gib to be one of [10 50], got 60")),
			},
		},
	})
}

func testAccWorkspace_recreate(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.Workspace
	rName := sdkacctest.RandString(8)
	domain := acctest.RandomDomainName()

	resourceName := "aws_workspaces_workspace.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheckDirectory(ctx, t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(ctx, t)
			acctest.PreCheckHasIAMRole(ctx, t, "workspaces_DefaultRole")
		},
		ErrorCheck:               acctest.ErrorCheck(t, strings.ToLower(workspaces.ServiceID)),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWorkspaceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkspaceConfig_basic(rName, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v),
				),
			},
			{
				Taint:  []string{resourceName}, // Force workspace re-creation
				Config: testAccWorkspaceConfig_basic(rName, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v),
				),
			},
		},
	})
}

func testAccWorkspace_timeout(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.Workspace
	rName := sdkacctest.RandString(8)
	domain := acctest.RandomDomainName()

	resourceName := "aws_workspaces_workspace.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheckDirectory(ctx, t)
			acctest.PreCheckDirectoryServiceSimpleDirectory(ctx, t)
			acctest.PreCheckHasIAMRole(ctx, t, "workspaces_DefaultRole")
		},
		ErrorCheck:               acctest.ErrorCheck(t, strings.ToLower(workspaces.ServiceID)),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWorkspaceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Destroy: false,
				Config:  testAccWorkspaceConfig_timeout(rName, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckWorkspaceExists(ctx, resourceName, &v),
				),
			},
		},
	})
}

func testAccCheckWorkspaceDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).WorkSpacesClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_workspaces_workspace" {
				continue
			}

			_, err := tfworkspaces.FindWorkspaceByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("WorkSpaces Workspace %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckWorkspaceExists(ctx context.Context, n string, v *types.Workspace) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).WorkSpacesClient(ctx)

		output, err := tfworkspaces.FindWorkspaceByID(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccWorkspaceConfig_Prerequisites(rName, domain string) string {
	return acctest.ConfigCompose(
		testAccDirectoryConfig_base(rName, domain),
		fmt.Sprintf(`
data "aws_workspaces_bundle" "test" {
  bundle_id = "wsb-bh8rsxt14" # Value with Windows 10 (English)
}

resource "aws_workspaces_directory" "test" {
  directory_id = aws_directory_service_directory.main.id

  tags = {
    Name = "tf-testacc-workspaces-directory-%[1]s"
  }
}
`, rName))
}

func testAccWorkspaceConfig_basic(rName, domain string) string {
	return acctest.ConfigCompose(
		testAccWorkspaceConfig_Prerequisites(rName, domain),
		fmt.Sprintf(`
resource "aws_workspaces_workspace" "test" {
  bundle_id    = data.aws_workspaces_bundle.test.id
  directory_id = aws_workspaces_directory.test.id

  # NOTE: WorkSpaces API doesn't allow creating users in the directory.
  # However, "Administrator"" user is always present in a bare directory.
  user_name = "Administrator"

  tags = {
    Name = "tf-testacc-workspaces-workspace-%[1]s"
  }
}
`, rName))
}

func testAccWorkspaceConfig_tagsA(rName, domain string) string {
	return acctest.ConfigCompose(
		testAccWorkspaceConfig_Prerequisites(rName, domain),
		fmt.Sprintf(`
resource "aws_workspaces_workspace" "test" {
  bundle_id    = data.aws_workspaces_bundle.test.id
  directory_id = aws_workspaces_directory.test.id

  # NOTE: WorkSpaces API doesn't allow creating users in the directory.
  # However, "Administrator"" user is always present in a bare directory.
  user_name = "Administrator"

  tags = {
    Name  = "tf-testacc-workspaces-workspace-%[1]s"
    Alpha = 1
  }
}
`, rName))
}

func testAccWorkspaceConfig_tagsB(rName, domain string) string {
	return acctest.ConfigCompose(
		testAccWorkspaceConfig_Prerequisites(rName, domain),
		fmt.Sprintf(`
resource "aws_workspaces_workspace" "test" {
  bundle_id    = data.aws_workspaces_bundle.test.id
  directory_id = aws_workspaces_directory.test.id

  # NOTE: WorkSpaces API doesn't allow creating users in the directory.
  # However, "Administrator"" user is always present in a bare directory.
  user_name = "Administrator"

  tags = {
    Name = "tf-testacc-workspaces-workspace-%[1]s"
    Beta = 2
  }
}
`, rName))
}

func testAccWorkspaceConfig_tagsC(rName, domain string) string {
	return acctest.ConfigCompose(
		testAccWorkspaceConfig_Prerequisites(rName, domain),
		fmt.Sprintf(`
resource "aws_workspaces_workspace" "test" {
  bundle_id    = data.aws_workspaces_bundle.test.id
  directory_id = aws_workspaces_directory.test.id

  # NOTE: WorkSpaces API doesn't allow creating users in the directory.
  # However, "Administrator"" user is always present in a bare directory.
  user_name = "Administrator"

  tags = {
    Name = "tf-testacc-workspaces-workspace-%[1]s"
  }
}
`, rName))
}

func testAccWorkspaceConfig_propertiesA(rName, domain string) string {
	return acctest.ConfigCompose(
		testAccWorkspaceConfig_Prerequisites(rName, domain),
		fmt.Sprintf(`
resource "aws_workspaces_workspace" "test" {
  bundle_id    = data.aws_workspaces_bundle.test.id
  directory_id = aws_workspaces_directory.test.id

  # NOTE: WorkSpaces API doesn't allow creating users in the directory.
  # However, "Administrator"" user is always present in a bare directory.
  user_name = "Administrator"

  workspace_properties {
    # NOTE: Compute type and volume size update not allowed within 6 hours after creation.
    running_mode                              = "AUTO_STOP"
    running_mode_auto_stop_timeout_in_minutes = 120
  }

  tags = {
    Name = "tf-testacc-workspaces-workspace-%[1]s"
  }
}
`, rName))
}

func testAccWorkspaceConfig_propertiesB(rName, domain string) string {
	return acctest.ConfigCompose(
		testAccWorkspaceConfig_Prerequisites(rName, domain),
		fmt.Sprintf(`
resource "aws_workspaces_workspace" "test" {
  bundle_id    = data.aws_workspaces_bundle.test.id
  directory_id = aws_workspaces_directory.test.id

  # NOTE: WorkSpaces API doesn't allow creating users in the directory.
  # However, "Administrator"" user is always present in a bare directory.
  user_name = "Administrator"

  workspace_properties {
    # NOTE: Compute type and volume size update not allowed within 6 hours after creation.
    running_mode = "ALWAYS_ON"
  }

  tags = {
    Name = "tf-testacc-workspaces-workspace-%[1]s"
  }
}
`, rName))
}

func testAccWorkspaceConfig_propertiesC(rName, domain string) string {
	return acctest.ConfigCompose(
		testAccWorkspaceConfig_Prerequisites(rName, domain),
		fmt.Sprintf(`
resource "aws_workspaces_workspace" "test" {
  bundle_id    = data.aws_workspaces_bundle.test.id
  directory_id = aws_workspaces_directory.test.id

  # NOTE: WorkSpaces API doesn't allow creating users in the directory.
  # However, "Administrator"" user is always present in a bare directory.
  user_name = "Administrator"

  workspace_properties {
  }

  tags = {
    Name = "tf-testacc-workspaces-workspace-%[1]s"
  }
}
`, rName))
}

func testAccWorkspaceConfig_propertiesAutoStopTimeoutInMinutes(rName, domain string, autoStoptimeoutInMinutes int) string {
	return acctest.ConfigCompose(
		testAccWorkspaceConfig_Prerequisites(rName, domain),
		fmt.Sprintf(`
resource "aws_workspaces_workspace" "test" {
  bundle_id    = data.aws_workspaces_bundle.test.id
  directory_id = aws_workspaces_directory.test.id

  # NOTE: WorkSpaces API doesn't allow creating users in the directory.
  # However, "Administrator"" user is always present in a bare directory.
  user_name = "Administrator"

  workspace_properties {
    # NOTE: Compute type and volume size update not allowed within 6 hours after creation.
    running_mode                              = "AUTO_STOP"
    running_mode_auto_stop_timeout_in_minutes = %[2]d
  }

  tags = {
    Name = "tf-testacc-workspaces-workspace-%[1]s"
  }
}
`, rName, autoStoptimeoutInMinutes))
}

func testAccWorkspaceConfig_validateRootVolumeSize(rName, domain string) string {
	return acctest.ConfigCompose(
		testAccWorkspaceConfig_Prerequisites(rName, domain),
		fmt.Sprintf(`
resource "aws_workspaces_workspace" "test" {
  bundle_id    = data.aws_workspaces_bundle.test.id
  directory_id = aws_workspaces_directory.test.id

  # NOTE: WorkSpaces API doesn't allow creating users in the directory.
  # However, "Administrator"" user is always present in a bare directory.
  user_name = "Administrator"

  workspace_properties {
    root_volume_size_gib = 90
    user_volume_size_gib = 50
  }

  tags = {
    Name = "tf-testacc-workspaces-workspace-%[1]s"
  }
}
`, rName))
}

func testAccWorkspaceConfig_validateUserVolumeSize(rName, domain string) string {
	return acctest.ConfigCompose(
		testAccWorkspaceConfig_Prerequisites(rName, domain),
		fmt.Sprintf(`
resource "aws_workspaces_workspace" "test" {
  bundle_id    = data.aws_workspaces_bundle.test.id
  directory_id = aws_workspaces_directory.test.id

  # NOTE: WorkSpaces API doesn't allow creating users in the directory.
  # However, "Administrator"" user is always present in a bare directory.
  user_name = "Administrator"

  workspace_properties {
    root_volume_size_gib = 80
    user_volume_size_gib = 60
  }

  tags = {
    Name = "tf-testacc-workspaces-workspace-%[1]s"
  }
}
`, rName))
}

func testAccWorkspaceConfig_timeout(rName, domain string) string {
	return acctest.ConfigCompose(
		testAccWorkspaceConfig_Prerequisites(rName, domain),
		fmt.Sprintf(`
resource "aws_workspaces_workspace" "test" {
  bundle_id    = data.aws_workspaces_bundle.test.id
  directory_id = aws_workspaces_directory.test.id

  # NOTE: WorkSpaces API doesn't allow creating users in the directory.
  # However, "Administrator"" user is always present in a bare directory.
  user_name = "Administrator"

  timeouts {
    create = "60m"
    update = "30m"
    delete = "30m"
  }

  tags = {
    Name = "tf-testacc-workspaces-workspace-%[1]s"
  }
}
`, rName))
}
