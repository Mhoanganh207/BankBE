package api

import (
	"github.com/Mhoanganh207/BankBE/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	config       *util.Config
	db           *gorm.DB
	router       *gin.Engine
	tokenService util.Generator
}

func NewServer(db *gorm.DB) *Server {
	config := util.LoadConfig()
	server := &Server{
		db: db,
	}
	server.config = &config
	server.tokenService = util.NewGeneratorToken(config.SecretKey)
	server.addRouter()
	return server
}

func (s *Server) Start() {
	s.router.Run("localhost:" + s.config.Port)
}
