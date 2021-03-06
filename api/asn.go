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
	"encoding/json"
	"github.com/Comcast/traffic_control/traffic_ops/goto2/db"
	_ "github.com/Comcast/traffic_control/traffic_ops/goto2/output_format" // needed for swagger
	"log"
	"time"
)

type Asn struct {
	Id          int64     `db:"id" json:"id"`
	Asn         int64     `db:"asn" json:"asn"`
	Cachegroup  int64     `db:"cachegroup" json:"cachegroup"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleAsn(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getAsn(id)
	} else if method == "POST" {
		return postAsn(payload)
	} else if method == "PUT" {
		return putAsn(id, payload)
	} else if method == "DELETE" {
		return delAsn(id)
	}
	return nil, nil
}

func getAsn(id int) (interface{}, error) {
	if id >= 0 {
		return getAsnById(id)
	} else {
		return getAsns()
	}
}

// @Title getAsnById
// @Description retrieves the asn information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Asn
// @Resource /api/2.0
// @Router /api/2.0/asn/{id} [get]
func getAsnById(id int) (interface{}, error) {
	ret := []Asn{}
	arg := Asn{Id: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from asn where id=:id`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getAsns
// @Description retrieves the asn information for a certain id
// @Accept  application/json
// @Success 200 {array}    Asn
// @Resource /api/2.0
// @Router /api/2.0/asn [get]
func getAsns() (interface{}, error) {
	ret := []Asn{}
	queryStr := "select * from asn"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postAsn
// @Description enter a new asn
// @Accept  application/json
// @Param                 Body body     Asn   true "Asn object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/asn [post]
func postAsn(payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO asn("
	sqlString += "asn"
	sqlString += ",cachegroup"
	sqlString += ") VALUES ("
	sqlString += ":asn"
	sqlString += ",:cachegroup"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putAsn
// @Description modify an existing asnentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Asn   true "Asn object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/asn/{id}  [put]
func putAsn(id int, payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE asn SET "
	sqlString += "asn = :asn"
	sqlString += ",cachegroup = :cachegroup"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delAsnById
// @Description deletes asn information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Asn
// @Resource /api/2.0
// @Router /api/2.0/asn/{id} [delete]
func delAsn(id int) (interface{}, error) {
	arg := Asn{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM asn WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
