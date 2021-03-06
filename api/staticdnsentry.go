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
	null "gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type Staticdnsentry struct {
	Id              int64     `db:"id" json:"id"`
	Host            string    `db:"host" json:"host"`
	Address         string    `db:"address" json:"address"`
	Type            int64     `db:"type" json:"type"`
	Ttl             int64     `db:"ttl" json:"ttl"`
	Deliveryservice int64     `db:"deliveryservice" json:"deliveryservice"`
	Cachegroup      null.Int  `db:"cachegroup" json:"cachegroup"`
	LastUpdated     time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleStaticdnsentry(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getStaticdnsentry(id)
	} else if method == "POST" {
		return postStaticdnsentry(payload)
	} else if method == "PUT" {
		return putStaticdnsentry(id, payload)
	} else if method == "DELETE" {
		return delStaticdnsentry(id)
	}
	return nil, nil
}

func getStaticdnsentry(id int) (interface{}, error) {
	if id >= 0 {
		return getStaticdnsentryById(id)
	} else {
		return getStaticdnsentrys()
	}
}

// @Title getStaticdnsentryById
// @Description retrieves the staticdnsentry information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Staticdnsentry
// @Resource /api/2.0
// @Router /api/2.0/staticdnsentry/{id} [get]
func getStaticdnsentryById(id int) (interface{}, error) {
	ret := []Staticdnsentry{}
	arg := Staticdnsentry{Id: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from staticdnsentry where id=:id`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getStaticdnsentrys
// @Description retrieves the staticdnsentry information for a certain id
// @Accept  application/json
// @Success 200 {array}    Staticdnsentry
// @Resource /api/2.0
// @Router /api/2.0/staticdnsentry [get]
func getStaticdnsentrys() (interface{}, error) {
	ret := []Staticdnsentry{}
	queryStr := "select * from staticdnsentry"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postStaticdnsentry
// @Description enter a new staticdnsentry
// @Accept  application/json
// @Param                 Body body     Staticdnsentry   true "Staticdnsentry object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/staticdnsentry [post]
func postStaticdnsentry(payload []byte) (interface{}, error) {
	var v Staticdnsentry
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO staticdnsentry("
	sqlString += "host"
	sqlString += ",address"
	sqlString += ",type"
	sqlString += ",ttl"
	sqlString += ",deliveryservice"
	sqlString += ",cachegroup"
	sqlString += ") VALUES ("
	sqlString += ":host"
	sqlString += ",:address"
	sqlString += ",:type"
	sqlString += ",:ttl"
	sqlString += ",:deliveryservice"
	sqlString += ",:cachegroup"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putStaticdnsentry
// @Description modify an existing staticdnsentryentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Staticdnsentry   true "Staticdnsentry object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/staticdnsentry/{id}  [put]
func putStaticdnsentry(id int, payload []byte) (interface{}, error) {
	var v Staticdnsentry
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE staticdnsentry SET "
	sqlString += "host = :host"
	sqlString += ",address = :address"
	sqlString += ",type = :type"
	sqlString += ",ttl = :ttl"
	sqlString += ",deliveryservice = :deliveryservice"
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

// @Title delStaticdnsentryById
// @Description deletes staticdnsentry information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Staticdnsentry
// @Resource /api/2.0
// @Router /api/2.0/staticdnsentry/{id} [delete]
func delStaticdnsentry(id int) (interface{}, error) {
	arg := Staticdnsentry{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM staticdnsentry WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
