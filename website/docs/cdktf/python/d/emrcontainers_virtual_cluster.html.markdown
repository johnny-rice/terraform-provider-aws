---
subcategory: "EMR Containers"
layout: "aws"
page_title: "AWS: aws_emrcontainers_virtual_cluster"
description: |-
  Retrieve information about an EMR Containers (EMR on EKS) Virtual Cluster
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_emrcontainers_virtual_cluster

Retrieve information about an EMR Containers (EMR on EKS) Virtual Cluster.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformOutput, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_emrcontainers_virtual_cluster import DataAwsEmrcontainersVirtualCluster
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = DataAwsEmrcontainersVirtualCluster(self, "example",
            virtual_cluster_id="example id"
        )
        TerraformOutput(self, "arn",
            value=example.arn
        )
        TerraformOutput(self, "name",
            value=example.name
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `virtual_cluster_id` - (Required) ID of the cluster.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - ID of the cluster.
* `name` - Name of the cluster.
* `arn` - ARN of the cluster.
* `container_provider` - Nested attribute containing information about the underlying container provider (EKS cluster) for your EMR Containers cluster.
    * `id` - The name of the container provider that is running your EMR Containers cluster
    * `info` - Nested list containing information about the configuration of the container provider
        * `eks_info` - Nested list containing EKS-specific information about the cluster where the EMR Containers cluster is running
            * `namespace` - The namespace where the EMR Containers cluster is running
    * `type` - The type of the container provider
* `created_at` - Unix epoch time stamp in seconds for when the cluster was created.
* `state` - Status of the EKS cluster. One of `RUNNING`, `TERMINATING`, `TERMINATED`, `ARRESTED`.
* `tags` - Key-value mapping of resource tags.

<!-- cache-key: cdktf-0.20.8 input-493ee07bc735014de653e26defab56c8d780c161054096685641402620f6eae1 -->