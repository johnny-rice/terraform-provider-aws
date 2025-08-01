// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	fwflex "github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @FrameworkDataSource("aws_cloudfront_origin_access_control", name="Origin Access Control")
func newOriginAccessControlDataSource(_ context.Context) (datasource.DataSourceWithConfigure, error) {
	d := &originAccessControlDataSource{}

	return d, nil
}

type originAccessControlDataSource struct {
	framework.DataSourceWithModel[originAccessControlDataSourceModel]
}

const (
	DSNameOriginAccessControl = "Origin Access Control Data Source"
)

func (d *originAccessControlDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			names.AttrARN: schema.StringAttribute{
				Computed: true,
			},
			names.AttrDescription: schema.StringAttribute{
				Computed: true,
			},
			"etag": schema.StringAttribute{
				Computed: true,
			},
			names.AttrID: schema.StringAttribute{
				Required: true,
			},
			names.AttrName: schema.StringAttribute{
				Computed: true,
			},
			"origin_access_control_origin_type": schema.StringAttribute{
				Computed: true,
			},
			"signing_behavior": schema.StringAttribute{
				Computed: true,
			},
			"signing_protocol": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *originAccessControlDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	conn := d.Meta().CloudFrontClient(ctx)
	var data originAccessControlDataSourceModel

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	output, err := findOriginAccessControlByID(ctx, conn, data.ID.ValueString())

	if err != nil {
		response.Diagnostics.AddError(
			create.ProblemStandardMessage(names.CloudFront, create.ErrActionReading, DSNameOriginAccessControl, data.ID.String(), err),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(fwflex.Flatten(ctx, output.OriginAccessControl.OriginAccessControlConfig, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	data.ARN = fwflex.StringValueToFramework(ctx, originAccessControlARN(ctx, d.Meta(), data.ID.ValueString()))
	data.Etag = fwflex.StringToFramework(ctx, output.ETag)

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

type originAccessControlDataSourceModel struct {
	ARN                           types.String `tfsdk:"arn"`
	Description                   types.String `tfsdk:"description"`
	Etag                          types.String `tfsdk:"etag"`
	ID                            types.String `tfsdk:"id"`
	Name                          types.String `tfsdk:"name"`
	OriginAccessControlOriginType types.String `tfsdk:"origin_access_control_origin_type"`
	SigningBehavior               types.String `tfsdk:"signing_behavior"`
	SigningProtocol               types.String `tfsdk:"signing_protocol"`
}
