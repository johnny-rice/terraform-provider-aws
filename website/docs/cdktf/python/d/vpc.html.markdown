---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_vpc"
description: |-
    Provides details about a specific VPC
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_vpc

`aws_vpc` provides details about a specific VPC.

This resource can prove useful when a module accepts a vpc id as
an input variable and needs to, for example, determine the CIDR block of that
VPC.

## Example Usage

The following example shows how one might accept a VPC id as a variable
and use this data source to obtain the data necessary to create a subnet
within it.

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformVariable, Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_vpc import DataAwsVpc
from imports.aws.subnet import Subnet
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        # Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
        #     You can read more about this at https://cdk.tf/variables
        vpc_id = TerraformVariable(self, "vpc_id")
        selected = DataAwsVpc(self, "selected",
            id=vpc_id.string_value
        )
        Subnet(self, "example",
            availability_zone="us-west-2a",
            cidr_block=Token.as_string(
                Fn.cidrsubnet(Token.as_string(selected.cidr_block), 4, 1)),
            vpc_id=Token.as_string(selected.id)
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `cidr_block` - (Optional) Cidr block of the desired VPC.
* `dhcp_options_id` - (Optional) DHCP options id of the desired VPC.
* `default` - (Optional) Boolean constraint on whether the desired VPC is
  the default VPC for the region.
* `filter` - (Optional) Custom filter block as described below.
* `id` - (Optional) ID of the specific VPC to retrieve.
* `state` - (Optional) Current state of the desired VPC.
  Can be either `"pending"` or `"available"`.
* `tags` - (Optional) Map of tags, each pair of which must exactly match
  a pair on the desired VPC.

More complex filters can be expressed using one or more `filter` sub-blocks,
which take the following arguments:

* `name` - (Required) Name of the field to filter by, as defined by
  [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html).
* `values` - (Required) Set of values that are accepted for the given field.
  A VPC will be selected if any one of the given values matches.

## Attribute Reference

All of the argument attributes except `filter` blocks are also exported as
result attributes. This data source will complete the data by populating
any fields that are not included in the configuration with the data for
the selected VPC.

The following attribute is additionally exported:

* `arn` - ARN of VPC
* `enable_dns_support` - Whether or not the VPC has DNS support
* `enable_network_address_usage_metrics` - Whether Network Address Usage metrics are enabled for your VPC
* `enable_dns_hostnames` - Whether or not the VPC has DNS hostname support
* `instance_tenancy` - Allowed tenancy of instances launched into the
  selected VPC. May be any of `"default"`, `"dedicated"`, or `"host"`.
* `ipv6_association_id` - Association ID for the IPv6 CIDR block.
* `ipv6_cidr_block` - IPv6 CIDR block.
* `main_route_table_id` - ID of the main route table associated with this VPC.
* `owner_id` - ID of the AWS account that owns the VPC.

`cidr_block_associations` is also exported with the following attributes:

* `association_id` - Association ID for the IPv4 CIDR block.
* `cidr_block` - CIDR block for the association.
* `state` - State of the association.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `read` - (Default `20m`)

<!-- cache-key: cdktf-0.20.8 input-a149abcdf6c93ceaf1a84563507c805586fc344d80b2fd597e9a8f4887a29342 -->