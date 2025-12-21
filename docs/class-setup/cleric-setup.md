# Cleric Setup

### `sub-class`

**Description:**
Domain is the subclass for Cleric. It can be any string and is not case sensitive

### `class-token`

**Description:**
Class Token is a way to track class specific tokens/charges/points. The token available to the cleric class is channel-divinity

**Fields:**
- `name`: "channel-divinity", (the only token available to this class is channel-divinity)
- `available`: 0, (current charges/tokens. Use `dndgo ctr recover` to set to maximum available to your level )
- `level`: 1, (channel-divinity is available from level 1)

### `prepared-spells`
**Description:**
Clerics can prepare a number of spells equal to their Wisdom modifier + cleric level. Domain spells are always prepared and do not count against this limit. Prepared spells can be changed after each long rest. Prepared spells must be in the list of your known spells in your character config

**Examples**: *not a comprehensive list*
- "Cure Wounds"
- "Bless"
- "Spiritual Weapon"
- "Guiding Bolt"
