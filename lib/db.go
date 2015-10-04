package db

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
    "model"
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

func InsertNewDataEntry(db *mgo.Database, vid string,
						atags []model.Atag, ptags []model.Ptag,
						imgtags []model.Imgtag) model.DataEntry {

	c := db.C("visitors")
<<<<<<< HEAD:lib/crnm/db/db.go
	a := []model.Atag{model.AtagDefault}
	p := []model.Ptag{model.PtagDefault}
	img := []model.Imgtag{model.Imgtag{"test", 15, 15, 0, 0, 0}}
	de := model.DataEntry{a, p, img}
=======
	de := model.DataEntry{atags, ptags, imgtags}
>>>>>>> architecture:lib/db/db.go
	c.Update(bson.M{"vid": vid}, bson.M{"$addToSet": bson.M{"data" : &de}})
	return de
}

func FetchMostRecentDataEntry(db *mgo.Database, vid string) model.Visitor {
	c := db.C("visitors")
	visitor := model.Visitor{}
	c.Find(bson.M{"vid": vid}).One(&visitor)
	return visitor
}