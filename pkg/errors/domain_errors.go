package errors

import "fmt"

type DomainError struct {
	Code    string
	Message string
	Cause   error
}

func (e *DomainError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *DomainError) Unwrap() error {
	return e.Cause
}

func NewDomainError(code, message string, cause error) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

var (
	ErrSceneNotFound = &DomainError{
		Code:    "SCENE_NOT_FOUND",
		Message: "Scene not found",
	}

	ErrCharacterNotFound = &DomainError{
		Code:    "CHARACTER_NOT_FOUND",
		Message: "Character not found",
	}

	ErrInvalidSceneID = &DomainError{
		Code:    "INVALID_SCENE_ID",
		Message: "Invalid scene ID",
	}

	ErrInvalidCharacterID = &DomainError{
		Code:    "INVALID_CHARACTER_ID",
		Message: "Invalid character ID",
	}

	ErrInvalidJumpTarget = &DomainError{
		Code:    "INVALID_JUMP_TARGET",
		Message: "Invalid jump target",
	}

	ErrLuaExecutionFailed = &DomainError{
		Code:    "LUA_EXECUTION_FAILED",
		Message: "Lua script execution failed",
	}

	ErrSceneLoadFailed = &DomainError{
		Code:    "SCENE_LOAD_FAILED",
		Message: "Failed to load scene",
	}

	ErrSaveDataNotFound = &DomainError{
		Code:    "SAVE_DATA_NOT_FOUND",
		Message: "Save data not found",
	}

	ErrInvalidSaveSlot = &DomainError{
		Code:    "INVALID_SAVE_SLOT",
		Message: "Invalid save slot",
	}

	ErrSaveFailed = &DomainError{
		Code:    "SAVE_FAILED",
		Message: "Failed to save game data",
	}

	ErrLoadFailed = &DomainError{
		Code:    "LOAD_FAILED",
		Message: "Failed to load game data",
	}
)

func WrapSceneNotFound(sceneID string, cause error) *DomainError {
	return &DomainError{
		Code:    "SCENE_NOT_FOUND",
		Message: fmt.Sprintf("Scene not found: %s", sceneID),
		Cause:   cause,
	}
}

func WrapCharacterNotFound(characterID string, cause error) *DomainError {
	return &DomainError{
		Code:    "CHARACTER_NOT_FOUND",
		Message: fmt.Sprintf("Character not found: %s", characterID),
		Cause:   cause,
	}
}

func WrapLuaExecutionFailed(functionName string, cause error) *DomainError {
	return &DomainError{
		Code:    "LUA_EXECUTION_FAILED",
		Message: fmt.Sprintf("Lua function execution failed: %s", functionName),
		Cause:   cause,
	}
}

func WrapSaveSlotError(slot int, cause error) *DomainError {
	return &DomainError{
		Code:    "INVALID_SAVE_SLOT",
		Message: fmt.Sprintf("Invalid save slot: %d", slot),
		Cause:   cause,
	}
}