// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamicdefault

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// StaticDynamic returns a static string value default handler.
//
// Use StaticDynamic if a static default value for a string should be set.
func StaticDynamic(defaultVal tftypes.Value) defaults.Dynamic {
	return staticDynamicDefault{
		defaultVal: defaultVal,
	}
}

// staticDynamicDefault is static value default handler that
// sets a value on a string attribute.
type staticDynamicDefault struct {
	defaultVal tftypes.Value
}

// Description returns a human-readable description of the default value handler.
func (d staticDynamicDefault) Description(_ context.Context) string {
	return fmt.Sprintf("value defaults to %s", d.defaultVal.String())
}

// MarkdownDescription returns a markdown description of the default value handler.
func (d staticDynamicDefault) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("value defaults to `%s`", d.defaultVal.String())
}

// DefaultDynamic implements the static default value logic.
func (d staticDynamicDefault) DefaultDynamic(_ context.Context, req defaults.DynamicRequest, resp *defaults.DynamicResponse) {
	resp.PlanValue = types.DynamicValue(&d.defaultVal)
}
