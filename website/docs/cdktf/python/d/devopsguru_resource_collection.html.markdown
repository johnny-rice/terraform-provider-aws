---
subcategory: "DevOps Guru"
layout: "aws"
page_title: "AWS: aws_devopsguru_resource_collection"
description: |-
  Terraform data source for managing an AWS DevOps Guru Resource Collection.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_devopsguru_resource_collection

Terraform data source for managing an AWS DevOps Guru Resource Collection.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_devopsguru_resource_collection import DataAwsDevopsguruResourceCollection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsDevopsguruResourceCollection(self, "example",
            type="AWS_SERVICE"
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `type` - (Required) Type of AWS resource collection to create. Valid values are `AWS_CLOUD_FORMATION`, `AWS_SERVICE`, and `AWS_TAGS`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `cloudformation` - A collection of AWS CloudFormation stacks. See [`cloudformation`](#cloudformation-attribute-reference) below for additional details.
* `id` - Type of AWS resource collection to create (same value as `type`).
* `tags` - AWS tags used to filter the resources in the resource collection. See [`tags`](#tags-attribute-reference) below for additional details.

### `cloudformation` Attribute Reference

* `stack_names` - Array of the names of the AWS CloudFormation stacks.

### `tags` Attribute Reference

* `app_boundary_key` - An AWS tag key that is used to identify the AWS resources that DevOps Guru analyzes.
* `tag_values` - Array of tag values.

<!-- cache-key: cdktf-0.20.8 input-cb687772fee35e1e103ec9995c0b5bb556649629014d46032a0a2db864bb4ce4 -->