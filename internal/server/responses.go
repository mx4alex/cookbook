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

func newStatusResponse(message string) statusResponse {
	return statusResponse{
        Status: message,
    }
}

type dishOutput struct {
	ID 			 int		`json:"id"`
	Name   	 	 string 	`json:"name"`
	Description  string     `json:"description"`
	Time         int 		`json:"time"`
}

type categoryInfo struct {
	Name 		string 	`json:"name"`
	Description string 	`json:"description"`
}

type cousineInfo struct {
	Name 		string 	`json:"name"`
	Description string 	`json:"description"`
}