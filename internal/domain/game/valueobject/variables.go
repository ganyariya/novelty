package valueobject

type Variables struct {
	vars map[string]int
}

func NewVariables() Variables {
	return Variables{
		vars: make(map[string]int),
	}
}

func NewVariablesFromMap(vars map[string]int) Variables {
	copiedVars := make(map[string]int)
	for k, v := range vars {
		copiedVars[k] = v
	}
	return Variables{vars: copiedVars}
}

func (v Variables) Get(name string) int {
	value, exists := v.vars[name]
	if exists {
		return value
	}
	return 0
}

func (v Variables) Set(name string, value int) Variables {
	newVars := make(map[string]int)
	for k, val := range v.vars {
		newVars[k] = val
	}
	newVars[name] = value
	return Variables{vars: newVars}
}

func (v Variables) Add(name string, value int) Variables {
	current := v.Get(name)
	return v.Set(name, current+value)
}

func (v Variables) ToMap() map[string]int {
	result := make(map[string]int)
	for k, val := range v.vars {
		result[k] = val
	}
	return result
}

func (v Variables) Has(name string) bool {
	_, exists := v.vars[name]
	return exists
}