package server

type statusResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Message string `json:"message"`
}

type statusID struct {
	ID int `json:"id"`
}

func StatusResponse(message string) statusResponse {
	return statusResponse{
        Status: message,
    }
}

type DishOutput struct {
	ID 			 int		`json:"id"`
	Name   	 	 string 	`json:"name"`
	Description  string     `json:"description"`
	Time         int 		`json:"time"`
}
