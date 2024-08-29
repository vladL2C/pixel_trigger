package trigger

const (
	tolerence   = 50
	targetColor = "red"
)

func GenerateDefaultConfig() *Config {
	return &Config{
		Tolerence:   tolerence,
		TargetColor: targetColor,
		TriggerKey:  leftAltRawCode,
	}
}
