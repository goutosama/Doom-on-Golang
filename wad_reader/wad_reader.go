package wad

import (
	"encoding/binary"
	"io"
	"os"
	"strings"
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
	Offset int32
	Size   int32
	Name   string
}

func (s WadReader) Find_lump_index_by_name(directory []Lump, n string) int {
	for i := 0; i < len(directory); i++ {
		if directory[i].Name == n {
			return i
		}
	}
	return -1
}

func (s WadReader) ReadDirectory(h Header) []Lump {
	var directory []Lump
	for i := 0; i < int(h.lumps_number); i++ {
		curr_offset := int64(h.table_offset) + int64(i*16)
		name := s.ReadString(curr_offset+8, 8)
		for strings.HasSuffix(name, "\x00") {
			name = strings.TrimSuffix(name, "\x00")
		}
		l := Lump{
			Offset: s.ReadInt32(curr_offset),
			Size:   s.ReadInt32(curr_offset + 4),
			Name:   name,
		}
		directory = append(directory, l)
	}
	return directory
}

func (s WadReader) ReadLinedef(offset int64) Linedef {
	return Linedef{
		St_Vertex:  s.ReadInt16(offset),
		End_Vertex: s.ReadInt16(offset + 2),
		Flags:      s.ReadInt16(offset + 4),
		Linetype:   s.ReadInt16(offset + 6),
		Sector_tag: s.ReadInt16(offset + 8),
		F_Sidedef:  s.ReadInt16(offset + 10),
		B_Sidedef:  s.ReadInt16(offset + 12),
	}
}

func (s WadReader) Get_Linedef_data(directory []Lump, index_of_map int, header_length int) []Linedef {
	l := directory[index_of_map+Lump_class["LINEDEFS"]]
	var Lump_data []Linedef
	for i := 0; i < (int(l.Size) / 14); i++ {
		offset := int(l.Offset) + i*14 + header_length
		Lump_data = append(Lump_data, s.ReadLinedef(int64(offset)))
	}
	return Lump_data
}

type Vertex struct {
	X int16
	Y int16
}

func (s WadReader) ReadVertex(offset int64) Vertex {
	return Vertex{
		X: s.ReadInt16(offset),
		Y: s.ReadInt16(offset + 2),
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

func (s WadReader) Get_Vertex_data(directory []Lump, index_of_map int, header_length int) []Vertex {
	l := directory[index_of_map+Lump_class["VERTEXES"]]
	var Lump_data []Vertex
	for i := 0; i < (int(l.Size) / 4); i++ {
		offset := int(l.Offset) + i*4 + header_length
		Lump_data = append(Lump_data, s.ReadVertex(int64(offset)))
	}
	return Lump_data
}

func Merge(a []int, b []int) []int {
	c := []int{}
	var i int = 0
	var j int = 0
	for k := 0; k < len(a)+len(b); k++ {
		if (j == len(b)) || ((i < len(a)) && (a[i] <= b[j])) {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	return c
}

func Merge_sort(v []int) []int {
	if len(v) <= 1 {
		return v
	}
	left := Merge_sort(v[:len(v)/2])
	right := Merge_sort(v[len(v)/2:])
	return Merge(left, right)
}

func Get_Map_Bounds(v []Vertex) []Vertex {
	var xes []int
	var yes []int
	for i := 0; i < len(v); i++ {
		xes = append(xes, int(v[i].X))
		yes = append(yes, int(v[i].Y))
	}
	xes = Merge_sort(xes)
	yes = Merge_sort(yes)
	return []Vertex{
		{
			X: int16(xes[0]),
			Y: int16(yes[0]),
		},
		{
			X: int16(xes[len(xes)-1]),
			Y: int16(yes[len(yes)-1]),
		},
	}
}

type Linedef struct {
	St_Vertex  int16
	End_Vertex int16
	Flags      int16
	Linetype   int16
	Sector_tag int16
	F_Sidedef  int16
	B_Sidedef  int16
}
