// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/hashicorp/aws-sdk-go-base/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	sdkid "github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_verifiedaccess_instance", name="Verified Access Instance")
// @Tags(identifierAttribute="id")
// @Testing(tagsTest=false)
func resourceVerifiedAccessInstance() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceVerifiedAccessInstanceCreate,
		ReadWithoutTimeout:   resourceVerifiedAccessInstanceRead,
		UpdateWithoutTimeout: resourceVerifiedAccessInstanceUpdate,
		DeleteWithoutTimeout: resourceVerifiedAccessInstanceDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			names.AttrCreationTime: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cidr_endpoints_custom_subdomain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			names.AttrDescription: {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fips_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			names.AttrLastUpdatedTime: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"verified_access_trust_providers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						names.AttrDescription: {
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_trust_provider_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trust_provider_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_trust_provider_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"verified_access_trust_provider_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"name_servers": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			names.AttrTags:    tftags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
		},
	}
}

func resourceVerifiedAccessInstanceCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	input := ec2.CreateVerifiedAccessInstanceInput{
		ClientToken:       aws.String(sdkid.UniqueId()),
		TagSpecifications: getTagSpecificationsIn(ctx, awstypes.ResourceTypeVerifiedAccessInstance),
	}

	if v, ok := d.GetOk("cidr_endpoints_custom_subdomain"); ok {
		input.CidrEndpointsCustomSubDomain = aws.String(v.(string))
	}

	if v, ok := d.GetOk(names.AttrDescription); ok {
		input.Description = aws.String(v.(string))
	}

	if v, ok := d.GetOk("fips_enabled"); ok {
		input.FIPSEnabled = aws.Bool(v.(bool))
	}

	output, err := conn.CreateVerifiedAccessInstance(ctx, &input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating Verified Access Instance: %s", err)
	}

	d.SetId(aws.ToString(output.VerifiedAccessInstance.VerifiedAccessInstanceId))

	return append(diags, resourceVerifiedAccessInstanceRead(ctx, d, meta)...)
}

func resourceVerifiedAccessInstanceRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	output, err := findVerifiedAccessInstanceByID(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] EC2 Verified Access Instance (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading Verified Access Instance (%s): %s", d.Id(), err)
	}

	d.Set(names.AttrCreationTime, output.CreationTime)
	if output.CidrEndpointsCustomSubDomain != nil {
		d.Set("cidr_endpoints_custom_subdomain", output.CidrEndpointsCustomSubDomain.SubDomain)
		d.Set("name_servers", output.CidrEndpointsCustomSubDomain.Nameservers)
	}
	d.Set(names.AttrDescription, output.Description)
	d.Set("fips_enabled", output.FipsEnabled)
	d.Set(names.AttrLastUpdatedTime, output.LastUpdatedTime)
	if v := output.VerifiedAccessTrustProviders; v != nil {
		if err := d.Set("verified_access_trust_providers", flattenVerifiedAccessTrustProviders(v)); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting verified_access_trust_providers: %s", err)
		}
	} else {
		d.Set("verified_access_trust_providers", nil)
	}

	setTagsOut(ctx, output.Tags)

	return diags
}

func resourceVerifiedAccessInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	if d.HasChangesExcept(names.AttrTags, names.AttrTagsAll) {
		input := ec2.ModifyVerifiedAccessInstanceInput{
			ClientToken:              aws.String(sdkid.UniqueId()),
			VerifiedAccessInstanceId: aws.String(d.Id()),
		}

		if d.HasChange("cidr_endpoints_custom_subdomain") {
			input.CidrEndpointsCustomSubDomain = aws.String(d.Get("cidr_endpoints_custom_subdomain").(string))
		}

		if d.HasChange(names.AttrDescription) {
			input.Description = aws.String(d.Get(names.AttrDescription).(string))
		}

		_, err := conn.ModifyVerifiedAccessInstance(ctx, &input)

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "updating Verified Access Instance (%s): %s", d.Id(), err)
		}
	}

	return append(diags, resourceVerifiedAccessInstanceRead(ctx, d, meta)...)
}

func resourceVerifiedAccessInstanceDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	log.Printf("[INFO] Deleting Verified Access Instance: %s", d.Id())
	input := ec2.DeleteVerifiedAccessInstanceInput{
		ClientToken:              aws.String(sdkid.UniqueId()),
		VerifiedAccessInstanceId: aws.String(d.Id()),
	}
	_, err := conn.DeleteVerifiedAccessInstance(ctx, &input)

	if tfawserr.ErrCodeEquals(err, errCodeInvalidVerifiedAccessInstanceIdNotFound) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting Verified Access Instance (%s): %s", d.Id(), err)
	}

	return diags
}

func flattenVerifiedAccessTrustProviders(apiObjects []awstypes.VerifiedAccessTrustProviderCondensed) []any {
	if len(apiObjects) == 0 {
		return nil
	}

	var tfList []any

	for _, apiObject := range apiObjects {
		v := flattenVerifiedAccessTrustProvider(apiObject)

		if len(v) > 0 {
			tfList = append(tfList, v)
		}
	}

	return tfList
}

func flattenVerifiedAccessTrustProvider(apiObject awstypes.VerifiedAccessTrustProviderCondensed) map[string]any {
	tfMap := map[string]any{
		"device_trust_provider_type": apiObject.DeviceTrustProviderType,
		"trust_provider_type":        apiObject.TrustProviderType,
		"user_trust_provider_type":   apiObject.UserTrustProviderType,
	}

	if v := apiObject.Description; v != nil {
		tfMap[names.AttrDescription] = aws.ToString(v)
	}

	if v := apiObject.VerifiedAccessTrustProviderId; v != nil {
		tfMap["verified_access_trust_provider_id"] = aws.ToString(v)
	}

	return tfMap
}
