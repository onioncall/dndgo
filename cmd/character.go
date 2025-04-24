package cmd

import (
	"strconv"
	"strings"

	"github.com/onioncall/dndgo/handlers"
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
			language, _ := cmd.Flags().GetString("language")
			backpack, _ := cmd.Flags().GetString("backpack")
			increaseLevel, _ := cmd.Flags().GetBool("level")
			equipment, _ := cmd.Flags().GetString("equipment")
			s, _ := cmd.Flags().GetString("spell")
			ss, _ := cmd.Flags().GetInt("spell-slots")

			c, err := handlers.LoadCharacter()
			if err != nil {
				panic(err)
			}
			
			if language != "" {
				c.AddLanguage(language)
			}
			if equipment != "" {
				equipmentArgs := strings.Split(equipment, " ")
				c.AddEquipment(equipmentArgs[0], strings.Join(equipmentArgs[1:], " ")) 
			}
			if backpack != "" {
				arg, quantity := getArgumentAndQuantity(backpack)
				c.AddItemToPack(arg, quantity)
			}
			if increaseLevel {
				// c.AddLevel()				
			}
			if ss > 0 {
				// add spell slot for level
			} 
			if s != "" {
				handlers.AddSpell(c, s)
			}
			
			handlers.SaveCharacterJson(c)
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
				panic(err)
			}

			if hp > 0 {
				c.DamageCharacter(hp)
			} else if u > 0 {
				c.UseSpellSlot(u);
			} 

			handlers.SaveCharacterJson(c)
			handlers.HandleCharacter(c)
		},
	}

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update character attributes",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := handlers.LoadCharacter()
			if err != nil {
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

			c, err := handlers.LoadCharacter()
			if err != nil {
				panic(err)
			}

			if bp != "" {
				arg, quantity := getArgumentAndQuantity(bp)
				c.RemoveItemFromPack(arg, quantity)
			} else if s > 0 {
				c.UseSpellSlot(s);
			}

			handlers.SaveCharacterJson(c)
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

			c, err := handlers.LoadCharacter()
			if err != nil {
				panic(err)
			}

			if a {
				c.Recover()
			} else if ss > 0 {
				c.RecoverSpellSlots(ss)	
			} else if hp > 0 {
				c.HealCharacter(hp)
			}

			handlers.SaveCharacterJson(c)
			handlers.HandleCharacter(c)
		},
	}
)

func init() {
	characterCmd.AddCommand(addCmd, removeCmd, updateCmd, useCmd, recoverCmd)

	addCmd.Flags().StringP("equipment", "e", "", "Equipment to add")
	addCmd.Flags().StringP("language", "l", "", "Language to add")
	addCmd.Flags().StringP("weapon", "w", "", "Weapon to add")
	addCmd.Flags().IntP("spell-slots", "x", 0, "Increase spell-slot by level")
	addCmd.Flags().StringP("spell", "s", "", "Add spell to list of character spells")
	addCmd.Flags().StringP("backpack", "b", "", "Item to add to backpack")

	removeCmd.Flags().StringP("language", "l", "", "Language to remove")
	removeCmd.Flags().StringP("weapon", "w", "", "Weapon to remove")
	removeCmd.Flags().StringP("backpack", "b", "", "Item to remove from backpack")
	removeCmd.Flags().IntP("hitpoints", "p", 0, "Include or modify hitpoints")

	useCmd.Flags().IntP("spell-slots", "x", 0, "Use spell-slot by level")
	useCmd.Flags().StringP("backpack", "b", "", "Use item from backpack")

	recoverCmd.Flags().IntP("spell-slots", "x", 0, "Use spell-slot by level")
	recoverCmd.Flags().BoolP("all", "a", false, "Use spell-slot by level")
	recoverCmd.Flags().IntP("hitpoints", "p", 0, "Include or modify hitpoints")
}

func getArgumentAndQuantity(input string) (string, int) {
	args := strings.Split(input, " ")
	quantityString := string(args[len(args)-1])
	args = args[:len(args)-1] 
	quantity := 1

	parsedQty, err := strconv.Atoi(quantityString)
	if err == nil {
		quantity = parsedQty
	}

	arg := strings.Join(args, " ")

	return arg, quantity
}
