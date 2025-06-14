package persistence

import (
	"context"
	"sync"

	"github.com/ganyariya/novelty/internal/domain/game/entity"
	scenarioEntity "github.com/ganyariya/novelty/internal/domain/scenario/entity"
	"github.com/ganyariya/novelty/internal/domain/scenario/repository"
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
	"github.com/ganyariya/novelty/internal/infrastructure/lua"
	"github.com/ganyariya/novelty/pkg/errors"
)

type FileScenarioRepository struct {
	engine    *lua.ScenarioEngine
	scenes    map[valueobject.SceneID]*scenarioEntity.Scene
	mu        sync.RWMutex
	scriptDir string
}

func NewFileScenarioRepository(scriptDir string) repository.ScenarioRepository {
	engine := lua.NewScenarioEngine(scriptDir)
	
	return &FileScenarioRepository{
		engine:    engine,
		scenes:    make(map[valueobject.SceneID]*scenarioEntity.Scene),
		scriptDir: scriptDir,
	}
}

func (r *FileScenarioRepository) LoadScene(ctx context.Context, sceneID valueobject.SceneID) (*scenarioEntity.Scene, error) {
	r.mu.RLock()
	if scene, exists := r.scenes[sceneID]; exists && scene.IsLoaded() {
		r.mu.RUnlock()
		return scene, nil
	}
	r.mu.RUnlock()
	
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if scene, exists := r.scenes[sceneID]; exists && scene.IsLoaded() {
		return scene, nil
	}
	
	scriptPath := sceneID.String()
	if err := r.engine.LoadScript(scriptPath); err != nil {
		return nil, err
	}
	
	scene := scenarioEntity.NewScene(sceneID, scriptPath)
	scene.SetLoaded(true)
	r.scenes[sceneID] = scene
	
	return scene, nil
}

func (r *FileScenarioRepository) LoadCharacters(ctx context.Context) (map[valueobject.CharacterID]*scenarioEntity.Character, error) {
	// TODO: Implement character loading from common/characters.lua
	return make(map[valueobject.CharacterID]*scenarioEntity.Character), nil
}

func (r *FileScenarioRepository) ExecuteSceneFunction(ctx context.Context, scene *scenarioEntity.Scene, functionName string, gameState interface{}) error {
	if gs, ok := gameState.(*entity.GameState); ok {
		r.engine.SetGameState(gs)
	} else {
		return errors.NewDomainError("INVALID_GAME_STATE", "Invalid game state type", nil)
	}
	
	return r.engine.ExecuteFunction(functionName)
}

func (r *FileScenarioRepository) ReloadScene(ctx context.Context, sceneID valueobject.SceneID) (*scenarioEntity.Scene, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	scriptPath := sceneID.String()
	if err := r.engine.ReloadScript(scriptPath); err != nil {
		return nil, err
	}
	
	scene := scenarioEntity.NewScene(sceneID, scriptPath)
	scene.SetLoaded(true)
	r.scenes[sceneID] = scene
	
	return scene, nil
}