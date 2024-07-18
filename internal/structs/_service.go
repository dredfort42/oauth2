package structs

// ServiceParamethers is a struct for service configuration
type ServiceParamethers struct {
	Host                             string
	Port                             string
	CorsStatus                       string
	DeviceVerificationURI            string
	DeviceVerificationCodeCharSet    string
	DeviceVerificationCodeLength     int
	DeviceVerificationCodeExpiration int
	DeviceVerificationCodeAttempts   int
}

// ServiceConfig holds service configuration
var ServiceConfig ServiceParamethers
