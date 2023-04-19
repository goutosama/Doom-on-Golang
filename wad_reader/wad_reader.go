package wad

import (
	"encoding/binary"
	"io"
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

func (s WadReader) ReadBytes(offset int64, num_bytes int) []byte {
	s.wad_file.Seek(offset, 0)
	buffer := make([]byte, num_bytes)
	io.ReadAtLeast(s.wad_file, buffer, num_bytes)
	return buffer
}

func (s WadReader) ReadString(offset int64, num_bytes int) string {
	byt := s.ReadBytes(offset, num_bytes)
	return string(byt)
}

func (s WadReader) ReadInt32(offset int64) int32 {
	byt := s.ReadBytes(offset, 4)
	return int32(binary.LittleEndian.Uint32(byt))
}

func (s WadReader) ReadInt16(offset int64) int16 {
	byt := s.ReadBytes(offset, 2)
	return int16(binary.LittleEndian.Uint16(byt))
}

type Header struct {
	signature    string
	lumps_number int32
	table_offset int32
}

func (s WadReader) ReadHeader() Header {
	return Header{
		signature:    s.ReadString(0, 4),
		lumps_number: s.ReadInt32(4),
		table_offset: s.ReadInt32(8),
	}
}

type Lump struct {
	offset int32
	size   int32
	name   string
}

func (s WadReader) Find_lump_index_by_name(directory []Lump, n string) int {
	for i := 0; i < len(directory); i++ {
		if directory[i].name == n {
			return i
		}
	}
	return -1
}

func (s WadReader) ReadDirectory(h Header) []Lump {
	var directory []Lump
	for i := 0; i < int(h.lumps_number); i++ {
		curr_offset := int64(h.table_offset) + int64(i*16)
		l := Lump{
			offset: s.ReadInt32(curr_offset),
			size:   s.ReadInt32(curr_offset + 4),
			name:   s.ReadString(curr_offset+8, 8),
		}
		directory = append(directory, l)
	}
	return directory
}

type Vertex struct {
	x int16
	y int16
}

func (s WadReader) ReadVertex(offset int64) Vertex {
	return Vertex{
		x: s.ReadInt16(offset),
		y: s.ReadInt16(offset + 2),
	}
}

var Lump_class = map[string]int{ //можно поменять на срез и его индекс+1
	"THINGS":   1,
	"LINEDEFS": 2,
	"SIDEDEFS": 3,
	"VERTEXES": 4,
	"SEGS":     5,
	"SSECTORS": 6,
	"NODES":    7,
	"SECTORS":  8,
	"REJECT":   9,
	"BLOCKMAP": 10,
}
