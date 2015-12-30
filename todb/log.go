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

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package todb

import (
	"encoding/json"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"time"
)

type Log struct {
	Id          int64       `db:"id" json:"id"`
	Level       null.String `db:"level" json:"level"`
	Message     string      `db:"message" json:"message"`
	TmUser      int64       `db:"tm_user" json:"tmUser"`
	Ticketnum   null.String `db:"ticketnum" json:"ticketnum"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleLog(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getLog(id)
	} else if method == "POST" {
		return postLog(payload)
	} else if method == "PUT" {
		return putLog(id, payload)
	} else if method == "DELETE" {
		return delLog(id)
	}
	return nil, nil
}

func getLog(id int) (interface{}, error) {
	ret := []Log{}
	if id >= 0 {
		err := globalDB.Select(&ret, "select * from log where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		queryStr := "select * from log"
		err := globalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postLog(payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO log("
	sqlString += "level"
	sqlString += ",message"
	sqlString += ",tm_user"
	sqlString += ",ticketnum"
	sqlString += ") VALUES ("
	sqlString += ":level"
	sqlString += ",:message"
	sqlString += ",:tm_user"
	sqlString += ",:ticketnum"
	sqlString += ")"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putLog(id int, payload []byte) (interface{}, error) {
	// Note this depends on the json having the correct id!
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE log SET "
	sqlString += "level = :level"
	sqlString += ",message = :message"
	sqlString += ",tm_user = :tm_user"
	sqlString += ",ticketnum = :ticketnum"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delLog(id int) (interface{}, error) {
	result, err := globalDB.Exec("DELETE FROM log WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}