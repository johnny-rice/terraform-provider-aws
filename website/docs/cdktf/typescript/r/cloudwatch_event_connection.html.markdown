---
subcategory: "EventBridge"
layout: "aws"
page_title: "AWS: aws_cloudwatch_event_connection"
description: |-
  Provides an EventBridge connection resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_cloudwatch_event_connection

Provides an EventBridge connection resource.

~> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CloudwatchEventConnection } from "./.gen/providers/aws/cloudwatch-event-connection";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new CloudwatchEventConnection(this, "test", {
      authParameters: {
        apiKey: {
          key: "x-signature",
          value: "1234",
        },
      },
      authorizationType: "API_KEY",
      description: "A connection description",
      name: "ngrok-connection",
    });
  }
}

```

## Example Usage Basic Authorization

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CloudwatchEventConnection } from "./.gen/providers/aws/cloudwatch-event-connection";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new CloudwatchEventConnection(this, "test", {
      authParameters: {
        basic: {
          password: "Pass1234!",
          username: "user",
        },
      },
      authorizationType: "BASIC",
      description: "A connection description",
      name: "ngrok-connection",
    });
  }
}

```

## Example Usage OAuth Authorization

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CloudwatchEventConnection } from "./.gen/providers/aws/cloudwatch-event-connection";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new CloudwatchEventConnection(this, "test", {
      authParameters: {
        oauth: {
          authorizationEndpoint: "https://auth.url.com/endpoint",
          clientParameters: {
            clientId: "1234567890",
            clientSecret: "Pass1234!",
          },
          httpMethod: "GET",
          oauthHttpParameters: {
            body: [
              {
                isValueSecret: false,
                key: "body-parameter-key",
                value: "body-parameter-value",
              },
            ],
            header: [
              {
                isValueSecret: false,
                key: "header-parameter-key",
                value: "header-parameter-value",
              },
            ],
            queryString: [
              {
                isValueSecret: false,
                key: "query-string-parameter-key",
                value: "query-string-parameter-value",
              },
            ],
          },
        },
      },
      authorizationType: "OAUTH_CLIENT_CREDENTIALS",
      description: "A connection description",
      name: "ngrok-connection",
    });
  }
}

```

## Example Usage Invocation Http Parameters

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CloudwatchEventConnection } from "./.gen/providers/aws/cloudwatch-event-connection";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new CloudwatchEventConnection(this, "test", {
      authParameters: {
        basic: {
          password: "Pass1234!",
          username: "user",
        },
        invocationHttpParameters: {
          body: [
            {
              isValueSecret: false,
              key: "body-parameter-key",
              value: "body-parameter-value",
            },
            {
              isValueSecret: true,
              key: "body-parameter-key2",
              value: "body-parameter-value2",
            },
          ],
          header: [
            {
              isValueSecret: false,
              key: "header-parameter-key",
              value: "header-parameter-value",
            },
          ],
          queryString: [
            {
              isValueSecret: false,
              key: "query-string-parameter-key",
              value: "query-string-parameter-value",
            },
          ],
        },
      },
      authorizationType: "BASIC",
      description: "A connection description",
      name: "ngrok-connection",
    });
  }
}

```

