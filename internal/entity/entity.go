package entity

type Ingredient struct {
	Name   	 	 	string 		`json:"name"`
	MeasureUnit 	string  	`json:"measure_unit"`
	Quantity		int  		`json:"quantity"`
	Protein 	 	int 		`json:"ingredient_protein"`
	Fats 		 	int 		`json:"ingredient_fats"`
	Carbohydrates  int	 		`json:"ingredient_carbohydrates"`
}

type Dish struct {
	ID       	 	int				`json:"id"`
	Name   	 	 	string 			`json:"name"`
	CategoryID	 	int 			`json:"category_id"`
	CousineID	 	int 			`json:"cousine_id"`
	Description  	string     		`json:"description"`
	Recipe 		 	string    	 	`json:"recipe"`
	Time        	int 			`json:"time"`
	Protein 	 	int 			`json:"dish_protein"`
	Fats 		 	int 			`json:"dish_fats"`
	Carbohydrates  int	 			`json:"dish_carbohydrates"`
	Kilocalories	int				`json:"kilocalories"`
	Ingredients  	[]Ingredient 	`json:"ingredients"`
}

type Cousine struct {
	ID 			int		`json:"id"`
	Name 		string 	`json:"name"`
	Description string 	`json:"description"`
}

type Category struct {
	ID 			int		`json:"id"`
	Name 		string 	`json:"name"`
	Description string 	`json:"description"`
}