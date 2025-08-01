---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "AWS: aws_ec2_spot_price"
description: |-
  Information about most recent Spot Price for a given EC2 instance.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ec2_spot_price

Information about most recent Spot Price for a given EC2 instance.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsEc2SpotPrice } from "./.gen/providers/aws/data-aws-ec2-spot-price";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsEc2SpotPrice(this, "example", {
      availabilityZone: "us-west-2a",
      filter: [
        {
          name: "product-description",
          values: ["Linux/UNIX"],
        },
      ],
      instanceType: "t3.medium",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `instanceType` - (Optional) Type of instance for which to query Spot Price information.
* `availabilityZone` - (Optional) Availability zone in which to query Spot price information.
* `filter` - (Optional) One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSpotPriceHistory.html) for supported filters. Detailed below.

### filter Argument Reference

* `name` - (Required) Name of the filter.
* `values` - (Required) List of one or more values for the filter.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - AWS Region.
* `spotPrice` - Most recent Spot Price value for the given instance type and AZ.
* `spotPriceTimestamp` - The timestamp at which the Spot Price value was published.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `read` - (Default `20m`)

<!-- cache-key: cdktf-0.20.8 input-e9b9dfd7bc6ff5b480d0b0f0cb7d874b0041eaf3adf19cce410878c7e2866dd6 -->