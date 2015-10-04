package db

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
    "../model"
)

func Connect() *mgo.Session {
    session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }
    return session;
}

func SetDB(s *mgo.Session) *mgo.Database {
    return s.DB("cranium")
}

func Disconnect(s *mgo.Session) {
    s.Close()
}

func InsertNewVisitor(db *mgo.Database, vid string) {
    c := db.C("visitors")
    err := c.Insert(&model.Visitor{Vid: vid})
    if err != nil {
        log.Fatal(err)
    }
}

func InsertNewDataEntry(db *mgo.Database, vid string) {
	c := db.C("visitors")
	a := []model.Atag{model.AtagDefault}
	p := []model.Ptag{model.PtagDefault}
	img := []model.Imgtag{model.Imgtag{"test", 15, 15, 0, 0, 0}}
	de := model.DataEntry{a, p, img}
	c.Update(bson.M{"vid": vid}, bson.M{"$addToSet": bson.M{"data" : &de}})
}