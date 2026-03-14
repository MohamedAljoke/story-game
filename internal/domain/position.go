package domain

type Position struct {
	X, Y float64
}

func (p *Position) Clamp(minX, minY, maxX, maxY float64) {
	if p.X < minX {
		p.X = minX
	}
	if p.Y < minY {
		p.Y = minY
	}
	if p.X > maxX {
		p.X = maxX
	}
	if p.Y > maxY {
		p.Y = maxY
	}
}
