// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoNeko

//---------------------------------------------------------

type NekoBase struct {
	Url  string `json:"url"`
	Why  string `json:"why"`
	Name string `json:"name"`
	Cat  string `json:"cat"`
	Fact string `json:"fact"`
	Owo  string `json:"owo"`
}

//---------------------------------------------------------

// TenorList is simply a list of tenor gif base.
type TenorList []TenorBase

// TenorBase is an structure for
// receiving and decoding the gif data from
// tenor.come written by ALiwoto.
type TenorBase struct {
	Url       string `json:"url"`  // [preview]
	Src       string `json:"src"`  // [real source]
	MediaType string `json:"type"` // [the type]
}

//---------------------------------------------------------

type tenorType string

type TenorContainer struct {
	Results []ResultsTenor `json:"results"`
	Next    string         `json:"next"`
}

type Mp4Tenor struct {
	Size     int     `json:"size"`
	Dims     []int   `json:"dims"`
	Preview  string  `json:"preview"`
	Duration float64 `json:"duration"`
	URL      string  `json:"url"`
}

type MediumgifTenor struct {
	Dims    []int  `json:"dims"`
	URL     string `json:"url"`
	Size    int    `json:"size"`
	Preview string `json:"preview"`
}

type NanogifTenor struct {
	URL     string `json:"url"`
	Dims    []int  `json:"dims"`
	Size    int    `json:"size"`
	Preview string `json:"preview"`
}

type TinywebmTenor struct {
	Dims    []int  `json:"dims"`
	Preview string `json:"preview"`
	URL     string `json:"url"`
	Size    int    `json:"size"`
}

type GifTenor struct {
	Size    int    `json:"size"`
	Dims    []int  `json:"dims"`
	Preview string `json:"preview"`
	URL     string `json:"url"`
}

type Tinymp4Tenor struct {
	Duration float64 `json:"duration"`
	Size     int     `json:"size"`
	URL      string  `json:"url"`
	Preview  string  `json:"preview"`
	Dims     []int   `json:"dims"`
}

type NanowebmTenor struct {
	Size    int    `json:"size"`
	URL     string `json:"url"`
	Dims    []int  `json:"dims"`
	Preview string `json:"preview"`
}

type Nanomp4Tenor struct {
	URL      string  `json:"url"`
	Size     int     `json:"size"`
	Dims     []int   `json:"dims"`
	Duration float64 `json:"duration"`
	Preview  string  `json:"preview"`
}

type Loopedmp4Tenor struct {
	Size     int     `json:"size"`
	Duration float64 `json:"duration"`
	Preview  string  `json:"preview"`
	URL      string  `json:"url"`
	Dims     []int   `json:"dims"`
}

type WebmTenor struct {
	Dims    []int  `json:"dims"`
	URL     string `json:"url"`
	Size    int    `json:"size"`
	Preview string `json:"preview"`
}

type TinygifTenor struct {
	Dims    []int  `json:"dims"`
	Preview string `json:"preview"`
	URL     string `json:"url"`
	Size    int    `json:"size"`
}

type MediaTenor struct {
	Mp4       *Mp4Tenor       `json:"mp4"`
	Mediumgif *MediumgifTenor `json:"mediumgif"`
	Nanogif   *NanogifTenor   `json:"nanogif"`
	Tinywebm  *TinywebmTenor  `json:"tinywebm"`
	Gif       *GifTenor       `json:"gif"`
	Tinymp4   *Tinymp4Tenor   `json:"tinymp4"`
	Nanowebm  *NanowebmTenor  `json:"nanowebm"`
	Nanomp4   *Nanomp4Tenor   `json:"nanomp4"`
	Loopedmp4 *Loopedmp4Tenor `json:"loopedmp4"`
	Webm      *WebmTenor      `json:"webm"`
	Tinygif   *TinygifTenor   `json:"tinygif"`
}

// the results for tenor search
// I removed these two as they are not really
// necessary:
//	Tags       []interface{} `json:"tags"`
//	Flags      []interface{} `json:"flags"`
type ResultsTenor struct {
	ID         string       `json:"id"`
	Title      string       `json:"title"`
	H1Title    string       `json:"h1_title"`
	Media      []MediaTenor `json:"media"`
	BgColor    string       `json:"bg_color"`
	Created    float64      `json:"created"`
	Itemurl    string       `json:"itemurl"`
	URL        string       `json:"url"`
	Shares     int          `json:"shares"`
	Hasaudio   bool         `json:"hasaudio"`
	Hascaption bool         `json:"hascaption"`
	SourceID   string       `json:"source_id"`
	Composite  interface{}  `json:"composite"`
}

//---------------------------------------------------------
