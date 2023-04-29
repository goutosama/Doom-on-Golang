package map_render

import (
	"image/color"

	vars "github.com/fagirton/Doom-on-Golang/vars"
	wad "github.com/fagirton/Doom-on-Golang/wad_reader"
	"github.com/hajimehoshi/ebiten/v2"
	vector "github.com/hajimehoshi/ebiten/v2/vector"
)

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
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

func Get_Map_Bounds(v []wad.Vertex) []wad.Vertex {
	var xes []int
	var yes []int
	for i := 0; i < len(v); i++ {
		xes = append(xes, int(v[i].X))
		yes = append(yes, int(v[i].Y))
	}
	xes = Merge_sort(xes)
	yes = Merge_sort(yes)
	return []wad.Vertex{
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

func Remap_X(v []wad.Vertex, n int) int {
	out_min := 30
	out_max := int(vars.WIN_RES[0]) - 30
	return ((Max(int(v[0].X), Min(n, int(v[1].X)))-int(v[0].X))*(out_max-out_min)/(int(v[1].X)-int(v[0].X)) + out_min)
}
func Remap_Y(v []wad.Vertex, n int) int {
	out_min := 30
	out_max := int(vars.WIN_RES[1]) - 30
	return int(vars.WIN_RES[1]) - ((Max(int(v[0].Y), Min(n, int(v[1].Y)))-int(v[0].Y))*(out_max-out_min)/(int(v[1].Y)-int(v[0].Y)) - out_min)
}

func RemapVertexes(v []wad.Vertex) []wad.Vertex {
	bounds := Get_Map_Bounds(v)
	for i := 0; i < len(v); i++ {
		v[i] = wad.Vertex{
			X: int16(Remap_X(bounds, int(v[i].X))),
			Y: int16(Remap_Y(bounds, int(v[i].Y))),
		}
	}
	return v
}

func Render_Map(image *ebiten.Image, Vertexes []wad.Vertex, Linedefs []wad.Linedef) {
	for i := 0; i < len(Vertexes); i++ {
		vector.DrawFilledCircle(image, float32(Vertexes[i].X), float32(Vertexes[i].Y), 1, color.White, true)
	}
	for i := 0; i < len(Linedefs); i++ {
		vector.StrokeLine(image, float32(Vertexes[Linedefs[i].Get_St_vertex()].X), float32(Vertexes[Linedefs[i].Get_St_vertex()].Y), float32(Vertexes[Linedefs[i].Get_End_vertex()].X), float32(Vertexes[Linedefs[i].Get_End_vertex()].Y), 2, color.White, true)
	}
}
