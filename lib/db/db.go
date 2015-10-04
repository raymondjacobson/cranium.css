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

func InsertNewDataEntry(db *mgo.Database, vid string,
						atags []model.Atag, ptags []model.Ptag,
						imgtags []model.Imgtag) model.DataEntry {

	c := db.C("visitors")
	de := model.DataEntry{atags, ptags, imgtags}
	c.Update(bson.M{"vid": vid}, bson.M{"$addToSet": bson.M{"data" : &de}})
	return de
}

func FetchVisitor(db *mgo.Database, vid string) model.Visitor {
	c := db.C("visitors")
	visitor := model.Visitor{}
	c.Find(bson.M{"vid": vid}).One(&visitor)
	return visitor
}

func UpdateAtagField(db *mgo.Database, vid string, tag model.Tag) {
    c := db.C("visitors")
    c.Update(bson.M{"vid": vid, "data.0.atags.id": tag.Id},
             bson.M{"$set": bson.M{"data.0.atags.$.hover": tag.Hover,
                                   "data.0.atags.$.click": tag.Click,
                                   "data.0.atags.$.frametime": tag.Frame, 
                                  }})
}

func UpdatePtagField(db *mgo.Database, vid string, tag model.Tag) {
    c := db.C("visitors")
    c.Update(bson.M{"vid": vid, "data.0.ptags.id": tag.Id},
             bson.M{"$set": bson.M{"data.0.ptags.$.hover": tag.Hover,
                                   "data.0.ptags.$.click": tag.Click,
                                   "data.0.ptags.$.frametime": tag.Frame, 
                                  }})
}

func UpdateImgtagField(db *mgo.Database, vid string, tag model.Tag) {
    c := db.C("visitors")
    c.Update(bson.M{"vid": vid, "data.0.imgtags.id": tag.Id},
             bson.M{"$set": bson.M{"data.0.imgtags.$.hover": tag.Hover,
                                   "data.0.imgtags.$.click": tag.Click,
                                   "data.0.imgtags.$.frametime": tag.Frame, 
                                  }})
}