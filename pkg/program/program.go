package program

import tea "github.com/charmbracelet/bubbletea"

type Settings struct {
	ToggleKey    string `json:"toggle_key"`
	Delay        int    `json:"delay"`
	Hotkey       string `json:"hotkey"`
	OutlineColor uint32 `json:"outline_color"`
	AutoStart    bool   `json:"auto_start"`
}

type Pixel struct {
	Settings     *Settings
	Action       string
	SettingsView tea.Model
}

func NewProgram(settings *Settings) *Pixel {
	return &Pixel{
		Settings: settings,
	}
}

func (p *Pixel) CurrentAction() string {
	return p.Action
}

func (p *Pixel) SetAction(action string) {
	p.Action = action
}

func (s *Settings) GetToggleKey() string {
	return s.ToggleKey
}

func (s *Settings) LoadSettingsFromJson() {

}
