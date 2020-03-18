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
