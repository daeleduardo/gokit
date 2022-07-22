package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"gokit/config"
	log "gokit/lib"
	"gokit/routes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	c := config.Config{}
	Default := config.DefaultConfig()

	flag.StringVar(&c.DbConn, "database", Default.DbConn, "Database DSN URL connection example: user:password@tcp(localhost:3303)/database")
	flag.StringVar(&c.CacheConn, "cache", Default.CacheConn, "Redis DSN URL connection example: redis://username:password@host:port/0")
	flag.StringVar(&c.GinMode, "gin-mode", Default.GinMode, "Gin Gonic Mode: dev|release")
	flag.StringVar(&c.JwtToken, "jwt-token", Default.JwtToken, "JWT token")
	flag.Int64Var(&c.JwtTimeout, "jwt-timeout", Default.JwtTimeout, "JWT token timeout in Time.Duration values default 1 Hour")
	flag.IntVar(&c.Port, "port", Default.Port, "HTTP port default 80")
	flag.Parse()
	c.Init()
	
	if c.GinMode == Default.GinMode {
		log.BuildLogger("debug")
		gin.SetMode(gin.DebugMode)
	} else {
		log.BuildLogger("info")
		gin.SetMode(gin.ReleaseMode)
	}
	
	TlsConfig := &tls.Config{
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},
	}

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", c.Port),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       10 * time.Second,
		TLSConfig:         TlsConfig,
		Handler: 		   routes.GetRouters(),
		//  MaxHeaderBytes: 1 << 20,
	}
	//Disable KeepAlive
	s.SetKeepAlivesEnabled(false)

	err := s.ListenAndServe()

	if err != nil {
		log.Log().Error(err.Error())
	}

}