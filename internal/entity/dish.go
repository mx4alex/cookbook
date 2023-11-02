package entity

type Ingredient struct {
	Name   	 	 string 	`json:"name"`
	MeasureUnit  string     `json:"measure_unit"`
	Quantity	 int  		`json:"quantity"`
}

type Dish struct {
	ID       	 int			`json:"id"`
	Name   	 	 string 		`json:"name"`
	CategoryID	 int 			`json:"category_id"`
	CousineID	 int 			`json:"cousine_id"`
	Description  string     	`json:"description"`
	Recipe 		 string    	 	`json:"recipe"`
	Time         int 			`json:"time"`
	Ingredients []Ingredient 	`json:"ingredients"`
}