---
subcategory: "Lambda"
layout: "aws"
page_title: "AWS: aws_lambda_invocation"
description: |-
  Manages an AWS Lambda Function invocation.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lambda_invocation

Manages an AWS Lambda Function invocation. Use this resource to invoke a Lambda function with the [RequestResponse](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax) invocation type.

~> **Note:** By default this resource _only_ invokes the function when the arguments call for a create or replace. After an initial invocation on _apply_, if the arguments do not change, a subsequent _apply_ does not invoke the function again. To dynamically invoke the function, see the `triggers` example below. To always invoke a function on each _apply_, see the [`aws_lambda_invocation` data source](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/lambda_invocation). To invoke the Lambda function when the Terraform resource is updated and deleted, see the [CRUD Lifecycle Management](#crud-lifecycle-management) example below.

~> **Note:** If you get a `KMSAccessDeniedException: Lambda was unable to decrypt the environment variables because KMS access was denied` error when invoking a Lambda function with environment variables, the IAM role associated with the function may have been deleted and recreated after the function was created. You can fix the problem two ways: 1) updating the function's role to another role and then updating it back again to the recreated role, or 2) by using Terraform to `taint` the function and `apply` your configuration again to recreate the function. (When you create a function, Lambda grants permissions on the KMS key to the function's IAM role. If the IAM role is recreated, the grant is no longer valid. Changing the function's role or recreating the function causes Lambda to update the grant.)

## Example Usage

### Basic Invocation

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformOutput, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LambdaFunction } from "./.gen/providers/aws/lambda-function";
import { LambdaInvocation } from "./.gen/providers/aws/lambda-invocation";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new LambdaFunction(this, "example", {
      filename: "function.zip",
      functionName: "data_processor",
      handler: "index.handler",
      role: lambdaRole.arn,
      runtime: "python3.12",
    });
    const awsLambdaInvocationExample = new LambdaInvocation(this, "example_1", {
      functionName: example.functionName,
      input: Token.asString(
        Fn.jsonencode({
          config: {
            debug: false,
            environment: "production",
          },
          operation: "initialize",
        })
      ),
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLambdaInvocationExample.overrideLogicalId("example");
    new TerraformOutput(this, "initialization_result", {
      value: Fn.lookupNested(
        Fn.jsondecode(Token.asString(awsLambdaInvocationExample.result)),
        ['"status"']
      ),
    });
  }
}

```

### Dynamic Invocation with Triggers

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LambdaInvocation } from "./.gen/providers/aws/lambda-invocation";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new LambdaInvocation(this, "example", {
      functionName: Token.asString(awsLambdaFunctionExample.functionName),
      input: Token.asString(
        Fn.jsonencode({
          batch_id: batchId.result,
          environment: environment.value,
          operation: "process_data",
        })
      ),
      triggers: {
        config_hash: Token.asString(
          Fn.sha256(
            Token.asString(
              Fn.jsonencode({
                environment: environment.value,
                timestamp: Fn.timestamp(),
              })
            )
          )
        ),
        function_version: Token.asString(awsLambdaFunctionExample.version),
      },
    });
  }
}

```

### CRUD Lifecycle Management

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LambdaInvocation } from "./.gen/providers/aws/lambda-invocation";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new LambdaInvocation(this, "example", {
      functionName: Token.asString(awsLambdaFunctionExample.functionName),
      input: Token.asString(
        Fn.jsonencode({
          credentials: {
            password: dbPassword.value,
            username: dbUsername.value,
          },
          database_url: awsDbInstanceExample.endpoint,
          resource_name: "database_setup",
        })
      ),
      lifecycleScope: "CRUD",
    });
  }
}

```

~> **Note:** `lifecycle_scope = "CRUD"` will inject a key `tf` in the input event to pass lifecycle information! This allows the Lambda function to handle different lifecycle transitions uniquely. If you need to use a key `tf` in your own input JSON, the default key name can be overridden with the `terraformKey` argument.

The lifecycle key gets added with subkeys:

* `action` - Action Terraform performs on the resource. Values are `create`, `update`, or `delete`.
* `prev_input` - Input JSON payload from the previous invocation. This can be used to handle update and delete events.

When the resource from the CRUD example above is created, the Lambda will receive the following JSON payload:

```json
{
  "resource_name": "database_setup",
  "database_url": "mydb.cluster-xyz.us-west-2.rds.amazonaws.com:5432",
  "credentials": {
    "username": "admin",
    "password": "secret123"
  },
  "tf": {
    "action": "create",
    "prev_input": null
  }
}
```

If the `databaseUrl` changes, the Lambda will be invoked again with:

```json
{
  "resource_name": "database_setup",
  "database_url": "mydb-new.cluster-abc.us-west-2.rds.amazonaws.com:5432",
  "credentials": {
    "username": "admin",
    "password": "secret123"
  },
  "tf": {
    "action": "update",
    "prev_input": {
      "resource_name": "database_setup",
      "database_url": "mydb.cluster-xyz.us-west-2.rds.amazonaws.com:5432",
      "credentials": {
        "username": "admin",
        "password": "secret123"
      }
    }
  }
}
```

When the invocation resource is removed, the final invocation will have:

```json
{
  "resource_name": "database_setup",
  "database_url": "mydb-new.cluster-abc.us-west-2.rds.amazonaws.com:5432",
  "credentials": {
    "username": "admin",
    "password": "secret123"
  },
  "tf": {
    "action": "delete",
    "prev_input": {
      "resource_name": "database_setup",
      "database_url": "mydb-new.cluster-abc.us-west-2.rds.amazonaws.com:5432",
      "credentials": {
        "username": "admin",
        "password": "secret123"
      }
    }
  }
}
```

## Argument Reference

The following arguments are required:

* `functionName` - (Required) Name of the Lambda function.
* `input` - (Required) JSON payload to the Lambda function.

The following arguments are optional:

* `lifecycleScope` - (Optional) Lifecycle scope of the resource to manage. Valid values are `CREATE_ONLY` and `CRUD`. Defaults to `CREATE_ONLY`. `CREATE_ONLY` will invoke the function only on creation or replacement. `CRUD` will invoke the function on each lifecycle event, and augment the input JSON payload with additional lifecycle information.
* `qualifier` - (Optional) Qualifier (i.e., version) of the Lambda function. Defaults to `$LATEST`.
* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `terraformKey` - (Optional) JSON key used to store lifecycle information in the input JSON payload. Defaults to `tf`. This additional key is only included when `lifecycleScope` is set to `CRUD`.
* `triggers` - (Optional) Map of arbitrary keys and values that, when changed, will trigger a re-invocation. To force a re-invocation without changing these keys/values, use the [`terraform taint` command](https://developer.hashicorp.com/terraform/cli/commands/taint).

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `result` - String result of the Lambda function invocation.

<!-- cache-key: cdktf-0.20.8 input-bb189426b52466acd39d2e4392b430bfe6857c05c18783873936ae5c8895a7c2 -->