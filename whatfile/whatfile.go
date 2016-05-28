package main

import (
	"fmt"
	"os"

	"github.com/akolb1/filetypes"
)

var names = map[filetypes.FileType]string{
	filetypes.Unknown: "Unknown",
	filetypes.Jpeg:    "JPEG",
	filetypes.TiffBE:  "TIFF, Big Endian",
	filetypes.TiffLE:  "TIFF, Little Endian",
	filetypes.CR2:     "Canon Raw",
	filetypes.BMP:     "BMP",
	filetypes.GIF:     "GIF",
	filetypes.Ico:     "Icon",
	filetypes.Pcx:     "PCX",
	filetypes.Arw:     "Sony Alpha Raw",
	filetypes.Crw:     "Camera Image File Format",
	filetypes.Nef:     "Nikon Electronic Format",
	filetypes.Orf:     "Olympus RAW",
	filetypes.Raf:     "FUJI Raw",
	filetypes.Rw2:     "Panasonic RAW",
	filetypes.Mov:     "Quick Time Movie",
	filetypes.MP3:     "MP3",
	filetypes.PS:      "PostScript",
	filetypes.PDF:     "PDF document",
	filetypes.DJVU:    "DjVu document",
	filetypes.ZIP:     "ZIP archive",
	filetypes.GZIP:    "GZIP archive",
	filetypes.LZW:     "LZW compressed archive",
	filetypes.LZH:     "LZH compressed archive",
	filetypes.BZ2:     "Bzip2 compressed archive",
	filetypes.EXE:     "DOS executable",
}

func main() {
	size := filetypes.MaxKeyLen()
	data := make([]byte, size, size)
	for _, name := range os.Args[1:] {
		fd, err := os.Open(name)
		if err != nil {
			fmt.Println("can't open", name, err)
			continue
		}
		n, err := fd.Read(data)
		if err != nil {
			fmt.Println("can't read", name, err)
			continue
		}
		fmt.Printf("%s: %s\n", name, names[filetypes.Match(data[:n])])
	}
}
