---
subcategory: "DataSync"
layout: "aws"
page_title: "AWS: aws_datasync_task"
description: |-
  Manages an AWS DataSync Task
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_datasync_task

Manages an AWS DataSync Task, which represents a configuration for synchronization. Starting an execution of these DataSync Tasks (actually synchronizing files) is performed outside of this Terraform resource.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Op, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.datasync_task import DatasyncTask
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DatasyncTask(self, "example",
            destination_location_arn=destination.arn,
            name="example",
            options=DatasyncTaskOptions(
                bytes_per_second=Token.as_number(Op.negate(1))
            ),
            source_location_arn=source.arn
        )
```

## Example Usage with Scheduling

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.datasync_task import DatasyncTask
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DatasyncTask(self, "example",
            destination_location_arn=destination.arn,
            name="example",
            schedule=DatasyncTaskSchedule(
                schedule_expression="cron(0 12 ? * SUN,WED *)"
            ),
            source_location_arn=source.arn
        )
```

## Example Usage with Filtering

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.datasync_task import DatasyncTask
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DatasyncTask(self, "example",
            destination_location_arn=destination.arn,
            excludes=DatasyncTaskExcludes(
                filter_type="SIMPLE_PATTERN",
                value="/folder1|/folder2"
            ),
            includes=DatasyncTaskIncludes(
                filter_type="SIMPLE_PATTERN",
                value="/folder1|/folder2"
            ),
            name="example",
            source_location_arn=source.arn
        )
