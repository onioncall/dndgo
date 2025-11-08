# Fighter Setup

### `archetype`
**Description:**
Archetype is the subclass for Fighter. It can be any string and is not case sensitive

### `fighting-style`
**Description:**
At 1st level, you adopt a particular style of fighting as your specialty. Choose one of the available fighting styles. Fighting Styles not found in the Players Handbook are not yet supported

**Allowed Values:**
- "Archery"
- "Defense"
- "Dueling"
- "Great Weapon Fighting"
- "Protection"
- "Two-Weapon Fighting"

### `second-wind`
**Description:**
You have a limited well of stamina that you can draw on to protect yourself from harm. On your turn, you can use a bonus action to regain hit points equal to 1d10 + your fighter level.

**Fields:**
- `available`: int, current charges/tokens
- `maximum`: int, maximum charges/tokens (always 1, recharges on short or long rest)

### `action-surge`
**Description:**
Starting at 2nd level, you can push yourself beyond your normal limits for a moment. On your turn, you can take one additional action. At 17th level, you can use it twice before a rest, but only once on the same turn.

**Fields:**
- `available`: int, current charges/tokens
- `maximum`: int, maximum charges/tokens (1 at levels 2-16, 2 at level 17+)

### `indomitable`
**Description:**
Beginning at 9th level, you can reroll a saving throw that you fail. If you do so, you must use the new roll. At 13th level, you can use this feature twice between long rests, and at 17th level, you can use it three times.

**Fields:**
- `available`: int, current charges/tokens
- `maximum`: int, maximum charges/tokens (1 at levels 9-12, 2 at levels 13-16, 3 at level 17+)
