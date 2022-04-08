package routers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"io"
	"kaya-backend/api/v1/controllers"
	"kaya-backend/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"

	"kaya-backend/library/httpserver/ginserver"
	kayaloger "kaya-backend/library/logger"
	"kaya-backend/library/tracing"
	"kaya-backend/utils/helper"

	//!! V1 API List
	activitasAPI "kaya-backend/api/v1/controllers/activitas"
	authAPI "kaya-backend/api/v1/controllers/auth"
	bannerAPI "kaya-backend/api/v1/controllers/banner"
	blogAPI "kaya-backend/api/v1/controllers/blog"
	customerAPI "kaya-backend/api/v1/controllers/customer"
	dashboardAPI "kaya-backend/api/v1/controllers/dashboard"
	propertyAPI "kaya-backend/api/v1/controllers/properties"
	tagAPI "kaya-backend/api/v1/controllers/tag"
	testimonialAPI "kaya-backend/api/v1/controllers/testimonial"
	titipJualAPI "kaya-backend/api/v1/controllers/titip-jual"
	transactionAPI "kaya-backend/api/v1/controllers/transaction"
	uploadMinioAPI "kaya-backend/api/v1/controllers/upload"
)

// ServerEnv ..
type ServerEnv struct {
	ServiceName     string `envconfig:"KAYA_API_SERVICE" default:"KAYA_API_SERVICE"`
	OpenTracingHost string `envconfig:"OPEN_TRACING_HOST" default:"0.0.0.0:4000"`
	DebugMode       string `envconfig:"DEBUG_MODE" default:"debug"`
	ReadTimeout     int    `envconfig:"READ_TIMEOUT" default:"120"`
	WriteTimeout    int    `envconfig:"WRITE_TIMEOUT" default:"120"`
}

var (
	server ServerEnv
)

func init() {
	if err := envconfig.Process("SERVER", &server); err != nil {
		fmt.Println("Failed to get SERVER env:", err)
	}
}

// Server ..
func Server(listenAddress string) error {
	sugarLogger := kayaloger.GetLogger()

	kayaRouter := KayaRouter{}
	kayaRouter.InitTracing()
	kayaRouter.Routers()
	defer kayaRouter.Close()

	err := ginserver.GinServerUp(listenAddress, kayaRouter.Router)

	if err != nil {
		fmt.Println("Error:", err)
		sugarLogger.Error("Error ", zap.Error(err))
		return err
	}

	fmt.Println("Server UP")
	sugarLogger.Info("Server UP ", zap.String("listenAddress", listenAddress))

	return err
}

// KayaRouter ..
type KayaRouter struct {
	Tracer   opentracing.Tracer
	Reporter jaeger.Reporter
	Closer   io.Closer
	Err      error
	GinFunc  gin.HandlerFunc
	Router   *gin.Engine
}

