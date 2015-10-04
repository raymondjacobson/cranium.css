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
	cranium := db.SetDB(s)

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
			db.InsertNewVisitor(cranium, newCraniumId)
			session.Set("cranium_id", newCraniumId)
			session.Save()
			title = newCraniumId
			readfile, _ := os.Open("templates/index-tmpl.html")
  			outfile, _ := os.Create("templates/index.html")
  			defer readfile.Close()
  			defer outfile.Close()
			imp_ids, nimp_ids := preprocess.GenerateTags(readfile, outfile)
			fmt.Println(imp_ids)
			fmt.Println(nimp_ids)
			// db.InsertNewDataEntry(cranium, "A8b839013jgkke", atags, ptags, etc.)
    	} else {
    		// The user already exists, so ask the database for attributes
    		title = craniumId.(string)
    	}
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "cranium.css | " + title,
        })
    })

	r.Run(":1225")
}