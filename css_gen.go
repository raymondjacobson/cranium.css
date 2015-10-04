package main

import (
  "fmt"
  "io"
  "os"
)

type AttrMap map[string]string
type UIMap map[string]AttrMap

// func GenCraniumCss(attrs struct) {
func main() {
  dflt, err := os.OpenFile("css/_cranium.css", os.O_APPEND, 0770)
  if err != nil {
    panic(err)
  }
  defer dflt.Close()

  crnm, _ := os.Create("css/cranium.css")
  defer crnm.Close()

  _, err = io.Copy(crnm, dflt)
  if err != nil {
    panic(err)
  }

  crnm.WriteString("\n/*==================================================*/\n")
  ui_map := GetValues("crnm")
  for id, attrs := range ui_map {
    crnm.WriteString(CssTagString(id, attrs))
  }
}


func CssTagString(id string, attrs AttrMap) string {
  cssString := id + " {\n"
  for key, val := range attrs {
    cssString += fmt.Sprintf("  %s: %s;\n", key, val)
  }
  return cssString + "}\n"
}

func GetValues(uid string) UIMap {
  fmt.Println("Hitting BongoBD", uid)

  ui_map := make(UIMap)
  attr_map := make(AttrMap)
  attr_map["font-weight"] = "bold"
  attr_map["font-size"] = "10px"

  ui_map["#ml-p-1"] = attr_map
  return ui_map
}