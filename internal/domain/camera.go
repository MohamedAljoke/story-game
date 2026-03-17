package domain

const CameraLerp = 0.08

type Camera struct {
	X, Y          float64
	Width, Height float64
}

func NewCamera(viewW, viewH float64) *Camera {
	return &Camera{Width: viewW, Height: viewH}
}

func (c *Camera) Follow(target *Position) {
	centerX := target.X + float64(TileSize)/2 - c.Width/2
	centerY := target.Y + float64(TileSize)/2 - c.Height/2

	c.X += (centerX - c.X) * CameraLerp
	c.Y += (centerY - c.Y) * CameraLerp
}

func (c *Camera) Clamp(maxX, maxY float64) {
	if c.X < 0 {
		c.X = 0
	}
	if c.Y < 0 {
		c.Y = 0
	}
	if c.X > maxX-c.Width {
		c.X = maxX - c.Width
	}
	if c.Y > maxY-c.Height {
		c.Y = maxY - c.Height
	}
}
