// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package glue

import (
	"context"
	"log"
	"time"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	awstypes "github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_glue_schema", name="Schema")
// @Tags(identifierAttribute="arn")
// @ArnIdentity
// @Testing(preIdentityVersion="v6.3.0")
// @Testing(existsType="github.com/aws/aws-sdk-go-v2/service/glue;glue.GetSchemaOutput")
func resourceSchema() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceSchemaCreate,
		ReadWithoutTimeout:   resourceSchemaRead,
		UpdateWithoutTimeout: resourceSchemaUpdate,
		DeleteWithoutTimeout: resourceSchemaDelete,

		Schema: map[string]*schema.Schema{
			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrDescription: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 2048),
			},
			"registry_arn": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: verify.ValidARN,
			},
			"registry_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"latest_schema_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"next_schema_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"schema_checkpoint": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"compatibility": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: enum.Validate[awstypes.Compatibility](),
			},
			"data_format": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: enum.Validate[awstypes.DataFormat](),
			},
			"schema_definition": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(1, 170000),
					validation.StringMatch(regexache.MustCompile(`.*\S.*`), ""),
				),
			},
			"schema_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(1, 255),
					validation.StringMatch(regexache.MustCompile(`[0-9A-Za-z_$#-]+$`), ""),
				),
			},
			names.AttrTags:    tftags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
		},
	}
}

func resourceSchemaCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).GlueClient(ctx)

	input := &glue.CreateSchemaInput{
		SchemaName:       aws.String(d.Get("schema_name").(string)),
		SchemaDefinition: aws.String(d.Get("schema_definition").(string)),
		DataFormat:       awstypes.DataFormat(d.Get("data_format").(string)),
		Tags:             getTagsIn(ctx),
	}

	if v, ok := d.GetOk("registry_arn"); ok {
		input.RegistryId = createRegistryID(v.(string))
	}

	if v, ok := d.GetOk(names.AttrDescription); ok {
		input.Description = aws.String(v.(string))
	}

	if v, ok := d.GetOk("compatibility"); ok {
		input.Compatibility = awstypes.Compatibility(v.(string))
	}

	log.Printf("[DEBUG] Creating Glue Schema: %+v", input)
	output, err := conn.CreateSchema(ctx, input)
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating Glue Schema: %s", err)
	}
	d.SetId(aws.ToString(output.SchemaArn))

	_, err = waitSchemaAvailable(ctx, conn, d.Id())
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for Glue Schema (%s) to be Available: %s", d.Id(), err)
	}

	return append(diags, resourceSchemaRead(ctx, d, meta)...)
}

func resourceSchemaRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).GlueClient(ctx)

	output, err := findSchemaByID(ctx, conn, d.Id())
	if err != nil {
		if errs.IsA[*awstypes.EntityNotFoundException](err) {
			log.Printf("[WARN] Glue Schema (%s) not found, removing from state", d.Id())
			d.SetId("")
			return diags
		}
		return sdkdiag.AppendErrorf(diags, "reading Glue Schema (%s): %s", d.Id(), err)
	}

	if output == nil {
		log.Printf("[WARN] Glue Schema (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	arn := aws.ToString(output.SchemaArn)
	d.Set(names.AttrARN, arn)
	d.Set(names.AttrDescription, output.Description)
	d.Set("schema_name", output.SchemaName)
	d.Set("compatibility", output.Compatibility)
	d.Set("data_format", output.DataFormat)
	d.Set("latest_schema_version", output.LatestSchemaVersion)
	d.Set("next_schema_version", output.NextSchemaVersion)
	d.Set("registry_arn", output.RegistryArn)
	d.Set("registry_name", output.RegistryName)
	d.Set("schema_checkpoint", output.SchemaCheckpoint)

	schemeDefOutput, err := findSchemaVersionByID(ctx, conn, d.Id())
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading Glue Schema Definition (%s): %s", d.Id(), err)
	}

	d.Set("schema_definition", schemeDefOutput.SchemaDefinition)

	return diags
}

func resourceSchemaUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).GlueClient(ctx)

	input := &glue.UpdateSchemaInput{
		SchemaId: createSchemaID(d.Id()),
		SchemaVersionNumber: &awstypes.SchemaVersionNumber{
			LatestVersion: true,
		},
	}
	update := false

	if d.HasChange(names.AttrDescription) {
		input.Description = aws.String(d.Get(names.AttrDescription).(string))
		update = true
	}

	if d.HasChange("compatibility") {
		input.Compatibility = awstypes.Compatibility(d.Get("compatibility").(string))
		update = true
	}

	if update {
		log.Printf("[DEBUG] Updating Glue Schema: %#v", input)
		_, err := conn.UpdateSchema(ctx, input)
		if err != nil {
			return sdkdiag.AppendErrorf(diags, "updating Glue Schema (%s): %s", d.Id(), err)
		}

		_, err = waitSchemaAvailable(ctx, conn, d.Id())
		if err != nil {
			return sdkdiag.AppendErrorf(diags, "waiting for Glue Schema (%s) to be Available: %s", d.Id(), err)
		}
	}

	if d.HasChange("schema_definition") {
		defInput := &glue.RegisterSchemaVersionInput{
			SchemaId:         createSchemaID(d.Id()),
			SchemaDefinition: aws.String(d.Get("schema_definition").(string)),
		}

		_, err := conn.RegisterSchemaVersion(ctx, defInput)
		if err != nil {
			return sdkdiag.AppendErrorf(diags, "updating Glue Schema Definition (%s): %s", d.Id(), err)
		}

		_, err = waitSchemaVersionAvailable(ctx, conn, d.Id())
		if err != nil {
			return sdkdiag.AppendErrorf(diags, "waiting for Glue Schema Version (%s) to be Available: %s", d.Id(), err)
		}
	}

	return append(diags, resourceSchemaRead(ctx, d, meta)...)
}

func resourceSchemaDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).GlueClient(ctx)

	log.Printf("[DEBUG] Deleting Glue Schema: %s", d.Id())
	input := &glue.DeleteSchemaInput{
		SchemaId: createSchemaID(d.Id()),
	}

	_, err := conn.DeleteSchema(ctx, input)
	if err != nil {
		if errs.IsA[*awstypes.EntityNotFoundException](err) {
			return diags
		}
		return sdkdiag.AppendErrorf(diags, "deleting Glue Schema (%s): %s", d.Id(), err)
	}

	_, err = waitSchemaDeleted(ctx, conn, d.Id())
	if err != nil {
		if errs.IsA[*awstypes.EntityNotFoundException](err) {
			return diags
		}
		return sdkdiag.AppendErrorf(diags, "waiting for Glue Schema (%s) to be deleted: %s", d.Id(), err)
	}

	return diags
}

func findSchemaByID(ctx context.Context, conn *glue.Client, id string) (*glue.GetSchemaOutput, error) {
	input := &glue.GetSchemaInput{
		SchemaId: createSchemaID(id),
	}

	output, err := conn.GetSchema(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// statusSchema fetches the Schema and its Status
func statusSchema(ctx context.Context, conn *glue.Client, id string) retry.StateRefreshFunc {
	const (
		schemaStatusUnknown = "Unknown"
	)
	return func() (any, string, error) {
		output, err := findSchemaByID(ctx, conn, id)

		if err != nil {
			return nil, schemaStatusUnknown, err
		}

		if output == nil {
			return output, schemaStatusUnknown, nil
		}

		return output, string(output.SchemaStatus), nil
	}
}

func waitSchemaAvailable(ctx context.Context, conn *glue.Client, registryID string) (*glue.GetSchemaOutput, error) { //nolint:unparam
	const (
		timeout = 2 * time.Minute
	)
	stateConf := &retry.StateChangeConf{
		Pending: enum.Slice(awstypes.SchemaStatusPending),
		Target:  enum.Slice(awstypes.SchemaStatusAvailable),
		Refresh: statusSchema(ctx, conn, registryID),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*glue.GetSchemaOutput); ok {
		return output, err
	}

	return nil, err
}

func waitSchemaDeleted(ctx context.Context, conn *glue.Client, registryID string) (*glue.GetSchemaOutput, error) {
	const (
		timeout = 2 * time.Minute
	)
	stateConf := &retry.StateChangeConf{
		Pending: enum.Slice(awstypes.SchemaStatusDeleting),
		Target:  []string{},
		Refresh: statusSchema(ctx, conn, registryID),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*glue.GetSchemaOutput); ok {
		return output, err
	}

	return nil, err
}

func findSchemaVersionByID(ctx context.Context, conn *glue.Client, id string) (*glue.GetSchemaVersionOutput, error) {
	input := &glue.GetSchemaVersionInput{
		SchemaId: createSchemaID(id),
		SchemaVersionNumber: &awstypes.SchemaVersionNumber{
			LatestVersion: true,
		},
	}

	output, err := conn.GetSchemaVersion(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// statusSchemaVersion fetches the Schema Version and its Status
func statusSchemaVersion(ctx context.Context, conn *glue.Client, id string) retry.StateRefreshFunc {
	const (
		schemaVersionStatusUnknown = "Unknown"
	)
	return func() (any, string, error) {
		output, err := findSchemaVersionByID(ctx, conn, id)

		if err != nil {
			return nil, schemaVersionStatusUnknown, err
		}

		if output == nil {
			return output, schemaVersionStatusUnknown, nil
		}

		return output, string(output.Status), nil
	}
}

func waitSchemaVersionAvailable(ctx context.Context, conn *glue.Client, registryID string) (*glue.GetSchemaVersionOutput, error) {
	const (
		timeout = 2 * time.Minute
	)
	stateConf := &retry.StateChangeConf{
		Pending: enum.Slice(awstypes.SchemaVersionStatusPending),
		Target:  enum.Slice(awstypes.SchemaVersionStatusAvailable),
		Refresh: statusSchemaVersion(ctx, conn, registryID),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*glue.GetSchemaVersionOutput); ok {
		return output, err
	}

	return nil, err
}
