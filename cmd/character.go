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
			
			handlers.SaveCharacterJson(c)
			handlers.HandleCharacter(c)
		},
	}

	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove character attributes",
		Run: func(cmd *cobra.Command, args []string) {
			backpack, _ := cmd.Flags().GetString("backpack")
			hp, _ := cmd.Flags().GetInt("hitpoints")

			c, err := handlers.LoadCharacter()
			if err != nil {
				panic(err)
			}
			
			if backpack != "" {
				arg, quantity := getArgumentAndQuantity(backpack)
				c.RemoveItemFromPack(arg, quantity)
			}
			if hp > 0 {
				c.DamageCharacter(hp)
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
)

func init() {
	characterCmd.AddCommand(addCmd, removeCmd, updateCmd)

	characterCmd.PersistentFlags().IntP("hitpoints", "p", 0, "Include or modify hitpoints")

	addCmd.Flags().StringP("equipment", "e", "", "Equipment to add")
	addCmd.Flags().StringP("language", "l", "", "Language to add")
	addCmd.Flags().StringP("weapon", "w", "", "Weapon to add")
	addCmd.Flags().StringP("backpack", "b", "", "Item to add to backpack")

	removeCmd.Flags().StringP("language", "l", "", "Language to remove")
	removeCmd.Flags().StringP("weapon", "w", "", "Weapon to remove")
	removeCmd.Flags().StringP("backpack", "b", "", "Item to remove from backpack")
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
