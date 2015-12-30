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

type Federation struct {
	Id          int64       `db:"id" json:"id"`
	Cname       string      `db:"cname" json:"cname"`
	Description null.String `db:"description" json:"description"`
	Ttl         int64       `db:"ttl" json:"ttl"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleFederation(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getFederation(id)
	} else if method == "POST" {
		return postFederation(payload)
	} else if method == "PUT" {
		return putFederation(id, payload)
	} else if method == "DELETE" {
		return delFederation(id)
	}
	return nil, nil
}

func getFederation(id int) (interface{}, error) {
	ret := []Federation{}
	if id >= 0 {
		err := globalDB.Select(&ret, "select * from federation where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		queryStr := "select * from federation"
		err := globalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postFederation(payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO federation("
	sqlString += "cname"
	sqlString += ",description"
	sqlString += ",ttl"
	sqlString += ") VALUES ("
	sqlString += ":cname"
	sqlString += ",:description"
	sqlString += ",:ttl"
	sqlString += ")"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putFederation(id int, payload []byte) (interface{}, error) {
	// Note this depends on the json having the correct id!
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE federation SET "
	sqlString += "cname = :cname"
	sqlString += ",description = :description"
	sqlString += ",ttl = :ttl"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delFederation(id int) (interface{}, error) {
	result, err := globalDB.Exec("DELETE FROM federation WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}