# Paladin Setup 

### `sub-class`
**Description**
Sacred Oath is the subclass for Paladin. It can be any string and is not case sensitive.

### `fighting-style`
**Description:**
At 1st level, you adopt a particular style of fighting as your specialty. Choose one of the available fighting styles. Fighting Styles not found in the Players Handbook are not yet supported

**Allowed Values:**
- "Defense"
- "Dueling"
- "Great Weapon Fighting"
- "Protection"

### `prepared-spells`
**Description:**
Paladins can prepare a number of spells equal to their Charisma modifier + paladin level (minimum of 1). Prepared spells are chosen from your spellbook and can be changed after each long rest. Prepared spells must be in the list of your known spells in your character config and be spelled the same.

**Examples**: *not a comprehensive list*
- "Aid"
- "Find Steed"
- "Lesser Restoration"
- "Magic Weapon"

### `oath-spells`
**Description:**
Paladins also have oath spells that they get access to at levels 3, 5, 9, 13, and 17. These spells are always prepared. oath-spells is a list of strings with the spell name similar to prepared-spells

### `class-tokens`
This is a list of class specific points/tokens/charges that you may want to keep track of. 

**Fields:**
- `name`: string
**Allowed Values:** 
    - "divine-sense"
    - "lay-on-hands"
- `level`: int, the level required before you can use each token type
- `available`: int, current charges/token. Once you run your first update cmd on your character, do a recover and it will set this to your maximum for each token
