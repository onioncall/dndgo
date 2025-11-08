# Druid Setup

### `circle`
**Description:**
Circle is the subclass for Druid. It can be any string and is not case sensitive

### `wild-shape`
**Description:**
Starting at 2nd level, you can use your action to magically assume the shape of a beast that you have seen before. You can use this feature twice. You regain expended uses when you finish a short or long rest. At higher levels, you can transform into more powerful beasts and stay in beast form for longer.

**Fields:**
- `available`: int, current uses
- `maximum`: int, maximum uses (always 2, recharges on short or long rest)

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


