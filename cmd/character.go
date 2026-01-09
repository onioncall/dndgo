package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/onioncall/dndgo/character-management/handlers"
	"github.com/onioncall/dndgo/logger"

	"github.com/spf13/cobra"
)

var buildMd bool

var (
	characterCmd = &cobra.Command{
		Use:     "character",
		Short:   "Manage character information",
		Aliases: []string{"ctr"},
	}

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add character attributes",
		Run: func(cmd *cobra.Command, args []string) {
			a, _ := cmd.Flags().GetString("ability-improvement")
			l, _ := cmd.Flags().GetString("language")
			bp, _ := cmd.Flags().GetString("backpack")
			il, _ := cmd.Flags().GetBool("level")
			e, _ := cmd.Flags().GetString("equipment")
			s, _ := cmd.Flags().GetString("spell")
			ss, _ := cmd.Flags().GetInt("spell-slots")
			q, _ := cmd.Flags().GetInt("quantity")
			t, _ := cmd.Flags().GetInt("temp-hp")
			n, _ := cmd.Flags().GetString("name")
			sc, _ := cmd.Flags().GetString("sub-class")
			ct, _ := cmd.Flags().GetString("class-type")

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to load character data")
				return
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to handle character")
				return
			}

			if l != "" {
				c.AddLanguage(l)
			}
			if e != "" {
				if n == "" {
					logger.PrintError("Name of equipment can not be left empty")
					return
				}

				c.AddEquipment(e, n)
			}
			if bp != "" {
				if q <= 0 {
					logger.PrintError("Must pass a positive quantity to add")
					return
				}

				c.AddItemToPack(bp, q)
			}
			if il {
				c.AddLevel()
			}
			if ss > 0 {
				// add spell slot for level
			}
			if s != "" {
				err = handlers.AddSpell(c, s)
			}
			if t != 0 {
				c.AddTempHp(t)
			}
			if sc != "" {
				err = c.AddSubClass(ct, sc)
				if err != nil {
					logger.Error(err)
					logger.PrintError("Failed to add subclass")
					return
				}
			}
			if a != "" {
				err = c.AddAbilityScoreImprovementItem(q, a)
				if err != nil {
					logger.PrintError(err.Error())
					return
				}
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			for _, class := range c.Classes {
				err = handlers.SaveClass(class)
				if err != nil {
					logger.Error(err)
					logger.PrintError(fmt.Sprintf("Failed to save data for class '%s'", class.GetClassType()))
					return
				}
			}

			if buildMd {
				err = handlers.BuildCharacterMarkdown(*c)
				if err != nil {
					logger.Error(err)
					logger.PrintError("failed to generate markdown file")
					return
				}
			}

			logger.PrintSuccess("Character Update Successful")
		},
	}

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get character specific data",
		Run: func(cmd *cobra.Command, args []string) {
			p, _ := cmd.Flags().GetString("path")
			t, _ := cmd.Flags().GetBool("tokens")
			n, _ := cmd.Flags().GetBool("character-names")

			if p != "" {
				var path string
				var err error

				if p == "config" {
					path, err = handlers.GetConfigPath()
					if err != nil {
						logger.Error(err)
						logger.PrintError("Failed to get config path")
						return
					}
				} else if p == "markdown" {
					c, err := handlers.LoadCharacter()
					if err != nil {
						logger.Error(err)
						logger.PrintError("Failed to load character")
						return
					}

					path = c.Path
					if path == "" {
						path, err = handlers.GetConfigPath()
					}
				} else {
					logger.PrintError("Path option not found")
				}

				fmt.Println(path)
			}

			if t {
				c, err := handlers.LoadCharacter()
				if err != nil {
					logger.Error(err)
					logger.PrintError("Failed to load character data")
					return
				}

				if c.Classes == nil {
					logger.PrintError("Class not properly configured")
					return
				}

				tokenMap := c.GetTokenNames()
				for classType, tokens := range tokenMap {
					fmt.Printf("Class: %s\n", classType)
					for _, token := range tokens {
						fmt.Printf("-> %s\n", token)
					}
				}
			}

			if n {
				names, err := handlers.GetCharacterNames()
				if err != nil {
					logger.Error(err)
					logger.PrintError("Failed to get character names")
					return
				}

				if len(names) == 0 {
					fmt.Println("No characters found")
					return
				}

				for _, name := range names {
					fmt.Println(name)
				}
			}
		},
	}

	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete character/class data",
		Run: func(cmd *cobra.Command, args []string) {
			c, _ := cmd.Flags().GetBool("class")
			n, _ := cmd.Flags().GetString("name")

			if c {
				// Delete class by type and character name
				logger.PrintError("Feature to delete single class is not yet supported")
			} else {
				err := handlers.DeleteCharacter(n)
				if err != nil {
					logger.Error(err)
					logger.PrintError("Failed to delete character")
					return
				}

				logger.PrintSuccess("Deleted character")
			}
		},
	}

	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove character attributes",
		Run: func(cmd *cobra.Command, args []string) {
			hp, _ := cmd.Flags().GetInt("hitpoints")
			u, _ := cmd.Flags().GetInt("use-slot")

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to load character")
				return
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to process character")
				return
			}

			if hp > 0 {
				c.DamageCharacter(hp)
			} else if u > 0 {
				c.UseSpellSlot(u)
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			if buildMd {
				err = handlers.BuildCharacterMarkdown(*c)
				if err != nil {
					logger.Error(err)
					logger.PrintError("failed to generate markdown file")
					return
				}
			}

			logger.PrintSuccess("Character Update Successful")
		},
	}

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update character attributes",
		Run: func(cmd *cobra.Command, args []string) {
			d, _ := cmd.Flags().GetString("default-character-name")
			if d != "" {
				err := handlers.SetDefaultCharacter(d)
				if err != nil {
					logger.Error(err)
					logger.PrintError("Failed to update default character")
					return
				}
			}

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to process character")
				return
			}

			if buildMd {
				err = handlers.BuildCharacterMarkdown(*c)
				if err != nil {
					logger.Error(err)
					logger.PrintError("failed to generate markdown file")
					return
				}
			}

			logger.PrintSuccess("Character Update Successful")
		},
	}

	useCmd = &cobra.Command{
		Use:   "use",
		Short: "Use character items/spell slots",
		Run: func(cmd *cobra.Command, args []string) {
			s, _ := cmd.Flags().GetInt("spell-slots")
			bp, _ := cmd.Flags().GetString("backpack")
			q, _ := cmd.Flags().GetInt("quantity")
			t, _ := cmd.Flags().GetString("class-tokens")
			ct, _ := cmd.Flags().GetString("class-type")

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to process character")
				return
			}

			if bp != "" {
				if q <= 0 {
					logger.PrintError("Must pass a positive quantity to use")
					return
				}

				err = c.RemoveItemFromPack(bp, q)
				logger.Error(err)
				logger.PrintError("Failed to remove item from pack")
				return
			} else if s > 0 {
				c.UseSpellSlot(s)
			} else if t != "" {
				q = max(q, 1) // If q isn't provided with a valid value, we use one by default
				c.UseClassTokens(t, ct, q)
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			for _, class := range c.Classes {
				err = handlers.SaveClass(class)
				if err != nil {
					logger.Error(err)
					logger.PrintError(fmt.Sprintf("Failed to save data for class '%s'", class.GetClassType()))
					return
				}
			}

			if buildMd {
				err = handlers.BuildCharacterMarkdown(*c)
				if err != nil {
					logger.Error(err)
					logger.PrintError("failed to generate markdown file")
					return
				}
			}

			logger.PrintSuccess("Character Update Successful")
		},
	}

	recoverCmd = &cobra.Command{
		Use:   "recover",
		Short: "Recover health and spell slot usage",
		Run: func(cmd *cobra.Command, args []string) {
			a, _ := cmd.Flags().GetBool("all")
			ss, _ := cmd.Flags().GetInt("spell-slots")
			hp, _ := cmd.Flags().GetInt("hitpoints")
			t, _ := cmd.Flags().GetString("class-tokens")
			ct, _ := cmd.Flags().GetString("class-type")
			q, _ := cmd.Flags().GetInt("quantity")

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to process character")
				return
			}

			if a {
				c.Recover()
			} else if ss > 0 {
				c.RecoverSpellSlots(ss, q)
			} else if hp > 0 {
				c.HealCharacter(hp)
			} else if t != "" {
				c.RecoverClassTokens(t, ct, q)
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			for _, class := range c.Classes {
				err = handlers.SaveClass(class)
				if err != nil {
					logger.Error(err)
					logger.PrintError(fmt.Sprintf("Failed to save data for class '%s'", class.GetClassType()))
					return
				}
			}

			if buildMd {
				err = handlers.BuildCharacterMarkdown(*c)
				if err != nil {
					logger.Error(err)
					logger.PrintError("failed to generate markdown file")
					return
				}
			}

			logger.PrintSuccess("Character Update Successful")
		},
	}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes a new character on this machine",
		Run: func(cmd *cobra.Command, args []string) {
			c, _ := cmd.Flags().GetStringSlice("class")
			n, _ := cmd.Flags().GetString("name")

			character, err := handlers.LoadCharacterTemplate(n, c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to load character template")
				return
			}

			err = handlers.CreateCharacter(character)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to create character")
				return
			}

			logger.PrintError("Character Creation Successful")
		},
	}

	equipCmd = &cobra.Command{
		Use:   "equip",
		Short: "Equips a weapon or shield",
		Run: func(cmd *cobra.Command, args []string) {
			p, _ := cmd.Flags().GetString("primary")
			s, _ := cmd.Flags().GetString("secondary")

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to process character")
				return
			}

			if p != "" {
				err = c.Equip(true, p)
				if err != nil {
				}
			}
			if s != "" {
				err = c.Equip(false, s)
				if err != nil {
				}
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			if buildMd {
				err = handlers.BuildCharacterMarkdown(*c)
				if err != nil {
					logger.Error(err)
					logger.PrintError("failed to generate markdown file")
					return
				}
			}

			logger.PrintSuccess("Character Update Successful")
		},
	}

	unequipCmd = &cobra.Command{
		Use:   "unequip",
		Short: "Unequips a weapon or shield",
		Run: func(cmd *cobra.Command, args []string) {
			p, _ := cmd.Flags().GetBool("primary")
			s, _ := cmd.Flags().GetBool("secondary")

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to load character data")
				return
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to process character")
				return
			}

			if p == true {
				c.Unequip(true)
			}
			if s == true {
				c.Unequip(false)
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			if buildMd {
				err = handlers.BuildCharacterMarkdown(*c)
				if err != nil {
					logger.Error(err)
					logger.PrintError("failed to generate markdown file")
					return
				}
			}

			logger.PrintError("Character Update Successful")
		},
	}

	modifyCmd = &cobra.Command{
		Use:   "modify",
		Short: "modify character attributes",
		Run: func(cmd *cobra.Command, args []string) {
			a, _ := cmd.Flags().GetString("ability-improvement")
			q, _ := cmd.Flags().GetInt("quantity")
			l, _ := cmd.Flags().GetInt("level")

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to load character data")
				return
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to process character")
				return
			}

			if a != "" {
				err = c.ModifyAbilityScoreImprovementItem(q, a)
				if err != nil {
					logger.Error(err)
					logger.PrintError("Failed to modify ability score improvement item")
					return
				}
			} else if l > 0 {
				if l > 20 {
					logger.PrintError("Level must be no more than 20")
				}

				c.SetLevel(l)
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			if buildMd {
				err = handlers.BuildCharacterMarkdown(*c)
				if err != nil {
					logger.Error(err)
					logger.PrintError("failed to generate markdown file")
					return
				}
			}

			logger.PrintSuccess("Character Update Successful")
		},
	}

	importCmd = &cobra.Command{
		Use:   "import",
		Short: "Imports a character or class, supports inserts and updates",
		Long: `Imports a character or class from a json file 
		Existing characters or classes can be exported via "export" command.
		Will update existing record if ID is provided in json.`,
		Run: func(cmd *cobra.Command, args []string) {
			var entity string
			classType, _ := cmd.Flags().GetString("class-type")
			filePath, _ := cmd.Flags().GetString("file")
			characterName, _ := cmd.Flags().GetString("character-name")

			bytes, err := os.ReadFile(filePath)
			if err != nil {
				logger.Errorf("Error reading file '%v':\n%v", filePath, err.Error())
				logger.PrintError("Failed to import character")
				return
			}

			if classType != "" {
				classType = strings.ToLower(classType)
				entity = "Class"
				err = handlers.ImportClassJson(bytes, characterName, classType)
				if err != nil {
					logger.Error(err)
					logger.PrintError("Failed to import class")
					return
				}
			} else {
				entity = "Character"
				err = handlers.ImportCharacterJson(bytes)
				if err != nil {
					logger.Error(err)
					logger.PrintError("Failed to import character")
					return
				}
			}

			logger.PrintSuccess(fmt.Sprintf("%v Import Successful", entity))
		},
	}

	exportCmd = &cobra.Command{
		Use:   "export",
		Short: "Exports a character or class to a file",
		Long: `Exports a character or class to a json file. 
		Can be altered and re-imported with the "import" command.
		Will update existing record if ID is provided in json.`,
		Run: func(cmd *cobra.Command, args []string) {
			var entity string
			var data []byte
			var err error
			name, _ := cmd.Flags().GetString("name")
			classType, _ := cmd.Flags().GetString("class-type")
			filePath, _ := cmd.Flags().GetString("file")

			if classType != "" {
				entity = "Class"
				data, err = handlers.ExportClassJson(name, classType)
				if err != nil {
					logger.Error(err)
					logger.PrintError("Failed to export class file")
					return
				}
			} else {
				entity = "Character"
				data, err = handlers.ExportCharacterJson(name)
				if err != nil {
					logger.Error(err)
					logger.PrintError("Failed to export character file")
					return
				}
			}

			err = os.WriteFile(filePath, data, 0o644)
			if err != nil {
				logger.PrintError(fmt.Sprintf("Failed to write file '%v'", filePath))
				return
			}

			logger.PrintSuccess(fmt.Sprintf("%v Export Successful", entity))
		},
	}

	classCmd = &cobra.Command{
		Use:   "class",
		Short: "Executes commands on the class",
		Long:  `Executes commands on the class via various flags.`,
		Run: func(cmd *cobra.Command, args []string) {
			e, _ := cmd.Flags().GetString("expertise")
			p, _ := cmd.Flags().GetString("prepared-spell")
			o, _ := cmd.Flags().GetString("oath-spell")
			f, _ := cmd.Flags().GetString("fighting-style")
			v, _ := cmd.Flags().GetString("favored-enemy")
			r, _ := cmd.Flags().GetBool("remove")
			ct, _ := cmd.Flags().GetString("class-type")

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to save character data")
				return
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Error(err)
				logger.PrintError("Failed to process character")
				return
			}

			if e != "" {
				if r {
					logger.PrintError("-> removing expertise skill is not implemented yet")
					return
				}

				err := c.AddExpertiseSkill(e, ct)
				if err != nil {
					logger.Error(err)
					logger.PrintError("Failed to add expertise skill")
					return
				}
			} else if p != "" {
				if r {
					err = c.RemovePreparedSpell(p, ct)
					if err != nil {
						logger.Error(err)
						logger.PrintError("Failed to remove prepared spell")
						return
					}
				} else {
					err = c.AddPreparedSpell(p, ct)
					if err != nil {
						logger.Error(err)
						logger.PrintError("Failed to add prepared spell")
						return
					}
				}
			} else if o != "" {
				if r {
					err = c.RemoveOathSpell(o, ct)
					if err != nil {
						logger.Error(err)
						logger.PrintError("Failed to remove oath spell")
						return
					}
				} else {
					err = c.AddOathSpell(o, ct)
					if err != nil {
						logger.Error(err)
						logger.PrintError("Failed to add oath spell")
						return
					}
				}
			} else if f != "" {
				if r {
					logger.PrintError("-> removing fighting style is not implemented")
					return
				} else {
					err = c.ModifyFightingStyle(f, ct)
					if err != nil {
						logger.Error(err)
						logger.PrintError("Failed to modify fighting style")
						return
					}
				}
			} else if v != "" {
				if r {
					err = c.RemoveFavoredEnemy(v, ct)
					if err != nil {
						logger.Error(err)
						logger.PrintError("Failed to remove favored enemy")
						return
					}
				} else {
					err = c.AddFavoredEnemy(v, ct)
					if err != nil {
						logger.Error(err)
						logger.PrintError("Failed to add favored enemy")
						return
					}
				}
			}

			for _, class := range c.Classes {
				err = handlers.SaveClass(class)
				if err != nil {
					logger.Error(err)
					logger.PrintError(fmt.Sprintf("Failed to save data for class '%s'", class.GetClassType()))
					return
				}
			}

			if buildMd {
				err = handlers.BuildCharacterMarkdown(*c)
				if err != nil {
					logger.Error(err)
					logger.PrintError("failed to generate markdown file")
					return
				}
			}

			logger.PrintSuccess("Character Update Successful")
		},
	}
)

