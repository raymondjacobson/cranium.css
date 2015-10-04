package cssgen

import (
  "fmt"
  "model"
)


func GenCss(attrs model.DataEntry) string {
  css := ""
  for _, atag := range attrs.Atags {
    css += AtagCss(atag)
  }

  for _, ptag := range attrs.Ptags {
    css += PtagCss(ptag)
  }

  for _, imgtag := range attrs.Imgtags {
    css += ImgtagCss(imgtag)
  }
  return css
}


func AtagCss(atag model.Atag) string {
  css := fmt.Sprintf("%s{\n", atag.Id)
  css += fmt.Sprintf("%s\n", atag.GetCssFontSize())
  css += fmt.Sprintf("%s\n", atag.GetCssFontStyle())
  css += fmt.Sprintf("%s\n", atag.GetCssPadding())
  css += "}\n"
  return css
}


func PtagCss(ptag model.Ptag) string {
  css := fmt.Sprintf("%s{\n", ptag.Id)
  css += fmt.Sprintf("%s\n", ptag.GetCssFontSize())
  css += fmt.Sprintf("%s\n", ptag.GetCssFontStyle())
  css += fmt.Sprintf("%s\n", ptag.GetCssPadding())
  css += "}\n"
  return css
}


func ImgtagCss(imgtag model.Imgtag) string {
  css := fmt.Sprintf("%s{\n", imgtag.Id)
  css += fmt.Sprintf("%s\n", imgtag.GetCssWidth())
  css += fmt.Sprintf("%s\n", imgtag.GetCssPadding())
  css += "}\n"
  return css
}
