---
subcategory: "Verified Access"
layout: "aws"
page_title: "AWS: aws_verifiedaccess_trust_provider"
description: |-
  Terraform resource for managing a Verified Access Trust Provider.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_verifiedaccess_trust_provider

Terraform resource for managing a Verified Access Trust Provider.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VerifiedaccessTrustProvider } from "./.gen/providers/aws/verifiedaccess-trust-provider";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new VerifiedaccessTrustProvider(this, "example", {
      policyReferenceName: "example",
      trustProviderType: "user",
      userTrustProviderType: "iam-identity-center",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `policyReferenceName` - (Required) The identifier to be used when working with policy rules.
* `trustProviderType` - (Required) The type of trust provider can be either user or device-based.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `description` - (Optional) A description for the AWS Verified Access trust provider.
* `deviceOptions` - (Optional) A block of options for device identity based trust providers.
* `deviceTrustProviderType` (Optional) The type of device-based trust provider.
* `nativeApplicationOidcOptions` - (Optional) The OpenID Connect details for an Native Application OIDC, user-identity based trust provider.
* `oidcOptions` - (Optional) The OpenID Connect details for an oidc-type, user-identity based trust provider.
* `tags` - (Optional) Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `userTrustProviderType` - (Optional) The type of user-based trust provider.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the AWS Verified Access trust provider.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `60m`)
* `update` - (Default `180m`)
* `delete` - (Default `90m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Transfer Workflows using the `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VerifiedaccessTrustProvider } from "./.gen/providers/aws/verifiedaccess-trust-provider";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    VerifiedaccessTrustProvider.generateConfigForImport(
      this,
      "example",
      "vatp-8012925589"
    );
  }
}

```

Using `terraform import`, import Transfer Workflows using the  `id`. For example:

```console
% terraform import aws_verifiedaccess_trust_provider.example vatp-8012925589
```

<!-- cache-key: cdktf-0.20.8 input-5e8db91fac8de5b8e24e4eb6d199ffed067199fbe846fba732cd10106d70d3ec -->