# Code Structure

## Files
- `server.go`: this contains the main server/session information

## Directories
- `data`: this is all static or dynamic data within the world.
    - `data/static`: this is data which does not change over the course of gameplay
    - `data/dyanmic`: this is data which changes over the course of gameplay
    - `data/config`: this is data which is provided once at start, but then doesn't change again

- `packet`: this contains all packet decoding/handling logic. The logic in here should be minimal
  and palm off actual work to the systems.

- `system`: this contains the various in-game systems. These are often long-running processes that
  are updated at regular intervals.

- `game`: this contains in-game logic (such as resolving combat between two different units).

### `data/dynamic`

This contains all information about in-game objects. There are a number of types:

Implemented (partially):
- `Container`: containers (i.e. item with slots).
- `GameObject`: super-class for all other objects.
- `Item`: items (e.g. sword/shield). This refers to specific instances of items.
- `Player`: human-controlled characters. This is a superset of `Unit`.
- `Unit`: computer-controlled characters (e.g. monsters, enemies, trainers, quest givers, ...).

Ones left to implement:
- `Corpse`: a `Player` or `Unit` corpse.
- `DynamicObject`: interactive, non-living objects (e.g. chests).
- `Pet`: computer-controlled `Player` companions (e.g. hunter pets, warlock summons).
- `Transport`: transport-only units (e.g. griffons for flight paths).

These all inherit from each other, in the following tree (partially complete):

```
GameObject
   |
   |---> Item ---> Container
   |
   |---> Unit ---> Player
   |
   |---> DynamicObject
```
