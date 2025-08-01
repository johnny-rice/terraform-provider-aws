---
subcategory: "QuickSight"
layout: "aws"
page_title: "AWS: aws_quicksight_data_set"
description: |-
  Manages a Resource QuickSight Data Set.
---

# Resource: aws_quicksight_data_set

Resource for managing a QuickSight Data Set.

## Example Usage

### Basic Usage

```terraform
resource "aws_quicksight_data_set" "example" {
  data_set_id = "example-id"
  name        = "example-name"
  import_mode = "SPICE"

  physical_table_map {
    physical_table_map_id = "example-id"
    s3_source {
      data_source_arn = aws_quicksight_data_source.example.arn
      input_columns {
        name = "Column1"
        type = "STRING"
      }
      upload_settings {
        format = "JSON"
      }
    }
  }
}
```

### With Column Level Permission Rules

```terraform
resource "aws_quicksight_data_set" "example" {
  data_set_id = "example-id"
  name        = "example-name"
  import_mode = "SPICE"

  physical_table_map {
    physical_table_map_id = "example-id"
    s3_source {
      data_source_arn = aws_quicksight_data_source.example.arn
      input_columns {
        name = "Column1"
        type = "STRING"
      }
      upload_settings {
        format = "JSON"
      }
    }
  }
  column_level_permission_rules {
    column_names = ["Column1"]
    principals   = [aws_quicksight_user.example.arn]
  }
}
```

### With Field Folders

```terraform
resource "aws_quicksight_data_set" "example" {
  data_set_id = "example-id"
  name        = "example-name"
  import_mode = "SPICE"

  physical_table_map {
    physical_table_map_id = "example-id"
    s3_source {
      data_source_arn = aws_quicksight_data_source.example.arn
      input_columns {
        name = "Column1"
        type = "STRING"
      }
      upload_settings {
        format = "JSON"
      }
    }
  }
  field_folders {
    field_folders_id = "example-id"
    columns          = ["Column1"]
    description      = "example description"
  }
}
```

### With Permissions

```terraform
resource "aws_quicksight_data_set" "example" {
  data_set_id = "example-id"
  name        = "example-name"
  import_mode = "SPICE"

  physical_table_map {
    physical_table_map_id = "example-id"
    s3_source {
      data_source_arn = aws_quicksight_data_source.example.arn
      input_columns {
        name = "Column1"
        type = "STRING"
      }
      upload_settings {
        format = "JSON"
      }
    }
  }
  permissions {
    actions = [
      "quicksight:DescribeDataSet",
      "quicksight:DescribeDataSetPermissions",
      "quicksight:PassDataSet",
      "quicksight:DescribeIngestion",
      "quicksight:ListIngestions",
    ]
    principal = aws_quicksight_user.example.arn
  }
}
```

### With Row Level Permission Tag Configuration

```terraform
resource "aws_quicksight_data_set" "example" {
  data_set_id = "example-id"
  name        = "example-name"
  import_mode = "SPICE"

  physical_table_map {
    physical_table_map_id = "example-id"
    s3_source {
      data_source_arn = aws_quicksight_data_source.example.arn
      input_columns {
        name = "Column1"
        type = "STRING"
      }
      upload_settings {
        format = "JSON"
      }
    }
  }
  row_level_permission_tag_configuration {
    status = "ENABLED"
    tag_rules {
      column_name               = "Column1"
      tag_key                   = "tagkey"
      match_all_value           = "*"
      tag_multi_value_delimiter = ","
    }
  }
}
```

## Argument Reference

The following arguments are required:

* `data_set_id` - (Required, Forces new resource) Identifier for the data set.
* `import_mode` - (Required) Indicates whether you want to import the data into SPICE. Valid values are `SPICE` and `DIRECT_QUERY`.
* `name` - (Required) Display name for the dataset.

The following arguments are optional:

