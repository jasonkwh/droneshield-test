package server

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"

	"github.com/gin-gonic/gin"
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
		zl: zl,
	}

	// default gin mode is DEBUG
	//gin.SetMode(gin.ReleaseMode)

	mux := gin.Default()

	// Add a ginzap middleware
	mux.Use(ginzap.Ginzap(zl, time.RFC3339, false))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	mux.Use(ginzap.RecoveryWithZap(zl, false))

	// handle websocket
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
}
