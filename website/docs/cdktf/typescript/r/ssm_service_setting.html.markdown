---
subcategory: "SSM (Systems Manager)"
layout: "aws"
page_title: "AWS: aws_ssm_service_setting"
description: |-
  Defines how a user interacts with or uses a service or a feature of a service.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ssm_service_setting

This setting defines how a user interacts with or uses a service or a feature of a service.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmServiceSetting } from "./.gen/providers/aws/ssm-service-setting";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SsmServiceSetting(this, "test_setting", {
      settingId:
        "arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled",
      settingValue: "true",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `settingId` - (Required) ID of the service setting.
* `settingValue` - (Required) Value of the service setting.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the service setting.
* `status` - Status of the service setting. Value can be `Default`, `Customized` or `PendingUpdate`.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import AWS SSM Service Setting using the `settingId`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmServiceSetting } from "./.gen/providers/aws/ssm-service-setting";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SsmServiceSetting.generateConfigForImport(
      this,
      "example",
      "arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled"
    );
  }
}

```

Using `terraform import`, import AWS SSM Service Setting using the `settingId`. For example:

```console
% terraform import aws_ssm_service_setting.example arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled
```

<!-- cache-key: cdktf-0.20.8 input-a6fa90154400e1c2ad4a6df2ea58199a9d9539c71573319019c6d0832606c0dd -->