package waveapi

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"http"
	"io/ioutil"
	"json"
)

type Handler func(*Event, *Wavelet)

type Robot struct {
	name       string
	imageUrl   string
	profileUrl string
	pathPrefix string
	handlers   map[string]Handler
}

func NewRobot(name, imageUrl, profileUrl, pathPrefix string) (r *Robot) {
	r = &Robot{
		name:       name,
		imageUrl:   imageUrl,
		profileUrl: profileUrl,
		pathPrefix: pathPrefix,
		handlers:   make(map[string]Handler),
	}
	return
}

func (r *Robot) RegisterHandler(eventType string, handler Handler) {
	r.handlers[eventType] = handler
}

type requestBundle struct {
	Events       []*Event
	Wavelet      *WaveletData
	Blips        map[string]*BlipData
	RobotAddress string
}

func (r *Robot) ServeHTTP(conn *http.Conn, req *http.Request) {
	log("Path:", req.URL.Path)
	switch req.URL.Path {
	case r.pathPrefix + "/robot/jsonrpc":
		// decode request bundle
		body, _ := ioutil.ReadAll(req.Body)
		logf("Request:\n%s\n\n", body)
		b := new(requestBundle)
		err := json.Unmarshal(body, b)
		if err != nil {
			http.Error(conn, "Bad request", http.StatusBadRequest)
			return
		}

		oq := NewOperationQueue()
		w := NewWavelet(b.Wavelet, b.Blips, r, oq)

		// handle events
		for _, e := range b.Events {
			if handler, ok := r.handlers[e.Type]; ok {
				handler(e, w)
			}
		}

		// output operation queue
		q, _ := json.Marshal(oq.pending)
		logf("Response:\n%s\n\n", q)
		conn.Write(q)
	case r.pathPrefix + "/capabilities.xml":
		fmt.Fprintf(conn, capabilitiesXMLstart, r.version())
		for k, _ := range r.handlers {
			fmt.Fprintf(conn, capabilityXML, k)
		}
		fmt.Fprint(conn, capabilitiesXMLend)
	case r.pathPrefix + "/robot/profile":
		p, _ := json.Marshal(map[string]string{
			"profileUrl": r.profileUrl,
			"imageUrl":   r.imageUrl,
			"name":       r.name,
		})
		conn.Write(p)
	default:
		http.NotFound(conn, req)
	}
}

func (r *Robot) version() string {
	v := sha1.New()
	for k, _ := range r.handlers {
		fmt.Fprint(v, k)
	}
	return hex.EncodeToString(v.Sum())
}

const capabilitiesXMLstart = `<?xml version="1.0"?>
<w:robot xmlns:w="http://wave.google.com/extensions/robots/1.0">
<w:version>%s</w:version>
<w:protocolversion>0.21</w:protocolversion>
<w:capabilities>
`
const capabilitiesXMLend = `
</w:capabilities>
</w:robot>
`
const capabilityXML = `<w:capability name="%s"/>`
