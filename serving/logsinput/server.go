package logsinput

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"github.com/walkline/ipcounter"
)

type Server struct {
	fasthttp.Server
	ipIndex ipcounter.IPIndex
	logger  log.Logger
}

func NewServer(i ipcounter.IPIndex, logger log.Logger) *Server {
	return &Server{
		ipIndex: i,
		logger:  logger,
	}
}

func (s *Server) ListenAndServe(addr string) error {
	s.Server.Handler = func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) == "/logs" && string(ctx.Method()) == http.MethodPost {
			s.handleLogMsg(ctx)
			return
		}

		_ = s.logger.Log("bad_route", "trying to route to "+string(ctx.Path())+" "+string(ctx.Method()))
		ctx.Error("not found", fasthttp.StatusNotFound)
	}

	return s.Server.ListenAndServe(addr)
}

func (s *Server) handleLogMsg(ctx *fasthttp.RequestCtx) {
	logMsg := &LogMsg{}
	if err := easyjson.Unmarshal(ctx.PostBody(), logMsg); err != nil {
		_ = s.logger.Log("err", err.Error())
		ctx.Error("bad log msg", fasthttp.StatusBadRequest)
		return
	}

	if err := s.ipIndex.Add(logMsg.IP); err != nil {
		_ = s.logger.Log("err", err.Error())
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusAccepted)
}
