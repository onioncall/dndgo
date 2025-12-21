# Sorcerer Setup

### `sub-class`
**Description:**
Sorcerous Origin is the subclass for Sorcerer. It can be any string and is not case sensitive

### `class-token`

**Description:**
Class Token is a way to track class specific tokens/charges/points. The token available to the sorcerer class is sorcery-points

**Fields:**
- `name`: "sorcery-points", (the only token available to this class is sorcery-points)
- `available`: int, current charges/tokens. Feel free to set this to 0, and use `dndgo ctr recover` to set to maximum available to your level 
- `level`: 1, (sorcery-points is available from level 1)

### `meta-magic-spells`
**Description:**
At 3rd level, you gain the ability to twist your spells to suit your needs. You gain two Metamagic options of your choice. You gain another one at 10th and 17th level. You can use only one Metamagic option on a spell when you cast it, unless otherwise noted.

**Fields:**
- `name`: string, the name of the metamagic option
- `level`: int, the level at which you gained this metamagic (3, 10, or 17)
- `details`: string, description of what the metamagic does

**Example:**
```json
"meta-magic-spells": [
  {
    "name": "Quickened Spell",
    "level": 3,
    "details": "When you cast a spell that has a casting time of 1 action, you can spend 2 sorcery points to change the casting time to 1 bonus action for this casting.\n"
  },
  {
    "name": "Twinned Spell",
    "level": 3,
    "details": "When you cast a spell that targets only one creature and doesn't have a range of self, you can spend a number of sorcery points equal to the spell's level to target a second creature in range with the same spell (1 sorcery point if the spell is a cantrip).\n"
  }
]
```
