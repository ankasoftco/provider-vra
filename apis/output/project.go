package models 
type project struct {
	administratorsField []Principal

	constraintsField map[string]ProjectConstraint

	costField ProjectCost

	descriptionField string

	idField string

	membersField []Principal

	nameField string

	operationTimeoutField int64

	orgIdField string

	propertiesField map[string]string

	sharedResourcesField bool

	supervisorsField []Principal

	viewersField []Principal
}

