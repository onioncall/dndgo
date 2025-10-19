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

When in doubt, don't use caps. I haven't done the conversion everywhere to lower case everywhere, so it may cause problems

## Setting up Character Config
- run "dndgo ctr init -c <class-name> -n <character-name>
- modify character details in the ~/.config/dndgo directory
- if you want your character markdown to be saved to a different file than your config files, you can specify a path in the character.json file. If you do this, make sure that you only specify the path from the home directory (e.g. "dnd/mdfiles", not "~/dnd/mdfiles") or you will create a ~ dir in the home directory. 
- run "dndgo ctr update" to generate your markdown file

Search Example

https://github.com/user-attachments/assets/0770df8d-d0c6-4dd4-ab1c-8e7098882262

Character Markdown Example

https://github.com/user-attachments/assets/85c1b130-808e-4889-aa60-fd8d45ae1190

