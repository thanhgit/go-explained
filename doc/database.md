# Database insider
- Go has some database interface standards for developers to develop database drivers with

## sql.Register in database/sql package
- used for registering database drivers when you use 3rd-party database drivers 
- Register(name string, driver driver.Driver)
```go
func init() {
    sql.Register('sqlite3', &SQLiteDriver{})
}
```
OR
```go
var d = Driver{proto: "tcp", raddr: "127.0.0.1:3306"}
func init() {
    Register("SET NAMES utf8")
    sql.Register("mymysql", &d)
}
```
- driver.Driver struct
```go
 type Driver interface {
    Open(name string) (Conn, error)
}
```
- driver.Conn struct
```go
type Conn interface {
    Prepare(query string) (Stmt, error)
    Close() error
    Begin() (Tx, error)
}
```
- driver.Stmt
```go
type Stmt interface {
    Close() error
    NumInput() int
    Exec(args []Value) (Result, error)
    Query(args []Value) (Rows, error)
}
```
- driver.Tx
```go
type Tx interface {
    Commit() error
    Rollback() error
}
```
- driver.Execer
```go
type Execer interface {
    Exec(query string, args []Value) (Result, error)
}
```
- driver.Result
```go
type Result interface {
    LastInsertId() (int64, error)
    RowsAffected() (int64, error)
}
```
- driver.Rows
```go
type Rows interface {
    Columns() []string
    Close() error
    Next(dest []Value) error
}
```
- driver.RowsAffected
```go
type RowsAffected int64

func (RowsAffected) LastInsertId() (int64, error)

func (v RowsAffected) RowsAffected() (int64, error)
```
- driver.Value
```go
type Value interface{}
```
- driver.ValueConverter
```go
type ValueConverter interface {
    ConvertValue(v interface{}) (Value, error)
}
// converts driver.Value to a user defined value in th scan() function
```
- driver.Valuer
```go
type Valuer interface {
    Value() (Value, error)
}
```
- connection pool
```go
type DB struct {
    driver   driver.Driver
    dsn      string
    mu       sync.Mutex // protects freeConn and closed
    freeConn []driver.Conn
    closed   bool
}
```
## MySQL
- https://github.com/go-sql-driver/mysql
- https://github.com/ziutek/mymysql
## Process
* ### Open connection
- Schema:
    ```text
    user@unix(/path/to/socket)/dbname?charset=utf8
    user:password@tcp(localhost:5555)/dbname?charset=utf8
    user:password@/dbname
    user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
    ```
