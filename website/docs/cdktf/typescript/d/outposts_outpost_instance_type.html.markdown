---
subcategory: "Outposts"
layout: "aws"
page_title: "AWS: aws_outposts_outpost_instance_type"
description: |-
  Information about single Outpost Instance Type.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_outposts_outpost_instance_type

Information about single Outpost Instance Type.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Ec2Instance } from "./.gen/providers/aws/";
import { DataAwsOutpostsOutpostInstanceType } from "./.gen/providers/aws/data-aws-outposts-outpost-instance-type";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new DataAwsOutpostsOutpostInstanceType(this, "example", {
      arn: Token.asString(dataAwsOutpostsOutpostExample.arn),
      preferredInstanceTypes: ["m5.large", "m5.4xlarge"],
    });
    const awsEc2InstanceExample = new Ec2Instance(this, "example_1", {
      instance_type: example.instanceType,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsEc2InstanceExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

The following arguments are required:

* `arn` - (Required) Outpost ARN.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `instanceType` - (Optional) Desired instance type. Conflicts with `preferredInstanceTypes`.
* `preferredInstanceTypes` - (Optional) Ordered list of preferred instance types. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned. Conflicts with `instanceType`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - Outpost identifier.

<!-- cache-key: cdktf-0.20.8 input-69ff1ef46ac8e63e892f42489f4fd3b1db4859cd1c18688be48db0003a2f958e -->