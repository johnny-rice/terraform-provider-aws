// {{ .UpdateTagsFunc }} updates {{ .ServicePackage }} service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
{{if  .TagTypeAddBoolElem -}}
func {{ .UpdateTagsFunc }}(ctx context.Context, conn {{ .ClientType }}, identifier{{ if .TagResTypeElem }}, resourceType{{ end }} string, oldTagsSet, newTagsSet any, optFns ...func(*{{ .AWSService }}.Options)) error {
	oldTags := {{ .KeyValueTagsFunc }}(ctx, oldTagsSet, identifier{{ if .TagResTypeElem }}, resourceType{{ end }})
	newTags := {{ .KeyValueTagsFunc }}(ctx, newTagsSet, identifier{{ if .TagResTypeElem }}, resourceType{{ end }})
{{- else -}}
func {{ .UpdateTagsFunc }}(ctx context.Context, conn {{ .ClientType }}, identifier{{ if .TagResTypeElem }}, resourceType{{ end }} string, oldTagsMap, newTagsMap any, optFns ...func(*{{ .AWSService }}.Options)) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)
{{- end }}

	ctx = tflog.SetField(ctx, logging.KeyResourceId, identifier)

	{{- if eq (.TagOp) (.UntagOp) }}

	removedTags := oldTags.Removed(newTags)
	{{- if .UpdateTagsIgnoreSystem }}
	removedTags = removedTags.IgnoreSystem(names.{{ .ProviderNameUpper }})
	{{- end }}
	updatedTags := oldTags.Updated(newTags)
	{{- if .UpdateTagsIgnoreSystem }}
	updatedTags = updatedTags.IgnoreSystem(names.{{ .ProviderNameUpper }})
	{{- end }}

	// Ensure we do not send empty requests.
	if len(removedTags) == 0 && len(updatedTags) == 0 {
		return nil
	}

	input := {{ .AWSService }}.{{ .TagOp }}Input{
		{{- if not ( .TagTypeIDElem ) }}
		{{- if .TagInIDNeedValueSlice }}
		{{ .TagInIDElem }}: []string{identifier},
		{{- else }}
		{{ .TagInIDElem }}:   aws.String(identifier),
		{{- end }}
		{{- if .TagResTypeElem }}
		{{- if .TagResTypeElemType }}
		{{ .TagResTypeElem }}:      awstypes.{{ .TagResTypeElemType }}(resourceType),
		{{- else }}
		{{ .TagResTypeElem }}:      aws.String(resourceType),
		{{- end }}
		{{- end }}
		{{- end }}
	}

	if len(updatedTags) > 0 {
		input.{{ .TagInTagsElem }} = {{ .TagsFunc }}(updatedTags)
	}

	if len(removedTags) > 0 {
		{{- if .UntagInNeedTagType }}
		input.{{ .UntagInTagsElem }} = {{ .TagsFunc }}(removedTags)
		{{- else if .UntagInNeedTagKeyType }}
		input.{{ .UntagInTagsElem }} = TagKeys(removedTags)
		{{- else if .UntagInCustomVal }}
		input.{{ .UntagInTagsElem }} = {{ .UntagInCustomVal }}
		{{- else }}
		input.{{ .UntagInTagsElem }} = removedTags.Keys()
		{{- end }}
	}
	{{ if .RetryTagOps }}
	_, err := tfresource.RetryWhenIsAErrorMessageContains[any, *{{ .RetryErrorCode }}](ctx, {{ .RetryTimeout }},
		func(ctx context.Context) (any, error) {
			return conn.{{ .TagOp }}(ctx, &input, optFns...)
		},
		"{{ .RetryErrorMessage }}",
	)
	{{ else }}
	_, err := conn.{{ .TagOp }}(ctx, &input, optFns...)
	{{- end }}

	if err != nil {
		return smarterr.NewError(err)
	}

	{{- else }}

	removedTags := oldTags.Removed(newTags)
	{{- if .UpdateTagsIgnoreSystem }}
	removedTags = removedTags.IgnoreSystem(names.{{ .ProviderNameUpper }})
	{{- end }}
	if len(removedTags) > 0 {
		{{- if .TagOpBatchSize }}
		for _, removedTags := range removedTags.Chunks({{ .TagOpBatchSize }}) {
		{{- end }}
		input := {{ .AWSService }}.{{ .UntagOp }}Input{
			{{- if not ( .TagTypeIDElem ) }}
			{{- if .TagInIDNeedValueSlice }}
			{{ .TagInIDElem }}: []string{identifier},
			{{- else }}
			{{ .TagInIDElem }}:   aws.String(identifier),
			{{- end }}
			{{- if .TagResTypeElem }}
		    {{- if .TagResTypeElemType }}
			{{ .TagResTypeElem }}: awstypes.{{ .TagResTypeElemType }}(resourceType),
		    {{- else }}
			{{ .TagResTypeElem }}: aws.String(resourceType),
			{{- end }}
			{{- end }}
			{{- end }}
			{{- if .UntagInNeedTagType }}
			{{ .UntagInTagsElem }}:       {{ .TagsFunc }}(removedTags),
			{{- else if .UntagInNeedTagKeyType }}
			{{ .UntagInTagsElem }}:       TagKeys(removedTags),
			{{- else if .UntagInCustomVal }}
			{{ .UntagInTagsElem }}:       {{ .UntagInCustomVal }},
			{{- else }}
			{{ .UntagInTagsElem }}:       removedTags.Keys(),
			{{- end }}
		}
		{{ if .RetryTagOps }}
		_, err := tfresource.RetryWhenIsAErrorMessageContains[any, *{{ .RetryErrorCode }}](ctx, {{ .RetryTimeout }},
			func(ctx context.Context) (any, error) {
				return conn.{{ .UntagOp }}(ctx, &input, optFns...)
			},
			"{{ .RetryErrorMessage }}",
		)
		{{ else }}
		_, err := conn.{{ .UntagOp }}(ctx, &input, optFns...)
		{{- end }}

		if err != nil {
			return smarterr.NewError(err)
		}
		{{- if .TagOpBatchSize }}
		}
		{{- end }}
	}

	updatedTags := oldTags.Updated(newTags)
	{{- if .UpdateTagsIgnoreSystem }}
	updatedTags = updatedTags.IgnoreSystem(names.{{ .ProviderNameUpper }})
	{{- end }}
	if len(updatedTags) > 0 {
		{{- if .TagOpBatchSize }}
		for _, updatedTags := range updatedTags.Chunks({{ .TagOpBatchSize }}) {
		{{- end }}
		input := {{ .AWSService }}.{{ .TagOp }}Input{
			{{- if not ( .TagTypeIDElem ) }}
			{{- if .TagInIDNeedValueSlice }}
			{{ .TagInIDElem }}: []string{identifier},
			{{- else }}
			{{ .TagInIDElem }}: aws.String(identifier),
			{{- end }}
			{{- if .TagResTypeElem }}
		    {{- if .TagResTypeElemType }}
			{{ .TagResTypeElem }}:    awstypes.{{ .TagResTypeElemType }}(resourceType),
		    {{- else }}
			{{ .TagResTypeElem }}:    aws.String(resourceType),
			{{- end }}
			{{- end }}
			{{- end }}
			{{- if .TagInCustomVal }}
			{{ .TagInTagsElem }}:       {{ .TagInCustomVal }},
			{{- else }}
			{{ .TagInTagsElem }}:       {{ .TagsFunc }}(updatedTags),
			{{- end }}
		}

		{{ if .RetryTagOps }}
		_, err := tfresource.RetryWhenIsAErrorMessageContains[any, *{{ .RetryErrorCode }}](ctx, {{ .RetryTimeout }},
			func(ctx context.Context) (any, error) {
				return conn.{{ .TagOp }}(ctx, &input, optFns...)
			},
			"{{ .RetryErrorMessage }}",
		)
		{{ else }}
		_, err := conn.{{ .TagOp }}(ctx, &input, optFns...)
		{{- end }}

		if err != nil {
			return smarterr.NewError(err)
		}
		{{- if .TagOpBatchSize }}
		}
		{{- end }}
	}

	{{- end }}

	{{ if .WaitForPropagation }}
	if len(removedTags) > 0 || len(updatedTags) > 0 {
		if err := {{ .WaitTagsPropagatedFunc }}(ctx, conn, identifier, newTags, optFns...); err != nil {
			return smarterr.NewError(err)
		}
	}
	{{- end }}

	return nil
}

{{- if .IsDefaultUpdateTags }}
// {{ .UpdateTagsFunc | Title }} updates {{ .ServicePackage }} service tags.
// It is called from outside this package.
func (p *servicePackage) {{ .UpdateTagsFunc | Title }}(ctx context.Context, meta any, identifier{{ if .TagResTypeElem }}, resourceType{{ end }} string, oldTags, newTags any) error {
	return  {{ .UpdateTagsFunc }}(ctx, meta.(*conns.AWSClient).{{ .ProviderNameUpper }}Client(ctx), identifier{{ if .TagResTypeElem }}, resourceType{{ end }}, oldTags, newTags)
}
{{- end }}
