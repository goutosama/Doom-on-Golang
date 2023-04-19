package wad

import (
	"os"
)

type WadReader struct {
	wad_file *os.File
}

func NewReader(wad_path string) WadReader {
	wad_file, _ := os.Open(wad_path)
	return WadReader{wad_file: wad_file}
}

func (s WadReader) CloseReader(w WadReader) error {
	return w.wad_file.Close()
}

func (s WadReader) Read4Bytes(offset int64) []byte {
	_, err0 := s.wad_file.Seek(offset, 0)
	if err0 == nil {
		panic(err0)
	}
	var buffer []byte = []byte{0, 0, 0, 0}
	_, err1 := s.wad_file.Read(buffer)
	if err1 == nil {
		panic(err1)
	}
	return buffer
}

func (s WadReader) ReadString(offset int64) string {
	byt := s.Read4Bytes(offset)
	return string(byt)
}