```

## Example Usage with Enhanced Task Mode

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.datasync_task import DatasyncTask
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DatasyncTask(self, "example",
            destination_location_arn=destination.arn,
            name="example",
            options=DatasyncTaskOptions(
                gid="NONE",
                posix_permissions="NONE",
                uid="NONE",
                verify_mode="ONLY_FILES_TRANSFERRED"
            ),
            source_location_arn=source.arn,
            task_mode="ENHANCED"
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `destination_location_arn` - (Required) Amazon Resource Name (ARN) of destination DataSync Location.
* `source_location_arn` - (Required) Amazon Resource Name (ARN) of source DataSync Location.
* `cloudwatch_log_group_arn` - (Optional) Amazon Resource Name (ARN) of the CloudWatch Log Group that is used to monitor and log events in the sync task.
* `excludes` - (Optional) Filter rules that determines which files to exclude from a task.
* `includes` - (Optional) Filter rules that determines which files to include in a task.
* `name` - (Optional) Name of the DataSync Task.
* `options` - (Optional) Configuration block containing option that controls the default behavior when you start an execution of this DataSync Task. For each individual task execution, you can override these options by specifying an overriding configuration in those executions.
* `schedule` - (Optional) Specifies a schedule used to periodically transfer files from a source to a destination location.
* `tags` - (Optional) Key-value pairs of resource tags to assign to the DataSync Task. If configured with a provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `task_mode` - (Optional) One of the following task modes for your data transfer:
    * `BASIC` (default) - Transfer files or objects between Amazon Web Services storage and on-premises, edge, or other cloud storage.
    * `ENHANCED` - Transfer virtually unlimited numbers of objects with enhanced metrics, more detailed logs, and higher performance than Basic mode. Currently available for transfers between Amazon S3 locations.
* `task_report_config` - (Optional) Configuration block containing the configuration of a DataSync Task Report. See [`task_report_config`](#task_report_config-argument-reference) below.

### options Argument Reference

~> **NOTE:** If `atime` is set to `BEST_EFFORT`, `mtime` must be set to `PRESERVE`. If `atime` is set to `NONE`, `mtime` must be set to `NONE`.

The `options` configuration block supports the following arguments:

* `atime` - (Optional) A file metadata that shows the last time a file was accessed (that is when the file was read or written to). If set to `BEST_EFFORT`, the DataSync Task attempts to preserve the original (that is, the version before sync `PREPARING` phase) `atime` attribute on all source files. Valid values: `BEST_EFFORT`, `NONE`. Default: `BEST_EFFORT`.
* `bytes_per_second` - (Optional) Limits the bandwidth utilized. For example, to set a maximum of 1 MB, set this value to `1048576`. Value values: `-1` or greater. Default: `-1` (unlimited).
* `gid` - (Optional) Group identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
* `log_level` - (Optional) Determines the type of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide. Valid values: `OFF`, `BASIC`, `TRANSFER`. Default: `OFF`.
* `mtime` - (Optional) A file metadata that indicates the last time a file was modified (written to) before the sync `PREPARING` phase. Value values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
* `object_tags` - (Optional) Specifies whether object tags are maintained when transferring between object storage systems. If you want your DataSync task to ignore object tags, specify the NONE value. Valid values: `PRESERVE`, `NONE`. Default value: `PRESERVE`.
* `overwrite_mode` - (Optional) Determines whether files at the destination should be overwritten or preserved when copying files. Valid values: `ALWAYS`, `NEVER`. Default: `ALWAYS`.
* `posix_permissions` - (Optional) Determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file. Valid values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
* `preserve_deleted_files` - (Optional) Whether files deleted in the source should be removed or preserved in the destination file system. Valid values: `PRESERVE`, `REMOVE`. Default: `PRESERVE`.
* `preserve_devices` - (Optional) Whether the DataSync Task should preserve the metadata of block and character devices in the source files system, and recreate the files with that device name and metadata on the destination. The DataSync Task can’t sync the actual contents of such devices, because many of the devices are non-terminal and don’t return an end of file (EOF) marker. Valid values: `NONE`, `PRESERVE`. Default: `NONE` (ignore special devices).
* `security_descriptor_copy_flags` - (Optional) Determines which components of the SMB security descriptor are copied from source to destination objects. This value is only used for transfers between SMB and Amazon FSx for Windows File Server locations, or between two Amazon FSx for Windows File Server locations. Valid values: `NONE`, `OWNER_DACL`, `OWNER_DACL_SACL`. Default: `OWNER_DACL`.
* `task_queueing` - (Optional) Determines whether tasks should be queued before executing the tasks. Valid values: `ENABLED`, `DISABLED`. Default `ENABLED`.
* `transfer_mode` - (Optional) Determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location. Valid values: `CHANGED`, `ALL`. Default: `CHANGED`
* `uid` - (Optional) User identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
* `verify_mode` - (Optional) Whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred. Valid values: `NONE`, `POINT_IN_TIME_CONSISTENT`, `ONLY_FILES_TRANSFERRED`. Default: `POINT_IN_TIME_CONSISTENT`.

### `task_report_config` Argument Reference

The following arguments are supported inside the `task_report_config` configuration block:

* `s3_destination` - (Required) Configuration block containing the configuration for the Amazon S3 bucket where DataSync uploads your task report. See [`s3_destination`](#s3_destination-argument-reference) below.
* `s3_object_versioning` - (Optional) Specifies whether your task report includes the new version of each object transferred into an S3 bucket. This only applies if you enable versioning on your bucket. Keep in mind that setting this to INCLUDE can increase the duration of your task execution. Valid values: `INCLUDE` and `NONE`.
* `output_type` - (Optional) Specifies the type of task report you'd like. Valid values: `SUMMARY_ONLY` and `STANDARD`.
* `report_overrides` - (Optional) Configuration block containing the configuration of the reporting level for aspects of your task report. See [`report_overrides`](#report_overrides-argument-reference) below.
* `report_level` - Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't. Valid values: `ERRORS_ONLY` and `SUCCESSES_AND_ERRORS`.

### `s3_destination` Argument Reference

The following arguments are supported inside the `s3_destination` configuration block:

* `bucket_access_role_arn` - (Required) Specifies the Amazon Resource Name (ARN) of the IAM policy that allows DataSync to upload a task report to your S3 bucket.
* `s3_bucket_arn` - (Required) Specifies the ARN of the S3 bucket where DataSync uploads your report.
* `subdirectory` - (Optional) Specifies a bucket prefix for your report.

### `report_overrides` Argument Reference

The following arguments are supported inside the `report_overrides` configuration block:

* `deleted_override` - (Optional) Specifies the level of reporting for the files, objects, and directories that DataSync attempted to delete in your destination location. This only applies if you configure your task to delete data in the destination that isn't in the source. Valid values: `ERRORS_ONLY` and `SUCCESSES_AND_ERRORS`.
* `skipped_override` - (Optional) Specifies the level of reporting for the files, objects, and directories that DataSync attempted to skip during your transfer. Valid values: `ERRORS_ONLY` and `SUCCESSES_AND_ERRORS`.
* `transferred_override` - (Optional) Specifies the level of reporting for the files, objects, and directories that DataSync attempted to transfer. Valid values: `ERRORS_ONLY` and `SUCCESSES_AND_ERRORS`.
* `verified_override` - (Optional) Specifies the level of reporting for the files, objects, and directories that DataSync attempted to verify at the end of your transfer. Valid values: `ERRORS_ONLY` and `SUCCESSES_AND_ERRORS`.

~> **NOTE:** If any `report_overrides` are set to the same value as `task_report_config.report_level`, they will always be flagged as changed. Only set overrides to a value that differs from `task_report_config.report_level`.

### Schedule

* `schedule_expression` - (Required) Specifies the schedule you want your task to use for repeated executions. For more information, see [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).

### excludes Argument Reference

* `filter_type` - (Optional) The type of filter rule to apply. Valid values: `SIMPLE_PATTERN`.
* `value` - (Optional) A single filter string that consists of the patterns to exclude. The patterns are delimited by "|" (that is, a pipe), for example: `/folder1|/folder2`

### includes Argument Reference

* `filter_type` - (Optional) The type of filter rule to apply. Valid values: `SIMPLE_PATTERN`.
* `value` - (Optional) A single filter string that consists of the patterns to include. The patterns are delimited by "|" (that is, a pipe), for example: `/folder1|/folder2`

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Amazon Resource Name (ARN) of the DataSync Task.
* `arn` - Amazon Resource Name (ARN) of the DataSync Task.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `5m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_datasync_task` using the DataSync Task Amazon Resource Name (ARN). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.datasync_task import DatasyncTask
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DatasyncTask.generate_config_for_import(self, "example", "arn:aws:datasync:us-east-1:123456789012:task/task-12345678901234567")
```

Using `terraform import`, import `aws_datasync_task` using the DataSync Task Amazon Resource Name (ARN). For example:

```console
% terraform import aws_datasync_task.example arn:aws:datasync:us-east-1:123456789012:task/task-12345678901234567
```

<!-- cache-key: cdktf-0.20.8 input-e345ae4963038d5596f0040049ba40af169207efd7ca0c2a4edceda7b6f026d2 -->