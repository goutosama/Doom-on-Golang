package main

import (
	"fmt"
	"log"

	wad "github.com/fagirton/Doom-on-Golang/wad_reader"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	num1 string
	num2 int32
	num3 int32
)

func init() {
	reader := wad.NewReader("doom1.wad")

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
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
