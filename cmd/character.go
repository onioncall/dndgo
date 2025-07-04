package cmd

import (
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
				panic(err)
			}
			
			if l != "" {
				c.AddLanguage(l)
			}
			if e != "" {
				if n == "" {
					panic("Name of equipment can not be left empty")
				}

				c.AddEquipment(e, n) 
			}
			if bp != "" {
				if q <= 0 {
					panic("Must pass a positive quantity to add")
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
			q, _ := cmd.Flags().GetInt("quantity")
			cd, _ := cmd.Flags().GetString("class-detail-slots")

			c, err := handlers.LoadCharacter()
			if err != nil {
				panic(err)
			}

			if bp != "" {
				if q <= 0 {
					panic("Must pass a positive quantity to use")
				}

				c.RemoveItemFromPack(bp, q)
			} else if s > 0 {
				c.UseSpellSlot(s);
			} else if cd != "" {
				c.UseClassSlots(cd)	
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
			cd, _ := cmd.Flags().GetString("class-detail-slots")

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
			} else if cd != "" {
				c.RecoverClassDetailSlots(cd)
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
	addCmd.Flags().IntP("quantity", "q", 0, "Modify quantity of something") 
	addCmd.Flags().StringP("name", "n", "", "Name of equipment to add") 

	removeCmd.Flags().StringP("language", "l", "", "Language to remove")
	removeCmd.Flags().StringP("weapon", "w", "", "Weapon to remove")
	removeCmd.Flags().StringP("backpack", "b", "", "Item to remove from backpack")
	removeCmd.Flags().IntP("hitpoints", "p", 0, "Include or modify hitpoints")

	useCmd.Flags().IntP("spell-slots", "x", 0, "Use spell-slot by level")
	useCmd.Flags().StringP("backpack", "b", "", "Use item from backpack")
	useCmd.Flags().IntP("quantity", "q", 0, "Modify quantity of something") 
	useCmd.Flags().StringP("class-detail-slots", "d", "", "Use class-detail-slot by slot name")

	recoverCmd.Flags().IntP("spell-slots", "x", 0, "Recover spell-slot by level")
	recoverCmd.Flags().BoolP("all", "a", false, "Recover all health and slots")
	recoverCmd.Flags().IntP("hitpoints", "p", 0, "Recover hitpoints")
	recoverCmd.Flags().StringP("class-detail-slots", "d", "", "Recover class-detail-slot by slot name")
}
