package valueobject

type DisplayMode int

const (
	DisplayModeADV DisplayMode = iota
	DisplayModeNVL
	DisplayModeHybrid
)

func (d DisplayMode) String() string {
	switch d {
	case DisplayModeADV:
		return "adv"
	case DisplayModeNVL:
		return "nvl"
	case DisplayModeHybrid:
		return "hybrid"
	default:
		return "adv"
	}
}

func NewDisplayModeFromString(mode string) DisplayMode {
	switch mode {
	case "adv":
		return DisplayModeADV
	case "nvl":
		return DisplayModeNVL
	case "hybrid":
		return DisplayModeHybrid
	default:
		return DisplayModeADV
	}
}

func (d DisplayMode) IsValid() bool {
	return d >= DisplayModeADV && d <= DisplayModeHybrid
}