package helpers

import (
	"image"
	"image/color"
	"os"

	eb "github.com/hajimehoshi/ebiten/v2"
)

func HexColor(v uint32) color.RGBA {
	return color.RGBA{R: uint8(v >> 16), G: uint8(v >> 8), B: uint8(v), A: 0xff}
}

func LoadSprite(path string) (*eb.Image, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return eb.NewImageFromImage(img), nil
}

func getFrame(sheet *eb.Image, col, row, frameW, frameH int) *eb.Image {
	x := col * frameW
	y := row * frameH

	return sheet.SubImage(
		image.Rect(x, y, x+frameW, y+frameH),
	).(*eb.Image)
}

func SliceSpriteSheet(sheet *eb.Image, rows, cols, frameW, frameH int) [][]*eb.Image {

	frames := make([][]*eb.Image, rows)

	for row := 0; row < rows; row++ {

		frames[row] = make([]*eb.Image, cols)

		for col := 0; col < cols; col++ {
			frames[row][col] = getFrame(sheet, col, row, frameW, frameH)
		}
	}

	return frames
}

func SliceRow(sheet *eb.Image, row int, start int, frameW int, frameH int, count int) []*eb.Image {
	frames := make([]*eb.Image, count)

	for i := 0; i < count; i++ {
		x := (start + i) * frameW
		y := row * frameH

		rect := image.Rect(x, y, x+frameW, y+frameH)
		frames[i] = sheet.SubImage(rect).(*eb.Image)
	}

	return frames
}
