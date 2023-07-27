package basetypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// DynamicTypable extends attr.Type for dynamic types.
// Implement this interface to create a custom DynamicType type.
type DynamicTypable interface {
	attr.Type
}

var _ DynamicTypable = DynamicType{}

// DynamicType is the base framework type for a pseudo-dynamic value.
// DynamicValue is the associated value type.
type DynamicType struct {
	inner tftypes.Type
}

// ApplyTerraform5AttributePathStep applies the given AttributePathStep to the
// type.
func (t DynamicType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	// WIP
	return tftypes.DynamicPseudoType, nil
}

// Equal returns true if the given type is equivalent.
func (t DynamicType) Equal(o attr.Type) bool {
	_, ok := o.(DynamicType)

	return ok
}

// String returns a human readable string of the type name.
func (t DynamicType) String() string {
	return "basetypes.DynamicType"
}

// TerraformType returns the tftypes.Type that should be used to represent this
// framework type.
func (t DynamicType) TerraformType(_ context.Context) tftypes.Type {
	return tftypes.DynamicPseudoType
}

// ValueFromTerraform returns a Value given a tftypes.Value.  This is meant to
// convert the tftypes.Value into a more convenient Go type for the provider to
// consume the data with.
func (t DynamicType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if !in.IsKnown() {
		return NewDynamicUnknown(), nil
	}

	if in.IsNull() {
		return NewDynamicNull(), nil
	}

	return NewDynamicValue(&in), nil
}

// ValueType returns the Value type.
func (t DynamicType) ValueType(_ context.Context) attr.Value {
	// This Value does not need to be valid.
	return DynamicValue{}
}
