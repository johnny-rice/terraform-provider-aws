---
subcategory: "Outposts (EC2)"
layout: "aws"
page_title: "AWS: aws_ec2_local_gateway"
description: |-
    Provides details about an EC2 Local Gateway
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ec2_local_gateway

Provides details about an EC2 Local Gateway.

## Example Usage

The following example shows how one might accept a local gateway id as a variable.

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformVariable, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_ec2_local_gateway import DataAwsEc2LocalGateway
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        # Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
        #     You can read more about this at https://cdk.tf/variables
        local_gateway_id = TerraformVariable(self, "local_gateway_id")
        DataAwsEc2LocalGateway(self, "selected",
            id=local_gateway_id.string_value
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `filter` - (Optional) Custom filter block as described below.
* `id` - (Optional) Id of the specific Local Gateway to retrieve.
* `state` - (Optional) Current state of the desired Local Gateway.
  Can be either `"pending"` or `"available"`.
* `tags` - (Optional) Mapping of tags, each pair of which must exactly match
  a pair on the desired Local Gateway.

The arguments of this data source act as filters for querying the available
Local Gateways in the current region. The given filters must match exactly one
Local Gateway whose data will be exported as attributes.

### `filter`

More complex filters can be expressed using one or more `filter` sub-blocks, which take the following arguments:

* `name` - (Required) Name of the field to filter by, as defined by
  [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLocalGateways.html).
* `values` - (Required) Set of values that are accepted for the given field.
  A Local Gateway will be selected if any one of the given values matches.

## Attribute Reference

All of the argument attributes except `filter` blocks are also exported as
result attributes. This data source will complete the data by populating
any fields that are not included in the configuration with the data for
the selected Local Gateway.

The following attributes are additionally exported:

* `outpost_arn` - ARN of Outpost
* `owner_id` - AWS account identifier that owns the Local Gateway.
* `state` - State of the local gateway.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `read` - (Default `20m`)

<!-- cache-key: cdktf-0.20.8 input-2ffa068d73e670d8d8ed36270fe6a9439e64bb15d1979e39b31372245469fa7e -->