---
subcategory: "SES (Simple Email)"
layout: "aws"
page_title: "AWS: aws_ses_receipt_filter"
description: |-
  Provides an SES receipt filter
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ses_receipt_filter

Provides an SES receipt filter resource

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SesReceiptFilter } from "./.gen/providers/aws/ses-receipt-filter";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SesReceiptFilter(this, "filter", {
      cidr: "10.10.10.10",
      name: "block-spammer",
      policy: "Block",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of the filter
* `cidr` - (Required) The IP address or address range to filter, in CIDR notation
* `policy` - (Required) Block or Allow

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The SES receipt filter name.
* `arn` - The SES receipt filter ARN.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SES Receipt Filter using their `name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SesReceiptFilter } from "./.gen/providers/aws/ses-receipt-filter";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SesReceiptFilter.generateConfigForImport(this, "test", "some-filter");
  }
}

```

Using `terraform import`, import SES Receipt Filter using their `name`. For example:

```console
% terraform import aws_ses_receipt_filter.test some-filter
```

<!-- cache-key: cdktf-0.20.8 input-48c3f7575e147ed9b42a3b5fa1d9cdb485e9f550dc7379b068bffcc5272dcbe5 -->