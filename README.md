# README 
---

<p align="center">
  <img width="800" height="450" alt="dndgo-logo-highres" src="https://github.com/user-attachments/assets/b4bfb7b8-16a5-40e3-b285-4b414a249ba7" />
 </p>

## About
dndgo is a terminal user interface and a command line application for Dungeons and Dragons built to make managing characters easier, and searching dnd info with the [5e-srd-api](https://5e-bits.github.io/docs/).

Configure your character, and this application will handle tracking your mods, bonuses, and class details. If your class has an ability that modifies your characters stats, it is handled for you by dndgo. You can also use the application to track your characters health, spell slots, and class specific slots (like bardic inspiration or rage). It will only apply class bonuses available to your class at your current level. This information is formatted for you in markdown, or rendered for you in a stateful terminal interface using [bubbletea](https://github.com/charmbracelet/bubbletea).

---

## Installation

If you have golang installed

`go install github.com/onioncall/dndgo@latest`

or to build from source (v 1.21 or later)

```bash
git clone https://github.com/onioncall/dndgo.git
cd dndgo
go build -o dndgo
sudo mv dndgo /usr/local/bin/
```

### Linux

*x86_64 (Intel/AMD 64-bit)*
```bash
curl -LO https://github.com/onioncall/dndgo/releases/download/v0.2.0/dndgo_Linux_x86_64.tar.gz
tar -xzf dndgo_Linux_x86_64.tar.gz
sudo mv dndgo /usr/local/bin/
```

### Mac OS

*Apple Silicon (M1/M2/M3)*
```bash
curl -LO https://github.com/onioncall/dndgo/releases/download/v0.2.0/dndgo_Darwin_arm64.tar.gz
tar -xzf dndgo_Darwin_arm64.tar.gz
sudo mv dndgo /usr/local/bin/
# Don't do the following step unless you trust this applciation, I don't pay apple $99.99 a year for a developer account to sign these binaries
xattr -d com.apple.quarantine /usr/local/bin/dndgo
```

### Windows

Download [dndgo_Windows_x86_64.zip](https://github.com/onioncall/dndgo/releases/download/v0.2.0/dndgo_Windows_x86_64.zip)
Extract the zip file
Move `dndgo.exe` to a folder in your PATH (e.g., `C:\Windows\System32`)

Once you have installed it, run `dndgo -v` to verify your installation version.

Hopefully we can get this in package managers if we ever get more than half a dozen people using this. 

---

## Tui 

### Character Management
Character management has tabs for basic info, spells, equipment, class info, and help for how to use the tui. Actions can be performed on the character like damage, recovering health, adding items, using spell slots and more. Detailed information on these commands can be found in the docs/tui directory.

*example*

https://github.com/user-attachments/assets/ff6a58e3-74a2-4461-a4f6-ebf140e1ac70

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

---

## Setting up Character Config
We recommend using the tui for this, but if you'd prefer, it can technically be done via cli.
- run `dndgo ctr init -c <class-name> -n <character-name>`
- export and modify your character details `dndgo ctr export -f <file name.json> -n <character name>`
- import your modified character `dndgo ctr import -f <file name.json>`
- if you want your character markdown to be saved to a different file than your config files, you can specify a path in the character.json file. If you do this, make sure that you only specify the path from the home directory (e.g. "dnd/mdfiles", not "~/dnd/mdfiles") or you will create a ~ dir in the home directory. 
- run `dndgo ctr update` or `dndgo` to generate your markdown file
- more details on how to set up character and class can be found in the docs directory

---

## Contributions

We'd love to have you contribute, either by creating a discussion and reporting issues/requesting features, or by picking up existing issues and submitting a PR! If you're interested, please read through the contribution.md file in our docs/ directory.

---

<small>*Special Thank you to Matt Evans for the dndgo logo, and to Renée French for the inspiration</small>
