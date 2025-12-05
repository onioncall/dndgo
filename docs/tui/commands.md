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

You can also swith tabs by entering in `/` followed by the first letter of the tab you'd like to switch to
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
    - details: if no argument is specified, we perform the equivilant of a long rest on your character.
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
    - details: if you don't specify primary or secondary, it will equip which ever is open (prioritizing primary). if neither are avialable and primary/secondary is not specified, it will replace the primary. 

- *unequip (string, primary/secondary/weapons name/shield/name)* 
    - example:  `unequip dagger` or `equip secondary` 

- *add-item (string, item name)/(optional int, quantity)*
    - example:  `add-item gold` or `add-item gold/5` 
    - details: if you don't specify a quantity, only one is added. If an item with that name (not case sensitive) is found in your inventory, we'll add to the quantity rather than adding a new item

- *remove-item (string, item name)/(optional int, quantity)*
    - example: `remove-item gold` or `remove-item gold/5`
