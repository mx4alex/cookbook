package entity

type DishInStorage struct {
	Id       	 int		`json:"id" bson:"_id"`
	Name   	 	 string 	`json:"name" bson:"name"`
	CategoryId	 int 		`json:"category_id" bson:"category_id"`
	CousineId	 int 		`json:"cousine_id" bson:"cousine_id"`
	Description  string     `json:"description" bson:"description"`
	Recipe 		 string     `json:"recipe" bson:"recipe"`
	Time         int 		`json:"time" bson:"time"`
}

type DishOutput struct {
	Name   	 	 string 	`json:"name" bson:"name"`
	Description  string     `json:"description" bson:"description"`
	Time         int 		`json:"time" bson:"time"`
}

type DishInfo struct {
	Name   	 	 string 		`json:"name" bson:"name"`
	Description  string     	`json:"description" bson:"description"`
	Recipe 		 string 	    `json:"recipe" bson:"recipe"`
	Time         int 			`json:"time" bson:"time"`
	Ingredients []Ingredient 	`json:"ingredients" bson:"ingredients"`
}

type Ingredient struct {
	Name   	 	 string 	`json:"name" bson:"name"`
	MeasureUnit  string     `json:"measure_unit" bson:"measure_unit"`
	Quantity	 int  		`json:"quantity" bson:"quantity"`
}

func NewIngredient(name string, measureUnit string, quantity int) Ingredient {
	return Ingredient {
		Name:   	 name,
		MeasureUnit: measureUnit,
		Quantity: 	 quantity,
	}
}

func NewDishOutput (name string, measureUnit string, quantity int) Ingredient {
	return Ingredient {
		Name:   	 name,
		MeasureUnit: measureUnit,
		Quantity: 	 quantity,
	}
}