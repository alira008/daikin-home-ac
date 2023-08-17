package server

import (
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	session *sql.DB
}

func (database *Database) loadState() *DaikinState {
	statement, err := database.session.Prepare("Select Temperature, Mode, Timer, TimerDelay, Power, FanSpeed, Swing, Powerful, Econo from State")

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
		err := results.Scan(&daikinState.Temperature, &daikinState.Mode, &daikinState.Timer, &daikinState.TimerDelay, &daikinState.Power, &daikinState.Power, &daikinState.FanSpeed, &daikinState.Swing, &daikinState.Powerful, &daikinState.Econo)

        // Error getting columns
		if err != nil {
			log.Fatal(err)
		}
	}

	defer statement.Close()

    return &daikinState
}

func (database *Database) saveState(state *DaikinState) {

}

func OpenDatabase() *Database {
	db, err := sql.Open("sqlite3", "./daikin.db")

	if err != nil {
		fmt.Println(err.Error())
	}

	return &Database{
		session: db,
	}
}

