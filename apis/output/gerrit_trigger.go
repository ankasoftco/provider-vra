package models 
type gerritTrigger struct {
	createTimeInMicrosField int64

	linkField string

	projectIdField string

	updateTimeInMicrosField int64

	branchField string

	configurationsField []*GerritTriggerGerritEventConfiguration

	createdAtField string

	createdByField string

	descriptionField *string

	enabledField bool

	exclusionsField []*GerritTriggerFileFilter

	gerritProjectField string

	idField string

	inclusionsField []*GerritTriggerFileFilter

	listenerField string

	nameField *string

	prioritizeExclusionField bool

	projectField string

	updatedAtField string

	updatedByField string

	versionField string
}

