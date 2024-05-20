package main

import (
    "github.com/fuki01/onion-architecture/infrastructure"
    "github.com/fuki01/onion-architecture/presentation/controller"
    "github.com/fuki01/onion-architecture/router"
    "github.com/fuki01/onion-architecture/usecase"
    "github.com/fuki01/onion-architecture/domain/task"
		"gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // データベース接続を設定
    db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
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
