---
subcategory: "Managed Grafana"
layout: "aws"
page_title: "AWS: aws_grafana_license_association"
description: |-
  Provides an Amazon Managed Grafana workspace license association resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_grafana_license_association

Provides an Amazon Managed Grafana workspace license association resource.

## Example Usage

### Basic configuration

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.grafana_license_association import GrafanaLicenseAssociation
from imports.aws.grafana_workspace import GrafanaWorkspace
from imports.aws.iam_role import IamRole
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        assume = IamRole(self, "assume",
            assume_role_policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "grafana.amazonaws.com"
                        },
                        "Sid": ""
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            name="grafana-assume"
        )
        example = GrafanaWorkspace(self, "example",
            account_access_type="CURRENT_ACCOUNT",
            authentication_providers=["SAML"],
            permission_type="SERVICE_MANAGED",
            role_arn=assume.arn
        )
        aws_grafana_license_association_example = GrafanaLicenseAssociation(self, "example_2",
            license_type="ENTERPRISE_FREE_TRIAL",
            workspace_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_grafana_license_association_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `grafana_token` - (Optional) A token from Grafana Labs that ties your AWS account with a Grafana Labs account.
* `license_type` - (Required) The type of license for the workspace license association. Valid values are `ENTERPRISE` and `ENTERPRISE_FREE_TRIAL`.
* `workspace_id` - (Required) The workspace id.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `free_trial_expiration` - If `license_type` is set to `ENTERPRISE_FREE_TRIAL`, this is the expiration date of the free trial.
* `license_expiration` - If `license_type` is set to `ENTERPRISE`, this is the expiration date of the enterprise license.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Grafana workspace license association using the workspace's `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.grafana_license_association import GrafanaLicenseAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GrafanaLicenseAssociation.generate_config_for_import(self, "example", "g-2054c75a02")
```

Using `terraform import`, import Grafana workspace license association using the workspace's `id`. For example:

```console
% terraform import aws_grafana_license_association.example g-2054c75a02
```

<!-- cache-key: cdktf-0.20.8 input-e849564597fa773fcafee168404f7ff43d95505cc5a173a11c3fe5c43191f71e -->