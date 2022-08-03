package models 
type NotificationScenarioConfig struct {

	// Notification scenario enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// Notification scenario category
	ScenarioCategory string `json:"scenarioCategory,omitempty"`

	// Notification scenario description
	ScenarioDescription string `json:"scenarioDescription,omitempty"`

	// Notification scenario id
	// Required: true
	ScenarioID *string `json:"scenarioId"`

	// Notification scenario name
	ScenarioName string `json:"scenarioName,omitempty"`
}

