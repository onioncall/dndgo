# Cleric Setup

### `domain`
**Description:**
Domain is the subclass for Cleric. It can be any string and is not case sensitive

### `channel-divinity`
**Description:**
Clerics have a tracked number of Channel Divinity uses they can expend before a rest. Channel Divinity is a special ability granted by your divine domain.

**Fields:**
- `available`: int, current charges/tokens
- `maximum`: int, maximum charges/tokens (1 at levels 2-5, 2 at levels 6-17, 3 at level 18+)

### `prepared-spells`
**Description:**
Clerics can prepare a number of spells equal to their Wisdom modifier + cleric level. Domain spells are always prepared and do not count against this limit. Prepared spells can be changed after each long rest. Prepared spells must be in the list of your known spells in your character config

**Examples**: *not a comprehensive list*
- "Cure Wounds"
- "Bless"
- "Spiritual Weapon"
- "Guiding Bolt"
