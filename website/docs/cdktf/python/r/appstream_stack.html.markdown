---
subcategory: "AppStream 2.0"
layout: "aws"
page_title: "AWS: aws_appstream_stack"
description: |-
  Provides an AppStream stack
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_appstream_stack

Provides an AppStream stack.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appstream_stack import AppstreamStack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppstreamStack(self, "example",
            application_settings=AppstreamStackApplicationSettings(
                enabled=True,
                settings_group="SettingsGroup"
            ),
            description="stack description",
            display_name="stack display name",
            feedback_url="http://your-domain/feedback",
            name="stack name",
            redirect_url="http://your-domain/redirect",
            storage_connectors=[AppstreamStackStorageConnectors(
                connector_type="HOMEFOLDERS"
            )
            ],
            tags={
                "TagName": "TagValue"
            },
            user_settings=[AppstreamStackUserSettings(
                action="AUTO_TIME_ZONE_REDIRECTION",
                permission="DISABLED"
            ), AppstreamStackUserSettings(
                action="CLIPBOARD_COPY_FROM_LOCAL_DEVICE",
                permission="ENABLED"
            ), AppstreamStackUserSettings(
                action="CLIPBOARD_COPY_TO_LOCAL_DEVICE",
                permission="ENABLED"
            ), AppstreamStackUserSettings(
                action="DOMAIN_PASSWORD_SIGNIN",
                permission="ENABLED"
            ), AppstreamStackUserSettings(
                action="DOMAIN_SMART_CARD_SIGNIN",
                permission="DISABLED"
            ), AppstreamStackUserSettings(
                action="FILE_DOWNLOAD",
                permission="ENABLED"
            ), AppstreamStackUserSettings(
                action="FILE_UPLOAD",
                permission="ENABLED"
            ), AppstreamStackUserSettings(
                action="PRINTING_TO_LOCAL_DEVICE",
                permission="ENABLED"
            )
            ]
        )
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Unique name for the AppStream stack.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `access_endpoints` - (Optional) Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
  See [`access_endpoints`](#access_endpoints) below.
* `application_settings` - (Optional) Settings for application settings persistence.
  See [`application_settings`](#application_settings) below.
* `description` - (Optional) Description for the AppStream stack.
* `display_name` - (Optional) Stack name to display.
* `embed_host_domains` - (Optional) Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
* `feedback_url` - (Optional) URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
* `redirect_url` - (Optional) URL that users are redirected to after their streaming session ends.
* `storage_connectors` - (Optional) Configuration block for the storage connectors to enable.
  See [`storage_connectors`](#storage_connectors) below.
* `user_settings` - (Optional) Configuration block for the actions that are enabled or disabled for users during their streaming sessions. If not provided, these settings are configured automatically by AWS. If provided, the terraform configuration should include a block for each configurable action.
  See [`user_settings`](#user_settings) below.
* `streaming_experience_settings` - (Optional) The streaming protocol you want your stack to prefer. This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.
  See [`streaming_experience_settings`](#streaming_experience_settings) below.
* `tags` - (Optional) Key-value mapping of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### `access_endpoints`

* `endpoint_type` - (Required) Type of the interface endpoint.
  See the [`AccessEndpoint` AWS API documentation](https://docs.aws.amazon.com/appstream2/latest/APIReference/API_AccessEndpoint.html) for valid values.
* `vpce_id` - (Optional) ID of the VPC in which the interface endpoint is used.

### `application_settings`

* `enabled` - (Required) Whether application settings should be persisted.
* `settings_group` - (Optional) Name of the settings group.
  Required when `enabled` is `true`.
  Can be up to 100 characters.

### `storage_connectors`

* `connector_type` - (Required) Type of storage connector.
  Valid values are `HOMEFOLDERS`, `GOOGLE_DRIVE`, or `ONE_DRIVE`.
* `domains` - (Optional) Names of the domains for the account.
* `resource_identifier` - (Optional) ARN of the storage connector.

### `user_settings`

* `action` - (Required) Action that is enabled or disabled.
  Valid values are `AUTO_TIME_ZONE_REDIRECTION`, `CLIPBOARD_COPY_FROM_LOCAL_DEVICE`, `CLIPBOARD_COPY_TO_LOCAL_DEVICE`, `DOMAIN_PASSWORD_SIGNIN`, `DOMAIN_SMART_CARD_SIGNIN`, `FILE_UPLOAD`, `FILE_DOWNLOAD`, or `PRINTING_TO_LOCAL_DEVICE`.
* `permission` - (Required) Whether the action is enabled or disabled.
  Valid values are `ENABLED` or `DISABLED`.

### `streaming_experience_settings`

* `preferred_protocol` - (Optional) The preferred protocol that you want to use while streaming your application.
  Valid values are `TCP` and `UDP`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the appstream stack.
* `created_time` - Date and time, in UTC and extended RFC 3339 format, when the stack was created.
* `id` - Unique ID of the appstream stack.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_appstream_stack` using the id. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appstream_stack import AppstreamStack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppstreamStack.generate_config_for_import(self, "example", "stackID")
```

Using `terraform import`, import `aws_appstream_stack` using the id. For example:

```console
% terraform import aws_appstream_stack.example stackID
```

<!-- cache-key: cdktf-0.20.8 input-cb2bdf6e130f23e2b7a12e4882a4ebdc1202f0e431368e87980d99b9785db0c0 -->