package defaults

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Dynamic is a schema default value for types.Dynamic attributes.
type Dynamic interface {
	Describer

	// DefaultDynamic should Dynamic the default value.
	DefaultDynamic(context.Context, DynamicRequest, *DynamicResponse)
}

type DynamicRequest struct {
	// Path contains the path of the attribute for Dynamicting the
	// default value. Use this path for any response diagnostics.
	Path path.Path
}

type DynamicResponse struct {
	// Diagnostics report errors or warnings related to Dynamicting the
	// default value resource configuration. An empty slice
	// indicates success, with no warnings or errors generated.
	Diagnostics diag.Diagnostics

	// PlanValue is the planned new state for the attribute.
	PlanValue types.Dynamic
}
