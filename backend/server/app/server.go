package app

import (
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/backend/server/app/config"
	"github.com/temphia/temphia/backend/server/app/routes"
	"github.com/temphia/temphia/backend/server/btypes"
	"gitlab.com/mr_balloon/golib"
)

var _ btypes.Server = (*Server)(nil)

type Server struct {
	app       *App
	routes    *routes.R
	ginEngine *gin.Engine
	config    *config.AppConfig
}

func (s *Server) BindRoutes() {
	s.routes = routes.New(s.app, s.config)
	s.buildRoutes(s.ginEngine)
}

func (s *Server) Listen() error {
	return s.listenHTTP()
}

func (s *Server) Close() error {
	return nil
}

func (s *Server) listenHTTP() error {
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	return s.ginEngine.RunListener(listener)
}

func (s *Server) listenHTTPS() error {
	return nil
}

func (a *Server) listenUnix() error {
	gen := gin.Default()

	gen.NoRoute(func(c *gin.Context) {
		c.Writer.Write([]byte(` fixme => give operator creds`))
	})

	sockpath := "tmp/temphia.sock" // fixme => make this configurable

	exist, err := golib.FileExists(sockpath)
	if err != nil {
		return err
	}
	if exist {
		os.Remove(sockpath)
	}

	return gen.RunUnix(sockpath)
}

/*

socat - UNIX-CONNECT:./tmp/temphia.sock

POST / HTTP/1.1
Host: localhost
User-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:93.0) Gecko/20100101 Firefox/93.0
action: ping

42
*/
