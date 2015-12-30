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
	"time"
)

type Region struct {
	Id          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Division    int64     `db:"division" json:"division"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleRegion(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getRegion(id)
	} else if method == "POST" {
		return postRegion(payload)
	} else if method == "PUT" {
		return putRegion(id, payload)
	} else if method == "DELETE" {
		return delRegion(id)
	}
	return nil, nil
}

func getRegion(id int) (interface{}, error) {
	ret := []Region{}
	if id >= 0 {
		err := globalDB.Select(&ret, "select * from region where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		queryStr := "select * from region"
		err := globalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postRegion(payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO region("
	sqlString += "name"
	sqlString += ",division"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:division"
	sqlString += ")"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putRegion(id int, payload []byte) (interface{}, error) {
	// Note this depends on the json having the correct id!
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE region SET "
	sqlString += "name = :name"
	sqlString += ",division = :division"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delRegion(id int) (interface{}, error) {
	result, err := globalDB.Exec("DELETE FROM region WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}