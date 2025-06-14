package valueobject

import "fmt"

type SaveSlot struct {
	value int
}

func NewSaveSlot(slot int) SaveSlot {
	return SaveSlot{value: slot}
}

func (s SaveSlot) Value() int {
	return s.value
}

func (s SaveSlot) String() string {
	return fmt.Sprintf("slot_%d", s.value)
}

func (s SaveSlot) IsValid() bool {
	return s.value >= 1 && s.value <= 9
}

func (s SaveSlot) Equals(other SaveSlot) bool {
	return s.value == other.value
}