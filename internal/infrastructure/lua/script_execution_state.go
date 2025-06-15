package lua

import (
	"fmt"

	"github.com/ganyariya/novelty/internal/domain/scenario/entity"
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type ScriptExecutionState struct {
	currentScene    valueobject.SceneID
	currentFunction string
	isWaitingInput  bool
	messageQueue    []*entity.Message
	nextAction      func() error
}

func NewScriptExecutionState() *ScriptExecutionState {
	return &ScriptExecutionState{
		messageQueue:   make([]*entity.Message, 0),
		isWaitingInput: false,
	}
}

func (s *ScriptExecutionState) AddMessage(message *entity.Message) {
	s.messageQueue = append(s.messageQueue, message)
	s.isWaitingInput = true
	fmt.Printf("[DEBUG] Message queued: %s, waiting for input\n", message.Content())
}

func (s *ScriptExecutionState) GetNextMessage() *entity.Message {
	if len(s.messageQueue) > 0 {
		message := s.messageQueue[0]
		s.messageQueue = s.messageQueue[1:]
		fmt.Printf("[DEBUG] Dequeued message: %s, remaining: %d\n", message.Content(), len(s.messageQueue))
		
		if len(s.messageQueue) == 0 {
			s.isWaitingInput = false
		}
		
		return message
	}
	return nil
}

func (s *ScriptExecutionState) IsWaitingInput() bool {
	return s.isWaitingInput
}

func (s *ScriptExecutionState) HasPendingMessages() bool {
	return len(s.messageQueue) > 0
}

func (s *ScriptExecutionState) ContinueExecution() error {
	if s.nextAction != nil {
		fmt.Printf("[DEBUG] Continuing execution with next action\n")
		err := s.nextAction()
		s.nextAction = nil
		return err
	}
	return nil
}

func (s *ScriptExecutionState) SetNextAction(action func() error) {
	s.nextAction = action
}