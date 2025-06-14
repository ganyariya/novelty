package entity

import (
	"time"

	"github.com/ganyariya/novelty/internal/domain/game/valueobject"
	scenarioEntity "github.com/ganyariya/novelty/internal/domain/scenario/entity"
	scenarioVO "github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type GameState struct {
	currentScene    scenarioVO.SceneID
	currentFunction string
	flags           valueobject.GameFlags
	variables       valueobject.Variables
	history         []*scenarioEntity.Message
	readHistory     map[string]bool
	startTime       time.Time
	playTime        time.Duration
}

func NewGameState(startScene scenarioVO.SceneID) *GameState {
	return &GameState{
		currentScene:    startScene,
		currentFunction: "start",
		flags:           valueobject.NewGameFlags(),
		variables:       valueobject.NewVariables(),
		history:         make([]*scenarioEntity.Message, 0),
		readHistory:     make(map[string]bool),
		startTime:       time.Now(),
		playTime:        0,
	}
}

func (g *GameState) CurrentScene() scenarioVO.SceneID {
	return g.currentScene
}

func (g *GameState) CurrentFunction() string {
	return g.currentFunction
}

func (g *GameState) Flags() valueobject.GameFlags {
	return g.flags
}

func (g *GameState) Variables() valueobject.Variables {
	return g.variables
}

func (g *GameState) History() []*scenarioEntity.Message {
	result := make([]*scenarioEntity.Message, len(g.history))
	copy(result, g.history)
	return result
}

func (g *GameState) ReadHistory() map[string]bool {
	result := make(map[string]bool)
	for k, v := range g.readHistory {
		result[k] = v
	}
	return result
}

func (g *GameState) PlayTime() time.Duration {
	return g.playTime + time.Since(g.startTime)
}

func (g *GameState) SetCurrentScene(sceneID scenarioVO.SceneID, functionName string) {
	g.currentScene = sceneID
	g.currentFunction = functionName
}

func (g *GameState) SetFlag(name string, value bool) {
	g.flags = g.flags.Set(name, value)
}

func (g *GameState) GetFlag(name string) bool {
	return g.flags.Get(name)
}

func (g *GameState) SetVariable(name string, value int) {
	g.variables = g.variables.Set(name, value)
}

func (g *GameState) GetVariable(name string) int {
	return g.variables.Get(name)
}

func (g *GameState) AddVariable(name string, value int) {
	g.variables = g.variables.Add(name, value)
}

func (g *GameState) AddToHistory(message *scenarioEntity.Message) {
	g.history = append(g.history, message)
}

func (g *GameState) MarkAsRead(messageID string) {
	g.readHistory[messageID] = true
}

func (g *GameState) IsRead(messageID string) bool {
	return g.readHistory[messageID]
}

func (g *GameState) UpdatePlayTime() {
	g.playTime = g.PlayTime()
	g.startTime = time.Now()
}