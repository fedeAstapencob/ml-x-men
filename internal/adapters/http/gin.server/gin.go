package gin_server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type GinServerMode int

const (
	DebugMode GinServerMode = iota
	ReleaseMode
	TestMode
)

// GinServer : the struct gathering all the server details
type GinServer struct {
	port   int
	Router *gin.Engine
}

// NewServer
func NewServer(port int, ginServerMode string) GinServer {
	s := GinServer{}
	s.port = port

	s.Router = gin.New()

	gin.SetMode(ginServerMode)

	s.Router.Use(gin.Recovery())

	SetCors(s.Router, "*")

	return s
}

// SetCors is a helper to set current engine cors
func SetCors(engine *gin.Engine, allowedOrigins string) {
	engine.Use(cors.Middleware(cors.Config{
		Origins:         allowedOrigins,
		Methods:         strings.Join([]string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions, http.MethodPatch}, ","),
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
}

// Start the server
func (s GinServer) Start() {
	err := s.Router.Run(":" + strconv.Itoa(int(s.port)))
	if err != nil {
		panic(fmt.Sprintf("Couldn't initialize the API, error %s", err.Error()))
	}
}
