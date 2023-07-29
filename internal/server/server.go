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

	redisPsCh := make(chan []byte)
	done := make(chan struct{})
	sp, err := NewSocketPublisher(wsc, h.rcfg, redisPsCh, done, h.zl)
	if err != nil {
		wsc.Close(websocket.StatusInternalError, err.Error())
		return
	}

	sub, err := NewSubscriber(h.rcfg, redisPsCh, h.zl)
	if err != nil {
		wsc.Close(websocket.StatusInternalError, err.Error())
		return
	}
	defer sub.Close()

	// listen on redis pubsub
	go func() {
		err := sub.Listen()
		if err != nil {
			h.zl.Error("redis pubsub subscriber error", zap.Error(err))

			close(redisPsCh)
			close(done)
		}
	}()

	err = sp.PublishLoop()
	if err != nil {
		wsc.Close(websocket.StatusInternalError, err.Error())
		return
	}

	wsc.Close(websocket.StatusNormalClosure, "")
}