func init() {
	characterCmd.AddCommand(addCmd,
		removeCmd,
		updateCmd,
		useCmd,
		recoverCmd,
		initCmd,
		getCmd,
		deleteCmd,
		equipCmd,
		unequipCmd,
		modifyCmd,
		importCmd,
		exportCmd,
		classCmd)

	characterCmd.Flags().BoolVar(&buildMd, "build-md", false, "generate markdown file")

	addCmd.Flags().StringP("ability-improvement", "a", "", "Ability Score Improvement item name, (use -q to specify a quantity)")
	addCmd.Flags().StringP("equipment", "e", "", "Kind of quipment to add 'armor, ring, etc'")
	addCmd.Flags().BoolP("level", "l", false, "Level to add")
	addCmd.Flags().StringP("language", "", "", "Language to add")
	addCmd.Flags().StringP("weapon", "w", "", "Weapon to add")
	addCmd.Flags().IntP("spell-slots", "s", 0, "Increase spell-slot max capacity by level")
	addCmd.Flags().StringP("spell", "x", "", "Add spell to list of character spells")
	addCmd.Flags().StringP("backpack", "b", "", "Item to add to backpack (use -q to specify quantity)")
	addCmd.Flags().IntP("quantity", "q", 0, "Modify quantity of something")
	addCmd.Flags().IntP("temp-hp", "t", 0, "Add temporary hp")
	addCmd.Flags().StringP("name", "n", "", "Name of equipment to add")
	addCmd.Flags().StringP("sub-class", "u", "", "Name of sub-class to add")
	addCmd.Flags().StringP("class-type", "c", "", "class type to modify (only required for multi-class)")

	removeCmd.Flags().StringP("language", "l", "", "Language to remove")
	removeCmd.Flags().StringP("weapon", "w", "", "Weapon to remove")
	removeCmd.Flags().StringP("backpack", "b", "", "Item to remove from backpack")
	removeCmd.Flags().IntP("hitpoints", "p", 0, "Include or modify hitpoints")

	useCmd.Flags().IntP("spell-slots", "s", 0, "Use spell-slot by level")
	useCmd.Flags().StringP("backpack", "b", "", "Use item from backpack")
	useCmd.Flags().IntP("quantity", "q", 0, "Modify quantity of something")
	useCmd.Flags().StringP("class-tokens", "t", "any", "Use class-tokens by token name")
	useCmd.Flags().StringP("class-type", "c", "", "class type to modify (only required for multi-class)")

	recoverCmd.Flags().IntP("spell-slots", "s", 0, "recover spell-slot by level")
	recoverCmd.Flags().BoolP("all", "a", false, "recover all health, slots, and tokens")
	recoverCmd.Flags().IntP("hitpoints", "p", 0, "recover hitpoints")
	recoverCmd.Flags().StringP("class-tokens", "t", "all", "recover class-tokens by token name")
	recoverCmd.Flags().StringP("class-type", "c", "", "class type to modify (only required for multi-class)")
	recoverCmd.Flags().IntP("quantity", "q", 0, "recover the quantity of something")

	initCmd.Flags().StringP("class", "c", "", "name of character class")
	initCmd.Flags().StringP("name", "n", "", "name of character")
	initCmd.MarkFlagRequired("class")
	initCmd.MarkFlagRequired("name")

	equipCmd.Flags().StringP("primary", "p", "", "equip primary weapon or shield")
	equipCmd.Flags().StringP("secondary", "s", "", "equip secondary weapon or shield")

	unequipCmd.Flags().BoolP("primary", "p", false, "equip primary weapon or shield")
	unequipCmd.Flags().BoolP("secondary", "s", false, "equip secondary weapon or shield")

	getCmd.Flags().StringP("path", "p", "", "get config or markdown path")
	getCmd.Flags().BoolP("tokens", "t", false, "get class tokens")
	getCmd.Flags().BoolP("character-names", "n", false, "get character names")

	deleteCmd.Flags().StringP("name", "n", "", "name of character to delete")
	deleteCmd.MarkFlagRequired("name")

	updateCmd.Flags().StringP("default-character-name", "d", "", "name of character to make default")

	importCmd.Flags().StringP("class-type", "c", "", "class type for class file import (default: character)")
	importCmd.Flags().StringP("file", "f", "", "relative path to json file")
	importCmd.Flags().StringP("character-name", "n", "", "name of character, only used when importing Class data")
	importCmd.MarkFlagRequired("file")
	importCmd.MarkFlagFilename("file")
	importCmd.MarkFlagsRequiredTogether("class-type", "character-name")

	exportCmd.Flags().StringP("class-type", "c", "", "class type for class file exports (default: character)")
	exportCmd.Flags().StringP("name", "n", "", "name of character to export data for")
	exportCmd.Flags().StringP("file", "f", "", "name of output file")
	exportCmd.MarkFlagRequired("name")
	exportCmd.MarkFlagRequired("file")

	modifyCmd.Flags().StringP("ability-improvement", "a", "", "Ability Score Improvement item name, (use -q to specify a quantity)")
	modifyCmd.Flags().IntP("quantity", "q", 0, "Modify quantity of something")

	classCmd.Flags().StringP("expertise", "e", "", "name of skill to add to expertise")
	classCmd.Flags().StringP("prepared-spell", "p", "", "name of spell to prepare")
	classCmd.Flags().StringP("fighting-style", "f", "", "name of fighting style to assign")
	classCmd.Flags().StringP("favored-enemy", "v", "", "name of favored to assign")
	classCmd.Flags().StringP("oath-spell", "o", "", "name of oath spell to add")
	classCmd.Flags().BoolP("remove", "r", false, "remove instead of add one of these things")
	classCmd.Flags().StringP("class-type", "c", "", "class type to modify (only required for multi-class)")
}
