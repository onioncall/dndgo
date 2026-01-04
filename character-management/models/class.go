package models

import (
	"fmt"
)

type BaseClass struct {
	SubClass    string `json:"sub-class" clover:"sub-class"`
	CharacterID string `json:"character-id" clover:"character-id"`
	ClassName   string `json:"class-name" clover:"class-name"`
	// This will be used when we implement multiclassing
	Level         int            `json:"level" clover:"level"`
	OtherFeatures []ClassFeature `json:"other-features" clover:"other-features"`
}

type Class interface {
	CalculateHitDice(level int) string
	ClassDetails(level int) string
	GetClassFeatures(level int) string
	SetSubClass(subClass string)
	GetSubClass(level int) string
	GetCharacterId() string
	SetCharacterId(id string)
	GetClassName() string
	SetClassName(name string)
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

func (c *BaseClass) GetClassName() string {
	return c.ClassName
}

func (c *BaseClass) SetClassName(name string) {
	c.ClassName = name
}

func (c *BaseClass) GetSubClass(level int) string {
	if level <= 2 {
		return "Too early, you need to level up!"
	}

	return c.SubClass
}

func (c *BaseClass) SetSubClass(subClass string) {
	c.SubClass = subClass
}

func (c *BaseClass) GetClassFeatures(level int) string {
	var s string
	if len(c.OtherFeatures) > 0 {
		for _, feature := range c.OtherFeatures {
			if feature.Level > level {
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
