# Flags 
---
to return a list of all spells/monsters/etc, pass l or list as an argument

-s for spells (ex. dndgo -s "acid arrow" -or- dndgo -s acid-arrow)

-m for monsters (ex. dndgo -m "adult black dragon" -or- dndgo -s adult-black-dragon)

ctr for character actions

dndgo ctr add -l "Elvish" (Add Language)
dndgo ctr add -b "Gold 17" (Add 17 gold)
dndgo ctr remove -b "Gold 2" (remove 2 gold)
dndgo ctr add -e "cloak: Cloak of Billowing" (adds Cloak of Billowing as a cloak to body equipment)
dndgo ctr add -lvl (increase level)
dndgo ctr add -p 12 (heal 12 hp)
dndgo ctr remove -p 12 (damage 12 hp)
dndgo ctr update (update markdown from existing character.json)

character flag hints 
    -p Hit Points (int)
    -l Language (string)
    -b Backpack (Name Quantity)
    -e Equipment (type: Name)
    -lvl Level

-tui (TODO: Implement alternative method of searching and viewing returned data with bubbletea)

upcoming TODOs: 
    - implement add spell and add weapon by using the 5e API. in theory, like dndgo ctr add -s "acid-arrow" or dndgo ctr add -w "Rapier" 

When in doubt, don't use caps. I haven't done the conversion everywhere to lower case everywhere, so it may cause problems
