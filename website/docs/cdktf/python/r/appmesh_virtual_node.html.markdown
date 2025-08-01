---
subcategory: "App Mesh"
layout: "aws"
page_title: "AWS: aws_appmesh_virtual_node"
description: |-
  Provides an AWS App Mesh virtual node resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_appmesh_virtual_node

Provides an AWS App Mesh virtual node resource.

## Breaking Changes

Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92)), `aws_appmesh_virtual_node` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:

* Rename the `service_name` attribute of the `dns` object to `hostname`.

* Replace the `backends` attribute of the `spec` object with one or more `backend` configuration blocks,
setting `virtual_service_name` to the name of the service.

The Terraform state associated with existing resources will automatically be migrated.

## Example Usage

### Basic

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_virtual_node import AppmeshVirtualNode
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshVirtualNode(self, "serviceb1",
            mesh_name=simple.id,
            name="serviceBv1",
            spec=AppmeshVirtualNodeSpec(
                backend=[AppmeshVirtualNodeSpecBackend(
                    virtual_service=AppmeshVirtualNodeSpecBackendVirtualService(
                        virtual_service_name="servicea.simpleapp.local"
                    )
                )
                ],
                listener=[AppmeshVirtualNodeSpecListener(
                    port_mapping=AppmeshVirtualNodeSpecListenerPortMapping(
                        port=8080,
                        protocol="http"
                    )
                )
                ],
                service_discovery=AppmeshVirtualNodeSpecServiceDiscovery(
                    dns=AppmeshVirtualNodeSpecServiceDiscoveryDns(
                        hostname="serviceb.simpleapp.local"
                    )
                )
            )
        )
```

### AWS Cloud Map Service Discovery

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_virtual_node import AppmeshVirtualNode
from imports.aws.service_discovery_http_namespace import ServiceDiscoveryHttpNamespace
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = ServiceDiscoveryHttpNamespace(self, "example",
            name="example-ns"
        )
        AppmeshVirtualNode(self, "serviceb1",
            mesh_name=simple.id,
            name="serviceBv1",
            spec=AppmeshVirtualNodeSpec(
                backend=[AppmeshVirtualNodeSpecBackend(
                    virtual_service=AppmeshVirtualNodeSpecBackendVirtualService(
                        virtual_service_name="servicea.simpleapp.local"
                    )
                )
                ],
                listener=[AppmeshVirtualNodeSpecListener(
                    port_mapping=AppmeshVirtualNodeSpecListenerPortMapping(
                        port=8080,
                        protocol="http"
                    )
                )
                ],
                service_discovery=AppmeshVirtualNodeSpecServiceDiscovery(
                    aws_cloud_map=AppmeshVirtualNodeSpecServiceDiscoveryAwsCloudMap(
                        attributes={
                            "stack": "blue"
                        },
                        namespace_name=example.name,
                        service_name="serviceb1"
                    )
                )
            )
        )
```

### Listener Health Check

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_virtual_node import AppmeshVirtualNode
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshVirtualNode(self, "serviceb1",
            mesh_name=simple.id,
            name="serviceBv1",
            spec=AppmeshVirtualNodeSpec(
                backend=[AppmeshVirtualNodeSpecBackend(
                    virtual_service=AppmeshVirtualNodeSpecBackendVirtualService(
                        virtual_service_name="servicea.simpleapp.local"
                    )
                )
                ],
                listener=[AppmeshVirtualNodeSpecListener(
                    health_check=AppmeshVirtualNodeSpecListenerHealthCheck(
                        healthy_threshold=2,
                        interval_millis=5000,
                        path="/ping",
                        protocol="http",
                        timeout_millis=2000,
                        unhealthy_threshold=2
                    ),
                    port_mapping=AppmeshVirtualNodeSpecListenerPortMapping(
                        port=8080,
                        protocol="http"
                    )
                )
                ],
                service_discovery=AppmeshVirtualNodeSpecServiceDiscovery(
                    dns=AppmeshVirtualNodeSpecServiceDiscoveryDns(
                        hostname="serviceb.simpleapp.local"
                    )
                )
            )
        )
