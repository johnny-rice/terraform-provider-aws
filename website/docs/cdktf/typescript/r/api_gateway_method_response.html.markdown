---
subcategory: "API Gateway"
layout: "aws"
page_title: "AWS: aws_api_gateway_method_response"
description: |-
  Provides an HTTP Method Response for an API Gateway Resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_api_gateway_method_response

Provides an HTTP Method Response for an API Gateway Resource. More information about API Gateway method responses can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-settings-method-response.html).

## Example Usage

### Basic Response

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApiGatewayIntegration } from "./.gen/providers/aws/api-gateway-integration";
import { ApiGatewayMethod } from "./.gen/providers/aws/api-gateway-method";
import { ApiGatewayMethodResponse } from "./.gen/providers/aws/api-gateway-method-response";
import { ApiGatewayResource } from "./.gen/providers/aws/api-gateway-resource";
import { ApiGatewayRestApi } from "./.gen/providers/aws/api-gateway-rest-api";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const myDemoApi = new ApiGatewayRestApi(this, "MyDemoAPI", {
      description: "This is my API for demonstration purposes",
      name: "MyDemoAPI",
    });
    const myDemoResource = new ApiGatewayResource(this, "MyDemoResource", {
      parentId: myDemoApi.rootResourceId,
      pathPart: "mydemoresource",
      restApiId: myDemoApi.id,
    });
    const myDemoMethod = new ApiGatewayMethod(this, "MyDemoMethod", {
      authorization: "NONE",
      httpMethod: "GET",
      resourceId: myDemoResource.id,
      restApiId: myDemoApi.id,
    });
    new ApiGatewayMethodResponse(this, "response_200", {
      httpMethod: myDemoMethod.httpMethod,
      resourceId: myDemoResource.id,
      restApiId: myDemoApi.id,
      statusCode: "200",
    });
    new ApiGatewayIntegration(this, "MyDemoIntegration", {
      httpMethod: myDemoMethod.httpMethod,
      resourceId: myDemoResource.id,
      restApiId: myDemoApi.id,
      type: "MOCK",
    });
  }
}

```

### Response with Custom Header and Model

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApiGatewayIntegration } from "./.gen/providers/aws/api-gateway-integration";
import { ApiGatewayMethod } from "./.gen/providers/aws/api-gateway-method";
import { ApiGatewayMethodResponse } from "./.gen/providers/aws/api-gateway-method-response";
import { ApiGatewayModel } from "./.gen/providers/aws/api-gateway-model";
import { ApiGatewayResource } from "./.gen/providers/aws/api-gateway-resource";
import { ApiGatewayRestApi } from "./.gen/providers/aws/api-gateway-rest-api";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const myDemoApi = new ApiGatewayRestApi(this, "MyDemoAPI", {
      description: "This is my API for demonstration purposes",
      name: "MyDemoAPI",
    });
    new ApiGatewayModel(this, "MyDemoResponseModel", {
      contentType: "application/json",
      description: "API response for MyDemoMethod",
      name: "MyDemoResponseModel",
      restApiId: myDemoApi.id,
      schema: Token.asString(
        Fn.jsonencode({
          $schema: "http://json-schema.org/draft-04/schema#",
          properties: {
            Message: {
              type: "string",
            },
          },
          title: "MyDemoResponse",
          type: "object",
        })
      ),
    });
    const myDemoResource = new ApiGatewayResource(this, "MyDemoResource", {
      parentId: myDemoApi.rootResourceId,
      pathPart: "mydemoresource",
      restApiId: myDemoApi.id,
    });
    const myDemoMethod = new ApiGatewayMethod(this, "MyDemoMethod", {
      authorization: "NONE",
      httpMethod: "GET",
      resourceId: myDemoResource.id,
      restApiId: myDemoApi.id,
    });
    new ApiGatewayMethodResponse(this, "response_200", {
      httpMethod: myDemoMethod.httpMethod,
      resourceId: myDemoResource.id,
      responseModels: {
        "application/json": "MyDemoResponseModel",
      },
      responseParameters: {
        "method-response-header.X-My-Demo-Header": false,
        "method.response.header.Content-Type": false,
      },
      restApiId: myDemoApi.id,
      statusCode: "200",
    });
    new ApiGatewayIntegration(this, "MyDemoIntegration", {
      httpMethod: myDemoMethod.httpMethod,
      resourceId: myDemoResource.id,
      restApiId: myDemoApi.id,
      type: "MOCK",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `restApiId` - (Required) The string identifier of the associated REST API.
* `resourceId` - (Required) The Resource identifier for the method resource.
* `httpMethod` - (Required) The HTTP verb of the method resource (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`).
* `statusCode` - (Required) The method response's status code.
* `responseModels` - (Optional) A map specifying the model resources used for the response's content type. Response models are represented as a key/value map, with a content type as the key and a Model name as the value.
* `responseParameters` - (Optional) A map specifying required or optional response parameters that API Gateway can send back to the caller. A key defines a method response header name and the associated value is a boolean flag indicating whether the method response parameter is required. The method response header names must match the pattern of `method.response.header.{name}`, where `name` is a valid and unique header name.

The response parameter names defined here are available in the integration response to be mapped from an integration response header expressed in `integration.response.header.{name}`, a static value enclosed within a pair of single quotes (e.g., '`application/json'`), or a JSON expression from the back-end response payload in the form of `integration.response.body.{JSON-expression}`, where `JSON-expression` is a valid JSON expression without the `$` prefix.)

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_api_gateway_method_response` using `REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApiGatewayMethodResponse } from "./.gen/providers/aws/api-gateway-method-response";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ApiGatewayMethodResponse.generateConfigForImport(
      this,
      "example",
      "12345abcde/67890fghij/GET/200"
    );
  }
}

```

Using `terraform import`, import `aws_api_gateway_method_response` using `REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE`. For example:

```console
% terraform import aws_api_gateway_method_response.example 12345abcde/67890fghij/GET/200
```

<!-- cache-key: cdktf-0.20.8 input-99186c6c0499f61bb223e736bf1cfab300f9c32ba5c22b8fb08eac3ed80762d2 -->