---
subcategory: "SSO Admin"
layout: "aws"
page_title: "AWS: aws_ssoadmin_application_providers"
description: |-
  Terraform data source for managing AWS SSO Admin Application Providers.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ssoadmin_application_providers

Terraform data source for managing AWS SSO Admin Application Providers.

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
from imports.aws.data_aws_ssoadmin_application_providers import DataAwsSsoadminApplicationProviders
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsSsoadminApplicationProviders(self, "example")
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - AWS region.
* `application_providers` - A list of application providers available in the current region. See [`application_providers`](#application_providers-attribute-reference) below.

### `application_providers` Attribute Reference

* `application_provider_arn` - ARN of the application provider.
* `display_data` - An object describing how IAM Identity Center represents the application provider in the portal. See [`display_data`](#display_data-attribute-reference) below.
* `federation_protocol` - Protocol that the application provider uses to perform federation. Valid values are `SAML` and `OAUTH`.

### `display_data` Attribute Reference

* `description` - Description of the application provider.
* `display_name` - Name of the application provider.
* `icon_url` - URL that points to an icon that represents the application provider.

<!-- cache-key: cdktf-0.20.8 input-a45aef2dc2fdc30ead19aab394d3ae5888b4a5e913575dcc95cb03e9c3d59e80 -->