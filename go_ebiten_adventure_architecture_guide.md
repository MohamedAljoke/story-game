# Go + Ebiten Adventure Game Architecture Guide

## Purpose

This guide explains how to build a **small 2D adventure game in Go**
using **Ebiten** while maintaining **high‑quality, production‑grade
architecture**.

The goal is not just to make a prototype, but to build a **foundation
that other developers can extend**, including:

-   modding support
-   alternative render engines
-   multiplayer reuse
-   long‑term maintainability

Key principles:

-   **Clean Architecture**
-   **Domain‑Driven Design (DDD) concepts**
-   **Engine‑agnostic game logic**
-   **Separation of concerns**
-   **Mod‑friendly structure**

The rendering engine (Ebiten) should be **replaceable** without
rewriting game logic.

------------------------------------------------------------------------

# High Level Architecture

    Input Layer
        ↓
    Application Layer (use cases)
        ↓
    Domain Layer (game rules)
        ↓
    Game State
        ↓
    Rendering Adapter (Ebiten)

The **domain layer must never depend on the game engine**.

This ensures we can later switch to:

-   Raylib
-   SDL
-   Web renderer
-   Dedicated server simulation

without rewriting the core game.

------------------------------------------------------------------------

# Folder Structure

Recommended project layout:

    game/
    │
    ├── cmd/
    │   └── game/
    │       └── main.go
    │
    ├── internal/
    │
    │   ├── domain/
    │   │   ├── character.go
    │   │   ├── position.go
    │   │   ├── world.go
    │   │   └── movement.go
    │   │
    │   ├── application/
    │   │   └── move_character.go
    │   │
    │   ├── engine/
    │   │   ├── renderer.go
    │   │   └── input.go
    │   │
    │   ├── ebiten/
    │   │   └── renderer.go
    │   │
    │   └── animation/
    │       └── animation.go
    │
    └── assets/
        ├── sprites/
        └── maps/

Important idea:

    Domain → does NOT depend on engine
    Engine adapters → depend on domain

------------------------------------------------------------------------

# Domain Layer

The domain contains **pure game rules**.

No rendering. No engine types. No external libraries.

Example:

## Position

``` go
type Position struct {
    X int
    Y int
}
```

## Character

``` go
type Character struct {
    ID   string
    Name string
}
```

Notice:

Position is **separate from character identity**.

This allows:

-   items
-   NPCs
-   projectiles
-   environment objects

to reuse the same spatial logic.

------------------------------------------------------------------------

# World State

``` go
type World struct {
    Characters map[string]*Character
    Positions  map[string]*Position
}
```

Movement becomes a **world operation**.

``` go
func (w *World) MoveCharacter(id string, dx, dy int) {

    pos := w.Positions[id]

    pos.X += dx
    pos.Y += dy
}
```

The world becomes the **aggregate root** managing game state.

------------------------------------------------------------------------

# Application Layer

Application logic coordinates domain operations.

Example use case:

``` go
type MoveCharacter struct {
    World *domain.World
}

func (uc *MoveCharacter) Execute(id string, dx, dy int) {
    uc.World.MoveCharacter(id, dx, dy)
}
```

Application layer:

-   handles player commands
-   triggers domain logic
-   coordinates systems

But still does not depend on rendering.

------------------------------------------------------------------------

# Animation System

Animation should not be inside the domain.

It is a **visual concern**.

Example animation state:

``` go
type AnimationState struct {
    Current string
    Frame   int
    Time    float64
}
```

Possible animations:

    idle
    walk
    attack
    interact

Movement can trigger animation state changes:

    move → walk animation
    stop → idle animation

------------------------------------------------------------------------

# Rendering Layer

Rendering reads game state and draws it.

Renderer interface:

``` go
type Renderer interface {
    Draw(world *domain.World)
}
```

Engine implementations provide the concrete renderer.

Example:

    EbitenRenderer
    RaylibRenderer
    SDLRenderer

This keeps the game engine **pluggable**.

------------------------------------------------------------------------

# Ebiten Renderer

Example adapter:

``` go
type EbitenRenderer struct {}

func (r *EbitenRenderer) Draw(screen *ebiten.Image, world *domain.World) {

    for id, character := range world.Characters {

        pos := world.Positions[id]

        op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(float64(pos.X), float64(pos.Y))

        screen.DrawImage(characterSprite, op)
    }
}
```

Only the adapter knows about Ebiten.

------------------------------------------------------------------------

# Game Loop

Ebiten requires this structure:

``` go
type Game struct {
    world *domain.World
}

func (g *Game) Update() error {

    // read input
    // call use cases

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    renderer.Draw(screen, g.world)
}
```

Important:

Game loop calls **application logic**, not domain directly.

------------------------------------------------------------------------

# Input Layer

Input should translate **player actions → commands**.

Example:

    W → move north
    A → move west
    TAB → switch character

Input is an **adapter**, not part of domain logic.

------------------------------------------------------------------------

# Supporting Modding

To support mods later:

Design systems around **data‑driven content**.

Examples:

Mods can define:

    characters
    items
    quests
    maps
    dialogue
    animations

Possible structure:

    mods/
        mod1/
            characters.json
            items.json
            sprites/

Game loads mod data at runtime.

------------------------------------------------------------------------

# Engine Independence

Because the domain is independent, we can replace the renderer.

Example:

    Ebiten → Raylib

Only these files change:

    internal/ebiten/

Everything else remains the same.

This dramatically increases project lifespan.

------------------------------------------------------------------------

# Testing

Since the domain has no engine dependencies, we can easily test it.

Example:

``` go
func TestCharacterMovement(t *testing.T) {

    world := NewWorld()

    world.MoveCharacter("hero", 1, 0)

    pos := world.Positions["hero"]

    if pos.X != 1 {
        t.Fail()
    }
}
```

Fast and reliable tests.

------------------------------------------------------------------------

# Future Extensions

Once the basic architecture works, the following systems can be added:

## Tile Maps

Load maps created in Tiled.

## Camera System

Follow active character.

## NPC Dialogue

## Combat System

## Inventory System

## Quest System

## Multiplayer Server

Because the architecture is clean, these can be implemented without
rewriting core systems.

------------------------------------------------------------------------

# Summary

Key design principles:

-   Domain logic must be **engine‑agnostic**
-   Rendering must be **replaceable**
-   Game rules must be **testable**
-   Systems should be **mod‑friendly**
-   Code must be **clean and extensible**

The result is a **professional‑grade indie game foundation**.

This architecture supports:

-   long‑term development
-   community contributions
-   modding
-   engine portability
