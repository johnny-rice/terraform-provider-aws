---
subcategory: "Bedrock Agents"
layout: "aws"
page_title: "AWS: aws_bedrockagent_knowledge_base"
description: |-
  Terraform resource for managing an AWS Agents for Amazon Bedrock Knowledge Base.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_bedrockagent_knowledge_base

Terraform resource for managing an AWS Agents for Amazon Bedrock Knowledge Base.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.bedrockagent_knowledge_base import BedrockagentKnowledgeBase
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BedrockagentKnowledgeBase(self, "example",
            knowledge_base_configuration=[BedrockagentKnowledgeBaseKnowledgeBaseConfiguration(
                type="VECTOR",
                vector_knowledge_base_configuration=[BedrockagentKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfiguration(
                    embedding_model_arn="arn:aws:bedrock:us-west-2::foundation-model/amazon.titan-embed-text-v2:0"
                )
                ]
            )
            ],
            name="example",
            role_arn=Token.as_string(aws_iam_role_example.arn),
            storage_configuration=[BedrockagentKnowledgeBaseStorageConfiguration(
                opensearch_serverless_configuration=[BedrockagentKnowledgeBaseStorageConfigurationOpensearchServerlessConfiguration(
                    collection_arn="arn:aws:aoss:us-west-2:123456789012:collection/142bezjddq707i5stcrf",
                    field_mapping=[BedrockagentKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMapping(
                        metadata_field="AMAZON_BEDROCK_METADATA",
                        text_field="AMAZON_BEDROCK_TEXT_CHUNK",
                        vector_field="bedrock-knowledge-base-default-vector"
                    )
                    ],
                    vector_index_name="bedrock-knowledge-base-default-index"
                )
                ],
                type="OPENSEARCH_SERVERLESS"
            )
            ]
        )
```

### With Supplemental Data Storage Configuration

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.bedrockagent_knowledge_base import BedrockagentKnowledgeBase
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BedrockagentKnowledgeBase(self, "example",
            knowledge_base_configuration=[BedrockagentKnowledgeBaseKnowledgeBaseConfiguration(
                type="VECTOR",
                vector_knowledge_base_configuration=[BedrockagentKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfiguration(
                    embedding_model_arn="arn:aws:bedrock:us-west-2::foundation-model/amazon.titan-embed-text-v2:0",
                    embedding_model_configuration=[BedrockagentKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationEmbeddingModelConfiguration(
                        bedrock_embedding_model_configuration=[BedrockagentKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationEmbeddingModelConfigurationBedrockEmbeddingModelConfiguration(
                            dimensions=1024,
                            embedding_data_type="FLOAT32"
                        )
                        ]
                    )
                    ],
                    supplemental_data_storage_configuration=[BedrockagentKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationSupplementalDataStorageConfiguration(
                        storage_location=[BedrockagentKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationSupplementalDataStorageConfigurationStorageLocation(
                            s3_location=[BedrockagentKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationSupplementalDataStorageConfigurationStorageLocationS3Location(
                                uri="s3://my-bucket/chunk-processor/"
                            )
                            ],
                            type="S3"
                        )
                        ]
                    )
                    ]
                )
                ]
            )
            ],
            name="example",
            role_arn=Token.as_string(aws_iam_role_example.arn),
            storage_configuration=[BedrockagentKnowledgeBaseStorageConfiguration(
                opensearch_serverless_configuration=[BedrockagentKnowledgeBaseStorageConfigurationOpensearchServerlessConfiguration(
                    collection_arn="arn:aws:aoss:us-west-2:123456789012:collection/142bezjddq707i5stcrf",
                    field_mapping=[BedrockagentKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMapping(
                        metadata_field="AMAZON_BEDROCK_METADATA",
                        text_field="AMAZON_BEDROCK_TEXT_CHUNK",
                        vector_field="bedrock-knowledge-base-default-vector"
                    )
                    ],
                    vector_index_name="bedrock-knowledge-base-default-index"
                )
                ],
                type="OPENSEARCH_SERVERLESS"
            )
            ]
        )
```

