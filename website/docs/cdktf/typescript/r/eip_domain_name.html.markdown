---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "AWS: aws_eip_domain_name"
description: |-
  Assigns a static reverse DNS record to an Elastic IP addresses
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_eip_domain_name

Assigns a static reverse DNS record to an Elastic IP addresses. See [Using reverse DNS for email applications](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html#Using_Elastic_Addressing_Reverse_DNS).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Eip } from "./.gen/providers/aws/eip";
import { EipDomainName } from "./.gen/providers/aws/eip-domain-name";
import { Route53Record } from "./.gen/providers/aws/route53-record";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new Eip(this, "example", {
      domain: "vpc",
    });
    const awsRoute53RecordExample = new Route53Record(this, "example_1", {
      name: "reverse",
      records: [example.publicIp],
      type: "A",
      zoneId: main.zoneId,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsRoute53RecordExample.overrideLogicalId("example");
    const awsEipDomainNameExample = new EipDomainName(this, "example_2", {
      allocationId: example.allocationId,
      domainName: Token.asString(awsRoute53RecordExample.fqdn),
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsEipDomainNameExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `allocationId` - (Required) The allocation ID.
* `domainName` - (Required) The domain name to modify for the IP address.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `ptrRecord` - The DNS pointer (PTR) record for the IP address.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `10m`)
- `update` - (Default `10m`)
- `delete` - (Default `10m`)

<!-- cache-key: cdktf-0.20.8 input-0c05e4f12f02cf0bf8184333431906c25d66e8f7b39d4ddafc82bbfbc41c7a4d -->