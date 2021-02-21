package schema

// Openshift is based on the Lagoon API type.
type Openshift struct {
	ID int `json:"id"`
	UpdateOpenshiftPatchInput
}

// UpdateOpenshiftInput is based on the Lagoon API type.
type UpdateOpenshiftInput struct {
	ID    int                       `json:"id"`
	Patch UpdateOpenshiftPatchInput `json:"patch"`
}

// UpdateOpenshiftPatchInput is based on the Lagoon API type.
type UpdateOpenshiftPatchInput struct {
	Name             string `json:"name,omitempty"`
	ConsoleURL       string `json:"consoleUrl,omitempty"`
	RouterPattern    string `json:"routerPattern,omitempty"`
	ProjectUser      string `json:"projectUser,omitempty"`
	SSHHost          string `json:"sshHost,omitempty"`
	SSHPort          string `json:"sshPort,omitempty"`
	Created          string `json:"created,omitempty"`
	Token            string `json:"token,omitempty"`
	MonitoringConfig string `json:"monitoringConfig,omitempty"`
}

// AddOpenshiftInput is based on the Lagoon API type.
type AddOpenshiftInput Openshift

// DeleteOpenshiftInput  is based on the Lagoon API type.
type DeleteOpenshiftInput struct {
	Name string `json:"name"`
}

// DeleteOpenshift is the response, it just contains "success".
type DeleteOpenshift struct {
	DeleteOpenshift string `json:"deleteOpenshift"`
}
