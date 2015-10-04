package main

import (
    "fmt"
    "html/template"
    "github.com/gin-gonic/contrib/sessions"
    "github.com/gin-gonic/gin"
    "net/http"
    "math/rand"
    "os"
    "os/exec"
    "strings"
    "strconv"
    "time"
    "./lib/cssgen"
    "./lib/db"
    "./lib/model"
    "./lib/preprocess"
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

func deserialize(input string) []model.Tag {
  var tags []model.Tag
 
  tag_names := strings.Split(input, ";")
 
  for i := 0; i < len(tag_names); i++ {
    split_tag := strings.Split(tag_names[i], ":")
    id := split_tag[0]
 
    nums := strings.Split(split_tag[1], ",")
 
    hover, _ := strconv.Atoi(nums[0])
    click, _ := strconv.Atoi(nums[1])
    frame, _ := strconv.Atoi(nums[2])
 
    tags = append(tags, model.Tag{id, hover, click, frame})
  }
 
  return tags
}

func main() {
    s := db.Connect()
    defer db.Disconnect(s)
    craniumDB := db.SetDB(s)

    r := gin.Default()

    // Set up Gin cookies
    store := sessions.NewCookieStore([]byte("pladjamaicanwifflepoof"))
    r.Use(sessions.Sessions("cranium_session", store))

    var title string

    var de model.DataEntry

    r.LoadHTMLGlob("templates/*")
    r.Static("/assets", "./assets")

    // Routes
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
            de = db.InsertNewDataEntry(craniumDB, newCraniumId, atags, ptags, imgtags)
        } else {
            readfile, _ := os.Open("templates/index-tmpl.html")
            outfile, _ := os.Create("templates/index.html")
            defer readfile.Close()
            defer outfile.Close()
            // Get all the Ids for tags in question
            preprocess.GenerateTags(readfile, outfile)
            // The user already exists, so ask the database for attributes
            title = craniumId.(string)
            visitor := db.FetchVisitor(craniumDB, craniumId.(string))
            de = visitor.Data[len(visitor.Data)-1]
        }
        craniumcss := cssgen.GenCss(de)
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "cranium.css | " + title,
            "craniumcss": template.CSS(craniumcss),
        })
    })

    type Entities struct {
        Data string `json:"data" binding:"required"`
    }

    r.POST("/data", func(c *gin.Context) {
        session := sessions.Default(c)
        craniumId := session.Get("cranium_id")
        var ents Entities
        c.BindJSON(&ents)
        c.JSON(200, ents)
        tags := deserialize(ents.Data)
        for _, tag := range tags {
            tag_pick := strings.Split(string(tag.Id), "-")
            if string(tag_pick[1][0]) == "a" {
                db.UpdateAtagField(craniumDB, craniumId.(string), tag)
            } else if string(tag_pick[1][0]) == "p" {
                db.UpdatePtagField(craniumDB, craniumId.(string), tag)
            } else if string(tag_pick[1][0]) == "i" {
                db.UpdateImgtagField(craniumDB, craniumId.(string), tag)
            }
        }
        // CSV the data
        app := "python"
        cmd := exec.Command(app, "learn/csvgen.py", craniumId.(string), "-de0")
        fmt.Println()
        out, err := cmd.Output()
        if err != nil {
            println(err.Error())
            return
        }
        fmt.Println(string(out))

        fmt.Println(craniumId.(string))

        cmd1 := exec.Command(app, "learn/classify.py", "learn/de_td0_a.csv", "a", craniumId.(string), "", "", "")
        cmd1.Output()

        cmd2 := exec.Command(app, "learn/classify.py", "learn/de_td0_p.csv", "p", craniumId.(string), "", "", "")
        cmd2.Output()

        cmd3 := exec.Command(app, "learn/classify.py", "learn/de_td0_img.csv", "img", craniumId.(string), "", "", "")
        cmd3.Output()

        cmd4 := exec.Command(app, "learn/classify.py", "", "solve", craniumId.(string), "learn/training/atag.csv", "learn/training/ptag.csv", "learn/training/imgtag.csv")
        cmd4.Output()

        // Run genetic

        // Save new DE entry

        // Export DE to CSV
        // cmd4 := exec.Command(app, "learn/csvgen.py", craniumId.(string))
        // cmd4.Output()

        // Concat CSV with train

        // run SVM
        // python learn/svm.py learn/training/imgtag.csv img
        // cmd5 := exec.Command(app, "learn/svm.py", "learn/training/atag.csv", "a")
        // cmd5.Output()

        // cmd6 := exec.Command(app, "learn/svm.py", "learn/training/ptag.csv", "p")
        // cmd5.Output()

        // cmd7 := exec.Command(app, "learn/svm.py", "learn/training/imgtag.csv", "img")
        // cmd5.Output()
    })

    r.Run(":1225")
}