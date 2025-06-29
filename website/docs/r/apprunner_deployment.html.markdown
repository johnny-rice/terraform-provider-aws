---
subcategory: "App Runner"
layout: "aws"
page_title: "AWS: aws_apprunner_deployment"
description: |-
  Manages an App Runner Deployment Operation.
---

# Resource: aws_apprunner_deployment

Manages an App Runner Deployment Operation.

## Example Usage

```terraform
resource "aws_apprunner_deployment" "example" {
  service_arn = aws_apprunner_service.example.arn
}
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `service_arn` - (Required) The Amazon Resource Name (ARN) of the App Runner service to start the deployment for.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - A unique identifier for the deployment.
* `operation_id` - The unique ID of the operation associated with deployment.
* `status` - The current status of the App Runner service deployment.
