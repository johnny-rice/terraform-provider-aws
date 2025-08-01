---
subcategory: "CloudWatch Logs"
layout: "aws"
page_title: "AWS: aws_cloudwatch_log_subscription_filter"
description: |-
  Provides a CloudWatch Logs subscription filter.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_cloudwatch_log_subscription_filter

Provides a CloudWatch Logs subscription filter resource.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudwatch_log_subscription_filter import CloudwatchLogSubscriptionFilter
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CloudwatchLogSubscriptionFilter(self, "test_lambdafunction_logfilter",
            destination_arn=test_logstream.arn,
            distribution="Random",
            filter_pattern="logtype test",
            log_group_name="/aws/lambda/example_lambda_name",
            name="test_lambdafunction_logfilter",
            role_arn=iam_for_lambda.arn
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) A name for the subscription filter
* `destination_arn` - (Required) The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
* `filter_pattern` - (Required) A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events. Use empty string `""` to match everything. For more information, see the [Amazon CloudWatch Logs User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).
* `log_group_name` - (Required) The name of the log group to associate the subscription filter with
* `role_arn` - (Optional) The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `aws_lambda_permission` resource for granting access from CloudWatch logs to the destination Lambda function.
* `distribution` - (Optional) The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import CloudWatch Logs subscription filter using the log group name and subscription filter name separated by `|`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudwatch_log_subscription_filter import CloudwatchLogSubscriptionFilter
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CloudwatchLogSubscriptionFilter.generate_config_for_import(self, "testLambdafunctionLogfilter", "/aws/lambda/example_lambda_name|test_lambdafunction_logfilter")
```

Using `terraform import`, import CloudWatch Logs subscription filter using the log group name and subscription filter name separated by `|`. For example:

```console
% terraform import aws_cloudwatch_log_subscription_filter.test_lambdafunction_logfilter "/aws/lambda/example_lambda_name|test_lambdafunction_logfilter"
```

<!-- cache-key: cdktf-0.20.8 input-7288759b6c5cff35c95fb747308a50c7c06f176906bfc5ee2d893f8aebb2f6f4 -->