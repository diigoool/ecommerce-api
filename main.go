package main

import (
	"context"
	"ecommerce-api/config"
	_ "ecommerce-api/docs"
	"ecommerce-api/handlers"
	"ecommerce-api/middlewares"
	"ecommerce-api/models"
	"ecommerce-api/repositories"
	"ecommerce-api/seeders"
	"ecommerce-api/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Ecommerce API
// @version 1.0
// @description Ecommerce API with Golang Chi
// @host localhost:3000
// @BasePath /
func main() {
	config.LoadEnv()

	config.ConnectDB()

	config.InitJWT()
	config.DB.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)

	seeders.SeedAdmin(config.DB)

	// product repo
	productRepo := repositories.NewProductRepository(config.DB)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// auth repo
	authRepo := repositories.NewUserRepository(config.DB)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// cart repo
	cartRepo := repositories.NewCartRepository(config.DB)
	cartService := services.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService)

	// order repo
	orderRepo := repositories.NewOrderRepository(config.DB)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	r := chi.NewRouter()
	port := os.Getenv("PORT")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	r.Use(middlewares.RequestID)
	r.Use(middlewares.Logger)

	r.Post("/api/auth/login", authHandler.Login)
	r.Post("/api/auth/register", authHandler.Register)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.JWTAuth)

		// routes product
		r.Get("/api/product", productHandler.GetProducts)
		r.Get("/api/product/{id}", productHandler.GetProductById)
		r.Post("/api/product/create", middlewares.AdminOnly(productHandler.CreateProduct))
		r.Delete("/api/product/delete/{id}", middlewares.AdminOnly(productHandler.DeleteProduct))
		r.Patch("/api/product/update/{id}", middlewares.AdminOnly(productHandler.UpdateProduct))

		// routes cart
		r.Post("/api/cart", cartHandler.AddToCart)
		r.Get("/api/cart", cartHandler.GetCart)
		r.Delete("/api/cart", cartHandler.RemoveItem)

		// routes order
		r.Post("/api/checkout", orderHandler.Checkout)
		r.Get("/api/orders", orderHandler.GetOrders)

	})

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	go func() {

		log.Printf("server running on %s", port)

		if err := srv.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("server error : %v", err)
		}

	}()
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf(
			"server forced to shutdown: %v", err,
		)
	}
	sqlDB, err := config.DB.DB()

	if err == nil {
		sqlDB.Close()
	}

	log.Println("server exited properly")

}
