---
subcategory: "Timestream Write"
layout: "aws"
page_title: "AWS: aws_timestreamwrite_database"
description: |-
  Provides a Timestream database resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_timestreamwrite_database

Provides a Timestream database resource.

## Example Usage

### Basic usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.timestreamwrite_database import TimestreamwriteDatabase
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        TimestreamwriteDatabase(self, "example",
            database_name="database-example"
        )
```

### Full usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.timestreamwrite_database import TimestreamwriteDatabase
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        TimestreamwriteDatabase(self, "example",
            database_name="database-example",
            kms_key_id=Token.as_string(aws_kms_key_example.arn),
            tags={
                "Name": "value"
            }
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `database_name` - (Required) The name of the Timestream database. Minimum length of 3. Maximum length of 64.
* `kms_key_id` - (Optional) The ARN (not Alias ARN) of the KMS key to be used to encrypt the data stored in the database. If the KMS key is not specified, the database will be encrypted with a Timestream managed KMS key located in your account. Refer to [AWS managed KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk) for more info.
* `tags` - (Optional) Map of tags to assign to this resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The name of the Timestream database.
* `arn` - The ARN that uniquely identifies this database.
* `kms_key_id` - The ARN of the KMS key used to encrypt the data stored in the database.
* `table_count` - The total number of tables found within the Timestream database.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Timestream databases using the `database_name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.timestreamwrite_database import TimestreamwriteDatabase
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        TimestreamwriteDatabase.generate_config_for_import(self, "example", "example")
```

Using `terraform import`, import Timestream databases using the `database_name`. For example:

```console
% terraform import aws_timestreamwrite_database.example example
```

<!-- cache-key: cdktf-0.20.8 input-0f9f6306b946c419825721a60d1ab5087901270b76af3550886c32e9dbb54bc6 -->