# README 
---
## About
dndgo is a terminal user interface and a command line application for Dungeons and Dragons built to make managing characters easier, and searching dnd info with the [5e-srd-api](https://5e-bits.github.io/docs/). 

Configure your character, and this application will handle tracking your mods, bonuses, and class details. If your class has an ability that modifies your characters stats, it is handled for you by dndgo. You can also use the application to track your characters health, spell slots, and class specific slots (like bardic inspiration or rage). It will only apply class bonuses available to your class at your current level. This information is formatted for you in markdown, or rendered for you in a stateful terminal interface using [bubbletea](https://github.com/charmbracelet/bubbletea).

---

## Tui 

### Character Management
Currently Character management is only implemented for basic info, spells, and equipment. Actions can be preformed on the character like damage, recovering health, adding items, using spell slots and more. More detailed information on these commands can be found in the docs/tui directory.

*example*

https://github.com/user-attachments/assets/c6c40922-ce7e-472e-8589-9180bf1d3421

### Search

Use tab and shift tab to navigate tabs or /s, /m, /e, /f in the search field to switch directly to a tab. 

You can enter a search term after the command to search in a specific tab without navigating to it first. 

ctrl+s will toggle the search field

*example*

https://github.com/user-attachments/assets/d64a8fd3-eece-4444-8ade-f5feaaac6082

### Create Character

If manually modifying json sounds exhausting to you, you can setup your character with the tui. 

*example*

Example pending, for now you'll have to take my word for it or try it yourself!

## CLI

You can use the CLI to directly make changes to the characters state (like health, items, spell slots, etc) as well as rendering out a markdown file with your character information, like in the example below

<img width="750" height="1202" alt="Screenshot 2025-11-26 at 10 31 32 PM" src="https://github.com/user-attachments/assets/afc1c638-c841-4eb8-919e-a4fa35cf3924" />
<img width="750" height="1307" alt="Screenshot 2025-11-26 at 10 32 07 PM" src="https://github.com/user-attachments/assets/64da0036-cfdb-4aba-8052-ad3c904efc82" />



## Commands

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
-  -b, --backpack string    Item to add to backpack (use -q to specify quantity)
-  -e, --equipment string   Kind of quipment to add 'armor, ring, etc'
-  -l, --language string    Language to add
-  -n, --name string        Name of equipment to add
-  -q, --quantity int       Modify quantity of something
-  -x, --spell string       Add spell to list of character spells
-  -s, --spell-slots int    Increase spell-slot max capacity by level
-  -t, --temp-hp int        Add temporary hp
-  -w, --weapon string      Weapon to add
  
*examples*

`dndgo ctr add -b "potion of greater healing" -q 1` - Add one potion of greater healing to your inventory

`dndgo ctr add -t 5` Add 5 temporary HP

---

`ctr remove`

*documentation needed*

---

`ctr use`

**Use Flags**
-  -b, --backpack string       Use item from backpack
-  -c, --class-tokens string   Use class-tokens by token name (default "any")
-  -q, --quantity int          Modify quantity of something
-  -s, --spell-slots int       Use spell-slot by level

*examples*

`dndgo ctr use -b Gold -q 10` -  Use 10 Gold

`dndgo ctr use -c any` -  Use 1 class token for a class that only uses one token

`dndgo ctr use -c divine-sense -q 2` -  Use 2 divine sense class tokens

`dndgo ctr use -s 2` -  Use a level 2 spell slot

---

`ctr recover`

**Recover Flags**
-  -a, --all                   Recover all health, slots, and tokens
-  -c, --class-tokens string   Recover class-tokens by token name (default "all"), if no quantity is specified, a full class token recovery is assumed
-  -p, --hitpoints int         Recover hitpoints
-  -q, --quantity int          Recover the quantity of something
-  -s, --spell-slots int       Recover spell-slot by level, if no quantity is specified, a full spell slot recovery is assumed for that level

*examples*

`dndgo ctr recover -a` - Full recovery equivalent to a long rest.

`dndgo ctr recover -p 10` - Recover 10 hp

`dndgo ctr recover -s 1 -q 2` - Recover 2 level 1 spell slots 

`dndgo ctr recover -s 2` - Recover all level 2 spell slots

`dndgo ctr recover -c any -q 2` - Recover 2 class tokens where one only token is available for that class

`dndgo ctr recover -c divine-sense -q 2` -  Recover 2 divine sense class tokens

---

`ctr get`

**Get Flags**
-  -p, --path string   Get config or markdown path
-  -t, --tokens        Get class tokens

*examples*

`dndgo ctr get -p` - Get config Path

`dndgo ctr get -t` - Get all class tokens available to character's class

---

`ctr equip`

*documentation needed*

`ctr unequip`

*documentation needed*

---

## Setting up Character Config
- run "dndgo ctr init -c <class-name> -n <character-name>
- modify character details in the ~/.config/dndgo directory
- if you want your character markdown to be saved to a different file than your config files, you can specify a path in the character.json file. If you do this, make sure that you only specify the path from the home directory (e.g. "dnd/mdfiles", not "~/dnd/mdfiles") or you will create a ~ dir in the home directory. 
- run "dndgo ctr update" to generate your markdown file
- more details on how to set up character and class can be found in the docs directory

Search Example

https://github.com/user-attachments/assets/0770df8d-d0c6-4dd4-ab1c-8e7098882262

Character Markdown Example

https://github.com/user-attachments/assets/85c1b130-808e-4889-aa60-fd8d45ae1190
