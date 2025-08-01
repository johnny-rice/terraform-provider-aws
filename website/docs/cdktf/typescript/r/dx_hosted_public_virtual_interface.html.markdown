---
subcategory: "Direct Connect"
layout: "aws"
page_title: "AWS: aws_dx_hosted_public_virtual_interface"
description: |-
  Provides a Direct Connect hosted public virtual interface resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_dx_hosted_public_virtual_interface

Provides a Direct Connect hosted public virtual interface resource. This resource represents the allocator's side of the hosted virtual interface.
A hosted virtual interface is a virtual interface that is owned by another AWS account.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DxHostedPublicVirtualInterface } from "./.gen/providers/aws/dx-hosted-public-virtual-interface";
interface MyConfig {
  ownerAccountId: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new DxHostedPublicVirtualInterface(this, "foo", {
      addressFamily: "ipv4",
      amazonAddress: "175.45.176.2/30",
      bgpAsn: 65352,
      connectionId: "dxcon-zzzzzzzz",
      customerAddress: "175.45.176.1/30",
      name: "vif-foo",
      routeFilterPrefixes: ["210.52.109.0/24", "175.45.176.0/22"],
      vlan: 4094,
      ownerAccountId: config.ownerAccountId,
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `addressFamily` - (Required) The address family for the BGP peer. `ipv4 ` or `ipv6`.
* `bgpAsn` - (Required) The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
* `connectionId` - (Required) The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
* `name` - (Required) The name for the virtual interface.
* `ownerAccountId` - (Required) The AWS account that will own the new virtual interface.
* `routeFilterPrefixes` - (Required) A list of routes to be advertised to the AWS network in this region.
* `vlan` - (Required) The VLAN ID.
* `amazonAddress` - (Optional) The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
* `bgpAuthKey` - (Optional) The authentication key for BGP configuration.
* `customerAddress` - (Optional) The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the virtual interface.
* `arn` - The ARN of the virtual interface.
* `awsDevice` - The Direct Connect endpoint on which the virtual interface terminates.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `10m`)
- `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Direct Connect hosted public virtual interfaces using the VIF `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DxHostedPublicVirtualInterface } from "./.gen/providers/aws/dx-hosted-public-virtual-interface";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    DxHostedPublicVirtualInterface.generateConfigForImport(
      this,
      "test",
      "dxvif-33cc44dd"
    );
  }
}

```

Using `terraform import`, import Direct Connect hosted public virtual interfaces using the VIF `id`. For example:

```console
% terraform import aws_dx_hosted_public_virtual_interface.test dxvif-33cc44dd
```

<!-- cache-key: cdktf-0.20.8 input-38fa3e2229ac9a4e76f540a30b62f705dac190ced8321255c81e6b8e1a4d824f -->