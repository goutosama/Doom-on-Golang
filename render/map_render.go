package map_render

import (
	"fmt"

	wad "github.com/fagirton/Doom-on-Golang/wad_reader"
)

func Render_Map(w wad.WadReader, d []wad.Lump, map_name string) {
	Vertexes := w.Get_Vertex_data(d, w.Find_lump_index_by_name(d, map_name), 4, 0)
	fmt.Println(Vertexes)
}
