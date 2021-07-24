package contextkeys

// ContextKey is really just a string
type ContextKey string

const (

	// XRequestID used to get xrequestID in context
	XRequestID ContextKey = "x-request-id"
)
