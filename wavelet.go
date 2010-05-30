package waveapi

import (
	"json"
)

type Wavelet struct {
	WaveletData
	blips   map[string]*Blip
	robot   *Robot
	opQueue *OperationQueue
}

// Wavelet data as represented in teh JSON-encoded requests
type WaveletData struct {
	CreationTime     int64
	Creator          string
	LastModifiedTime int64
	Participants     []string
	RootBlipId       string
	Title            string
	Version          int
	WaveId           string
	WaveletId        string
	DataDocuments    map[string]string
}

func NewWavelet(wd *WaveletData, bd map[string]*BlipData, r *Robot, oq *OperationQueue) *Wavelet {
	// create Blips from BlipData
	b := make(map[string]*Blip)
	for k, v := range bd {
		b[k] = NewBlip(v, r, oq)
	}
	// create wavelet
	return &Wavelet{*wd, b, r, oq}
}

func (w *Wavelet) Reply(content string) (b *Blip) {
	bd := w.opQueue.WaveletAppendBlip(w.WaveId, w.WaveletId, content)
	b = newBlipFromHash(bd, w.robot, w.opQueue)
	// TODO: create Blip instance and return it
	return
}

func (w *Wavelet) Delete(blipId string) {
	w.opQueue.BlipDelete(w.WaveId, w.WaveletId, blipId)
}

func (w *Wavelet) SetDataDoc(key, value string) {
	w.opQueue.WaveletDatadocSet(w.WaveId, w.WaveletId, key, value)
	w.DataDocuments[key] = value
}


// TODO(adg): Remove the below. It is a hack for the I/O demo.
const embedYoutubeJSON = `
{
	"modifyAction": {
		"modifyHow": "INSERT_AFTER",
		"elements": [
			{
				"type": "GADGET",
				"properties": {
					"url": "http://wh3rd.net/dump/youtube.xml"
				}
			}
		]
	}
}`

func (w *Wavelet) EmbedVideo(blipId string) {
	v := new(Params)
	err := json.Unmarshal([]byte(embedYoutubeJSON), v)
	if err != nil {
		println(err.String())
		return
	}
	(*v)["blipId"] = blipId
	w.opQueue.NewOperation(DOCUMENT_MODIFY, w.WaveId, w.WaveletId, *v)
}
