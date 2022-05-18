// tinyfontgen -yadvance 0 --package meteoicon --fontname MeteoIcon48pt meteo-icon-regular-48pt.bdf --output MeteoIcon-Regular-48pt.go --all --verbose

package meteoicon

import (
	"tinygo.org/x/tinyfont"
)

var MeteoIcon48pt = tinyfont.Font{
	BBox: [4]int8{50, 50, 0, -47},
	Glyphs: []tinyfont.Glyph{
		/* # */ tinyfont.Glyph{Rune: 0x23, Width: 0x32, Height: 0x32, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0x3, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xff, 0x80, 0x0, 0x0, 0x0, 0x7, 0xff, 0xf8, 0x0, 0x0, 0x0, 0x7, 0xff, 0xff, 0x80, 0x0, 0x0, 0x3, 0xff, 0xff, 0xf0, 0x0, 0x0, 0x1, 0xff, 0xff, 0xfe, 0x0, 0x0, 0x3, 0xff, 0xff, 0xff, 0xf8, 0x0, 0xf, 0xff, 0xff, 0xff, 0xff, 0xc0, 0xf, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x7, 0xff, 0xff, 0xff, 0xff, 0xff, 0x83, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x1f, 0xff, 0xff, 0xff, 0xff, 0xfe, 0x3, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x7f, 0xff, 0xff, 0xff, 0xff, 0x80, 0x7, 0xff, 0xff, 0xff, 0xff, 0x80, 0x0, 0x0, 0xff, 0xff, 0xc0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xc0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xf4, 0x0, 0x30, 0x0, 0x0, 0x0, 0x7f, 0x80, 0x1e, 0x0, 0xc, 0x0, 0x39, 0xc0, 0x7, 0x80, 0x3, 0x30, 0x1e, 0xe0, 0x31, 0xe3, 0x0, 0xfc, 0x7, 0xf8, 0x1e, 0xfd, 0xe1, 0xfe, 0x0, 0x3f, 0x7, 0xff, 0xf8, 0x39, 0x80, 0xc, 0x80, 0xff, 0xfc, 0x6, 0x78, 0x2, 0x0, 0xf, 0x3c, 0x1, 0xfe, 0x0, 0x0, 0x3, 0x87, 0x0, 0xfc, 0x0, 0x0, 0x0, 0xf3, 0xc0, 0x13, 0x0, 0x0, 0x0, 0xff, 0xfc, 0x0, 0xc0, 0x0, 0x0, 0x7f, 0xff, 0x80, 0x0, 0x0, 0x0, 0x1f, 0xff, 0xe0, 0x0, 0x0, 0x0, 0x3, 0x1e, 0x30, 0x0, 0x0, 0x0, 0x0, 0x7, 0x80, 0x0, 0x0, 0x0, 0x0, 0x1, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x30, 0x0, 0x0, 0x0}},
		/* % */ tinyfont.Glyph{Rune: 0x25, Width: 0x32, Height: 0x28, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0x0, 0x7, 0x80, 0x0, 0x0, 0x0, 0x0, 0xf, 0xfe, 0x0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0x0, 0x7, 0xff, 0xfc, 0x0, 0x0, 0x0, 0x3, 0xff, 0xff, 0x0, 0x0, 0x0, 0xf, 0xff, 0xff, 0xfe, 0x0, 0x0, 0xf, 0xff, 0xff, 0xff, 0xe0, 0x0, 0x7, 0xff, 0xff, 0xff, 0xfc, 0x0, 0x3, 0xff, 0xff, 0xff, 0xff, 0x80, 0x1, 0xff, 0xff, 0xff, 0xff, 0xe0, 0x0, 0x7f, 0xff, 0xff, 0xff, 0xfc, 0x0, 0x1f, 0xff, 0xff, 0xff, 0xff, 0x0, 0x7, 0xff, 0xff, 0xff, 0xff, 0xc0, 0x1, 0xff, 0xff, 0xff, 0xff, 0xf0, 0x0, 0x7f, 0xff, 0xff, 0xff, 0xfc, 0x3, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xc3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe, 0x3f, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7, 0xff, 0xff, 0xff, 0xff, 0xff, 0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xc0, 0xf, 0xff, 0xff, 0xff, 0xff, 0xc0, 0x0, 0xfd, 0xff, 0xff, 0xef, 0xc0, 0x0, 0x0, 0x1f, 0xff, 0xf0, 0x0, 0x0, 0x0, 0x3, 0xff, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0xe0, 0x0, 0x0}},
		/* ) */ tinyfont.Glyph{Rune: 0x29, Width: 0x28, Height: 0x13, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0x7, 0x0, 0x0, 0xf0, 0x1c, 0x7, 0x3, 0xc0, 0xf0, 0x1c, 0x6, 0x3, 0xe0, 0xf8, 0x1c, 0xe, 0x7, 0xe0, 0xfc, 0x1c, 0xe, 0x7, 0xe0, 0xfc, 0x1c, 0xc, 0x7, 0xf0, 0xfe, 0x1c, 0x1c, 0xf, 0x70, 0xff, 0x1c, 0x1c, 0xe, 0x70, 0xf7, 0x9c, 0x18, 0xe, 0x78, 0xf7, 0x9c, 0x38, 0x1e, 0x38, 0xf3, 0xdc, 0x38, 0x1c, 0x38, 0xf1, 0xfc, 0x30, 0x1f, 0xfc, 0xf0, 0xfc, 0x70, 0x3f, 0xfc, 0xf0, 0xfc, 0x70, 0x3f, 0xfc, 0xf0, 0x7c, 0x60, 0x38, 0x1e, 0xf0, 0x3c, 0xe0, 0x78, 0xe, 0xf0, 0x3c, 0xe0, 0x70, 0xe, 0x0, 0x0, 0xc0, 0x70, 0xf, 0x0, 0x1, 0xc0, 0x80, 0x0}},
		/* 4 */ tinyfont.Glyph{Rune: 0x34, Width: 0x32, Height: 0x2e, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0x0, 0x3, 0x80, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0xfc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3d, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xe, 0x70, 0x0, 0x0, 0x0, 0x0, 0x7, 0x9e, 0x0, 0x0, 0x0, 0x0, 0x1, 0xc3, 0x80, 0x0, 0x0, 0x0, 0x0, 0x70, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1c, 0x3c, 0x0, 0x0, 0x0, 0x0, 0x7, 0x7, 0x80, 0x0, 0x0, 0x0, 0x1, 0xc0, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x70, 0x3e, 0x0, 0x0, 0x0, 0x0, 0x1c, 0x7, 0xfb, 0xe0, 0x0, 0x0, 0x1f, 0xe0, 0x7f, 0xf8, 0x0, 0x0, 0x3f, 0xff, 0xf, 0xfe, 0x0, 0x0, 0x3f, 0xff, 0xf0, 0x3f, 0x0, 0x0, 0x1f, 0xff, 0xfe, 0x3, 0xc0, 0x0, 0xf, 0xff, 0xff, 0xc3, 0xe0, 0x0, 0x7, 0xff, 0xff, 0xfb, 0xf0, 0x0, 0x3f, 0xff, 0xff, 0xff, 0xf8, 0x0, 0x3f, 0xff, 0xff, 0xff, 0xff, 0x0, 0x3f, 0xff, 0xff, 0xff, 0xff, 0xf0, 0x1f, 0xff, 0xff, 0xff, 0xff, 0xfe, 0xf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xc7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xbf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf0, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xf8, 0xf, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x1, 0xff, 0xff, 0xff, 0xff, 0xfe, 0x0, 0xf, 0xdf, 0xff, 0xfe, 0xfc, 0x0, 0x0, 0x3, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x3f, 0xff, 0x0, 0x0, 0x0, 0x0, 0x3, 0xff, 0x0, 0x0, 0x0}},
		/* 5 */ tinyfont.Glyph{Rune: 0x35, Width: 0x32, Height: 0x20, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x7, 0xff, 0x80, 0x0, 0x0, 0x0, 0x7, 0xff, 0xf8, 0x0, 0x0, 0x0, 0x3, 0xff, 0xff, 0x0, 0x0, 0x0, 0x1, 0xff, 0xff, 0xe0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xfc, 0x0, 0x0, 0x0, 0x7f, 0xff, 0xff, 0x80, 0x0, 0x7, 0xff, 0xff, 0xff, 0xff, 0x80, 0x7, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x3, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x3f, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7, 0xff, 0xff, 0xff, 0xff, 0xff, 0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0x80, 0xf, 0xff, 0xff, 0xff, 0xff, 0xc0, 0x0, 0x70, 0xff, 0xff, 0xc3, 0x0, 0x0, 0x0, 0x1f, 0xff, 0xe0, 0x0, 0x0, 0x0, 0x1, 0xff, 0xe0, 0x0, 0x0, 0x0, 0x0, 0xf, 0xc0, 0x0, 0x0}},
		/* 6 */ tinyfont.Glyph{Rune: 0x36, Width: 0x32, Height: 0x2f, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0x3, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xff, 0x80, 0x0, 0x0, 0x0, 0x7, 0xff, 0xf8, 0x0, 0x0, 0x0, 0x7, 0xff, 0xff, 0x80, 0x0, 0x0, 0x3, 0xff, 0xff, 0xf0, 0x0, 0x0, 0x1, 0xff, 0xff, 0xfe, 0x0, 0x0, 0x3, 0xff, 0xff, 0xff, 0xf8, 0x0, 0xf, 0xff, 0xff, 0xff, 0xff, 0xc0, 0xf, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x7, 0xff, 0xff, 0xff, 0xff, 0xff, 0x83, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x1f, 0xff, 0xff, 0xff, 0xff, 0xfe, 0x3, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x7f, 0xff, 0xff, 0xff, 0xff, 0x80, 0x7, 0xff, 0xff, 0xff, 0xff, 0x80, 0x0, 0x0, 0xff, 0xff, 0xc0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xf, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xfc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x80, 0x0, 0x0, 0x0, 0x0, 0x1, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0}},
		/* 7 */ tinyfont.Glyph{Rune: 0x37, Width: 0x32, Height: 0x2c, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0x1, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xff, 0x80, 0x0, 0x0, 0x0, 0x7, 0xff, 0xf8, 0x0, 0x0, 0x0, 0x7, 0xff, 0xff, 0x80, 0x0, 0x0, 0x3, 0xff, 0xff, 0xf0, 0x0, 0x0, 0x1, 0xff, 0xff, 0xfe, 0x0, 0x0, 0x2, 0x7f, 0xff, 0xff, 0xa0, 0x0, 0xf, 0xff, 0xff, 0xff, 0xff, 0xc0, 0x7, 0xff, 0xff, 0xff, 0xff, 0xf8, 0x7, 0xff, 0xff, 0xff, 0xff, 0xff, 0x83, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x1f, 0xff, 0xff, 0xff, 0xff, 0xfe, 0x3, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x7f, 0xff, 0xff, 0xff, 0xff, 0x80, 0x7, 0xff, 0xff, 0xff, 0xff, 0x80, 0x0, 0x0, 0xff, 0xff, 0xc0, 0x0, 0x0, 0x0, 0x1f, 0xff, 0xe0, 0x0, 0x0, 0x0, 0x1, 0xff, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x78, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1e, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xfc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x80, 0x0, 0x0}},
		/* 8 */ tinyfont.Glyph{Rune: 0x38, Width: 0x32, Height: 0x32, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0x3, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xff, 0x80, 0x0, 0x0, 0x0, 0x7, 0xff, 0xf8, 0x0, 0x0, 0x0, 0x7, 0xff, 0xff, 0x80, 0x0, 0x0, 0x3, 0xff, 0xff, 0xf0, 0x0, 0x0, 0x1, 0xff, 0xff, 0xfe, 0x0, 0x0, 0x3, 0xff, 0xff, 0xff, 0xf8, 0x0, 0xf, 0xff, 0xff, 0xff, 0xff, 0xc0, 0xf, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x7, 0xff, 0xff, 0xff, 0xff, 0xff, 0x83, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0x1f, 0xff, 0xff, 0xff, 0xff, 0xfe, 0x3, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x7f, 0xff, 0xff, 0xff, 0xff, 0x80, 0x7, 0xff, 0xff, 0xff, 0xff, 0x80, 0x0, 0x0, 0xff, 0xff, 0xc0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xc0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x0, 0x0, 0x40, 0x0, 0x0, 0xf8, 0x0, 0x0, 0x38, 0x0, 0x0, 0x1c, 0x0, 0x0, 0xe, 0x0, 0x0, 0x0, 0x0, 0x10, 0x7, 0xc0, 0x0, 0x0, 0x0, 0xc, 0x1, 0xf0, 0x0, 0x0, 0x0, 0x7, 0x80, 0xfe, 0x0, 0x0, 0x0, 0x1, 0xe0, 0x3f, 0x80, 0x0, 0x0, 0x0, 0xfc, 0x7, 0xc0, 0x0, 0x0, 0x0, 0x3f, 0x0, 0xe0, 0x0, 0x0, 0x0, 0xf, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x78, 0x0, 0x0, 0x0}},
		/* B */ tinyfont.Glyph{Rune: 0x42, Width: 0x28, Height: 0x28, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0x18, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x0, 0x0, 0x3, 0x0, 0x18, 0x0, 0xc0, 0x7, 0x80, 0x0, 0x1, 0xe0, 0x7, 0xc0, 0x0, 0x3, 0xe0, 0x3, 0xc0, 0x0, 0x3, 0xc0, 0x1, 0xe0, 0x7e, 0x7, 0x80, 0x0, 0xc1, 0xff, 0x83, 0x0, 0x0, 0x7, 0xff, 0xe0, 0x0, 0x0, 0xf, 0xc3, 0xf0, 0x0, 0x0, 0x1f, 0x0, 0xf8, 0x0, 0x0, 0x1e, 0x0, 0x78, 0x0, 0x0, 0x3c, 0x0, 0x3c, 0x0, 0x0, 0x38, 0x0, 0x1c, 0x0, 0x0, 0x78, 0x0, 0x1e, 0x0, 0x0, 0x70, 0x0, 0xe, 0x0, 0xfc, 0x70, 0x0, 0xe, 0x3f, 0xfc, 0x70, 0x0, 0xe, 0x3f, 0xfc, 0x70, 0x0, 0xe, 0x3f, 0x0, 0x78, 0x0, 0x1e, 0x0, 0x0, 0x38, 0x0, 0x1c, 0x0, 0x0, 0x3c, 0x0, 0x3c, 0x0, 0x0, 0x3c, 0x0, 0x3c, 0x0, 0x0, 0x1e, 0x0, 0xf8, 0x0, 0x0, 0xf, 0x81, 0xf0, 0x0, 0x0, 0x7, 0xff, 0xe0, 0x0, 0x0, 0x83, 0xff, 0xc0, 0x0, 0x1, 0xc0, 0xfe, 0x3, 0x80, 0x3, 0xe0, 0x0, 0x7, 0xc0, 0x7, 0xc0, 0x0, 0x3, 0xe0, 0x7, 0x80, 0x0, 0x1, 0xe0, 0x3, 0x0, 0x18, 0x0, 0xc0, 0x0, 0x0, 0x3c, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x0, 0x0, 0x0, 0x0, 0x18, 0x0, 0x0}},
		/* C */ tinyfont.Glyph{Rune: 0x43, Width: 0x18, Height: 0x18, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x1, 0xc0, 0x0, 0x7, 0xc0, 0x0, 0xf, 0xc0, 0x0, 0x1f, 0xc0, 0x0, 0x3f, 0xc0, 0x0, 0x7b, 0xc0, 0x0, 0x79, 0xc0, 0x0, 0x71, 0xc0, 0x0, 0xe1, 0xc0, 0x0, 0xe1, 0xe0, 0x0, 0xe0, 0xe0, 0x0, 0xe0, 0xf0, 0x0, 0xe0, 0x78, 0x0, 0xe0, 0x3e, 0x0, 0xe0, 0x1f, 0xdf, 0xf0, 0xf, 0xff, 0x70, 0x7, 0xff, 0x78, 0x0, 0x7e, 0x3c, 0x0, 0x1e, 0x3e, 0x0, 0x3c, 0x1f, 0x80, 0xf8, 0xf, 0xff, 0xf0, 0x3, 0xff, 0xe0, 0x0, 0xff, 0x80}},
		/* H */ tinyfont.Glyph{Rune: 0x48, Width: 0x32, Height: 0x32, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0x18, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x70, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x0, 0x0, 0x0, 0x0, 0x38, 0x0, 0x0, 0xe, 0x0, 0x0, 0x1f, 0x0, 0x0, 0x7, 0x80, 0x0, 0x3, 0xe0, 0x0, 0x3, 0xc0, 0x0, 0x0, 0x78, 0x0, 0x0, 0xe0, 0x0, 0x0, 0xc, 0x7, 0xf0, 0x30, 0x0, 0x0, 0x0, 0x7, 0xff, 0x0, 0x0, 0x0, 0x0, 0x7, 0xff, 0xe0, 0x0, 0x0, 0x0, 0x3, 0xf0, 0x7e, 0x0, 0x0, 0x0, 0x1, 0xf0, 0x7, 0x80, 0x0, 0x0, 0x0, 0x70, 0x0, 0xf0, 0x0, 0x0, 0x0, 0x3c, 0x0, 0x1c, 0x0, 0x0, 0x0, 0xe, 0x0, 0x3, 0x80, 0x0, 0x0, 0x3, 0x80, 0x0, 0xe0, 0x0, 0x3, 0xf1, 0xe0, 0x1f, 0xf8, 0x7c, 0x0, 0xfc, 0x78, 0x3f, 0xff, 0x3f, 0x80, 0x3f, 0x1e, 0x3f, 0xff, 0xf7, 0xc0, 0x0, 0x3, 0x9f, 0x80, 0xfe, 0x0, 0x0, 0x0, 0xef, 0x80, 0x7, 0xc0, 0x0, 0x0, 0x3f, 0x80, 0x0, 0x78, 0x0, 0x0, 0x3f, 0xc0, 0x0, 0xf, 0xf0, 0x0, 0x7f, 0xe0, 0x0, 0x1, 0xff, 0x80, 0x3f, 0xf8, 0x0, 0x0, 0x7f, 0xf0, 0x1f, 0x4, 0x0, 0x0, 0x10, 0x3e, 0xf, 0x0, 0x0, 0x0, 0x0, 0x3, 0xc7, 0x80, 0x0, 0x0, 0x0, 0x0, 0x79, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xe, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0xb8, 0x0, 0x0, 0x0, 0x0, 0x0, 0xfe, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x80, 0x0, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3d, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xe, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0x9e, 0x0, 0x0, 0x0, 0x0, 0x1, 0xe3, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xf0, 0x7c, 0xc, 0x0, 0x0, 0xc0, 0xf8, 0xf, 0xff, 0x80, 0x0, 0x7f, 0xfc, 0x1, 0xff, 0xf8, 0x0, 0x7f, 0xfe, 0x0, 0xf, 0xdf, 0x80, 0x7e, 0xfc, 0x0, 0x0, 0x3, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x3f, 0xff, 0x0, 0x0, 0x0, 0x0, 0x3, 0xff, 0x0, 0x0, 0x0}},
		/* M */ tinyfont.Glyph{Rune: 0x4d, Width: 0x20, Height: 0x15, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x7f, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0xff, 0x7f, 0xff, 0xff, 0xfe, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7f, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0xff, 0x7f, 0xff, 0xff, 0xfe, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7f, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0xff, 0x7f, 0xff, 0xff, 0xfe, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7f, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0xff, 0x7f, 0xff, 0xff, 0xfe}},
		/* N */ tinyfont.Glyph{Rune: 0x4e, Width: 0x32, Height: 0x1f, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0xf, 0xfc, 0x0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xfc, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x1f, 0x80, 0x0, 0x0, 0x3, 0xe0, 0x1, 0xf0, 0x0, 0x0, 0x1, 0xe0, 0x0, 0x1e, 0x0, 0x0, 0xf, 0xf0, 0x0, 0x3, 0xfc, 0x0, 0x1f, 0xf8, 0x0, 0x0, 0x7f, 0xe0, 0xf, 0xfe, 0x0, 0x0, 0x1f, 0xfc, 0x7, 0xc1, 0x0, 0x0, 0x4, 0xf, 0x83, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xf1, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1e, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0x9c, 0x0, 0x0, 0x0, 0x0, 0x0, 0xee, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0x80, 0x0, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x80, 0x0, 0x0, 0x0, 0x0, 0xf, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0x9c, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe7, 0x80, 0x0, 0x0, 0x0, 0x0, 0x78, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x1f, 0x3, 0x0, 0x0, 0x30, 0x3e, 0x3, 0xff, 0xe0, 0x0, 0x1f, 0xff, 0x0, 0x7f, 0xfe, 0x0, 0x1f, 0xff, 0x80, 0x3, 0xf7, 0xe0, 0x1f, 0xbf, 0x0, 0x0, 0x0, 0xff, 0xff, 0xc0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xc0, 0x0, 0x0}},
		/* P */ tinyfont.Glyph{Rune: 0x50, Width: 0x32, Height: 0x2f, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0xf, 0xfc, 0x0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xfc, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x1f, 0x80, 0x0, 0x0, 0x3, 0xe0, 0x1, 0xf0, 0x0, 0x0, 0x1, 0xe0, 0x0, 0x1e, 0x0, 0x0, 0xf, 0xf0, 0x0, 0x3, 0xfc, 0x0, 0x1f, 0xf8, 0x0, 0x0, 0x7f, 0xe0, 0xf, 0xfe, 0x0, 0x0, 0x1f, 0xfc, 0x7, 0xc1, 0x0, 0x0, 0x4, 0xf, 0x83, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xf1, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1e, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0x9c, 0x0, 0x0, 0x0, 0x0, 0x0, 0xee, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0x80, 0x0, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x80, 0x0, 0x0, 0x0, 0x0, 0xf, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0x9c, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe7, 0x80, 0x0, 0x0, 0x0, 0x0, 0x78, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x1f, 0x3, 0x0, 0x0, 0x30, 0x3e, 0x3, 0xff, 0xf0, 0x4, 0x3f, 0xff, 0x0, 0x7f, 0xfe, 0x1, 0x1f, 0xff, 0x80, 0x3, 0xf7, 0xe0, 0x87, 0xbf, 0x0, 0x0, 0x0, 0xf0, 0x63, 0xc0, 0x0, 0x0, 0x0, 0x8, 0x38, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x1c, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xfc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x80, 0x0, 0x0, 0x0, 0x0, 0x1, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0}},
		/* Q */ tinyfont.Glyph{Rune: 0x51, Width: 0x32, Height: 0x2b, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0xf, 0xfc, 0x0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xfc, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x1f, 0x80, 0x0, 0x0, 0x3, 0xe0, 0x1, 0xf0, 0x0, 0x0, 0x1, 0xe0, 0x0, 0x1e, 0x0, 0x0, 0xf, 0xf0, 0x0, 0x3, 0xfc, 0x0, 0x1f, 0xf8, 0x0, 0x0, 0x7f, 0xe0, 0xf, 0xfe, 0x0, 0x0, 0x1f, 0xfc, 0x7, 0xc1, 0x0, 0x0, 0x4, 0xf, 0x83, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xf1, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1e, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0xbc, 0x0, 0x0, 0x0, 0x0, 0x0, 0xfe, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x80, 0x0, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xf, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0x9e, 0x0, 0x0, 0x0, 0x0, 0x1, 0xe3, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xf0, 0x7c, 0xc, 0x0, 0x0, 0xc0, 0xf8, 0xf, 0xff, 0x80, 0x0, 0x7f, 0xfc, 0x1, 0xff, 0xf8, 0x0, 0x7f, 0xfe, 0x0, 0xf, 0xdf, 0x80, 0x7e, 0xfc, 0x0, 0x0, 0x3, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x3f, 0xff, 0x0, 0x0, 0x0, 0x0, 0x3, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x78, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xfc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1e, 0x0, 0x0, 0x0}},
		/* R */ tinyfont.Glyph{Rune: 0x52, Width: 0x32, Height: 0x32, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0xf, 0xfc, 0x0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xfc, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x1f, 0x80, 0x0, 0x0, 0x3, 0xe0, 0x1, 0xf0, 0x0, 0x0, 0x1, 0xe0, 0x0, 0x1e, 0x0, 0x0, 0xf, 0xf0, 0x0, 0x3, 0xfc, 0x0, 0x1f, 0xf8, 0x0, 0x0, 0x7f, 0xe0, 0xf, 0xfe, 0x0, 0x0, 0x1f, 0xfc, 0x7, 0xc1, 0x0, 0x0, 0x4, 0xf, 0x83, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xf1, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1e, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0x9c, 0x0, 0x0, 0x0, 0x0, 0x0, 0xee, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0x80, 0x0, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x80, 0x0, 0x0, 0x0, 0x0, 0xf, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0x9c, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe7, 0x80, 0x0, 0x0, 0x0, 0x0, 0x78, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x1f, 0x3, 0x0, 0x0, 0x30, 0x3e, 0x3, 0xff, 0xe0, 0x0, 0x1f, 0xff, 0x0, 0x7f, 0xfe, 0x0, 0x1f, 0xff, 0x80, 0x3, 0xf7, 0xe0, 0x1f, 0xbf, 0x0, 0x0, 0x0, 0xff, 0xff, 0xc0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xc0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x18, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe0, 0x0, 0x0, 0x40, 0x0, 0x0, 0xf8, 0x0, 0x0, 0x38, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x0, 0x0, 0x10, 0x7, 0xc0, 0x0, 0x0, 0x0, 0xc, 0x1, 0xf0, 0x0, 0x0, 0x0, 0x7, 0x80, 0xfe, 0x0, 0x0, 0x0, 0x1, 0xe0, 0x3f, 0x80, 0x0, 0x0, 0x0, 0xfc, 0x7, 0xc0, 0x0, 0x0, 0x0, 0x3f, 0x0, 0xe0, 0x0, 0x0, 0x0, 0xf, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x78, 0x0, 0x0, 0x0}},
		/* W */ tinyfont.Glyph{Rune: 0x57, Width: 0x32, Height: 0x32, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0xf, 0xfc, 0x0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xfc, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x1f, 0x80, 0x0, 0x0, 0x3, 0xe0, 0x1, 0xf0, 0x0, 0x0, 0x1, 0xe0, 0x0, 0x1e, 0x0, 0x0, 0xf, 0xf0, 0x0, 0x3, 0xfc, 0x0, 0x1f, 0xf8, 0x0, 0x0, 0x7f, 0xe0, 0xf, 0xfe, 0x0, 0x0, 0x1f, 0xfc, 0x7, 0xc1, 0x0, 0x0, 0x4, 0xf, 0x83, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xf1, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1e, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0x9c, 0x0, 0x0, 0x0, 0x0, 0x0, 0xee, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0x80, 0x0, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x80, 0x0, 0x0, 0x0, 0x0, 0xf, 0x70, 0x0, 0x0, 0x0, 0x0, 0x3, 0x9c, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe7, 0x80, 0x0, 0x0, 0x0, 0x0, 0x78, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x1f, 0x3, 0x0, 0x0, 0x30, 0x3e, 0x3, 0xff, 0xe0, 0x0, 0x1f, 0xff, 0x0, 0x7f, 0xfe, 0x0, 0x1f, 0xff, 0x80, 0x3, 0xf7, 0xe0, 0x1f, 0xbf, 0x0, 0x0, 0x0, 0xff, 0xff, 0xc0, 0x0, 0x0, 0x0, 0xf, 0xff, 0xc0, 0x0, 0x0, 0xc, 0x0, 0xff, 0xc0, 0x0, 0x0, 0x13, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf6, 0x0, 0x30, 0x0, 0x0, 0x0, 0x7f, 0x80, 0x1e, 0x0, 0xc, 0x0, 0x79, 0x80, 0x7, 0x80, 0x3, 0x30, 0x1e, 0xe0, 0x31, 0xe3, 0x0, 0xfc, 0x0, 0xfc, 0x1e, 0xfd, 0xe1, 0xfe, 0x0, 0x33, 0x7, 0xff, 0xf8, 0x39, 0x80, 0xc, 0x0, 0xff, 0xfc, 0x6, 0x78, 0x0, 0x0, 0xf, 0x3c, 0x1, 0xfe, 0x0, 0x0, 0x3, 0x87, 0x0, 0xfc, 0x0, 0x0, 0x0, 0xf3, 0xc0, 0x13, 0x0, 0x0, 0x0, 0xff, 0xfc, 0x0, 0xc0, 0x0, 0x0, 0x7f, 0xff, 0x80, 0x0, 0x0, 0x0, 0x1f, 0xff, 0xe0, 0x0, 0x0, 0x0, 0x3, 0x1e, 0x30, 0x0, 0x0, 0x0, 0x0, 0x7, 0x80, 0x0, 0x0, 0x0, 0x0, 0x1, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x30, 0x0, 0x0, 0x0}},
		/* Y */ tinyfont.Glyph{Rune: 0x59, Width: 0x32, Height: 0x28, XAdvance: 0x32, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0xfc, 0x0, 0x0, 0x0, 0x0, 0x7, 0xff, 0xc0, 0x0, 0x0, 0x0, 0x7, 0xff, 0xf8, 0x0, 0x0, 0x0, 0x3, 0xe0, 0x1f, 0x0, 0x0, 0x0, 0xf, 0xf0, 0x1, 0xfc, 0x0, 0x0, 0xf, 0xf8, 0x0, 0x3f, 0xc0, 0x0, 0x7, 0xfc, 0x0, 0x7, 0xf8, 0x0, 0x3, 0xe0, 0x0, 0x0, 0xf, 0x0, 0x1, 0xe1, 0xfe, 0x0, 0x1, 0xe0, 0x0, 0x7b, 0xff, 0xf0, 0x0, 0x38, 0x0, 0x1f, 0xff, 0xff, 0x0, 0x7, 0x0, 0x7, 0xf8, 0x7, 0xe0, 0x1, 0xc0, 0x3, 0xf8, 0x0, 0x7c, 0x0, 0x70, 0x0, 0x78, 0x0, 0x7, 0x80, 0x1c, 0x3, 0xfc, 0x0, 0x0, 0xff, 0xf, 0x7, 0xfe, 0x0, 0x0, 0x1f, 0xfb, 0xc3, 0xff, 0x80, 0x0, 0x7, 0xff, 0xe1, 0xf0, 0x40, 0x0, 0x1, 0x3, 0xf0, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x3c, 0x78, 0x0, 0x0, 0x0, 0x0, 0x7, 0x1c, 0x0, 0x0, 0x0, 0x0, 0x1, 0xe7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3b, 0x80, 0x0, 0x0, 0x0, 0x0, 0xf, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x80, 0x0, 0x0, 0x0, 0x0, 0x7, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xdc, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x39, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x1e, 0x3c, 0x0, 0x0, 0x0, 0x0, 0xf, 0x7, 0xc0, 0xc0, 0x0, 0xc, 0xf, 0x80, 0xff, 0xf8, 0x0, 0x7, 0xff, 0xc0, 0x1f, 0xff, 0x80, 0x7, 0xff, 0xe0, 0x0, 0xfd, 0xf8, 0x7, 0xef, 0xc0, 0x0, 0x0, 0x3f, 0xff, 0xf0, 0x0, 0x0, 0x0, 0x3, 0xff, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0xf0, 0x0, 0x0}},
	},

	YAdvance: 0x30,
}
