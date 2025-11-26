package cmd

import (
	"fmt"

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
			l, _ := cmd.Flags().GetString("language")
			bp, _ := cmd.Flags().GetString("backpack")
			il, _ := cmd.Flags().GetBool("level")
			e, _ := cmd.Flags().GetString("equipment")
			s, _ := cmd.Flags().GetString("spell")
			ss, _ := cmd.Flags().GetInt("spell-slots")
			q, _ := cmd.Flags().GetInt("quantity")
			t, _ := cmd.Flags().GetInt("temp-hp")
			n, _ := cmd.Flags().GetString("name")

			c, err := handlers.LoadCharacter()
			if err != nil {
				logger.Info("Failed to save character data")
				panic(err)
			}

			if l != "" {
				c.AddLanguage(l)
			}
			if e != "" {
				if n == "" {
					logger.Info("Name of equipment can not be left empty")
					return
				}

				c.AddEquipment(e, n)
			}
			if bp != "" {
				if q <= 0 {
					logger.Info("Must pass a positive quantity to add")
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

			err = handlers.SaveCharacterJson(c)
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

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get character specific data",
		Run: func(cmd *cobra.Command, args []string) {
			p, _ := cmd.Flags().GetString("path")
			t, _ := cmd.Flags().GetBool("tokens")

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

			err = handlers.SaveCharacterJson(c)
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

				c.RemoveItemFromPack(bp, q)
			} else if s > 0 {
				c.UseSpellSlot(s)
			} else if ct != "" {
				q = max(q, 1) // If q isn't provided with a valid value, we use one by default
				c.UseClassTokens(ct, q)
			}

			err = handlers.SaveCharacterJson(c)
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

			err = handlers.SaveCharacterJson(c)
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
			err = handlers.SaveCharacterJson(character)
			if err != nil {
				logger.Info("Failed to save new character data")
				panic(err)
			}

			if c != "" {
				class, err := handlers.LoadClassTemplate(c)
				if err != nil {
					errMsg := "Failed to load class template"
					logger.Info(errMsg)
					panic(fmt.Errorf("%s: %w", errMsg, err))
				}
				err = handlers.SaveClassHandler(class)
				if err != nil {
					logger.Info("Failed to save new class data")
					panic(err)
				}
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
				c.Equip(true, p)
			}
			if s != "" {
				c.Equip(false, s)
			}

			err = handlers.SaveCharacterJson(c)
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
				logger.Info("Failed to save character data")
				panic(err)
			}

			if p == true {
				c.Unequip(true)
			}
			if s == true {
				c.Unequip(false)
			}

			err = handlers.SaveCharacterJson(c)
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
)

func init() {
	characterCmd.AddCommand(addCmd, removeCmd, updateCmd, useCmd, recoverCmd, initCmd, getCmd, equipCmd, unequipCmd)

	addCmd.Flags().StringP("equipment", "e", "", "Kind of quipment to add 'armor, ring, etc'")
	addCmd.Flags().StringP("language", "l", "", "Language to add")
	addCmd.Flags().StringP("weapon", "w", "", "Weapon to add")
	addCmd.Flags().IntP("spell-slots", "s", 0, "Increase spell-slot max capacity by level")
	addCmd.Flags().StringP("spell", "x", "", "Add spell to list of character spells")
	addCmd.Flags().StringP("backpack", "b", "", "Item to add to backpack (use -q to specify quantity)")
	addCmd.Flags().IntP("quantity", "q", 0, "Modify quantity of something")
	addCmd.Flags().IntP("temp-hp", "t", 0, "Add temporary hp")
	addCmd.Flags().StringP("name", "n", "", "Name of equipment to add")

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
	initCmd.MarkFlagRequired("name")

	equipCmd.Flags().StringP("primary", "p", "", "Equip primary weapon or shield")
	equipCmd.Flags().StringP("secondary", "s", "", "Equip secondary weapon or shield")
	unequipCmd.Flags().BoolP("primary", "p", false, "Equip primary weapon or shield")
	unequipCmd.Flags().BoolP("secondary", "s", false, "Equip secondary weapon or shield")

	getCmd.Flags().StringP("path", "p", "", "Get config or markdown path")
	getCmd.Flags().BoolP("tokens", "t", false, "Get class tokens")
}
