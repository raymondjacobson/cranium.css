package model

type Ptag struct {
    Id string
    FontSize int
    FontWeight int
    FontStyle int
    TextDecoration int
    Padding int
    Hover int
    Click int
    FrameTime int
}

type Atag struct {
    Id string
    FontSize int
    FontWeight int
    FontStyle int
    TextDecoration int
    Padding int
    Hover int
    Click int
    FrameTime int
}

type Imgtag struct {
    Id string
    Area int
    Padding int
    Hover int
    Click int
    FrameTime int
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