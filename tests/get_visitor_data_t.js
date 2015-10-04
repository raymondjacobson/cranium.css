conn = new Mongo();
db = conn.getDB("cranium");

printjson(db.visitors.find().toArray());