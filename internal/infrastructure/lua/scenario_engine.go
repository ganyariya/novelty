package lua

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ganyariya/novelty/internal/domain/game/entity"
	"github.com/ganyariya/novelty/pkg/errors"
	lua "github.com/yuin/gopher-lua"
)

type ScenarioEngine struct {
	state           *lua.LState
	scriptDir       string
	gameState       *entity.GameState
	functionRegistry *FunctionRegistry
}

func NewScenarioEngine(scriptDir string) *ScenarioEngine {
	L := lua.NewState()
	
	engine := &ScenarioEngine{
		state:     L,
		scriptDir: scriptDir,
	}
	
	engine.functionRegistry = NewFunctionRegistry(L, engine)
	
	return engine
}

func (e *ScenarioEngine) SetGameState(gameState *entity.GameState) {
	e.gameState = gameState
}

func (e *ScenarioEngine) LoadScript(scriptPath string) error {
	fullPath := filepath.Join(e.scriptDir, scriptPath)
	
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return errors.WrapSceneNotFound(scriptPath, err)
	}
	
	if err := e.state.DoString(string(content)); err != nil {
		return errors.WrapLuaExecutionFailed(scriptPath, err)
	}
	
	return nil
}

func (e *ScenarioEngine) ExecuteFunction(functionName string) error {
	if e.gameState == nil {
		return errors.NewDomainError("GAME_STATE_NOT_SET", "Game state not set", nil)
	}
	
	fn := e.state.GetGlobal(functionName)
	if fn == lua.LNil {
		return errors.WrapLuaExecutionFailed(functionName, fmt.Errorf("function %s not found", functionName))
	}
	
	e.state.Push(fn)
	if err := e.state.PCall(0, 0, nil); err != nil {
		return errors.WrapLuaExecutionFailed(functionName, err)
	}
	
	return nil
}

func (e *ScenarioEngine) ReloadScript(scriptPath string) error {
	fullPath := filepath.Join(e.scriptDir, scriptPath)
	
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return errors.WrapSceneNotFound(scriptPath, err)
	}
	
	if err := e.state.DoString(string(content)); err != nil {
		return errors.WrapLuaExecutionFailed(scriptPath, err)
	}
	
	return nil
}

func (e *ScenarioEngine) Close() {
	if e.state != nil {
		e.state.Close()
	}
}

func (e *ScenarioEngine) GetGameState() *entity.GameState {
	return e.gameState
}