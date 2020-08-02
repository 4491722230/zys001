package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"net/http"
)

//声明全局变量
var (
	mysql_url = "root:root@/a?charset=utf8&parseTime=True&loc=Local"
	DB, _     = gorm.Open("mysql", mysql_url)
	// 初始化admin 还有其他的，比如API
	API = admin.New(&qor.Config{DB: DB})
)

func init() {
	//用户
	type User struct {
		gorm.Model
		Email    string
		Password string
		Name     sql.NullString
		Gender   string
		Active   bool
	}
	DB.AutoMigrate(&User{}) //自动创建表。
	// Add it to Admin
	user := API.AddResource(&User{})
	//user.Action(&admin.Action{
	//	Name: "enable",
	//	Handler: func(actionArgument *admin.ActionArgument) error {
	//		// `FindSelectedRecords` => in bulk action mode, will return all checked records, in other mode, will return current record
	//		for _, record := range actionArgument.FindSelectedRecords() {
	//			actionArgument.Context.DB.Model(record.(*models.User)).Update("Active", true)
	//		}
	//		return nil
	//	},
	//})
	fmt.Println(user)
}

func main() {
	// 启动服务
	mux := http.NewServeMux()
	API.MountTo("/admin", mux)
	http.ListenAndServe(":9000", mux)
}
