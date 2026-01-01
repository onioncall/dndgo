package cmd

import (
	"fmt"
	"os"

	"github.com/onioncall/dndgo/character-management/db"
	"github.com/onioncall/dndgo/character-management/handlers"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/logger"

	"github.com/spf13/cobra"
)

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

			c, err := db.Repo.GetCharacter()
			if err != nil {
				logger.Info("Failed to load character data")
				panic(err)
			}

			if l != "" {
				c.AddLanguage(l)
			}
			if e != "" {
				if n == "" {
					fmt.Println("Name of equipment can not be left empty")
					return
				}

				c.AddEquipment(e, n)
			}
			if bp != "" {
				if q <= 0 {
					fmt.Println("Must pass a positive quantity to add")
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
				c.AddSubClass(sc)
			}
			if a != "" {
				err = c.AddAbilityScoreImprovementItem(q, a)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			err = db.Repo.SyncCharacter(*c)
			if err != nil {
				logger.Info("Failed to save character data")
				panic(err)
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Info("Failed to process character")
				panic(err)
			}

			logger.Info("Character Update Successful")
		},
	}

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get character specific data",
		Run: func(cmd *cobra.Command, args []string) {
			p, _ := cmd.Flags().GetString("path")
			t, _ := cmd.Flags().GetBool("tokens")
			c, _ := cmd.Flags().GetBool("character-names")

			if p != "" {
				var path string
				var err error

				if p == "config" {
					path, err = handlers.GetConfigPath()
					if err != nil {
						logger.Info("Failed to get config path")
						panic(err)
					}
				} else if p == "markdown" {
					c, err := handlers.LoadCharacter()
					if err != nil {
						logger.Info("Failed to load character")
						panic(err)
					}

					path = c.Path
					if path == "" {
						path, err = handlers.GetConfigPath()
					}
				} else {
					logger.Info("path option not found")
				}

				fmt.Printf("Path: %s\n", path)
			}

			if t {
				c, err := handlers.LoadCharacter()
				if err != nil {
					logger.Info("Failed to save character data")
					panic(err)
				}

				if c.Class == nil {
					fmt.Println("Class not properly configured")
					return
				}

				var tokenClass models.TokenClass
				tokenClass, ok := c.Class.(models.TokenClass)
				if !ok {
					fmt.Println("Class does not implement TokenClass")
					return
				}

				tokens := tokenClass.GetTokens()
				if len(tokens) == 0 {
					fmt.Println("Class has no tokens implemented")
				} else if len(tokens) == 1 {
					fmt.Println("(Class only has one token, when modifying token values for this class you may enter any value)")
					fmt.Printf("-> %s\n", tokens[0])
				} else {
					fmt.Println("Tokens:")
					for _, token := range tokens {
						fmt.Printf("%s\n", token)
					}
				}
			}

			if c {
				names, err := handlers.GetCharacterNames()
				if err != nil {
					panic(err)
				}

				if len(names) == 0 {
					fmt.Println("No characters found")
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

			if n == "" {
				fmt.Println("Name cannot be empty")
				return
			}

			if c {
				// Delete class by type and character name
				fmt.Println("Feature to delete single class is not yet supported")
			} else {
				err := handlers.DeleteCharacter(n)
				if err != nil {
					fmt.Println("Failed to delete character")
					logger.Error(err)
				}
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
				logger.Info("Failed to load character")
				panic(err)
			}

			if hp > 0 {
				c.DamageCharacter(hp)
			} else if u > 0 {
				c.UseSpellSlot(u)
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Info("Failed to save character data")
				panic(err)
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Info("Failed to process character")
				panic(err)
			}

			logger.Info("Character Update Successful")
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
					fmt.Println("Failed to update default character")
				}
			}

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Info("Failed to save character data")
				panic(err)
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Info("Failed to process character")
				panic(err)
			}

			logger.Info("Character Update Successful")
		},
	}

	useCmd = &cobra.Command{
		Use:   "use",
		Short: "Use character items/spell slots",
		Run: func(cmd *cobra.Command, args []string) {
			s, _ := cmd.Flags().GetInt("spell-slots")
			bp, _ := cmd.Flags().GetString("backpack")
			q, _ := cmd.Flags().GetInt("quantity")
			ct, _ := cmd.Flags().GetString("class-tokens")

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Info("Failed to save character data")
				panic(err)
			}

			if bp != "" {
				if q <= 0 {
					logger.Info("Must pass a positive quantity to use")
					return
				}

				err = c.RemoveItemFromPack(bp, q)
				logger.Info(err)
			} else if s > 0 {
				c.UseSpellSlot(s)
			} else if ct != "" {
				q = max(q, 1) // If q isn't provided with a valid value, we use one by default
				c.UseClassTokens(ct, q)
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Info("Failed to save character data")
				panic(err)
			}

			err = handlers.SaveClassHandler(c.Class)
			if err != nil {
				logger.Info("Failed to save class data")
				panic(err)
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Info("Failed to process character")
				panic(err)
			}

			logger.Info("Character Update Successful")
		},
	}

	recoverCmd = &cobra.Command{
		Use:   "recover",
		Short: "Recover health and spell slot usage",
		Run: func(cmd *cobra.Command, args []string) {
			a, _ := cmd.Flags().GetBool("all")
			ss, _ := cmd.Flags().GetInt("spell-slots")
			hp, _ := cmd.Flags().GetInt("hitpoints")
			ct, _ := cmd.Flags().GetString("class-tokens")
			q, _ := cmd.Flags().GetInt("quantity")

			c, err := handlers.LoadCharacter()
			if err != nil {
				errMsg := "Failed to save character data"
				logger.Info(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			if a {
				c.Recover()
			} else if ss > 0 {
				c.RecoverSpellSlots(ss, q)
			} else if hp > 0 {
				c.HealCharacter(hp)
			} else if ct != "" {
				c.RecoverClassTokens(ct, q)
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Info("Failed to save character data")
				panic(err)
			}

			err = handlers.SaveClassHandler(c.Class)
			if err != nil {
				logger.Info("Failed to save class data")
				panic(err)
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Info("Failed to process character")
				panic(err)
			}

			logger.Info("Character Update Successful")
		},
	}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes a new character on this machine",
		Run: func(cmd *cobra.Command, args []string) {
			c, _ := cmd.Flags().GetString("class")
			n, _ := cmd.Flags().GetString("name")

			character, err := handlers.LoadCharacterTemplate(n, c)
			if err != nil {
				logger.Info("Failed to load character template")
				panic(err)
			}

			cid, err := db.Repo.InsertCharacter(*character)
			if err != nil {
				logger.Info("Failed to save new character data")
				panic(err)
			}

			class, err := handlers.LoadClassTemplate(c)
			if err != nil {
				errMsg := "Failed to load class template"
				logger.Info(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			class.SetCharacterId(cid)
			class.SetClassName(c)

			if err = db.Repo.InsertClass(class); err != nil {
				errMsg := "Failed to save class data"
				logger.Info(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			logger.Info("Character Creation Successful")
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
				logger.Info("Failed to save character data")
				panic(err)
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
				logger.Info("Failed to save character data")
				panic(err)
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Info("Failed to process character")
				panic(err)
			}

			logger.Info("Character Update Successful")
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
				logger.Error("Failed to save character data")
				panic(err)
			}

			if p == true {
				c.Unequip(true)
			}
			if s == true {
				c.Unequip(false)
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Error("Failed to save character data")
				panic(err)
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Error("Failed to process character")
				panic(err)
			}

			logger.Info("Character Update Successful")
		},
	}

	modifyCmd = &cobra.Command{
		Use:   "modify",
		Short: "modify character attributes",
		Run: func(cmd *cobra.Command, args []string) {
			a, _ := cmd.Flags().GetString("ability-improvement")
			q, _ := cmd.Flags().GetInt("quantity")

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Info("Failed to save character data")
				panic(err)
			}

			if a != "" {
				err = c.ModifyAbilityScoreImprovementItem(q, a)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			err = handlers.SaveCharacter(c)
			if err != nil {
				logger.Error("Failed to save character data")
				panic(err)
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				logger.Info("Failed to process character")
				panic(err)
			}

			logger.Info("Character Update Successful")
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
			isClass, _ := cmd.Flags().GetBool("class")
			filePath, _ := cmd.Flags().GetString("file")
			characterName, _ := cmd.Flags().GetString("character-name")

			bytes, err := os.ReadFile(filePath)
			if err != nil {
				logger.Errorf("Error reading file '%v':\n%v", filePath, err.Error())
			}

			if isClass {
				entity = "Class"
				handlers.ImportClassJson(bytes, characterName)
			} else {
				entity = "Character"
				err = handlers.ImportCharacterJson(bytes)
				if err != nil {
					logger.Error(err)
					fmt.Println("Failed to import character character")
				}
			}

			logger.Infof("%v Import Successful", entity)
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
			isClass, _ := cmd.Flags().GetBool("class")
			filePath, _ := cmd.Flags().GetString("file")

			if isClass {
				entity = "Class"
				data, err = handlers.ExportClassJson(name)
			} else {
				entity = "Character"
				data, err = handlers.ExportCharacterJson(name)
			}

			err = os.WriteFile(filePath, data, 0o644)
			if err != nil {
				logger.Infof("Failed to write file '%v'", filePath)
				panic(err)
			}

			logger.Infof("%v Export Successful", entity)
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
		exportCmd)

	characterCmd.Flags().StringP("default", "d", "", "Name of character to update to default")

	addCmd.Flags().StringP("ability-improvement", "a", "", "Ability Score Improvement item name, (use -q to specify a quantity)")
	addCmd.Flags().StringP("equipment", "e", "", "Kind of quipment to add 'armor, ring, etc'")
	addCmd.Flags().StringP("language", "l", "", "Language to add")
	addCmd.Flags().StringP("weapon", "w", "", "Weapon to add")
	addCmd.Flags().IntP("spell-slots", "s", 0, "Increase spell-slot max capacity by level")
	addCmd.Flags().StringP("spell", "x", "", "Add spell to list of character spells")
	addCmd.Flags().StringP("backpack", "b", "", "Item to add to backpack (use -q to specify quantity)")
	addCmd.Flags().IntP("quantity", "q", 0, "Modify quantity of something")
	addCmd.Flags().IntP("temp-hp", "t", 0, "Add temporary hp")
	addCmd.Flags().StringP("name", "n", "", "Name of equipment to add")
	addCmd.Flags().StringP("sub-class", "c", "", "Name of sub-class to add")

	removeCmd.Flags().StringP("language", "l", "", "Language to remove")
	removeCmd.Flags().StringP("weapon", "w", "", "Weapon to remove")
	removeCmd.Flags().StringP("backpack", "b", "", "Item to remove from backpack")
	removeCmd.Flags().IntP("hitpoints", "p", 0, "Include or modify hitpoints")

	useCmd.Flags().IntP("spell-slots", "s", 0, "Use spell-slot by level")
	useCmd.Flags().StringP("backpack", "b", "", "Use item from backpack")
	useCmd.Flags().IntP("quantity", "q", 0, "Modify quantity of something")
	useCmd.Flags().StringP("class-tokens", "c", "any", "Use class-tokens by token name")

	recoverCmd.Flags().IntP("spell-slots", "s", 0, "Recover spell-slot by level")
	recoverCmd.Flags().BoolP("all", "a", false, "Recover all health, slots, and tokens")
	recoverCmd.Flags().IntP("hitpoints", "p", 0, "Recover hitpoints")
	recoverCmd.Flags().StringP("class-tokens", "c", "all", "Recover class-tokens by token name")
	recoverCmd.Flags().IntP("quantity", "q", 0, "Recover the quantity of something")

	initCmd.Flags().StringP("class", "c", "", "Name of character class")
	initCmd.Flags().StringP("name", "n", "", "Name of character")
	initCmd.MarkFlagRequired("class")
	initCmd.MarkFlagRequired("name")

	equipCmd.Flags().StringP("primary", "p", "", "Equip primary weapon or shield")
	equipCmd.Flags().StringP("secondary", "s", "", "Equip secondary weapon or shield")

	unequipCmd.Flags().BoolP("primary", "p", false, "Equip primary weapon or shield")
	unequipCmd.Flags().BoolP("secondary", "s", false, "Equip secondary weapon or shield")

	getCmd.Flags().StringP("path", "p", "", "Get config or markdown path")
	getCmd.Flags().BoolP("tokens", "t", false, "Get class tokens")
	getCmd.Flags().BoolP("character-names", "c", false, "Get character names")

	deleteCmd.Flags().StringP("name", "n", "", "Name of character to delete")
	deleteCmd.MarkFlagRequired("name")

	updateCmd.Flags().StringP("default-character-name", "d", "", "Name of character to make default")

	importCmd.Flags().BoolP("class", "c", false, "Import Class file (default: Character)")
	importCmd.Flags().StringP("file", "f", "", "Relative path to json file")
	importCmd.Flags().StringP("character-name", "n", "", "Name of character, only used when importing Class data")
	importCmd.MarkFlagRequired("file")
	importCmd.MarkFlagFilename("file")
	importCmd.MarkFlagsRequiredTogether("class", "character-name")

	exportCmd.Flags().BoolP("class", "c", false, "Export Class data (will otherwise default to Character data)")
	exportCmd.Flags().StringP("name", "n", "", "Name of Character to export data for")
	exportCmd.Flags().StringP("file", "f", "", "Name of output file")
	exportCmd.MarkFlagRequired("name")
	exportCmd.MarkFlagRequired("file")

	modifyCmd.Flags().StringP("ability-improvement", "a", "", "Ability Score Improvement item name, (use -q to specify a quantity)")
	modifyCmd.Flags().IntP("quantity", "q", 0, "Modify quantity of something")
}
