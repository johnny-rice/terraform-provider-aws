---
subcategory: "Kendra"
layout: "aws"
page_title: "AWS: aws_kendra_experience"
description: |-
  Terraform resource for managing an AWS Kendra Experience.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_kendra_experience

Terraform resource for managing an AWS Kendra Experience.

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
from imports.aws.kendra_experience import KendraExperience
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        KendraExperience(self, "example",
            configuration=KendraExperienceConfiguration(
                content_source_configuration=KendraExperienceConfigurationContentSourceConfiguration(
                    direct_put_content=True,
                    faq_ids=[Token.as_string(aws_kendra_faq_example.faq_id)]
                ),
                user_identity_configuration=KendraExperienceConfigurationUserIdentityConfiguration(
                    identity_attribute_name="12345ec453-1546651e-79c4-4554-91fa-00b43ccfa245"
                )
            ),
            description="My Kendra Experience",
            index_id=Token.as_string(aws_kendra_index_example.id),
            name="example",
            role_arn=Token.as_string(aws_iam_role_example.arn)
        )
```

## Argument Reference

The following arguments are required:

* `index_id` - (Required, Forces new resource) The identifier of the index for your Amazon Kendra experience.
* `name` - (Required) A name for your Amazon Kendra experience.
* `role_arn` - (Required) The Amazon Resource Name (ARN) of a role with permission to access `Query API`, `QuerySuggestions API`, `SubmitFeedback API`, and `AWS SSO` that stores your user and group information. For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `description` - (Optional, Forces new resource if removed) A description for your Amazon Kendra experience.
* `configuration` - (Optional) Configuration information for your Amazon Kendra experience. Terraform will only perform drift detection of its value when present in a configuration. [Detailed below](#configuration).

~> **NOTE:** By default of the AWS Kendra API, updates to an existing `aws_kendra_experience` resource (e.g. updating the `name`) will also update the `configuration.content_source_configuration.direct_put_content` parameter to `false` if not already provided.

### `configuration`

~> **NOTE:** By default of the AWS Kendra API, the `content_source_configuration.direct_put_content` parameter will be set to `false` if not provided.  

The `configuration` configuration block supports the following arguments:

* `content_source_configuration` - (Optional, Required if `user_identity_configuration` not provided) The identifiers of your data sources and FAQs. Or, you can specify that you want to use documents indexed via the `BatchPutDocument API`. Terraform will only perform drift detection of its value when present in a configuration. [Detailed below](#content_source_configuration).
* `user_identity_configuration` - (Optional, Required if `content_source_configuration` not provided) The AWS SSO field name that contains the identifiers of your users, such as their emails. [Detailed below](#user_identity_configuration).

### `content_source_configuration`

The `content_source_configuration` configuration block supports the following arguments:

* `data_source_ids` - (Optional) The identifiers of the data sources you want to use for your Amazon Kendra experience. Maximum number of 100 items.
* `direct_put_content` - (Optional) Whether to use documents you indexed directly using the `BatchPutDocument API`. Defaults to `false`.
* `faq_ids` - (Optional) The identifier of the FAQs that you want to use for your Amazon Kendra experience. Maximum number of 100 items.

### `user_identity_configuration`

The `user_identity_configuration` configuration block supports the following argument:

* `identity_attribute_name` - (Required) The AWS SSO field name that contains the identifiers of your users, such as their emails.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The unique identifiers of the experience and index separated by a slash (`/`).
* `arn` - ARN of the Experience.
* `endpoints` - Shows the endpoint URLs for your Amazon Kendra experiences. The URLs are unique and fully hosted by AWS.
    * `endpoint` - The endpoint of your Amazon Kendra experience.
    * `endpoint_type` - The type of endpoint for your Amazon Kendra experience.
* `experience_id` - The unique identifier of the experience.
* `status` - The current processing status of your Amazon Kendra experience.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Kendra Experience using the unique identifiers of the experience and index separated by a slash (`/`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.kendra_experience import KendraExperience
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        KendraExperience.generate_config_for_import(self, "example", "1045d08d-66ef-4882-b3ed-dfb7df183e90/b34dfdf7-1f2b-4704-9581-79e00296845f")
```

Using `terraform import`, import Kendra Experience using the unique identifiers of the experience and index separated by a slash (`/`). For example:

```console
% terraform import aws_kendra_experience.example 1045d08d-66ef-4882-b3ed-dfb7df183e90/b34dfdf7-1f2b-4704-9581-79e00296845f
```

<!-- cache-key: cdktf-0.20.8 input-84518e9e182ca7bc4cf838bb011fd91dd8377b6c0776cd49eec16f6e5099e127 -->