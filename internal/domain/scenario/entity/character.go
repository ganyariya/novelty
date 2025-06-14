package entity

import (
	"github.com/ganyariya/novelty/internal/domain/presentation/valueobject"
	scenarioVO "github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type Character struct {
	id          scenarioVO.CharacterID
	name        string
	displayName string
	colorTheme  valueobject.ColorTheme
	voiceDir    string
}

func NewCharacter(
	id scenarioVO.CharacterID,
	name string,
	displayName string,
	colorTheme valueobject.ColorTheme,
	voiceDir string,
) *Character {
	return &Character{
		id:          id,
		name:        name,
		displayName: displayName,
		colorTheme:  colorTheme,
		voiceDir:    voiceDir,
	}
}

func (c *Character) ID() scenarioVO.CharacterID {
	return c.id
}

func (c *Character) Name() string {
	return c.name
}

func (c *Character) DisplayName() string {
	return c.displayName
}

func (c *Character) ColorTheme() valueobject.ColorTheme {
	return c.colorTheme
}

func (c *Character) VoiceDir() string {
	return c.voiceDir
}

func (c *Character) SetDisplayName(displayName string) {
	c.displayName = displayName
}

func (c *Character) SetColorTheme(colorTheme valueobject.ColorTheme) {
	c.colorTheme = colorTheme
}

func (c *Character) SetVoiceDir(voiceDir string) {
	c.voiceDir = voiceDir
}