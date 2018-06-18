package models

import (
	"sync"

	"gopkg.in/go-playground/validator.v9"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	// Server Server string.
	Server string
	// Database Database string.
	Database string
	// Validator Models validator.
	Validator *validator.Validate

	once sync.Once
	db   *mgo.Database
)

// Dao interface.
type Dao interface {
}

// Dao struct definition.
type dao struct {
}

func init() {
	// Register driver.
	Server = beego.AppConfig.String("database::server")
	Database = beego.AppConfig.String("database::database")

	// Init validator.
	Validator = validator.New()
}

// newDB Creates a new session.
var newDB = func() {
	session, err := mgo.Dial(Server)
	if err != nil {
		logs.Error(err.Error())
	}
	db = session.DB(Database)
}

// GetDB Return database session.
func GetDB() *mgo.Database {
	// Create new database session.
	once.Do(newDB)
	return db
}

// Insert Insert model in to database collection.
func Insert(collection string, model interface{}) error {
	d := GetDB()
	err := d.C(collection).Insert(model)
	return err
}

// Read Read from database collection.
func Read(collection string, id string, model interface{}) error {
	d := GetDB()
	err := d.C(collection).FindId(bson.ObjectIdHex(id)).One(model)
	return err
}

// ReadAll Read all from database collection.
func ReadAll(collection string, model interface{}) error {
	d := GetDB()
	err := d.C(collection).Find(bson.M{}).All(model)
	return err
}

// Update Update model.
func Update(collection string, id string, model interface{}) error {
	d := GetDB()
	err := d.C(collection).UpdateId(bson.ObjectIdHex(id), model)
	return err
}

// Delete Delete model.
func Delete(collection string, id string) error {
	d := GetDB()
	err := d.C(collection).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
