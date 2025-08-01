---
subcategory: "Bedrock"
layout: "aws"
page_title: "AWS: aws_bedrock_provisioned_model_throughput"
description: |-
  Manages Provisioned Throughput for an Amazon Bedrock model.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_bedrock_provisioned_model_throughput

Manages [Provisioned Throughput](https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html) for an Amazon Bedrock model.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.bedrock_provisioned_model_throughput import BedrockProvisionedModelThroughput
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BedrockProvisionedModelThroughput(self, "example",
            commitment_duration="SixMonths",
            model_arn="arn:aws:bedrock:us-east-1::foundation-model/anthropic.claude-v2",
            model_units=1,
            provisioned_model_name="example-model"
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `commitment_duration` - (Optional) Commitment duration requested for the Provisioned Throughput. For custom models, you can purchase on-demand Provisioned Throughput by omitting this argument. Valid values: `OneMonth`, `SixMonths`.
* `model_arn` - (Required) ARN of the model to associate with this Provisioned Throughput.
* `model_units` - (Required) Number of model units to allocate. A model unit delivers a specific throughput level for the specified model.
* `provisioned_model_name` - (Required) Unique name for this Provisioned Throughput.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `provisioned_model_arn` - The ARN of the Provisioned Throughput.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Provisioned Throughput using the `provisioned_model_arn`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.bedrock_provisioned_model_throughput import BedrockProvisionedModelThroughput
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BedrockProvisionedModelThroughput.generate_config_for_import(self, "example", "arn:aws:bedrock:us-west-2:123456789012:provisioned-model/1y5n57gh5y2e")
```

Using `terraform import`, import Provisioned Throughput using the `provisioned_model_arn`. For example:

```console
% terraform import aws_bedrock_provisioned_model_throughput.example arn:aws:bedrock:us-west-2:123456789012:provisioned-model/1y5n57gh5y2e
```

<!-- cache-key: cdktf-0.20.8 input-40ebdc15f78fbeef79d58a6dd67f830f1a83bce550db2622f392ddd0664ebb45 -->