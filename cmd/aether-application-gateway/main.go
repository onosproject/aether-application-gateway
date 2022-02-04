// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/onosproject/aether-application-gateway/internal/config"
	"github.com/onosproject/aether-application-gateway/internal/router"
	promApi "github.com/prometheus/client_golang/api"
	promApiV1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.New()

	flag.IntVar(&cfg.Server.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Server.Env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	t := http.DefaultTransport.(*http.Transport).Clone()
	httpClient := &http.Client{
		Timeout:   time.Second * 20,
		Transport: t,
	}

	pc, err := promApi.NewClient(promApi.Config{
		Address: cfg.Prometheus.Acc,
	})
	if err != nil {
		fmt.Printf("Error creating prometheus client: %v\n", err)
		os.Exit(1)
	}
	promV1 := promApiV1.NewAPI(pc)

	r := router.Setup(httpClient, &cfg.Roc, promV1)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Aether Application Gateway API")
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
