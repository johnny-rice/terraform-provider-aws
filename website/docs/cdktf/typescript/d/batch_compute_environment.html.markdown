---
subcategory: "Batch"
layout: "aws"
page_title: "AWS: aws_batch_compute_environment"
description: |-
    Provides details about a batch compute environment
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_batch_compute_environment

The Batch Compute Environment data source allows access to details of a specific
compute environment within AWS Batch.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsBatchComputeEnvironment } from "./.gen/providers/aws/data-aws-batch-compute-environment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsBatchComputeEnvironment(this, "batch-mongo", {
      name: "batch-mongo-production",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name of the Batch Compute Environment

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the compute environment.
* `ecsClusterArn` - ARN of the underlying Amazon ECS cluster used by the compute environment.
* `serviceRole` - ARN of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
* `type` - Type of the compute environment (for example, `MANAGED` or `UNMANAGED`).
* `status` - Current status of the compute environment (for example, `CREATING` or `VALID`).
* `statusReason` - Short, human-readable string to provide additional details about the current status of the compute environment.
* `state` - State of the compute environment (for example, `ENABLED` or `DISABLED`). If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues.
* `updatePolicy` - Specifies the infrastructure update policy for the compute environment.
* `tags` - Key-value map of resource tags

<!-- cache-key: cdktf-0.20.8 input-5222910e308777681b3cf743aa91fb7eca92faaa4af674d030446a2f131a05df -->