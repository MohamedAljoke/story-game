# Dream Walker — Roadmap

A step-by-step plan to turn the current prototype into a simple playable game. Each feature builds on the previous one.

## Goal

A single playable dream level where the player explores a tile map, avoids obstacles, talks to an NPC, finds a key, and unlocks a door. Simple but complete — covers the fundamentals of 2D game development.

---

## Phase 1: Tile Map

**What:** Replace the empty grid with a real tile map loaded from data.

**Why:** The map is the foundation. Everything else (camera, collision, interaction) depends on having a proper tile map.

**How it fits the architecture:**
- `domain/` — new `TileMap` struct: 2D grid of tile types, loaded from a simple format (2D int array or JSON)
- `domain/` — tile types: `Grass`, `Wall`, `Water`, `Path`, `Door`
- World holds the tile map as part of its state
- Renderer draws tiles instead of (or under) the debug grid

**Tiles needed:**
- Ground tiles (grass, path, stone floor)
- Wall/obstacle tiles
- At least one special tile (door)

**Result:** Player walks on a visible map with different tile types drawn on screen.

---

## Phase 2: Camera

**What:** A viewport that follows the player across a map larger than the screen.

**Why:** Right now the world IS the screen (800x600). A real game has maps bigger than what fits on screen.

**How it fits the architecture:**
- `domain/` — new `Camera` struct: position (X, Y), viewport size, `Follow(target Position)` with smooth lerp
- World holds the camera
- Renderer offsets all drawing by camera position: `drawX = worldX - camera.X`
- Map can now be e.g. 2400x1800 (3x3 screens)

**Key concepts to learn:**
- World coordinates vs screen coordinates
- Smooth camera follow (lerp: `camera.X += (target.X - camera.X) * 0.1`)
- Only draw tiles visible in the viewport (culling)

**Result:** Player moves through a large map, camera smoothly follows.

---

## Phase 3: Collision

**What:** Solid tiles that block movement.

**Why:** Without collision, the map is just decoration. Walls need to stop the player.

**How it fits the architecture:**
- Each tile type has a `Solid` property
- `World.MoveCharacter()` checks the target position against the tile map before applying movement
- Collision check: "is the tile at the new position solid? If yes, don't move."
- Works per-pixel against tile grid: convert position to tile coords, check solidity

**Key concepts to learn:**
- Tile-based collision (grid lookup, not pixel-perfect)
- AABB (axis-aligned bounding box) — the player occupies a rectangle, check all tiles it would overlap
- Sliding along walls (block X movement but allow Y, and vice versa)

**Result:** Player can't walk through walls. Map layout creates paths and rooms.

---

## Phase 4: Interaction

**What:** Player presses a key near an object to trigger an action.

**Why:** This is what makes a game a game — the player does things.

**How it fits the architecture:**
- `engine/` — new command: `CmdInteract`
- `domain/` — interactable objects placed on the map (signs, chests, doors, NPCs)
- `domain/` — interaction check: "is there an interactable object in the tile the player is facing?"
- `application/` — new `Interact` use case that resolves what happens

**Interactions for the demo level:**
- **Sign** — displays a text message
- **Chest** — gives the player an item (the key)
- **Door** — opens if player has the key, blocked otherwise
- **NPC** — triggers dialogue

**Result:** Player can walk up to things and interact with them.

---

## Phase 5: Dialogue

**What:** A text box that displays NPC dialogue or sign text.

**Why:** Dialogue gives the world personality and guides the player.

**How it fits the architecture:**
- `domain/` — `Dialogue` struct: list of lines, current line index, active flag
- World holds the active dialogue state
- Renderer draws a text box overlay when dialogue is active
- Input routes to "advance dialogue" instead of movement when dialogue is open
- `GameLoop` checks if dialogue is active and changes command routing

**Key concepts to learn:**
- Game state modes (exploring vs dialogue vs menu)
- UI overlay rendering (draw on top of the game world)
- Text rendering with word wrap

**Result:** Player talks to NPC, reads signs. Text appears in a box at the bottom of the screen.

---

## Phase 6: Game State

**What:** Track player progress — inventory, flags, events.

**Why:** Needed for "find key, open door" gameplay.

**How it fits the architecture:**
- `domain/` — `Inventory` (simple list of item IDs)
- `domain/` — `Flags` (map of string → bool for one-off events: "talked_to_npc", "chest_opened")
- World holds inventory and flags
- Interaction outcomes check and modify flags/inventory
- Door interaction checks inventory for the key item

**Result:** Player picks up key from chest, uses it to open the door. Game remembers what happened.

---

## The Demo Level

Putting it all together — one small dream world:

```
+------------------------------------------+
|  ####################################    |
|  #  .  .  .  .  .  .  .  .  .  .  #    |
|  #  .  NPC  .  .  .  .  .  .  .  .#    |
|  #  .  .  .  .  ####  .  .  .  .  #    |
|  #  .  .  .  .  #  .  .  .  .  .  #    |
|  #  .  .  .  .  #  . [chest] .  . #    |
|  #  .  .  .  .  ####  .  .  .  .  #    |
|  #  .  .  .  .  .  .  .  .  .  .  #    |
|  #  . [player] [cat]  .  .  .  .  #    |
|  #  .  .  .  .  .  .  .  .  [door]#    |
|  ####################################    |
+------------------------------------------+
```

**Flow:**
1. Player spawns in a room, cat follows
2. NPC says: "The exit is locked. I saw a key hidden somewhere..."
3. Player explores, finds chest behind wall section
4. Chest gives the key
5. Player goes to the door, presses interact
6. Door opens — level complete

---

## What Each Phase Teaches

| Phase | Game Dev Concept |
|---|---|
| Tile Map | Data-driven levels, tile rendering, map formats |
| Camera | World vs screen coordinates, viewport, culling |
| Collision | Physics basics, AABB, tile-based blocking |
| Interaction | Game verbs, proximity detection, triggers |
| Dialogue | UI state machines, overlay rendering, text display |
| Game State | Progression, inventory, event flags |

Each phase is independently buildable and testable. The game gets more complete with each one.
