---
subcategory: "WAF Classic Regional"
layout: "aws"
page_title: "AWS: aws_wafregional_subscribed_rule_group"
description: |-
  retrieves information about a Managed WAF Rule Group from AWS Marketplace for use in WAF Regional.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_wafregional_subscribed_rule_group

`aws_wafregional_subscribed_rule_group` retrieves information about a Managed WAF Rule Group from AWS Marketplace for use in WAF Regional (needs to be subscribed to first).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsWafregionalSubscribedRuleGroup } from "./.gen/providers/aws/data-aws-wafregional-subscribed-rule-group";
import { WafregionalWebAcl } from "./.gen/providers/aws/wafregional-web-acl";
interface MyConfig {
  defaultAction: any;
  metricName: any;
  name: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    const byMetricName = new DataAwsWafregionalSubscribedRuleGroup(
      this,
      "by_metric_name",
      {
        metricName: "F5BotDetectionSignatures",
      }
    );
    const byName = new DataAwsWafregionalSubscribedRuleGroup(this, "by_name", {
      name: "F5 Bot Detection Signatures For AWS WAF",
    });
    new WafregionalWebAcl(this, "acl", {
      rules: [
        {
          priority: 1,
          rule_id: byName.id,
          type: "GROUP",
        },
        {
          priority: 2,
          rule_id: byMetricName.id,
          type: "GROUP",
        },
      ],
      defaultAction: config.defaultAction,
      metricName: config.metricName,
      name: config.name,
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Optional) Name of the WAF rule group.
* `metricName` - (Optional) Name of the WAF rule group.

At least one of `name` or `metricName` must be configured.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - ID of the WAF rule group.

<!-- cache-key: cdktf-0.20.8 input-bb70bc005581ca28e3644651f9c5f2d63143581a987b63ace552d5a355591bf9 -->