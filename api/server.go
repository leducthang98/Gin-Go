package server

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"go-graphql-boilderplate/api/controller"
	"go-graphql-boilderplate/api/ioc"
	"go-graphql-boilderplate/config"
	"go-graphql-boilderplate/infrastructure/db/rdb"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	appName = "golang-graphql-boilerplate"
)

type Server struct {
	cfg *config.Server
}

func New(cfg *config.Server) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Start(ctx context.Context) (err error) {
	port := strconv.Itoa(int(s.cfg.Port))
	router := chi.NewRouter()

	// Init DB
	rdb.Init(&s.cfg.Mysql)

	// Ioc
	resolverCfg := ioc.ResolverCfg{
		SQLHandler: rdb.GetInstance(),
	}
	iocContext := ioc.NewResolver(resolverCfg)

	// Middleware
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{""},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	// Routers
	controller.MountHealthCheckRouters(router)
	controller.MountUserRouters(router, iocContext.UserService)

	// Start server
	httpSrv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Println("Start server at port:", port)

	idleConnClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
		sig := <-sigint
		log.Printf("Received signal: %s", sig.String())
		if err = httpSrv.Shutdown(ctx); err != nil {
			log.Printf("Shutting down server error: %v", err)
		}
		close(idleConnClosed)
	}()
	if err = httpSrv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server error: %v", err)
	}

	<-idleConnClosed

	return nil
}

func (s *Server) Close() {
	log.Print("Closing server, release resources")
}
