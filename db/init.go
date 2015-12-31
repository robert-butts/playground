// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	GlobalDB     sqlx.DB
	DatabaseName string
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func InitializeDatabase(dbtype, username, password, environment string) {
	connString := ""
	if dbtype == "mysql" {
		connString = username + ":" + password + "@tcp(localhost:3306)/" + environment + "?parseTime=True"
	} else if dbtype == "postgres" {
		connString = "user=" + username + " dbname=" + environment + " password=" + password + " sslmode=disable"
	}

	db, err := sqlx.Connect(dbtype, connString)
	check(err)
	GlobalDB = *db
	DatabaseName = environment
}