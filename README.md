# snmp-go-example

## 介绍
  
  snmp网段扫描接口列表
  支持单个ip或者ip范围或者一个网段，目前只支持/24网段

## 使用
   `go build`
   `./snmp-go-example -n 192.168.1.222` 
      or
   `./snmp-go-example -n 192.168.1.222-192.168.1.228`
      or
   `./snmp-go-example -n 192.168.1.0/24`      
   
