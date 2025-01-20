package types

type HeadersConfig struct {
	ModifyRequest ModificationConfig `yaml:"modifyRequest"`
	ModifyResponse ModificationConfig `yaml:"modifyResponse"`
}

type ModificationConfig struct {
	DoIt bool `yaml:"doIt"`
	Actions Actions `yaml:"actions"`
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

func (config *HeadersConfig) CanModifyRequest() bool { return config.ModifyRequest.DoIt }
func (config *HeadersConfig) CanModifyResponse() bool { return config.ModifyResponse.DoIt }

func (config *HeadersConfig) CanModifyRequestHeader() bool { return config.CanModifyRequest() && config.ModifyRequest.Actions.ModifyHeaders.DoIt }
func (config *HeadersConfig) CanModifyResponseHeader() bool { return config.CanModifyResponse() && config.ModifyResponse.Actions.ModifyHeaders.DoIt }

func (config *HeadersConfig) CanRemoveRequestHeader() bool {return config.CanModifyRequestHeader() && config.ModifyRequest.Actions.RemoveHeaders.DoIt }
func (config *HeadersConfig) CanRemoveResponseHeader() bool {return config.CanModifyResponseHeader() && config.ModifyResponse.Actions.RemoveHeaders.DoIt }

func (config *HeadersConfig) GetRequestModifyHeadersMap() map[string]string { 
	if config.CanModifyRequest() {
		return config.ModifyRequest.Actions.ModifyHeaders.HeadersToModify
	}
	return nil
}

func (config *HeadersConfig) GetResponseModifyHeadersMap() map[string]string { 
	if config.CanModifyResponse() {
		return config.ModifyResponse.Actions.ModifyHeaders.HeadersToModify
	}
	return nil
}

func (config *HeadersConfig) GetRequestHeadersToRemove() []string {
	if config.CanModifyRequest() {
		return config.ModifyRequest.Actions.RemoveHeaders.HeadersToRemove
	}
	return nil
}

func (config *HeadersConfig) GetResponseHeadersToRemove() []string {
	if config.CanModifyResponse() {
		return config.ModifyResponse.Actions.RemoveHeaders.HeadersToRemove
	}
	return nil
}