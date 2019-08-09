package packet

// CharErrorCode is a common type for all error codes returned
// during the character creation/deletion process.
//go:generate stringer -type=CharErrorCode -trimprefix=CharErrorCode
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
