package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:root1234@(38.6.164.167:13306)/test?charset=utf8&parseTime=True&loc=Local"
	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// //关闭数据库
	// defer db.Close()

	// 自动迁移
	db.AutoMigrate(&UserInfo{})
	u1 := UserInfo{1, "小明", "男", "篮球"}
	u2 := UserInfo{2, "娜扎", "女", "乒乓球"}

	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)
}
