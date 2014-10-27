package generic

// Type is the placeholder type that indicates a generic value.
// When genny is executed, variables of this type will be replaced with
// references to the specific types.
//      var GenericType generic.Type
type Type interface{}
