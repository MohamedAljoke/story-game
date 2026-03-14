# Dream Walker

A 2D adventure game written in Go where a kid enters other people's dreams to cure their nightmares. Each dream is a self-contained world with its own atmosphere — surreal environments, emotional storytelling, light puzzle-solving.

Built as a personal project to explore clean architecture patterns in a game context.

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.21+ |
| Rendering | [Ebiten v2](https://ebitengine.org/) |
| Architecture | Domain-Driven Design + Ports & Adapters (Hexagonal) |

---

## Architecture

The codebase is deliberately structured so that **game logic has zero knowledge of the rendering engine**. Ebiten is a swappable adapter — replacing it with another engine (e.g. Raylib) requires no changes to domain or application code.

```
cmd/game/              Entry point — wires all layers together
internal/
  domain/              Pure game rules. No external imports, stdlib only.
  engine/              Interfaces (ports) for input and rendering
  application/         Use cases — orchestrate domain operations
  ebiten/              Ebiten adapter (implements engine interfaces)
assets/
  sprites/
  maps/
```

### Dependency Flow

```
domain  ←  engine
domain  ←  application
domain  ←  ebiten  →  ebiten/v2
```

Domain sits at the center and depends on nothing. All other layers depend inward on it.

---

## Key Design Decisions

**Domain is engine-agnostic.** `internal/domain` contains only pure Go structs and logic. It has no knowledge of Ebiten, rendering, or input systems. This makes it trivially testable and portable.

**Ports & Adapters for I/O.** `engine.InputReader` and `engine.Renderer` are interfaces defined in the engine layer. The Ebiten implementation lives in `internal/ebiten` and can be swapped without touching anything else.

**World as aggregate root.** All game state mutations (character movement, world transitions) go through the `World` struct in the domain layer. No other layer writes state directly.

**Use cases as explicit objects.** Each game action (e.g. moving a character) is a dedicated struct in `internal/application/` with an `Execute()` method, making the game loop readable and the logic easy to extend.

---

## Game Loop

```
EbitenInput.ReadCommands()     keys held → Command values
      ↓
Game.Update()                  routes commands to use cases
      ↓
MoveCharacter.Execute()        calls World.MoveCharacter() → clamps to bounds
      ↓
Game.Draw()                    delegates to EbitenRenderer.Draw()
```

---

## Running Locally

```sh
git clone <repo>
cd story-game
go run ./cmd/game/
```

Controls: `WASD` or arrow keys to move, `ESC` to quit.

---

## Project Status

Early stage — core movement, world bounds, and rendering pipeline are in place. Dream world content, puzzles, and NPC interactions are next.
