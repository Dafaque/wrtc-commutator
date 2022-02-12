package errcodes

type ErrorCode uint8

var (
	ERROR_CODE_NONE                ErrorCode = 0
	ERROR_CODE_UNKNOWN             ErrorCode = 1
	ERROR_CODE_TARGET_UNACCESSABLE ErrorCode = 4
	ERROR_CODE_ALREADY_ONLINE      ErrorCode = 5
	ERROR_CODE_NOT_ONLINE          ErrorCode = 100
	ERROR_CODE_CANT_WRITE_BACK     ErrorCode = 101
	ERROR_CODE_MUMBLING            ErrorCode = 102
	ERROR_CODE_BAD_SIGNATURE       ErrorCode = 103
	ERROR_CODE_INVALID_SYNTAX      ErrorCode = 104
)
