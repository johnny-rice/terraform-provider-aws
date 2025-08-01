---
subcategory: "Cloud Map"
layout: "aws"
page_title: "AWS: aws_service_discovery_private_dns_namespace"
description: |-
  Provides a Service Discovery Private DNS Namespace resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_service_discovery_private_dns_namespace

Provides a Service Discovery Private DNS Namespace resource.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.service_discovery_private_dns_namespace import ServiceDiscoveryPrivateDnsNamespace
from imports.aws.vpc import Vpc
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = Vpc(self, "example",
            cidr_block="10.0.0.0/16"
        )
        aws_service_discovery_private_dns_namespace_example =
        ServiceDiscoveryPrivateDnsNamespace(self, "example_1",
            description="example",
            name="hoge.example.local",
            vpc=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_service_discovery_private_dns_namespace_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of the namespace.
* `vpc` - (Required) The ID of VPC that you want to associate the namespace with.
* `description` - (Optional) The description that you specify for the namespace when you create it.
* `tags` - (Optional) A map of tags to assign to the namespace. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of a namespace.
* `arn` - The ARN that Amazon Route 53 assigns to the namespace when you create it.
* `hosted_zone` - The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Service Discovery Private DNS Namespace using the namespace ID and VPC ID. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.service_discovery_private_dns_namespace import ServiceDiscoveryPrivateDnsNamespace
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ServiceDiscoveryPrivateDnsNamespace.generate_config_for_import(self, "example", "0123456789:vpc-123345")
```

Using `terraform import`, import Service Discovery Private DNS Namespace using the namespace ID and VPC ID. For example:

```console
% terraform import aws_service_discovery_private_dns_namespace.example 0123456789:vpc-123345
```

<!-- cache-key: cdktf-0.20.8 input-d6c1ed864d293a758ae77c9c19ff9a4838756189f976fb27557d8890e92245e0 -->