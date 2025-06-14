package entity

import (
	"github.com/ganyariya/novelty/internal/domain/scenario/valueobject"
)

type ChoiceID struct {
	value string
}

func NewChoiceID(value string) ChoiceID {
	return ChoiceID{value: value}
}

func (c ChoiceID) String() string {
	return c.value
}

type Choice struct {
	id         ChoiceID
	text       string
	condition  string
	jumpTarget *valueobject.JumpTarget
}

func NewChoice(id ChoiceID, text string, condition string, jumpTarget *valueobject.JumpTarget) *Choice {
	return &Choice{
		id:         id,
		text:       text,
		condition:  condition,
		jumpTarget: jumpTarget,
	}
}

func (c *Choice) ID() ChoiceID {
	return c.id
}

func (c *Choice) Text() string {
	return c.text
}

func (c *Choice) Condition() string {
	return c.condition
}

func (c *Choice) JumpTarget() *valueobject.JumpTarget {
	return c.jumpTarget
}

func (c *Choice) HasCondition() bool {
	return c.condition != ""
}

func (c *Choice) HasJumpTarget() bool {
	return c.jumpTarget != nil && c.jumpTarget.IsValid()
}