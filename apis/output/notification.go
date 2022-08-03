package models 
type notification struct {
	eventGroupField string

	jiraResponseField map[string]string

	providerTypeField string

	responseCodeField string

	responseMessageField string

	stageField string

	successField bool

	taskField string
}

