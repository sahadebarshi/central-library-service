package main

import (
	"log"
	"net/http"
	"os"
	"rest/cls/controller"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var mainServerHealthy atomic.Bool

func main() {
	// Set environment variable for the certificates
	server_cert := os.Getenv("SERVER_CERT")
	server_key := os.Getenv("SERVER_KEY")
	server_ca_cert := os.Getenv("SERVER_CA_CERT")

	// Setting helth server instance in go routine and traking main server health
	mainServerHealthy.Store(true)
	go func() {
		healthRouter := gin.Default()
		controller.HealthController(healthRouter, &mainServerHealthy)
		log.Println("Hello World !!")
		log.Fatal(healthRouter.Run(":8083"))
	}()

	// mTLS server instance
	server := gin.Default()
	controller.RouteController(server)

	server.GET("/", func(context *gin.Context) {
		if context.Request.TLS != nil && len(context.Request.TLS.PeerCertificates) > 0 {
			clientCert := context.Request.TLS.PeerCertificates[0]
			context.String(http.StatusOK, "Hello, %s!\n", clientCert.Subject.CommonName)
		} else {
			context.String(http.StatusUnauthorized, "No client certificate provided")
		}
	})

	// TLS config for mTLS
	tlsConfig := createTLSConfig(server_cert, server_key, server_ca_cert)

	// Wrap the Gin router with http.Server to attach TLS config
	httpServer := &http.Server{
		Addr:      "0.0.0.0:8443",
		Handler:   server,
		TLSConfig: tlsConfig,
	}

	log.Println("🚀 Server running at port 8443 with mTLS")
	if err := httpServer.ListenAndServeTLS("", ""); err != nil {
		mainServerHealthy.Store(false)
		log.Printf("Server failed: %v", err)
	}
	// Keep the main goroutine alive (infinite block)
	select {}
}
