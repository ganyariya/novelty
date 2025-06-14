package valueobject

type CharacterID struct {
	value string
}

func NewCharacterID(id string) CharacterID {
	return CharacterID{value: id}
}

func (c CharacterID) String() string {
	return c.value
}

func (c CharacterID) Value() string {
	return c.value
}

func (c CharacterID) IsEmpty() bool {
	return c.value == ""
}

func (c CharacterID) Equals(other CharacterID) bool {
	return c.value == other.value
}