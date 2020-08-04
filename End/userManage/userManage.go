package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)
var db *sql.DB
func initDB()(err error)  {
	/*数据库信息*/
	dsn:= ""
		/*已做加密处理*/
	db,err = sql.Open("mysql",dsn) //不会校验用户名密码是否正确

	if err!=nil{ //格式不正确
		return
	}

	err = db.Ping()
	if err!=nil{
		return
	}

	db.SetMaxOpenConns(10)
	return
}

func main() {
	initDB()
	engin:=gin.Default()
	engin.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))



	routerGroup := engin.Group("/student")


	routerGroup.POST("/register", registerHandle)
	routerGroup.POST("/login", loginHandle)
	routerGroup.POST("/changeInfo", changeInfoHandle)
	routerGroup.POST("/changePd", changePdHandle)



	engin.Run(":8088")

}

type Person struct {
	Name string `from:"name" `
	Schoolid int `from:"schoolid"`
	Password string `from:"password"`
	Phonenumber int `from:"phonenumber"`

}

type chPerson struct {
	Name string `from:"name" `
	Schoolid int `from:"schoolid"`
	Password string `from:"password"`
	Phonenumber int `from:"phonenumber"`
	Oldpassword string `from:"oldpassword"`
}



func changeInfoHandle(context *gin.Context) {
	fmt.Println(context.FullPath())
	var chperson chPerson
	if err:=context.BindJSON(&chperson); err !=nil{
		log.Fatal(err.Error())
	}

	_,err := db.Exec("update student set name= ?,phonenumber=? where schoolid = ?;" ,
		chperson.Name ,chperson.Phonenumber,chperson.Schoolid)
	if err != nil {
		fmt.Println(err)
		context.JSON(200,map[string]interface{}{
			"code": 402, /*未知错误*/
			"message": "未知错误",
		})
		return
	}else{
		context.JSON(200,map[string]interface{}{
			"code": 200,
			"message": "OK",
		})
		return
	}




}


func changePdHandle(context *gin.Context) {
	fmt.Println(context.FullPath())
	var chperson chPerson
	if err:=context.BindJSON(&chperson); err !=nil{
		log.Fatal(err.Error())
	}

	//首先查询用户和密码
	var confPasswd = ""
	db.QueryRow("SELECT password FROM student WHERE schoolid = ?",chperson.Schoolid).Scan(&confPasswd)

	if confPasswd!=chperson.Oldpassword{
		context.JSON(200,map[string]interface{}{
			"code": 401,
			"message": "原来的密码错误",
		})
		return
	} else{
		_,err := db.Exec("update student set password = ? where schoolid = ?;" ,
			chperson.Password  ,chperson.Schoolid)
		if err != nil {
			fmt.Println(err)
			context.JSON(200,map[string]interface{}{
				"code": 402, /*未知错误*/
				"message": "未知错误",
			})
			return
		}else{
			context.JSON(200,map[string]interface{}{
				"code": 200,
				"message": "OK",
			})
			return
		}

	}


}





func registerHandle(context *gin.Context) {
	fmt.Println(context.FullPath())
	var person Person
	if err:=context.BindJSON(&person); err !=nil{
		log.Fatal(err.Error())
	}

	//首先查询用户是否存在
	var confId = ""
	db.QueryRow("SELECT name FROM student WHERE schoolid = ?",person.Schoolid).Scan(&confId)

	fmt.Println(confId)
	if confId!=""{
		context.JSON(200,map[string]interface{}{
			"code": 401, /*用户已经存在*/
			"message": "用户已经存在",
		})
		return
	} else{
		_,err := db.Exec("insert into student(name,schoolid,password,phonenumber)"+
			" values(?,?,?,?);" ,
			person.Name,person.Schoolid,person.Password,person.Phonenumber)
		if err != nil {
			context.JSON(200,map[string]interface{}{
				"code": 402, /*未知错误*/
				"message": "未知错误",
			})
			return
		}else{
			context.JSON(200,map[string]interface{}{
				"code": 200,
				"message": "OK",
			})
			return
		}

	}


}

func loginHandle(context *gin.Context) {
	var person Person
	var person2 Person
	fmt.Println(context.FullPath())
	if err:=context.BindJSON(&person); err !=nil{
		log.Fatal(err.Error())
	}
	fmt.Println(person)
	var oriSchoolid = person.Schoolid
	var oriPassword = person.Password
	var aimPassword string

	err := db.QueryRow("SELECT password FROM student WHERE schoolid = ?",oriSchoolid).Scan(&person2.Password)

	aimPassword=person2.Password
	if person2.Password==""{
		fmt.Println("no")
		context.JSON(200,map[string]interface{}{
			"code": 401, /*用户未注册*/
			"message": "未注册",
			"data": oriSchoolid,
		})
		return
	}

	if err != nil {
		log.Fatalln(err)
		return
	}
	if oriPassword == aimPassword {
		fmt.Println("yes")
		context.JSON(200,map[string]interface{}{
			"code": 200, /*登入成功*/
			"message": "OK",
			"data": oriSchoolid,
		})
	} else
	{
		fmt.Println("no")
		context.JSON(200,map[string]interface{}{
			"code": 402, /*密码错误*/
			"message": "密码错误",
			"data": oriSchoolid,
		})
	}




}


