package models 
type updateProjectSpecification struct {
	constraintsField map[string]ProjectConstraint

	descriptionField string

	nameField string

	operationTimeoutField int64

	propertiesField map[string]string

	sharedResourcesField bool
}

