---
subcategory: "DocumentDB"
layout: "aws"
page_title: "AWS: aws_docdb_orderable_db_instance"
description: |-
  Information about DocumentDB orderable DB instances.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_docdb_orderable_db_instance

Information about DocumentDB orderable DB instances.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsDocdbOrderableDbInstance } from "./.gen/providers/aws/data-aws-docdb-orderable-db-instance";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsDocdbOrderableDbInstance(this, "test", {
      engine: "docdb",
      engineVersion: "3.6.0",
      licenseModel: "na",
      preferredInstanceClasses: ["db.r5.large", "db.r4.large", "db.t3.medium"],
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `engine` - (Optional) DB engine. Default: `docdb`
* `engineVersion` - (Optional) Version of the DB engine.
* `instanceClass` - (Optional) DB instance class. Examples of classes are `db.r5.12xlarge`, `db.r5.24xlarge`, `db.r5.2xlarge`, `db.r5.4xlarge`, `db.r5.large`, `db.r5.xlarge`, and `db.t3.medium`. (Conflicts with `preferredInstanceClasses`.)
* `licenseModel` - (Optional) License model. Default: `na`
* `preferredInstanceClasses` - (Optional) Ordered list of preferred DocumentDB DB instance classes. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned. (Conflicts with `instanceClass`.)
* `vpc` - (Optional) Enable to show only VPC.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `availabilityZones` - Availability zones where the instance is available.

<!-- cache-key: cdktf-0.20.8 input-290904cc8a40704b832a7c27048240d52b95a7189ab6d04cc55d799a806f1c00 -->