package silkroad

import "errors"

var (
	errMissingClientParams        = errors.New("Client: Missing parameters for the Client. client ID or Secret cannot be empty.")
	errInvalidEnvironment         = errors.New("Client: Environment is not valid.")
	errInvalidJWTSigningMethod    = errors.New("Client: Invalid JWT Signing Method.")
	errInvalidTokenExpirationTime = errors.New("Client: Invalid TokenExpirationTime. Allowed range: 1-3600 seconds.")
	errClientNotAuthorized        = errors.New("Client: Client not authorized.")
	errJWTEncodingError           = errors.New("JWT: Encoding Error")
	errResponseError              = errors.New("HTTP: Response error")
	errURLParse                   = errors.New("HTTP: URL Parse Error")
	errJSONUnmarshalError         = errors.New("Encoding: JSON Unmarshal error")
)
