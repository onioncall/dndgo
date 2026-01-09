package models

import (
	"fmt"
)

type BaseClass struct {
	SubClass      string         `json:"sub-class" clover:"sub-class"`
	CharacterID   string         `json:"character-id" clover:"character-id"`
	ClassType     string         `json:"class-type" clover:"class-type"`
	Level         int            `json:"level" clover:"level"`
	OtherFeatures []ClassFeature `json:"other-features" clover:"other-features"`
}

type Class interface {
	CalculateHitDice() string
	ClassDetails() string
	GetClassFeatures() string
	SetSubClass(subClass string)
	GetClassType() string
	GetClassLevel() int
	SetClassLevel(int)
	GetSubClass() string
	GetCharacterId() string
	SetCharacterId(id string)
	SetClassType(name string)
}

type PostCalculator interface {
	ExecutePostCalculateMethods(c *Character)
}

type PreCalculator interface {
	ExecutePreCalculateMethods(c *Character)
}

type TokenClass interface {
	GetTokens() []string
	UseClassTokens(string, int)
	RecoverClassTokens(string, int)
}

type ExpertiseClass interface {
	AddExpertiseSkill(skill string) error
}

type PreparedSpellClass interface {
	AddPreparedSpell(spell string) error
	RemovePreparedSpell(spell string) error
	GetPreparedSpells() []string
}

type SpellCasterClass interface {
	UseSpellSlot(level int) error
	RecoverSpellSlots(level int, quantity int)
}

type OathSpellClass interface {
	AddOathSpell(spell string) error
	RemoveOathSpell(spell string) error
}

type FightingStyleClass interface {
	ModifyFightingStyle(fightingStyle string) error
}

type FavoredEnemyClass interface {
	AddFavoredEnemy(favoredEnemy string) error
	RemoveFavoredEnemy(favoredEnemy string) error
}

type ClassFeature struct {
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Details string `json:"details"`
}

func (c *BaseClass) GetCharacterId() string {
	return c.CharacterID
}

func (c *BaseClass) SetCharacterId(id string) {
	c.CharacterID = id
}

func (c *BaseClass) GetClassType() string {
	return c.ClassType
}

func (c *BaseClass) GetClassLevel() int {
	return c.Level
}

func (c *BaseClass) SetClassLevel(level int) {
	c.Level += level
}

func (c *BaseClass) SetClassType(name string) {
	c.ClassType = name
}

func (c *BaseClass) GetSubClass() string {
	return c.SubClass
}

func (c *BaseClass) SetSubClass(subClass string) {
	c.SubClass = subClass
}

func (c *BaseClass) GetClassFeatures() string {
	var s string
	if len(c.OtherFeatures) > 0 {
		for _, feature := range c.OtherFeatures {
			if feature.Level > c.Level {
				continue
			}

			featureName := fmt.Sprintf("---\n**%s**\n", feature.Name)
			s += featureName
			features := fmt.Sprintf("%s\n", feature.Details)
			s += features
		}
	}

	return s
}
