package main

import (
	"http"
	"waveapi"
)

func handlerFunc(e *waveapi.Event, w *waveapi.Wavelet) {
	w.Reply("Some response")
}

func main() {
	r := waveapi.NewRobot(
		"Example robot",
		"http://exmaple.com/avatar.png",
		"http://example.com/profile.html",
		"")
	r.RegisterHandler(waveapi.E_WaveletSelfAdded, handlerFunc)
	http.ListenAndServe(":8080", r)
}
