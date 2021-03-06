package defs

type Err struct {
	ErrStr  string `err_str`
	ErrCode string `err_code`
}

type ErrResponse struct {
	HttpErr  int
	InnerErr Err
}

var (
	ErrParseFailed = ErrResponse{HttpErr: 400, InnerErr: Err{ErrStr: "parse element failed", ErrCode: "0001"}}
	ErrInvalidUser = ErrResponse{HttpErr: 401, InnerErr: Err{ErrStr: "user is not authentication", ErrCode: "0002"}}
	ErrInnerFailed = ErrResponse{HttpErr: 500, InnerErr: Err{ErrStr: "inner error", ErrCode: "0003"}}
	ErrDBErr       = ErrResponse{HttpErr: 500, InnerErr: Err{ErrStr: "database err", ErrCode: "0004"}}
)
