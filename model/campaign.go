package model

// Campaign constant values
const (
	Paused         string = "Paused"
	Implemented    string = "Implemented"
	LeadGeneration string = "LeadGeneration"
	Conversions    string = "Conversions"
	Impressions    string = "Impressions"
)

// Campaign represents the campaign object
type Campaign struct {
	ID          string   `json:"id,omitempty"`
	Template    string   `json:"template"`
	AdTitle     []string `json:"adTitle"`
	AdCopy      []string `json:"adCopy"`
	AdImage     []string `json:"image"`
	Objective   string   `json:"objective"`
	Status      string   `json:"status"`
	AdNetworkID string   `json:"adNetworkId"`
}