// Routers ..
func (kayaRouter *KayaRouter) Routers() {
	gin.SetMode(server.DebugMode)
	gen := new(models.GeneralModel)
	transactionController := transactionAPI.InitiateTransactionInterface(gen)
	propertyController := propertyAPI.InitiatePropertiesInterface(gen)
	customerController := customerAPI.InitiateCustomerInterface(gen)
	authController := authAPI.InitiateAuthInterface(gen)
	activitasController := activitasAPI.InitiateActivitasInterface(gen)
	dashboardController := dashboardAPI.InitiateCustomerInterface(gen)
	bannerController := bannerAPI.InitiateBannerInterface(gen)
	blogController := blogAPI.InitiateBlogInterface(gen)
	tagController := tagAPI.InitiateTagInterface(gen)
	titipJualController := titipJualAPI.InitiateTitipJualInterface(gen)
	testimonialController := testimonialAPI.InitiateTestimonialInterface(gen)

	// programmatically set swagger info

	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(kayaRouter.GinFunc)
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowOrigins: []string{"*"},
		MaxAge:           86400,
	}))

	kayaAPI := router.Group("/api")
	{
		v1 := kayaAPI.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/login", authController.Login)
				auth.GET("/logout", helper.TokenAuthMiddleware(), authController.Logout)
				auth.GET("/me", helper.TokenAuthMiddleware(), authController.Me)
				auth.POST("/refresh-token", authController.RefreshToken)
				auth.POST("/send-reset-password", authController.SendEmailResetPassword)
				auth.POST("/reset-password", authController.ResetPassword)
				auth.POST("/send-verification", helper.TokenAuthMiddleware(), authController.SendVerification)
			}

			minio := v1.Group("/minio")
			{
				minio.POST("/upload", uploadMinioAPI.Upload)
				minio.POST("/send-email", uploadMinioAPI.Email)
			}

			customer := v1.Group("/customer")
			{
				customer.POST("/register", customerController.Register)
				customer.POST("/upload", helper.TokenAuthMiddleware(), customerController.Upload)
				customer.POST("/verify", helper.TokenAuthMiddleware(), customerController.Verify)
				customer.POST("/update-password", helper.TokenAuthMiddleware(), customerController.UpdatePassword)
				customer.POST("/update-account", helper.TokenAuthMiddleware(), customerController.UpdateCustomerAccount)
			}

			property := v1.Group("/property")
			{
				primary := property.Group("/primary")
				{
					primary.POST("/list", propertyController.List)
					primary.GET("/detail/:slug", propertyController.Detail)
				}

				secondary := property.Group("/secondary")
				{
					secondary.POST("/list", propertyController.ListSecondary)
					secondary.POST("/list-me", helper.TokenAuthMiddleware(), propertyController.ListSecondaryMe)
					secondary.POST("/sum-list-me", helper.TokenAuthMiddleware(), propertyController.SumSecondaryMe)
					secondary.GET("/detail/:slug", propertyController.DetailSecondary)
					secondary.GET("/detail-id/:id", propertyController.DetailSecondaryByID)
					secondary.GET("/detail-list/:slug", propertyController.DetailListSecondary)
				}

				portfolio := property.Group("/portfolio")
				{
					portfolio.POST("/list", propertyController.ListPortolio)
					portfolio.POST("/list-me", helper.TokenAuthMiddleware(), propertyController.ListPortofolioMe)
					portfolio.POST("/count-me", helper.TokenAuthMiddleware(), propertyController.PortfolioCount)
					portfolio.POST("/detail", helper.TokenAuthMiddleware(), propertyController.PortfolioDetail)
				}
			}

			transaction := v1.Group("/transaction")
			{
				transaction.POST("/checkout", helper.TokenAuthMiddleware(), transactionController.Checkout)
				transaction.POST("/checkout-secondary", helper.TokenAuthMiddleware(), transactionController.CheckoutSecondary)
				transaction.POST("/withdrawal", helper.TokenAuthMiddleware(), transactionController.Withdrawal)
				transaction.POST("/topup", helper.TokenAuthMiddleware(), transactionController.Topup)
				transaction.GET("/:noOrder", helper.TokenAuthMiddleware(), transactionController.Detail)
			}

			activitas := v1.Group("/activitas")
			{
				activitas.POST("/history-topup", helper.TokenAuthMiddleware(), activitasController.HistoryTopup)
				activitas.POST("/history-transaction", helper.TokenAuthMiddleware(), activitasController.HistoryTransaction)
			}

			dashboard := v1.Group("/dashboard")
			{
				dashboard.GET("/detail", helper.TokenAuthMiddleware(), dashboardController.Detail)
				dashboard.POST("/tickets", helper.TokenAuthMiddleware(), dashboardController.Tickets)
				dashboard.POST("/save-ticket", helper.TokenAuthMiddleware(), dashboardController.SaveTicket)
				dashboard.POST("/save-ticket-comment", helper.TokenAuthMiddleware(), dashboardController.SaveTicketComment)
				dashboard.GET("/detail-ticket/:id", helper.TokenAuthMiddleware(), dashboardController.DetailTicket)
				dashboard.PUT("/close-ticket/:id", helper.TokenAuthMiddleware(), dashboardController.CloseTicket)
				dashboard.POST("/cancel-sell", helper.TokenAuthMiddleware(), dashboardController.CancelSell)
				dashboard.POST("/sell-lot", helper.TokenAuthMiddleware(), dashboardController.SellLot)
			}

			banner := v1.Group("/banner")
			{
				banner.GET("/", bannerController.Banner)
			}

			testimonial := v1.Group("/testimonial")
			{
				testimonial.GET("/", testimonialController.GetAll)
			}

			tag := v1.Group("/tag")
			{
				tag.POST("/filter", tagController.Filter)
			}

			blog := v1.Group("/blog")
			{
				blog.POST("/filter", blogController.Filter)
				blog.POST("/filter-tag", blogController.FilterBlogTag)
				blog.GET("/detail/:slug", blogController.Detail)
			}

			options := v1.Group("/options")
			{
				options.GET("/province", controllers.ProvinceOptions)
				options.GET("/city", controllers.CityOptions)
				options.GET("/all-city", controllers.AllCityOptions)
				options.GET("/bank", controllers.BankOptions)
			}

			titipJual := v1.Group("/titip-jual")
			{
				titipJual.POST("/save", titipJualController.Save)
			}

			guest := v1.Group("guest")
			{
				guest.POST("/", customerController.SaveGuest)
			}

		}
	}

	kayaRouter.Router = router

}

// CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// InitTracing ..
func (kayaRouter *KayaRouter) InitTracing() {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "PROD"
	}

	tracer, reporter, closer, err := tracing.InitTracing(fmt.Sprintf("%s::%s", server.ServiceName, hostName), server.OpenTracingHost, tracing.WithEnableInfoLog(true))
	if err != nil {
		fmt.Println("Error :", err)
	}
	opentracing.SetGlobalTracer(tracer)

	kayaRouter.Closer = closer
	kayaRouter.Reporter = reporter
	kayaRouter.Tracer = tracer
	kayaRouter.Err = err
	kayaRouter.GinFunc = tracing.OpenTracer([]byte("api-request-"))
}

// Close ..
func (kayaRouter *KayaRouter) Close() {
	kayaRouter.Closer.Close()
	kayaRouter.Reporter.Close()
}
