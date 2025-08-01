---
subcategory: "Network Manager"
layout: "aws"
page_title: "AWS: aws_networkmanager_transit_gateway_route_table_attachment"
description: |-
  Manages a Network Manager transit gateway route table attachment.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_networkmanager_transit_gateway_route_table_attachment

Manages a Network Manager transit gateway route table attachment.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkmanager_transit_gateway_route_table_attachment import NetworkmanagerTransitGatewayRouteTableAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NetworkmanagerTransitGatewayRouteTableAttachment(self, "example",
            peering_id=Token.as_string(aws_networkmanager_transit_gateway_peering_example.id),
            transit_gateway_route_table_arn=Token.as_string(aws_ec2_transit_gateway_route_table_example.arn)
        )
```

## Argument Reference

The following arguments are required:

* `peering_id` - (Required) ID of the peer for the attachment.
* `transit_gateway_route_table_arn` - (Required) ARN of the transit gateway route table for the attachment.

The following arguments are optional:

* `tags` - (Optional) Key-value tags for the attachment. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Attachment ARN.
* `attachment_policy_rule_number` - Policy rule number associated with the attachment.
* `attachment_type` - Type of attachment.
* `core_network_arn` - ARN of the core network.
* `core_network_id` - ID of the core network.
* `edge_location` - Edge location for the peer.
* `id` - ID of the attachment.
* `owner_account_id` - ID of the attachment account owner.
* `resource_arn` - Attachment resource ARN.
* `segment_name` - Name of the segment attachment.
* `state` - State of the attachment.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `10m`)
* `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_networkmanager_transit_gateway_route_table_attachment` using the attachment ID. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkmanager_transit_gateway_route_table_attachment import NetworkmanagerTransitGatewayRouteTableAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NetworkmanagerTransitGatewayRouteTableAttachment.generate_config_for_import(self, "example", "attachment-0f8fa60d2238d1bd8")
```

Using `terraform import`, import `aws_networkmanager_transit_gateway_route_table_attachment` using the attachment ID. For example:

```console
% terraform import aws_networkmanager_transit_gateway_route_table_attachment.example attachment-0f8fa60d2238d1bd8
```

<!-- cache-key: cdktf-0.20.8 input-48c988a2a19cb9db3ba7901f8a78831fe0cb80082a5f76e7980ec29b5bbfe8f4 -->