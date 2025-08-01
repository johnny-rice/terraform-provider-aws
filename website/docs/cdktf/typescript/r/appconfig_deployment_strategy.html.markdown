---
subcategory: "AppConfig"
layout: "aws"
page_title: "AWS: aws_appconfig_deployment_strategy"
description: |-
  Provides an AppConfig Deployment Strategy resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_appconfig_deployment_strategy

Provides an AppConfig Deployment Strategy resource.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AppconfigDeploymentStrategy } from "./.gen/providers/aws/appconfig-deployment-strategy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new AppconfigDeploymentStrategy(this, "example", {
      deploymentDurationInMinutes: 3,
      description: "Example Deployment Strategy",
      finalBakeTimeInMinutes: 4,
      growthFactor: 10,
      growthType: "LINEAR",
      name: "example-deployment-strategy-tf",
      replicateTo: "NONE",
      tags: {
        Type: "AppConfig Deployment Strategy",
      },
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `deploymentDurationInMinutes` - (Required) Total amount of time for a deployment to last. Minimum value of 0, maximum value of 1440.
* `growthFactor` - (Required) Percentage of targets to receive a deployed configuration during each interval. Minimum value of 1.0, maximum value of 100.0.
* `name` - (Required, Forces new resource) Name for the deployment strategy. Must be between 1 and 64 characters in length.
* `replicateTo` - (Required, Forces new resource) Where to save the deployment strategy. Valid values: `NONE` and `SSM_DOCUMENT`.
* `description` - (Optional) Description of the deployment strategy. Can be at most 1024 characters.
* `finalBakeTimeInMinutes` - (Optional) Amount of time AWS AppConfig monitors for alarms before considering the deployment to be complete and no longer eligible for automatic roll back. Minimum value of 0, maximum value of 1440.
* `growthType` - (Optional) Algorithm used to define how percentage grows over time. Valid value: `LINEAR` and `EXPONENTIAL`. Defaults to `LINEAR`.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - AppConfig deployment strategy ID.
* `arn` - ARN of the AppConfig Deployment Strategy.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import AppConfig Deployment Strategies using their deployment strategy ID. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AppconfigDeploymentStrategy } from "./.gen/providers/aws/appconfig-deployment-strategy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    AppconfigDeploymentStrategy.generateConfigForImport(
      this,
      "example",
      "11xxxxx"
    );
  }
}

```

Using `terraform import`, import AppConfig Deployment Strategies using their deployment strategy ID. For example:

```console
% terraform import aws_appconfig_deployment_strategy.example 11xxxxx
```

<!-- cache-key: cdktf-0.20.8 input-fe7a795a7a405b9426f1dc2cf3b7197814636553f8fda741e17e6baa30584494 -->