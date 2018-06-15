package models

import (
	"sync"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	Server   string
	Database string

	once sync.Once
	db   *mgo.Database
)

func init() {
	// Register driver.
	Server = beego.AppConfig.String("database::server")
	Database = beego.AppConfig.String("database::database")

}

var NewDB = func() {
	session, err := mgo.Dial(Server)
	if err != nil {
		logs.Error(err.Error())
	}
	db = session.DB(Database)
}

func GetDB() *mgo.Database {
	// Create new ormer.
	once.Do(NewDB)
	return db
}

func Insert(collection string, model interface{}) error {
	d := GetDB()
	err := d.C(collection).Insert(model)
	return err
}

func Read(collection string, id string, model interface{}) error {
	d := GetDB()
	err := d.C(collection).FindId(bson.ObjectIdHex(id)).One(model)
	return err
}

func ReadAll(collection string, model interface{}) error {
	d := GetDB()
	err := d.C(collection).Find(bson.M{}).All(model)
	return err
}

func Update(collection string, id string, model interface{}) error {
	d := GetDB()
	err := d.C(collection).UpdateId(bson.ObjectIdHex(id), model)
	return err
}

func Delete(collection string, model interface{}) error {
	d := GetDB()
	err := d.C(collection).Remove(model)
	return err
}
