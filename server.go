package main

import (
    "fmt"
    "github.com/gin-gonic/contrib/sessions"
    "github.com/gin-gonic/gin"
    "net/http"
    "math/rand"
    "os"
    "time"
    "./lib"
    "./lib/db"
    "./lib/model"
)

func randomString(n int) string {
    rand.Seed(time.Now().UnixNano())
    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func main() {

    s := db.Connect()
    defer db.Disconnect(s)
    craniumDB := db.SetDB(s)

    r := gin.Default()

    // Set up Gin cookies
    store := sessions.NewCookieStore([]byte("pladjamaicanwifflepoof"))
    r.Use(sessions.Sessions("cranium_session", store))

    var title string;

    r.LoadHTMLGlob("templates/*")
    r.GET("/", func(c *gin.Context) {
        session := sessions.Default(c)
        craniumId := session.Get("cranium_id")
        if craniumId == nil {
            // Insert a new user if the user has no cookie
            newCraniumId := randomString(30);
            db.InsertNewVisitor(craniumDB, newCraniumId)
            session.Set("cranium_id", newCraniumId)
            session.Save()
            title = newCraniumId
            readfile, _ := os.Open("templates/index-tmpl.html")
            outfile, _ := os.Create("templates/index.html")
            defer readfile.Close()
            defer outfile.Close()
            // Get all the Ids for tags in question
            aImpIds, pImpIds, imgImpIds, aNimpIds, pNimpIds, imgNimpIds := preprocess.GenerateTags(readfile, outfile)
            fmt.Println(aImpIds, pImpIds, imgImpIds, aNimpIds, pNimpIds, imgNimpIds)
            var atags []model.Atag
            var ptags []model.Ptag
            var imgtags []model.Imgtag

            for _, impId := range aImpIds {
                atag := model.AtagDefault;
                atag.Id = impId
                atag.Important = true
                atags = append(atags, atag)
            }
            for _, impId := range pImpIds {
                ptag := model.PtagDefault;
                ptag.Id = impId
                ptag.Important = true
                ptags = append(ptags, ptag)
            }
            for _, impId := range imgImpIds {
                imgtag := model.ImgtagDefault;
                imgtag.Id = impId
                imgtag.Important = true
                imgtags = append(imgtags, imgtag)
            }
            for _, impId := range aNimpIds {
                atag := model.AtagDefault;
                atag.Id = impId
                atag.Important = false
                atags = append(atags, atag)
            }
            for _, impId := range pNimpIds {
                ptag := model.PtagDefault;
                ptag.Id = impId
                ptag.Important = false
                ptags = append(ptags, ptag)
            }
            for _, impId := range imgNimpIds {
                imgtag := model.ImgtagDefault;
                imgtag.Id = impId
                imgtag.Important = false
                imgtags = append(imgtags, imgtag)
            }
            de := db.InsertNewDataEntry(craniumDB, newCraniumId, atags, ptags, imgtags)

            fmt.Println(de)
        } else {
            // The user already exists, so ask the database for attributes
            title = craniumId.(string)
            visitor := db.FetchMostRecentDataEntry(craniumDB, craniumId.(string))
            fmt.Println(visitor.Data[0])
        }
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "cranium.css | " + title,
        })
    })

    r.Run(":1225")
}