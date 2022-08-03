package models 
type gitWebhook struct {
	createTimeInMicrosField int64

	linkField string

	projectIdField string

	updateTimeInMicrosField int64

	branchNameField string

	createdAtField string

	createdByField string

	delayTimeInMinsField int32

	descriptionField *string

	endpointField string

	eventCategoryField string

	exclusionsField []*GitWebhookFileFilter

	externalListenerLinkField string

	idField string

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

	updatedAtField string

	updatedByField string

	versionField string
}

