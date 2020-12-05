package jsonTest

import (
	"encoding/json"
	"testing"
)

func TestJson(t *testing.T) {
	type tcpClient struct {
		IPAddress  string
		PortNumber uint32
	}

	type tcpServer struct {
		IPAddress  string
		PortNumber uint32
	}

	type outputType struct {
		TCPServer tcpServer
		TCPClient []tcpClient
	}

	type OutputInfo struct {
		Command    string
		OutputType outputType
	}

	myOutputInfo := OutputInfo{
		Command: "SetParam",
		OutputType: outputType{
			TCPServer: tcpServer{
				IPAddress:  "127.0.0.1",
				PortNumber: 10000,
			},
			TCPClient: []tcpClient{
				{
					IPAddress:  "10.197.237.61",
					PortNumber: 10001,
				},
				{
					IPAddress:  "10.197.237.62",
					PortNumber: 10002,
				},
			},
		},
	}

	dataIntent, err := json.MarshalIndent(myOutputInfo, "", "    ")
	if err != nil {
		t.Logf("Error, json marshalIntent failed: %s", err)
	}
	t.Logf("\n%s", dataIntent)

	dataOrigin, err := json.Marshal(myOutputInfo)
	if err != nil {
		t.Logf("Error, json marshal failed: %s", err)
	}
	t.Logf("\n%s", dataOrigin)

	// 反序列化
	var newOutputInfo OutputInfo
	err = json.Unmarshal(dataOrigin, &newOutputInfo)
	if err != nil {
		t.Logf("Error, json UnMarshal failed: %s", err)
	}
	t.Logf("\n完全结构体反序列化: %v", newOutputInfo)

	var OutputType struct{ OutputType outputType }

	err = json.Unmarshal(dataOrigin, &OutputType)
	if err != nil {
		t.Logf("Error, json Unmarshal failed: %s", err)
	}
	t.Logf("\n部分结构体反序列化: %v", OutputType)
}
