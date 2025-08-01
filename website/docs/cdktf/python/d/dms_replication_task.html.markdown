---
subcategory: "DMS (Database Migration)"
layout: "aws"
page_title: "AWS: aws_dms_replication_task"
description: |-
  Terraform data source for managing an AWS DMS (Database Migration) Replication Task.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_dms_replication_task

Terraform data source for managing an AWS DMS (Database Migration) Replication Task.

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
from imports.aws.data_aws_dms_replication_task import DataAwsDmsReplicationTask
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsDmsReplicationTask(self, "test",
            replication_task_id=Token.as_string(aws_dms_replication_task_test.replication_task_id)
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `replication_task_id` - (Required) The replication task identifier.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `cdc_start_position` - (Conflicts with `cdc_start_time`) Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
* `cdc_start_time` - (Conflicts with `cdc_start_position`) The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
* `migration_type` - The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
* `replication_instance_arn` - The Amazon Resource Name (ARN) of the replication instance.
* `replication_task_settings` - An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
* `source_endpoint_arn` - The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
* `start_replication_task` -  Whether to run or stop the replication task.
* `status` - Replication Task status.
* `table_mappings` - An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
* `target_endpoint_arn` - The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
* `replication_task_arn` - The Amazon Resource Name (ARN) for the replication task.

<!-- cache-key: cdktf-0.20.8 input-bd68a0d3d48ee95e21b042c28d8b92612fdcb7f3700886e85dd6551b43189b4d -->