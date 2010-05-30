package waveapi

import (
	"bytes"
	"http"
	"io/ioutil"
	"testing"
	"time"
)

func handlerFunc(e *Event, w *Wavelet) {
	w.Reply("Test response")
}

func setupRobot() {
	r := NewRobot("Test robot", "http://avatar-url", "http://profile-url", "/_wave")
	r.RegisterHandler(E_BlipSubmitted, handlerFunc)
	http.ListenAndServe(":8080", r)
}

func TestCapabilities(t *testing.T) {
	go setupRobot()
	time.Sleep(2e8)

	body := bytes.NewBufferString(test_req)
	resp, _ := http.Post(
		"http://127.0.0.1:8080/_wave/robot/jsonrpc",
		"text/json",
		body)
	respBody, _ := ioutil.ReadAll(resp.Body)
	if string(respBody) != test_resp {
		t.Error("Response body not as expected.")
	}
}

const test_req = `{
  "events":[{
    "modifiedBy": "user@example.com",
    "timestamp": 1255935016481,
    "type": "BLIP_SUBMITTED",
    "properties": {
      "blipId": "b+ja8F_Hw4J"
    }
  }],
  "wavelet": {
    "creationTime": 1255934856713,
    "creator": "user@example.com",
    "lastModifiedTime": 1255935016481,
    "participants": [ "user@example.com","user2@example.com" ],
    "rootBlipId": "b+ja8F_Hw4J",
    "title": "",
    "version": 11,
    "waveId": "example.com!w+ja8F_Hw4I",
    "waveletId": "example.com!conv+root",
    "dataDocuments": {
    }
  },
  "blips": {
    "b+ja8F_Hw4J": {
      "annotations": [{
        "range": {
          "start": 0,
          "end": 1
        },
        "name": "conv/title",
        "value": ""
      }],
      "elements": {},
      "blipId": "b+ja8F_Hw4J",
      "childBlipIds": [],
      "contributors": ["user@example.com"],
      "creator": "user@example.com",
      "content": "\n",
      "lastModifiedTime": 1255934856708,
      "version": 6,
      "waveId": "google.com!w+ja8F_Hw4I",
      "waveletId": "example.com!conv+root"
    }
  },
  "robotAddress": "myrobot@example.com"
}
`
const test_resp = `[{"method":"wavelet.appendBlip","id":"op%1","params":{"blipData":{"blipId":"TBD_example.com!conv+root_7fcfd52","content":"Test response","parentBlipId":"","waveId":"example.com!w+ja8F_Hw4I","waveletId":"example.com!conv+root"},"waveId":"example.com!w+ja8F_Hw4I","waveletId":"example.com!conv+root"}}]`
