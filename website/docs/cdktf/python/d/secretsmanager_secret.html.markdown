---
subcategory: "Secrets Manager"
layout: "aws"
page_title: "AWS: aws_secretsmanager_secret"
description: |-
  Retrieve metadata information about a Secrets Manager secret
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_secretsmanager_secret

Retrieve metadata information about a Secrets Manager secret. To retrieve a secret value, see the [`aws_secretsmanager_secret_version` data source](/docs/providers/aws/d/secretsmanager_secret_version.html).

## Example Usage

### ARN

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_secretsmanager_secret import DataAwsSecretsmanagerSecret
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsSecretsmanagerSecret(self, "by-arn",
            arn="arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456"
        )
```

### Name

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_secretsmanager_secret import DataAwsSecretsmanagerSecret
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsSecretsmanagerSecret(self, "by-name",
            name="example"
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `arn` - (Optional) ARN of the secret to retrieve.
* `name` - (Optional) Name of the secret to retrieve.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the secret.
* `created_date` - Created date of the secret in UTC.
* `description` - Description of the secret.
* `kms_key_id` - Key Management Service (KMS) Customer Master Key (CMK) associated with the secret.
* `id` - ARN of the secret.
* `last_changed_date` - Last updated date of the secret in UTC.
* `policy` - Resource-based policy document that's attached to the secret.
* `tags` - Tags of the secret.

<!-- cache-key: cdktf-0.20.8 input-38c85767e85f8518c95a3e8cded6487df7c19d31be18afc5dce39e7ae3153c3a -->