package main

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
	"fmt"
)

type UserInfo struct {
	Uid int
	Username string
	Department string
	Created time.Time
}

func main() {

	db, err := sql.Open("mymysql", "test/xiemgjun/123456")
	if err != nil {
		panic(err)
	}

	orm := beedb.New(db)
	beedb.OnDebug = true

	// Create a dummy user
	newUser := &UserInfo{Uid: 1, Username: "larry", Department: "Cinema", Created: time.Now()}

	insertUser(&orm, newUser)

	fmt.Println(newUser.Uid)

	// Update the user
	t := make(map[string]interface{})
	t["Username"] = "seedollar"
	orm.SetTable("userinfo").SetPK("uid").Where(1).Update(t)




}

func insertUser(orm *beedb.Model, user *UserInfo) {
		orm.Save(user)
}
