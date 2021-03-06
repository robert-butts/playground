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

type FederationDeliveryservice struct {
	Federation      int64     `db:"federation" json:"federation"`
	Deliveryservice int64     `db:"deliveryservice" json:"deliveryservice"`
	LastUpdated     time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleFederationDeliveryservice(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getFederationDeliveryservice(id)
	} else if method == "POST" {
		return postFederationDeliveryservice(payload)
	} else if method == "PUT" {
		return putFederationDeliveryservice(id, payload)
	} else if method == "DELETE" {
		return delFederationDeliveryservice(id)
	}
	return nil, nil
}

func getFederationDeliveryservice(id int) (interface{}, error) {
	if id >= 0 {
		return getFederationDeliveryserviceById(id)
	} else {
		return getFederationDeliveryservices()
	}
}

// @Title getFederationDeliveryserviceById
// @Description retrieves the federation_deliveryservice information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    FederationDeliveryservice
// @Resource /api/2.0
// @Router /api/2.0/federation_deliveryservice/{id} [get]
func getFederationDeliveryserviceById(id int) (interface{}, error) {
	ret := []FederationDeliveryservice{}
	arg := FederationDeliveryservice{Federation: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from federation_deliveryservice where federation=:federation`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getFederationDeliveryservices
// @Description retrieves the federation_deliveryservice information for a certain id
// @Accept  application/json
// @Success 200 {array}    FederationDeliveryservice
// @Resource /api/2.0
// @Router /api/2.0/federation_deliveryservice [get]
func getFederationDeliveryservices() (interface{}, error) {
	ret := []FederationDeliveryservice{}
	queryStr := "select * from federation_deliveryservice"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postFederationDeliveryservice
// @Description enter a new federation_deliveryservice
// @Accept  application/json
// @Param                 Body body     FederationDeliveryservice   true "FederationDeliveryservice object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federation_deliveryservice [post]
func postFederationDeliveryservice(payload []byte) (interface{}, error) {
	var v FederationDeliveryservice
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO federation_deliveryservice("
	sqlString += "federation"
	sqlString += ",deliveryservice"
	sqlString += ") VALUES ("
	sqlString += ":federation"
	sqlString += ",:deliveryservice"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putFederationDeliveryservice
// @Description modify an existing federation_deliveryserviceentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     FederationDeliveryservice   true "FederationDeliveryservice object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federation_deliveryservice/{id}  [put]
func putFederationDeliveryservice(id int, payload []byte) (interface{}, error) {
	var v FederationDeliveryservice
	err := json.Unmarshal(payload, &v)
	v.Federation = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE federation_deliveryservice SET "
	sqlString += "federation = :federation"
	sqlString += ",deliveryservice = :deliveryservice"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE federation=:federation"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delFederationDeliveryserviceById
// @Description deletes federation_deliveryservice information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    FederationDeliveryservice
// @Resource /api/2.0
// @Router /api/2.0/federation_deliveryservice/{id} [delete]
func delFederationDeliveryservice(id int) (interface{}, error) {
	arg := FederationDeliveryservice{Federation: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM federation_deliveryservice WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
