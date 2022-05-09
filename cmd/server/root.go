package server

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/Amazeful/Amazeful-Backend/api/auth"
// 	v1 "github.com/Amazeful/Amazeful-Backend/api/v1"
// 	"github.com/Amazeful/Amazeful-Backend/config"
// 	"github.com/Amazeful/Amazeful-Backend/util"
// 	"github.com/go-chi/chi/v5"
// 	chimid "github.com/go-chi/chi/v5/middleware"
// 	"github.com/go-chi/httprate"
// 	"github.com/joho/godotenv"
// 	"github.com/spf13/cobra"
// )

// var (
// 	reqTimeout   = 2 * time.Minute
// 	requestLimit = 10
// 	limitTimeout = 10 * time.Second
// 	ctx, cancel  = context.WithTimeout(context.Background(), 5*time.Minute)
// )

// // rootCmd represents the base command when called without any subcommands
// var serverCmd = &cobra.Command{
// 	Use:   "serve",
// 	Short: "Start the webserver",
// 	Long:  "Start the webserver",
// 	// Uncomment the following line if your bare application
// 	// has an action associated with it:
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		initConfig()
// 		initServices()
// 		defer cleanupServices()
// 		initServer()
// 	},
// }

// func initConfig() {
// 	godotenv.Load()
// 	err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatalf("failed to init config -- %v", err)
// 	}
// }

// func initServices() {
// 	cfg := config.GetConfig()
// 	err := util.InitAllServices(ctx, cfg)
// 	if err != nil {
// 		log.Fatalf("failed to services -- %v", err)
// 	}
// }

// func initServer() {
// 	cfg := config.GetConfig()
// 	r := chi.NewRouter()
// 	r.Use(chimid.RequestID)
// 	r.Use(chimid.RealIP)
// 	r.Use(chimid.Logger)
// 	r.Use(chimid.Recoverer)
// 	r.Use(chimid.Timeout(reqTimeout))
// 	r.Use(httprate.Limit(requestLimit, limitTimeout, httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint)))

// 	//routers
// 	r.Route("/auth", auth.ProcessRoutes)
// 	r.Route("/v1", v1.ProcessRoutes)

// 	var err error
// 	if cfg.ServerConfig.TLS {
// 		err = http.ListenAndServeTLS(cfg.ServerConfig.IpAddress+":"+cfg.ServerConfig.Port, cfg.ServerConfig.CertPath, cfg.ServerConfig.KeyPath, r)
// 	} else {
// 		err = http.ListenAndServe(cfg.ServerConfig.IpAddress+":"+cfg.ServerConfig.Port, r)
// 	}
// 	if err != nil {
// 		log.Fatalf("failed to init server -- %v", err)
// 	}
// }

// func cleanupServices() {
// 	cancel()
// 	util.DB().Disconnect(context.Background())
// }
