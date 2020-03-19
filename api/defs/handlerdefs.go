package defs

type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type VideoInfo struct {
	ID          string
	AuthorID    int
	Name        string
	CreateTime  string
	DisplayTime string
}

type Comment struct {
	ID         string
	VideoID    string
	AuthorName string
	Content    string
}

type SimpleSession struct {
	UserName string
	TTL      int64
}

type SingUp struct {
	Success   bool
	SessionID string
}
