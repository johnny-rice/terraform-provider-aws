// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_security_group", name="Security Group")
// @Tags
// @Testing(tagsIdentifierAttribute="id")
func dataSourceSecurityGroup() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceSecurityGroupRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrDescription: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrFilter: customFiltersSchema(),
			names.AttrID: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			names.AttrName: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			names.AttrTags: tftags.TagsSchemaComputed(),
			names.AttrVPCID: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceSecurityGroupRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := meta.(*conns.AWSClient)
	conn := c.EC2Client(ctx)

	input := &ec2.DescribeSecurityGroupsInput{
		Filters: newAttributeFilterList(
			map[string]string{
				"group-name": d.Get(names.AttrName).(string),
				"vpc-id":     d.Get(names.AttrVPCID).(string),
			},
		),
	}

	if v, ok := d.GetOk(names.AttrID); ok {
		input.GroupIds = []string{v.(string)}
	}

	input.Filters = append(input.Filters, newTagFilterList(
		svcTags(tftags.New(ctx, d.Get(names.AttrTags).(map[string]any))),
	)...)

	input.Filters = append(input.Filters, newCustomFilterList(
		d.Get(names.AttrFilter).(*schema.Set),
	)...)

	if len(input.Filters) == 0 {
		// Don't send an empty filters list; the EC2 API won't accept it.
		input.Filters = nil
	}

	sg, err := findSecurityGroup(ctx, conn, input)

	if err != nil {
		return sdkdiag.AppendFromErr(diags, tfresource.SingularDataSourceFindError("EC2 Security Group", err))
	}

	d.SetId(aws.ToString(sg.GroupId))
	ownerID := aws.ToString(sg.OwnerId)
	d.Set(names.AttrARN, securityGroupARN(ctx, c, ownerID, d.Id()))
	d.Set(names.AttrDescription, sg.Description)
	d.Set(names.AttrName, sg.GroupName)
	d.Set(names.AttrVPCID, sg.VpcId)

	setTagsOut(ctx, sg.Tags)

	return diags
}
