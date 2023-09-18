package models

type Recipe struct {
	ID                    int
	Title                 string
	Image                 string
	UsedIngredientCount   int
	MissedIngredientCount int
	MissedIngredients     []MissedIngredient
}

type MissedIngredient struct {
	ID    int
	Name  string
	Image string
}

type SavedRecipe struct {
	ID           int
	Title        string
	Instructions string
	Image        string
}
