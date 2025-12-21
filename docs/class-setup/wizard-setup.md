# Wizard Setup

### `sub-class`
**Description:**
Arcane Tradition is the subclass for Wizard. It can be any string and is not case sensitive

### `prepared-spells`
**Description:**
Wizards can prepare a number of spells equal to their Intelligence modifier + wizard level (minimum of 1). Prepared spells are chosen from your spellbook and can be changed after each long rest. Prepared spells must be in the list of your known spells in your character config.

**Examples**: *not a comprehensive list*
- "Fireball"
- "Counterspell"
- "Misty Step"
- "Shield"
- "Magic Missile"

### `signature-spells`
**Description:**
At 20th level, you gain mastery over two powerful spells and can cast them with little effort. Choose two 3rd-level wizard spells in your spellbook as your signature spells. You always have these spells prepared, they don't count against the number of spells you can prepare, and you can cast each of them once at 3rd level without expending a spell slot.

**Allowed Values:** Array of exactly two 3rd-level spell names from your spellbook
- Only available at level 20
- Must be 3rd-level spells
- Must be spells in your spellbook

**Examples**: *not a comprehensive list*
- "Fireball"
- "Counterspell"

