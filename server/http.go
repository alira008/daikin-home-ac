package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type HttpServer struct {
	Database *Database
}

func convertStrModeToEnum(strMode string) (Mode, error) {

	modeInt, err := strconv.Atoi(strMode)
	if err != nil {
		return Cold, err
	}

	var mode Mode

	if modeInt == int(Heat) {
		mode = Heat
	} else if modeInt == int(Fan) {
		mode = Fan
	} else if modeInt == int(Dry) {
		mode = Dry
	} else if modeInt == int(Auto) {
		mode = Auto
	} else {
		mode = Cold
	}

	return mode, err
}

func convertStrFanSpeedToEnum(strFanSpeed string) (FanSpeed, error) {

	fanSpeedInt, err := strconv.Atoi(strFanSpeed)
	if err != nil {
		return One, err
	}

	var fanSpeed FanSpeed

	if fanSpeedInt == int(Silent) {
		fanSpeed = Silent
	} else if fanSpeedInt == int(Two) {
		fanSpeed = Two
	} else if fanSpeedInt == int(Three) {
		fanSpeed = Three
	} else if fanSpeedInt == int(Four) {
		fanSpeed = Four
	} else if fanSpeedInt == int(Five) {
		fanSpeed = Five
	} else if fanSpeedInt == int(Automatic) {
		fanSpeed = Automatic
	} else {
		fanSpeed = One
	}

	return fanSpeed, err
}

func (hs *HttpServer) ChangeTemperature(w http.ResponseWriter, r *http.Request) {

	// retrieve state from db
	daikin := hs.Database.LoadState()
	vars := mux.Vars(r)
	strTemp := vars["temp"]

	temp, err := strconv.Atoi(strTemp)
	if err != nil {
		http.Error(w, "invalid mode passed", http.StatusBadRequest)
	}

	// change state
	daikin.Temperature = temp

	// send new state
	daikin.Send()

	// save state to db
	hs.Database.SaveState(daikin)
}

func (hs *HttpServer) ChangeMode(w http.ResponseWriter, r *http.Request) {

	// retrieve state from db
	daikin := hs.Database.LoadState()
	vars := mux.Vars(r)
	strMode := vars["state"]

	mode, err := convertStrModeToEnum(strMode)
	if err != nil {
		http.Error(w, "invalid mode passed", http.StatusBadRequest)
	}

	// change state
	daikin.Mode = mode

	// send new state
	daikin.Send()

	// save state to db
	hs.Database.SaveState(daikin)
}

func (hs *HttpServer) ChangeTimerState(w http.ResponseWriter, r *http.Request) {

	// retrieve state from db
	daikin := hs.Database.LoadState()
	vars := mux.Vars(r)
	timer := vars["timer"]
	strState := vars["state"]
	strTimeDiff := vars["timeDiff"]

	state, err := strconv.ParseBool(strState)
	if err != nil {
		http.Error(w, "invalid timer state passed", http.StatusBadRequest)
	}

	timeDiff, err := strconv.Atoi(strTimeDiff)
	if err != nil {
		http.Error(w, "invalid timer difference passed", http.StatusBadRequest)
	}

	// change state
	if timer == "on" {
		daikin.OnTimer = state
	} else if timer == "off" {
		daikin.OffTimer = state
	} else {
		http.Error(w, "invalid format for timer given", http.StatusBadRequest)
	}

	// set delay
	daikin.TimerDelay = timeDiff

	// send new state
	daikin.Send()

	// save state to db
	hs.Database.SaveState(daikin)
}

func (hs *HttpServer) ChangePowerState(w http.ResponseWriter, r *http.Request) {

	// retrieve state from db
	daikin := hs.Database.LoadState()
	vars := mux.Vars(r)
	strState := vars["state"]

	state, err := strconv.ParseBool(strState)
	if err != nil {
		http.Error(w, "invalid power state passed", http.StatusBadRequest)
	}

	// change state
	daikin.Power = state

	// send new state
	daikin.Send()

	// save state to db
	hs.Database.SaveState(daikin)
}

func (hs *HttpServer) ChangeFanSpeed(w http.ResponseWriter, r *http.Request) {

	// retrieve state from db
	daikin := hs.Database.LoadState()
	vars := mux.Vars(r)
	strFanSpeed := vars["state"]

	fanSpeed, err := convertStrFanSpeedToEnum(strFanSpeed)
	if err != nil {
		http.Error(w, "invalid fanspeed passed", http.StatusBadRequest)
	}

	// change state
	daikin.FanSpeed = fanSpeed

	// send new state
	daikin.Send()

	// save state to db
	hs.Database.SaveState(daikin)
}

func (hs *HttpServer) ChangeSwingState(w http.ResponseWriter, r *http.Request) {

	// retrieve state from db
	daikin := hs.Database.LoadState()

	vars := mux.Vars(r)
	strState := vars["state"]

	state, err := strconv.ParseBool(strState)
	if err != nil {
		http.Error(w, "invalid power state passed", http.StatusBadRequest)
	}

	// change state
	daikin.Swing = state

	// send new state
	daikin.Send()

	// save state to db
	hs.Database.SaveState(daikin)
}

func (hs *HttpServer) ChangePowerfulState(w http.ResponseWriter, r *http.Request) {

	// retrieve state from db
	daikin := hs.Database.LoadState()
	vars := mux.Vars(r)
	strState := vars["state"]

	state, err := strconv.ParseBool(strState)
	if err != nil {
		http.Error(w, "invalid power state passed", http.StatusBadRequest)
	}

	// change state
	daikin.Powerful = state

	// send new state
	daikin.Send()

	// save state to db
	hs.Database.SaveState(daikin)
}

func (hs *HttpServer) ChangeEconoState(w http.ResponseWriter, r *http.Request) {

	// retrieve state from db
	daikin := hs.Database.LoadState()
	vars := mux.Vars(r)
	strState := vars["state"]

	state, err := strconv.ParseBool(strState)
	if err != nil {
		http.Error(w, "invalid power state passed", http.StatusBadRequest)
	}

	// change state
	daikin.Econo = state

	// send new state
	daikin.Send()

	// save state to db
	hs.Database.SaveState(daikin)
}

func (hs *HttpServer) GetTemperature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	daikin := hs.Database.LoadState()

	json.NewEncoder(w).Encode(daikin.Temperature)
}

func NewHttpServer(addr string) *http.Server {
	database := OpenDatabase()
	server := &HttpServer{
		Database: database,
	}

	r := mux.NewRouter()
	r.HandleFunc("/temperature", server.GetTemperature).Methods("GET")
	r.HandleFunc("/temperature/{temp}", server.ChangeTemperature).Methods("GET")
	r.HandleFunc("/mode/{state}", server.ChangeMode).Methods("GET")
	r.HandleFunc("/timer/{timer}/{state}/{timeDiff}", server.ChangeTimerState).Methods("GET")
	r.HandleFunc("/power/{state}", server.ChangePowerState).Methods("GET")
	r.HandleFunc("/fanspeed/{state}", server.ChangeFanSpeed).Methods("GET")
	r.HandleFunc("/swing/{state}", server.ChangeSwingState).Methods("GET")
	r.HandleFunc("/powerful/{state}", server.ChangePowerfulState).Methods("GET")
	r.HandleFunc("/econo/{state}", server.ChangeEconoState).Methods("GET")

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET"}),
		handlers.AllowedHeaders([]string{"X-RequestedWith", "Content-Type", "Authorization", "Accept"}),
	)(r)

	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
