---
subcategory: "ECR Public"
layout: "aws"
page_title: "AWS: aws_ecrpublic_repository"
description: |-
  Provides a Public Elastic Container Registry Repository.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ecrpublic_repository

Provides a Public Elastic Container Registry Repository.

~> **NOTE:** This resource can only be used in the `us-east-1` region.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcrpublicRepository } from "./.gen/providers/aws/ecrpublic-repository";
import { AwsProvider } from "./.gen/providers/aws/provider";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const usEast1 = new AwsProvider(this, "aws", {
      alias: "us_east_1",
      region: "us-east-1",
    });
    new EcrpublicRepository(this, "foo", {
      catalogData: {
        aboutText: "About Text",
        architectures: ["ARM"],
        description: "Description",
        logoImageBlob: Token.asString(Fn.filebase64(png)),
        operatingSystems: ["Linux"],
        usageText: "Usage Text",
      },
      provider: usEast1,
      repositoryName: "bar",
      tags: {
        env: "production",
      },
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `repositoryName` - (Required) Name of the repository.
* `catalogData` - (Optional) Catalog data configuration for the repository. See [below for schema](#catalog_data).
* `tags` - (Optional) Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### catalog_data

* `aboutText` - (Optional) A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format.
* `architectures` - (Optional) The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: `ARM`, `ARM 64`, `x86`, `x86-64`
* `description` - (Optional) A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery.
* `logoImageBlob` - (Optional) The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute.
* `operatingSystems` -  (Optional) The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters: `Linux`, `Windows`
* `usageText` -  (Optional) Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Full ARN of the repository.
* `id` - The repository name.
* `registryId` - The registry ID where the repository was created.
* `repositoryUri` - The URI of the repository.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `delete` - (Default `20m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ECR Public Repositories using the `repositoryName`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcrpublicRepository } from "./.gen/providers/aws/ecrpublic-repository";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    EcrpublicRepository.generateConfigForImport(this, "example", "example");
  }
}

```

Using `terraform import`, import ECR Public Repositories using the `repositoryName`. For example:

```console
% terraform import aws_ecrpublic_repository.example example
```

<!-- cache-key: cdktf-0.20.8 input-4de95f66ada8bf61a8a550573f5ee857d57e4f09d28bb8444c629bbffb6e9856 -->