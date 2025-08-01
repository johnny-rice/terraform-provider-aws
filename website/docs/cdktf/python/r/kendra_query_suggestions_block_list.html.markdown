---
subcategory: "Kendra"
layout: "aws"
page_title: "AWS: aws_kendra_query_suggestions_block_list"
description: |-
  Terraform resource for managing an AWS Kendra block list used for query suggestions for an index
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_kendra_query_suggestions_block_list

Use the `aws_kendra_index_block_list` resource to manage an AWS Kendra block list used for query suggestions for an index.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.kendra_query_suggestions_block_list import KendraQuerySuggestionsBlockList
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        KendraQuerySuggestionsBlockList(self, "example",
            index_id=Token.as_string(aws_kendra_index_example.id),
            name="Example",
            role_arn=Token.as_string(aws_iam_role_example.arn),
            source_s3_path=KendraQuerySuggestionsBlockListSourceS3Path(
                bucket=Token.as_string(aws_s3_bucket_example.id),
                key="example/suggestions.txt"
            ),
            tags={
                "Name": "Example Kendra Index"
            }
        )
```

## Argument Reference

The following arguments are required:

* `index_id` - (Required, Forces New Resource) Identifier of the index for a block list.
* `name` - (Required) Name for the block list.
* `role_arn` - (Required) IAM (Identity and Access Management) role used to access the block list text file in S3.
* `source_s3_path` - (Required) S3 path where your block list text file is located. See details below.

The `source_s3_path` configuration block supports the following arguments:

* `bucket` - (Required) Name of the S3 bucket that contains the file.
* `key` - (Required) Name of the file.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `description` - (Optional) Description for a block list.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block), tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the block list.
* `query_suggestions_block_list_id` - Unique identifier of the block list.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider's [default_tags configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

Configuration options for operation timeouts can be found [here](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts).

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `aws_kendra_query_suggestions_block_list` resource using the unique identifiers of the block list and index separated by a slash (`/`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.kendra_query_suggestions_block_list import KendraQuerySuggestionsBlockList
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        KendraQuerySuggestionsBlockList.generate_config_for_import(self, "example", "blocklist-123456780/idx-8012925589")
```

Using `terraform import`, import the `aws_kendra_query_suggestions_block_list` resource using the unique identifiers of the block list and index separated by a slash (`/`). For example:

```console
% terraform import aws_kendra_query_suggestions_block_list.example blocklist-123456780/idx-8012925589
```

<!-- cache-key: cdktf-0.20.8 input-c3f0258d125ae302dc1cb1de441abe7a58be01006304bedbd844661b54c0661d -->