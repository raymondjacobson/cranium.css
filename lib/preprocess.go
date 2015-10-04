//----------------------------------------------------
// CSS Preprocessor for cranium.css
//  Converts owner-defined html "important" tags into
//  unique identifiers
//----------------------------------------------------

package preprocess

import (
  "os"
  "regexp"
  "strconv"
  "strings"
)

const prefix = "imp"
var tags = []string {"p", "a", "img"}

func readFile(file *os.File) string {
  bufsize := 1024

  data := make([]byte, bufsize)
  count, _ := file.Read(data)
  var fileread string

  for count > 0 {
    fileread += string(data[:count])

    data = make([]byte, bufsize)
    count, _ = file.Read(data)
  }

  return fileread
}

func GenerateTags(rf *os.File, of *os.File) ([]string, []string, []string, []string, []string, []string) {
  readfile := readFile(rf)

  re := regexp.MustCompile("<[a-zA-Z0-9= \"]*>")
  tags := re.FindAllIndex([]byte(readfile), -1)

  a_count := 0
  p_count := 0
  img_count := 0

  n_a_count := 0
  n_p_count := 0
  n_img_count := 0

  outstring := ""
  var a_imp_ids []string
  var p_imp_ids []string
  var img_imp_ids []string
  var a_nimp_ids []string
  var p_nimp_ids []string
  var img_nimp_ids []string

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

    var gen_id string

    if id == prefix {
      stag_count := ""
      if tag_type == "a" {
        stag_count = strconv.Itoa(a_count)
        a_count += 1
        gen_id = prefix+"-"+tag_type+"-"+stag_count
        a_imp_ids = append(a_imp_ids, gen_id)
      } else if tag_type == "p" {
        stag_count = strconv.Itoa(p_count)
        p_count += 1
        gen_id = prefix+"-"+tag_type+"-"+stag_count
        p_imp_ids = append(p_imp_ids, gen_id)
      } else if tag_type == "img" {
        stag_count = strconv.Itoa(img_count)
        img_count += 1
        gen_id = prefix+"-"+tag_type+"-"+stag_count
        img_imp_ids = append(img_imp_ids, gen_id)
      }

      outstring += pre_id+"id=\""+gen_id+"\""+post_id
    } else {
      nimp := true
      stag_count := ""
      if tag_type == "a" {
        stag_count = strconv.Itoa(n_a_count)
        n_a_count += 1
        gen_id = "n"+prefix+"-"+tag_type+"-"+stag_count
        a_nimp_ids = append(a_nimp_ids, gen_id)
      } else if tag_type == "p" {
        stag_count = strconv.Itoa(n_p_count)
        n_p_count += 1
        gen_id = "n"+prefix+"-"+tag_type+"-"+stag_count
        p_nimp_ids = append(p_nimp_ids, gen_id)
      } else if tag_type == "img" {
        stag_count = strconv.Itoa(n_img_count)
        n_img_count += 1
        gen_id = "n"+prefix+"-"+tag_type+"-"+stag_count
        img_nimp_ids = append(img_nimp_ids, gen_id)
      } else {
        nimp = false
      }

      outstring += full_tag

      if nimp {
        outstring += " id=\""+gen_id+"\""
      }
    }
    outstring += ">"
  }
  outstring += readfile[tags[len(tags)-1][1]:]
  of.WriteString(outstring)

  return a_imp_ids, p_imp_ids, img_imp_ids, a_nimp_ids, p_nimp_ids, img_nimp_ids
}