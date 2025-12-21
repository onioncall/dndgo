# Character Setup Guide

This is to aid your setup of your character for the character.json file. As a note, this is not meant as a comprehensive guide on how to set up a DND character period. You should still [RTFM](https://archive.org/details/players-handbook_202308). This is purely meant as a guide to help set up a character of this particular application 

## Basic Character Details

### `path`

**Description:**
This is the relative path to the markdown, if it's blank it will default to ".config/dndgo" so it's not far from the config, we recommend you store it here.

### `name`

**Description:**
The name of your character, whatever you want I guess...


### `level`

**Description:**
The level of your character.

### `class-name`

**Description:**
The class of your character. *Currently, multiclassing is not supported

**Allowed Values:**
- "barbarian"
- "bard"
- "cleric"
- "druid"
- "fighter"
- "monk"
- "paladin" *not implemented yet
- "rogue"
- "sorcerer" *not implemented yet
- "warlock"
- "wizard"

### `race`

**Description:**
The race of your character.

**Allowed Values:**
- "aasimar"
- "dragonborn"
- "dwarf"
- "elf"
- "gnome"
- "half-elf"
- "half-orc"
- "halfling"
- "human"
- "variant-human"
- "tiefling"
- "goliath"
- "tabaxi"
- "kenku"
- "triton"
- "warforged"
- "changeling"
- "kalashtar"
- "shifter"
- "orc"
- "kobold"
- "lizardfolk"
- "tortle"
- "githyanki"
- "githzerai"

### `background`

**Description:** *optional*
The background of your character.

### `feats`

**Description:** *optional*
A list of feats your character has.

**Fields:**
- `name`: string, name of your feat
- `description`: string, any text you want to describe your feat

### `languages`

**Description:** 
A list of languages your character has.

**Examples**: *not a comprehensive list*
- "Common"
- "Elvish"
- "Draconic"

### `hp-current`

**Description:** 
An int representing your current HP. When setting up your character, just make this the same as your max HP.

### `hp-max`

**Description:** 
An int representing your maximum HP. You will need to adjust this every level, as a dice roll is used to determine your gained HP as you level up.

### `hp-temp`

**Description:** 
An int representing your temporary HP. Set this up as zero. When you add damage to your character, if you have temporary HP, it will be decreased before your current HP

### `speed`

**Description:** 
The speed your character can travel on flat land. Should be an int determined by your race.

### `abilities`

**Description:** 
A list of your abilities. The base stats should never change, the modifiers will be calculated for you at runtime.

**Fields:**
- `name`: string
**Allowed Values:** *not case sensitive*
    - "Strength"
    - "Dexterity"
    - "Constitution"
    - "Intelligence"
    - "Wisdom"
    - "Charisma"

- `base`: int from 1 to 20
- `saving-throws-proficient`: bool (true or false)

### `skills`

**Description:** 
A list of your skills. Each skill is tied to one of the six abilities listed above.

**Fields:**
- `ability`: string, ability that skill is a member of
**Allowed Values:** *not case sensitive*
    - "Strength"
    - "Dexterity"
    - "Constitution"
    - "Intelligence"
    - "Wisdom"
    - "Charisma"
- `name`: string
**Allowed Values:** *not case sensitive*
    - "Acrobatics"
    - "Animal Handling"
    - "Arcana"
    - "Athletics"
    - "Deception"
    - "History"
    - "Insight"
    - "Intimidation"
    - "Investigation"
    - "Medicine"
    - "Nature"
    - "Perception"
    - "Performance"
    - "Persuasion"
    - "Religion"
    - "Sleight of Hand"
    - "Stealth"
    - "Survival"
- `proficient`: bool (true or false)
---

## Character Spells

As a note, if your character class has no spell save DC, then the spell commands for the CLI/TUI will not be available to you. 

### `spells`

**Description:** 
A list of your known spells.

**Fields:**
- `slot-level`:  int from 0 to 9
- `ritual`: bool (true or false)
- `name`: string, name of spell. I'd recommend using whatever name is used to query the api when searching, [found here](https://www.dnd5eapi.co/) 

### `spell-slots`

**Description:** 
A List of spell slots available to your character at their current level.

**Fields:**
- `level`:  int from 0 to 9
- `available`: int, currently available spell slots for that level. When setting this up, just make it equal to the maximum
- `maximum`: int, maximum spell slots for this level

---

## Equipment

### `weapons`

**Description:** 
A List of all weapons in your characters possession, not just equipped ones

**Fields:**
- `name`: string, name of weapon
- `damage`: string, die value representing the weapon damage (don't worry about bonuses, those are handled at runtime)
- `proficient`: bool (true or false)
- `ranged`: bool (true or false). If it isn't clear, melee weapons are false, all others are true. If it is a throwable weapon, still mark this as false and just make sure to put in the normal/long range
- `range`: 
    - normal-range: int 
    - long-range: int 
- `type`: string 
**Allowed Values:** *not case sensitive*
    - "Slashing"
    - "Piercing"
    - "Bludgeoning"
- `properties`: string, a list of weapon properties 
**Allowed Values:** 
    - "ammunition"
    - "finesse"
    - "heavy"
    - "light"
    - "loading"
    - "reach"
    - "special"
    - "thrown"
    - "two-handed"
    - "versatile"
    - "monk"

### `primary-equipped`

**Description:** 
Your primary equipped item, should be a string that matches a weapon or your shield from worn equipment. Make sure they are spelled the same.

### `secondary-equipped`

**Description:** 
Your secondary equipped item, should be a string that matches a weapon or your shield from worn equipment. Make sure they are spelled the same. It should be different from your primary weapon, you can not equip the same primary and secondary weapons unless you have multiple of that weapon listed under `weapons`

### `worn-equipment`

**Description:** 
Equipment that is specifically worn by your character

**Fields:**
- `head`: string
- `amulet`: string
- `cloak`: string
- `hands-arms`: string
- `ring`: string
- `ring2`: string
- `belt`: string
- `boots`: string
- `shield`: string
- `armor`: 
**Fields:**
    - `name`: string
    - `proficient`: bool (true or false)
    - `class`: int, represents part of what makes up your AC
    - `type`: string, represents the durability/weight of your armor
**Allowed Values:** 
        - "light"
        - "medium"
        - "heavy"

### `backpack`

**Description:** 
A list of backpack items. You can be as granular as you'd like.

**Fields:**
- `name`: string
- `quantity`: int
