# DNDGO General Info
---

This page is intended to outline what that application specifically does. Implementation details for these commands are more explicitly covered in the tui and cli documentation.

Currently, we only support 5th Edition, and the application will only handle cases outlined in the players handbook, auxilary sources are not issues, but we can't guarantee implementation for features found in sources outside the players handbook.

## Search

The initial idea for this project was was a search client. Currently we support querying the [5e-srd-api](https://5e-bits.github.io/docs/) api for data on Spells, Monsters, Equipment, and Features. This is supported by both the CLI and TUI.

## Character Management

The project scope is now largely character management. The goal of character management is for dndgo to handle the calculation of all mods and bonuses that are able to be derived from the characters class and configuration.

dndgo will handle calculating-

- Proficiency: We derive your proficiency bonus from your base level
- Ability Mods: We derive your ability mods from you ability base (1-20) 
- Skill Mods: We derive your skill mods from your ability mods
- AC: We derive your AC from your armor and your ability mods
- Passive Perception, and Passive Insight: We derive these from your skill/ability mods and proficiency
- Weapon Bonuses: We derive these from your ability mods and weapon properties 
- Hit Dice: We derive these by class
- Spell Save DC: We derive this from mods and proficiency
- Spell Attack Mod: We derive this from your ability mods and proficiency 
- Class Token Maximums: Class tokens like bardic inspiration and rage have their max uses calculated by dndgo

These are the initial basic transformations that apply for all classes. Each individual class further modifies these values based on configuration and level.

Classes have features that modify our characters stats, dndgo will track these on your behalf. For instance, if a class has a feature that adds a proficiency to a skill, that skill will be recalculated with the new proficiency bonus added.

Below are features that classes have that aren't listed under class setup because they require no configuration, but that are happening in the background once a certain class level has been reached. It's recommended to read the class setup doc for your class to understand all the configuration you are able to do on that class.

### Bard 

Level 3
Jack of All Trades: Adds half of their proficiency bonus rounded down to any skill they are not proficient in

## Druid

Level 20
Arch Druid: Unlimited wild shape transformations. We set it to zero, but that's our way of showing that it does not need to be tracked anymore

### Monk

Diamond Soul: Proficiency is automatically applied to all saving throws. This will be rendered in the tui and in the markdown, but does not get saved to the json, so don't be alarmed if you don't see it in the json

### Warlock

Eldritch Invocations: These can be added for the warlock class, but full disclosure they don't do anything.

---
### Class Modifications 

## Work Arounds 

There are many, many things you need to be able to add to your character, and currently we don't have commands implemented for all of them. For the ones that don't have commands, you must directly modify the json. 

---
### Ability Score Improvement Items

We represent Ability Score improvement items as an array of objects. 

  ```
  "ability-score-improvement": [
    {
      "ability": "dexterity",
      "bonus": 2
    }
  ]
  ```
When you add an ability score improvement item, you get to select adding two items with a bonus of one, or one item with a bonus of two. If you choose the same ability for multiple Ability Score Improvements (say, two dexterity at level 4, and two more dexterity at level 8), add a second "dexterity" element with a bonus of two.

---
### Sub Class Details

Fully implementing Sub Class is a massive amount of text to add, and we haven't started on implementing them. The current recommended way of adding sub class information to your character, is by adding them as class features. 

```
  "other-features": [
    {
      "name": "Bardic Inspiration",
      "level": 1,
      "details": "As a bonus action, give an ally a d6 (scales with level) they can add to one attack roll, ability check, or saving throw within the next 10 minutes. You have a number of uses equal to your Charisma modifier (minimum 1), and regain all uses on a long rest.\n"
    },
    {
      "name": "Example Sub Class Item Title",
      "level": 1,
      "details": "Example Sub Class Item Details..."
    }
]
```

Just add a feature element to the array with the title of your sub class item/feature/detail to add, and the content in the "details" section

If there are any sub class features that modify your mods or bonuses directly, *THEY ARE NOT* accounted for in the mods/bonuses that you see in the markdown or in the tui. You will need to add them separately. This is different than core class features, which will be added to your mods and bonuses when you reach a level that makes that feature available to your character.

### Feats

While we do have feats as a field on our character struct, if you are predominantly using the TUI, we recommend adding the feat to your class details the same way we recommend adding sub class details.

### Class Detail Modifications

If you want to rewrite the information that comes by default with your class features, we welcome it! You'll want to modify it the in the class.json file, and you should be good.

---
### Primal Knowledge

Primal knowledge grants Barbarians the ability at level 3 and level 10 to add a proficiency to a skill from their list. Add/Modify this in your class.json file.

```
  "primal-knowledge": [
    "animal-handling"
  ],
```
---
### Expertise

Bards and Rogues get to pick skills they are proficient in, and double the proficiency bonus. To use this feature, add this to your class json with the skills you are choosing

```
  "expertise": [
    "insight",
    "persuasion"
  ]
```

As a note, there are sub-classes and feats that grant expertise, but they are not handled by dndgo and will have to be tracked separately. This is something we will consider implementing in the future. If you'd like this to be added sooner rather than later, raise an issue in the github!

---
### Prepared Spells

This one should get a CLI command pretty soon, but in the meantime... 
Cleric, Druid, Paladin, and Wizard have a list of prepared spells separate from their known spells. Until a command is implemented to set add these, you'll have to modify them in the class.json file.

```
  "prepared-spells": [
    "some spell name"
  ],
```
Add the spells that you prepare to this array, and make sure that they match the spelling and casing of the spells in your "known-spells" array in your character.json file, and you should be good to go.

---
### Oath Spells

Paladins apparently get everything, which includes their own special list of spells. To add and remove, you'll want to add/modify the following in your class.json file.

```
  "oath-spells": [
    "Oath of Devotion"
  ],
```

---
### Fighting Style

Fighters, Paladins, and Rangers get to select a fighting style, this one is pretty simple to add.

```
  "fighting-style": "two-weapon-fighting",
```

Supported fighting styles are 
- archery
- defense
- dueling
- two-weapon-fighting
- great-weapon-fighting
- protection

Though not all classes listed above implement all these styles.

---
### Favored Enemies

Rangers that want to add favored enemies will currently need to add/modify them directly in their class.json

```
  "favored-enemy": [
    "Beasts"
  ],
```
---
