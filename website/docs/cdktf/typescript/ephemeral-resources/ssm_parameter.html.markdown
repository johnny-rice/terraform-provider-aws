---
subcategory: "SSM (Systems Manager)"
layout: "aws"
page_title: "AWS: aws_ssm_parameter"
description: |-
  Retrieve information about an SSM parameter, including its value. To retrieve parameter metadata, see the `aws_ssm_parameter` data source.
---


<!-- Please do not edit this file, it is generated. -->
# Ephemeral: aws_ssm_parameter

Retrieve information about an SSM parameter, including its value.

~> **NOTE:** Ephemeral resources are a new feature and may evolve as we continue to explore their most effective uses. [Learn more](https://developer.hashicorp.com/terraform/language/resources/ephemeral).

## Example Usage

### Retrieve an SSM parameter

By default, this ephemeral resource attempts to return decrypted values for secure string parameters.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `arn` - (Required) The Amazon Resource Name (ARN) of the parameter that you want to query
* `withDecryption` - (Optional) Return decrypted values for a secure string parameter (Defaults to `true`).

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `name` - The name of the parameter.
* `type` - The type of parameter.
* `value` - The parameter value.
* `version` - The parameter version.
* `withDecryption` - Indicates whether the secure string parameters were decrypted.

<!-- cache-key: cdktf-0.20.8 input-6fc477cd1a26e75163f3bc9e97e96d1cd971e51c67a5bada2cc3806403da32f4 -->