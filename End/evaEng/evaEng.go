package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
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

type Set struct {
	m map[string]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[string]bool{},
	}
}

func (s *Set) Add(item string) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := []string{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}





func main() {
	initDB()
	engin:=gin.Default()
	//跨域
	engin.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))



	routerGroup := engin.Group("/evaluate")
	routerGroup.POST("/insert", insertHandle)
	routerGroup.POST("/queryTeacher", queryTeacherHandle)
	routerGroup.POST("/queryHighRageTeacher", queryHighRateHandle)
	routerGroup.POST("/querySubject", querySubjectHandle)
	engin.Run(":8089")
}

type Classes struct {
	Subject string `from:"subject"`
	Teacher string `from:"teacher"`
	Evtext string `from:"evtext"`
	Rate int `from:"rate"`
	StudentId int `from:"studentid"`
}



func insertHandle(context *gin.Context) {
	fmt.Println(context.FullPath())
	var classes Classes
	if err:=context.BindJSON(&classes); err !=nil{
		log.Fatal(err.Error())
	}

	_,err := db.Exec("insert into evaluate(subject,teacher,evtext,rate,studentid)"+
		" values(?,?,?,?,?);" ,
		classes.Subject ,classes.Teacher,classes.Evtext,classes.Rate,classes.StudentId)

	if err != nil {
		log.Fatalln(err)
		context.JSON(200,map[string]interface{}{
			"code": 401,
			"message": "No",
		})
	}else{
		context.JSON(200,map[string]interface{}{
			"code": 200,
			"message": "OK",
		})
	}
	return

}

func queryTeacherHandle(context *gin.Context) {
	fmt.Println(context.FullPath())
	var classes Classes
	if err:=context.BindJSON(&classes); err !=nil{
		log.Fatal(err.Error())
	}
	teacherClasses := make([]Classes, 0)
	fmt.Println(classes.Subject)
	rows,err := db.Query("select subject,teacher,evtext,rate,studentid from evaluate where subject = ?", classes.Subject)

	if err != nil {
		log.Fatalln(err)
		context.JSON(200,map[string]interface{}{
			"code": 401,
			"message": "No",
		})
	}else{

		for rows.Next() {
			var tempClasses Classes
			rows.Scan(&tempClasses.Subject,&tempClasses.Teacher,&tempClasses.Evtext,&tempClasses.Rate,&tempClasses.StudentId)
			teacherClasses = append(teacherClasses,tempClasses)
			fmt.Println(teacherClasses)
		}
		context.JSON(200,map[string]interface{}{
			"code": 200,
			"message": "Ok",
			"data":teacherClasses,
		})
	}
	return

}

func queryHighRateHandle(context *gin.Context) {
	fmt.Println(context.FullPath())
	teacherClasses := make([]Classes, 0)
	rows,err := db.Query("select subject,teacher,evtext,rate,studentid from evaluate where rate = 5")

	if err != nil {
		log.Fatalln(err)
		context.JSON(200,map[string]interface{}{
			"code": 401,
			"message": "No",
		})
	}else{

		for rows.Next() {
			var tempClasses Classes
			rows.Scan(&tempClasses.Subject,&tempClasses.Teacher,&tempClasses.Evtext,&tempClasses.Rate,&tempClasses.StudentId)
			teacherClasses = append(teacherClasses,tempClasses)
			fmt.Println(teacherClasses)
		}
		context.JSON(200,map[string]interface{}{
			"code": 200,
			"message": "Ok",
			"data":teacherClasses,
		})
	}
	return

}

func querySubjectHandle(context *gin.Context) {
	fmt.Println(context.FullPath())
	//创建subject的集合
	s:=New()

	rows,err := db.Query("select subject from evaluate")

	if err != nil {
		log.Fatalln(err)
		context.JSON(200,map[string]interface{}{
			"code": 401,
			"message": "No",
		})
	}else{

		for rows.Next() {
			var subjectName string
			rows.Scan(&subjectName)
			s.Add(subjectName)
		}

		keys := s.List()

		fmt.Println(keys)
		context.JSON(200,map[string]interface{}{
			"code": 200,
			"message": "Ok",
			"data":keys,
		})
	}
	return

}
