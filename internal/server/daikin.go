package server

type DaikinState struct {
	Temperature int  `json:"temperature"`
	Mode        bool `json:"mode"`
	Timer       bool `json:"timer"`
	TimerDelay  int  `json:"timerDelay"`
	Power       bool `json:"power"`
	FanSpeed    int  `json:"fanSpeed"`
	Swing       bool `json:"swing"`
	Powerful    bool `json:"powerful"`
	Econo       bool `json:"econo"`
}
