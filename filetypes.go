// Package filetypes determines file type by its magic number

package filetypes

import "github.com/akolb1/bytetrie"

// FileType is an enum defining various file types
type FileType int

const (
	Unknown FileType = iota
	Jpeg
	TiffBE
	TiffLE
	CR2 // Canon Raw version 2
	PSD // Photoshop Document file, Adobe Photoshop's native file format
	PNG // Image encoded in the Portable Network Graphics format
	BMP
	GIF
	Ico // Computer icon encoded in ICO file format
	Pcx
	Arw // Sony Alpha RAW
	Crw // Camera Image File Format
	Nef // Nikon Electronic Format)
	Orf // Olympus RAW
	Raf // FUJI RAW
	Rw2 // Panasonic raw
	Mov // Quick Time Movie
	MP3
	PS // PostScript
	PDF
	DJVU
	ZIP
	GZIP
	DMG // Apple Disk Image
)

var typeTree *bytetrie.Trie

// https://en.wikipedia.org/wiki/List_of_file_signatures
func init() {
	typeTree = bytetrie.New()
	typeTree.Insert(Jpeg, []byte{0xff, 0xd8})
	typeTree.Insert(TiffLE, []byte("II"), []byte{0x2a, 0x00})
	typeTree.Insert(TiffBE, []byte("MM"), []byte{0x00, 0x2a})
	typeTree.Insert(CR2, []byte("II*"),
		[]byte{0x00, 0x10, 0x00, 0x00, 0x00}, []byte("CR"))
	typeTree.Insert(PSD, []byte("8BPS"))
	typeTree.Insert(PNG, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A,
		0x0A})
	typeTree.Insert(BMP, []byte("BM"))
	typeTree.Insert(GIF, []byte("GIF87a"))
	typeTree.Insert(BMP, []byte("GIF89a"))
	typeTree.Insert(Ico, []byte{0x00, 0x00, 0x01, 0x00})
	typeTree.Insert(Pcx, []byte{0x0A, 0x00, 0x01})
	typeTree.Insert(Pcx, []byte{0x0A, 0x02, 0x01})
	typeTree.Insert(Pcx, []byte{0x0A, 0x03, 0x01})
	typeTree.Insert(Pcx, []byte{0x0A, 0x05, 0x01})
	typeTree.Insert(Arw, []byte("II"), []byte{0x2a, 0x00, 0x08, 0x00})
	typeTree.Insert(Crw, []byte("II"), []byte{0x1a, 0x00, 0x00, 0x00},
		[]byte("HEAPCCDR"))
	typeTree.Insert(Nef, []byte("MM"), []byte{0x00, 0x2a, 0x00, 0x00, 0x00,
		0x80, 0x00})
	typeTree.Insert(Orf, []byte("IIRO"), []byte{0x08, 0x00})
	typeTree.Insert(Orf, []byte("IIRS"), []byte{0x08, 0x00})
	typeTree.Insert(Raf, []byte("FUJIFILMCCD-RAW"))
	typeTree.Insert(Rw2, []byte("II"), []byte{0x55, 0x00})
	typeTree.Insert(Mov, []byte{0x00, 0x00, 0x00, 0x18, 0x66, 0x74, 0x79,
		0x70, 0x71, 0x74, 0x20, 0x20})
	typeTree.Insert(MP3, []byte{0x49,0x44, 0x33})
	typeTree.Insert(PS, []byte("%!PS"))
	typeTree.Insert(PDF, []byte("%PDF"))
	typeTree.Insert(ZIP, []byte("PK"), []byte{0x03, 0x04})
	typeTree.Insert(DJVU, []byte("AT&TFORM"))
	typeTree.Insert(GZIP, []byte{0x1F, 0x8B})
	// DMG is very strange - I see several different signatures
	// TODO: Investigate
	typeTree.Insert(DMG, []byte{0x78, 0xda, 0x73, 0x0D, 0x62, 0x62, 0x60})
	typeTree.Insert(DMG, []byte{0x78, 0xda, 0x63, 0x60, 0x18, 0x05})
	typeTree.Insert(DMG, []byte{0x78, 0x01, 0x63, 0x60, 0x18, 0x05})
}

// Match tries to determine file type for a sequence of bytes at the beginning
// of the file. For unknown types 'Unknown' is returned.
func Match(data []byte) FileType {
	if result, ok := typeTree.Match(data); ok {
		return result.(FileType)
	}
	return Unknown
}

// MaxKeyLen returns maximum amount of bytes needed to determine file type.
func MaxKeyLen() int {
	return typeTree.MaxDepth()
}
