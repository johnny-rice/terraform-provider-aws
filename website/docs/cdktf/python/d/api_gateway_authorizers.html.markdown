---
subcategory: "API Gateway"
layout: "aws"
page_title: "AWS: aws_api_gateway_authorizers"
description: |-
  Provides details about multiple API Gateway Authorizers.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_api_gateway_authorizers

Provides details about multiple API Gateway Authorizers.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_api_gateway_authorizers import DataAwsApiGatewayAuthorizers
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsApiGatewayAuthorizers(self, "example",
            rest_api_id=Token.as_string(aws_api_gateway_rest_api_example.id)
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `rest_api_id` - (Required) ID of the associated REST API.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `ids` - List of Authorizer identifiers.

<!-- cache-key: cdktf-0.20.8 input-06923b48b8049f6fb82863075767bf1841a89fa4c09f8a0c9b0255d2668ef551 -->