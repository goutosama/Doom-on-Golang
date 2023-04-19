package main

import (
	"fmt"
	"log"

	wad "github.com/fagirton/Doom-on-Golang/wad_reader"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	reader := wad.NewReader("/doom1.wad")
	var (
		num1 string
		num2 string
		num3 string
	)
	num1 = reader.ReadString(0)
	num2 = string(reader.Read4Bytes(4))
	num3 = string(reader.Read4Bytes(8))
	fmt.Println(num1, num2, num3)
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
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
