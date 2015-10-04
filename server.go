package main

import (
  // "fmt"
  "github.com/gin-gonic/gin"
  "./lib/crnm/db"
  // "./lib/crnm/css_gen"
  // "./lib/crnm/model"

)

func main() {

  s := db.Connect()
  defer db.Disconnect(s)
  cranium := db.SetDB(s)
  db.InsertNewVisitor(cranium, "A8b839013jgkke")
  db.InsertNewDataEntry(cranium, "A8b839013jgkke")

  // TODO: Call Peter's code

  r := gin.Default()
  r.StaticFile("/", "./index.html")
  r.Run(":1225")
}