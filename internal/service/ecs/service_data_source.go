// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_ecs_service", name="Service")
// @Tags
func dataSourceService() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceServiceRead,

		Schema: map[string]*schema.Schema{
			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_zone_rebalancing": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_arn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"desired_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"launch_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"load_balancer": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"container_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"container_port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"elb_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_group_arn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"advanced_configuration": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alternate_target_group_arn": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"production_listener_rule": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"test_listener_rule": {
										Type:     schema.TypeString,
										Computed: true,
									},
									names.AttrRoleARN: {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"scheduling_strategy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrServiceName: {
				Type:     schema.TypeString,
				Required: true,
			},
			"task_definition": {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrTags: tftags.TagsSchemaComputed(),
		},
	}
}

func dataSourceServiceRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).ECSClient(ctx)

	service, err := findServiceByTwoPartKey(ctx, conn, d.Get(names.AttrServiceName).(string), d.Get("cluster_arn").(string))

	if err != nil {
		return sdkdiag.AppendFromErr(diags, tfresource.SingularDataSourceFindError("ECS Service", err))
	}

	arn := aws.ToString(service.ServiceArn)
	d.SetId(arn)
	d.Set(names.AttrARN, arn)
	d.Set("availability_zone_rebalancing", service.AvailabilityZoneRebalancing)
	d.Set("cluster_arn", service.ClusterArn)
	d.Set("desired_count", service.DesiredCount)
	d.Set("launch_type", service.LaunchType)
	if service.LoadBalancers != nil {
		if err := d.Set("load_balancer", flattenServiceLoadBalancers(service.LoadBalancers)); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting load_balancer: %s", err)
		}
	}
	d.Set("scheduling_strategy", service.SchedulingStrategy)
	d.Set(names.AttrServiceName, service.ServiceName)
	d.Set("task_definition", service.TaskDefinition)

	setTagsOut(ctx, service.Tags)

	return diags
}
