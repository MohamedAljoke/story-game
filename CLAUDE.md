# Dream Walker

2D adventure game built in Go with Ebiten. Clean architecture — domain logic is engine-agnostic, Ebiten is a replaceable rendering adapter.

## Game Concept

**Dream Walker** — A kid who can enter other people's dreams to cure their nightmares. Each dream is a small self-contained world with its own visual style and rules. Surreal environments, emotional storytelling, light puzzle-solving.

The player navigates distinct dream worlds, each belonging to a different dreamer. Worlds shift in atmosphere — a horror corridor, a childhood memory, a fantasy realm — and the player must explore, interact with dream fragments, and resolve the nightmare at the core.

WHEN WE MAKE AN ARCHITECTURE CHANGE UPDATE /docs/architecture.md

## Quick Start

```sh
go build ./cmd/game/
go run ./cmd/game/
```

- 800x600 window, "Dream Walker" title
- WASD or arrow keys to move
- ESC to quit

Terminal stub (no graphics):

```sh
go run ./cmd/game-raylib/
```

- w/a/s/d to move, q to quit

## Project Structure

```
cmd/game/main.go                   Entry point (Ebiten)
cmd/game-raylib/main.go            Entry point (Raylib terminal stub)
internal/
  domain/                           Pure game rules, zero external imports
    position.go                     Position{X,Y float64}, Clamp()
    direction.go                    Direction enum (Up/Down/Left/Right), Delta()
    character.go                    Character{ID, Name}
    world.go                        Aggregate root, MoveCharacter(), constants
  engine/                           Engine-agnostic interfaces
    input.go                        Command enum, InputReader interface
    renderer.go                     Renderer interface (screen is `any`)
    engine.go                       Engine interface (Run() error)
  application/                      Use cases coordinating domain operations
    move_character.go               MoveCharacter use case
    game_loop.go                    GameLoop: routes commands to use cases
  adapter/
    ebiten/                         Ebiten adapter
      input.go                      EbitenInput: keys → commands
      renderer.go                   EbitenRenderer: draws world state
      game.go                       Game struct, Run() wraps ebiten.RunGame
    raylib/                         Raylib terminal stub adapter
      input.go                      RaylibInput: stdin → commands
      renderer.go                   RaylibRenderer: prints positions to stdout
      game.go                       Game struct, Run() imperative loop
assets/
  sprites/                          Sprite assets (empty, placeholder)
  maps/                             Map assets (empty, placeholder)
```

## Architecture & Dependency Rules

```
domain/         → nothing (stdlib only, NO engine imports)
engine/         → domain
application/    → domain, engine
adapter/ebiten/ → domain, engine, application, ebiten/v2
adapter/raylib/ → domain, engine, application, stdlib
cmd/game/       → adapter/ebiten, application, domain
cmd/game-raylib/→ adapter/raylib, application, domain
```

Domain must never import engine packages. Engine adapters depend inward on domain. Application layer may import engine (for Command type) but not adapters.

## Domain Constants

Defined in `internal/domain/world.go`:

| Constant    | Value |
| ----------- | ----- |
| WorldWidth  | 800   |
| WorldHeight | 600   |
| TileSize    | 32    |
| MoveSpeed   | 3.0   |

## How the Game Loop Works

1. Adapter's `InputReader.ReadCommands()` translates input into `Command` values
2. `GameLoop.ProcessCommands()` routes commands to use cases (shared across all adapters)
3. `MoveCharacter` calls `World.MoveCharacter()` which applies direction delta \* speed, then clamps to bounds
4. Adapter's `Renderer.Draw()` reads world state and renders output

## Adding a New Feature

**New domain concept** — add to `internal/domain/`, no engine imports allowed. Update `World` if it's aggregate state.

**New input command** — add constant to `internal/engine/input.go`, handle in `internal/application/game_loop.go` ProcessCommands(), add key mapping in each adapter's input.go.

**New use case** — add file in `internal/application/`, inject into `GameLoop`, route in `ProcessCommands()`.

**New visual element** — add to each adapter's renderer.go Draw method.

**Swap engine** — create new adapter package under `internal/adapter/`, implement `engine.InputReader`, `engine.Renderer`, and `engine.Engine` interfaces. Create a new `cmd/` entry point that wires the adapter with `GameLoop`. See `internal/adapter/raylib/` for a minimal example.

## Conventions

- Ebiten v2 imports are aliased as `eb` in adapter files
- `cmd/game/main.go` uses import alias `ebitenpkg` for `internal/adapter/ebiten`
- World is the aggregate root — all game state mutations go through it
- Renderer receives `screen any` and type-asserts to the engine's image type
- Command routing lives in `GameLoop.ProcessCommands()`, not in adapter code
- `go vet ./...` must pass clean