## Example Usage CMK Encryption

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CloudwatchEventConnection } from "./.gen/providers/aws/cloudwatch-event-connection";
import { DataAwsCallerIdentity } from "./.gen/providers/aws/data-aws-caller-identity";
import { DataAwsPartition } from "./.gen/providers/aws/data-aws-partition";
import { KmsKey } from "./.gen/providers/aws/kms-key";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new CloudwatchEventConnection(this, "test", {
      authParameters: {
        basic: {
          password: "Pass1234!",
          username: "user",
        },
      },
      authorizationType: "BASIC",
      description: "A connection description",
      kmsKeyIdentifier: example.id,
      name: "ngrok-connection",
    });
    const current = new DataAwsCallerIdentity(this, "current", {});
    const dataAwsPartitionCurrent = new DataAwsPartition(this, "current_2", {});
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsPartitionCurrent.overrideLogicalId("current");
    const awsKmsKeyTest = new KmsKey(this, "test_3", {
      deletionWindowInDays: 7,
      policy: Token.asString(
        Fn.jsonencode({
          Id: "key-policy-example",
          Statement: [
            {
              Action: "kms:*",
              Effect: "Allow",
              Principal: {
                AWS:
                  "arn:${" +
                  dataAwsPartitionCurrent.partition +
                  "}:iam::${" +
                  current.accountId +
                  "}:root",
              },
              Resource: "*",
              Sid: "Enable IAM User Permissions",
            },
            {
              Action: ["kms:DescribeKey", "kms:Decrypt", "kms:GenerateDataKey"],
              Condition: {
                StringLike: {
                  "kms:EncryptionContext:SecretARN": [
                    "arn:aws:secretsmanager:*:*:secret:events!connection/*",
                  ],
                  "kms:ViaService": "secretsmanager.*.amazonaws.com",
                },
              },
              Effect: "Allow",
              Principal: {
                AWS:
                  "arn:${" +
                  dataAwsPartitionCurrent.partition +
                  "}:iam::${" +
                  current.accountId +
                  "}:root",
              },
              Resource: "*",
              Sid: "Allow use of the key",
            },
          ],
          Version: "2012-10-17",
        })
      ),
      tags: {
        EventBridgeApiDestinations: "true",
      },
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsKmsKeyTest.overrideLogicalId("test");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name for the connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
* `description` - (Optional) Description for the connection. Maximum of 512 characters.
* `authorizationType` - (Required) Type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
* `authParameters` - (Required) Parameters used for authorization. A maximum of 1 are allowed. Documented below.
* `invocationConnectivityParameters` - (Optional) Parameters to use for invoking a private API. Documented below.
* `kmsKeyIdentifier` - (Optional) Identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this connection. The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.

`authParameters` support the following:

* `apiKey` - (Optional) Parameters used for API_KEY authorization. An API key to include in the header for each authentication request. A maximum of 1 are allowed. Conflicts with `basic` and `oauth`. Documented below.
* `basic` - (Optional) Parameters used for BASIC authorization. A maximum of 1 are allowed. Conflicts with `apiKey` and `oauth`. Documented below.
* `invocationHttpParameters` - (Optional) Invocation Http Parameters are additional credentials used to sign each Invocation of the ApiDestination created from this Connection. If the ApiDestination Rule Target has additional HttpParameters, the values will be merged together, with the Connection Invocation Http Parameters taking precedence. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.
* `oauth` - (Optional) Parameters used for OAUTH_CLIENT_CREDENTIALS authorization. A maximum of 1 are allowed. Conflicts with `basic` and `apiKey`. Documented below.

`apiKey` support the following:

* `key` - (Required) Header Name.
* `value` - (Required) Header Value. Created and stored in AWS Secrets Manager.

`basic` support the following:

* `username` - (Required) A username for the authorization.
* `password` - (Required) A password for the authorization. Created and stored in AWS Secrets Manager.

`oauth` support the following:

* `authorizationEndpoint` - (Required) The URL to the authorization endpoint.
* `httpMethod` - (Required) A password for the authorization. Created and stored in AWS Secrets Manager.
* `clientParameters` - (Required) Contains the client parameters for OAuth authorization. Contains the following two parameters.
    * `clientId` - (Required) The client ID for the credentials to use for authorization. Created and stored in AWS Secrets Manager.
    * `clientSecret` - (Required) The client secret for the credentials to use for authorization. Created and stored in AWS Secrets Manager.
* `oauthHttpParameters` - (Required) OAuth Http Parameters are additional credentials used to sign the request to the authorization endpoint to exchange the OAuth Client information for an access token. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.

`invocationHttpParameters` and `oauthHttpParameters` support the following:

* `body` - (Optional) Contains additional body string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
    * `key` - (Required) The key for the parameter.
    * `value` - (Required) The value associated with the key. Created and stored in AWS Secrets Manager if is secret.
    * `isValueSecret` - (Optional) Specified whether the value is secret.

* `header` - (Optional) Contains additional header parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
    * `key` - (Required) The key for the parameter.
    * `value` - (Required) The value associated with the key. Created and stored in AWS Secrets Manager if is secret.
    * `isValueSecret` - (Optional) Specified whether the value is secret.

* `queryString` - (Optional) Contains additional query string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
    * `key` - (Required) The key for the parameter.
    * `value` - (Required) The value associated with the key. Created and stored in AWS Secrets Manager if is secret.
    * `isValueSecret` - (Optional) Specified whether the value is secret.

`invocationConnectivityParameters` supports the following:

* `resourceParameters` - (Required) The parameters for EventBridge to use when invoking the resource endpoint. Documented below.

`resourceParameters` supports the following:

* `resourceConfigurationArn` - (Required) ARN of the Amazon VPC Lattice [resource configuration](vpclattice_resource_configuration) for the resource endpoint.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) of the connection.
* `secretArn` - The Amazon Resource Name (ARN) of the secret created from the authorization parameters specified for the connection.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import EventBridge connection using the `name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CloudwatchEventConnection } from "./.gen/providers/aws/cloudwatch-event-connection";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    CloudwatchEventConnection.generateConfigForImport(
      this,
      "test",
      "ngrok-connection"
    );
  }
}

```

Using `terraform import`, import EventBridge EventBridge connection using the `name`. For example:

```console
% terraform import aws_cloudwatch_event_connection.test ngrok-connection
```

<!-- cache-key: cdktf-0.20.8 input-89571db5a946668df93cfaee7902913652ae1dd15f0f4ecd31915846dd20405b -->