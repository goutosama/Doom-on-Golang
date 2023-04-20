package main

import (
	"fmt"
	"log"

	map_render "github.com/fagirton/Doom-on-Golang/map_render"
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
	DOOM_RES          = []float32{320, 200}
	SCALE             = float32(4.0)
	WIN_RES           = []float32{DOOM_RES[0] * SCALE, DOOM_RES[1] * SCALE}
	H_wIDTH, H_HEIGHT = WIN_RES[0], WIN_RES[1]
)

func init() {
	reader := wad.NewReader("doom1.wad")
	map_render.Render_Map(reader)
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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(int(WIN_RES[0]), int(WIN_RES[1]))
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
