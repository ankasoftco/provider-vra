package models 
type batchUserOperationResponse struct {
	errorsField []BatchUserOperationStatus

	failureField int32

	successField int32
}

