---
subcategory: "ElastiCache"
layout: "aws"
page_title: "AWS: aws_elasticache_user"
description: |-
  Provides an ElastiCache user.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_elasticache_user

Provides an ElastiCache user resource.

~> **Note:** All arguments including the username and passwords will be stored in the raw state as plain-text.
[Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ElasticacheUser } from "./.gen/providers/aws/elasticache-user";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ElasticacheUser(this, "test", {
      accessString:
        "on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
      engine: "redis",
      passwords: ["password123456789"],
      userId: "testUserId",
      userName: "testUserName",
    });
  }
}

```

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ElasticacheUser } from "./.gen/providers/aws/elasticache-user";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ElasticacheUser(this, "test", {
      accessString: "on ~* +@all",
      authenticationMode: {
        type: "iam",
      },
      engine: "redis",
      userId: "testUserId",
      userName: "testUserName",
    });
  }
}

```

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ElasticacheUser } from "./.gen/providers/aws/elasticache-user";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ElasticacheUser(this, "test", {
      accessString: "on ~* +@all",
      authenticationMode: {
        passwords: ["password1", "password2"],
        type: "password",
      },
      engine: "redis",
      userId: "testUserId",
      userName: "testUserName",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `accessString` - (Required) Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
* `engine` - (Required) The current supported values are `redis`, `valkey` (case insensitive).
* `userId` - (Required) The ID of the user.
* `userName` - (Required) The username of the user.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `authenticationMode` - (Optional) Denotes the user's authentication properties. Detailed below.
* `noPasswordRequired` - (Optional) Indicates a password is not required for this user.
* `passwords` - (Optional) Passwords used for this user. You can create up to two passwords for each user.
* `tags` - (Optional) A list of tags to be added to this resource. A tag is a key-value pair.

### authentication_mode Configuration Block

* `passwords` - (Optional) Specifies the passwords to use for authentication if `type` is set to `password`.
* `type` - (Required) Specifies the authentication type. Possible options are: `password`, `no-password-required` or `iam`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The ARN of the created ElastiCache User.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `5m`)
- `read` - (Default `5m`)
- `update` - (Default `5m`)
- `delete` - (Default `5m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ElastiCache users using the `userId`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ElasticacheUser } from "./.gen/providers/aws/elasticache-user";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ElasticacheUser.generateConfigForImport(this, "myUser", "userId1");
  }
}

```

Using `terraform import`, import ElastiCache users using the `userId`. For example:

```console
% terraform import aws_elasticache_user.my_user userId1
```

<!-- cache-key: cdktf-0.20.8 input-6c06a82050e24759a7072986f5befc6084861761f3a2f07716ea62a32b75a1c4 -->