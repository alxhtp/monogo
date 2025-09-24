package message

import "fmt"

type ResponseMessage string

const (
	SuccessCreated ResponseMessage = "Successfully created a"
	SuccessUpdated ResponseMessage = "Successfully updated a"
	SuccessDeleted ResponseMessage = "Successfully deleted a"
	SuccessList    ResponseMessage = "Successfully got a list of"
	SuccessGetByID ResponseMessage = "Successfully got a"
	FailedCreated  ResponseMessage = "Failed to create a"
	FailedUpdated  ResponseMessage = "Failed to update a"
	FailedDeleted  ResponseMessage = "Failed to delete a"
	FailedList     ResponseMessage = "Failed to get a list of"
	FailedUpdate   ResponseMessage = "Failed to update a"
	FailedDelete   ResponseMessage = "Failed to delete a"
	FailedGetByID  ResponseMessage = "Failed to get a"
)

func GetResponseMessage(message ResponseMessage, entity string) string {
	return fmt.Sprintf("%s %s", string(message), entity)
}
