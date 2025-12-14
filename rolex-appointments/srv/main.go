package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"my-rolex/rolexserver"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
    var tokenFlag bool
    var helpFlag bool
    var debugFlag bool
    var configPath string
    flag.BoolVar(&tokenFlag, "t", false, "create token")
    flag.BoolVar(&helpFlag, "h", false, "print usages")
    flag.StringVar(&configPath, "c", "", "config file path")
    flag.BoolVar(&debugFlag, "debug", false, "debug info")
    flag.Parse()

    if err := rolexserver.InitServer(configPath); err != nil {
        slog.Error("Failed to load config, exiting...", "error", err)
        os.Exit(1)
    }
    if debugFlag {
        logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
            Level: slog.LevelDebug, // 设置日志级别为 Debug
        }))
    
        // 将 logger 设置为全局默认 logger
        slog.SetDefault(logger)
    }
    if tokenFlag {
        token()
    } else if helpFlag {
        flag.Usage()
    } else {
        srv()
    }
}

func token() {
    token, err := rolexserver.GenerateToken()
    if err == nil {
        fmt.Println(token)
    }
}

func srv() {
    var wait time.Duration
    flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()

    r := mux.NewRouter()
    // Add your routes as needed
    r.HandleFunc("/agent",rolexserver.RegisterAgentHandler).Methods(http.MethodPost)
    r.Use(
        func (next http.Handler) http.Handler  {
            return handlers.LoggingHandler(os.Stdout, next)
        },
        rolexserver.AuthMiddleware)

    srv := &http.Server{
        Addr:         "0.0.0.0:8080",
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: r, // Pass our instance of gorilla/mux in.
    }

    // Run our server in a goroutine so that it doesn't block.
    go func() {
        if err := srv.ListenAndServe(); err != nil {
            slog.Error("Server error","error", err)
        }
    }()

    c := make(chan os.Signal, 1)
    // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
    signal.Notify(c, os.Interrupt)

    // Block until we receive our signal.
    <-c

    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    srv.Shutdown(ctx)
    // Optionally, you could run srv.Shutdown in a goroutine and block on
    // <-ctx.Done() if your application should wait for other services
    // to finalize based on context cancellation.
    slog.Info("shutting down")
    os.Exit(0)
}
