package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codersgyan/students-api/internal/config"
	student "github.com/codersgyan/students-api/internal/http/handlers"
	config "github.com/codersgyan/students-api/internal"
)

// func main() {
// 	// fmt.Println("Welcome to go API")

//
// func main() {
// 	// fmt.Println("Welcome to go API")

// 	//load config
// 	cfg := config.MustLoad()
// 	//database setup
// 	//setup router

// 	//setup server
// 	router := http.NewServeMux()
// 	router.HandleFunc("POST /api/students", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Welcome to students api"))
// 	})
// 	// setup server
// 	server := http.Server{
// 		Addr:    cfg.Addr,
// 		Handler: router,
// 	}
// 	fmt.Println("Server Started")
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatal("failed to start server")
// 	}
// }

//
// func main() {
// 	// fmt.Println("Welcome to go API")

// 	//load config
// 	cfg := config.MustLoad()
// 	//database setup
// 	//setup router

// 	//setup server
// 	router := http.NewServeMux()
// 	router.HandleFunc("POST /api/students", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Welcome to students api"))
// 	})
// 	// setup server
// 	server := http.Server{
// 		Addr:    cfg.Addr,
// 		Handler: router,
// 	}
// 	fmt.Println("Server Started %s", cfg.HTTPServer.Addr)
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatal("failed to start server")
// 	}
// }

//
// func main() {
// 	//load config
// 	cfg := config.MustLoad()
// 	//database setup
// 	//setup server
// 	router := http.NewServeMux()
// 	router.HandleFunc("POST /api/students", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Welcome to students api"))
// 	})
// 	// setup server
// 	server := http.Server{
// 		Addr:    cfg.Addr,
// 		Handler: router,
// 	}
// 	slog.Info("Server Started", slog.String("address", cfg.Addr))
// 	done := make(chan os.Signal, 1)
// 	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
// 	go func() {
// 		err := server.ListenAndServe()
// 		if err != nil {
// 			log.Fatal("failed to start server")
// 		}
// 	}()
// 	<-done
// 	slog.Info("Shutting down the server")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
// 	defer cancel()
// 	if err := server.Shutdown(ctx); err != nil {
// 		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
// 	}
// 	slog.Info("shutting down the server")
// }

// Now importing student routes
func main() {
	//load config
	cfg := config.MustLoad()
	//database setup
		storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	//setup server
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("Server Started", slog.String("address", cfg.Addr))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done
	slog.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("shutting down the server")
}

//

// 	//load config
// 	cfg := config.MustLoad()
// 	//database setup
// 	//setup router

// 	//setup server
// 	router := http.NewServeMux()
// 	router.HandleFunc("POST /api/students", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Welcome to students api"))
// 	})
// 	// setup server
// 	server := http.Server{
// 		Addr:    cfg.Addr,
// 		Handler: router,
// 	}
// 	fmt.Println("Server Started")
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatal("failed to start server")
// 	}
// }

// router.HandleFunc("POST /api/students", student.New(storage))

//
// import (
// 	"context"
// 	"log"
// 	"log/slog"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/codersgyan/students-api/internal/config"
// 	"github.com/codersgyan/students-api/internal/http/handlers/student"
// 	"github.com/codersgyan/students-api/internal/storage/sqlite"
// )

// func main() {
// 	// load confi
// 	// database setup

// 	storage, err := sqlite.New(cfg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

// 	// setup router
// 	router := http.NewServeMux()

//

// 	slog.Info("server started", slog.String("address", cfg.Addr))

// 	done := make(chan os.Signal, 1)

// 	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

// 	<-done

//

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	slog.Info("server shutdown successfully")

// }
