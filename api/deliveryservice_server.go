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

type DeliveryserviceServer struct {
	Deliveryservice int64     `db:"deliveryservice" json:"deliveryservice"`
	Server          int64     `db:"server" json:"server"`
	LastUpdated     time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleDeliveryserviceServer(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getDeliveryserviceServer(id)
	} else if method == "POST" {
		return postDeliveryserviceServer(payload)
	} else if method == "PUT" {
		return putDeliveryserviceServer(id, payload)
	} else if method == "DELETE" {
		return delDeliveryserviceServer(id)
	}
	return nil, nil
}

func getDeliveryserviceServer(id int) (interface{}, error) {
	if id >= 0 {
		return getDeliveryserviceServerById(id)
	} else {
		return getDeliveryserviceServers()
	}
}

// @Title getDeliveryserviceServerById
// @Description retrieves the deliveryservice_server information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    DeliveryserviceServer
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_server/{id} [get]
func getDeliveryserviceServerById(id int) (interface{}, error) {
	ret := []DeliveryserviceServer{}
	arg := DeliveryserviceServer{Deliveryservice: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from deliveryservice_server where deliveryservice=:deliveryservice`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getDeliveryserviceServers
// @Description retrieves the deliveryservice_server information for a certain id
// @Accept  application/json
// @Success 200 {array}    DeliveryserviceServer
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_server [get]
func getDeliveryserviceServers() (interface{}, error) {
	ret := []DeliveryserviceServer{}
	queryStr := "select * from deliveryservice_server"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postDeliveryserviceServer
// @Description enter a new deliveryservice_server
// @Accept  application/json
// @Param                 Body body     DeliveryserviceServer   true "DeliveryserviceServer object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_server [post]
func postDeliveryserviceServer(payload []byte) (interface{}, error) {
	var v DeliveryserviceServer
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO deliveryservice_server("
	sqlString += "deliveryservice"
	sqlString += ",server"
	sqlString += ") VALUES ("
	sqlString += ":deliveryservice"
	sqlString += ",:server"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putDeliveryserviceServer
// @Description modify an existing deliveryservice_serverentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     DeliveryserviceServer   true "DeliveryserviceServer object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_server/{id}  [put]
func putDeliveryserviceServer(id int, payload []byte) (interface{}, error) {
	var v DeliveryserviceServer
	err := json.Unmarshal(payload, &v)
	v.Deliveryservice = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE deliveryservice_server SET "
	sqlString += "deliveryservice = :deliveryservice"
	sqlString += ",server = :server"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE deliveryservice=:deliveryservice"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delDeliveryserviceServerById
// @Description deletes deliveryservice_server information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    DeliveryserviceServer
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_server/{id} [delete]
func delDeliveryserviceServer(id int) (interface{}, error) {
	arg := DeliveryserviceServer{Deliveryservice: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM deliveryservice_server WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
