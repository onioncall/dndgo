# Barbarian Setup

### `sub-class`
**Description:**
Path is the subclass for Barbarian. It can be any string and is not case sensitive

### `class-token`
**Description:**
Class Token is a way to track class specific tokens/charges/points. The token available to the barbarian class is rage

**Fields:**
- `name`: "rage", (the only token available to this class is rage)
- `available`: int, current charges/tokens. Feel free to set this to 0, and use `dndgo ctr recover` to set to maximum available to your level 
- `level`: 1, (rage is available from level 1)

### `primal-knowledge`
**Description:**
When you reach 3rd level and again at 10th level, you gain proficiency in one skill of your choice from the list of skills available to barbarians at 1st level. Only use this if you are including *Tasha's Cauldron of Everything* as part of your campaign

**Allowed Values:** 
- "Animal Handling"
- "Athletics"
- "Intimidation"
- "Nature"
- "Perception"
- "Survival"

