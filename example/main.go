package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	// p "github.com/shubhamdwivedii/particle-engine/particle"
	e "github.com/shubhamdwivedii/particle-engine/emitter"
)

type Game struct{}

// var particle *p.Particle
var emitter *e.Emitter

var cpx, cpy int

func init() {
	img, _, err := ebitenutil.NewImageFromFile("./assets/particle.png")
	if err != nil {
		log.Fatal(err)
	}
	// particle = p.New(img, 160, 120, 0, math.Pi/2, 0, 0.02, color.RGBA{128, 0, 0, 255}, 2, 6000)
	options := e.NewEmitterOptions()
	emitter = e.New([]*ebiten.Image{img}, 160, 120, []color.Color{
		color.RGBA{147, 231, 251, 255},
		color.RGBA{192, 246, 251, 255},
		color.RGBA{240, 250, 255, 255},
		color.RGBA{224, 255, 255, 255},
	}, options)
}

func (g *Game) Update() error {
	cpx, cpy = ebiten.CursorPosition()
	emitter.MoveTo(float64(cpx), float64(cpy))
	emitter.Generate()
	emitter.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v %v", cpx, cpy))
	emitter.Draw(screen)
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
