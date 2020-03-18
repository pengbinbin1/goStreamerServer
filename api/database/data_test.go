package data

import (
	"fmt"
	"testing"
)

func clearTables() {
	dbConn.Exec("truncate users ")
	dbConn.Exec("truncate sessions ")
	dbConn.Exec("truncate video_info ")

}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestAll(t *testing.T) {
	t.Run("add", testAddUser)
	t.Run("del", testDelUser)
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
