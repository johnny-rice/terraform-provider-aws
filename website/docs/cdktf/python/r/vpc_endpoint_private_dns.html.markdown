---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_vpc_endpoint_private_dns"
description: |-
  Terraform resource for enabling private DNS on an AWS VPC (Virtual Private Cloud) Endpoint.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpc_endpoint_private_dns

Terraform resource for enabling private DNS on an AWS VPC (Virtual Private Cloud) Endpoint.

~> When using this resource, the `private_dns_enabled` argument should be omitted on the parent `aws_vpc_endpoint` resource.
Setting the value both places can lead to unintended behavior and persistent differences.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.vpc_endpoint_private_dns import VpcEndpointPrivateDns
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        VpcEndpointPrivateDns(self, "example",
            private_dns_enabled=True,
            vpc_endpoint_id=Token.as_string(aws_vpc_endpoint_example.id)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `private_dns_enabled` - (Required) Indicates whether a private hosted zone is associated with the VPC. Only applicable for `Interface` endpoints.
* `vpc_endpoint_id` - (Required) VPC endpoint identifier.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import a VPC (Virtual Private Cloud) Endpoint Private DNS using the `vpc_endpoint_id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.vpc_endpoint_private_dns import VpcEndpointPrivateDns
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        VpcEndpointPrivateDns.generate_config_for_import(self, "example", "vpce-abcd-1234")
```

Using `terraform import`, import a VPC (Virtual Private Cloud) Endpoint Private DNS using the `vpc_endpoint_id`. For example:

```console
% terraform import aws_vpc_endpoint_private_dns.example vpce-abcd-1234
```

<!-- cache-key: cdktf-0.20.8 input-fb61383c27ad6532408d0d99242fa17f2ccc7b0a5214264644e63a609400534d -->