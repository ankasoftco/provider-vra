package models 
type gerritTriggerSpec struct {
	branchField *string

	configurationsField []*GerritTriggerSpecGerritEventConfiguration

	descriptionField *string

	enabledField bool

	exclusionsField []*GerritTriggerFileFilter

	gerritProjectField *string

	inclusionsField []*GerritTriggerFileFilter

	listenerField *string

	nameField *string

	prioritizeExclusionField bool

	projectField string
}

