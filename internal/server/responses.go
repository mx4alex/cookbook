package server

type statusResponse struct {
	Status string `json:"status"`
}
type errorResponse struct {
	Message string `json:"message"`
}

func StatusResponse(message string) statusResponse {
	return statusResponse{
        Status: message,
    }
}