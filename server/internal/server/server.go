package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Handler interface {
	Path() string
	Handler() func(http.ResponseWriter, *http.Request)
	Methods() string
}

type Server struct {
	handlers []Handler
	port     uint16
}

func NewServer(port uint16) *Server {
	return &Server{port: port}
}

func (s *Server) AddHandler(h Handler) {
	s.handlers = append(s.handlers, h)
}

func (s *Server) Start() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(
		cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"chrome-extension://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST"},
			AllowedHeaders:   []string{"Accept", "Content-Type"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
	)

	for _, h := range s.handlers {
		router.MethodFunc(h.Methods(), h.Path(), h.Handler())
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		log.Printf("Server started on port %d\n", s.port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
