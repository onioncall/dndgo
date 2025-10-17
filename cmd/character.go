package cmd

import (
	"fmt"
	"os"

	defaultjsonconfigs "github.com/onioncall/dndgo/default-json-configs"
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
				err := fmt.Errorf("Failed to load character: %v", err)
				panic(err)
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
				handlers.AddSpell(c, s)
			}

			handlers.SaveCharacterJson(c)
			handlers.SaveClassHandler(c.Class)
			handlers.HandleCharacter(c)
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
				err := fmt.Errorf("Failed to load character: %v", err)
				panic(err)
			}

			if hp > 0 {
				c.DamageCharacter(hp)
			} else if u > 0 {
				c.UseSpellSlot(u)
			}

			handlers.SaveCharacterJson(c)
			handlers.SaveClassHandler(c.Class)
			handlers.HandleCharacter(c)
		},
	}

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update character attributes",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := handlers.LoadCharacter()
			if err != nil {
				err := fmt.Errorf("Failed to load character: %v", err)
				panic(err)
			}

			handlers.HandleCharacter(c)
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
				err := fmt.Errorf("Failed to load character: %v", err)
				panic(err)
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

			handlers.SaveCharacterJson(c)
			handlers.SaveClassHandler(c.Class)
			handlers.HandleCharacter(c)
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
				err := fmt.Errorf("Failed to load character: %v", err)
				panic(err)
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

			handlers.SaveCharacterJson(c)
			handlers.SaveClassHandler(c.Class)
			handlers.HandleCharacter(c)
		},
	}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes a new character on this machine",
		Run: func(cmd *cobra.Command, args []string) {
			c, _ := cmd.Flags().GetString("class")

			var templateName string
			switch c {
			case "bard":
				templateName = "bard-class.json"
			case "barbarian":
				templateName = "barbarian-class.json"
			case "ranger":
				templateName = "default-ranger.json"
			default:
				panic(fmt.Sprintf("Unsupported class '%v'", c))
			}

			classBytes, err := defaultjsonconfigs.Content.ReadFile(fmt.Sprintf("default-json-configs/%v", templateName))
			if err != nil {
				panic(fmt.Errorf("Failed to read content of %v: %w", templateName, err))
			}
			charBytes, err := defaultjsonconfigs.Content.ReadFile("default-json-configs/default-character.json")
			if err != nil {
				panic(fmt.Errorf("Failed to read content of %v: %w", templateName, err))
			}

			dirname := "$HOME/.config/dndgo"
			classFile := fmt.Sprintf("%v/class.json", dirname)
			charFile := fmt.Sprintf("%v/character.json", dirname)

			if err = os.MkdirAll(classFile, 0755); err != nil {
				panic(fmt.Errorf("Failed to create dir %v: %w", dirname, err))
			}

			if err = os.WriteFile(classFile, classBytes, 0655); err != nil {
				panic(fmt.Errorf("Failed to create dir %v: %w", classFile, err))
			}

			if err = os.WriteFile(charFile, charBytes, 0655); err != nil {
				panic(fmt.Errorf("Failed to write file %v: %w", charFile, err))
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
}
