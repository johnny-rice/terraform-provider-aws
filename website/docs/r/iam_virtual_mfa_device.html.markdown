---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "AWS: aws_iam_virtual_mfa_device"
description: |-
  Provides an IAM Virtual MFA Device
---

# Resource: aws_iam_virtual_mfa_device

Provides an IAM Virtual MFA Device.

~> **Note:** All attributes will be stored in the raw state as plain-text.
[Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

~> **Note:** A virtual MFA device cannot be directly associated with an IAM User from Terraform.
  To associate the virtual MFA device with a user and enable it, use the code returned in either `base_32_string_seed` or `qr_code_png` to generate TOTP authentication codes.
  The authentication codes can then be used with the AWS CLI command [`aws iam enable-mfa-device`](https://docs.aws.amazon.com/cli/latest/reference/iam/enable-mfa-device.html) or the AWS API call [`EnableMFADevice`](https://docs.aws.amazon.com/IAM/latest/APIReference/API_EnableMFADevice.html).

## Example Usage

**Using certs on file:**

```terraform
resource "aws_iam_virtual_mfa_device" "example" {
  virtual_mfa_device_name = "example"
}
```

## Argument Reference

This resource supports the following arguments:

* `virtual_mfa_device_name` - (Required) The name of the virtual MFA device. Use with path to uniquely identify a virtual MFA device.
* `path` - (Optional) The path for the virtual MFA device.
* `tags` - (Optional) Map of resource tags for the virtual mfa device. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) specifying the virtual mfa device.
* `base_32_string_seed` - The base32 seed defined as specified in [RFC3548](https://tools.ietf.org/html/rfc3548.txt). The `base_32_string_seed` is base64-encoded.
* `enable_date` - The date and time when the virtual MFA device was enabled.
* `qr_code_png` -  A QR code PNG image that encodes `otpauth://totp/$virtualMFADeviceName@$AccountName?secret=$Base32String` where `$virtualMFADeviceName` is one of the create call arguments. AccountName is the user name if set (otherwise, the account ID), and Base32String is the seed in base32 format.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).
* `user_name` - The associated IAM User name if the virtual MFA device is enabled.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import IAM Virtual MFA Devices using the `arn`. For example:

```terraform
import {
  to = aws_iam_virtual_mfa_device.example
  id = "arn:aws:iam::123456789012:mfa/example"
}
```

Using `terraform import`, import IAM Virtual MFA Devices using the `arn`. For example:

```console
% terraform import aws_iam_virtual_mfa_device.example arn:aws:iam::123456789012:mfa/example
```
