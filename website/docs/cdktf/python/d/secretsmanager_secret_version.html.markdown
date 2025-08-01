---
subcategory: "Secrets Manager"
layout: "aws"
page_title: "AWS: aws_secretsmanager_secret_version"
description: |-
  Retrieve information about a Secrets Manager secret version including its secret value
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_secretsmanager_secret_version

Retrieve information about a Secrets Manager secret version, including its secret value. To retrieve secret metadata, see the [`aws_secretsmanager_secret` data source](/docs/providers/aws/d/secretsmanager_secret.html).

## Example Usage

### Retrieve Current Secret Version

By default, this data sources retrieves information based on the `AWSCURRENT` staging label.

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_secretsmanager_secret_version import DataAwsSecretsmanagerSecretVersion
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsSecretsmanagerSecretVersion(self, "secret-version",
            secret_id=Token.as_string(example.id)
        )
```

### Retrieve Specific Secret Version

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_secretsmanager_secret_version import DataAwsSecretsmanagerSecretVersion
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsSecretsmanagerSecretVersion(self, "by-version-stage",
            secret_id=Token.as_string(example.id),
            version_stage="example"
        )
```

### Handling Key-Value Secret Strings in JSON

Reading key-value pairs from JSON back into a native Terraform map can be accomplished in Terraform 0.12 and later with the [`jsondecode()` function](https://www.terraform.io/docs/configuration/functions/jsondecode.html):

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformOutput, Fn, Token, TerraformStack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        TerraformOutput(self, "example",
            value=Fn.lookup_nested(
                Fn.jsondecode(Token.as_string(example.secret_string)), ["\"key1\""])
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `secret_id` - (Required) Specifies the secret containing the version that you want to retrieve. You can specify either the ARN or the friendly name of the secret.
* `version_id` - (Optional) Specifies the unique identifier of the version of the secret that you want to retrieve. Overrides `version_stage`.
* `version_stage` - (Optional) Specifies the secret version that you want to retrieve by the staging label attached to the version. Defaults to `AWSCURRENT`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the secret.
* `created_date` - Created date of the secret in UTC.
* `id` - Unique identifier of this version of the secret.
* `secret_string` - Decrypted part of the protected secret information that was originally provided as a string.
* `secret_binary` - Decrypted part of the protected secret information that was originally provided as a binary.
* `version_id` - Unique identifier of this version of the secret.

<!-- cache-key: cdktf-0.20.8 input-a07848bd8d2d1006ae852e0eda3bdf0f366fa0d55c1bcd7bad0d4870afb4175b -->