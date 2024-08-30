package trigger

const (
	tolerence   = 10
	targetColor = "red"
)

func GenerateDefaultConfig() *Config {
	return &Config{
		Tolerence:   tolerence,
		TargetColor: targetColor,
		TriggerKey:  leftAltRawCode,
	}
}
