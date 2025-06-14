package entity

import (
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type ScenarioID struct {
	value string
}

func NewScenarioID(value string) ScenarioID {
	return ScenarioID{value: value}
}

func (s ScenarioID) String() string {
	return s.value
}

type Scenario struct {
	id         ScenarioID
	title      string
	scenes     map[valueobject.SceneID]*Scene
	characters map[valueobject.CharacterID]*Character
	startScene valueobject.SceneID
}

func NewScenario(id ScenarioID, title string, startScene valueobject.SceneID) *Scenario {
	return &Scenario{
		id:         id,
		title:      title,
		scenes:     make(map[valueobject.SceneID]*Scene),
		characters: make(map[valueobject.CharacterID]*Character),
		startScene: startScene,
	}
}

func (s *Scenario) ID() ScenarioID {
	return s.id
}

func (s *Scenario) Title() string {
	return s.title
}

func (s *Scenario) StartScene() valueobject.SceneID {
	return s.startScene
}

func (s *Scenario) AddScene(scene *Scene) {
	s.scenes[scene.ID()] = scene
}

func (s *Scenario) GetScene(sceneID valueobject.SceneID) (*Scene, bool) {
	scene, exists := s.scenes[sceneID]
	return scene, exists
}

func (s *Scenario) GetAllScenes() map[valueobject.SceneID]*Scene {
	result := make(map[valueobject.SceneID]*Scene)
	for k, v := range s.scenes {
		result[k] = v
	}
	return result
}

func (s *Scenario) AddCharacter(character *Character) {
	s.characters[character.ID()] = character
}

func (s *Scenario) GetCharacter(characterID valueobject.CharacterID) (*Character, bool) {
	character, exists := s.characters[characterID]
	return character, exists
}

func (s *Scenario) GetAllCharacters() map[valueobject.CharacterID]*Character {
	result := make(map[valueobject.CharacterID]*Character)
	for k, v := range s.characters {
		result[k] = v
	}
	return result
}

func (s *Scenario) SetStartScene(sceneID valueobject.SceneID) {
	s.startScene = sceneID
}