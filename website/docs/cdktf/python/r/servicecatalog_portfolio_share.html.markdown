---
subcategory: "Service Catalog"
layout: "aws"
page_title: "AWS: aws_servicecatalog_portfolio_share"
description: |-
  Manages a Service Catalog Portfolio Share
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_servicecatalog_portfolio_share

Manages a Service Catalog Portfolio Share. Shares the specified portfolio with the specified account or organization node. You can share portfolios to an organization, an organizational unit, or a specific account.

If the portfolio share with the specified account or organization node already exists, using this resource to re-create the share will have no effect and will not return an error. You can then use this resource to update the share.

~> **NOTE:** Shares to an organization node can only be created by the management account of an organization or by a delegated administrator. If a delegated admin is de-registered, they can no longer create portfolio shares.

~> **NOTE:** AWSOrganizationsAccess must be enabled in order to create a portfolio share to an organization node.

~> **NOTE:** You can't share a shared resource, including portfolios that contain a shared product.

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
from imports.aws.servicecatalog_portfolio_share import ServicecatalogPortfolioShare
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ServicecatalogPortfolioShare(self, "example",
            portfolio_id=Token.as_string(aws_servicecatalog_portfolio_example.id),
            principal_id="012128675309",
            type="ACCOUNT"
        )
```

## Argument Reference

The following arguments are required:

* `portfolio_id` - (Required) Portfolio identifier.
* `principal_id` - (Required) Identifier of the principal with whom you will share the portfolio. Valid values AWS account IDs and ARNs of AWS Organizations and organizational units.
* `type` - (Required) Type of portfolio share. Valid values are `ACCOUNT` (an external account), `ORGANIZATION` (a share to every account in an organization), `ORGANIZATIONAL_UNIT`, `ORGANIZATION_MEMBER_ACCOUNT` (a share to an account in an organization).

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `accept_language` - (Optional) Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
* `share_principals` - (Optional) Enables or disables Principal sharing when creating the portfolio share. If this flag is not provided, principal sharing is disabled.
* `share_tag_options` - (Optional) Whether to enable sharing of `aws_servicecatalog_tag_option` resources when creating the portfolio share.
* `wait_for_acceptance` - (Optional) Whether to wait (up to the timeout) for the share to be accepted. Organizational shares are automatically accepted.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `accepted` - Whether the shared portfolio is imported by the recipient account. If the recipient is organizational, the share is automatically imported, and the field is always set to true.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `3m`)
- `read` - (Default `10m`)
- `update` - (Default `3m`)
- `delete` - (Default `3m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_servicecatalog_portfolio_share` using the portfolio share ID. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.servicecatalog_portfolio_share import ServicecatalogPortfolioShare
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ServicecatalogPortfolioShare.generate_config_for_import(self, "example", "port-12344321:ACCOUNT:123456789012")
```

Using `terraform import`, import `aws_servicecatalog_portfolio_share` using the portfolio share ID. For example:

```console
% terraform import aws_servicecatalog_portfolio_share.example port-12344321:ACCOUNT:123456789012
```

<!-- cache-key: cdktf-0.20.8 input-4cbf119057df42a6a2e1b0ab1f227e2cb35a63f0921dff43497dc1cc9842b138 -->