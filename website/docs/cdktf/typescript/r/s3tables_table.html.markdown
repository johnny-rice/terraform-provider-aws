---
subcategory: "S3 Tables"
layout: "aws"
page_title: "AWS: aws_s3tables_table"
description: |-
  Terraform resource for managing an Amazon S3 Tables Table.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_s3tables_table

Terraform resource for managing an Amazon S3 Tables Table.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { S3TablesNamespace } from "./.gen/providers/aws/s3-tables-namespace";
import { S3TablesTable } from "./.gen/providers/aws/s3-tables-table";
import { S3TablesTableBucket } from "./.gen/providers/aws/s3-tables-table-bucket";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new S3TablesTableBucket(this, "example", {
      name: "example-bucket",
    });
    const awsS3TablesNamespaceExample = new S3TablesNamespace(
      this,
      "example_1",
      {
        namespace: "example_namespace",
        tableBucketArn: example.arn,
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3TablesNamespaceExample.overrideLogicalId("example");
    const awsS3TablesTableExample = new S3TablesTable(this, "example_2", {
      format: "ICEBERG",
      name: "example_table",
      namespace: Token.asString(awsS3TablesNamespaceExample.namespace),
      tableBucketArn: Token.asString(
        awsS3TablesNamespaceExample.tableBucketArn
      ),
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3TablesTableExample.overrideLogicalId("example");
  }
}

```

### With Metadata Schema

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { S3TablesNamespace } from "./.gen/providers/aws/s3-tables-namespace";
import { S3TablesTable } from "./.gen/providers/aws/s3-tables-table";
import { S3TablesTableBucket } from "./.gen/providers/aws/s3-tables-table-bucket";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new S3TablesTableBucket(this, "example", {
      name: "example-bucket",
    });
    const awsS3TablesNamespaceExample = new S3TablesNamespace(
      this,
      "example_1",
      {
        namespace: "example_namespace",
        tableBucketArn: example.arn,
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3TablesNamespaceExample.overrideLogicalId("example");
    const awsS3TablesTableExample = new S3TablesTable(this, "example_2", {
      format: "ICEBERG",
      metadata: [
        {
          iceberg: [
            {
              schema: [
                {
                  field: [
                    {
                      name: "id",
                      required: true,
                      type: "long",
                    },
                    {
                      name: "name",
                      required: true,
                      type: "string",
                    },
                    {
                      name: "created_at",
                      required: false,
                      type: "timestamp",
                    },
                    {
                      name: "price",
                      required: false,
                      type: "decimal(10,2)",
                    },
                  ],
                },
              ],
            },
          ],
        },
      ],
      name: "example_table",
      namespace: Token.asString(awsS3TablesNamespaceExample.namespace),
      tableBucketArn: Token.asString(
        awsS3TablesNamespaceExample.tableBucketArn
      ),
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3TablesTableExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

The following arguments are required:

* `format` - (Required) Format of the table.
  Must be `ICEBERG`.
* `name` - (Required) Name of the table.
  Must be between 1 and 255 characters in length.
  Can consist of lowercase letters, numbers, and underscores, and must begin and end with a lowercase letter or number.
  A full list of table naming rules can be found in the [S3 Tables documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-buckets-naming.html#naming-rules-table).
* `namespace` - (Required) Name of the namespace for this table.
  Must be between 1 and 255 characters in length.
  Can consist of lowercase letters, numbers, and underscores, and must begin and end with a lowercase letter or number.
* `tableBucketArn` - (Required, Forces new resource) ARN referencing the Table Bucket that contains this Namespace.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `encryptionConfiguration` - (Optional) A single table bucket encryption configuration object.
  [See `encryptionConfiguration` below](#encryption_configuration).
* `maintenanceConfiguration` - (Optional) A single table bucket maintenance configuration object.
  [See `maintenanceConfiguration` below](#maintenance_configuration).
* `metadata` - (Optional) Contains details about the table metadata. This configuration specifies the metadata format and schema for the table. Currently only supports Iceberg format.
  [See `metadata` below](#metadata).

### `encryptionConfiguration`

The `encryptionConfiguration` object supports the following arguments:

* `kmsKeyArn` - (Optional) The ARN of a KMS Key to be used with `aws:kms` `sseAlgorithm`
* `sseAlgorithm` - (Required) One of `aws:kms` or `AES256`

### `maintenanceConfiguration`

The `maintenanceConfiguration` object supports the following arguments:

* `iceberg_compaction` - (Required) A single Iceberg compaction settings object.
  [See `iceberg_compaction` below](#iceberg_compaction).
* `iceberg_snapshot_management` - (Required) A single Iceberg snapshot management settings object.
  [See `iceberg_snapshot_management` below](#iceberg_snapshot_management).

### `iceberg_compaction`

The `iceberg_compaction` object supports the following arguments:

* `settings` - (Required) Settings object for compaction.
  [See `iceberg_compaction.settings` below](#iceberg_compactionsettings).
* `status` - (Required) Whether the configuration is enabled.
  Valid values are `enabled` and `disabled`.

### `iceberg_compaction.settings`

The `iceberg_compaction.settings` object supports the following argument:

* `target_file_size_mb` - (Required) Data objects smaller than this size may be combined with others to improve query performance.
  Must be between `64` and `512`.

### `iceberg_snapshot_management`

The `iceberg_snapshot_management` configuration block supports the following arguments:

* `settings` - (Required) Settings object for snapshot management.
  [See `iceberg_snapshot_management.settings` below](#iceberg_snapshot_managementsettings).
* `status` - (Required) Whether the configuration is enabled.
  Valid values are `enabled` and `disabled`.

### `iceberg_snapshot_management.settings`

The `iceberg_snapshot_management.settings` object supports the following argument:

* `max_snapshot_age_hours` - (Required) Snapshots older than this will be marked for deletiion.
  Must be at least `1`.
* `min_snapshots_to_keep` - (Required) Minimum number of snapshots to keep.
  Must be at least `1`.

### `metadata`

The `metadata` configuration block supports the following argument:

* `iceberg` - (Optional) Contains details about the metadata for an Iceberg table. This block defines the schema structure for the Apache Iceberg table format.
  [See `iceberg` below](#iceberg).

### `iceberg`

The `iceberg` configuration block supports the following argument:

* `schema` - (Required) Schema configuration for the Iceberg table.
  [See `schema` below](#schema).

### `schema`

The `schema` configuration block supports the following argument:

* `field` - (Required) List of schema fields for the Iceberg table. Each field defines a column in the table schema.
  [See `field` below](#field).

### `field`

The `field` configuration block supports the following arguments:

* `name` - (Required) The name of the field.
* `type` - (Required) The field type. S3 Tables supports all Apache Iceberg primitive types including: `boolean`, `int`, `long`, `float`, `double`, `decimal(precision,scale)`, `date`, `time`, `timestamp`, `timestamptz`, `string`, `uuid`, `fixed(length)`, `binary`.
* `required` - (Optional) A Boolean value that specifies whether values are required for each row in this field. Defaults to `false`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the table.
* `createdAt` - Date and time when the namespace was created.
* `createdBy` - Account ID of the account that created the namespace.
* `metadataLocation` - Location of table metadata.
* `modifiedAt` - Date and time when the namespace was last modified.
* `modifiedBy` - Account ID of the account that last modified the namespace.
* `ownerAccountId` - Account ID of the account that owns the namespace.
* `type` - Type of the table.
  One of `customer` or `aws`.
* `versionToken` - Identifier for the current version of table data.
* `warehouseLocation` - S3 URI pointing to the S3 Bucket that contains the table data.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import S3 Tables Table using the `tableBucketArn`, the value of `namespace`, and the value of `name`, separated by a semicolon (`;`). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { S3TablesTable } from "./.gen/providers/aws/s3-tables-table";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    S3TablesTable.generateConfigForImport(
      this,
      "example",
      "arn:aws:s3tables:us-west-2:123456789012:bucket/example-bucket;example-namespace;example-table"
    );
  }
}

```

Using `terraform import`, import S3 Tables Table using the `tableBucketArn`, the value of `namespace`, and the value of `name`, separated by a semicolon (`;`). For example:

```console
% terraform import aws_s3tables_table.example 'arn:aws:s3tables:us-west-2:123456789012:bucket/example-bucket;example-namespace;example-table'
```

<!-- cache-key: cdktf-0.20.8 input-eef39b038544ae73eabd552c24c04ba424d0b404744c63136e10ed9a763c4a80 -->