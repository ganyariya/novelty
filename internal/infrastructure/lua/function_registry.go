package lua

import (
	"fmt"
	"time"

	"github.com/ganyariya/novelty/internal/domain/scenario/entity"
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
	lua "github.com/yuin/gopher-lua"
)

type FunctionRegistry struct {
	luaState *lua.LState
	engine   *ScenarioEngine
}

func NewFunctionRegistry(L *lua.LState, engine *ScenarioEngine) *FunctionRegistry {
	registry := &FunctionRegistry{
		luaState: L,
		engine:   engine,
	}
	
	registry.registerFunctions()
	return registry
}

func (r *FunctionRegistry) registerFunctions() {
	r.luaState.SetGlobal("text", r.luaState.NewFunction(r.text))
	r.luaState.SetGlobal("narration", r.luaState.NewFunction(r.narration))
	r.luaState.SetGlobal("nvl_text", r.luaState.NewFunction(r.nvlText))
	r.luaState.SetGlobal("choice", r.luaState.NewFunction(r.choice))
	r.luaState.SetGlobal("jump_to_scene", r.luaState.NewFunction(r.jumpToScene))
	r.luaState.SetGlobal("jump_to_label", r.luaState.NewFunction(r.jumpToLabel))
	r.luaState.SetGlobal("set_flag", r.luaState.NewFunction(r.setFlag))
	r.luaState.SetGlobal("get_flag", r.luaState.NewFunction(r.getFlag))
	r.luaState.SetGlobal("set_var", r.luaState.NewFunction(r.setVar))
	r.luaState.SetGlobal("get_var", r.luaState.NewFunction(r.getVar))
	r.luaState.SetGlobal("add_var", r.luaState.NewFunction(r.addVar))
	r.luaState.SetGlobal("set_display_mode", r.luaState.NewFunction(r.setDisplayMode))
	r.luaState.SetGlobal("register_character", r.luaState.NewFunction(r.registerCharacter))
	r.luaState.SetGlobal("play_voice", r.luaState.NewFunction(r.playVoice))
	r.luaState.SetGlobal("voice_text", r.luaState.NewFunction(r.voiceText))
}