```go
db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/test?charset=utf8")
checkErr(err)
```
* ### Insert data
```go
stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
checkErr(err)

res, err := stmt.Exec("astaxie", "研发部门","2012-12-09")
checkErr(err)

id, err := res.LastInsertId()
checkErr(err)
```
* ### Update data
```go
stmt, err = db.Prepare("update userinfo set username=? where uid=?")
checkErr(err)

res, err = stmt.Exec("astaxieupdate", id)
checkErr(err)

affect, err := res.RowsAffected()
checkErr(err)

fmt.Println(affect)
```
* ### Query data
```go
rows, err := db.Query("SELECT * FROM userinfo")
checkErr(err)

for rows.Next() {
    var uid int
    var username string
    var department string
    var created string
    err = rows.Scan(&uid, &username, &department, &created)
    checkErr(err)
    fmt.Println(uid)
    fmt.Println(username)
    fmt.Println(department)
    fmt.Println(created)
}
```
* ### Delete data
```go
stmt, err = db.Prepare("delete from userinfo where uid=?")
checkErr(err)

res, err = stmt.Exec(id)
checkErr(err)

affect, err = res.RowsAffected()
checkErr(err)

fmt.Println(affect)
```
* ### Close connection
```go
db.Close()
```
* ### Utils function
```go
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
```
# NOSQL
## Redis
- github.com/garyburd/redigo/redis
* ### Open connection 
```go
 var (
    Pool *redis.Pool
)

func init() {
    redisHost := ":6379"
    Pool = newPool(redisHost)
    close()
}

func newPool(server string) *redis.Pool {

    return &redis.Pool{

        MaxIdle:     3,
        IdleTimeout: 240 * time.Second,

        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", server)
            if err != nil {
                return nil, err
            }
            return c, err
        },

        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
    }
}
```
* ### Close connection
```go
func close() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, syscall.SIGTERM)
    signal.Notify(c, syscall.SIGKILL)
    go func() {
        <-c
        Pool.Close()
        os.Exit(0)
    }()
}
```
* ### Get value with key
```go
func Get(key string) ([]byte, error) {

    conn := Pool.Get()
    defer conn.Close()

    var data []byte
    data, err := redis.Bytes(conn.Do("GET", key))
    if err != nil {
        return data, fmt.Errorf("error get key %s: %v", key, err)
    }
    return data, err
}
```
* ### Set key-value
```go
func Set(key string, value string) error {
    conn := Pool.Get()
    defer conn.Close()

    var data []byte
    res, err := redis.Bytes(conn.Do("APPEND", key, value))
    if err != nil {
        return fmt.Errorf("error set key %s: %v", key, err)
    }

    return err
}
```
## MongoDB
- github.com/globalsign/mgo
* ### Config Database
```go
package config

type config struct {
	MongoDBHost string
	MongoDBUser string
	MongoDBPwd  string
	Database    string
}

var (
	AppConfig config
)

func init() {
	AppConfig = config{
		MongoDBHost: "localhost:27017",
		MongoDBUser: "",
		MongoDBPwd: "",
		Database: "demo",
	}
}

//---
package config

import (
	"github.com/globalsign/mgo"
	"log"
	"time"
)

type Session struct {
	session *mgo.Session
}

func NewSession() *Session {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.MongoDBUser,
		Password: AppConfig.MongoDBPwd,
		Timeout:  60 * time.Second,
	})

	if err != nil {
		log.Fatalf("[ConnectDB]: %s\n", err)
	}
	session.SetMode(mgo.Monotonic, true)

	return &Session{session}
}

func (s *Session) Copy() *mgo.Session {
	return s.session.Copy()
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
```
* ### Model
```go
package model

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
)

type User struct {
	Id 			bson.ObjectId 	`bson:"_id,omitempty" json:"id"`
	Username	string			`json:"username"`
	Password	string 			`json:"password"`
}

func (user *User) String() string  {
	data, _ :=json.Marshal(&user)
	return string(data)
}
```
### Reposiroty
```go
package repository

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/thanhgit/realtime-mongodb/config"
	"github.com/thanhgit/realtime-mongodb/model"
)

type IUserRepo interface {
	GetAllUser() []model.User
}

type UserRepository struct {}

func (userRepo *UserRepository)GetAllUser() []model.User  {
	s := config.NewSession()
	defer s.Close()

	c := s.Copy().DB(config.AppConfig.Database).C("user")
	users := []model.User{}
	c.Find(nil).All(&users)

	return users
}
//-- realtime
func (userRepo *UserRepository)ChangeStreamWatcher(ctx context.Context) {
	s := config.NewSession()
	defer s.Close()

	coll := s.Copy().DB(config.AppConfig.Database).C("user")

	changeStream, err := coll.Watch(nil, mgo.ChangeStreamOptions{
		BatchSize: 10,
		//data can come from the stream simultaneously
		//this parameter is like a buffer size
	})
    defer changeStream.Close()
    
	if err != nil {
		fmt.Errorf("Failed to open change stream")
		return
	}

	//Handling change stream in a cycle
	for {
		select {
		case <-ctx.Done(): //if parent context was cancelled
			err := changeStream.Close()
			if err != nil {
				fmt.Errorf("Change stream closed")
			}
			return
		default:
			//making a struct for unmarshalling
			changeDoc := struct {
				FullDocument model.User `bson:"fullDocument"`
			}{}

			//getting next item from the steam
			ok := changeStream.Next(&changeDoc)

			//if data from the stream wasn't unmarshaled, we get ok == false as a result
			//so we need to call Err() method to get info why
			//it'll be nil if we just have no data
			if !ok {
				err := changeStream.Err()
				if err != nil {
					return
				}
			}

			if ok {
                elem := changeDoc.FullDocument
                // Result
				println(elem.String())
			}
		}

	}
}
```
* ### Find max-depth of directory
```text
find . -type d -printf '%d\n' | sort -rn | head -1
```

* ### CRUD Gorm with gin framework
    - Get all
    ```go
     var people []Person
     if err := db.Find(&people).Error; err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
     } else {
        c.JSON(200, people)
     }
    ```
    - Get by id
    ```go
     if err := db.Where(“id = ?”, id).First(&person).Error; err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
     } else {
        c.JSON(200, person)
     }
    ```
    - Create
    ```go
     var person Person
     c.BindJSON(&person)
     db.Create(&person)
     c.JSON(200, person)
    ```
    - Update
    ```go
     var person Person
     id := c.Params.ByName(“id”)
     if err := db.Where(“id = ?”, id).First(&person).Error; err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
     }
     c.BindJSON(&person)
     db.Save(&person)
     c.JSON(200, person)
    ```
    - Delete
    ```go
     id := c.Params.ByName(“id”)
     var person Person
     d := db.Where(“id = ?”, id).Delete(&person)
     fmt.Println(d)
     c.JSON(200, gin.H{“id #” + id: “deleted”})
    ````