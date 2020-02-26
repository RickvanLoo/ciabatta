package main

//AddFlower adds flower to the recipe
func (r *Recipe) AddFlower(amount int) {
	flower := Ingredient{"Flower", amount, 1}
	r.Flower = flower
}

//AddIngredient adds an ingredient to the recipe
func (r *Recipe) AddIngredient(name string, amount int) {
	fweight := r.Flower.Amount

	ing := Ingredient{name, amount, float64(amount) / float64(fweight)}
	r.Ingredients = append(r.Ingredients, ing)
}

//AddMixM adds a Multiple Mixing Step
func (r *Recipe) AddMixM(indices []int) {
	var ings []Ingredient
	for _, i := range indices {
		ings = append(ings, r.Ingredients[i])
	}

	step := new(Step)
	step.Type = MixMT
	step.Time = 0
	step.Factor = 1
	step.Ingredients = ings
	step.Comment = ""

	r.Process = append(r.Process, *step)
}
