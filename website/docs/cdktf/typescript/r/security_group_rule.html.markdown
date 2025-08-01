---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_security_group_rule"
description: |-
  Provides an security group rule resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_security_group_rule

Provides a security group rule resource. Represents a single `ingress` or `egress` group rule, which can be added to external Security Groups.

~> **NOTE:** Avoid using the `aws_security_group_rule` resource, as it struggles with managing multiple CIDR blocks, and, due to the historical lack of unique IDs, tags and descriptions. To avoid these problems, use the current best practice of the [`aws_vpc_security_group_egress_rule`](vpc_security_group_egress_rule.html) and [`aws_vpc_security_group_ingress_rule`](vpc_security_group_ingress_rule.html) resources with one CIDR block per rule.

!> **WARNING:** You should not use the `aws_security_group_rule` resource in conjunction with [`aws_vpc_security_group_egress_rule`](vpc_security_group_egress_rule.html) and [`aws_vpc_security_group_ingress_rule`](vpc_security_group_ingress_rule.html) resources or with an [`aws_security_group`](security_group.html) resource that has in-line rules. Doing so may cause rule conflicts, perpetual differences, and result in rules being overwritten.

~> **NOTE:** Setting `protocol = "all"` or `protocol = -1` with `fromPort` and `toPort` will result in the EC2 API creating a security group rule with all ports open. This API behavior cannot be controlled by Terraform and may generate warnings in the future.

~> **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).

## Example Usage

Basic usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecurityGroupRule } from "./.gen/providers/aws/security-group-rule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SecurityGroupRule(this, "example", {
      cidrBlocks: [Token.asString(awsVpcExample.cidrBlock)],
      fromPort: 0,
      ipv6CidrBlocks: [Token.asString(awsVpcExample.ipv6CidrBlock)],
      protocol: "tcp",
      securityGroupId: "sg-123456",
      toPort: 65535,
      type: "ingress",
    });
  }
}

```

### Usage With Prefix List IDs

Prefix Lists are either managed by AWS internally, or created by the customer using a
[Managed Prefix List resource](ec2_managed_prefix_list.html). Prefix Lists provided by
AWS are associated with a prefix list name, or service name, that is linked to a specific region.

Prefix list IDs are exported on VPC Endpoints, so you can use this format:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecurityGroupRule } from "./.gen/providers/aws/security-group-rule";
import { VpcEndpoint } from "./.gen/providers/aws/vpc-endpoint";
interface MyConfig {
  vpcId: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    const myEndpoint = new VpcEndpoint(this, "my_endpoint", {
      vpcId: config.vpcId,
    });
    new SecurityGroupRule(this, "allow_all", {
      fromPort: 0,
      prefixListIds: [myEndpoint.prefixListId],
      protocol: "-1",
      securityGroupId: "sg-123456",
      toPort: 0,
      type: "egress",
    });
  }
}

```

You can also find a specific Prefix List using the [`aws_prefix_list`](/docs/providers/aws/d/prefix_list.html)
or [`ec2_managed_prefix_list`](/docs/providers/aws/d/ec2_managed_prefix_list.html) data sources:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsPrefixList } from "./.gen/providers/aws/data-aws-prefix-list";
import { DataAwsRegion } from "./.gen/providers/aws/data-aws-region";
import { SecurityGroupRule } from "./.gen/providers/aws/security-group-rule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const current = new DataAwsRegion(this, "current", {});
    const s3 = new DataAwsPrefixList(this, "s3", {
      name: "com.amazonaws.${" + current.region + "}.s3",
    });
    new SecurityGroupRule(this, "s3_gateway_egress", {
      description: "S3 Gateway Egress",
      fromPort: 443,
      prefixListIds: [Token.asString(s3.id)],
      protocol: "tcp",
      securityGroupId: "sg-123456",
      toPort: 443,
      type: "egress",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `fromPort` - (Required) Start port (or ICMP type number if protocol is "icmp" or "icmpv6").
* `protocol` - (Required) Protocol. If not icmp, icmpv6, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
* `securityGroupId` - (Required) Security group to apply this rule to.
* `toPort` - (Required) End port (or ICMP code if protocol is "icmp").
* `type` - (Required) Type of rule being created. Valid options are `ingress` (inbound)
or `egress` (outbound).

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
~> **Note** Although `cidrBlocks`, `ipv6CidrBlocks`, `prefixListIds`, and `sourceSecurityGroupId` are all marked as optional, you _must_ provide one of them in order to configure the source of the traffic.

* `cidrBlocks` - (Optional) List of CIDR blocks. Cannot be specified with `sourceSecurityGroupId` or `self`.
* `description` - (Optional) Description of the rule.
* `ipv6CidrBlocks` - (Optional) List of IPv6 CIDR blocks. Cannot be specified with `sourceSecurityGroupId` or `self`.
* `prefixListIds` - (Optional) List of Prefix List IDs.
* `self` - (Optional) Whether the security group itself will be added as a source to this ingress rule. Cannot be specified with `cidrBlocks`, `ipv6CidrBlocks`, or `sourceSecurityGroupId`.
* `sourceSecurityGroupId` - (Optional) Security group id to allow access to/from, depending on the `type`. Cannot be specified with `cidrBlocks`, `ipv6CidrBlocks`, or `self`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ID of the security group rule.
* `securityGroupRuleId` - If the `aws_security_group_rule` resource has a single source or destination then this is the AWS Security Group Rule resource ID. Otherwise it is empty.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `5m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Security Group Rules using the `securityGroupId`, `type`, `protocol`, `fromPort`, `toPort`, and source(s)/destination(s) (such as a `cidrBlock`) separated by underscores (`_`). All parts are required. For example:

**NOTE:** Not all rule permissions (e.g., not all of a rule's CIDR blocks) need to be imported for Terraform to manage rule permissions. However, importing some of a rule's permissions but not others, and then making changes to the rule will result in the creation of an additional rule to capture the updated permissions. Rule permissions that were not imported are left intact in the original rule.

Import an ingress rule in security group `sg-6e616f6d69` for TCP port 8000 with an IPv4 destination CIDR of `10.0.3.0/24`:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecurityGroupRule } from "./.gen/providers/aws/security-group-rule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SecurityGroupRule.generateConfigForImport(
      this,
      "ingress",
      "sg-6e616f6d69_ingress_tcp_8000_8000_10.0.3.0/24"
    );
  }
}

```

Import a rule with various IPv4 and IPv6 source CIDR blocks:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecurityGroupRule } from "./.gen/providers/aws/security-group-rule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SecurityGroupRule.generateConfigForImport(
      this,
      "ingress",
      "sg-4973616163_ingress_tcp_100_121_10.1.0.0/16_2001:db8::/48_10.2.0.0/16_2002:db8::/48"
    );
  }
}

