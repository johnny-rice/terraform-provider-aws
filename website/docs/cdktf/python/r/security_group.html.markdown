---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_security_group"
description: |-
  Provides a security group resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_security_group

Provides a security group resource.

~> **NOTE:** Avoid using the `ingress` and `egress` arguments of the `aws_security_group` resource to configure in-line rules, as they struggle with managing multiple CIDR blocks, and, due to the historical lack of unique IDs, tags and descriptions. To avoid these problems, use the current best practice of the [`aws_vpc_security_group_egress_rule`](vpc_security_group_egress_rule.html) and [`aws_vpc_security_group_ingress_rule`](vpc_security_group_ingress_rule.html) resources with one CIDR block per rule.

!> **WARNING:** You should not use the `aws_security_group` resource with _in-line rules_ (using the `ingress` and `egress` arguments of `aws_security_group`) in conjunction with the [`aws_vpc_security_group_egress_rule`](vpc_security_group_egress_rule.html) and [`aws_vpc_security_group_ingress_rule`](vpc_security_group_ingress_rule.html) resources or the [`aws_security_group_rule`](security_group_rule.html) resource. Doing so may cause rule conflicts, perpetual differences, and result in rules being overwritten.

~> **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).

~> **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), security groups associated with Lambda Functions can take up to 45 minutes to successfully delete. Terraform AWS Provider version 2.31.0 and later automatically handles this increased timeout, however prior versions require setting the [customizable deletion timeout](#timeouts) to 45 minutes (`delete = "45m"`). AWS and HashiCorp are working together to reduce the amount of time required for resource deletion and updates can be tracked in this [GitHub issue](https://github.com/hashicorp/terraform-provider-aws/issues/10329).

~> **NOTE:** The `cidr_blocks` and `ipv6_cidr_blocks` parameters are optional in the `ingress` and `egress` blocks. If nothing is specified, traffic will be blocked as described in _NOTE on Egress rules_ later.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.security_group import SecurityGroup
from imports.aws.vpc_security_group_egress_rule import VpcSecurityGroupEgressRule
from imports.aws.vpc_security_group_ingress_rule import VpcSecurityGroupIngressRule
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        allow_tls = SecurityGroup(self, "allow_tls",
            description="Allow TLS inbound traffic and all outbound traffic",
            name="allow_tls",
            tags={
                "Name": "allow_tls"
            },
            vpc_id=main.id
        )
        VpcSecurityGroupEgressRule(self, "allow_all_traffic_ipv4",
            cidr_ipv4="0.0.0.0/0",
            ip_protocol="-1",
            security_group_id=allow_tls.id
        )
        VpcSecurityGroupEgressRule(self, "allow_all_traffic_ipv6",
            cidr_ipv6="::/0",
            ip_protocol="-1",
            security_group_id=allow_tls.id
        )
        VpcSecurityGroupIngressRule(self, "allow_tls_ipv4",
            cidr_ipv4=main.cidr_block,
            from_port=443,
            ip_protocol="tcp",
            security_group_id=allow_tls.id,
            to_port=443
        )
        VpcSecurityGroupIngressRule(self, "allow_tls_ipv6",
            cidr_ipv6=main.ipv6_cidr_block,
            from_port=443,
            ip_protocol="tcp",
            security_group_id=allow_tls.id,
            to_port=443
        )
```

~> **NOTE on Egress rules:** By default, AWS creates an `ALLOW ALL` egress rule when creating a new Security Group inside of a VPC. When creating a new Security Group inside a VPC, **Terraform will remove this default rule**, and require you specifically re-create it if you desire that rule. We feel this leads to fewer surprises in terms of controlling your egress rules. If you desire this rule to be in place, you can use this `egress` block:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.security_group import SecurityGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SecurityGroup(self, "example",
            egress=[SecurityGroupEgress(
                cidr_blocks=["0.0.0.0/0"],
                from_port=0,
                ipv6_cidr_blocks=["::/0"],
                protocol="-1",
                to_port=0
            )
            ]
        )
```

### Usage With Prefix List IDs

Prefix Lists are either managed by AWS internally, or created by the customer using a
[Prefix List resource](ec2_managed_prefix_list.html). Prefix Lists provided by
AWS are associated with a prefix list name, or service name, that is linked to a specific region.
Prefix list IDs are exported on VPC Endpoints, so you can use this format:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.security_group import SecurityGroup
from imports.aws.vpc_endpoint import VpcEndpoint
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, vpcId):
        super().__init__(scope, name)
        my_endpoint = VpcEndpoint(self, "my_endpoint",
            vpc_id=vpc_id
        )
        SecurityGroup(self, "example",
            egress=[SecurityGroupEgress(
                from_port=0,
                prefix_list_ids=[my_endpoint.prefix_list_id],
                protocol="-1",
                to_port=0
            )
            ]
        )
