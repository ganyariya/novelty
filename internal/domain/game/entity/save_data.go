package entity

import (
	"time"

	"github.com/ganyariya/novelty/internal/domain/game/valueobject"
	scenarioVO "github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type SaveData struct {
	slot            valueobject.SaveSlot
	currentScene    scenarioVO.SceneID
	currentFunction string
	flags           map[string]bool
	variables       map[string]int
	readHistory     map[string]bool
	playTime        time.Duration
	saveTime        time.Time
	description     string
}

func NewSaveData(
	slot valueobject.SaveSlot,
	gameState *GameState,
	description string,
) *SaveData {
	return &SaveData{
		slot:            slot,
		currentScene:    gameState.CurrentScene(),
		currentFunction: gameState.CurrentFunction(),
		flags:           gameState.Flags().ToMap(),
		variables:       gameState.Variables().ToMap(),
		readHistory:     gameState.ReadHistory(),
		playTime:        gameState.PlayTime(),
		saveTime:        time.Now(),
		description:     description,
	}
}

func (s *SaveData) Slot() valueobject.SaveSlot {
	return s.slot
}

func (s *SaveData) CurrentScene() scenarioVO.SceneID {
	return s.currentScene
}

func (s *SaveData) CurrentFunction() string {
	return s.currentFunction
}

func (s *SaveData) Flags() map[string]bool {
	result := make(map[string]bool)
	for k, v := range s.flags {
		result[k] = v
	}
	return result
}

func (s *SaveData) Variables() map[string]int {
	result := make(map[string]int)
	for k, v := range s.variables {
		result[k] = v
	}
	return result
}

func (s *SaveData) ReadHistory() map[string]bool {
	result := make(map[string]bool)
	for k, v := range s.readHistory {
		result[k] = v
	}
	return result
}

func (s *SaveData) PlayTime() time.Duration {
	return s.playTime
}

func (s *SaveData) SaveTime() time.Time {
	return s.saveTime
}

func (s *SaveData) Description() string {
	return s.description
}

func (s *SaveData) SetDescription(description string) {
	s.description = description
}