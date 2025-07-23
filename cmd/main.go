package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Vladimir5577/go_shop_meat_factures/internal/config"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/handler"
	authMiddleware "github.com/Vladimir5577/go_shop_meat_factures/internal/middleware"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/repository"
	"github.com/Vladimir5577/go_shop_meat_factures/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	envConfigs, err := config.NewEnvConfig()
	if err != nil {
		log.Fatalf("Can't read .env file, %v\n", err)
	}
	db, err := config.ConnectDB(envConfigs)
	if err != nil {
		log.Fatalf("Can't connect to database, %v\n", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	// Home page.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Meet factures."))
	})

	r.Post("/register", userHandler.Register())
	r.Post("/login", userHandler.Login())

	r.Get("/products", productHandler.GetAllProducts())

	r.With(authMiddleware.IsAuthed).Post("/orders", orderHandler.CreateOrder())
	r.Get("/orders", orderHandler.GetOrdersByUser())

	fmt.Println("Server up and running on the port", envConfigs.ServicePort)
	http.ListenAndServe(":"+envConfigs.ServicePort, r)
}
