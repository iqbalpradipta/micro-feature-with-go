package helpers

type SuccessResponse struct {
	Messages 	string
	Data		interface{}
}

type FailedResponse struct {
	Messages 	string
	Code		int
}