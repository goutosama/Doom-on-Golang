package wad

import "os"

type WadReader struct {
	wad_file os.File
}

func (s WadReader) NewReader(wad_path string) (*WadReader, error) {
	wad_file, err := os.Open(wad_path)

	return &WadReader{wad_file: *wad_file}, err
}

func (s WadReader) CloseReader(w WadReader) error {
	return w.wad_file.Close()
}

type WadData struct {
	wad WadReader
}

func (s WadData) Read4Bytes(offset int64) ([]byte, error) {
	_, err0 := s.wad.wad_file.Seek(offset, 0)
	if err0 != nil {
		return []byte{0, 0, 0, 0}, err0
	}
	var buffer []byte = []byte{0, 0, 0, 0}
	_, err1 := s.wad.wad_file.Read(buffer)
	if err1 != nil {
		return []byte{0, 0, 0, 0}, err0
	}
	return buffer, nil
}

func (s WadData) ReadString(offset int64) string {
	byt, _ := s.Read4Bytes(offset)
	return string(byt)
}
