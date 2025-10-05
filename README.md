# README 
---

## Search CLI usage
to return a list of all spells/monsters/etc, pass l or list as an argument

-s for spells (ex. dndgo search -s "acid arrow" -or- dndgo -s acid-arrow)

-m for monsters (ex. dndgo search -m "adult black dragon" -or- dndgo -s adult-black-dragon)

-e for equipment (ex. dndgo search -e "soap" -or- dndgo -e soap)

## Character CLI usage
dndgo ctr add -l "Elvish" (add Language)
dndgo ctr add -b Gold -q 17 (add 17 gold)
dndgo ctr remove -b Gold -q 2 (remove 2 gold)
dndgo ctr add -e cloak -n "Cloak of Billowing" (adds Cloak of Billowing as a cloak to body equipment)
dndgo ctr add -lvl (increase level)
dndgo ctr recover -p 12 (heal 12 hp)
dndgo ctr use -s 1 (uses level one spell slot)
dndgo ctr remove -p 12 (damage 12 hp)
dndgo ctr update (update markdown from existing character.json)

character flag hints 
    -p Hit Points (int)
    -l Language (string)
    -b Backpack (Name Quantity)
    -e Equipment (type: Name)
    -lvl Level

When in doubt, don't use caps. I haven't done the conversion everywhere to lower case everywhere, so it may cause problems

## Setting up Character Config
- copy character file from the "default-json-configs" directory
- modify it to your character
- find your class and copy the file from "default-json-configs" directory
- modify it to your character
- rename files "character.json" and "class.json" respectively
- move them to your ~/.config/dndgo directory
- run "dndgo ctr update" to generate your markdown file

Search Example

https://github.com/user-attachments/assets/d7c95bd0-3388-4514-ae1c-aa1fadc10b02

Character Markdown Example

https://github.com/user-attachments/assets/730cd6f4-2f8b-4afa-a784-067a18f9bf3f
