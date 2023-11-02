package entity

type DishOutput struct {
	ID 			 int		`json:"id"`
	Name   	 	 string 	`json:"name"`
	Description  string     `json:"description"`
	Time         int 		`json:"time"`
}

type DishInfo struct {
	ID 			 int 			`json:"id"`
	Name   	 	 string 		`json:"name"`
	Description  string     	`json:"description"`
	Recipe 		 string 	    `json:"recipe"`
	Time         int 			`json:"time"`
	Ingredients []Ingredient 	`json:"ingredients"`
}

type Ingredient struct {
	Name   	 	 string 	`json:"name"`
	MeasureUnit  string     `json:"measure_unit"`
	Quantity	 int  		`json:"quantity"`
}

type DishInput struct {
	ID       	 int			`json:"id"`
	Name   	 	 string 		`json:"name"`
	CategoryId	 int 			`json:"category_id"`
	CousineId	 int 			`json:"cousine_id"`
	Description  string     	`json:"description"`
	Recipe 		 string    	 	`json:"recipe"`
	Time         int 			`json:"time"`
	Ingredients []Ingredient 	`json:"ingredients"`
}