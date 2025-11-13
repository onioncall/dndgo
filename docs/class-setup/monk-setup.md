# Monk Setup

### `mosaic-tradition`
**Description:**
Monastic Tradition is the subclass for Monk. It can be any string and is not case sensitive

### `class-token`

**Description:**
Class Token is a way to track class specific tokens/charges/points. The token available to the monk class is ki-points

**Fields:**
- `name`: "ki-points", (the only token available to this class is ki-points)
- `available`: int, current charges/tokens. Feel free to set this to 0, and use `dndgo ctr recover` to set to maximum available to your level 
- `level`: 1, (ki-points is available from level 1)
