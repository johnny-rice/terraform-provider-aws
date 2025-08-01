---
subcategory: "Macie"
layout: "aws"
page_title: "AWS: aws_macie2_custom_data_identifier"
description: |-
  Provides a resource to manage an AWS Macie Custom Data Identifier.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_macie2_custom_data_identifier

Provides a resource to manage an [AWS Macie Custom Data Identifier](https://docs.aws.amazon.com/macie/latest/APIReference/custom-data-identifiers-id.html).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Macie2Account } from "./.gen/providers/aws/macie2-account";
import { Macie2CustomDataIdentifier } from "./.gen/providers/aws/macie2-custom-data-identifier";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new Macie2Account(this, "example", {});
    const awsMacie2CustomDataIdentifierExample = new Macie2CustomDataIdentifier(
      this,
      "example_1",
      {
        dependsOn: [test],
        description: "DESCRIPTION",
        ignoreWords: ["ignore"],
        keywords: ["keyword"],
        maximumMatchDistance: 10,
        name: "NAME OF CUSTOM DATA IDENTIFIER",
        regex: "[0-9]{3}-[0-9]{2}-[0-9]{4}",
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsMacie2CustomDataIdentifierExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `regex` - (Optional) The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
* `keywords` -  (Optional) An array that lists specific character sequences (keywords), one of which must be within proximity (`maximumMatchDistance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren't case sensitive.
* `ignoreWords` - (Optional) An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
* `name` - (Optional) A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, Terraform will assign a random, unique name. Conflicts with `namePrefix`.
* `namePrefix` -  (Optional) Creates a unique name beginning with the specified prefix. Conflicts with `name`.
* `description` - (Optional) A custom description of the custom data identifier. The description can contain as many as 512 characters.
* `maximumMatchDistance` - (Optional) The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The unique identifier (ID) of the macie custom data identifier.
* `deleted` - Specifies whether the custom data identifier was deleted. If you delete a custom data identifier, Amazon Macie doesn't delete it permanently. Instead, it soft deletes the identifier.
* `createdAt` - The date and time, in UTC and extended RFC 3339 format, when the Amazon Macie account was created.
* `arn` - The Amazon Resource Name (ARN) of the custom data identifier.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_macie2_custom_data_identifier` using the id. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Macie2CustomDataIdentifier } from "./.gen/providers/aws/macie2-custom-data-identifier";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    Macie2CustomDataIdentifier.generateConfigForImport(
      this,
      "example",
      "abcd1"
    );
  }
}

```

Using `terraform import`, import `aws_macie2_custom_data_identifier` using the id. For example:

```console
% terraform import aws_macie2_custom_data_identifier.example abcd1
```

<!-- cache-key: cdktf-0.20.8 input-4124c76245725a8b00487b8803a0d2aa927bd775cd3d35323d70147e836a4cd4 -->