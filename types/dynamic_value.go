package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type Dynamic = basetypes.DynamicValue

// DynamicNull creates a DynamicValue with a null value. Determine whether the
// value is null via the Dynamic type IsNull method.
func DynamicNull() basetypes.DynamicValue {
	return basetypes.NewDynamicNull()
}

// DynamicUnknown creates a DynamicValue with an unknown value. Determine whether
// the value is unknown via the Dynamic type IsUnknown method.
func DynamicUnknown() basetypes.DynamicValue {
	return basetypes.NewDynamicUnknown()
}

// DynamicValue creates a DynamicValue with a known value. If the given value
// is nil, a null DynamicValue is created.
func DynamicValue(value *tftypes.Value) basetypes.DynamicValue {
	return basetypes.NewDynamicValue(value)
}
