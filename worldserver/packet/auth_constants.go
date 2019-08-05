package packet

// AuthErrorCode is a common type for all error codes returned
// during the login process.
//go:generate stringer -type=AuthErrorCode
type AuthErrorCode uint8

// ErrorCodes used the login process.
const (
	AuthOK                  AuthErrorCode = 0x0C
	AuthFailed              AuthErrorCode = 0x0D
	AuthReject              AuthErrorCode = 0x0E
	AuthBadServerProof      AuthErrorCode = 0x0F
	AuthUnavailable         AuthErrorCode = 0x10
	AuthSystemError         AuthErrorCode = 0x11
	AuthBillingError        AuthErrorCode = 0x12
	AuthBillingExpired      AuthErrorCode = 0x13
	AuthVersionMismatch     AuthErrorCode = 0x14
	AuthUnknownAccount      AuthErrorCode = 0x15
	AuthIncorrectPassword   AuthErrorCode = 0x16
	AuthSessionExpired      AuthErrorCode = 0x17
	AuthServerShuttingDown  AuthErrorCode = 0x18
	AuthAlreadyLoggingIn    AuthErrorCode = 0x19
	AuthLoginServerNotFound AuthErrorCode = 0x1A
	AuthWaitQueue           AuthErrorCode = 0x1B
	AuthBanned              AuthErrorCode = 0x1C
	AuthAlreadyOnline       AuthErrorCode = 0x1D
	AuthNoTime              AuthErrorCode = 0x1E
	AuthDbBusy              AuthErrorCode = 0x1F
	AuthSuspended           AuthErrorCode = 0x20
	AuthParentalControl     AuthErrorCode = 0x21
)
