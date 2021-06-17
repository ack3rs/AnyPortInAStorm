package database

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/acky666/AnyPortInAStorm/PortController"
	"github.com/acky666/AnyPortInAStorm/PortDomainServiceServer/config"
	l "github.com/acky666/ackyLog"
	_ "github.com/go-sql-driver/mysql"
)

func Save(portName string, pi PortController.PortInfo) (bool, error) {
	defer l.TIMED("Saving %s to Database", portName)()

	Alias, _ := json.Marshal(pi.Alias)
	Regions, _ := json.Marshal(pi.Regions)
	Unlocs, _ := json.Marshal(pi.Unlocs)

	sqlQuery := "REPLACE INTO portnames SET portname=?,name=?,city=?,country=?,alias=?,regions=?,lat=?,lon=?,province=?,timezone=?,unlocs=?,code=?"

	// Cords may be blank.
	if len(pi.Coordinates) == 0 {
		pi.Coordinates = append(pi.Coordinates, -1)
		pi.Coordinates = append(pi.Coordinates, -1)
	}

	db, err := sql.Open("mysql", config.Get("DatabaseDSN"))
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		return false, err
	}
	defer db.Close()

	_, err = db.Exec(sqlQuery, portName, pi.Name, pi.City, pi.Country, Alias, Regions, pi.Coordinates[0], pi.Coordinates[1], pi.Province, pi.Timezone, Unlocs, pi.Code)

	if err != nil {
		return false, err
	}

	//LastInsertedID, _ := Result.LastInsertId()
	//RowsAffected, _ := Result.RowsAffected()
	//l.INFO("Executed - LastID %v and Rows Effected %v", LastInsertedID, RowsAffected)

	return true, nil
}
