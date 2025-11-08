# Monk Setup

### `mosaic-tradition`
**Description:**
Monastic Tradition is the subclass for Monk. It can be any string and is not case sensitive

### `ki-points`
**Description:**
Starting at 2nd level, your training allows you to harness the mystic energy of ki. You have a number of ki points equal to your monk level. You can spend these points to fuel various ki features.

**Fields:**
- `available`: int, current ki points
- `maximum`: int, maximum ki points (equal to your monk level, and set at runtime)
- `spell-save-dc`: int, the DC for ki abilities that require a saving throw (8 + proficiency bonus + Wisdom modifier)