func (r *FunctionRegistry) text(L *lua.LState) int {
	characterID := L.CheckString(1)
	content := L.CheckString(2)
	
	// デバッグ情報
	fmt.Printf("[DEBUG] text called: %s: %s\n", characterID, content)
	
	gameState := r.engine.GetGameState()
	if gameState == nil {
		fmt.Printf("[DEBUG] game state is nil\n")
		L.Push(lua.LBool(false))
		L.Push(lua.LString("game state not set"))
		return 2
	}
	
	message := entity.NewMessage(
		entity.NewMessageID(fmt.Sprintf("msg_%d", time.Now().UnixNano())),
		valueobject.NewCharacterID(characterID),
		content,
		valueobject.DisplayModeADV,
		"",
	)
	
	gameState.AddToHistory(message)
	fmt.Printf("[DEBUG] message added to history. Total messages: %d\n", len(gameState.History()))
	
	// メッセージキューシステムを使用
	r.engine.GetExecutionState().AddMessage(message)
	
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) narration(L *lua.LState) int {
	content := L.CheckString(1)
	
	fmt.Printf("[DEBUG] narration called: %s\n", content)
	
	gameState := r.engine.GetGameState()
	if gameState == nil {
		fmt.Printf("[DEBUG] game state is nil\n")
		L.Push(lua.LBool(false))
		L.Push(lua.LString("game state not set"))
		return 2
	}
	
	message := entity.NewNarrationMessage(
		entity.NewMessageID(fmt.Sprintf("msg_%d", time.Now().UnixNano())),
		content,
		valueobject.DisplayModeADV,
	)
	
	gameState.AddToHistory(message)
	fmt.Printf("[DEBUG] narration message added to history. Total messages: %d\n", len(gameState.History()))
	
	// メッセージキューシステムを使用
	r.engine.GetExecutionState().AddMessage(message)
	
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) nvlText(L *lua.LState) int {
	content := L.CheckString(1)
	
	gameState := r.engine.GetGameState()
	if gameState == nil {
		L.Push(lua.LBool(false))
		L.Push(lua.LString("game state not set"))
		return 2
	}
	
	message := entity.NewNarrationMessage(
		entity.NewMessageID(fmt.Sprintf("msg_%d", time.Now().UnixNano())),
		content,
		valueobject.DisplayModeNVL,
	)
	
	gameState.AddToHistory(message)
	
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) choice(L *lua.LState) int {
	// TODO: Implement choice system
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) jumpToScene(L *lua.LState) int {
	scenePath := L.CheckString(1)
	functionName := L.CheckString(2)
	
	gameState := r.engine.GetGameState()
	if gameState == nil {
		L.Push(lua.LBool(false))
		L.Push(lua.LString("game state not set"))
		return 2
	}
	
	sceneID := valueobject.NewSceneIDFromPath(scenePath)
	gameState.SetCurrentScene(sceneID, functionName)
	
	// Load and execute the new scene
	if err := r.engine.LoadScript(scenePath); err != nil {
		L.Push(lua.LBool(false))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	
	if err := r.engine.ExecuteFunction(functionName); err != nil {
		L.Push(lua.LBool(false))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) jumpToLabel(L *lua.LState) int {
	functionName := L.CheckString(1)
	
	gameState := r.engine.GetGameState()
	if gameState == nil {
		L.Push(lua.LBool(false))
		L.Push(lua.LString("game state not set"))
		return 2
	}
	
	gameState.SetCurrentScene(gameState.CurrentScene(), functionName)
	
	// Execute the function in current scene
	if err := r.engine.ExecuteFunction(functionName); err != nil {
		L.Push(lua.LBool(false))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) setFlag(L *lua.LState) int {
	name := L.CheckString(1)
	value := L.CheckBool(2)
	
	gameState := r.engine.GetGameState()
	if gameState == nil {
		L.Push(lua.LBool(false))
		return 1
	}
	
	gameState.SetFlag(name, value)
	
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) getFlag(L *lua.LState) int {
	name := L.CheckString(1)
	
	gameState := r.engine.GetGameState()
	if gameState == nil {
		L.Push(lua.LBool(false))
		return 1
	}
	
	value := gameState.GetFlag(name)
	L.Push(lua.LBool(value))
	return 1
}

func (r *FunctionRegistry) setVar(L *lua.LState) int {
	name := L.CheckString(1)
	value := int(L.CheckNumber(2))
	
	gameState := r.engine.GetGameState()
	if gameState == nil {
		L.Push(lua.LBool(false))
		return 1
	}
	
	gameState.SetVariable(name, value)
	
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) getVar(L *lua.LState) int {
	name := L.CheckString(1)
	
	gameState := r.engine.GetGameState()
	if gameState == nil {
		L.Push(lua.LNumber(0))
		return 1
	}
	
	value := gameState.GetVariable(name)
	L.Push(lua.LNumber(float64(value)))
	return 1
}

func (r *FunctionRegistry) addVar(L *lua.LState) int {
	name := L.CheckString(1)
	value := int(L.CheckNumber(2))
	
	gameState := r.engine.GetGameState()
	if gameState == nil {
		L.Push(lua.LBool(false))
		return 1
	}
	
	gameState.AddVariable(name, value)
	
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) setDisplayMode(L *lua.LState) int {
	mode := L.CheckString(1)
	
	// TODO: Implement display mode switching
	_ = valueobject.NewDisplayModeFromString(mode)
	
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) registerCharacter(L *lua.LState) int {
	// TODO: Implement character registration
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) playVoice(L *lua.LState) int {
	// TODO: Implement voice playback (future feature)
	L.Push(lua.LBool(true))
	return 1
}

func (r *FunctionRegistry) voiceText(L *lua.LState) int {
	characterID := L.CheckString(1)
	content := L.CheckString(2)
	voiceFile := L.CheckString(3)
	
	gameState := r.engine.GetGameState()
	if gameState == nil {
		L.Push(lua.LBool(false))
		L.Push(lua.LString("game state not set"))
		return 2
	}
	
	message := entity.NewMessage(
		entity.NewMessageID(fmt.Sprintf("msg_%d", time.Now().UnixNano())),
		valueobject.NewCharacterID(characterID),
		content,
		valueobject.DisplayModeADV,
		voiceFile,
	)
	
	gameState.AddToHistory(message)
	
	// TODO: Play voice file
	
	L.Push(lua.LBool(true))
	return 1
}