## Argument Reference

The following arguments are required:

* `knowledge_base_configuration` - (Required, Forces new resource) Details about the embeddings configuration of the knowledge base. See [`knowledge_base_configuration` block](#knowledge_base_configuration-block) for details.
* `name` - (Required) Name of the knowledge base.
* `role_arn` - (Required) ARN of the IAM role with permissions to invoke API operations on the knowledge base.
* `storage_configuration` - (Required, Forces new resource) Details about the storage configuration of the knowledge base. See [`storage_configuration` block](#storage_configuration-block) for details.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `description` - (Optional) Description of the knowledge base.
* `tags` - (Optional) Map of tags assigned to the resource. If configured with a provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### `knowledge_base_configuration` block

The `knowledge_base_configuration` configuration block supports the following arguments:

* `type` - (Required) Type of data that the data source is converted into for the knowledge base. Valid Values: `VECTOR`.
* `vector_knowledge_base_configuration` - (Optional) Details about the embeddings model that'sused to convert the data source. See [`vector_knowledge_base_configuration` block](#vector_knowledge_base_configuration-block) for details.

### `vector_knowledge_base_configuration` block

The `vector_knowledge_base_configuration` configuration block supports the following arguments:

* `embedding_model_arn` - (Required) ARN of the model used to create vector embeddings for the knowledge base.
* `embedding_model_configuration` - (Optional) The embeddings model configuration details for the vector model used in Knowledge Base.  See [`embedding_model_configuration` block](#embedding_model_configuration-block) for details.
* `supplemental_data_storage_configuration` - (Optional) supplemental_data_storage_configuration.  See [`supplemental_data_storage_configuration` block](#supplemental_data_storage_configuration-block) for details.

### `embedding_model_configuration` block

The `embedding_model_configuration` configuration block supports the following arguments:

* `bedrock_embedding_model_configuration` - (Optional) The vector configuration details on the Bedrock embeddings model.  See [`bedrock_embedding_model_configuration` block](#bedrock_embedding_model_configuration-block) for details.

### `bedrock_embedding_model_configuration` block

The `bedrock_embedding_model_configuration` configuration block supports the following arguments:

* `dimensions` - (Optional) Dimension details for the vector configuration used on the Bedrock embeddings model.
* `embedding_data_type` - (Optional) Data type for the vectors when using a model to convert text into vector embeddings. The model must support the specified data type for vector embeddings.  Valid values are `FLOAT32` and `BINARY`.

### `supplemental_data_storage_configuration` block

The `supplemental_data_storage_configuration` configuration block supports the following arguments:

* `storage_location` - (Required) A storage location specification for images extracted from multimodal documents in your data source.  See [`storage_location` block](#storage_location-block) for details.

### `storage_location` block

The `storage_location` configuration block supports the following arguments:

* `type` - (Required) Storage service used for this location. `S3` is the only valid value.
* `s3_location` - (Optional) Contains information about the Amazon S3 location for the extracted images.  See [`s3_location` block](#s3_location-block) for details.

### `s3_location` block

The `s3_location` configuration block supports the following arguments:

* `uri` - (Required) URI of the location.

### `storage_configuration` block

The `storage_configuration` configuration block supports the following arguments:

* `type` - (Required) Vector store service in which the knowledge base is stored. Valid Values: `OPENSEARCH_SERVERLESS`, `PINECONE`, `REDIS_ENTERPRISE_CLOUD`, `RDS`.
* `opensearch_serverless_configuration` - (Optional) The storage configuration of the knowledge base in Amazon OpenSearch Service. See [`opensearch_serverless_configuration` block](#opensearch_serverless_configuration-block) for details.
* `pinecone_configuration` - (Optional)  The storage configuration of the knowledge base in Pinecone. See [`pinecone_configuration` block](#pinecone_configuration-block) for details.
* `rds_configuration` - (Optional) Details about the storage configuration of the knowledge base in Amazon RDS. For more information, see [Create a vector index in Amazon RDS](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup.html). See [`rds_configuration` block](#rds_configuration-block) for details.
* `redis_enterprise_cloud_configuration` - (Optional) The storage configuration of the knowledge base in Redis Enterprise Cloud. See [`redis_enterprise_cloud_configuration` block](#redis_enterprise_cloud_configuration-block) for details.

### `opensearch_serverless_configuration` block

The `opensearch_serverless_configuration` configuration block supports the following arguments:

* `collection_arn` - (Required) ARN of the OpenSearch Service vector store.
* `field_mapping` - (Required) The names of the fields to which to map information about the vector store. This block supports the following arguments:
    * `metadata_field` - (Required) Name of the field in which Amazon Bedrock stores metadata about the vector store.
    * `text_field` - (Required) Name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.
    * `vector_field` - (Required) Name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.
* `vector_index_name` - (Required) Name of the vector store.

### `pinecone_configuration` block

The `pinecone_configuration` configuration block supports the following arguments:

* `connection_string` - (Required) Endpoint URL for your index management page.
* `credentials_secret_arn` - (Required) ARN of the secret that you created in AWS Secrets Manager that is linked to your Pinecone API key.
* `field_mapping` - (Required) The names of the fields to which to map information about the vector store. This block supports the following arguments:
    * `metadata_field` - (Required) Name of the field in which Amazon Bedrock stores metadata about the vector store.
    * `text_field` - (Required) Name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.
* `namespace` - (Optional) Namespace to be used to write new data to your database.

### `rds_configuration` block

 The `rds_configuration` configuration block supports the following arguments:

* `credentials_secret_arn` - (Required) ARN of the secret that you created in AWS Secrets Manager that is linked to your Amazon RDS database.
* `database_name` - (Required) Name of your Amazon RDS database.
* `field_mapping` - (Required) Names of the fields to which to map information about the vector store. This block supports the following arguments:
    * `metadata_field` - (Required) Name of the field in which Amazon Bedrock stores metadata about the vector store.
    * `primary_key_field` - (Required) Name of the field in which Amazon Bedrock stores the ID for each entry.
    * `text_field` - (Required) Name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.
    * `vector_field` - (Required) Name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.
* `resource_arn` - (Required) ARN of the vector store.
* `table_name` - (Required) Name of the table in the database.

### `redis_enterprise_cloud_configuration` block

The `redis_enterprise_cloud_configuration` configuration block supports the following arguments:

* `credentials_secret_arn` - (Required) ARN of the secret that you created in AWS Secrets Manager that is linked to your Redis Enterprise Cloud database.
* `endpoint` - (Required) Endpoint URL of the Redis Enterprise Cloud database.
* `field_mapping` - (Required) The names of the fields to which to map information about the vector store. This block supports the following arguments:
    * `metadata_field` - (Required) Name of the field in which Amazon Bedrock stores metadata about the vector store.
    * `text_field` - (Required) Name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.
    * `vector_field` - (Required) Name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.
* `vector_index_name` - (Required) Name of the vector index.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the knowledge base.
* `created_at` - Time at which the knowledge base was created.
* `id` - Unique identifier of the knowledge base.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).
* `updated_at` - Time at which the knowledge base was last updated.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Agents for Amazon Bedrock Knowledge Base using the knowledge base ID. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.bedrockagent_knowledge_base import BedrockagentKnowledgeBase
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BedrockagentKnowledgeBase.generate_config_for_import(self, "example", "EMDPPAYPZI")
```

Using `terraform import`, import Agents for Amazon Bedrock Knowledge Base using the knowledge base ID. For example:

```console
% terraform import aws_bedrockagent_knowledge_base.example EMDPPAYPZI
```

<!-- cache-key: cdktf-0.20.8 input-0e03b5b72c6d03d9e151bacea847398de5d7dc3b9f8939d76792e3bc7dc540bd -->