package models 
type UpdateFlavorProfileSpecification struct {

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Map between global fabric flavor keys <String> and fabric flavor descriptions <FabricFlavorDescription>
	// Example: { \"small\": { \"name\": \"t2.small\" }, \"medium\": { \"name\": \"t2.medium\"}}, \"vSphere_small\": { \"cpuCount\": \"2\", \"memoryInMB\": \"2048\"}, \"vSphere_medium\": { \"cpuCount\": \"4\", \"memoryInMB\": \"4096\"}}
	// Required: true
	FlavorMapping map[string]FabricFlavorDescription `json:"flavorMapping"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`
}

