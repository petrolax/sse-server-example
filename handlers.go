package sse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	service *Service
	retry   int
}

func NewHandlers(service *Service, retry int) *Handlers {
	return &Handlers{
		service: service,
		retry:   retry,
	}
}

func (h *Handlers) ListeHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Allow", "GET")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Content-Type", "text/event-stream; charset=utf-8")
	if req.Method == http.MethodHead {
		return
	}
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	flusher := w.(http.Flusher)

	buffer := bufio.NewWriter(w)
	buffer.WriteString(OpenEvent().String())
	buffer.Flush()
	flusher.Flush()

	ticker := time.NewTicker(time.Second * 5)
	i := 1
	event := &Event{
		Event: "message",
		Data:  []byte(h.service.GetValue()),
		ID:    fmt.Sprint(i),
		Retry: h.retry,
	}
	for {
		select {
		case <-ticker.C:
		case <-req.Context().Done():
			log.Println("Context dead", req.Context())
			buffer.WriteString(CloseEvent().String())
			buffer.Flush()
			return
		}

		event.ID = fmt.Sprint(i)
		event.Data = []byte(h.service.GetValue())
		buffer.WriteString(event.String())
		buffer.Flush()
		flusher.Flush()
		log.Println("Sent data")
	}
}

type ReplaceWord struct {
	Word string `json:"word"`
}

func (h *Handlers) SayHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Allow", "POST")
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	word := new(ReplaceWord)
	if err := json.NewDecoder(req.Body).Decode(word); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: can't decode body")
		return
	}

	h.service.SetValue(word.Word)

	w.WriteHeader(http.StatusOK)
}
