package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/07-apis/configs"
	_ "github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/07-apis/docs"
	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/07-apis/internal/entity"
	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/07-apis/internal/infra/database"
	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/07-apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Vitor Camargo
// @contact.url    https://www.vitorlrrcamargo.com
// @contact.email  contato@vitorlrrcamargo.com

// @license.name  Vitor Camargo License
// @license.url   https://www.vitorlrrcamargo.com/license

// @host                        localhost:8000
// @BasePath                    /
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn))
	// r.Use(LogRequest)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
