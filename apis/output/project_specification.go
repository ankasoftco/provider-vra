package models 
type projectSpecification struct {
	administratorsField []Principal

	constraintsField map[string]ProjectConstraint

	costField ProjectCost

	descriptionField string

	membersField []Principal

	nameField *string

	operationTimeoutField *int64

	propertiesField map[string]string

	sharedResourcesField *bool

	viewersField []Principal
}

