---
subcategory: "Redshift"
layout: "aws"
page_title: "AWS: aws_redshift_cluster_iam_roles"
description: |-
  Provides a Redshift Cluster IAM Roles resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_redshift_cluster_iam_roles

Provides a Redshift Cluster IAM Roles resource.

~> **NOTE:** A Redshift cluster's default IAM role can be managed both by this resource's `defaultIamRoleArn` argument and the [`aws_redshift_cluster`](redshift_cluster.html) resource's `defaultIamRoleArn` argument. Do not configure different values for both arguments. Doing so will cause a conflict of default IAM roles.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RedshiftClusterIamRoles } from "./.gen/providers/aws/redshift-cluster-iam-roles";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new RedshiftClusterIamRoles(this, "example", {
      clusterIdentifier: Token.asString(
        awsRedshiftClusterExample.clusterIdentifier
      ),
      iamRoleArns: [Token.asString(awsIamRoleExample.arn)],
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `clusterIdentifier` - (Required) The name of the Redshift Cluster IAM Roles.
* `iamRoleArns` - (Optional) A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
* `defaultIamRoleArn` - (Optional) The Amazon Resource Name (ARN) for the IAM role that was set as default for the cluster when the cluster was created.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The Redshift Cluster ID.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Redshift Cluster IAM Roless using the `clusterIdentifier`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RedshiftClusterIamRoles } from "./.gen/providers/aws/redshift-cluster-iam-roles";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    RedshiftClusterIamRoles.generateConfigForImport(
      this,
      "examplegroup1",
      "example"
    );
  }
}

```

Using `terraform import`, import Redshift Cluster IAM Roless using the `clusterIdentifier`. For example:

```console
% terraform import aws_redshift_cluster_iam_roles.examplegroup1 example
```

<!-- cache-key: cdktf-0.20.8 input-ce4a486cf010df1204adfdcdf32c77c32ccd00a01fe077b54dcf9395da9aad3d -->