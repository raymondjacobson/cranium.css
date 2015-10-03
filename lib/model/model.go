package model

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
    Padding int
    Hover int
    Click int
    FrameTime int
}

type Imgtag struct {
    Id string
    Width int
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