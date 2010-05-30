package waveapi

import (
	"fmt"
)

const (
	WAVELET_APPEND_BLIP             = "wavelet.appendBlip"
	WAVELET_SET_TITLE               = "wavelet.setTitle"
	WAVELET_ADD_PARTICIPANT         = "wavelet.participant.add"
	WAVELET_DATADOC_SET             = "wavelet.datadoc.set"
	WAVELET_MODIFY_TAG              = "wavelet.modifyTag"
	WAVELET_MODIFY_PARTICIPANT_ROLE = "wavelet.modifyParticipantRole"
	BLIP_CREATE_CHILD               = "blip.createChild"
	BLIP_DELETE                     = "blip.delete"
	DOCUMENT_APPEND_MARKUP          = "document.appendMarkup"
	DOCUMENT_INLINE_BLIP_INSERT     = "document.inlineBlip.insert"
	DOCUMENT_MODIFY                 = "document.modify"
	ROBOT_CREATE_WAVELET            = "robot.createWavelet"
	ROBOT_FETCH_WAVE                = "robot.fetchWave"
	ROBOT_NOTIFY_CAPABILITIES_HASH  = "robot.notifyCapabilitiesHash"
)

type Operation struct {
	method string
	id     string
	params Params
}

type Params map[string]interface{}

func NewOperation(method, id string, params Params) *Operation {
	if params == nil {
		params = make(Params)
	}
	return &Operation{
		method: method,
		id:     id,
		params: params,
	}
}

type OperationQueue struct {
	nextOpId int
	pending  []*Operation
}

const operationQueueDefaultLen = 10

func NewOperationQueue() *OperationQueue {
	return &OperationQueue{
		nextOpId: 1,
		pending:  make([]*Operation, 0, 10),
	}
}

func (oq *OperationQueue) NewOperation(method, waveId, waveletId string, params Params) (o *Operation) {
	if params == nil {
		params = make(Params)
	}
	params["waveId"] = waveId
	params["waveletId"] = waveletId

	o = NewOperation(method, fmt.Sprintf("op%d", oq.nextOpId), params)
	oq.nextOpId++

	// grow queue if necessary
	l := len(oq.pending)
	if l == cap(oq.pending) {
		q := make([]*Operation, l, l+operationQueueDefaultLen)
		copy(q, oq.pending)
		oq.pending = q
	}
	// update queue
	oq.pending = oq.pending[0 : l+1] // grow slice
	oq.pending[l] = o

	return
}

func (oq *OperationQueue) WaveletAppendBlip(waveId, waveletId, content string) (bd map[string]string) {
	bd = newBlipHash(waveId, waveletId, content, "")
	oq.NewOperation(WAVELET_APPEND_BLIP, waveId, waveletId,
		Params{"blipData": bd})
	return
}

func (oq *OperationQueue) WaveletDatadocSet(waveId, waveletId, key, value string) {
	oq.NewOperation(WAVELET_DATADOC_SET, waveId, waveletId,
		Params{"datadocName": key, "datadocValue": value})
}

func (oq *OperationQueue) BlipDelete(waveId, waveletId, blipId string) {
	oq.NewOperation(BLIP_DELETE, waveId, waveletId,
		Params{"blipId": blipId})
}

