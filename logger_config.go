package gologger

// Config for configuring logger
type Config struct {

	// Log Level
	LogLevel uint8

	// File flush for determining to flush log to file or not
	FileFlush bool

	// LogDir directory at which log files will be flushed
	LogDir string

	// FormatString configurable how user wants
	// TODO: need to add support in near future
	FormatString string

	// Pretty print for logging arrays and struct default false
	PrettyPrint bool
}
