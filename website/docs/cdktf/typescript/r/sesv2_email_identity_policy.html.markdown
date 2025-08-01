---
subcategory: "SESv2 (Simple Email V2)"
layout: "aws"
page_title: "AWS: aws_sesv2_email_identity_policy"
description: |-
  Terraform resource for managing an AWS SESv2 (Simple Email V2) Email Identity Policy.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_sesv2_email_identity_policy

Terraform resource for managing an AWS SESv2 (Simple Email V2) Email Identity Policy.

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
import { Sesv2EmailIdentity } from "./.gen/providers/aws/sesv2-email-identity";
import { Sesv2EmailIdentityPolicy } from "./.gen/providers/aws/sesv2-email-identity-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new Sesv2EmailIdentity(this, "example", {
      emailIdentity: "testing@example.com",
    });
    const awsSesv2EmailIdentityPolicyExample = new Sesv2EmailIdentityPolicy(
      this,
      "example_1",
      {
        emailIdentity: example.emailIdentity,
        policy:
          '{\n  "Id":"ExampleAuthorizationPolicy",\n  "Version":"2012-10-17",\n  "Statement":[\n    {\n      "Sid":"AuthorizeIAMUser",\n      "Effect":"Allow",\n      "Resource":"${' +
          example.arn +
          '}",\n      "Principal":{\n        "AWS":[\n          "arn:aws:iam::123456789012:user/John",\n          "arn:aws:iam::123456789012:user/Jane"\n        ]\n      },\n      "Action":[\n        "ses:DeleteEmailIdentity",\n        "ses:PutEmailIdentityDkimSigningAttributes"\n      ]\n    }\n  ]\n}\n\n',
        policyName: "example",
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsSesv2EmailIdentityPolicyExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `emailIdentity` - (Required) The email identity.
* `policyName` - (Required) - The name of the policy.
* `policy` - (Required) - The text of the policy in JSON format.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SESv2 (Simple Email V2) Email Identity Policy using the `id` (`email_identity|policy_name`). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Sesv2EmailIdentityPolicy } from "./.gen/providers/aws/sesv2-email-identity-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    Sesv2EmailIdentityPolicy.generateConfigForImport(
      this,
      "example",
      "example_email_identity|example_policy_name"
    );
  }
}

```

Using `terraform import`, import SESv2 (Simple Email V2) Email Identity Policy using the `example_id_arg`. For example:

```console
% terraform import aws_sesv2_email_identity_policy.example example_email_identity|example_policy_name
```

<!-- cache-key: cdktf-0.20.8 input-ce24a2cee08dd43e3384c0ee51e3f6935e0859695d697797d1393bfd4fff7e2f -->