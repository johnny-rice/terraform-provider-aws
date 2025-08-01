// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflow_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfappflow "github.com/hashicorp/terraform-provider-aws/internal/service/appflow"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccAppFlowConnectorProfile_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var connectorProfiles types.ConnectorProfile
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_appflow_connector_profile.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppFlowServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckConnectorProfileDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccConnectorProfileConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConnectorProfileExists(ctx, resourceName, &connectorProfiles),
					acctest.CheckResourceAttrRegionalARNFormat(ctx, resourceName, names.AttrARN, "appflow", "connectorprofile/{name}"),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttrSet(resourceName, "connection_mode"),
					resource.TestCheckResourceAttrSet(resourceName, "connector_profile_config.#"),
					resource.TestCheckResourceAttrSet(resourceName, "connector_profile_config.0.connector_profile_credentials.#"),
					resource.TestCheckResourceAttrSet(resourceName, "connector_profile_config.0.connector_profile_properties.#"),
					resource.TestCheckResourceAttrSet(resourceName, "connector_type"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connector_profile_config.0.connector_profile_credentials"},
			},
		},
	})
}

func TestAccAppFlowConnectorProfile_update(t *testing.T) {
	ctx := acctest.Context(t)
	var connectorProfiles types.ConnectorProfile
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_appflow_connector_profile.test"
	testPrefix := "test-prefix"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppFlowServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckConnectorProfileDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccConnectorProfileConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConnectorProfileExists(ctx, resourceName, &connectorProfiles),
				),
			},
			{
				Config: testAccConnectorProfileConfig_update(rName, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConnectorProfileExists(ctx, resourceName, &connectorProfiles),
					resource.TestCheckResourceAttr(resourceName, "connector_profile_config.0.connector_profile_properties.0.redshift.0.bucket_prefix", testPrefix),
				),
			},
		},
	})
}

func TestAccAppFlowConnectorProfile_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var connectorProfiles types.ConnectorProfile
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_appflow_connector_profile.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppFlowServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckConnectorProfileDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccConnectorProfileConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConnectorProfileExists(ctx, resourceName, &connectorProfiles),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfappflow.ResourceConnectorProfile(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckConnectorProfileDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).AppFlowClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_appflow_connector_profile" {
				continue
			}

			_, err := tfappflow.FindConnectorProfileByName(ctx, conn, rs.Primary.Attributes[names.AttrName])

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("AppFlow Connector Profile %s still exists", rs.Primary.Attributes[names.AttrName])
		}

		return nil
	}
}

func testAccCheckConnectorProfileExists(ctx context.Context, n string, v *types.ConnectorProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).AppFlowClient(ctx)

		output, err := tfappflow.FindConnectorProfileByName(ctx, conn, rs.Primary.Attributes[names.AttrName])

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccConnectorProfileConfig_base(rName string, redshiftPassword string, redshiftUsername string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

data "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id
}

resource "aws_route" "test" {
  route_table_id         = data.aws_route_table.test.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.test.id
}

resource "aws_redshift_subnet_group" "test" {
  name       = %[1]q
  subnet_ids = aws_subnet.test[*].id
}

data "aws_iam_policy" "test" {
  name = "AmazonRedshiftFullAccess"
}

resource "aws_iam_role" "test" {
  name = %[1]q

  managed_policy_arns = [data.aws_iam_policy.test.arn]

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "appflow.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_security_group" "test" {
  name   = %[1]q
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_security_group_rule" "test" {
  type = "ingress"

  security_group_id = aws_security_group.test.id

  from_port   = 0
  to_port     = 65535
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_redshift_cluster" "test" {
  cluster_identifier = %[1]q

  availability_zone         = data.aws_availability_zones.available.names[0]
  cluster_subnet_group_name = aws_redshift_subnet_group.test.name
  vpc_security_group_ids    = [aws_security_group.test.id]

  master_password = %[2]q
  master_username = %[3]q

  publicly_accessible = false

  node_type           = "ra3.large"
  skip_final_snapshot = true
  encrypted           = true
}
`, rName, redshiftPassword, redshiftUsername))
}

func testAccConnectorProfileConfig_basic(rName string) string {
	const redshiftPassword = "testPassword123!"
	const redshiftUsername = "testusername"

	return acctest.ConfigCompose(
		testAccConnectorProfileConfig_base(rName, redshiftPassword, redshiftUsername),
		fmt.Sprintf(`
resource "aws_appflow_connector_profile" "test" {
  name            = %[1]q
  connector_type  = "Redshift"
  connection_mode = "Public"

  connector_profile_config {

    connector_profile_credentials {
      redshift {
        password = aws_redshift_cluster.test.master_password
        username = aws_redshift_cluster.test.master_username
      }
    }

    connector_profile_properties {
      redshift {
        bucket_name        = %[1]q
        cluster_identifier = aws_redshift_cluster.test.cluster_identifier
        database_name      = "dev"
        database_url       = "jdbc:redshift://${aws_redshift_cluster.test.endpoint}/dev"
        data_api_role_arn  = aws_iam_role.test.arn
        role_arn           = aws_iam_role.test.arn
      }
    }
  }

  depends_on = [
    aws_route.test,
    aws_security_group_rule.test,
  ]
}
`, rName, redshiftPassword, redshiftUsername))
}

func testAccConnectorProfileConfig_update(rName string, bucketPrefix string) string {
	const redshiftPassword = "testPassword123!"
	const redshiftUsername = "testusername"

	return acctest.ConfigCompose(
		testAccConnectorProfileConfig_base(rName, redshiftPassword, redshiftUsername),
		fmt.Sprintf(`
resource "aws_appflow_connector_profile" "test" {
  name            = %[1]q
  connector_type  = "Redshift"
  connection_mode = "Public"

  connector_profile_config {

    connector_profile_credentials {
      redshift {
        password = aws_redshift_cluster.test.master_password
        username = aws_redshift_cluster.test.master_username
      }
    }

    connector_profile_properties {
      redshift {
        bucket_name        = %[1]q
        bucket_prefix      = %[4]q
        cluster_identifier = aws_redshift_cluster.test.cluster_identifier
        database_name      = "dev"
        database_url       = "jdbc:redshift://${aws_redshift_cluster.test.endpoint}/dev"
        data_api_role_arn  = aws_iam_role.test.arn
        role_arn           = aws_iam_role.test.arn
      }
    }
  }

  depends_on = [
    aws_route.test,
    aws_security_group_rule.test,
  ]
}
`, rName, redshiftPassword, redshiftUsername, bucketPrefix))
}
