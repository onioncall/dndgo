# Tui Commands
---

## Search
- tab to switch tabs to the right, shift+tab to switch to tabs on the left
- ctrl+s to show or hide search bar

When you are on the tab you'd like to search for, just enter your query. For example,

*Spell tab*
- `acid arrow`

*Monster tab*
- `adult black dragon`

*Equipment tab*
- `dagger`

You can also switch tabs by entering in `/` followed by the first letter of the tab you'd like to switch to
If you'd like you can switch and query in the same command. For example,

*Spell tab*
- `/m adult black dragon`

This searches for adult black dragon as a monster even when you aren't on the monster tab

## Manage
Character Management is the central feature of this application. To Navigate,

- tab to switch tabs to the right, shift+tab to switch to tabs on the left
- ctrl+s to show or hide cmd bar, or clear an error

### Basic Info
Commands available to basic info 

- *damage (int, damage amount)* example, `damage 3` removes three hp from your character, if your character has temp hp, that is removed first
- *recover (optional int, recover amount)* 
    - example: `recover 3` recovers three hp for your character.
    - details: if no argument is specified, we perform the equivilent of a long rest on your character.
- *temp (int, temp hp amount)* example, `temp 5` adds five temporary hp

### Spells
Commands available to spells

- *use-slot (int, level)* example, `use-slot 1` uses a single level one spell slot
- *recover-slot (int, level)* example, `recover-slot 1` recovers a single level one spell slot

### Equipment
Commands available to equipment

- *add-equipment (string, worn equipment type)/(string, equipment name)* example, `add-equipment amulet/clockwork amulet` 
- *equip (string, weapon or shield name)/(optional string, primary or secondary)*
    - example:  `equip dagger` or `equip rapier primary` 
    - details: if you don't specify primary or secondary, it will equip which ever is open (prioritizing primary). if neither are available and primary/secondary is not specified, it will replace the primary. 

- *unequip (string, primary/secondary/weapons name/shield/name)* 
    - example:  `unequip dagger` or `unequip secondary` 

- *add-item (string, item name)/(optional int, quantity)*
    - example:  `add-item gold` or `add-item gold/5` 
    - details: if you don't specify a quantity, only one is added. If an item with that name (not case sensitive) is found in your inventory, we'll add to the quantity rather than adding a new item

- *remove-item (string, item name)/(optional int, quantity)*
    - example: `remove-item gold` or `remove-item gold/5`

### Class
Commands available to class

- *use-token (optional string, token name)/(optional int, quantity)*
    - example:  `use-token` or `use-token /2` or `use-token divine-sense` or `use-token divine-sense/2`
    - details: if you don't specify a quantity, only one is used. A token name is only required if there are multiple tokens available to that class, otherwise any (or an empty) string will do

- *recover-token (optional string, token name)/(optional int, quantity)*
    - example:  `recover-token` or `recover-token /2` or `recover-token divine-sense` or `recover-token divine-sense/2`
    - details: if you don't specify a quantity, a full token recovery is performed. A token name is only required if there are multiple tokens available to that class, otherwise any (or an empty) string will do

