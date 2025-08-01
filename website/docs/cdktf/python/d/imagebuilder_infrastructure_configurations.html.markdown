---
subcategory: "EC2 Image Builder"
layout: "aws"
page_title: "AWS: aws_imagebuilder_infrastructure_configurations"
description: |-
    Get information on Image Builder Infrastructure Configurations.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_imagebuilder_infrastructure_configurations

Use this data source to get the ARNs and names of Image Builder Infrastructure Configurations matching the specified criteria.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_imagebuilder_infrastructure_configurations import DataAwsImagebuilderInfrastructureConfigurations
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsImagebuilderInfrastructureConfigurations(self, "example",
            filter=[DataAwsImagebuilderInfrastructureConfigurationsFilter(
                name="name",
                values=["example"]
            )
            ]
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `filter` - (Optional) Configuration block(s) for filtering. Detailed below.

## filter Configuration Block

The `filter` configuration block supports the following arguments:

* `name` - (Required) Name of the filter field. Valid values can be found in the [Image Builder ListInfrastructureConfigurations API Reference](https://docs.aws.amazon.com/imagebuilder/latest/APIReference/API_ListInfrastructureConfigurations.html).
* `values` - (Required) Set of values that are accepted for the given filter field. Results will be selected if any given value matches.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arns` - Set of ARNs of the matched Image Builder Infrastructure Configurations.
* `names` - Set of names of the matched Image Builder Infrastructure Configurations.

<!-- cache-key: cdktf-0.20.8 input-a1cd72071b141f397ead5f974367942298426b6da2b2d44014362e54e3ab1408 -->