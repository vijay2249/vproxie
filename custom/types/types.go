package types

type HeadersConfig struct {
	ModifyRequest ModificationConfig `yaml:"modifyRequest"`
	ModifyResponse ModificationConfig `yaml:"modifyResponse"`
}

type ModificationConfig struct {
	DoIt bool `yaml:"doIt"`
	Actions []Actions `yaml:"actions"`
}

type Actions struct {
	ModifyHeaders struct {
		DoIt bool `yaml:"doIt"`
		HeadersToModify map[string]string `yaml:"headersToModify"`
	} `yaml:"modifyHeaders"`

	RemoveHeaders struct {
		DoIt bool `yaml:"doIt"`
		HeadersToRemove []string `yaml:"headersToRemove"`
	} `yaml:"removeHeaders"`
}

type ForwardRequestToConfig struct {
	BackendServiceRoute map[string]string `yaml:"forwardRequestTo"`
}