```

Import a rule, applicable to all ports, with a protocol other than TCP/UDP/ICMP/ICMPV6/ALL, e.g., Multicast Transport Protocol (MTP), using the IANA protocol number. For example: 92.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecurityGroupRule } from "./.gen/providers/aws/security-group-rule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SecurityGroupRule.generateConfigForImport(
      this,
      "ingress",
      "sg-6777656e646f6c796e_ingress_92_0_65536_10.0.3.0/24_10.0.4.0/24"
    );
  }
}

```

Import a default any/any egress rule to 0.0.0.0/0:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecurityGroupRule } from "./.gen/providers/aws/security-group-rule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SecurityGroupRule.generateConfigForImport(
      this,
      "defaultEgress",
      "sg-6777656e646f6c796e_egress_all_0_0_0.0.0.0/0"
    );
  }
}

```

Import an egress rule with a prefix list ID destination:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecurityGroupRule } from "./.gen/providers/aws/security-group-rule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SecurityGroupRule.generateConfigForImport(
      this,
      "egress",
      "sg-62726f6479_egress_tcp_8000_8000_pl-6469726b"
    );
  }
}

```

Import a rule applicable to all protocols and ports with a security group source:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecurityGroupRule } from "./.gen/providers/aws/security-group-rule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SecurityGroupRule.generateConfigForImport(
      this,
      "ingressRule",
      "sg-7472697374616e_ingress_all_0_65536_sg-6176657279"
    );
  }
}

```

Import a rule that has itself and an IPv6 CIDR block as sources:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecurityGroupRule } from "./.gen/providers/aws/security-group-rule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SecurityGroupRule.generateConfigForImport(
      this,
      "ruleName",
      "sg-656c65616e6f72_ingress_tcp_80_80_self_2001:db8::/48"
    );
  }
}

```

**Using `terraform import` to import** Security Group Rules using the `securityGroupId`, `type`, `protocol`, `fromPort`, `toPort`, and source(s)/destination(s) (such as a `cidrBlock`) separated by underscores (`_`). All parts are required. For example:

**NOTE:** Not all rule permissions (e.g., not all of a rule's CIDR blocks) need to be imported for Terraform to manage rule permissions. However, importing some of a rule's permissions but not others, and then making changes to the rule will result in the creation of an additional rule to capture the updated permissions. Rule permissions that were not imported are left intact in the original rule.

Import an ingress rule in security group `sg-6e616f6d69` for TCP port 8000 with an IPv4 destination CIDR of `10.0.3.0/24`:

```console
% terraform import aws_security_group_rule.ingress sg-6e616f6d69_ingress_tcp_8000_8000_10.0.3.0/24
```

Import a rule with various IPv4 and IPv6 source CIDR blocks:

```console
% terraform import aws_security_group_rule.ingress sg-4973616163_ingress_tcp_100_121_10.1.0.0/16_2001:db8::/48_10.2.0.0/16_2002:db8::/48
```

Import a rule, applicable to all ports, with a protocol other than TCP/UDP/ICMP/ICMPV6/ALL, e.g., Multicast Transport Protocol (MTP), using the IANA protocol number. For example: 92.

```console
% terraform import aws_security_group_rule.ingress sg-6777656e646f6c796e_ingress_92_0_65536_10.0.3.0/24_10.0.4.0/24
```

Import a default any/any egress rule to 0.0.0.0/0:

```console
% terraform import aws_security_group_rule.default_egress sg-6777656e646f6c796e_egress_all_0_0_0.0.0.0/0
```

Import an egress rule with a prefix list ID destination:

```console
% terraform import aws_security_group_rule.egress sg-62726f6479_egress_tcp_8000_8000_pl-6469726b
```

Import a rule applicable to all protocols and ports with a security group source:

```console
% terraform import aws_security_group_rule.ingress_rule sg-7472697374616e_ingress_all_0_65536_sg-6176657279
```

Import a rule that has itself and an IPv6 CIDR block as sources:

```console
% terraform import aws_security_group_rule.rule_name sg-656c65616e6f72_ingress_tcp_80_80_self_2001:db8::/48
```

<!-- cache-key: cdktf-0.20.8 input-404769c8224b238d376812dbcce66e6e4526200026487201eb9a559700c3f558 -->