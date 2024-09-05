package api

import (
	handlerC "api/api/handler"

	_ "api/docs"
	_ "api/genproto/auth"
	_ "api/genproto/gym"

	"api/api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Budgeting SYSTEM API
// @version 1.0
// @description Developing a platform that helps users track their spending, set a budget and manage their financial goals
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin( /*AuthConn, */ budgetingConn *grpc.ClientConn) *gin.Engine {
	budgeting := handlerC.NewBudgetingHandler(budgetingConn)
	ca := middleware.CasbinEnforcer()
	err := ca.LoadPolicy()
	if err != nil{
		panic(err)
	}
	router := gin.Default()

	// sw := router.Group("/")
	// router.Use(middleware.JWTMiddleware(), middleware.CasbinMiddleware(ca))

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	account := router.Group("/gym").Use(middleware.JWTMiddleware(),middleware.CasbinMiddleware(ca))
	{
		account.POST("/create", budgeting.CreateGym)
		account.PUT("/update", budgeting.UpdateGym)
		account.GET("/get/:id", budgeting.GetGym)
		account.DELETE("/delete/:id", budgeting.DeleteGym)
		account.GET("/list", budgeting.ListGyms)
	}
	budget := router.Group("/facility").Use(middleware.JWTMiddleware(),middleware.CasbinMiddleware(ca))
	{
		budget.POST("/create", budgeting.CreateFacility)
		budget.PUT("/update", budgeting.UpdateFacility)
		budget.GET("/get/:id", budgeting.GetFacility)
		budget.DELETE("/delete/:id", budgeting.DeleteFacility)
		budget.GET("/list", budgeting.ListFacilitys)
	}
	category := router.Group("/gymfacility").Use(middleware.JWTMiddleware(),middleware.CasbinMiddleware(ca))
	{
		category.POST("/create", budgeting.CreateGymFacility)
		category.PUT("/update", budgeting.UpdateGymFacility)
		category.GET("/get/:id", budgeting.GetGymFacility)
		category.DELETE("/delete/:id", budgeting.DeleteGymFacility)
		category.GET("/list", budgeting.ListGymFacilitys)
	}

	return router
}
