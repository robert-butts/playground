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

package api

import (
	"../db"
	"encoding/json"
	"fmt"
	"time"
)

type FederationFederationResolver struct {
	Federation         int64     `db:"federation" json:"federation"`
	FederationResolver int64     `db:"federation_resolver" json:"federationResolver"`
	LastUpdated        time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleFederationFederationResolver(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getFederationFederationResolver(id)
	} else if method == "POST" {
		return postFederationFederationResolver(payload)
	} else if method == "PUT" {
		return putFederationFederationResolver(id, payload)
	} else if method == "DELETE" {
		return delFederationFederationResolver(id)
	}
	return nil, nil
}

func getFederationFederationResolver(id int) (interface{}, error) {
	ret := []FederationFederationResolver{}
	if id >= 0 {
		err := db.GlobalDB.Select(&ret, "select * from federation_federation_resolver where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		queryStr := "select * from federation_federation_resolver"
		err := db.GlobalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postFederationFederationResolver(payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO federation_federation_resolver("
	sqlString += "federation"
	sqlString += ",federation_resolver"
	sqlString += ") VALUES ("
	sqlString += ":federation"
	sqlString += ",:federation_resolver"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putFederationFederationResolver(id int, payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwirte the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE federation_federation_resolver SET "
	sqlString += "federation = :federation"
	sqlString += ",federation_resolver = :federation_resolver"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delFederationFederationResolver(id int) (interface{}, error) {
	result, err := db.GlobalDB.Exec("DELETE FROM federation_federation_resolver WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}