```

### Logging

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_virtual_node import AppmeshVirtualNode
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshVirtualNode(self, "serviceb1",
            mesh_name=simple.id,
            name="serviceBv1",
            spec=AppmeshVirtualNodeSpec(
                backend=[AppmeshVirtualNodeSpecBackend(
                    virtual_service=AppmeshVirtualNodeSpecBackendVirtualService(
                        virtual_service_name="servicea.simpleapp.local"
                    )
                )
                ],
                listener=[AppmeshVirtualNodeSpecListener(
                    port_mapping=AppmeshVirtualNodeSpecListenerPortMapping(
                        port=8080,
                        protocol="http"
                    )
                )
                ],
                logging=AppmeshVirtualNodeSpecLogging(
                    access_log=AppmeshVirtualNodeSpecLoggingAccessLog(
                        file=AppmeshVirtualNodeSpecLoggingAccessLogFile(
                            path="/dev/stdout"
                        )
                    )
                ),
                service_discovery=AppmeshVirtualNodeSpecServiceDiscovery(
                    dns=AppmeshVirtualNodeSpecServiceDiscoveryDns(
                        hostname="serviceb.simpleapp.local"
                    )
                )
            )
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name to use for the virtual node. Must be between 1 and 255 characters in length.
* `mesh_name` - (Required) Name of the service mesh in which to create the virtual node. Must be between 1 and 255 characters in length.
* `mesh_owner` - (Optional) AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider][1] is currently connected to.
* `spec` - (Required) Virtual node specification to apply.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

The `spec` object supports the following:

* `backend` - (Optional) Backends to which the virtual node is expected to send outbound traffic.
* `backend_defaults` - (Optional) Defaults for backends.
* `listener` - (Optional) Listeners from which the virtual node is expected to receive inbound traffic.
* `logging` - (Optional) Inbound and outbound access logging information for the virtual node.
* `service_discovery` - (Optional) Service discovery information for the virtual node.

The `backend` object supports the following:

* `virtual_service` - (Required) Virtual service to use as a backend for a virtual node.

The `virtual_service` object supports the following:

* `client_policy` - (Optional) Client policy for the backend.
* `virtual_service_name` - (Required) Name of the virtual service that is acting as a virtual node backend. Must be between 1 and 255 characters in length.

The `client_policy` object supports the following:

* `tls` - (Optional) Transport Layer Security (TLS) client policy.

The `tls` object supports the following:

* `certificate` (Optional) Virtual node's client's Transport Layer Security (TLS) certificate.
* `enforce` - (Optional) Whether the policy is enforced. Default is `true`.
* `ports` - (Optional) One or more ports that the policy is enforced for.
* `validation` - (Required) TLS validation context.

The `certificate` object supports the following:

* `file` - (Optional) Local file certificate.
* `sds` - (Optional) A [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.

The `file` object supports the following:

* `certificate_chain` - (Required) Certificate chain for the certificate.
* `private_key` - (Required) Private key for a certificate stored on the file system of the mesh endpoint that the proxy is running on.

The `sds` object supports the following:

* `secret_name` - (Required) Name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.

The `validation` object supports the following:

* `subject_alternative_names` - (Optional) SANs for a TLS validation context.
* `trust` - (Required) TLS validation context trust.

The `subject_alternative_names` object supports the following:

* `match` - (Required) Criteria for determining a SAN's match.

The `match` object supports the following:

* `exact` - (Required) Values sent must match the specified values exactly.

The `trust` object supports the following:

* `acm` - (Optional) TLS validation context trust for an AWS Certificate Manager (ACM) certificate.
* `file` - (Optional) TLS validation context trust for a local file certificate.
* `sds` - (Optional) TLS validation context trust for a [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.

The `acm` object supports the following:

* `certificate_authority_arns` - (Required) One or more ACM ARNs.

The `file` object supports the following:

* `certificate_chain` - (Required) Certificate trust chain for a certificate stored on the file system of the virtual node that the proxy is running on. Must be between 1 and 255 characters in length.

The `sds` object supports the following:

* `secret_name` - (Required) Name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.

The `backend_defaults` object supports the following:

* `client_policy` - (Optional) Default client policy for virtual service backends. See above for details.

The `listener` object supports the following:

* `port_mapping` - (Required) Port mapping information for the listener.
* `connection_pool` - (Optional) Connection pool information for the listener.
* `health_check` - (Optional) Health check information for the listener.
* `outlier_detection` - (Optional) Outlier detection information for the listener.
* `timeout` - (Optional) Timeouts for different protocols.
* `tls` - (Optional) Transport Layer Security (TLS) properties for the listener

The `logging` object supports the following:

* `access_log` - (Optional) Access log configuration for a virtual node.

The `access_log` object supports the following:

* `file` - (Optional) File object to send virtual node access logs to.

The `file` object supports the following:

* `format` - (Optional) The specified format for the logs.
* `path` - (Required) File path to write access logs to. You can use `/dev/stdout` to send access logs to standard out. Must be between 1 and 255 characters in length.

The `format` object supports the following:

* `json` - (Optional) The logging format for JSON.
* `text` - (Optional) The logging format for text. Must be between 1 and 1000 characters in length.

The `json` object supports the following:

* `key` - (Required) The specified key for the JSON. Must be between 1 and 100 characters in length.
* `value` - (Required) The specified value for the JSON. Must be between 1 and 100 characters in length.

The `service_discovery` object supports the following:

* `aws_cloud_map` - (Optional) Any AWS Cloud Map information for the virtual node.
* `dns` - (Optional) DNS service name for the virtual node.

The `aws_cloud_map` object supports the following:

* `attributes` - (Optional) String map that contains attributes with values that you can use to filter instances by any custom attribute that you specified when you registered the instance. Only instances that match all of the specified key/value pairs will be returned.
* `namespace_name` - (Required) Name of the AWS Cloud Map namespace to use.
Use the [`aws_service_discovery_http_namespace`](/docs/providers/aws/r/service_discovery_http_namespace.html) resource to configure a Cloud Map namespace. Must be between 1 and 1024 characters in length.
* `service_name` - (Required) Name of the AWS Cloud Map service to use. Use the [`aws_service_discovery_service`](/docs/providers/aws/r/service_discovery_service.html) resource to configure a Cloud Map service. Must be between 1 and 1024 characters in length.

The `dns` object supports the following:

* `hostname` - (Required) DNS host name for your virtual node.
* `ip_preference` - (Optional) The preferred IP version that this virtual node uses. Valid values: `IPv6_PREFERRED`, `IPv4_PREFERRED`, `IPv4_ONLY`, `IPv6_ONLY`.
* `response_type` - (Optional) The DNS response type for the virtual node. Valid values: `LOADBALANCER`, `ENDPOINTS`.

The `port_mapping` object supports the following:

* `port` - (Required) Port used for the port mapping.
* `protocol` - (Required) Protocol used for the port mapping. Valid values are `http`, `http2`, `tcp` and `grpc`.

The `connection_pool` object supports the following:

* `grpc` - (Optional) Connection pool information for gRPC listeners.
* `http` - (Optional) Connection pool information for HTTP listeners.
* `http2` - (Optional) Connection pool information for HTTP2 listeners.
* `tcp` - (Optional) Connection pool information for TCP listeners.

The `grpc` connection pool object supports the following:

* `max_requests` - (Required) Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster. Minimum value of `1`.

The `http` connection pool object supports the following:

* `max_connections` - (Required) Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster. Minimum value of `1`.
* `max_pending_requests` - (Optional) Number of overflowing requests after `max_connections` Envoy will queue to upstream cluster. Minimum value of `1`.

The `http2` connection pool object supports the following:

* `max_requests` - (Required) Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster. Minimum value of `1`.

The `tcp` connection pool object supports the following:

* `max_connections` - (Required) Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster. Minimum value of `1`.

The `health_check` object supports the following:

* `healthy_threshold` - (Required) Number of consecutive successful health checks that must occur before declaring listener healthy.
* `interval_millis`- (Required) Time period in milliseconds between each health check execution.
* `protocol` - (Required) Protocol for the health check request. Valid values are `http`, `http2`, `tcp` and `grpc`.
* `timeout_millis` - (Required) Amount of time to wait when receiving a response from the health check, in milliseconds.
* `unhealthy_threshold` - (Required) Number of consecutive failed health checks that must occur before declaring a virtual node unhealthy.
* `path` - (Optional) Destination path for the health check request. This is only required if the specified protocol is `http` or `http2`.
* `port` - (Optional) Destination port for the health check request. This port must match the port defined in the `port_mapping` for the listener.

The `outlier_detection` object supports the following:

* `base_ejection_duration` - (Required) Base amount of time for which a host is ejected.
* `interval` - (Required) Time interval between ejection sweep analysis.
* `max_ejection_percent` - (Required) Maximum percentage of hosts in load balancing pool for upstream service that can be ejected. Will eject at least one host regardless of the value.
Minimum value of `0`. Maximum value of `100`.
* `max_server_errors` - (Required) Number of consecutive `5xx` errors required for ejection. Minimum value of `1`.

The `base_ejection_duration` and `interval` objects support the following:

* `unit` - (Required) Unit of time. Valid values: `ms`, `s`.
* `value` - (Required) Number of time units. Minimum value of `0`.

The `timeout` object supports the following:

* `grpc` - (Optional) Timeouts for gRPC listeners.
* `http` - (Optional) Timeouts for HTTP listeners.
* `http2` - (Optional) Timeouts for HTTP2 listeners.
* `tcp` - (Optional) Timeouts for TCP listeners.

The `grpc` timeout object supports the following:

* `idle` - (Optional) Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
* `per_request` - (Optional) Per request timeout.

The `idle` and `per_request` objects support the following:

* `unit` - (Required) Unit of time. Valid values: `ms`, `s`.
* `value` - (Required) Number of time units. Minimum value of `0`.

The `http` and `http2` timeout objects support the following:

* `idle` - (Optional) Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
* `per_request` - (Optional) Per request timeout.

The `idle` and `per_request` objects support the following:

* `unit` - (Required) Unit of time. Valid values: `ms`, `s`.
* `value` - (Required) Number of time units. Minimum value of `0`.

The `tcp` timeout object supports the following:

* `idle` - (Optional) Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.

The `idle` object supports the following:

* `unit` - (Required) Unit of time. Valid values: `ms`, `s`.
* `value` - (Required) Number of time units. Minimum value of `0`.

The `tls` object supports the following:

* `certificate` - (Required) Listener's TLS certificate.
* `mode`- (Required) Listener's TLS mode. Valid values: `DISABLED`, `PERMISSIVE`, `STRICT`.
* `validation`- (Optional) Listener's Transport Layer Security (TLS) validation context.

The `certificate` object supports the following:

* `acm` - (Optional) An AWS Certificate Manager (ACM) certificate.
* `file` - (Optional) Local file certificate.
* `sds` - (Optional) A [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.

The `acm` object supports the following:

* `certificate_arn` - (Required) ARN for the certificate.

The `file` object supports the following:

* `certificate_chain` - (Required) Certificate chain for the certificate. Must be between 1 and 255 characters in length.
* `private_key` - (Required) Private key for a certificate stored on the file system of the virtual node that the proxy is running on. Must be between 1 and 255 characters in length.

The `sds` object supports the following:

* `secret_name` - (Required) Name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.

The `validation` object supports the following:

* `subject_alternative_names` - (Optional) SANs for a TLS validation context.
* `trust` - (Required) TLS validation context trust.

The `subject_alternative_names` object supports the following:

* `match` - (Required) Criteria for determining a SAN's match.

The `match` object supports the following:

* `exact` - (Required) Values sent must match the specified values exactly.

The `trust` object supports the following:

* `file` - (Optional) TLS validation context trust for a local file certificate.
* `sds` - (Optional) TLS validation context trust for a [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.

The `file` object supports the following:

* `certificate_chain` - (Required) Certificate trust chain for a certificate stored on the file system of the mesh endpoint that the proxy is running on. Must be between 1 and 255 characters in length.

The `sds` object supports the following:

* `secret_name` - (Required) Name of the secret for a virtual node's Transport Layer Security (TLS) Secret Discovery Service validation context trust.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ID of the virtual node.
* `arn` - ARN of the virtual node.
* `created_date` - Creation date of the virtual node.
* `last_updated_date` - Last update date of the virtual node.
* `resource_owner` - Resource owner's AWS account ID.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import App Mesh virtual nodes using `mesh_name` together with the virtual node's `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.appmesh_virtual_node import AppmeshVirtualNode
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AppmeshVirtualNode.generate_config_for_import(self, "serviceb1", "simpleapp/serviceBv1")
```

Using `terraform import`, import App Mesh virtual nodes using `mesh_name` together with the virtual node's `name`. For example:

```console
% terraform import aws_appmesh_virtual_node.serviceb1 simpleapp/serviceBv1
```

[1]: /docs/providers/aws/index.html

<!-- cache-key: cdktf-0.20.8 input-897fc148eebb7acd39cbc2b383d7226226e543105c8facd9fcf57226ff439e7f -->