package server

import (
	"net/http"

	"github.com/jasonkwh/droneshield-test/internal/config"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

type Handler struct {
	rcfg config.RedisConfig
	zl   *zap.Logger
}

// NewServer - creates new server for sending out websocket messages
func NewServer(scfg config.ServerConfig, rcfg config.RedisConfig, zl *zap.Logger) *http.Server {
	hdl := &Handler{
		rcfg: rcfg,
		zl:   zl,
	}

	// handle websocket
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(hdl.handleSock))
	srv := http.Server{
		Addr:    ":" + scfg.Port,
		Handler: mux,
	}

	return &srv
}

func (h *Handler) handleSock(w http.ResponseWriter, r *http.Request) {
	h.zl.Info("received connection")

	wsc, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		CompressionMode:    websocket.CompressionDisabled,
		InsecureSkipVerify: true, // enable for testing purposes
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sp, err := NewSocketPublisher(wsc, h.rcfg, h.zl)
	if err != nil {
		wsc.Close(websocket.StatusInternalError, err.Error())
		return
	}

	err = sp.PublishLoop()
	if err != nil {
		wsc.Close(websocket.StatusInternalError, err.Error())
		return
	}

	wsc.Close(websocket.StatusNormalClosure, "")
}
