package types

type LoggingConfig struct {
	LoggerConfig struct {
		LogLevel string `yaml:"logLevel"`
		FileName string `yaml:"fileName"`
	} `yaml:"loggerConfig"`
}

func (config *LoggingConfig) GetLoggingLevel() string { return config.LoggerConfig.LogLevel }
func (config *LoggingConfig) SetLoggingLevel(logLevel string) { config.LoggerConfig.LogLevel = logLevel }

func (config *LoggingConfig) GetFileName() string { return config.LoggerConfig.FileName }
func (config *LoggingConfig) SetFileName(fileName string) { config.LoggerConfig.FileName = fileName }