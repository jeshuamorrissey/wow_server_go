package static

// CharErrorCode is a common type for all error codes returned
// during the character creation/deletion process.
type CharErrorCode uint8

// ErrorCodes used the character creation/deletion process.
const (
	CharErrorCodeCreateAccountLimit            CharErrorCode = 0x35
	CharErrorCodeCreateDisabled                CharErrorCode = 0x32
	CharErrorCodeCreateError                   CharErrorCode = 0x2F
	CharErrorCodeCreateFailed                  CharErrorCode = 0x30
	CharErrorCodeCreateInProgress              CharErrorCode = 0x2D
	CharErrorCodeCreateNameInUse               CharErrorCode = 0x31
	CharErrorCodeCreateOnlyExisting            CharErrorCode = 0x37
	CharErrorCodeCreatePvpTeamsViolation       CharErrorCode = 0x33
	CharErrorCodeCreateServerLimit             CharErrorCode = 0x34
	CharErrorCodeCreateServerQueue             CharErrorCode = 0x36
	CharErrorCodeCreateSuccess                 CharErrorCode = 0x2E
	CharErrorCodeDeleteFailed                  CharErrorCode = 0x3A
	CharErrorCodeDeleteFailedLockedForTransfer CharErrorCode = 0x3B
	CharErrorCodeDeleteInProgress              CharErrorCode = 0x38
	CharErrorCodeDeleteSuccess                 CharErrorCode = 0x39
	CharErrorCodeListFailed                    CharErrorCode = 0x2C
	CharErrorCodeListRetrieved                 CharErrorCode = 0x2B
	CharErrorCodeListRetrieving                CharErrorCode = 0x2A
	CharErrorCodeNameConsecutiveSpaces         CharErrorCode = 0x50
	CharErrorCodeNameFailure                   CharErrorCode = 0x51
	CharErrorCodeNameInvalidApostrophe         CharErrorCode = 0x4C
	CharErrorCodeNameInvalidCharacter          CharErrorCode = 0x48
	CharErrorCodeNameInvalidSpace              CharErrorCode = 0x4F
	CharErrorCodeNameMixedLanguages            CharErrorCode = 0x49
	CharErrorCodeNameMultipleApostrophes       CharErrorCode = 0x4D
	CharErrorCodeNameNoName                    CharErrorCode = 0x45
	CharErrorCodeNameProfane                   CharErrorCode = 0x4A
	CharErrorCodeNameReserved                  CharErrorCode = 0x4B
	CharErrorCodeNameSuccess                   CharErrorCode = 0x52
	CharErrorCodeNameThreeConsecutive          CharErrorCode = 0x4E
	CharErrorCodeNameTooLong                   CharErrorCode = 0x47
	CharErrorCodeNameTooShort                  CharErrorCode = 0x46
)

// AuthErrorCode is a common type for all error codes returned
// during the login process.
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
