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
cmd/game/              Entry point (Ebiten)
cmd/game-raylib/       Entry point (terminal stub)
internal/
  domain/              Pure game rules. No external imports, stdlib only.
  engine/              Interfaces (ports): InputReader, Renderer, Engine
  application/         Use cases — MoveCharacter, GameLoop
  adapter/
    ebiten/            Ebiten adapter (implements engine interfaces)
    raylib/            Raylib terminal stub adapter
assets/
  sprites/
  maps/
```

### Dependency Flow

```
domain  ←  engine
domain  ←  application  ←  engine
domain  ←  adapter/*  →  engine, application, (engine-specific lib)
```

Domain sits at the center and depends on nothing. Command routing lives in `GameLoop.ProcessCommands()` — shared across all adapters.

---

## Key Design Decisions

**Domain is engine-agnostic.** `internal/domain` contains only pure Go structs and logic. It has no knowledge of Ebiten, rendering, or input systems. This makes it trivially testable and portable.

**Ports & Adapters for I/O.** `engine.InputReader`, `engine.Renderer`, and `engine.Engine` are interfaces defined in the engine layer. Adapters live under `internal/adapter/` and can be swapped without touching domain or application code.

**World as aggregate root.** All game state mutations (character movement, world transitions) go through the `World` struct in the domain layer. No other layer writes state directly.

**Use cases as explicit objects.** Each game action (e.g. moving a character) is a dedicated struct in `internal/application/` with an `Execute()` method, making the game loop readable and the logic easy to extend.

---

## Game Loop

```
InputReader.ReadCommands()       keys/stdin → Command values
      ↓
GameLoop.ProcessCommands()       routes commands to use cases (shared)
      ↓
MoveCharacter.Execute()          calls World.MoveCharacter() → clamps to bounds
      ↓
Renderer.Draw()                  renders world state (Ebiten or stdout)
```

---

## Running Locally

```sh
git clone <repo>
cd story-game
make run        # Ebiten window
make run-term   # terminal stub (w/a/s/d + q)
```

Controls: `WASD` or arrow keys to move, `ESC` to quit.

---

## Project Status

Early stage — core movement, world bounds, and rendering pipeline are in place. Dream world content, puzzles, and NPC interactions are next.
