package waveapi

import (
	"fmt"
	"rand"
)

type Blip struct {
	BlipData
	robot   *Robot
	opQueue *OperationQueue
}

// Blip data as represented in the JSON-encoded requests
type BlipData struct {
	BlipId           string
	ChildBlipIds     []string
	Contributors     []string
	Creator          string
	LastModifiedTime int64
	Content          string
	Version          int
	WaveId           string
	WaveletId        string
}

func NewBlip(bd *BlipData, r *Robot, oq *OperationQueue) *Blip {
	return &Blip{*bd, r, oq}
}

func newBlipHash(waveId, waveletId, content, parentBlipId string) map[string]string {
	tempBlipId := fmt.Sprintf("TBD_%s_%x", waveletId, rand.Int())
	return map[string]string{
		"waveId":       waveId,
		"waveletId":    waveletId,
		"blipId":       tempBlipId,
		"content":      content,
		"parentBlipId": parentBlipId,
	}
}

func newBlipFromHash(bd map[string]string, r *Robot, oq *OperationQueue) (b *Blip) {
	b = NewBlip(&BlipData{
		WaveId: bd["waveId"],
		WaveletId: bd["waveletId"],
		BlipId: bd["blipId"],
		Content: bd["content"],
	}, r, oq)
	return
}

// TODO
// Find ?
// Append
// Reply
// AppendMarkup
// InsertInlineBlip

