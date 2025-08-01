---
subcategory: "CodeBuild"
layout: "aws"
page_title: "AWS: aws_codebuild_report_group"
description: |-
  Provides a CodeBuild Report Group resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_codebuild_report_group

Provides a CodeBuild Report Groups Resource.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CodebuildReportGroup } from "./.gen/providers/aws/codebuild-report-group";
import { DataAwsCallerIdentity } from "./.gen/providers/aws/data-aws-caller-identity";
import { DataAwsIamPolicyDocument } from "./.gen/providers/aws/data-aws-iam-policy-document";
import { KmsKey } from "./.gen/providers/aws/kms-key";
import { S3Bucket } from "./.gen/providers/aws/s3-bucket";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new S3Bucket(this, "example", {
      bucket: "my-test",
    });
    const current = new DataAwsCallerIdentity(this, "current", {});
    const dataAwsIamPolicyDocumentExample = new DataAwsIamPolicyDocument(
      this,
      "example_2",
      {
        statement: [
          {
            actions: ["kms:*"],
            effect: "Allow",
            principals: [
              {
                identifiers: ["arn:aws:iam::${" + current.accountId + "}:root"],
                type: "AWS",
              },
            ],
            resources: ["*"],
            sid: "Enable IAM User Permissions",
          },
        ],
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsIamPolicyDocumentExample.overrideLogicalId("example");
    const awsKmsKeyExample = new KmsKey(this, "example_3", {
      deletionWindowInDays: 7,
      description: "my test kms key",
      policy: Token.asString(dataAwsIamPolicyDocumentExample.json),
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsKmsKeyExample.overrideLogicalId("example");
    const awsCodebuildReportGroupExample = new CodebuildReportGroup(
      this,
      "example_4",
      {
        exportConfig: {
          s3Destination: {
            bucket: example.id,
            encryptionDisabled: false,
            encryptionKey: Token.asString(awsKmsKeyExample.arn),
            packaging: "NONE",
            path: "/some",
          },
          type: "S3",
        },
        name: "my test report group",
        type: "TEST",
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsCodebuildReportGroupExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of a Report Group.
* `type` - (Required) The type of the Report Group. Valid value are `TEST` and `CODE_COVERAGE`.
* `exportConfig` - (Required) Information about the destination where the raw data of this Report Group is exported. see [Export Config](#export-config) documented below.
* `deleteReports` - (Optional) If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
* `tags` - (Optional) Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### Export Config

* `type` - (Required) The export configuration type. Valid values are `S3` and `NO_EXPORT`.
* `s3Destination` - (Required) contains information about the S3 bucket where the run of a report is exported. see [S3 Destination](#s3-destination) documented below.

#### S3 Destination

* `bucket`- (Required) The name of the S3 bucket where the raw data of a report are exported.
* `encryptionKey` - (Required) The encryption key for the report's encrypted raw data. The KMS key ARN.
* `encryptionDisabled`- (Optional) A boolean value that specifies if the results of a report are encrypted.
 **Note: the API does not currently allow setting encryption as disabled**
* `packaging` - (Optional) The type of build output artifact to create. Valid values are: `NONE` (default) and `ZIP`.
* `path` - (Optional) The path to the exported report's raw data results.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ARN of Report Group.
* `arn` - The ARN of Report Group.
* `created` - The date and time this Report Group was created.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import CodeBuild Report Group using the CodeBuild Report Group arn. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CodebuildReportGroup } from "./.gen/providers/aws/codebuild-report-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    CodebuildReportGroup.generateConfigForImport(
      this,
      "example",
      "arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name"
    );
  }
}

```

Using `terraform import`, import CodeBuild Report Group using the CodeBuild Report Group arn. For example:

```console
% terraform import aws_codebuild_report_group.example arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
```

<!-- cache-key: cdktf-0.20.8 input-25b9334b88ebbecd60d2c5e4dcf48a26e87fafd1bd6398a5a7960787a264ec0f -->