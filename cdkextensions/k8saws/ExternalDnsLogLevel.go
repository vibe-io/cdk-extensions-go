package k8saws


// Verbosity of the logs generated by the external-dns service.
type ExternalDnsLogLevel string

const (
	// Set log level to 'panic'.
	ExternalDnsLogLevel_PANIC ExternalDnsLogLevel = "PANIC"
	// Set log level to 'debug'.
	ExternalDnsLogLevel_DEBUG ExternalDnsLogLevel = "DEBUG"
	// Set log level to 'info'.
	ExternalDnsLogLevel_INFO ExternalDnsLogLevel = "INFO"
	// Set log level to 'warning'.
	ExternalDnsLogLevel_WARNING ExternalDnsLogLevel = "WARNING"
	// Set log level to 'error'.
	ExternalDnsLogLevel_ERROR ExternalDnsLogLevel = "ERROR"
	// Set log level to 'fatal'.
	ExternalDnsLogLevel_FATAL ExternalDnsLogLevel = "FATAL"
	// Set log level to 'trace'.
	ExternalDnsLogLevel_TRACE ExternalDnsLogLevel = "TRACE"
)
