---
subcategory: "CodeGuru Profiler"
layout: "aws"
page_title: "AWS: aws_codeguruprofiler_profiling_group"
description: |-
  Terraform resource for managing an AWS CodeGuru Profiler Profiling Group.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_codeguruprofiler_profiling_group

Terraform resource for managing an AWS CodeGuru Profiler Profiling Group.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CodeguruprofilerProfilingGroup } from "./.gen/providers/aws/codeguruprofiler-profiling-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new CodeguruprofilerProfilingGroup(this, "example", {
      agentOrchestrationConfig: [
        {
          profilingEnabled: true,
        },
      ],
      computePlatform: "Default",
      name: "example",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `agentOrchestrationConfig` - (Required) Specifies whether profiling is enabled or disabled for the created profiling. See [Agent Orchestration Config](#agent-orchestration-config) for more details.
* `name` - (Required) Name of the profiling group.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `computePlatform` - (Optional) Compute platform of the profiling group.
* `tags` - (Optional) Map of tags assigned to the resource. If configured with a provider [`defaultTags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the profiling group.
* `id` - Name of the profiling group.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

### Agent Orchestration Config

* `profilingEnabled` - (Required) Boolean that specifies whether the profiling agent collects profiling data or

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import CodeGuru Profiler Profiling Group using the `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CodeguruprofilerProfilingGroup } from "./.gen/providers/aws/codeguruprofiler-profiling-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    CodeguruprofilerProfilingGroup.generateConfigForImport(
      this,
      "example",
      "profiling_group-name-12345678"
    );
  }
}

```

Using `terraform import`, import CodeGuru Profiler Profiling Group using the `id`. For example:

```console
% terraform import aws_codeguruprofiler_profiling_group.example profiling_group-name-12345678
```

<!-- cache-key: cdktf-0.20.8 input-52ebe093e45d478f72beed478958b9875d941b71304383ce57ba89350891a519 -->