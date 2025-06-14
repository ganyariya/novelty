package valueobject

type ColorTheme struct {
	name    string
	primary string
	accent  string
}

func NewColorTheme(name, primary, accent string) ColorTheme {
	return ColorTheme{
		name:    name,
		primary: primary,
		accent:  accent,
	}
}

func (c ColorTheme) Name() string {
	return c.name
}

func (c ColorTheme) Primary() string {
	return c.primary
}

func (c ColorTheme) Accent() string {
	return c.accent
}

func DefaultColorThemes() map[string]ColorTheme {
	return map[string]ColorTheme{
		"blue":   NewColorTheme("blue", "#0066cc", "#99ccff"),
		"red":    NewColorTheme("red", "#cc0000", "#ff9999"),
		"green":  NewColorTheme("green", "#008800", "#99ff99"),
		"purple": NewColorTheme("purple", "#8800cc", "#cc99ff"),
		"orange": NewColorTheme("orange", "#ff6600", "#ffcc99"),
		"cyan":   NewColorTheme("cyan", "#00cccc", "#99ffff"),
	}
}