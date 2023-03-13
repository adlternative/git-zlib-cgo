package git_zlib_cgo

// Allowed flush values
const (
	Z_NO_FLUSH      = 0
	Z_PARTIAL_FLUSH = 1
	Z_SYNC_FLUSH    = 2
	Z_FULL_FLUSH    = 3
	Z_FINISH        = 4
	Z_BLOCK         = 5
	Z_TREES         = 6
)

// Return codes
const (
	Z_OK            = 0
	Z_STREAM_END    = 1
	Z_NEED_DICT     = 2
	Z_ERRNO         = -1
	Z_STREAM_ERROR  = -2
	Z_DATA_ERROR    = -3
	Z_MEM_ERROR     = -4
	Z_BUF_ERROR     = -5
	Z_VERSION_ERROR = -6
)
