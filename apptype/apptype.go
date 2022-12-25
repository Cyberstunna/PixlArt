package apptype

import (
	"fyne.io/fyne/v2"
)

type BrushType int

type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	pxRows, PxCols int
	PxSize         int
}
