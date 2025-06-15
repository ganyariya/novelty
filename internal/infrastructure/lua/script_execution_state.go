package lua

import (
	"github.com/ganyariya/novelty/internal/domain/scenario/entity"
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
	"github.com/ganyariya/novelty/pkg/logger"
)

type ScriptExecutionState struct {
	currentScene    valueobject.SceneID
	currentFunction string
	isWaitingInput  bool
	MessageQueue    []*entity.Message // 公開フィールドに変更
	nextAction      func() error
}

func NewScriptExecutionState() *ScriptExecutionState {
	return &ScriptExecutionState{
		MessageQueue:   make([]*entity.Message, 0),
		isWaitingInput: false,
	}
}

func (s *ScriptExecutionState) AddMessage(message *entity.Message) {
	s.MessageQueue = append(s.MessageQueue, message)
	s.isWaitingInput = true
	logger.Debug("[DEBUG] Message queued: %s, waiting for input\n", message.Content())
}

func (s *ScriptExecutionState) GetNextMessage() *entity.Message {
	if len(s.MessageQueue) > 0 {
		message := s.MessageQueue[0]
		s.MessageQueue = s.MessageQueue[1:]
		logger.Debug("[DEBUG] Dequeued message: %s, remaining: %d", message.Content(), len(s.MessageQueue))

		if len(s.MessageQueue) == 0 {
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
	return len(s.MessageQueue) > 0
}

func (s *ScriptExecutionState) ContinueExecution() error {
	if s.nextAction != nil {
		logger.Debug("[DEBUG] Continuing execution with next action\n")
		err := s.nextAction()
		s.nextAction = nil
		return err
	}
	return nil
}

func (s *ScriptExecutionState) SetNextAction(action func() error) {
	s.nextAction = action
}
