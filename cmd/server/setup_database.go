package main

import (
	"fmt"
	"herman-technical-julo/config"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type databases struct {
	julodb *gorm.DB
}

func buildDatabases() (*databases, error) {
	julodb, err := setupJULODBConnection()
	if err != nil {
		log.Printf("build JULO database error - %v\n", err)
		return nil, err
	}
	return &databases{
		julodb: julodb,
	}, nil
}

func setupJULODBConnection() (*gorm.DB, error) {
	juloConnStr := fmt.Sprintf("sqlserver://%s:%s@%s?database=JULO&encrypt=disable", config.Username(), config.Password(), config.DBURL())
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	db, err := gorm.Open(sqlserver.Open(juloConnStr), gormConfig)
	if err != nil {
		log.Printf("Failed init JULO - %v\n", err)
		return nil, err
	}

	log.Println("Connected to JULO Database")
	return db, nil
}
