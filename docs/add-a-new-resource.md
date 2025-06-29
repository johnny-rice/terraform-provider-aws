<!-- markdownlint-configure-file { "code-block-style": false } -->
# Adding a New Resource

New resources are required when AWS adds a new service, or adds new features within an existing service which would require a new resource to manage in Terraform. Typically anything with a new set of CRUD API endpoints is a great candidate for a new resource.

Each resource should be submitted for review in isolation. Pull requests containing multiple resources are harder to review and the maintainers will normally ask for them to be broken apart.

## Prerequisites

If this is the first resource for a new service, please ensure the Service Client for the new service has been added and merged. See [Adding a new Service](add-a-new-service.md) for details.

## Steps to Add a Resource

### Fork the provider and create a feature branch

For new resources use a branch named `f-{resource name}` for example: `f-ec2-vpc`. See [Raising a Pull Request](raising-a-pull-request.md) for more details.

### Create and Name the Resource

See the [Naming Guide](naming.md#resources-and-data-sources) for details on how to name the new resource and the resource file. Not following the naming standards will cause extra delay as maintainers request that you make changes.

Use the [skaff](skaff.md) provider scaffolding tool to generate new resource and test templates using your chosen name. Doing so will ensure that any boilerplate code, structural best practices and repetitive naming are done for you and always represent our most current standards.

### Fill out the Resource Schema

In the `internal/service/<service>/<service>.go` file you will see a `Schema` property which exists as a map of `Schema` objects. This relates the AWS API data model with the Terraform resource itself. For each property you want to make available in Terraform, you will need to add it as an attribute, choose the correct data type and supply the correct [Schema Behaviors](https://www.terraform.io/plugin/sdkv2/schemas/schema-behaviors) to ensure Terraform knows how to correctly handle the value.

Typically you will add arguments to represent the values that are under control by Terraform, and attributes to supply read-only values as references for Terraform. These are distinguished by Schema Behavior.

Attribute names are to be specified in `snake_case` as opposed to the AWS API which is `CamelCase`.

### Implement CRUD handlers

These will map the planned Terraform state to the AWS API call, or an AWS API response to an applied Terraform state. You will also need to handle different response types (including errors correctly). For complex attributes, you will need to implement Flattener or Expander functions. The [Data Handling and Conversion Guide](data-handling-and-conversion.md) covers everything you need to know for mapping AWS API responses to Terraform State and vice-versa. The [Error Handling Guide](error-handling.md) covers everything you need to know about handling AWS API responses consistently.

### Register Resource to the provider

Resources use a self-registration process that adds them to the provider using the `@FrameworkResource()` or `@SDKResource()` annotation in the resource's comments. Run `make gen` to register the resource. This will add an entry to the `service_package_gen.go` file located in the service package folder.

=== "Terraform Plugin Framework (Preferred)"

    ```go
    package something

    import (
        "github.com/hashicorp/terraform-plugin-framework/resource"
        "github.com/hashicorp/terraform-provider-aws/internal/framework"
    )

    // @FrameworkResource("aws_something_example", name="Example")
    func newExampleResource(_ context.Context) (resource.ResourceWithConfigure, error) {
    	return &resourceExample{}, nil
    }

    type exampleResource struct {
    	framework.ResourceWithModel[exampleResourceModel]
    }

    type exampleResourceModel struct {
    	// Fields corresponding to attributes in the Schema.
    }
    ```

=== "Terraform Plugin SDK V2"

    ```go
    package something

    import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

    // @SDKResource("aws_something_example", name="Example)
    func ResourceExample() *schema.Resource {
    	return &schema.Resource{
    	    // some configuration
    	}
    }
    ```

### Write passing Acceptance Tests

To adequately test the resource we will need to write a complete set of Acceptance Tests. You will need an AWS account for this which allows the creation of that resource. See [Writing Acceptance Tests](running-and-writing-acceptance-tests.md) for a detailed guide on how to approach these.

You will need at a minimum:

- Basic Test - Tests full lifecycle (CRUD + Import) of a minimal configuration (all required fields, no optional).
- Disappears Test - Tests what Terraform does if a resource it is tracking can no longer be found.
- Argument Tests - All arguments should be tested in a pragmatic way. Ensure that each argument can be initially set, updated, and cleared, as applicable. Depending on the logic and interaction of arguments, this may take one to several separate tests.

### Create documentation for the resource

Add a file covering the use of the new resource in `website/docs/r/<service>_<name>.md`. Add more examples if it is complex or relies on resources in another service. This documentation will appear on the [Terraform Registry](https://registry.terraform.io/providers/hashicorp/aws/latest) when the resource is made available in a provider release. Link to AWS Documentation where appropriate, particularly for values which are likely to change.

### Ensure format and lint checks are passing locally

Format your code and check linters to detect various issues.

```sh
make fmt
make tools     # install linters and dependencies
make lint      # run provider linters
make docs-lint # run documentation linters
```

### Raise a Pull Request

See [Raising a Pull Request](raising-a-pull-request.md).

### Wait for Prioritization

In general, pull requests are triaged within a few days of creation and are prioritized based on community reactions. Please view our [Prioritization Guide](prioritization.md) for full details of the process.
