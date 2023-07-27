package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

type Handler struct {
	zl *zap.Logger
}

func NewServer(zl *zap.Logger) *http.Server {
	hdl := &Handler{
		zl: zl,
	}

	// start web server
	// mux := http.NewServeMux()
	// mux.Handle("/", http.HandlerFunc(hdl.handleSock))
	// srv := http.Server{
	// 	Addr:    ":" + scfg.Port,
	// 	Handler: mux,
	// }
	mux := gin.Default()
	mux.GET("/", hdl.handleSock())
	srv := http.Server{
		Addr:    ":" + scfg.Port,
		Handler: mux,
	}

	return &srv
}

func (h *Handler) handleSock() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.zl.Info("received connection")

		wsc, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
			CompressionMode: websocket.CompressionDisabled,
		})
		if err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}
