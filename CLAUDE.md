# Dream Walker

2D adventure game built in Go with Ebiten. Clean architecture — domain logic is engine-agnostic, Ebiten is a replaceable rendering adapter.

## Game Concept

**Dream Walker** — A kid who can enter other people's dreams to cure their nightmares. Each dream is a small self-contained world with its own visual style and rules. Surreal environments, emotional storytelling, light puzzle-solving.

The player navigates distinct dream worlds, each belonging to a different dreamer. Worlds shift in atmosphere — a horror corridor, a childhood memory, a fantasy realm — and the player must explore, interact with dream fragments, and resolve the nightmare at the core.

## Quick Start

```sh
go build ./cmd/game/
go run ./cmd/game/
```

- 800x600 window, "Dream Walker" title
- WASD or arrow keys to move
- ESC to quit

## Project Structure

```
cmd/game/main.go              Entry point, wires all layers together
internal/
  domain/                      Pure game rules, zero external imports
    position.go                Position{X,Y float64}, Clamp()
    direction.go               Direction enum (Up/Down/Left/Right), Delta()
    character.go               Character{ID, Name}
    world.go                   Aggregate root, MoveCharacter(), constants
  engine/                      Engine-agnostic interfaces
    input.go                   Command enum, InputReader interface
    renderer.go                Renderer interface (screen is `any`)
  application/                 Use cases coordinating domain operations
    move_character.go          MoveCharacter use case
  ebiten/                      Ebiten adapter implementations
    input.go                   EbitenInput: keys → commands
    renderer.go                EbitenRenderer: draws world state
    game.go                    Game struct implementing ebiten.Game
assets/
  sprites/                     Sprite assets (empty, placeholder)
  maps/                        Map assets (empty, placeholder)
```

## Architecture & Dependency Rules

```
domain/       → nothing (stdlib only, NO engine imports)
engine/       → domain
application/  → domain
ebiten/       → domain, engine, application, ebiten/v2
cmd/game/     → domain, internal/ebiten, ebiten/v2
```

Domain must never import engine packages. Engine adapters depend inward on domain.

## Domain Constants

Defined in `internal/domain/world.go`:

| Constant    | Value |
|-------------|-------|
| WorldWidth  | 800   |
| WorldHeight | 600   |
| TileSize    | 32    |
| MoveSpeed   | 3.0   |

## How the Game Loop Works

1. `EbitenInput.ReadCommands()` translates held keys into `Command` values
2. `Game.Update()` iterates commands, routes move commands through `MoveCharacter.Execute()`
3. `MoveCharacter` calls `World.MoveCharacter()` which applies direction delta * speed, then clamps to bounds
4. `Game.Draw()` delegates to `EbitenRenderer.Draw()` which reads world state and renders background, grid, characters, HUD

## Adding a New Feature

**New domain concept** — add to `internal/domain/`, no engine imports allowed. Update `World` if it's aggregate state.

**New input command** — add constant to `internal/engine/input.go`, handle key mapping in `internal/ebiten/input.go`, route in `internal/ebiten/game.go` Update().

**New use case** — add file in `internal/application/`, inject `World`, wire in `cmd/game/main.go` and `internal/ebiten/game.go`.

**New visual element** — add to `internal/ebiten/renderer.go` Draw method.

**Swap engine** — implement `engine.InputReader` and `engine.Renderer` interfaces, create new adapter package (e.g. `internal/raylib/`), update `cmd/game/main.go` wiring.

## Conventions

- Ebiten v2 imports are aliased as `eb` in adapter files to avoid collision with the `internal/ebiten` package name
- `cmd/game/main.go` uses import alias `ebitenpkg` for `internal/ebiten`
- World is the aggregate root — all game state mutations go through it
- Renderer receives `screen any` and type-asserts to the engine's image type
- `go vet ./...` must pass clean
