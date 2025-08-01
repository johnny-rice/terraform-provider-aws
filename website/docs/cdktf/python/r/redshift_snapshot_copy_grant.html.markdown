---
subcategory: "Redshift"
layout: "aws"
page_title: "AWS: aws_redshift_snapshot_copy_grant"
description: |-
  Creates a snapshot copy grant that allows AWS Redshift to encrypt copied snapshots with a customer master key from AWS KMS in a destination region.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_redshift_snapshot_copy_grant

Creates a snapshot copy grant that allows AWS Redshift to encrypt copied snapshots with a customer master key from AWS KMS in a destination region.

Note that the grant must exist in the destination region, and not in the region of the cluster.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.redshift_cluster import RedshiftCluster
from imports.aws.redshift_snapshot_copy_grant import RedshiftSnapshotCopyGrant
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, clusterIdentifier, nodeType):
        super().__init__(scope, name)
        test = RedshiftSnapshotCopyGrant(self, "test",
            snapshot_copy_grant_name="my-grant"
        )
        aws_redshift_cluster_test = RedshiftCluster(self, "test_1",
            snapshot_copy=[{
                "destination_region": "us-east-2",
                "grant_name": test.snapshot_copy_grant_name
            }
            ],
            cluster_identifier=cluster_identifier,
            node_type=node_type
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_redshift_cluster_test.override_logical_id("test")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `snapshot_copy_grant_name` - (Required, Forces new resource) A friendly name for identifying the grant.
* `kms_key_id` - (Optional, Forces new resource) The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN. If not specified, the default key is used.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of snapshot copy grant
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Redshift Snapshot Copy Grants by name. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.redshift_snapshot_copy_grant import RedshiftSnapshotCopyGrant
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        RedshiftSnapshotCopyGrant.generate_config_for_import(self, "test", "my-grant")
```

Using `terraform import`, import Redshift Snapshot Copy Grants by name. For example:

```console
% terraform import aws_redshift_snapshot_copy_grant.test my-grant
```

<!-- cache-key: cdktf-0.20.8 input-0ca88ea1a32ffa9dd3d9cd934727bd97e2768aeecd7c529b79f393913a9df9ab -->