---
subcategory: "Managed Streaming for Kafka"
layout: "aws"
page_title: "AWS: aws_msk_replicator"
description: |-
  Terraform resource for managing an AWS Managed Streaming for Kafka Replicator.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_msk_replicator

Terraform resource for managing an AWS Managed Streaming for Kafka Replicator.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { MskReplicator } from "./.gen/providers/aws/msk-replicator";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new MskReplicator(this, "test", {
      description: "test-description",
      kafkaCluster: [
        {
          amazonMskCluster: {
            mskClusterArn: source.arn,
          },
          vpcConfig: {
            securityGroupsIds: [Token.asString(awsSecurityGroupSource.id)],
            subnetIds: Token.asList(
              Fn.lookupNested(awsSubnetSource, ["*", "id"])
            ),
          },
        },
        {
          amazonMskCluster: {
            mskClusterArn: target.arn,
          },
          vpcConfig: {
            securityGroupsIds: [Token.asString(awsSecurityGroupTarget.id)],
            subnetIds: Token.asList(
              Fn.lookupNested(awsSubnetTarget, ["*", "id"])
            ),
          },
        },
      ],
      replicationInfoList: {
        consumerGroupReplication: [
          {
            consumerGroupsToReplicate: [".*"],
          },
        ],
        sourceKafkaClusterArn: source.arn,
        targetCompressionType: "NONE",
        targetKafkaClusterArn: target.arn,
        topicReplication: [
          {
            startingPosition: {
              type: "LATEST",
            },
            topicNameConfiguration: {
              type: "PREFIXED_WITH_SOURCE_CLUSTER_ALIAS",
            },
            topicsToReplicate: [".*"],
          },
        ],
      },
      replicatorName: "test-name",
      serviceExecutionRoleArn: Token.asString(awsIamRoleSource.arn),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `replicatorName` - (Required) The name of the replicator.
* `kafkaCluster` - (Required) A list of Kafka clusters which are targets of the replicator.
* `serviceExecutionRoleArn` - (Required) The ARN of the IAM role used by the replicator to access resources in the customer's account (e.g source and target clusters).
* `replicationInfoList` - (Required) A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
* `description` - (Optional) A summary description of the replicator.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### kafka_cluster Argument Reference

* `amazonMskCluster` - (Required) Details of an Amazon MSK cluster.
* `vpcConfig` - (Required) Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.

### amazon_msk_cluster Argument Reference

* `mskClusterArn` - (Required) The ARN of an Amazon MSK cluster.

### vpc_config Argument Reference

* `subnetIds` - (Required) The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets to allow communication between your Kafka Cluster and the replicator.
* `securityGroupsIds` - (Required) The AWS security groups to associate with the ENIs used by the replicator. If a security group is not specified, the default security group associated with the VPC is used.

### replication_info_list Argument Reference

* `sourceKafkaClusterArn` - (Required) The ARN of the source Kafka cluster.
* `targetKafkaClusterArn` - (Required) The ARN of the target Kafka cluster.
* `targetCompressionType` - (Required) The type of compression to use writing records to target Kafka cluster.
* `topicReplication` - (Required) Configuration relating to topic replication.
* `consumerGroupReplication` - (Required) Configuration relating to consumer group replication.

### topic_replication Argument Reference

* `topicNameConfiguration` - (Optional) Configuration for specifying replicated topic names should be the same as their corresponding upstream topics or prefixed with source cluster alias.
* `topicsToReplicate` - (Required) List of regular expression patterns indicating the topics to copy.
* `topicsToExclude` - (Optional) List of regular expression patterns indicating the topics that should not be replica.
* `detectAndCopyNewTopics` - (Optional) Whether to periodically check for new topics and partitions.
* `copyAccessControlListsForTopics` - (Optional) Whether to periodically configure remote topic ACLs to match their corresponding upstream topics.
* `copyTopicConfigurations` - (Optional) Whether to periodically configure remote topics to match their corresponding upstream topics.
* `startingPosition` - (Optional) Configuration for specifying the position in the topics to start replicating from.

### consumer_group_replication Argument Reference

* `consumerGroupsToReplicate` - (Required) List of regular expression patterns indicating the consumer groups to copy.
* `consumerGroupsToExclude` - (Optional) List of regular expression patterns indicating the consumer groups that should not be replicated.
* `detectAndCopyNewConsumerGroups` - (Optional) Whether to periodically check for new consumer groups.
* `synchroniseConsumerGroupOffsets` - (Optional) Whether to periodically write the translated offsets to __consumer_offsets topic in target cluster.

### topic_name_configuration

* `type` - (optional) The type of topic configuration name. Supports `PREFIXED_WITH_SOURCE_CLUSTER_ALIAS` and `IDENTICAL`.

### starting_position

* `type` - (Optional) The type of replication starting position. Supports `LATEST` and `EARLIEST`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Replicator.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `60m`)
* `update` - (Default `180m`)
* `delete` - (Default `90m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import MSK replicators using the replicator ARN. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { MskReplicator } from "./.gen/providers/aws/msk-replicator";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    MskReplicator.generateConfigForImport(
      this,
      "example",
      "arn:aws:kafka:us-west-2:123456789012:configuration/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3"
    );
  }
}

```

Using `terraform import`, import MSK replicators using the replicator ARN. For example:

```console
% terraform import aws_msk_replicator.example arn:aws:kafka:us-west-2:123456789012:configuration/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
```

<!-- cache-key: cdktf-0.20.8 input-8bf4594598d13dce6022640b42dd1f6965ce218f5432cc4d9d9f97068e0b2ddd -->