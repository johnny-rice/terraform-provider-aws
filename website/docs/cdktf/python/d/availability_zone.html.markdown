---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "AWS: aws_availability_zone"
description: |-
    Provides details about a specific availability zone
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_availability_zone

`aws_availability_zone` provides details about a specific availability zone (AZ)
in the current Region.

This can be used both to validate an availability zone given in a variable
and to split the AZ name into its component parts of an AWS Region and an
AZ identifier letter. The latter may be useful e.g., for implementing a
consistent subnet numbering scheme across several regions by mapping both
the region and the subnet letter to network numbers.

This is different from the `aws_availability_zones` (plural) data source,
which provides a list of the available zones.

## Example Usage

The following example shows how this data source might be used to derive
VPC and subnet CIDR prefixes systematically for an availability zone.

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformVariable, Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_availability_zone import DataAwsAvailabilityZone
from imports.aws.subnet import Subnet
from imports.aws.vpc import Vpc
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        # Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
        #     You can read more about this at https://cdk.tf/variables
        az_number = TerraformVariable(self, "az_number",
            default=[{
                "a": 1,
                "b": 2,
                "c": 3,
                "d": 4,
                "e": 5,
                "f": 6
            }
            ]
        )
        region_number = TerraformVariable(self, "region_number",
            default=[{
                "ap-northeast-1": 5,
                "eu-central-1": 4,
                "us-east-1": 1,
                "us-west-1": 2,
                "us-west-2": 3
            }
            ]
        )
        example = DataAwsAvailabilityZone(self, "example",
            name="eu-central-1a"
        )
        aws_vpc_example = Vpc(self, "example_3",
            cidr_block=Token.as_string(
                Fn.cidrsubnet("10.0.0.0/8", 4,
                    Token.as_number(Fn.lookup_nested(region_number.value, [example.region]))))
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_vpc_example.override_logical_id("example")
        aws_subnet_example = Subnet(self, "example_4",
            cidr_block=Token.as_string(
                Fn.cidrsubnet(
                    Token.as_string(aws_vpc_example.cidr_block), 4,
                    Token.as_number(Fn.lookup_nested(az_number.value, [example.name_suffix])))),
            vpc_id=Token.as_string(aws_vpc_example.id)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_subnet_example.override_logical_id("example")
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `all_availability_zones` - (Optional) Set to `true` to include all Availability Zones and Local Zones regardless of your opt in status.
* `filter` - (Optional) Configuration block(s) for filtering. Detailed below.
* `name` - (Optional) Full name of the availability zone to select.
* `state` - (Optional) Specific availability zone state to require. May be any of `"available"`, `"information"` or `"impaired"`.
* `zone_id` - (Optional) Zone ID of the availability zone to select.

The arguments of this data source act as filters for querying the available
availability zones. The given filters must match exactly one availability
zone whose data will be exported as attributes.

### filter Configuration Block

The `filter` configuration block supports the following arguments:

* `name` - (Required) Name of the filter field. Valid values can be found in the [EC2 DescribeAvailabilityZones API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html).
* `values` - (Required) Set of values that are accepted for the given filter field. Results will be selected if any given value matches.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `group_long_name` - The long name of the Availability Zone group, Local Zone group, or Wavelength Zone group.
* `group_name` - The name of the zone group. For example: `us-east-1-zg-1`, `us-west-2-lax-1`, or `us-east-1-wl1-bos-wlz-1`.
* `name_suffix` - Part of the AZ name that appears after the region name, uniquely identifying the AZ within its region.
For Availability Zones this is usually a single letter, for example `a` for the `us-west-2a` zone.
For Local and Wavelength Zones this is a longer string, for example `wl1-sfo-wlz-1` for the `us-west-2-wl1-sfo-wlz-1` zone.
* `network_border_group` - The name of the location from which the address is advertised.
* `opt_in_status` - For Availability Zones, this always has the value of `opt-in-not-required`. For Local Zones, this is the opt in status. The possible values are `opted-in` and `not-opted-in`.
* `parent_zone_id` - ID of the zone that handles some of the Local Zone or Wavelength Zone control plane operations, such as API calls.
* `parent_zone_name` - Name of the zone that handles some of the Local Zone or Wavelength Zone control plane operations, such as API calls.
* `zone_type` - Type of zone. Values are `availability-zone`, `local-zone`, and `wavelength-zone`.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `read` - (Default `20m`)

<!-- cache-key: cdktf-0.20.8 input-6b07dd8dc1751fdfa8957fbe337cde58e49d6979232e629aa58a1e04861a38e8 -->