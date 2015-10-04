package model

import (
    "strconv"
)

type Ptag struct {
    Id string
    FontSize int
    FontStyle int
    Padding int
    Hover int
    Click int
    FrameTime int
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
}

func getFontStyle string {
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
}

func (p Ptag) GetCssFontStyle string {
    return getFontStyle(p.FontStyle)
}

func (p Ptag) GetCssFontSize string {
    return "font-size: " + strconv.Itoa(p.FontSize) + ";"
}

func (p Ptag) GetCssPadding string {
    return "padding: " + strconv.Itoa(p.Padding) + "px 0 " + strconv.Itoa(p.Padding) + "px 0;"
}

func (a Atag) GetCssFontStyle string {
    return getFontStyle(a.FontStyle)
}

func (a Atag) GetCssFontSize string {
    return "font-size: " + strconv.Itoa(a.FontSize) + ";"
}

func (a Atag) GetCssPadding string {
    return "padding: " + strconv.Itoa(a.Padding) + "px 0 " + strconv.Itoa(a.Padding) + "px 0;"
}

func (a Atag) GetCssFont string {
    if a.color {
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
}

func (i Imgtag) GetCssWidth string {
    return "width: " + strconv.Itoa(i.Width) + "px;"
}

func (i Imgtag) GetCssPadding string {
    return "padding: " + strconv.Itoa(i.Padding) + "px 0 " + strconv.Itoa(i.Padding) + "px 0;"
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

const PtagDefault = Ptag {
    FontSize: 16,
    FontStyle: 0,
    Padding: 10
}

const AtagDefault = Atag {
    FontSize: 16,
    FontStyle: 3,
    Color: 1,
    Padding: 10
}

const ImgtagDefault = Imgtag {
    Width: 100,
    Padding: 10
}