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
			s, _ := cmd.Flags().GetInt("spell-slots")
			hp, _ := cmd.Flags().GetInt("hitpoints")

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
			if hp > 0 {
				c.HealCharacter(hp)
			}
			if s > 0 {
				c.RecoverSpellSlots()
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
			}
			if u > 0 {
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
			}

			if s > 0 {
				c.UseSpellSlot(s);
			}

			handlers.SaveCharacterJson(c)
			handlers.HandleCharacter(c)
		},
	}

	recoverCmd = &cobra.Command{
		Use:   "recover",
		Short: "Recover all health and spell slots",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := handlers.LoadCharacter()
			if err != nil {
				panic(err)
			}

			c.Recover()
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
	addCmd.Flags().IntP("hitpoints", "p", 0, "Include or modify hitpoints")
	addCmd.Flags().IntP("spell-slots", "s", 0, "Use spell-slot by level")
	addCmd.Flags().StringP("backpack", "b", "", "Item to remove from backpack")

	removeCmd.Flags().StringP("language", "l", "", "Language to remove")
	removeCmd.Flags().StringP("weapon", "w", "", "Weapon to remove")
	removeCmd.Flags().StringP("backpack", "b", "", "Item to remove from backpack")
	removeCmd.Flags().IntP("hitpoints", "p", 0, "Include or modify hitpoints")

	useCmd.Flags().IntP("spell-slots", "s", 0, "Use spell-slot by level")
	useCmd.Flags().StringP("backpack", "b", "", "Use item from backpack")
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
