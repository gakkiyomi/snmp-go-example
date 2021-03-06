// snmp-go-example - Golang snmp case
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package main

import (
	"regexp"
	"sync"
	"time"

	"github.com/alouca/gosnmp"
	"github.com/gakkiyomi/galang/net"
	"github.com/songtianyi/rrframework/logs"
)

func Sacn(target string) {

	is_Range, _ := regexp.MatchString(net.RANGE_REG, target)
	is_Ip, _ := regexp.MatchString(net.IP_REG, target)
	is_Cidr, _ := regexp.MatchString(net.CIDR_REG, target)

	if is_Range == true {
		list, err := net.Network.GetRangeAddrList(target)
		if err != nil {
			logs.Error(err.Error())
		}
		for _, t := range list {
			start(t)
		}

		return
	}

	if is_Ip == true {
		start(target)
		return
	}

	if is_Cidr == true {
		var wg sync.WaitGroup
		list, err := net.Network.GetCIDRAvailableAddrList(target)
		if err != nil {
			logs.Error(err.Error())
		}
		wg.Add(len(list))
		for _, t := range list {
			go startSync(t, &wg)
		}
		wg.Wait()
		return
	}

}

func start(target string) {

	wrapper, err := net.SNMP.NewSnmpWrapper(target, "public", gosnmp.Version2c, 3)
	if err != nil {
		logs.Error(err.Error())
	}
	intf, err := wrapper.Interfaces()
	if err == nil {
		for _, v := range intf {
			time.Sleep(time.Duration(500) * time.Millisecond)
			logs.Info(v.ToString())
		}
	}
}

func startSync(target string, wg *sync.WaitGroup) {

	defer wg.Done()

	wrapper, err := net.SNMP.NewSnmpWrapper(target, "public", gosnmp.Version2c, 3)
	if err != nil {
		logs.Error(err.Error())
	}
	intf, err := wrapper.Interfaces()
	if err == nil {
		for _, v := range intf {
			time.Sleep(time.Duration(500) * time.Millisecond)
			logs.Info(v.ToString())
		}
	}
}
