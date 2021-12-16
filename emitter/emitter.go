package Emitter

import (
	"container/list"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	p "github.com/shubhamdwivedii/particle-engine/particle"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

type Emitter struct {
	X         float64 // emmiter location
	Y         float64
	Particles *list.List // list of pointers
	Textures  []*ebiten.Image
	Colors    []color.Color
}

func New(textures []*ebiten.Image, x, y float64, colors []color.Color) *Emitter {
	// var particles []*p.Particle
	particles := list.New()
	return &Emitter{x, y, particles, textures, colors}
}

func GetRandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func (e *Emitter) Generate() {
	img := e.Textures[rand.Intn(len(e.Textures))]
	x, y := e.X, e.Y
	velocity := GetRandomFloat64(-2, 2)
	direction := GetRandomFloat64(-math.Pi, math.Pi)
	angularV := GetRandomFloat64(0.02, 0.4)

	col := e.Colors[rand.Intn(len(e.Colors))]
	size := GetRandomFloat64(0.5, 1.5)

	ttl := GetRandomFloat64(20, 40)

	particle := p.New(img, x, y, velocity, direction, 0, angularV, col, size, ttl)

	e.Particles.PushBack(particle)
}

func (e *Emitter) MoveTo(x, y float64) {
	e.X = x
	e.Y = y
}

func (e *Emitter) MoveBy(dx, dy float64) {
	e.X += dx
	e.Y += dy
}

func (e *Emitter) Update() {
	for part := e.Particles.Front(); part != nil; part = part.Next() {
		particle := part.Value.(*p.Particle)
		particle.Update()
		if particle.TTL <= 0 {
			defer e.Particles.Remove(part)
		}
	}
}

func (e *Emitter) Draw(screen *ebiten.Image) {
	for part := e.Particles.Front(); part != nil; part = part.Next() {
		particle := part.Value.(*p.Particle)
		particle.Draw(screen)
	}
}
