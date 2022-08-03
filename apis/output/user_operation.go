package models 
type userOperation struct {
	createTimeInMicrosField int64

	expiresOnInMicrosField int64

	linkField string

	projectIdField string

	requestedOnInMicrosField int64

	updateTimeInMicrosField int64

	approversField []string

	cancelPreviousPendingUserOpField bool

	createdAtField string

	createdByField string

	descriptionField *string

	endpointField string

	executionIdField string

	executionIndexField int64

	executionLinkField string

	expirationField int32

	expirationInDaysField int32

	expirationUnitField string

	idField string

	indexField string

	nameField *string

	pipelineIdField string

	pipelineNameField string

	projectField string

	requestedByField string

	respondedByField string

	responseMessageField string

	sendemailField bool

	statusField string

	summaryField string

	updatedAtField string

	updatedByField string

	versionField string
}

