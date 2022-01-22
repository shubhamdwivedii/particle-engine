package particle

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	scr "github.com/shubhamdwivedii/particle-engine/screen"
)

type Particle struct {
	Img             *ebiten.Image
	X               float64
	Y               float64
	RenderOffsetX   float64
	RenderOffsetY   float64
	Velocity        float64
	Direction       float64 // direction of movement
	Angle           float64
	AngularVelocity float64 // rotation speed
	Color           color.Color
	FadeRate        float64
	Scale           float64
	ScaleRate       float64
	TTL             float64
	OP              *ebiten.DrawImageOptions
	Alpha           float64
	ChangeColor     bool
}

func New(img *ebiten.Image, x, y, offsetX, offsetY, v, direction, angle, angV float64, col color.Color, fadeR, scale, scaleR, ttl float64, changeColor bool) *Particle {
	op := &ebiten.DrawImageOptions{}
	_, _, _, a := col.RGBA()
	alpha := float64(uint8(a>>8)) / 255
	return &Particle{img, x, y, offsetX, offsetY, v, direction, angle, angV, col, fadeR, scale, scaleR, ttl, op, alpha, changeColor}
}

func (p *Particle) Update() {
	p.TTL--

	vx := math.Cos(p.Direction) * p.Velocity
	vy := math.Sin(p.Direction) * p.Velocity

	p.X += vx
	p.Y += vy

	p.Angle += p.AngularVelocity
	p.Alpha -= p.FadeRate

	if p.Scale > 0 {
		p.Scale += p.ScaleRate
	} else {
		if p.Scale < 0 {
			p.Scale = 0
		}
	}
}

func (p *Particle) Draw(screen scr.Screen) {
	p.OP.GeoM.Reset()
	sx, sy := p.Img.Size()
	p.OP.GeoM.Translate(-float64(sx)/2, -float64(sy)/2) // Move pivot to centre
	p.OP.GeoM.Rotate(p.Angle)
	p.OP.GeoM.Scale(p.Scale, p.Scale)

	R, G, B, _ := p.Color.RGBA()

	// Set color
	if p.ChangeColor {
		p.OP.ColorM.Scale(0, 0, 0, p.Alpha)
		r := float64(uint8(R>>8)) / 255
		g := float64(uint8(G>>8)) / 255
		b := float64(uint8(B>>8)) / 255
		p.OP.ColorM.Translate(r, g, b, 0)
	} else {
		p.OP.ColorM.Scale(1, 1, 1, p.Alpha)
	}

	p.OP.GeoM.Translate(p.X+p.RenderOffsetX, p.Y+p.RenderOffsetY)
	screen.DrawImage(p.Img, p.OP)
}