```

You can also find a specific Prefix List using the `aws_prefix_list` data source.

### Removing All Ingress and Egress Rules

The `ingress` and `egress` arguments are processed in [attributes-as-blocks](https://developer.hashicorp.com/terraform/language/attr-as-blocks) mode. Due to this, removing these arguments from the configuration will **not** cause Terraform to destroy the managed rules. To subsequently remove all managed ingress and egress rules:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.security_group import SecurityGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SecurityGroup(self, "example",
            egress=[],
            ingress=[],
            name="sg",
            vpc_id=Token.as_string(aws_vpc_example.id)
        )
```

### Recreating a Security Group

A simple security group `name` change "forces new" the security group--Terraform destroys the security group and creates a new one. (Likewise, `description`, `name_prefix`, or `vpc_id` [cannot be changed](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/working-with-security-groups.html#creating-security-group).) Attempting to recreate the security group leads to a variety of complications depending on how it is used.

Security groups are generally associated with other resources--**more than 100** AWS Provider resources reference security groups. Referencing a resource from another resource creates a one-way dependency. For example, if you create an EC2 `aws_instance` that has a `vpc_security_group_ids` argument that refers to an `aws_security_group` resource, the `aws_security_group` is a dependent of the `aws_instance`. Because of this, Terraform will create the security group first so that it can then be associated with the EC2 instance.

However, the dependency relationship actually goes both directions causing the _Security Group Deletion Problem_. AWS does not allow you to delete the security group associated with another resource (_e.g._, the `aws_instance`).

Terraform does [not model bi-directional dependencies](https://developer.hashicorp.com/terraform/internals/graph) like this, but, even if it did, simply knowing the dependency situation would not be enough to solve it. For example, some resources must always have an associated security group while others don't need to. In addition, when the `aws_security_group` resource attempts to recreate, it receives a dependent object error, which does not provide information on whether the dependent object is a security group rule or, for example, an associated EC2 instance. Within Terraform, the associated resource (_e.g._, `aws_instance`) does not receive an error when the `aws_security_group` is trying to recreate even though that is where changes to the associated resource would need to take place (_e.g._, removing the security group association).

Despite these sticky problems, below are some ways to improve your experience when you find it necessary to recreate a security group.

#### `create_before_destroy`

(This example is one approach to [recreating security groups](#recreating-a-security-group). For more information on the challenges and the _Security Group Deletion Problem_, see [the section above](#recreating-a-security-group).)

Normally, Terraform first deletes the existing security group resource and then creates a new one. When a security group is associated with a resource, the delete won't succeed. You can invert the default behavior using the [`create_before_destroy` meta argument](https://www.terraform.io/language/meta-arguments/lifecycle#create_before_destroy):

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from cdktf import TerraformResourceLifecycle
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.security_group import SecurityGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SecurityGroup(self, "example",
            lifecycle=TerraformResourceLifecycle(
                create_before_destroy=True
            ),
            name="changeable-name"
        )
```

#### `replace_triggered_by`

(This example is one approach to [recreating security groups](#recreating-a-security-group). For more information on the challenges and the _Security Group Deletion Problem_, see [the section above](#recreating-a-security-group).)

To replace a resource when a security group changes, use the [`replace_triggered_by` meta argument](https://www.terraform.io/language/meta-arguments/lifecycle#replace_triggered_by). Note that in this example, the `aws_instance` will be destroyed and created again when the `aws_security_group` changes.

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from cdktf import TerraformResourceLifecycle
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.instance import Instance
from imports.aws.security_group import SecurityGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = SecurityGroup(self, "example",
            name="sg"
        )
        aws_instance_example = Instance(self, "example_1",
            instance_type="t3.small",
            lifecycle=TerraformResourceLifecycle(
                replace_triggered_by=[example]
            ),
            vpc_security_group_ids=[test.id]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_instance_example.override_logical_id("example")
```

#### Shorter timeout

(This example is one approach to [recreating security groups](#recreating-a-security-group). For more information on the challenges and the _Security Group Deletion Problem_, see [the section above](#recreating-a-security-group).)

If destroying a security group takes a long time, it may be because Terraform cannot distinguish between a dependent object (_e.g._, a security group rule or EC2 instance) that is _in the process of being deleted_ and one that is not. In other words, it may be waiting for a train that isn't scheduled to arrive. To fail faster, shorten the `delete` [timeout](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) from the default timeout:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.security_group import SecurityGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SecurityGroup(self, "example",
            name="izizavle",
            timeouts=[{
                "delete": "2m"
            }
            ]
        )
```

#### Provisioners

(This example is one approach to [recreating security groups](#recreating-a-security-group). For more information on the challenges and the _Security Group Deletion Problem_, see [the section above](#recreating-a-security-group).)

**DISCLAIMER:** We **_HIGHLY_** recommend using one of the above approaches and _NOT_ using local provisioners. Provisioners, like the one shown below, should be considered a **last resort** since they are _not readable_, _require skills outside standard Terraform configuration_, are _error prone_ and _difficult to maintain_, are not compatible with cloud environments and upgrade tools, require AWS CLI installation, and are subject to AWS CLI and Terraform changes outside the AWS Provider.

```terraform
data "aws_security_group" "default" {
  name = "default"
  # ... other configuration ...
}

resource "aws_security_group" "example" {
  name = "sg"
  # ... other configuration ...

  # The downstream resource must have at least one SG attached, therefore we
  # attach the default SG of the VPC temporarily and remove it later on
  provisioner "local-exec" {
    when       = destroy
    command    = <<DOC
            ENDPOINT_ID=`aws ec2 describe-vpc-endpoints --filters "Name=tag:Name,Values=${self.tags.workaround1}" --query "VpcEndpoints[0].VpcEndpointId" --output text` &&
            aws ec2 modify-vpc-endpoint --vpc-endpoint-id $${ENDPOINT_ID} --add-security-group-ids ${self.tags.workaround2} --remove-security-group-ids ${self.id}
        DOC
    on_failure = continue
  }

  tags = {
    workaround1 = "tagged-name"
    workaround2 = data.aws_security_group.default.id
  }
}

resource "null_resource" "example" {
  provisioner "local-exec" {
    command    = <<DOC
            aws ec2 modify-vpc-endpoint --vpc-endpoint-id ${aws_vpc_endpoint.example.id} --remove-security-group-ids ${data.aws_security_group.default.id}
        DOC
    on_failure = continue
  }

  triggers = {
    rerun_upon_change_of = join(",", aws_vpc_endpoint.example.security_group_ids)
  }
}
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `description` - (Optional, Forces new resource) Security group description. Defaults to `Managed by Terraform`. Cannot be `""`. **NOTE**: This field maps to the AWS `GroupDescription` attribute, for which there is no Update API. If you'd like to classify your security groups in a way that can be updated, use `tags`.
* `egress` - (Optional, VPC only) Configuration block for egress rules. Can be specified multiple times for each egress rule. Each egress block supports fields documented below. This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
* `ingress` - (Optional) Configuration block for ingress rules. Can be specified multiple times for each ingress rule. Each ingress block supports fields documented below. This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
* `name_prefix` - (Optional, Forces new resource) Creates a unique name beginning with the specified prefix. Conflicts with `name`.
* `name` - (Optional, Forces new resource) Name of the security group. If omitted, Terraform will assign a random, unique name.
* `revoke_rules_on_delete` - (Optional) Instruct Terraform to revoke all of the Security Groups attached ingress and egress rules before deleting the rule itself. This is normally not needed, however certain AWS services such as Elastic Map Reduce may automatically add required rules to security groups used with the service, and those rules may contain a cyclic dependency that prevent the security groups from being destroyed without removing the dependency first. Default `false`.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `vpc_id` - (Optional, Forces new resource) VPC ID. Defaults to the region's default VPC.

### ingress

This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).

The following arguments are required:

* `from_port` - (Required) Start port (or ICMP type number if protocol is `icmp` or `icmpv6`).
* `to_port` - (Required) End range port (or ICMP code if protocol is `icmp`).
* `protocol` - (Required) Protocol. If you select a protocol of `-1` (semantically equivalent to `all`, which is not a valid value here), you must specify a `from_port` and `to_port` equal to 0. The supported values are defined in the `IpProtocol` argument on the [IpPermission](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IpPermission.html) API reference. This argument is normalized to a lowercase value to match the AWS API requirement when using with Terraform 0.12.x and above, please make sure that the value of the protocol is specified as lowercase when using with older version of Terraform to avoid an issue during upgrade.

The following arguments are optional:

~> **Note** Although `cidr_blocks`, `ipv6_cidr_blocks`, `prefix_list_ids`, and `security_groups` are all marked as optional, you _must_ provide one of them in order to configure the source of the traffic.

* `cidr_blocks` - (Optional) List of CIDR blocks.
* `description` - (Optional) Description of this ingress rule.
* `ipv6_cidr_blocks` - (Optional) List of IPv6 CIDR blocks.
* `prefix_list_ids` - (Optional) List of Prefix List IDs.
* `security_groups` - (Optional) List of security groups. A group name can be used relative to the default VPC. Otherwise, group ID.
* `self` - (Optional) Whether the security group itself will be added as a source to this ingress rule.

### egress

This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).

The following arguments are required:

* `from_port` - (Required) Start port (or ICMP type number if protocol is `icmp`)
* `to_port` - (Required) End range port (or ICMP code if protocol is `icmp`).

The following arguments are optional:

~> **Note** Although `cidr_blocks`, `ipv6_cidr_blocks`, `prefix_list_ids`, and `security_groups` are all marked as optional, you _must_ provide one of them in order to configure the destination of the traffic.

* `cidr_blocks` - (Optional) List of CIDR blocks.
* `description` - (Optional) Description of this egress rule.
* `ipv6_cidr_blocks` - (Optional) List of IPv6 CIDR blocks.
* `prefix_list_ids` - (Optional) List of Prefix List IDs.
* `protocol` - (Required) Protocol. If you select a protocol of `-1` (semantically equivalent to `all`, which is not a valid value here), you must specify a `from_port` and `to_port` equal to 0. The supported values are defined in the `IpProtocol` argument in the [IpPermission](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IpPermission.html) API reference. This argument is normalized to a lowercase value to match the AWS API requirement when using Terraform 0.12.x and above. Please make sure that the value of the protocol is specified as lowercase when used with older version of Terraform to avoid issues during upgrade.
* `security_groups` - (Optional) List of security groups. A group name can be used relative to the default VPC. Otherwise, group ID.
* `self` - (Optional) Whether the security group itself will be added as a source to this egress rule.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the security group.
* `id` - ID of the security group.
* `owner_id` - Owner ID.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `10m`)
- `delete` - (Default `15m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Security Groups using the security group `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.security_group import SecurityGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SecurityGroup.generate_config_for_import(self, "elbSg", "sg-903004f8")
```

Using `terraform import`, import Security Groups using the security group `id`. For example:

```console
% terraform import aws_security_group.elb_sg sg-903004f8
```

<!-- cache-key: cdktf-0.20.8 input-30b7664a66000ed494f6d87b6ac7d76cbb048ba1cd60f878dd231a8b46a8de1f -->