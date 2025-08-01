---
subcategory: "Direct Connect"
layout: "aws"
page_title: "AWS: aws_dx_bgp_peer"
description: |-
  Provides a Direct Connect BGP peer resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_dx_bgp_peer

Provides a Direct Connect BGP peer resource.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DxBgpPeer } from "./.gen/providers/aws/dx-bgp-peer";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DxBgpPeer(this, "peer", {
      addressFamily: "ipv6",
      bgpAsn: 65351,
      virtualInterfaceId: foo.id,
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `addressFamily` - (Required) The address family for the BGP peer. `ipv4 ` or `ipv6`.
* `bgpAsn` - (Required) The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
* `virtualInterfaceId` - (Required) The ID of the Direct Connect virtual interface on which to create the BGP peer.
* `amazonAddress` - (Optional) The IPv4 CIDR address to use to send traffic to Amazon.
Required for IPv4 BGP peers on public virtual interfaces.
* `bgpAuthKey` - (Optional) The authentication key for BGP configuration.
* `customerAddress` - (Optional) The IPv4 CIDR destination address to which Amazon should send traffic.
Required for IPv4 BGP peers on public virtual interfaces.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the BGP peer resource.
* `bgpStatus` - The Up/Down state of the BGP peer.
* `bgpPeerId` - The ID of the BGP peer.
* `awsDevice` - The Direct Connect endpoint on which the BGP peer terminates.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `10m`)
- `delete` - (Default `10m`)

<!-- cache-key: cdktf-0.20.8 input-f6db2d6f7b72373cb61adc2ed77c8b225974fb80ac0814395a34fd69e63ad008 -->