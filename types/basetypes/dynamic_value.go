package basetypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ DynamicValuable = DynamicValue{}
)

// DynamicValuable extends attr.Value for dynamic value types.
// Implement this interface to create a custom dynamic value type.
type DynamicValuable interface {
	attr.Value
}

// DynamicValuableWithSemanticEquals extends DynamicValuable with semantic
// equality logic.
type DynamicValuableWithSemanticEquals interface {
	DynamicValuable

	// DynamicValueSemanticEquals should return true if the given value is
	// semantically equal to the current value. This logic is used to prevent
	// Terraform data consistency errors and resource drift where a value change
	// may have inconsequential differences.
	//
	// Only known values are compared with this method as changing a value's
	// state implicitly represents a different value.
	DynamicValueSemanticEquals(context.Context, DynamicValuable) (bool, diag.Diagnostics)
}

// NewDynamicNull creates a DynamicValue with a null value. Determine whether the value is
// null via the Dynamic type IsNull method.
func NewDynamicNull() DynamicValue {
	return DynamicValue{
		state: attr.ValueStateNull,
	}
}

// NewDynamicUnknown creates a Dynamic with an unknown value. Determine whether the
// value is unknown via the Dynamic type IsUnknown method.
func NewDynamicUnknown() DynamicValue {
	return DynamicValue{
		state: attr.ValueStateUnknown,
	}
}

// NewDynamicValue creates a Dynamic with a known value. If the given value is
// nil, a null Dynamic is created.
func NewDynamicValue(value *tftypes.Value) DynamicValue {
	if value == nil {
		return NewDynamicNull()
	}

	return DynamicValue{
		state: attr.ValueStateKnown,
		value: value,
	}
}

// DynamicValue represents a pseudo-dynamic value.
type DynamicValue struct {
	// state represents whether the value is null, unknown, or known. The
	// zero-value is null.
	state attr.ValueState

	// value contains the known value, if not null or unknown.
	value *tftypes.Value
}

// Type returns a DynamicType.
func (d DynamicValue) Type(_ context.Context) attr.Type {
	return DynamicType{
		inner: d.value.Type(),
	}
}

// ToTerraformValue returns the data contained in the DynamicValue as a tftypes.Value.
func (d DynamicValue) ToTerraformValue(_ context.Context) (tftypes.Value, error) {
	switch d.state {
	case attr.ValueStateKnown:
		if d.value == nil {
			return tftypes.NewValue(tftypes.DynamicPseudoType, nil), nil
		}

		return *d.value, nil
	case attr.ValueStateNull:
		return tftypes.NewValue(tftypes.DynamicPseudoType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled DynamicValue state in ToTerraformValue: %s", d.state))
	}
}

// Equal returns true if `other` is a DynamicValue and has the same value as `d`.
func (d DynamicValue) Equal(other attr.Value) bool {
	o, ok := other.(DynamicValue)

	if !ok {
		return false
	}

	if d.state != o.state {
		return false
	}

	if d.state != attr.ValueStateKnown {
		return true
	}

	if d.value == nil || o.value == nil {
		return d.value == o.value
	}

	return d.value.Equal(*o.value)
}

// IsNull returns true if the DynamicValue represents a null value.
func (d DynamicValue) IsNull() bool {
	return d.state == attr.ValueStateNull
}

// IsUnknown returns true if the DynamicValue represents a currently unknown value.
func (d DynamicValue) IsUnknown() bool {
	return d.state == attr.ValueStateUnknown
}

// String returns a human-readable representation of the DynamicValue.
// The string returned here is not protected by any compatibility guarantees,
// and is intended for logging and error reporting.
func (d DynamicValue) String() string {
	if d.IsUnknown() {
		return attr.UnknownValueString
	}

	if d.IsNull() {
		return attr.NullValueString
	}

	return d.value.String()
}
