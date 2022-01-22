package screen

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Screen interface {
	DrawImage(image *ebiten.Image, op *ebiten.DrawImageOptions)
}
