package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID            int64          `orm:"column(id)"`
	Name          string         `orm:"column(name)"`
	Age           int            `orm:"column(age)"`
	Gender        string         `orm:"column(gender)"`
	HobbyList     []string       `orm:"-"`
	HobbyListJSON string         `orm:"column(hobby_list)"`
	Scores        map[string]int `orm:"-"`
	ScoresJSON    string         `orm:"column(scores)"`
}

// database
var (
	DBDriver           = "mysql"
	DBConnectionString = "root:root@(localhost:3306)/goquickview?charset=utf8mb4&loc=Asia%2FShanghai"
	DBOpenConnCount    = 10
	DBIdleConnCount    = 10
	DBShowSQL          = true
)

func init() {
	// register database
	err := orm.RegisterDataBase("default", DBDriver, DBConnectionString, DBIdleConnCount, DBOpenConnCount)
	if err != nil {
		panic(fmt.Sprintf("Err: db init error, %s", err.Error()))
	}

	// register models
	orm.RegisterModel(new(Student))
}

func main() {
	jemy := Student{
		ID:        26015147,
		Name:      "金茧",
		Age:       29,
		Gender:    "男",
		HobbyList: []string{"吃饭", "睡觉", "工作", "玩耍"},
		Scores: map[string]int{
			"语文": 110,
			"数学": 120,
			"英语": 130,
		},
	}

	hobbyListJSON, _ := json.Marshal(jemy.HobbyList)
	jemy.HobbyListJSON = string(hobbyListJSON)

	scoresJSON, _ := json.Marshal(jemy.Scores)
	jemy.ScoresJSON = string(scoresJSON)

	_, insertErr := orm.NewOrm().Insert(&jemy)
	if insertErr != nil {
		fmt.Println("Err: wow!!!!", insertErr)
		return
	}

}
