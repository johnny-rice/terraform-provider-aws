---
subcategory: "OpenSearch Ingestion"
layout: "aws"
page_title: "AWS: aws_osis_pipeline"
description: |-
  Terraform resource for managing an AWS OpenSearch Ingestion Pipeline.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_osis_pipeline

Terraform resource for managing an AWS OpenSearch Ingestion Pipeline.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_region import DataAwsRegion
from imports.aws.iam_role import IamRole
from imports.aws.osis_pipeline import OsisPipeline
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = IamRole(self, "example",
            assume_role_policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "osis-pipelines.amazonaws.com"
                        },
                        "Sid": ""
                    }
                    ],
                    "Version": "2012-10-17"
                }))
        )
        current = DataAwsRegion(self, "current")
        aws_osis_pipeline_example = OsisPipeline(self, "example_2",
            max_units=1,
            min_units=1,
            pipeline_configuration_body="version: \"2\"\nexample-pipeline:\n  source:\n    http:\n      path: \"/example\"\n  sink:\n    - s3:\n        aws:\n          sts_role_arn: \"${" + example.arn + "}\"\n          region: \"${" + current.region + "}\"\n        bucket: \"example\"\n        threshold:\n          event_collect_timeout: \"60s\"\n        codec:\n          ndjson:\n\n",
            pipeline_name="example"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_osis_pipeline_example.override_logical_id("example")
```

### Using file function

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.osis_pipeline import OsisPipeline
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        OsisPipeline(self, "example",
            max_units=1,
            min_units=1,
            pipeline_configuration_body=Token.as_string(Fn.file("example.yaml")),
            pipeline_name="example"
        )
```

## Argument Reference

The following arguments are required:

* `max_units` - (Required) The maximum pipeline capacity, in Ingestion Compute Units (ICUs).
* `min_units` - (Required) The minimum pipeline capacity, in Ingestion Compute Units (ICUs).
* `pipeline_configuration_body` - (Required) The pipeline configuration in YAML format. This argument accepts the pipeline configuration as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \n.
* `pipeline_name` - (Required) The name of the OpenSearch Ingestion pipeline to create. Pipeline names are unique across the pipelines owned by an account within an AWS Region.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `buffer_options` - (Optional) Key-value pairs to configure persistent buffering for the pipeline. See [`buffer_options`](#buffer_options) below.
* `encryption_at_rest_options` - (Optional) Key-value pairs to configure encryption for data that is written to a persistent buffer. See [`encryption_at_rest_options`](#encryption_at_rest_options) below.
* `log_publishing_options` - (Optional) Key-value pairs to configure log publishing. See [`log_publishing_options`](#log_publishing_options) below.
* `tags` - (Optional) A map of tags to assign to the pipeline. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `vpc_options` - (Optional) Container for the values required to configure VPC access for the pipeline. If you don't specify these values, OpenSearch Ingestion creates the pipeline with a public endpoint. See [`vpc_options`](#vpc_options) below.

### buffer_options

* `persistent_buffer_enabled` - (Required) Whether persistent buffering should be enabled.

### encryption_at_rest_options

* `kms_key_arn` - (Required) The ARN of the KMS key used to encrypt data-at-rest in OpenSearch Ingestion. By default, data is encrypted using an AWS owned key.

### log_publishing_options

* `cloudwatch_log_destination` - (Optional) The destination for OpenSearch Ingestion logs sent to Amazon CloudWatch Logs. This parameter is required if IsLoggingEnabled is set to true. See [`cloudwatch_log_destination`](#cloudwatch_log_destination) below.
* `is_logging_enabled` - (Optional) Whether logs should be published.

### cloudwatch_log_destination

* `log_group` - (Required) The name of the CloudWatch Logs group to send pipeline logs to. You can specify an existing log group or create a new one. For example, /aws/OpenSearchService/IngestionService/my-pipeline.

### vpc_options

* `subnet_ids` - (Required) A list of subnet IDs associated with the VPC endpoint.
* `security_group_ids` - (Optional) A list of security groups associated with the VPC endpoint.
* `vpc_endpoint_management` - (Optional) Whether you or Amazon OpenSearch Ingestion service create and manage the VPC endpoint configured for the pipeline. Valid values are `CUSTOMER` or `SERVICE`

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Unique identifier for the pipeline.
* `ingest_endpoint_urls` - The list of ingestion endpoints for the pipeline, which you can send data to.
* `pipeline_arn` - Amazon Resource Name (ARN) of the pipeline.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `45m`)
* `update` - (Default `45m`)
* `delete` - (Default `45m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import OpenSearch Ingestion Pipeline using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.osis_pipeline import OsisPipeline
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        OsisPipeline.generate_config_for_import(self, "example", "example")
```

Using `terraform import`, import OpenSearch Ingestion Pipeline using the `id`. For example:

```console
% terraform import aws_osis_pipeline.example example
```

<!-- cache-key: cdktf-0.20.8 input-817bcbbf4e9bfa4e3a60baf5c33695fdcd1f212f25e98f4625d85e2863577c16 -->