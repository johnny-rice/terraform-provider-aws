---
subcategory: "IoT Core"
layout: "aws"
page_title: "AWS: aws_iot_certificate"
description: |-
    Creates and manages an AWS IoT certificate.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_iot_certificate

Creates and manages an AWS IoT certificate.

## Example Usage

### With CSR

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.iot_certificate import IotCertificate
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        IotCertificate(self, "cert",
            active=True,
            csr=Token.as_string(Fn.file("/my/csr.pem"))
        )
```

### Without CSR

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.iot_certificate import IotCertificate
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        IotCertificate(self, "cert",
            active=True
        )
```

### From existing certificate without a CA

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.iot_certificate import IotCertificate
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        IotCertificate(self, "cert",
            active=True,
            certificate_pem=Token.as_string(Fn.file("/my/cert.pem"))
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `active` - (Required)  Boolean flag to indicate if the certificate should be active
* `csr` - (Optional) The certificate signing request. Review
  [CreateCertificateFromCsr](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
  for more information on generating a certificate from a certificate signing request (CSR).
  If none is specified both the certificate and keys will be generated, review [CreateKeysAndCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateKeysAndCertificate.html)
  for more information on generating keys and a certificate.
* `certificate_pem` - (Optional) The certificate to be registered. If `ca_pem` is unspecified, review
  [RegisterCertificateWithoutCA](https://docs.aws.amazon.com/iot/latest/apireference/API_RegisterCertificateWithoutCA.html).
  If `ca_pem` is specified, review
  [RegisterCertificate](https://docs.aws.amazon.com/iot/latest/apireference/API_RegisterCertificate.html)
  for more information on registering a certificate.
* `ca_pem` - (Optional) The CA certificate for the certificate to be registered. If this is set, the CA needs to be registered with AWS IoT beforehand.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The internal ID assigned to this certificate.
* `arn` - The ARN of the created certificate.
* `ca_certificate_id` - The certificate ID of the CA certificate used to sign the certificate.
* `certificate_pem` - The certificate data, in PEM format.
* `public_key` - When neither CSR nor certificate is provided, the public key.
* `private_key` - When neither CSR nor certificate is provided, the private key.

<!-- cache-key: cdktf-0.20.8 input-4dc5a24579e02caa93934cb266036c6a14f9df13bb9b4b183037cdb4ca450b80 -->