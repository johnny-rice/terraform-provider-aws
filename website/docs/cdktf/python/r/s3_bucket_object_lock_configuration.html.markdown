---
subcategory: "S3 (Simple Storage)"
layout: "aws"
page_title: "AWS: aws_s3_bucket_object_lock_configuration"
description: |-
  Provides an S3 bucket Object Lock configuration resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_s3_bucket_object_lock_configuration

Provides an S3 bucket Object Lock configuration resource. For more information about Object Locking, go to [Using S3 Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock.html) in the Amazon S3 User Guide.

-> This resource can be used enable Object Lock for **new** and **existing** buckets.

-> This resource cannot be used with S3 directory buckets.

## Example Usage

### Object Lock configuration for new or existing buckets

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.s3_bucket import S3Bucket
from imports.aws.s3_bucket_object_lock_configuration import S3BucketObjectLockConfigurationA
from imports.aws.s3_bucket_versioning import S3BucketVersioningA
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = S3Bucket(self, "example",
            bucket="mybucket"
        )
        aws_s3_bucket_object_lock_configuration_example =
        S3BucketObjectLockConfigurationA(self, "example_1",
            bucket=example.id,
            rule=S3BucketObjectLockConfigurationRuleA(
                default_retention=S3BucketObjectLockConfigurationRuleDefaultRetentionA(
                    days=5,
                    mode="COMPLIANCE"
                )
            )
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_s3_bucket_object_lock_configuration_example.override_logical_id("example")
        aws_s3_bucket_versioning_example = S3BucketVersioningA(self, "example_2",
            bucket=example.id,
            versioning_configuration=S3BucketVersioningVersioningConfiguration(
                status="Enabled"
            )
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_s3_bucket_versioning_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `bucket` - (Required, Forces new resource) Name of the bucket.
* `expected_bucket_owner` - (Optional, Forces new resource) Account ID of the expected bucket owner.
* `object_lock_enabled` - (Optional, Forces new resource) Indicates whether this bucket has an Object Lock configuration enabled. Defaults to `Enabled`. Valid values: `Enabled`.
* `rule` - (Optional) Configuration block for specifying the Object Lock rule for the specified object. [See below](#rule).
* `token` - (Optional,Deprecated) This argument is deprecated and no longer needed to enable Object Lock.
To enable Object Lock for an existing bucket, you must first enable versioning on the bucket and then enable Object Lock. For more details on versioning, see the [`aws_s3_bucket_versioning` resource](s3_bucket_versioning.html.markdown).

### rule

The `rule` configuration block supports the following arguments:

* `default_retention` - (Required) Configuration block for specifying the default Object Lock retention settings for new objects placed in the specified bucket. [See below](#default_retention).

### default_retention

The `default_retention` configuration block supports the following arguments:

* `days` - (Optional, Required if `years` is not specified) Number of days that you want to specify for the default retention period.
* `mode` - (Required) Default Object Lock retention mode you want to apply to new objects placed in the specified bucket. Valid values: `COMPLIANCE`, `GOVERNANCE`.
* `years` - (Optional, Required if `days` is not specified) Number of years that you want to specify for the default retention period.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The `bucket` or `bucket` and `expected_bucket_owner` separated by a comma (`,`) if the latter is provided.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import an S3 bucket Object Lock Configuration using one of two forms. If the owner (account ID) of the source bucket is the same account used to configure the Terraform AWS Provider, import using the `bucket`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.s3_bucket_object_lock_configuration import S3BucketObjectLockConfigurationA
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        S3BucketObjectLockConfigurationA.generate_config_for_import(self, "example", "bucket-name")
```

If the owner (account ID) of the source bucket differs from the account used to configure the Terraform AWS Provider, import using the `bucket` and `expected_bucket_owner`, separated by a comma (`,`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.s3_bucket_object_lock_configuration import S3BucketObjectLockConfigurationA
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        S3BucketObjectLockConfigurationA.generate_config_for_import(self, "example", "bucket-name,123456789012")
```

**Using `terraform import`**, import an S3 bucket Object Lock Configuration using one of two forms. If the owner (account ID) of the source bucket is the same account used to configure the Terraform AWS Provider, import using the `bucket`. For example:

```console
% terraform import aws_s3_bucket_object_lock_configuration.example bucket-name
```

If the owner (account ID) of the source bucket differs from the account used to configure the Terraform AWS Provider, import using the `bucket` and `expected_bucket_owner`, separated by a comma (`,`). For example:

```console
% terraform import aws_s3_bucket_object_lock_configuration.example bucket-name,123456789012
```

<!-- cache-key: cdktf-0.20.8 input-e355b18418eb80788255f2bd42c221a7ba48a5d3d66f655f3641bc428392b9f8 -->