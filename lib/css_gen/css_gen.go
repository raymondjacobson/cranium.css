package css_gen

import (
  "fmt"
  "../model"
)


func GenCss(attrs model.DataEntry) {
  css := ""
  for atag := range attrs.Atag {
    css += AtagCss(atag)
  }

  for ptag := range attrs.Ptag {
    css += PtagCss(ptag)
  }

  for imgtag := range attrs.Imgtag {
    css += ImgtagCss(imgtag)
  }
  return css
}


func AtagCss(atag model.Atag) string {
  css := fmt.Sprintf("%s{", atag.Id)
  css += fmt.Sprintf("font-size:%s;", atag.GetCssFontSize)
  css += fmt.Sprintf("font-style:%s;", atag.GetCssFontStyle)
  css += fmt.Sprintf("padding:%s;", atag.GetCssPadding)
  css += "}"
  return css
}


func PtagCss(ptag model.Ptag) string {
  css := fmt.Sprintf("%s{", ptag.Id)
  css += fmt.Sprintf("font-size:%s;", ptag.GetCssFontSize)
  css += fmt.Sprintf("font-style:%s;", ptag.GetCssFontStyle)
  css += fmt.Sprintf("padding:%s;", ptag.GetCssPadding)
  css += "}"
  return css
}


func ImgtagCss(imgtag model.Imgtag) string {
  css := fmt.Sprintf("%s{", imgtag.Id)
  css += fmt.Sprintf("width:%s;", imgtag.GetCssWidth)
  css += fmt.Sprintf("padding:%s;", imgtag.GetCssPadding)
  css += "}"
  return css
}
