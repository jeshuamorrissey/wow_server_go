package static

// LoginErrorCode is a common type for all error codes returned
// during the login process.
type LoginErrorCode uint8

// ErrorCodes used the login process.
const (
	LoginOK              LoginErrorCode = 0x00
	LoginFailed          LoginErrorCode = 0x01 // "unable to connect"
	LoginFailed2         LoginErrorCode = 0x02 // "unable to connect"
	LoginBanned          LoginErrorCode = 0x03 // "this account has been closed"
	LoginUnknownAccount  LoginErrorCode = 0x04 // "information is not valid"
	LoginUnknownAccount3 LoginErrorCode = 0x05 // "information is not valid"
	LoginAlreadyOnline   LoginErrorCode = 0x06 // "this account is already logged in"
	LoginNoTime          LoginErrorCode = 0x07 // "you have no time left on this account"
	LoginDBBusy          LoginErrorCode = 0x08 // "could not log in at this time, try again later"
	LoginBadVersion      LoginErrorCode = 0x09 // "unable to validate game version"
	LoginDownloadFile    LoginErrorCode = 0x0A
	LoginFailed3         LoginErrorCode = 0x0B // "unable to connect"
	LoginSuspended       LoginErrorCode = 0x0C // "this account has been temporarily suspended"
	LoginFailed4         LoginErrorCode = 0x0D // "unable to connect"
	LoginConnected       LoginErrorCode = 0x0E
	LoginParentalControl LoginErrorCode = 0x0F // "blocked by parental controls"
	LoginLockedEnforced  LoginErrorCode = 0x10 // "disconnected from server"
)

// Game version information which must be present in the
// ClientLoingChallenge packet for a connection to be established.
const (
	SupportedGameName  = "WoW"
	SupportedGameBuild = 5875
)

// Game version information (which has to be variable because it is a slice).
var (
	SupportedGameVersion = [3]uint8{1, 12, 1}
)
