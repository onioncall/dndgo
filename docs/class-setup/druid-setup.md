# Druid Setup

### `circle`
**Description:**
Circle is the subclass for Druid. It can be any string and is not case sensitive

### `class-token`

**Description:**
Class Token is a way to track class specific tokens/charges/points. The token available to the druid class is wild-shape

**Fields:**
- `name`: "wild-shape", (the only token available to this class is wild-shape)
- `available`: int, current charges/tokens. Feel free to set this to 0, and use `dndgo ctr recover` to set to maximum available to your level 
- `level`: 1, (wild-shape is available from level 1)

### `prepared-spells`
**Description:**
Druids can prepare a number of spells equal to their Wisdom modifier + druid level (minimum of 1). Circle spells are always prepared and do not count against this limit. Prepared spells can be changed after each long rest. Prepared spells must be in the list of your known spells in your character config.


**Allowed Values:** Array of spell names from the druid spell list
- Must be spell names as strings
- Can include cantrips or leveled spells
- Total prepared spells (excluding circle spells) = Wisdom modifier + druid level

**Examples**: *not a comprehensive list*
- "Healing Word"
- "Entangle"
- "Moonbeam"
- "Pass Withou Grace"


