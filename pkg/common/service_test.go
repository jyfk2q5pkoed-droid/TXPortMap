package common

import (
	ejson "encoding/json"
	"github.com/4dogs-cn/TXPortMap/pkg/common/rangectl"
	"github.com/4dogs-cn/TXPortMap/pkg/output"
	"testing"
)

func TestComparePacketsMysql(t *testing.T) {
	banner := []byte(">\x00\x00\x00\x0a5.0.51a-3ubuntu5\x00\x0e\x00\x00\x00pf.Q.2Mn\x00,ª\x08\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00c'4pXG<56Oh?\x00\x10\x00\x00\x01ÿ\x13\x04Bad handshake")

	size := len(banner)
	var szBan string
	var szSvcName string

	num := ComparePackets(banner, size, &szBan, &szSvcName)

	if num == 0 {
		t.Error("unknown service")
		return
	}

	t.Log(szBan)
	t.Log(szSvcName)
}

func TestScan(t *testing.T) {
	handle := func(event *output.ResultEvent) {
		eventjs, _ := ejson.Marshal(event)
		t.Log(string(eventjs))
	}
	//var err error
	//Writer, err = output.NewStandardWriter(nocolor, json, rstfile, tracelog)
	//if err != nil {
	//	//return err
	//}
	//ScannerWithHandle("127.0.0.1", 3306, handle)
	e := CreateEngine()
	e.SetHandle(handle)
	//en.TaskIps =
	re, _ := rangectl.ParseIpv4Range("192.168.20.199")
	e.TaskIps = append(e.TaskIps, re)
	p, _ := rangectl.ParsePortRange("19912")
	e.TaskPorts = append(e.TaskPorts, p)
	e.RunWithHandle()
	e.Wg.Wait()
	//if Writer != nil {
	//	Writer.Close()
	//}
}
