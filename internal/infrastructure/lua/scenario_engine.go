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
	executionState  *ScriptExecutionState
}

func NewScenarioEngine(scriptDir string) *ScenarioEngine {
	L := lua.NewState()
	
	engine := &ScenarioEngine{
		state:          L,
		scriptDir:      scriptDir,
		executionState: NewScriptExecutionState(),
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
	fmt.Printf("[DEBUG] ExecuteFunction called: %s\n", functionName)
	
	if e.gameState == nil {
		fmt.Printf("[DEBUG] Game state is nil\n")
		return errors.NewDomainError("GAME_STATE_NOT_SET", "Game state not set", nil)
	}
	
	fn := e.state.GetGlobal(functionName)
	if fn == lua.LNil {
		fmt.Printf("[DEBUG] Function %s not found in Lua state\n", functionName)
		return errors.WrapLuaExecutionFailed(functionName, fmt.Errorf("function %s not found", functionName))
	}
	
	fmt.Printf("[DEBUG] Executing Lua function: %s\n", functionName)
	e.state.Push(fn)
	if err := e.state.PCall(0, 0, nil); err != nil {
		fmt.Printf("[DEBUG] Lua function execution failed: %v\n", err)
		return errors.WrapLuaExecutionFailed(functionName, err)
	}
	
	fmt.Printf("[DEBUG] Lua function executed successfully\n")
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

func (e *ScenarioEngine) GetExecutionState() *ScriptExecutionState {
	return e.executionState
}