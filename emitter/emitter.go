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
	OP        EmitterOptions
}

type EmitterOptions struct {
	MinVelocity        float64
	MaxVelocity        float64
	MinDirection       float64
	MaxDirection       float64
	MinAngularVelocity float64
	MaxAngularVelocity float64
	MinSize            float64
	MaxSize            float64
	MinTTL             float64
	MaxTTL             float64
	MinFadeRate        float64
	MaxFadeRate        float64
}

func NewEmitterOptions() EmitterOptions {
	return EmitterOptions{-2, 2, -math.Pi, math.Pi, 0.02, 0.4, 0.5, 1.5, 20, 40, 0.002, 0.01}
}

func New(textures []*ebiten.Image, x, y float64, colors []color.Color, options EmitterOptions) *Emitter {
	// var particles []*p.Particle
	particles := list.New()
	return &Emitter{x, y, particles, textures, colors, options}
}

func GetRandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func (e *Emitter) Generate() {
	img := e.Textures[rand.Intn(len(e.Textures))]
	x, y := e.X, e.Y
	velocity := GetRandomFloat64(e.OP.MinVelocity, e.OP.MaxVelocity)
	direction := GetRandomFloat64(e.OP.MinDirection, e.OP.MaxDirection)
	angularV := GetRandomFloat64(e.OP.MinAngularVelocity, e.OP.MaxAngularVelocity)
	col := e.Colors[rand.Intn(len(e.Colors))]
	size := GetRandomFloat64(e.OP.MinSize, e.OP.MaxSize)
	ttl := GetRandomFloat64(e.OP.MinTTL, e.OP.MaxTTL)
	fadeRate := GetRandomFloat64(e.OP.MinFadeRate, e.OP.MaxFadeRate)

	particle := p.New(img, x, y, velocity, direction, 0, angularV, col, fadeRate, size, ttl)
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
