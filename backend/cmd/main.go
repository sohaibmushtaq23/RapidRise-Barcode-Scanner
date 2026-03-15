package main

import (
	"bytes"
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"rapid-rise-backend/config"
	"rapid-rise-backend/middlewares"
	"rapid-rise-backend/routes"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed all:dist
var frontendFS embed.FS

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	if err := config.InitDB(); err != nil {
		log.Fatal("DB connection failed:", err)
	}
	log.Println("Connected to SQL Server!")

	 // Create the main Chi router
	 r := chi.NewRouter()

	 // Apply global middlewares (order matters: logger, recoverer, cors)
	 r.Use(middlewares.Logger)
	 r.Use(middlewares.Recoverer)
	 r.Use(middlewares.Cors)
 
	 // Mount API subrouter under /api
	 r.Mount("/api", routes.RegisterRoutes())
 
	 // Optional: health check endpoint
	 r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		 w.Write([]byte("server working"))
	 })

	// Create filesystem from embedded files
	distFS, err := fs.Sub(frontendFS, "dist")
	if err != nil {
		log.Fatal(err)
	}
	printEmbeddedFS(frontendFS)
	
	// Serve frontend – static files or SPA fallback
	r.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Trim leading slash for fs.FS
		path := strings.TrimPrefix(req.URL.Path, "/")
		f, err := distFS.Open(path)
		if err == nil {
			defer f.Close()

			stat, err := f.Stat()
			if err != nil {
				http.Error(w, "failed to stat file", http.StatusInternalServerError)
				return
			}

			content, err := io.ReadAll(f)
			if err != nil {
				http.Error(w, "failed to read file", http.StatusInternalServerError)
				return
			}

			// Optional: log that we're serving a file
			log.Printf("Serving file: %s", path)

			http.ServeContent(w, req, path, stat.ModTime(), bytes.NewReader(content))
			return
		}

		// If file doesn't exist, serve index.html (SPA fallback)
		indexFile, err := distFS.Open("index.html")
		if err != nil {
			http.Error(w, "index.html not found", http.StatusInternalServerError)
			return
		}
		defer indexFile.Close()

		stat, err := indexFile.Stat()
		if err != nil {
			http.Error(w, "failed to stat index.html", http.StatusInternalServerError)
			return
		}

		content, err := io.ReadAll(indexFile)
		if err != nil {
			http.Error(w, "failed to read index.html", http.StatusInternalServerError)
			return
		}

		http.ServeContent(w, req, "index.html", stat.ModTime(), bytes.NewReader(content))
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	certFile := os.Getenv("TLS_CERT")
	keyFile := os.Getenv("TLS_KEY")

	// If cert files are provided, use HTTPS; otherwise HTTP (for dev)
	if certFile != "" && keyFile != "" {
		log.Printf("Starting HTTPS server on :%s", port)
		err := http.ListenAndServeTLS(":"+port, certFile, keyFile, r)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("Starting HTTP server on :%s", port)
		err := http.ListenAndServe(":"+port, r)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func printEmbeddedFS(fsys fs.FS) {
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		log.Printf("Embedded: %s", path)
		return nil
	})
	if err != nil {
		log.Printf("Error walking embedded FS: %v", err)
	}
}