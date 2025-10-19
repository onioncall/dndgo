package cmd

import (
	"fmt"

	"github.com/onioncall/dndgo/handlers"
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
			n, _ := cmd.Flags().GetString("name")

			c, err := handlers.LoadCharacter()
			if err != nil {
				errMsg := "Failed to save character json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			if l != "" {
				c.AddLanguage(l)
			}
			if e != "" {
				if n == "" {
					logger.HandleInfo("Name of equipment can not be left empty")
					return
				}

				c.AddEquipment(e, n)
			}
			if bp != "" {
				if q <= 0 {
					logger.HandleInfo("Must pass a positive quantity to add")
					return
				}

				c.AddItemToPack(bp, q)
			}
			if il {
				// c.AddLevel()
			}
			if ss > 0 {
				// add spell slot for level
			}
			if s != "" {
				err = handlers.AddSpell(c, s)

			}

			err = handlers.SaveCharacterJson(c)
			if err != nil {
				errMsg := "Failed to save character json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			handlers.SaveClassHandler(c.Class)
			if err != nil {
				errMsg := "Failed to save class json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				errMsg := "Failed to handle character"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
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
				errMsg := "Failed to save character json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			if hp > 0 {
				c.DamageCharacter(hp)
			} else if u > 0 {
				c.UseSpellSlot(u)
			}

			err = handlers.SaveCharacterJson(c)
			if err != nil {
				errMsg := "Failed to save character json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			handlers.SaveClassHandler(c.Class)
			if err != nil {
				errMsg := "Failed to save class json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				errMsg := "Failed to handle character"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}
		},
	}

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update character attributes",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := handlers.LoadCharacter()
			if err != nil {
				errMsg := "Failed to save character json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				errMsg := "Failed to handle character"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}
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
				errMsg := "Failed to save character json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			if bp != "" {
				if q <= 0 {
					logger.HandleInfo("Must pass a positive quantity to use")
					return
				}

				c.RemoveItemFromPack(bp, q)
			} else if s > 0 {
				c.UseSpellSlot(s)
			} else if ct != "" {
				c.UseClassTokens(ct)
			}

			err = handlers.SaveCharacterJson(c)
			if err != nil {
				errMsg := "Failed to save character json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			handlers.SaveClassHandler(c.Class)
			if err != nil {
				errMsg := "Failed to save class json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				errMsg := "Failed to handle character"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}
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
				errMsg := "Failed to save character json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			if a {
				c.Recover()
			} else if ss > 0 {
				c.RecoverSpellSlots(ss)
			} else if hp > 0 {
				c.HealCharacter(hp)
			} else if ct != "" {
				c.RecoverClassTokens(ct, q)
			}

			err = handlers.SaveCharacterJson(c)
			if err != nil {
				errMsg := "Failed to save character json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			handlers.SaveClassHandler(c.Class)
			if err != nil {
				errMsg := "Failed to save class json"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			err = handlers.HandleCharacter(c)
			if err != nil {
				errMsg := "Failed to handle character"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}
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
				errMsg := "Failed to load character template"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}
			err = handlers.SaveCharacterJson(character)
			if err != nil {
				errMsg := "Failed to save new character data"
				logger.HandleInfo(errMsg)
				panic(fmt.Errorf("%s: %w", errMsg, err))
			}

			if c != "" {
				class, err := handlers.LoadClassTemplate(c)
				if err != nil {
					errMsg := "Failed to load character class template"
					logger.HandleInfo(errMsg)
					panic(fmt.Errorf("%s: %w", errMsg, err))
				}
				err = handlers.SaveClassHandler(class)
				if err != nil {
					errMsg := "Failed to save new character class data"
					logger.HandleInfo(errMsg)
					panic(fmt.Errorf("%s: %w", errMsg, err))
				}
			}
		},
	}
)

func init() {
	characterCmd.AddCommand(addCmd, removeCmd, updateCmd, useCmd, recoverCmd, initCmd)

	addCmd.Flags().StringP("equipment", "e", "", "Kind of quipment to add 'armor, ring, etc'")
	addCmd.Flags().StringP("language", "l", "", "Language to add")
	addCmd.Flags().StringP("weapon", "w", "", "Weapon to add")
	addCmd.Flags().IntP("spell-slots", "s", 0, "Increase spell-slot max capacity by level")
	addCmd.Flags().StringP("spell", "x", "", "Add spell to list of character spells")
	addCmd.Flags().StringP("backpack", "b", "", "Item to add to backpack (use -q to specify quantity)")
	addCmd.Flags().IntP("quantity", "q", 0, "Modify quantity of something")
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
}
