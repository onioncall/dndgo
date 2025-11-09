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

### `class-tokens`
This is a list of class specific points/tokens/charges that you may want to keep track of. 

**Fields:**
- `name`: string
**Allowed Values:** 
    - "second-wind"
    - "action-surge"
    - "indomitable"
- `level`: int, the level required before you can use each token type
- `available`: int, current charges/token. Once you run your first update cmd on your character, do a recover and it will set this to your maximum for each token
