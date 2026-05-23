package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iuriGaldino/Axion/backend/internal/application/usecase"
	"github.com/iuriGaldino/Axion/backend/internal/infrastructure/persistence/postgres"
	"github.com/iuriGaldino/Axion/backend/internal/infrastructure/security"
	"github.com/iuriGaldino/Axion/backend/internal/interface/api/handler"
	"github.com/iuriGaldino/Axion/backend/internal/interface/api/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Database connection
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Printf("Warning: Could not connect to database: %v", err)
	}

	app := fiber.New(fiber.Config{
		AppName: "Axion API",
	})

	middleware.InitLogger()
	app.Use(middleware.StructuredLogger())
	app.Use(recover.New())

	prometheus := fiberprometheus.New("axion-api")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Services
	hashService := security.NewArgon2idHasher()
	jwtService := security.NewJWTService(os.Getenv("JWT_SECRET"))

	// Repositories
	userRepo := postgres.NewPostgresUserRepository(db)
	categoryRepo := postgres.NewPostgresCategoryRepository(db)
	transactionRepo := postgres.NewPostgresTransactionRepository(db)
	budgetRepo := postgres.NewPostgresBudgetRepository(db)
	goalRepo := postgres.NewPostgresGoalRepository(db)

	// Use Cases
	authUseCase := usecase.NewAuthUseCase(userRepo, hashService, jwtService)
	userUseCase := usecase.NewUserUseCase(userRepo)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	transactionUseCase := usecase.NewTransactionUseCase(transactionRepo)
	budgetGoalUseCase := usecase.NewBudgetGoalUseCase(budgetRepo, goalRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authUseCase)
	userHandler := handler.NewUserHandler(userUseCase)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)
	transactionHandler := handler.NewTransactionHandler(transactionUseCase)
	budgetGoalHandler := handler.NewBudgetGoalHandler(budgetGoalUseCase)

	// Routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	v1.Use(middleware.AuthMiddleware(os.Getenv("JWT_SECRET")))

	categories := v1.Group("/categories")
	categories.Post("/", categoryHandler.Create)
	categories.Get("/", categoryHandler.List)

	transactions := v1.Group("/transactions")
	transactions.Post("/", transactionHandler.Create)
	transactions.Get("/", transactionHandler.List)

	goals := v1.Group("/goals")
	goals.Post("/", budgetGoalHandler.CreateGoal)
	goals.Get("/", budgetGoalHandler.ListGoals)

	users := v1.Group("/users")
	users.Get("/profile", userHandler.GetProfile)
	users.Put("/profile", userHandler.UpdateProfile)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "up"})
	})

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
