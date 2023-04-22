package main

import (
	"fmt"
	"log"

	map_render "github.com/fagirton/Doom-on-Golang/render"
	wad "github.com/fagirton/Doom-on-Golang/wad_reader"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	num1 string
	num2 int32
	num3 int32
)

var (
	DOOM_RES          = []int{320, 200}
	SCALE             = float32(4.0)
	WIN_RES           = []int{DOOM_RES[0] * int(SCALE), DOOM_RES[1] * int(SCALE)}
	H_wIDTH, H_HEIGHT = WIN_RES[0] / 2, WIN_RES[1] / 2
)
var (
	reader       = wad.NewReader("doom1.wad")
	directory    = reader.ReadDirectory(reader.ReadHeader())
	map_name     = "E1M1"
	curr_map_idx = reader.Find_lump_index_by_name(directory, map_name)
	VertexesE1M1 = map_render.RemapVertexes(reader.Get_Vertex_data(directory, curr_map_idx, 0))
	LinedefsE1M1 = reader.Get_Linedef_data(directory, curr_map_idx, 0)
)

func init() {
	num1 = reader.ReadString(0, 4)
	num2 = reader.ReadInt32(4)
	num3 = reader.ReadInt32(8)
	fmt.Println(num1, num2, num3)
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, num1)
	map_render.Render_Map(screen, VertexesE1M1, LinedefsE1M1)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIN_RES[0], WIN_RES[1]
}

func main() {
	ebiten.SetWindowSize(int(WIN_RES[0]), int(WIN_RES[1]))
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
