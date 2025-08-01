---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "AWS: aws_ec2_instance_state"
description: |-
  Provides an EC2 instance state resource. This allows managing an instance power state. 
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ec2_instance_state

Provides an EC2 instance state resource. This allows managing an instance power state.

~> **NOTE on Instance State Management:** AWS does not currently have an EC2 API operation to determine an instance has finished processing user data. As a result, this resource can interfere with user data processing. For example, this resource may stop an instance while the user data script is in mid run.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsAmi } from "./.gen/providers/aws/data-aws-ami";
import { Ec2InstanceState } from "./.gen/providers/aws/ec2-instance-state";
import { Instance } from "./.gen/providers/aws/instance";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const ubuntu = new DataAwsAmi(this, "ubuntu", {
      filter: [
        {
          name: "name",
          values: ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"],
        },
        {
          name: "virtualization-type",
          values: ["hvm"],
        },
      ],
      mostRecent: true,
      owners: ["099720109477"],
    });
    const test = new Instance(this, "test", {
      ami: Token.asString(ubuntu.id),
      instanceType: "t3.micro",
      tags: {
        Name: "HelloWorld",
      },
    });
    const awsEc2InstanceStateTest = new Ec2InstanceState(this, "test_2", {
      instanceId: test.id,
      state: "stopped",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsEc2InstanceStateTest.overrideLogicalId("test");
  }
}

```

## Argument Reference

The following arguments are required:

* `instanceId` - (Required) ID of the instance.
* `state` - (Required) - State of the instance. Valid values are `stopped`, `running`.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `force` - (Optional) Whether to request a forced stop when `state` is `stopped`. Otherwise (_i.e._, `state` is `running`), ignored. When an instance is forced to stop, it does not flush file system caches or file system metadata, and you must subsequently perform file system check and repair. Not recommended for Windows instances. Defaults to `false`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ID of the instance (matches `instanceId`).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `10m`)
* `update` - (Default `10m`)
* `delete` - (Default `1m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_ec2_instance_state` using the `instanceId` attribute. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Ec2InstanceState } from "./.gen/providers/aws/ec2-instance-state";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    Ec2InstanceState.generateConfigForImport(
      this,
      "test",
      "i-02cae6557dfcf2f96"
    );
  }
}

```

Using `terraform import`, import `aws_ec2_instance_state` using the `instanceId` attribute. For example:

```console
% terraform import aws_ec2_instance_state.test i-02cae6557dfcf2f96
```

<!-- cache-key: cdktf-0.20.8 input-e81886535285ebdaa3178d6e3d2e03d19e266715ad698aea787d8451ac6755da -->