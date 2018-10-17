package model

// Ad represents the ad object
type Ad struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
	Copy  string `json:"copy"`
	Image string `json:"image"`
}
