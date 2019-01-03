package client

// clientError implements net.Error
type clientError struct {
	errString string
	reconnect bool
	timeout   bool
	auth      bool
}

func (c clientError) Error() string {
	return c.errString
}

func (c clientError) Temporary() bool {
	return c.reconnect
}

func (c clientError) Timeout() bool {
	return c.timeout
}

var (
	// ErrRetryTimedOut is returned when Reconnect() time exceeds MaxElapsedTime.
	ErrRetryTimedOut = clientError{timeout: true, errString: "retry timed out"}

	// ErrBadToken is returned when a usable token can not be generated by the authorizer.
	ErrBadToken = clientError{errString: "bad auth token"}

	// ErrRetryFailed is returned when retry attempts fail.
	ErrRetryFailed = clientError{errString: "retry failed"}

	// ErrClientReconnecting is returned when the connection is reconnecting.
	// This is a temporary error, and callers should retry the operation after
	// a delay.
	ErrClientReconnecting = clientError{errString: "client reconnecting", reconnect: true}

	// ErrClientClosed is returned from an Accept call when the client is closed.
	ErrClientClosed = clientError{errString: "client closed"}

	// ErrAuthFailed is returned when authentication with the proxy fails
	ErrAuthFailed = clientError{errString: "auth failed", auth: true}
)