package models 
type gitWebhookSpec struct {
	branchNameField string

	delayTimeInMinsField int32

	descriptionField *string

	endpointField string

	eventCategoryField string

	exclusionsField []*GitWebhookFileFilter

	externalListenerLinkField string

	inclusionsField []*GitWebhookFileFilter

	inputField interface{}

	insecureSslField bool

	nameField *string

	pipelineField string

	prioritizeExclusionField bool

	projectField string

	refreshTokenField string

	repoNameField string

	secretTokenField string

	serverTypeField string

	serverWebhookIdField string
}

