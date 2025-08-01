---
subcategory: "RDS (Relational Database)"
layout: "aws"
page_title: "AWS: aws_rds_integration"
description: |-
  Terraform resource for managing an AWS RDS (Relational Database) Integration.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_rds_integration

Terraform resource for managing an AWS RDS (Relational Database) zero-ETL integration. You can refer to the [User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/zero-etl.setting-up.html).

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
import { RdsIntegration } from "./.gen/providers/aws/rds-integration";
import { RedshiftserverlessNamespace } from "./.gen/providers/aws/redshiftserverless-namespace";
import { RedshiftserverlessWorkgroup } from "./.gen/providers/aws/redshiftserverless-workgroup";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new RedshiftserverlessNamespace(this, "example", {
      namespaceName: "redshift-example",
    });
    const awsRedshiftserverlessWorkgroupExample =
      new RedshiftserverlessWorkgroup(this, "example_1", {
        baseCapacity: 8,
        configParameter: [
          {
            parameterKey: "enable_case_sensitive_identifier",
            parameterValue: "true",
          },
        ],
        namespaceName: example.namespaceName,
        publiclyAccessible: false,
        subnetIds: [example1.id, example2.id, example3.id],
        workgroupName: "example-workspace",
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsRedshiftserverlessWorkgroupExample.overrideLogicalId("example");
    const awsRdsIntegrationExample = new RdsIntegration(this, "example_2", {
      integrationName: "example",
      lifecycle: {
        ignoreChanges: [kmsKeyId],
      },
      sourceArn: Token.asString(awsRdsClusterExample.arn),
      targetArn: example.arn,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsRdsIntegrationExample.overrideLogicalId("example");
  }
}

```

### Use own KMS key

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsCallerIdentity } from "./.gen/providers/aws/data-aws-caller-identity";
import { DataAwsIamPolicyDocument } from "./.gen/providers/aws/data-aws-iam-policy-document";
import { KmsKey } from "./.gen/providers/aws/kms-key";
import { RdsIntegration } from "./.gen/providers/aws/rds-integration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const current = new DataAwsCallerIdentity(this, "current", {});
    const keyPolicy = new DataAwsIamPolicyDocument(this, "key_policy", {
      statement: [
        {
          actions: ["kms:*"],
          principals: [
            {
              identifiers: ["arn:aws:iam::${" + current.accountId + "}:root"],
              type: "AWS",
            },
          ],
          resources: ["*"],
        },
        {
          actions: ["kms:CreateGrant"],
          principals: [
            {
              identifiers: ["redshift.amazonaws.com"],
              type: "Service",
            },
          ],
          resources: ["*"],
        },
      ],
    });
    const example = new KmsKey(this, "example", {
      deletionWindowInDays: 10,
      policy: Token.asString(keyPolicy.json),
    });
    const awsRdsIntegrationExample = new RdsIntegration(this, "example_3", {
      additionalEncryptionContext: {
        example: "test",
      },
      integrationName: "example",
      kmsKeyId: example.arn,
      sourceArn: Token.asString(awsRdsClusterExample.arn),
      targetArn: Token.asString(awsRedshiftserverlessNamespaceExample.arn),
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsRdsIntegrationExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

The following arguments are required:

* `integrationName` - (Required, Forces new resources) Name of the integration.
* `sourceArn` - (Required, Forces new resources) ARN of the database to use as the source for replication.
* `targetArn` - (Required, Forces new resources) ARN of the Redshift data warehouse to use as the target for replication.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `additionalEncryptionContext` - (Optional, Forces new resources) Set of non-secret key–value pairs that contains additional contextual information about the data.
For more information, see the [User Guide](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context).
You can only include this parameter if you specify the `kmsKeyId` parameter.
* `dataFilter` - (Optional, Forces new resources) Data filters for the integration.
These filters determine which tables from the source database are sent to the target Amazon Redshift data warehouse.
The value should match the syntax from the AWS CLI which includes an `include:` or `exclude:` prefix before a filter expression.
Multiple expressions are separated by a comma.
See the [Amazon RDS data filtering guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/zero-etl.filtering.html) for additional details.
* `kmsKeyId` - (Optional, Forces new resources) KMS key identifier for the key to use to encrypt the integration.
If you don't specify an encryption key, RDS uses a default AWS owned key.
If you use the default AWS owned key, you should ignore `kmsKeyId` parameter by using [`lifecycle` parameter](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle#ignore_changes) to avoid unintended change after the first creation.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

For more detailed documentation about each argument, refer to the [AWS official documentation](https://docs.aws.amazon.com/cli/latest/reference/rds/create-integration.html).

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Integration.
* `id` - (**Deprecated**, use `arn` instead) ARN of the Integration.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `60m`)
* `update` - (Default `10m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import RDS (Relational Database) Integration using the `arn`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RdsIntegration } from "./.gen/providers/aws/rds-integration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    RdsIntegration.generateConfigForImport(
      this,
      "example",
      "arn:aws:rds:us-west-2:123456789012:integration:abcdefgh-0000-1111-2222-123456789012"
    );
  }
}

```

Using `terraform import`, import RDS (Relational Database) Integration using the `arn`. For example:

```console
% terraform import aws_rds_integration.example arn:aws:rds:us-west-2:123456789012:integration:abcdefgh-0000-1111-2222-123456789012
```

<!-- cache-key: cdktf-0.20.8 input-e1d279db05e5935cabfde974497d8d4bde326c92e9257925cf70f61a7d11fbc9 -->