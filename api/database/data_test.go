package data

import (
	"fmt"
	"testing"
)

func clearTables() {
	dbConn.Exec("truncate users ")
	dbConn.Exec("truncate sessions ")
	dbConn.Exec("truncate video_info ")
	dbConn.Exec("truncate comments")

}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	//clearTables()
}

func TestUserAll(t *testing.T) {
	t.Run("add", testAddUser)
	//t.Run("del", testDelUser)
}
func TestVideoAll(t *testing.T) {
	t.Run("add", testAddVideoInfo)
	//t.Run("get", testGetVideo)
	//t.Run("delete", testDelVideo)
}

func TestComment(t *testing.T) {
	t.Run("add", testAddComment)
	t.Run("list", testGetComment)
}
func testAddUser(t *testing.T) {
	fmt.Println("now test add users")
	err := AddUserCredential("pengbb", "peng123456")
	if err != nil {
		t.Errorf("Add user failed")
		fmt.Println("err:", err)
	}
}

func testDelUser(t *testing.T) {
	fmt.Println("now test delete users")
	err := DelUser("pengbb", "peng123456")
	if err != nil {
		fmt.Println("err:", err)
		t.Errorf("delet user faild")
	}
}

var tempUUID string

func testAddVideoInfo(t *testing.T) {
	res, err := AddVideo(1, "youku")
	tempUUID = res.ID
	if err != nil {
		fmt.Println("add video failed")
		t.Errorf("add video failed")
	}
}
func testDelVideo(t *testing.T) {
	err := DeleVideo(tempUUID)
	if err != nil {
		fmt.Println("delete video failed")
		t.Errorf("delete video failed")
	}
}

func testGetVideo(t *testing.T) {
	res, err := GetOneVideo(tempUUID)
	if err != nil {
		fmt.Println("delete video failed")
		t.Errorf("delete video failed")
	}
	fmt.Println("ID:", res.ID, "name:", res.Name, "authorID:", res.AuthorID, "time:", res.DisplayTime)
}

func testAddComment(t *testing.T) {
	err := AddNewComments(tempUUID, 1, "i like it")
	if err != nil {
		fmt.Println("delete video failed")
		t.Errorf("delete video failed")
	}

}
func testGetComment(t *testing.T) {
	res, err := ListComments(tempUUID, 84590672, 1584591672)
	if err != nil {
		fmt.Println("list comment failed:", err)
		t.Errorf("list comment failed")
	}
	fmt.Println("res:", res[0].Content)
}
