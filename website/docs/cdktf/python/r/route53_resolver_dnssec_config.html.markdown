---
subcategory: "Route 53 Resolver"
layout: "aws"
page_title: "AWS: aws_route53_resolver_dnssec_config"
description: |-
  Provides a Route 53 Resolver DNSSEC config resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_route53_resolver_dnssec_config

Provides a Route 53 Resolver DNSSEC config resource.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.route53_resolver_dnssec_config import Route53ResolverDnssecConfig
from imports.aws.vpc import Vpc
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = Vpc(self, "example",
            cidr_block="10.0.0.0/16",
            enable_dns_hostnames=True,
            enable_dns_support=True
        )
        aws_route53_resolver_dnssec_config_example =
        Route53ResolverDnssecConfig(self, "example_1",
            resource_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_route53_resolver_dnssec_config_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `resource_id` - (Required) The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The ARN for a configuration for DNSSEC validation.
* `id` - The ID for a configuration for DNSSEC validation.
* `owner_id` - The owner account ID of the virtual private cloud (VPC) for a configuration for DNSSEC validation.
* `validation_status` - The validation status for a DNSSEC configuration. The status can be one of the following: `ENABLING`, `ENABLED`, `DISABLING` and `DISABLED`.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import  Route 53 Resolver DNSSEC configs using the Route 53 Resolver DNSSEC config ID. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.route53_resolver_dnssec_config import Route53ResolverDnssecConfig
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Route53ResolverDnssecConfig.generate_config_for_import(self, "example", "rdsc-be1866ecc1683e95")
```

Using `terraform import`, import  Route 53 Resolver DNSSEC configs using the Route 53 Resolver DNSSEC config ID. For example:

```console
% terraform import aws_route53_resolver_dnssec_config.example rdsc-be1866ecc1683e95
```

<!-- cache-key: cdktf-0.20.8 input-25355a87d3de9c2a770f20c93030966e085ecfc0cf8a7652e10a4089dab21b5d -->