---
subcategory: "SQS (Simple Queue)"
layout: "aws"
page_title: "AWS: aws_sqs_queue_redrive_policy"
description: |-
  Provides a SQS Queue Redrive Policy resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_sqs_queue_redrive_policy

Allows you to set a redrive policy of an SQS Queue
while referencing ARN of the dead letter queue inside the redrive policy.

This is useful when you want to set a dedicated
dead letter queue for a standard or FIFO queue, but need
the dead letter queue to exist before setting the redrive policy.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sqs_queue import SqsQueue
from imports.aws.sqs_queue_redrive_policy import SqsQueueRedrivePolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        q = SqsQueue(self, "q",
            name="examplequeue"
        )
        ddl = SqsQueue(self, "ddl",
            name="examplequeue-ddl",
            redrive_allow_policy=Token.as_string(
                Fn.jsonencode({
                    "redrive_permission": "byQueue",
                    "source_queue_arns": [q.arn]
                }))
        )
        aws_sqs_queue_redrive_policy_q = SqsQueueRedrivePolicy(self, "q_2",
            queue_url=q.id,
            redrive_policy=Token.as_string(
                Fn.jsonencode({
                    "dead_letter_target_arn": ddl.arn,
                    "max_receive_count": 4
                }))
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_sqs_queue_redrive_policy_q.override_logical_id("q")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `queue_url` - (Required) The URL of the SQS Queue to which to attach the policy
* `redrive_policy` - (Required) The JSON redrive policy for the SQS queue. Accepts two key/val pairs: `deadLetterTargetArn` and `maxReceiveCount`. Learn more in the [Amazon SQS dead-letter queues documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html).

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SQS Queue Redrive Policies using the queue URL. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sqs_queue_redrive_policy import SqsQueueRedrivePolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SqsQueueRedrivePolicy.generate_config_for_import(self, "test", "https://queue.amazonaws.com/123456789012/myqueue")
```

Using `terraform import`, import SQS Queue Redrive Policies using the queue URL. For example:

```console
% terraform import aws_sqs_queue_redrive_policy.test https://queue.amazonaws.com/123456789012/myqueue
```

<!-- cache-key: cdktf-0.20.8 input-28acac650d5066bd826ddd5e642e0cd427fccd7be1cf280f8df9e3a8bf7bd5ea -->