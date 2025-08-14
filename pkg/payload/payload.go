package payload

type FailurePayload struct {
	Status string `json:"status"`
	Message string `json:"message"`
}