---
subcategory: "KMS (Key Management)"
layout: "aws"
page_title: "AWS: aws_kms_custom_key_store"
description: |-
  Get information on a AWS Key Management Service (KMS) Custom Key Store
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_kms_custom_key_store

Use this data source to get the metadata KMS custom key store.
By using this data source, you can reference KMS custom key store
without having to hard code the ID as input.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsKmsCustomKeyStore } from "./.gen/providers/aws/data-aws-kms-custom-key-store";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsKmsCustomKeyStore(this, "keystore", {
      customKeyStoreName: "my_cloudhsm",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `customKeyStoreId` - (Optional) The ID for the custom key store.
* `customKeyStoreName` - (Optional) The user-specified friendly name for the custom key store.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - The ID for the custom key store.
* `cloudhsm_cluster_id` - ID for the CloudHSM cluster that is associated with the custom key store.
* `connectionState` - Indicates whether the custom key store is connected to its CloudHSM cluster.
* `creationDate` - The date and time when the custom key store was created.
* `trustAnchorCertificate` - The trust anchor certificate of the associated CloudHSM cluster.

<!-- cache-key: cdktf-0.20.8 input-521afc4759abd94646fb15395b70fca2daf4d26406d5273ead2d288d6aa13682 -->