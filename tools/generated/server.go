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

type Server struct {
	Id             int64       `db:"id" json:"id"`
	HostName       string      `db:"host_name" json:"hostName"`
	DomainName     string      `db:"domain_name" json:"domainName"`
	TcpPort        null.Int    `db:"tcp_port" json:"tcpPort"`
	XmppId         null.String `db:"xmpp_id" json:"xmppId"`
	XmppPasswd     null.String `db:"xmpp_passwd" json:"xmppPasswd"`
	InterfaceName  string      `db:"interface_name" json:"interfaceName"`
	IpAddress      string      `db:"ip_address" json:"ipAddress"`
	IpNetmask      string      `db:"ip_netmask" json:"ipNetmask"`
	IpGateway      string      `db:"ip_gateway" json:"ipGateway"`
	Ip6Address     null.String `db:"ip6_address" json:"ip6Address"`
	Ip6Gateway     null.String `db:"ip6_gateway" json:"ip6Gateway"`
	InterfaceMtu   int64       `db:"interface_mtu" json:"interfaceMtu"`
	PhysLocation   int64       `db:"phys_location" json:"physLocation"`
	Rack           null.String `db:"rack" json:"rack"`
	Cachegroup     int64       `db:"cachegroup" json:"cachegroup"`
	Type           int64       `db:"type" json:"type"`
	Status         int64       `db:"status" json:"status"`
	UpdPending     int64       `db:"upd_pending" json:"updPending"`
	Profile        int64       `db:"profile" json:"profile"`
	CdnId          int64       `db:"cdn_id" json:"cdnId"`
	MgmtIpAddress  null.String `db:"mgmt_ip_address" json:"mgmtIpAddress"`
	MgmtIpNetmask  null.String `db:"mgmt_ip_netmask" json:"mgmtIpNetmask"`
	MgmtIpGateway  null.String `db:"mgmt_ip_gateway" json:"mgmtIpGateway"`
	IloIpAddress   null.String `db:"ilo_ip_address" json:"iloIpAddress"`
	IloIpNetmask   null.String `db:"ilo_ip_netmask" json:"iloIpNetmask"`
	IloIpGateway   null.String `db:"ilo_ip_gateway" json:"iloIpGateway"`
	IloUsername    null.String `db:"ilo_username" json:"iloUsername"`
	IloPassword    null.String `db:"ilo_password" json:"iloPassword"`
	RouterHostName null.String `db:"router_host_name" json:"routerHostName"`
	RouterPortName null.String `db:"router_port_name" json:"routerPortName"`
	LastUpdated    time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleServer(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getServer(id)
	} else if method == "POST" {
		return postServer(payload)
	} else if method == "PUT" {
		return putServer(id, payload)
	} else if method == "DELETE" {
		return delServer(id)
	}
	return nil, nil
}

func getServer(id int) (interface{}, error) {
	ret := []Server{}
	if id >= 0 {
		err := globalDB.Select(&ret, "select * from server where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		queryStr := "select * from server"
		err := globalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postServer(payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO server("
	sqlString += "host_name"
	sqlString += ",domain_name"
	sqlString += ",tcp_port"
	sqlString += ",xmpp_id"
	sqlString += ",xmpp_passwd"
	sqlString += ",interface_name"
	sqlString += ",ip_address"
	sqlString += ",ip_netmask"
	sqlString += ",ip_gateway"
	sqlString += ",ip6_address"
	sqlString += ",ip6_gateway"
	sqlString += ",interface_mtu"
	sqlString += ",phys_location"
	sqlString += ",rack"
	sqlString += ",cachegroup"
	sqlString += ",type"
	sqlString += ",status"
	sqlString += ",upd_pending"
	sqlString += ",profile"
	sqlString += ",cdn_id"
	sqlString += ",mgmt_ip_address"
	sqlString += ",mgmt_ip_netmask"
	sqlString += ",mgmt_ip_gateway"
	sqlString += ",ilo_ip_address"
	sqlString += ",ilo_ip_netmask"
	sqlString += ",ilo_ip_gateway"
	sqlString += ",ilo_username"
	sqlString += ",ilo_password"
	sqlString += ",router_host_name"
	sqlString += ",router_port_name"
	sqlString += ") VALUES ("
	sqlString += ":host_name"
	sqlString += ",:domain_name"
	sqlString += ",:tcp_port"
	sqlString += ",:xmpp_id"
	sqlString += ",:xmpp_passwd"
	sqlString += ",:interface_name"
	sqlString += ",:ip_address"
	sqlString += ",:ip_netmask"
	sqlString += ",:ip_gateway"
	sqlString += ",:ip6_address"
	sqlString += ",:ip6_gateway"
	sqlString += ",:interface_mtu"
	sqlString += ",:phys_location"
	sqlString += ",:rack"
	sqlString += ",:cachegroup"
	sqlString += ",:type"
	sqlString += ",:status"
	sqlString += ",:upd_pending"
	sqlString += ",:profile"
	sqlString += ",:cdn_id"
	sqlString += ",:mgmt_ip_address"
	sqlString += ",:mgmt_ip_netmask"
	sqlString += ",:mgmt_ip_gateway"
	sqlString += ",:ilo_ip_address"
	sqlString += ",:ilo_ip_netmask"
	sqlString += ",:ilo_ip_gateway"
	sqlString += ",:ilo_username"
	sqlString += ",:ilo_password"
	sqlString += ",:router_host_name"
	sqlString += ",:router_port_name"
	sqlString += ")"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putServer(id int, payload []byte) (interface{}, error) {
	// Note this depends on the json having the correct id!
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE server SET "
	sqlString += "host_name = :host_name"
	sqlString += ",domain_name = :domain_name"
	sqlString += ",tcp_port = :tcp_port"
	sqlString += ",xmpp_id = :xmpp_id"
	sqlString += ",xmpp_passwd = :xmpp_passwd"
	sqlString += ",interface_name = :interface_name"
	sqlString += ",ip_address = :ip_address"
	sqlString += ",ip_netmask = :ip_netmask"
	sqlString += ",ip_gateway = :ip_gateway"
	sqlString += ",ip6_address = :ip6_address"
	sqlString += ",ip6_gateway = :ip6_gateway"
	sqlString += ",interface_mtu = :interface_mtu"
	sqlString += ",phys_location = :phys_location"
	sqlString += ",rack = :rack"
	sqlString += ",cachegroup = :cachegroup"
	sqlString += ",type = :type"
	sqlString += ",status = :status"
	sqlString += ",upd_pending = :upd_pending"
	sqlString += ",profile = :profile"
	sqlString += ",cdn_id = :cdn_id"
	sqlString += ",mgmt_ip_address = :mgmt_ip_address"
	sqlString += ",mgmt_ip_netmask = :mgmt_ip_netmask"
	sqlString += ",mgmt_ip_gateway = :mgmt_ip_gateway"
	sqlString += ",ilo_ip_address = :ilo_ip_address"
	sqlString += ",ilo_ip_netmask = :ilo_ip_netmask"
	sqlString += ",ilo_ip_gateway = :ilo_ip_gateway"
	sqlString += ",ilo_username = :ilo_username"
	sqlString += ",ilo_password = :ilo_password"
	sqlString += ",router_host_name = :router_host_name"
	sqlString += ",router_port_name = :router_port_name"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delServer(id int) (interface{}, error) {
	result, err := globalDB.Exec("DELETE FROM server WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}