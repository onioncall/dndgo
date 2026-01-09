## CLI Commands
---

These are flags for general application configuration, and are mostly logging related
**Base Flags**
- --clear-log    clear the log file, no shorthand flag
- -h, --help     help for dndgo *This applies for all of the following sub commands, we will not list it every time
- --log string   log output path, use ':stdout' for stdout (default "/Users/onioncall/.local/state/dndgo/dndgo.log"), no shorthand flag
- -v, --version  version for dndgo

---
You can use the CLI to directly make changes to the characters state (like health, items, spell slots, etc) as well as rendering out a markdown file with your character information

### Search
**Search Flags**
 - -s (spells)
 - -m (monsters)
 - -e (equipment)
 - -f (features)

`search`

`search list`

*examples*

`dndgo search -s acid-arrow` - Look up the spell acid arrow

`dndgo search list -s` - Get a list of all spells available to this api

### Character

`ctr`

`ctr init`

**Init Flags**
-  -c, --class string   Name of character class
-  -n, --name string    Name of character

*examples*

`dndgo ctr init -c bard -n Nim` - Create character with a class of bard and a name of Nim

---

`ctr add`

**Add Flags**
-  -a  --ability-improvement    Ability Score Improvement item name, (use -q to specify a quantity)
-  -b, --backpack string        Item to add to backpack (use -q to specify quantity)
-  -e, --equipment string       Kind of equipment to add 'armor, ring, etc'
-  --language string            Name of language to add
-  -n, --name string            Name of equipment to add
-  -q, --quantity int           Modify quantity of something
-  -x, --spell string           Add spell to list of character spells
-  -s, --spell-slots int        Increase spell-slot max capacity by level
-  -t, --temp-hp int            Add temporary hp
-  -w, --weapon string          Weapon to add
  
*examples*

`dndgo ctr add -b "potion of greater healing" -q 1` - Add one potion of greater healing to your inventory

`dndgo ctr add -t 5` - Add 5 temporary HP

---

`ctr remove`

*documentation needed*

---

`ctr use`

**Use Flags**
-  -b, --backpack string       Use item from backpack
-  -t, --class-tokens string   Use class-tokens by token name (default "any")
-  -q, --quantity int          Modify quantity of something
-  -s, --spell-slots int       Use spell-slot by level

*examples*

`dndgo ctr use -b Gold -q 10` - Use 10 Gold

`dndgo ctr use -c any` - Use 1 class token for a class that only uses one token

`dndgo ctr use -c divine-sense -q 2` - Use 2 divine sense class tokens

`dndgo ctr use -s 2` - Use a level 2 spell slot

---

`ctr recover`

**Recover Flags**
-  -a, --all                   Recover all health, slots, and tokens
-  -t, --class-tokens string   Recover class-tokens by token name (default "all"), if no quantity is specified, a full class token recovery is assumed
-  -p, --hitpoints int         Recover hitpoints
-  -q, --quantity int          Recover the quantity of something
-  -s, --spell-slots int       Recover spell-slot by level, if no quantity is specified, a full spell slot recovery is assumed for that level

*examples*

`dndgo ctr recover -a` - Full recovery equivalent to a long rest.

`dndgo ctr recover -p 10` - Recover 10 hp

`dndgo ctr recover -s 1 -q 2` - Recover 2 level 1 spell slots 

`dndgo ctr recover -s 2` - Recover all level 2 spell slots

`dndgo ctr recover -c any -q 2` - Recover 2 class tokens where only one token is available for that class

`dndgo ctr recover -c divine-sense -q 2` -  Recover 2 divine sense class tokens

---

`ctr get`

**Get Flags**
-  -p, --path string     Get config or markdown path
-  -t, --tokens          Get class tokens
-  -n, --character-names Get character names

*examples*

`dndgo ctr get -p` - Get config Path

`dndgo ctr get -t` - Get all class tokens available to character's class

---

`ctr equip`

*documentation needed*

`ctr unequip`

*documentation needed*

---

`ctr modify`

**Modify Flags**
-a, --ability-improvement string   Ability Score Improvement item name, (use -q to specify a quantity)

*examples*

`dndgo ctr modify -a dexterity -q 4`

--- 

`ctr import`, `ctr export`

**Import Flags**
- -n, --character-name string   Name of character, only used when importing Class data
- -c, --class                   Import Class file (default: Character)
- -f, --file string             Relative path to json file

*examples*

`dndgo ctr import -f nim.json` - Imports character from json file 'nim.json' in your current directory

`dndgo ctr import -f nim-class.json -n Nim -c` - Imports class file 'nim-class.json' for character name 'Nim'

---

`ctr delete`

**Delete Flags**
- -n, --character-name string   Name of character to delete

*examples*

`dndgo ctr delete -n Nim`

---

`ctr update`

**Update Flags**
- -d, --default-character-name string   Name of character to make default

*examples*

`dndgo ctr update -d Nim`

---

`ctr class`

**Expertise Flags**
- -e, --expertise string        name of skill to add to expertise (remove does not apply)
- -f, --fighting-style string   name of fighting style to assign (remove does not apply)
- -p, --prepared-spell string   name of spell to prepare
- -r, --remove                  remove instead of add one of these things

*examples*

`dndgo ctr class -e perception` - adds perception to your expertise skills

`dndgo ctr class -p "Healing Word" -r`  - removes healing word from prepared spells

---
