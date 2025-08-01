---
subcategory: "RDS (Relational Database)"
layout: "aws"
page_title: "AWS: aws_rds_orderable_db_instance"
description: |-
  Information about RDS orderable DB instances.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_rds_orderable_db_instance

Information about RDS orderable DB instances and valid parameter combinations.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsRdsOrderableDbInstance } from "./.gen/providers/aws/data-aws-rds-orderable-db-instance";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsRdsOrderableDbInstance(this, "test", {
      engine: "mysql",
      engineVersion: "5.7.22",
      licenseModel: "general-public-license",
      preferredInstanceClasses: ["db.r6.xlarge", "db.m4.large", "db.t3.small"],
      storageType: "standard",
    });
  }
}

```

Valid parameter combinations can also be found with `preferredEngineVersions` and/or `preferredInstanceClasses`.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsRdsOrderableDbInstance } from "./.gen/providers/aws/data-aws-rds-orderable-db-instance";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsRdsOrderableDbInstance(this, "test", {
      engine: "mysql",
      licenseModel: "general-public-license",
      preferredEngineVersions: ["5.6.35", "5.6.41", "5.6.44"],
      preferredInstanceClasses: ["db.t2.small", "db.t3.medium", "db.t3.large"],
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `availabilityZoneGroup` - (Optional) Availability zone group.
* `engineLatestVersion` - (Optional) When set to `true`, the data source attempts to return the most recent version matching the other criteria you provide. You must use `engineLatestVersion` with `preferredInstanceClasses` and/or `preferredEngineVersions`. Using `engineLatestVersion` will avoid `multiple RDS DB Instance Classes` errors. If you use `engineLatestVersion` with `preferredInstanceClasses`, the data source returns the latest version for the _first_ matching instance class (instance class priority). **Note:** The data source uses a best-effort approach at selecting the latest version but due to the complexity of version identifiers across engines, using `engineLatestVersion` may _not_ return the latest version in every situation.
* `engineVersion` - (Optional) Version of the DB engine. If none is provided, the data source tries to use the AWS-defined default version that matches any other criteria.
* `engine` - (Required) DB engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
* `instanceClass` - (Optional) DB instance class. Examples of classes are `db.m3.2xlarge`, `db.t2.small`, and `db.m3.medium`.
* `licenseModel` - (Optional) License model. Examples of license models are `general-public-license`, `bring-your-own-license`, and `amazon-license`.
* `preferredEngineVersions` - (Optional) Ordered list of preferred RDS DB instance engine versions. When `engineLatestVersion` is not set, the data source will return the first match in this list that matches any other criteria. If the data source finds no preferred matches or multiple matches without `engineLatestVersion`, it returns an error. **CAUTION:** We don't recommend using `preferredEngineVersions` without `preferredInstanceClasses` since the data source returns an arbitrary `instanceClass` based on the first one AWS returns that matches the engine version and any other criteria.
* `preferredInstanceClasses` - (Optional) Ordered list of preferred RDS DB instance classes. The data source will return the first match in this list that matches any other criteria. If the data source finds no preferred matches or multiple matches without `engineLatestVersion`, it returns an error. If you use `preferredInstanceClasses` without `preferredEngineVersions` or `engineLatestVersion`, the data source returns an arbitrary `engineVersion` based on the first one AWS returns matching the instance class and any other criteria.
* `readReplicaCapable` - (Optional) Whether a DB instance can have a read replica.
* `storageType` - (Optional) Storage types. Examples of storage types are `standard`, `io1`, `gp2`, and `aurora`.
* `supportedEngineModes` - (Optional) Use to limit results to engine modes such as `provisioned`.
* `supportedNetworkTypes` - (Optional) Use to limit results to network types `IPV4` or `DUAL`.
* `supportsClusters` - (Optional) Whether to limit results to instances that support clusters.
* `supportsMultiAz` - (Optional) Whether to limit results to instances that are multi-AZ capable.
* `supportsEnhancedMonitoring` - (Optional) Enable this to ensure a DB instance supports Enhanced Monitoring at intervals from 1 to 60 seconds.
* `supportsGlobalDatabases` - (Optional) Enable this to ensure a DB instance supports Aurora global databases with a specific combination of other DB engine attributes.
* `supportsIamDatabaseAuthentication` - (Optional) Enable this to ensure a DB instance supports IAM database authentication.
* `supportsIops` - (Optional) Enable this to ensure a DB instance supports provisioned IOPS.
* `supportsKerberosAuthentication` - (Optional) Enable this to ensure a DB instance supports Kerberos Authentication.
* `supportsPerformanceInsights` - (Optional) Enable this to ensure a DB instance supports Performance Insights.
* `supportsStorageAutoscaling` - (Optional) Enable this to ensure Amazon RDS can automatically scale storage for DB instances that use the specified DB instance class.
* `supportsStorageEncryption` - (Optional) Enable this to ensure a DB instance supports encrypted storage.
* `vpc` - (Optional) Boolean that indicates whether to show only VPC or non-VPC offerings.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `availabilityZones` - Availability zones where the instance is available.
* `maxIopsPerDbInstance` - Maximum total provisioned IOPS for a DB instance.
* `maxIopsPerGib` - Maximum provisioned IOPS per GiB for a DB instance.
* `maxStorageSize` - Maximum storage size for a DB instance.
* `minIopsPerDbInstance` - Minimum total provisioned IOPS for a DB instance.
* `minIopsPerGib` - Minimum provisioned IOPS per GiB for a DB instance.
* `minStorageSize` - Minimum storage size for a DB instance.
* `multiAzCapable` - Whether a DB instance is Multi-AZ capable.
* `outpostCapable` - Whether a DB instance supports RDS on Outposts.

<!-- cache-key: cdktf-0.20.8 input-fbc85aec937bf802d9383955af5a81fc82b7deebbe576e0f9ccf828fde6c4470 -->