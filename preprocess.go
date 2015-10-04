//----------------------------------------------------
// CSS Preprocessor for cranium.css
//  Converts owner-defined html "important" tags into
//  unique identifiers
//----------------------------------------------------

package main

import (
  "os"
  "fmt"
  "strings"
  "regexp"
  "strconv"
)

const prefix = "ml"
var tags = []string {"p", "a", "img"}

func readFile(filename string) string {
  bufsize := 1024
  file, _ := os.Open(filename)

  data := make([]byte, bufsize)
  count, _ := file.Read(data)
  var fileread string

  for count > 0 {
    fileread += string(data[:count])

    data = make([]byte, bufsize)
    count, _ = file.Read(data)
  }

  defer file.Close()

  return fileread
}

func generateTags(infile string, outfile string) {
  readfile := readFile(infile)
  out_f, _ := os.Create(outfile)

  re := regexp.MustCompile("<[a-zA-Z0-9= \"]*>")
  tags := re.FindAllIndex([]byte(readfile), -1)

  a_counter := 0
  p_counter := 0
  img_counter := 0

  outstring := ""
  for i := 0; i < len(tags); i++ {
    prev_end := 0
    if i == 0 {
      prev_end = 0
    } else {
      prev_end = tags[i-1][1]
    }

    start_i := tags[i][0]+1
    end_i := tags[i][1]-1

    outstring += readfile[prev_end:start_i]

    full_tag := readfile[start_i:end_i]
    split := strings.Split(full_tag, " ")

    tag_type := split[0]
    id := ""

    pre_id := ""
    post_id := ""
    pre := true
    for j := 0; j < len(split); j++ {
      piece := split[j]
      if len(piece) > 6 && piece[:6] == "class=" {
        id = piece[7:len(piece)-1]
        pre = false
      } else {
        if pre {
          pre_id += piece + " "
        } else {
          post_id += " " + piece
        }
      }
    }

    pre_id = pre_id[:len(pre_id)]
    post_id = post_id[:len(post_id)]

    if id == "ml" {
      stag_counter := ""
      if tag_type == "a" {
        stag_counter = strconv.Itoa(a_counter)
        a_counter += 1
      } else if tag_type == "p" {
        stag_counter = strconv.Itoa(p_counter)
        p_counter += 1
      } else if tag_type == "img" {
        stag_counter = strconv.Itoa(img_counter)
        img_counter += 1
      }

      outstring += pre_id+"id=\"ml-"+tag_type+"-"+stag_counter+"\""+post_id
    } else {
      outstring += full_tag
    }
    outstring += ">"
  }
  outstring += readfile[tags[len(tags)-1][1]:]
  out_f.WriteString(outstring)

  defer out_f.Close()
}

func main() {
  fmt.Println("")

  var infile = "dummy.html"
  var outfile = "replaced.html"
  generateTags(infile, outfile)
}