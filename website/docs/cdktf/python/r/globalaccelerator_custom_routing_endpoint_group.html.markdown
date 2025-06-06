---
subcategory: "Global Accelerator"
layout: "aws"
page_title: "AWS: aws_globalaccelerator_custom_routing_endpoint_group"
description: |-
  Provides a Global Accelerator custom routing endpoint group.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_globalaccelerator_custom_routing_endpoint_group

Provides a Global Accelerator custom routing endpoint group.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.globalaccelerator_custom_routing_endpoint_group import GlobalacceleratorCustomRoutingEndpointGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GlobalacceleratorCustomRoutingEndpointGroup(self, "example",
            destination_configuration=[GlobalacceleratorCustomRoutingEndpointGroupDestinationConfiguration(
                from_port=80,
                protocols=["TCP"],
                to_port=8080
            )
            ],
            endpoint_configuration=[GlobalacceleratorCustomRoutingEndpointGroupEndpointConfiguration(
                endpoint_id=Token.as_string(aws_subnet_example.id)
            )
            ],
            listener_arn=Token.as_string(aws_globalaccelerator_custom_routing_listener_example.arn)
        )
```

## Argument Reference

This resource supports the following arguments:

* `listener_arn` - (Required) The Amazon Resource Name (ARN) of the custom routing listener.
* `destination_configuration` - (Required) The port ranges and protocols for all endpoints in a custom routing endpoint group to accept client traffic on. Fields documented below.
* `endpoint_configuration` - (Optional) The list of endpoint objects. Fields documented below.
* `endpoint_group_region` (Optional) - The name of the AWS Region where the custom routing endpoint group is located.

`destination_configuration` supports the following arguments:

* `from_port` - (Required) The first port, inclusive, in the range of ports for the endpoint group that is associated with a custom routing accelerator.
* `protocols` - (Required) The protocol for the endpoint group that is associated with a custom routing accelerator. The protocol can be either `"TCP"` or `"UDP"`.
* `to_port` - (Required) The last port, inclusive, in the range of ports for the endpoint group that is associated with a custom routing accelerator.

`endpoint_configuration` supports the following arguments:

* `endpoint_id` - (Optional) An ID for the endpoint. For custom routing accelerators, this is the virtual private cloud (VPC) subnet ID.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The Amazon Resource Name (ARN) of the custom routing endpoint group.
* `arn` - The Amazon Resource Name (ARN) of the custom routing endpoint group.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Global Accelerator custom routing endpoint groups using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.globalaccelerator_custom_routing_endpoint_group import GlobalacceleratorCustomRoutingEndpointGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GlobalacceleratorCustomRoutingEndpointGroup.generate_config_for_import(self, "example", "arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxx/endpoint-group/xxxxxxxx")
```

Using `terraform import`, import Global Accelerator custom routing endpoint groups using the `id`. For example:

```console
% terraform import aws_globalaccelerator_custom_routing_endpoint_group.example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxx/endpoint-group/xxxxxxxx
```

<!-- cache-key: cdktf-0.20.8 input-28ce95de83be8065a28d60402f41212a510b8bb7ecc34edb2b8cc92538380bb7 -->