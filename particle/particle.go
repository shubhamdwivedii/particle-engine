package particle

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle struct {
	Img             *ebiten.Image
	X               float64
	Y               float64
	Velocity        float64
	Direction       float64 // direction of movement
	Angle           float64
	AngularVelocity float64 // rotation speed
	Color           color.Color
	Size            float64
	TTL             float64
	OP              *ebiten.DrawImageOptions
}

func New(img *ebiten.Image, x, y, v, direction, angle, angV float64, col color.Color, size, ttl float64) *Particle {
	op := &ebiten.DrawImageOptions{}
	return &Particle{img, x, y, v, direction, angle, angV, col, size, ttl, op}
}

func (p *Particle) Update() {
	p.TTL--

	vx := math.Cos(p.Direction) * p.Velocity
	vy := math.Sin(p.Direction) * p.Velocity

	p.X += vx
	p.Y += vy

	p.Angle += p.AngularVelocity
}

func (p *Particle) Draw(screen *ebiten.Image) {
	p.OP.GeoM.Reset()
	sx, sy := p.Img.Size()
	p.OP.GeoM.Translate(-float64(sx)/2, -float64(sy)/2) // Move pivot to centre
	p.OP.GeoM.Rotate(p.Angle)
	p.OP.GeoM.Scale(p.Size, p.Size)

	R, G, B, _ := p.Color.RGBA()
	// rr, gg, bb, aa := float64(uint8(r))/255, float64(uint8(g))/255, float64(uint8(b))/255, float64(uint8(a))/255

	// fmt.Println(rr*255, gg*255, bb*255, aa*255)
	// // p.OP.ColorM.Scale(rr, gg, bb, aa)
	// p.OP.ColorM.Translate(255, 0, 0, 255)

	// Set color
	p.OP.ColorM.Scale(0, 0, 0, 1)
	r := float64(uint8(R>>8)) / 255
	g := float64(uint8(G>>8)) / 255
	b := float64(uint8(B>>8)) / 255
	p.OP.ColorM.Translate(r, g, b, 0)

	// p.OP.ColorM.Translate(rr, gg, bb, aa)
	p.OP.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Img, p.OP)
}
