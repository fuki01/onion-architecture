package main

import (
	"fmt"
	"os"

	"github.com/fuki01/onion-architecture/domain/task"
	"github.com/fuki01/onion-architecture/infrastructure"
	"github.com/fuki01/onion-architecture/infrastructure/config"
	"github.com/fuki01/onion-architecture/presentation/controller"
	"github.com/fuki01/onion-architecture/presentation/router"
	"github.com/fuki01/onion-architecture/usecase"
	"github.com/joho/godotenv"
)

func loadEnv(envfile string) {
	err := godotenv.Load(envfile)
	fmt.Println("err: ", err)
	if err != nil {
		panic("no env file")
	}
}

func main() {
	// 環境変数を読み込む
	loadEnv(".env")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || dbname == "" {
		panic("failed to load env")
	}

	db, err := config.NewDatabase(user, pass, host, dbname).Connect()
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&task.Task{})
	if err != nil {
		panic("failed to migrate database")
	}

	// TaskRepositoryの実装を初期化
	taskRepository := infrastructure.NewArticlePersistence(db)

	// UseCaseを初期化
	taskUseCase := usecase.NewTaskUsecase(taskRepository)

	// Controllerを初期化
	taskController := controller.NewTaskController(taskUseCase)

	// ルーティングを設定
	r := router.SetupRouter(taskController)

	// サーバーを起動
	r.Run(":8080")
}
