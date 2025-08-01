---
subcategory: "QuickSight"
layout: "aws"
page_title: "AWS: aws_quicksight_dashboard"
description: |-
  Manages a QuickSight Dashboard.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_quicksight_dashboard

Resource for managing a QuickSight Dashboard.

## Example Usage

### From Source Template

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.quicksight_dashboard import QuicksightDashboard
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        QuicksightDashboard(self, "example",
            dashboard_id="example-id",
            name="example-name",
            source_entity=QuicksightDashboardSourceEntity(
                source_template=QuicksightDashboardSourceEntitySourceTemplate(
                    arn=source.arn,
                    data_set_references=[QuicksightDashboardSourceEntitySourceTemplateDataSetReferences(
                        data_set_arn=dataset.arn,
                        data_set_placeholder="1"
                    )
                    ]
                )
            ),
            version_description="version"
        )
```

### With Definition

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.quicksight_dashboard import QuicksightDashboard
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        QuicksightDashboard(self, "example",
            dashboard_id="example-id",
            definition={
                "data_set_identifiers_declarations": [{
                    "data_set_arn": dataset.arn,
                    "identifier": "1"
                }
                ],
                "sheets": [{
                    "sheet_id": "Example1",
                    "title": "Example",
                    "visuals": [{
                        "line_chart_visual": {
                            "chart_configuration": {
                                "field_wells": {
                                    "line_chart_aggregated_field_wells": {
                                        "category": [{
                                            "categorical_dimension_field": {
                                                "column": {
                                                    "column_name": "Column1",
                                                    "data_set_identifier": "1"
                                                },
                                                "field_id": "1"
                                            }
                                        }
                                        ],
                                        "values": [{
                                            "categorical_measure_field": {
                                                "aggregation_function": "COUNT",
                                                "column": {
                                                    "column_name": "Column1",
                                                    "data_set_identifier": "1"
                                                },
                                                "field_id": "2"
                                            }
                                        }
                                        ]
                                    }
                                }
                            },
                            "title": {
                                "format_text": {
                                    "plain_text": "Line Chart Example"
                                }
                            },
                            "visual_id": "LineChart"
                        }
                    }
                    ]
                }
                ]
            },
            name="example-name",
            version_description="version"
        )
```

## Argument Reference

The following arguments are required:

* `dashboard_id` - (Required, Forces new resource) Identifier for the dashboard.
* `name` - (Required) Display name for the dashboard.
* `version_description` - (Required) A description of the current dashboard version being created/updated.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `aws_account_id` - (Optional, Forces new resource) AWS account ID.
* `dashboard_publish_options` - (Optional) Options for publishing the dashboard. See [dashboard_publish_options](#dashboard_publish_options).
* `definition` - (Optional) A detailed dashboard definition. Only one of `definition` or `source_entity` should be configured. See [definition](#definition).
* `parameters` - (Optional) The parameters for the creation of the dashboard, which you want to use to override the default settings. A dashboard can have any type of parameters, and some parameters might accept multiple values. See [parameters](#parameters).
* `permissions` - (Optional) A set of resource permissions on the dashboard. Maximum of 64 items. See [permissions](#permissions).
* `source_entity` - (Optional) The entity that you are using as a source when you create the dashboard (template). Only one of `definition` or `source_entity` should be configured. See [source_entity](#source_entity).
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `theme_arn` - (Optional) The Amazon Resource Name (ARN) of the theme that is being used for this dashboard. The theme ARN must exist in the same AWS account where you create the dashboard.

### permissions

* `actions` - (Required) List of IAM actions to grant or revoke permissions on.
* `principal` - (Required) ARN of the principal. See the [ResourcePermission documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ResourcePermission.html) for the applicable ARN values.

### source_entity

* `source_template` - (Optional) The source template. See [source_template](#source_template).

### source_template

* `arn` - (Required) The Amazon Resource Name (ARN) of the resource.
* `data_set_references` - (Required) List of dataset references. See [data_set_references](#data_set_references).

### data_set_references

* `data_set_arn` - (Required) Dataset Amazon Resource Name (ARN).
* `data_set_placeholder` - (Required) Dataset placeholder.

### dashboard_publish_options

* `ad_hoc_filtering_option` - (Optional) Ad hoc (one-time) filtering option. See [ad_hoc_filtering_option](#ad_hoc_filtering_option).
* `data_point_drill_up_down_option` - (Optional) The drill-down options of data points in a dashboard. See [data_point_drill_up_down_option](#data_point_drill_up_down_option).
* `data_point_menu_label_option` - (Optional) The data point menu label options of a dashboard. See [data_point_menu_label_option](#data_point_menu_label_option).
* `data_point_tooltip_option` - (Optional) The data point tool tip options of a dashboard. See [data_point_tooltip_option](#data_point_tooltip_option).
* `export_to_csv_option` - (Optional) Export to .csv option. See [export_to_csv_option](#export_to_csv_option).
* `export_with_hidden_fields_option` - (Optional) Determines if hidden fields are exported with a dashboard. See [export_with_hidden_fields_option](#export_with_hidden_fields_option).
* `sheet_controls_option` - (Optional) Sheet controls option. See [sheet_controls_option](#sheet_controls_option).
* `sheet_layout_element_maximization_option` - (Optional) The sheet layout maximization options of a dashboard. See [sheet_layout_element_maximization_option](#sheet_layout_element_maximization_option).
* `visual_axis_sort_option` - (Optional) The axis sort options of a dashboard. See [visual_axis_sort_option](#visual_axis_sort_option).
* `visual_menu_option` - (Optional) The menu options of a visual in a dashboard. See [visual_menu_option](#visual_menu_option).

### ad_hoc_filtering_option

* `availability_status` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### data_point_drill_up_down_option

* `availability_status` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### data_point_menu_label_option

* `availability_status` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### data_point_tooltip_option

* `availability_status` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### export_to_csv_option

* `availability_status` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### export_with_hidden_fields_option

* `availability_status` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### sheet_controls_option

* `visibility_state` - (Optional) Visibility state. Possibles values: EXPANDED, COLLAPSED.

### sheet_layout_element_maximization_option

* `availability_status` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### visual_axis_sort_option

* `availability_status` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### visual_menu_option

* `availability_status` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### parameters

* `date_time_parameters` - (Optional) A list of parameters that have a data type of date-time. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DateTimeParameter.html).
* `decimal_parameters` - (Optional) A list of parameters that have a data type of decimal. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DecimalParameter.html).
* `integer_parameters` - (Optional) A list of parameters that have a data type of integer. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_IntegerParameter.html).
* `string_parameters` - (Optional) A list of parameters that have a data type of string. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_StringParameter.html).

### definition

* `data_set_identifiers_declarations` - (Required) A list dataset identifier declarations. With this mapping,you can use dataset identifiers instead of dataset Amazon Resource Names (ARNs) throughout the dashboard's sub-structures. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSetIdentifierDeclaration.html).
* `analysis_defaults` - (Optional) The configuration for default analysis settings. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_AnalysisDefaults.html).
* `calculated_fields` - (Optional) A list of calculated field definitions for the dashboard. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CalculatedField.html).
* `column_configurations` - (Optional) A list of dashboard-level column configurations. Column configurations are used to set default formatting for a column that's used throughout a dashboard. See [AWS API Documentation for complete description](ttps://docs.aws.amazon.com/quicksight/latest/APIReference/API_ColumnConfiguration.html).
* `filter_groups` - (Optional) A list of filter definitions for a dashboard. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_FilterGroup.html). For more information, see [Filtering Data](https://docs.aws.amazon.com/quicksight/latest/user/filtering-visual-data.html) in Amazon QuickSight User Guide.
* `parameters_declarations` - (Optional) A list of parameter declarations for a dashboard. Parameters are named variables that can transfer a value for use by an action or an object. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ParameterDeclaration.html). For more information, see [Parameters in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-in-quicksight.html) in the Amazon QuickSight User Guide.
* `sheets` - (Optional) A list of sheet definitions for a dashboard. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_SheetDefinition.html).

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the dashboard.
* `created_time` - The time that the dashboard was created.
* `id` - A comma-delimited string joining AWS account ID and dashboard ID.
* `last_updated_time` - The time that the dashboard was last updated.
* `source_entity_arn` - Amazon Resource Name (ARN) of a template that was used to create this dashboard.
* `status` - The dashboard creation status.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block).
* `version_number` - The version number of the dashboard version.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `5m`)
* `update` - (Default `5m`)
* `delete` - (Default `5m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import a QuickSight Dashboard using the AWS account ID and dashboard ID separated by a comma (`,`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.quicksight_dashboard import QuicksightDashboard
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        QuicksightDashboard.generate_config_for_import(self, "example", "123456789012,example-id")
```

Using `terraform import`, import a QuickSight Dashboard using the AWS account ID and dashboard ID separated by a comma (`,`). For example:

```console
% terraform import aws_quicksight_dashboard.example 123456789012,example-id
```

<!-- cache-key: cdktf-0.20.8 input-0dde427a0ea9320283cfea5d88db74f856433081838c7dfd8d83488fb5d678d1 -->