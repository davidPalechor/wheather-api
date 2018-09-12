package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateRequestModel_20180912_110831 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateRequestModel_20180912_110831{}
	m.Created = "20180912_110831"

	migration.Register("CreateRequestModel_20180912_110831", m)
}

// Run the migrations
func (m *CreateRequestModel_20180912_110831) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE IF NOT EXISTS 'request' (
				'id' integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
				'location_name' varchar(255) NOT NULL DEFAULT '' ,
				'temperature' varchar(255) NOT NULL DEFAULT '' ,
				'wind' varchar(255) NOT NULL DEFAULT '' ,
				'pressure' varchar(255) NOT NULL DEFAULT '' ,
				'humidity' varchar(255) NOT NULL DEFAULT '' ,
				'sunrise' datetime NOT NULL,
				'sunset' datetime NOT NULL,
				'long' double precision NOT NULL DEFAULT 0 ,
				'lat' double precision NOT NULL DEFAULT 0 ,
				'requested_time' datetime NOT NULL,
				'timestamp' datetime NOT NULL
			) ENGINE=InnoDB;`)

}

// Reverse the migrations
func (m *CreateRequestModel_20180912_110831) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE 'request'")
}
