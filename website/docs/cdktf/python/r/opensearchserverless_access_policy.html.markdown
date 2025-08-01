---
subcategory: "OpenSearch Serverless"
layout: "aws"
page_title: "AWS: aws_opensearchserverless_access_policy"
description: |-
  Terraform resource for managing an AWS OpenSearch Serverless Access Policy.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_opensearchserverless_access_policy

Terraform resource for managing an AWS OpenSearch Serverless Access Policy. See AWS documentation for [data access policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html) and [supported data access policy permissions](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html#serverless-data-supported-permissions).

## Example Usage

### Grant all collection and index permissions

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_caller_identity import DataAwsCallerIdentity
from imports.aws.opensearchserverless_access_policy import OpensearchserverlessAccessPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        current = DataAwsCallerIdentity(self, "current")
        OpensearchserverlessAccessPolicy(self, "example",
            description="read and write permissions",
            name="example",
            policy=Token.as_string(
                Fn.jsonencode([{
                    "Principal": [current.arn],
                    "Rules": [{
                        "Permission": ["aoss:*"],
                        "Resource": ["index/example-collection/*"],
                        "ResourceType": "index"
                    }, {
                        "Permission": ["aoss:*"],
                        "Resource": ["collection/example-collection"],
                        "ResourceType": "collection"
                    }
                    ]
                }
                ])),
            type="data"
        )
```

### Grant read-only collection and index permissions

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_caller_identity import DataAwsCallerIdentity
from imports.aws.opensearchserverless_access_policy import OpensearchserverlessAccessPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        current = DataAwsCallerIdentity(self, "current")
        OpensearchserverlessAccessPolicy(self, "example",
            description="read-only permissions",
            name="example",
            policy=Token.as_string(
                Fn.jsonencode([{
                    "Principal": [current.arn],
                    "Rules": [{
                        "Permission": ["aoss:DescribeIndex", "aoss:ReadDocument"],
                        "Resource": ["index/example-collection/*"],
                        "ResourceType": "index"
                    }, {
                        "Permission": ["aoss:DescribeCollectionItems"],
                        "Resource": ["collection/example-collection"],
                        "ResourceType": "collection"
                    }
                    ]
                }
                ])),
            type="data"
        )
```

### Grant SAML identity permissions

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.opensearchserverless_access_policy import OpensearchserverlessAccessPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        OpensearchserverlessAccessPolicy(self, "example",
            description="saml permissions",
            name="example",
            policy=Token.as_string(
                Fn.jsonencode([{
                    "Principal": ["saml/123456789012/myprovider/user/Annie", "saml/123456789012/anotherprovider/group/Accounting"
                    ],
                    "Rules": [{
                        "Permission": ["aoss:*"],
                        "Resource": ["index/example-collection/*"],
                        "ResourceType": "index"
                    }, {
                        "Permission": ["aoss:*"],
                        "Resource": ["collection/example-collection"],
                        "ResourceType": "collection"
                    }
                    ]
                }
                ])),
            type="data"
        )
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the policy.
* `policy` - (Required) JSON policy document to use as the content for the new policy
* `type` - (Required) Type of access policy. Must be `data`.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `description` - (Optional) Description of the policy. Typically used to store information about the permissions defined in the policy.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `policy_version` - Version of the policy.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import OpenSearchServerless Access Policy using the `name` and `type` arguments separated by a slash (`/`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.opensearchserverless_access_policy import OpensearchserverlessAccessPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        OpensearchserverlessAccessPolicy.generate_config_for_import(self, "example", "example/data")
```

Using `terraform import`, import OpenSearchServerless Access Policy using the `name` and `type` arguments separated by a slash (`/`). For example:

```console
% terraform import aws_opensearchserverless_access_policy.example example/data
```

<!-- cache-key: cdktf-0.20.8 input-79b73f7bbd824877f38f16c1a540c4b05a5807db4052e3b10f6a16c8844610da -->