* `aws_account_id` - (Optional, Forces new resource) AWS account ID. Defaults to automatically determined account ID of the Terraform AWS provider.
* `column_groups` - (Optional) Groupings of columns that work together in certain Amazon QuickSight features. Currently, only geospatial hierarchy is supported. See [column_groups](#column_groups).
* `column_level_permission_rules` - (Optional) A set of 1 or more definitions of a [ColumnLevelPermissionRule](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ColumnLevelPermissionRule.html). See [column_level_permission_rules](#column_level_permission_rules).
* `data_set_usage_configuration` - (Optional) The usage configuration to apply to child datasets that reference this dataset as a source. See [data_set_usage_configuration](#data_set_usage_configuration).
* `field_folders` - (Optional) The folder that contains fields and nested subfolders for your dataset. See [field_folders](#field_folders).
* `logical_table_map` - (Optional) Configures the combination and transformation of the data from the physical tables. Maximum of 1 entry. See [logical_table_map](#logical_table_map).
* `permissions` - (Optional) A set of resource permissions on the data source. Maximum of 64 items. See [permissions](#permissions).
* `physical_table_map` - (Optional) Declares the physical tables that are available in the underlying data sources. See [physical_table_map](#physical_table_map).
* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `row_level_permission_data_set` - (Optional) The row-level security configuration for the data that you want to create. See [row_level_permission_data_set](#row_level_permission_data_set).
* `row_level_permission_tag_configuration` - (Optional) The configuration of tags on a dataset to set row-level security. Row-level security tags are currently supported for anonymous embedding only. See [row_level_permission_tag_configuration](#row_level_permission_tag_configuration).
* `refresh_properties` - (Optional) The refresh properties for the data set. **NOTE**: Only valid when `import_mode` is set to `SPICE`. See [refresh_properties](#refresh_properties).
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### physical_table_map

For a `physical_table_map` item to be valid, only one of `custom_sql`, `relational_table`, or `s3_source` should be configured.

* `physical_table_map_id` - (Required) Key of the physical table map.
* `custom_sql` - (Optional) A physical table type built from the results of the custom SQL query. See [custom_sql](#custom_sql).
* `relational_table` - (Optional) A physical table type for relational data sources. See [relational_table](#relational_table).
* `s3_source` - (Optional) A physical table type for as S3 data source. See [s3_source](#s3_source).

### custom_sql

* `data_source_arn` - (Required) ARN of the data source.
* `name` - (Required) Display name for the SQL query result.
* `sql_query` - (Required) SQL query.
* `columns` - (Optional) Column schema from the SQL query result set. See [columns](#columns).

### columns

* `name` - (Required) Name of this column in the underlying data source.
* `type` - (Required) Data type of the column.

### relational_table

* `data_source_arn` - (Required) ARN of the data source.
* `input_columns` - (Required) Column schema of the table. See [input_columns](#input_columns).
* `name` - (Required) Name of the relational table.
* `catalog` - (Optional) Catalog associated with the table.
* `schema` - (Optional) Schema name. This name applies to certain relational database engines.

### input_columns

* `name` - (Required) Name of this column in the underlying data source.
* `type` - (Required) Data type of the column.

### s3_source

* `data_source_arn` - (Required) ARN of the data source.
* `input_columns` - (Required) Column schema of the table. See [input_columns](#input_columns).
* `upload_settings` - (Required) Information about the format for the S3 source file or files. See [upload_settings](#upload_settings).

### upload_settings

* `contains_header` - (Optional) Whether the file has a header row, or the files each have a header row.
* `delimiter` - (Optional) Delimiter between values in the file.
* `format` - (Optional) File format. Valid values are `CSV`, `TSV`, `CLF`, `ELF`, `XLSX`, and `JSON`.
* `start_from_row` - (Optional) A row number to start reading data from.
* `text_qualifier` - (Optional) Text qualifier. Valid values are `DOUBLE_QUOTE` and `SINGLE_QUOTE`.

### column_groups

* `geo_spatial_column_group` - (Optional) Geospatial column group that denotes a hierarchy. See [geo_spatial_column_group](#geo_spatial_column_group).

### geo_spatial_column_group

* `columns` - (Required) Columns in this hierarchy.
* `country_code` - (Required) Country code. Valid values are `US`.
* `name` - (Required) A display name for the hierarchy.

### column_level_permission_rules

* `column_names` - (Optional) An array of column names.
* `principals` - (Optional) An array of ARNs for Amazon QuickSight users or groups.

### data_set_usage_configuration

* `disable_use_as_direct_query_source` - (Optional) Controls whether a child dataset of a direct query can use this dataset as a source.
* `disable_use_as_imported_source` - (Optional) Controls whether a child dataset that's stored in QuickSight can use this dataset as a source.

### field_folders

* `field_folders_id` - (Required) Key of the field folder map.
* `columns` - (Optional) An array of column names to add to the folder. A column can only be in one folder.
* `description` - (Optional) Field folder description.

### logical_table_map

* `alias` - (Required) A display name for the logical table.
* `logical_table_map_id` - (Required) Key of the logical table map.
* `data_transforms` - (Optional) Transform operations that act on this logical table. For this structure to be valid, only one of the attributes can be non-null. See [data_transforms](#data_transforms).
* `source` - (Optional) Source of this logical table. See [source](#source).

### data_transforms

* `cast_column_type_operation` - (Optional) A transform operation that casts a column to a different type. See [cast_column_type_operation](#cast_column_type_operation).
* `create_columns_operation` - (Optional) An operation that creates calculated columns. Columns created in one such operation form a lexical closure. See [create_columns_operation](#create_columns_operation).
* `filter_operation` - (Optional) An operation that filters rows based on some condition. See [filter_operation](#filter_operation).
* `project_operation` - (Optional) An operation that projects columns. Operations that come after a projection can only refer to projected columns. See [project_operation](#project_operation).
* `rename_column_operation` - (Optional) An operation that renames a column. See [rename_column_operation](#rename_column_operation).
* `tag_column_operation` - (Optional) An operation that tags a column with additional information. See [tag_column_operation](#tag_column_operation).
* `untag_column_operation` - (Optional) A transform operation that removes tags associated with a column. See [untag_column_operation](#untag_column_operation).

### cast_column_type_operation

* `column_name` - (Required) Column name.
* `new_column_type` - (Required) New column data type. Valid values are `STRING`, `INTEGER`, `DECIMAL`, `DATETIME`.
* `format` - (Optional) When casting a column from string to datetime type, you can supply a string in a format supported by Amazon QuickSight to denote the source data format.

### create_columns_operation

* `columns` - (Required) Calculated columns to create. See [columns](#columns-1).

### columns

* `column_id` - (Required) A unique ID to identify a calculated column. During a dataset update, if the column ID of a calculated column matches that of an existing calculated column, Amazon QuickSight preserves the existing calculated column.
* `column_name` - (Required) Column name.
* `expression` - (Required) An expression that defines the calculated column.

### filter_operation

* `condition_expression` - (Required) An expression that must evaluate to a Boolean value. Rows for which the expression evaluates to true are kept in the dataset.

### project_operation

* `projected_columns` - (Required) Projected columns.

### rename_column_operation

* `column_name` - (Required) Column to be renamed.
* `new_column_name` - (Required) New name for the column.

### tag_column_operation

* `column_name` - (Required) Column name.
* `tags` - (Required) The dataset column tag, currently only used for geospatial type tagging. See [tags](#tags).

### tags

* `column_description` - (Optional) A description for a column. See [column_description](#column_description).
* `column_geographic_role` - (Optional) A geospatial role for a column. Valid values are `COUNTRY`, `STATE`, `COUNTY`, `CITY`, `POSTCODE`, `LONGITUDE`, and `LATITUDE`.

### column_description

* `text` - (Optional) The text of a description for a column.

### untag_column_operation

* `column_name` - (Required) Column name.
* `tag_names` - (Required) The column tags to remove from this column.

### source

* `data_set_arn` - (Optional) ARN of the parent data set.
* `join_instruction` - (Optional) Specifies the result of a join of two logical tables. See [join_instruction](#join_instruction).
* `physical_table_id` - (Optional) Physical table ID.

### join_instruction

* `left_operand` - (Required) Operand on the left side of a join.
* `on_clause` - (Required) Join instructions provided in the ON clause of a join.
* `right_operand` - (Required) Operand on the right side of a join.
* `type` - (Required) Type of join. Valid values are `INNER`, `OUTER`, `LEFT`, and `RIGHT`.
* `left_join_key_properties` - (Optional) Join key properties of the left operand. See [left_join_key_properties](#left_join_key_properties).
* `right_join_key_properties` - (Optional) Join key properties of the right operand. See [right_join_key_properties](#right_join_key_properties).

### left_join_key_properties

* `unique_key` - (Optional) A value that indicates that a row in a table is uniquely identified by the columns in a join key. This is used by Amazon QuickSight to optimize query performance.

### right_join_key_properties

* `unique_key` - (Optional) A value that indicates that a row in a table is uniquely identified by the columns in a join key. This is used by Amazon QuickSight to optimize query performance.

### permissions

* `actions` - (Required) List of IAM actions to grant or revoke permissions on.
* `principal` - (Required) ARN of the principal. See the [ResourcePermission documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ResourcePermission.html) for the applicable ARN values.

### row_level_permission_data_set

* `arn` - (Required) ARN of the dataset that contains permissions for RLS.
* `permission_policy` - (Required) Type of permissions to use when interpreting the permissions for RLS. Valid values are `GRANT_ACCESS` and `DENY_ACCESS`.
* `format_version` - (Optional) User or group rules associated with the dataset that contains permissions for RLS.
* `namespace` - (Optional) Namespace associated with the dataset that contains permissions for RLS.
* `status` - (Optional) Status of the row-level security permission dataset. If enabled, the status is `ENABLED`. If disabled, the status is `DISABLED`.

### row_level_permission_tag_configuration

* `tag_rules` - (Required) A set of rules associated with row-level security, such as the tag names and columns that they are assigned to. See [tag_rules](#tag_rules).
* `status` - (Optional) The status of row-level security tags. If enabled, the status is `ENABLED`. If disabled, the status is `DISABLED`.

### refresh_properties

* `refresh_configuration` - (Required) The refresh configuration for the data set. See [refresh_configuration](#refresh_configuration).

### refresh_configuration

* `incremental_refresh` - (Required) The incremental refresh for the data set. See [incremental_refresh](#incremental_refresh).

### incremental_refresh

* `lookback_window` - (Required) The lookback window setup for an incremental refresh configuration. See [lookback_window](#lookback_window).

### lookback_window

* `column_name` - (Required) The name of the lookback window column.
* `size` - (Required) The lookback window column size.
* `size_unit` - (Required) The size unit that is used for the lookback window column. Valid values for this structure are `HOUR`, `DAY`, and `WEEK`.

### tag_rules

* `column_name` - (Required) Column name that a tag key is assigned to.
* `tag_key` - (Required) Unique key for a tag.
* `match_all_value` - (Optional) A string that you want to use to filter by all the values in a column in the dataset and don’t want to list the values one by one.
* `tag_multi_value_delimiter` - (Optional) A string that you want to use to delimit the values when you pass the values at run time.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the data set.
* `id` - A comma-delimited string joining AWS account ID and data set ID.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import a QuickSight Data Set using the AWS account ID and data set ID separated by a comma (`,`). For example:

```terraform
import {
  to = aws_quicksight_data_set.example
  id = "123456789012,example-id"
}
```

Using `terraform import`, import a QuickSight Data Set using the AWS account ID and data set ID separated by a comma (`,`). For example:

```console
% terraform import aws_quicksight_data_set.example 123456789012,example-id
```
