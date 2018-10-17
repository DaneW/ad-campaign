package model

// Template contains title, ad copy, and compain objective.
type Template struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Ads              []Ad   `json:"ads"`
	CompainObjective string `json:"compainObjective"`
}
