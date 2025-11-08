# Bard Setup

### `college`
**Description:**
College is the subclass for Bard. It can be any string and is not case sensitive

### `bardic-inspiration`
**Description:**
Bards have a tracked number of Bardic Inspiration dice they can use before a rest.

**Fields:**
- `available`: int, current charges/tokens
- `maximum`: int, maximum charges/tokens (equal to your Charisma modifier)

### `expertise`
**Description:**
At 3rd level, choose two of your skill proficiencies. Your proficiency bonus is doubled for any ability check you make that uses either of the chosen proficiencies. At 10th level, you can choose another two skill proficiencies to gain this benefit.

**Allowed Values:** Any skill you are proficient in
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
