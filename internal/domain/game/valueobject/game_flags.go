package valueobject

type GameFlags struct {
	flags map[string]bool
}

func NewGameFlags() GameFlags {
	return GameFlags{
		flags: make(map[string]bool),
	}
}

func NewGameFlagsFromMap(flags map[string]bool) GameFlags {
	copiedFlags := make(map[string]bool)
	for k, v := range flags {
		copiedFlags[k] = v
	}
	return GameFlags{flags: copiedFlags}
}

func (g GameFlags) Get(name string) bool {
	value, exists := g.flags[name]
	return exists && value
}

func (g GameFlags) Set(name string, value bool) GameFlags {
	newFlags := make(map[string]bool)
	for k, v := range g.flags {
		newFlags[k] = v
	}
	newFlags[name] = value
	return GameFlags{flags: newFlags}
}

func (g GameFlags) ToMap() map[string]bool {
	result := make(map[string]bool)
	for k, v := range g.flags {
		result[k] = v
	}
	return result
}

func (g GameFlags) Has(name string) bool {
	_, exists := g.flags[name]
	return exists
}