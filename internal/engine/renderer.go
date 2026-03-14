package engine

import "story-game/internal/domain"

type Renderer interface {
	Draw(screen any, world *domain.World)
}
