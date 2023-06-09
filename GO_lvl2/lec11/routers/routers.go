package routers

import (
	"github.com/babaevs2023/GO/GO_lvl2/lec11/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default() // аналог mux.NewRouter()
	//В Gin принято группировать ресурсы
	apiV1Group := router.Group("/api/v1") // Указываем префикс
	{
		apiV1Group.GET("article", handlers.GetAllArticles)
		apiV1Group.POST("article", handlers.PostNewArticle)
		apiV1Group.GET("article/:id", handlers.GetArticleById)
		apiV1Group.PUT("article/:id", handlers.UpdateArticleById)
		apiV1Group.DELETE("article/:id", handlers.DeleteArticleById)
	}

	//... инициализируем все остальное
	return router

}
