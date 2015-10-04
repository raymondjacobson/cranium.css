package model

import (
    "fmt"
)

type Tag struct {
    Id string
    Hover int
    Click int
    Frame int
}

type Ptag struct {
    Id string
    FontSize int
    FontStyle int
    Padding int
    Hover int
    Click int
    FrameTime int
    Important bool
}

type Atag struct {
    Id string
    FontSize int
    FontStyle int
    Color bool
    Padding int
    Hover int
    Click int
    FrameTime int
    Important bool
}

func getFontStyle(FontStyle int) string {
    switch {
        case FontStyle == 0: // Normal
            return "font-style: normal; font-weight: normal; text-decoration: none;"
        case FontStyle == 1: // Italic
            return "font-style: italic; font-weight: normal; text-decoration: none;"
        case FontStyle == 2: // Bold
            return "font-style: normal; font-weight: bold; text-decoration: none;"
        case FontStyle == 3: // Underline
            return "font-style: normal; font-weight: normal; text-decoration: underline;"
        case FontStyle == 4: // Italic Underline
            return "font-style: italic; font-weight: bold; text-decoration: underline;"
        case FontStyle == 5: // Italic Bold
            return "font-style: italic; font-weight: bold; text-decoration: none;"
        case FontStyle == 6: // Bold Underline
            return "font-style: normal; font-weight: bold; text-decoration: underline;"
    }
    return "font-style: normal; font-weight: normal; text-decoration: none;"
}

func (p Ptag) GetCssFontStyle() string {
  return getFontStyle(p.FontStyle)
}

func (p Ptag) GetCssFontSize() string {
  return fmt.Sprintf("font-size: %d;", p.FontSize)
}

func (p Ptag) GetCssPadding() string {
  return fmt.Sprintf("padding: %dpx 0 %dpx 0;", p.Padding, p.Padding)
}

func (a Atag) GetCssFontStyle() string {
  return getFontStyle(a.FontStyle)
}

func (a Atag) GetCssFontSize() string {
  return fmt.Sprintf("font-size: %d;", a.FontSize)
}

func (a Atag) GetCssPadding() string {
  return fmt.Sprintf("padding: %dpx 0 %dpx 0;", a.Padding, a.Padding)
}

func (a Atag) GetCssColor() string {
  if a.Color {
    return "color: #765DB6;"
  }
  return "color: #222;"
}

type Imgtag struct {
    Id string
    Width int
    Padding int
    Hover int
    Click int
    FrameTime int
    Important bool
}

func (i Imgtag) GetCssWidth() string {
  return fmt.Sprintf("width: %d%%;", i.Width)
}

func (i Imgtag) GetCssPadding() string {
  return fmt.Sprintf("padding: %dpx 0 %dpx 0;", i.Padding, i.Padding)
}

type DataEntry struct {
    Atags []Atag
    Ptags []Ptag
    Imgtags []Imgtag
}

type Visitor struct {
    Vid string
    Data []DataEntry
}

var PtagDefault = Ptag {
    FontSize: 16,
    FontStyle: 0,
    Padding: 10,
}

var AtagDefault = Atag {
    FontSize: 16,
    FontStyle: 3,
    Color: true,
    Padding: 10,
}

var ImgtagDefault = Imgtag {
    Width: 100,
    Padding: 10,
}
