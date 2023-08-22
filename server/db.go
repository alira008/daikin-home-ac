package server

import (
	"database/sql"
	"log"
    _ "github.com/mattn/go-sqlite3"
)

type Database struct {
	session *sql.DB
}

func (database *Database) LoadState() *DaikinState {
	statement, err := database.session.Prepare("Select Temperature, Mode, TimerState, TimerDelay, PowerState, FanSpeed, SwingState, PowerfulState, EconoState, ComfortState from State")

	// Error creating statement
	if err != nil {
		log.Fatal(err)
	}

	results, err := statement.Query()

	// Error executing query
	if err != nil {
		log.Fatal(err)
	}

	daikinState := DaikinState{}

	for results.Next() {
		err := results.Scan(&daikinState.Temperature, &daikinState.Mode, &daikinState.Timer, &daikinState.TimerDelay, &daikinState.Power, &daikinState.FanSpeed, &daikinState.Swing, &daikinState.Powerful, &daikinState.Econo, &daikinState.Comfort)

		// Error getting columns
		if err != nil {
			log.Fatal(err)
		}
	}

	defer statement.Close()

	return &daikinState
}

func (database *Database) SaveState(state *DaikinState) {
	statement, err := database.session.Prepare("UPDATE State SET Temperature = ?, Mode = ?, TimerState = ?, TimerDelay = ?, PowerState = ?, FanSpeed = ?, SwingState = ?, PowerfulState = ?, EconoState = ?, ComfortState = ?")

	// Error creating statement
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(state.Temperature, state.Mode, state.Timer, state.TimerDelay, state.Power, state.FanSpeed, state.Swing, state.Powerful, state.Econo, state.Comfort)

	// Error executing query
	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()
}

func OpenDatabase() *Database {
	db, err := sql.Open("sqlite3", "./daikin.db")

	if err != nil {
        log.Fatal(err)
	}

	return &Database{
		session: db,
	}
}
