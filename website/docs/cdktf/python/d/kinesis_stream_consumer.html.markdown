---
subcategory: "Kinesis"
layout: "aws"
page_title: "AWS: aws_kinesis_stream_consumer"
description: |-
  Provides details about a Kinesis Stream Consumer.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_kinesis_stream_consumer

Provides details about a Kinesis Stream Consumer.

For more details, see the [Amazon Kinesis Stream Consumer Documentation][1].

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_kinesis_stream_consumer import DataAwsKinesisStreamConsumer
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsKinesisStreamConsumer(self, "example",
            name="example-consumer",
            stream_arn=Token.as_string(aws_kinesis_stream_example.arn)
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `arn` - (Optional) ARN of the stream consumer.
* `name` - (Optional) Name of the stream consumer.
* `stream_arn` - (Required) ARN of the data stream the consumer is registered with.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `creation_timestamp` - Approximate timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of when the stream consumer was created.
* `id` - ARN of the stream consumer.
* `status` - Current status of the stream consumer.

[1]: https://docs.aws.amazon.com/streams/latest/dev/amazon-kinesis-consumers.html

<!-- cache-key: cdktf-0.20.8 input-77f6ef084f65362b06bfa52aa7180df46916b9f41414986b7770811ac59bdba2 -->