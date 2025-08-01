---
subcategory: "WAF Classic Regional"
layout: "aws"
page_title: "AWS: aws_wafregional_regex_match_set"
description: |-
  Provides a AWS WAF Regional Regex Match Set resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_wafregional_regex_match_set

Provides a WAF Regional Regex Match Set Resource

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { WafregionalRegexMatchSet } from "./.gen/providers/aws/wafregional-regex-match-set";
import { WafregionalRegexPatternSet } from "./.gen/providers/aws/wafregional-regex-pattern-set";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new WafregionalRegexPatternSet(this, "example", {
      name: "example",
      regexPatternStrings: ["one", "two"],
    });
    const awsWafregionalRegexMatchSetExample = new WafregionalRegexMatchSet(
      this,
      "example_1",
      {
        name: "example",
        regexMatchTuple: [
          {
            fieldToMatch: {
              data: "User-Agent",
              type: "HEADER",
            },
            regexPatternSetId: example.id,
            textTransformation: "NONE",
          },
        ],
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsWafregionalRegexMatchSetExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name or description of the Regex Match Set.
* `regexMatchTuple` - (Required) The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.

### Nested Arguments

#### `regexMatchTuple`

* `fieldToMatch` - (Required) The part of a web request that you want to search, such as a specified header or a query string.
* `regexPatternSetId` - (Required) The ID of a [Regex Pattern Set](/docs/providers/aws/r/waf_regex_pattern_set.html).
* `textTransformation` - (Required) Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
  e.g., `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
  See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-TextTransformation)
  for all supported values.

#### `fieldToMatch`

* `data` - (Optional) When `type` is `HEADER`, enter the name of the header that you want to search, e.g., `User-Agent` or `Referer`.
  If `type` is any other value, omit this field.
* `type` - (Required) The part of the web request that you want AWS WAF to search for a specified string.
  e.g., `HEADER`, `METHOD` or `BODY`.
  See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
  for all supported values.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the WAF Regional Regex Match Set.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import WAF Regional Regex Match Set using the id. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { WafregionalRegexMatchSet } from "./.gen/providers/aws/wafregional-regex-match-set";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    WafregionalRegexMatchSet.generateConfigForImport(
      this,
      "example",
      "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc"
    );
  }
}

```

Using `terraform import`, import WAF Regional Regex Match Set using the id. For example:

```console
% terraform import aws_wafregional_regex_match_set.example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```

<!-- cache-key: cdktf-0.20.8 input-a88eeea459a04f400e98231ea3f0e16206c98f972102d0d251d94f97795baada -->