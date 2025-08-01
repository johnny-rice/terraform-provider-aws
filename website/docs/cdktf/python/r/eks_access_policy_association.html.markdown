---
subcategory: "EKS (Elastic Kubernetes)"
layout: "aws"
page_title: "AWS: aws_eks_access_policy_associattion"
description: |-
  Access Entry Policy Association for an EKS Cluster.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_eks_access_policy_association

Access Entry Policy Association for an EKS Cluster.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.eks_access_policy_association import EksAccessPolicyAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        EksAccessPolicyAssociation(self, "example",
            access_scope=EksAccessPolicyAssociationAccessScope(
                namespaces=["example-namespace"],
                type="namespace"
            ),
            cluster_name=Token.as_string(aws_eks_cluster_example.name),
            policy_arn="arn:aws:eks::aws:cluster-access-policy/AmazonEKSViewPolicy",
            principal_arn=Token.as_string(aws_iam_user_example.arn)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `cluster_name` - (Required) Name of the EKS Cluster.
* `policy_arn` - (Required) The ARN of the access policy that you're associating.
* `principal_arn` - (Required) The IAM Principal ARN which requires Authentication access to the EKS cluster.
* `access_scope` - (Required) The configuration block to determine the scope of the access. See [`access_scope` Block](#access_scope-block) below.

### `access_scope` Block

The `access_scope` block supports the following arguments.

* `type` - (Required) Valid values are `namespace` or `cluster`.
* `namespaces` - (Optional) The namespaces to which the access scope applies when type is namespace.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `associated_access_policy` - Contains information about the access policy associatieon. See [`associated_access_policy` Block](#associated_access_policy-block) below.

### `associated_access_policy` Block

The `associated_access_policy` block has the following attributes.

* `associated_at` - Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the policy was associated.
* `modified_at` - Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the policy was updated.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `20m`)
* `update` - (Default `20m`)
* `delete` - (Default `40m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import EKS add-on using the `cluster_name`, `principal_arn`and `policy_arn` separated by an octothorp (`#`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.eks_access_policy_association import EksAccessPolicyAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        EksAccessPolicyAssociation.generate_config_for_import(self, "myEksEntry", "my_cluster_name#my_principal_arn#my_policy_arn")
```

Using `terraform import`, import EKS access entry using the `cluster_name` `principal_arn` and `policy_arn` separated by an octothorp (`#`). For example:

```console
% terraform import aws_eks_access_policy_association.my_eks_access_entry my_cluster_name#my_principal_arn#my_policy_arn
```

<!-- cache-key: cdktf-0.20.8 input-d2b4280cd293788b0f38d656ad4c7aeddfb7b78b462d6f75a40a3f1f4602dd69 -->