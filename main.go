package main

import (
	"log"
	"wad"

	"github.com/hajimehoshi/ebiten/v2"
)

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
	reader := wad.WadData{wad.WadReader.NewReader()}
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
