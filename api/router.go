package api

import "github.com/gin-gonic/gin"

type Routes struct {
	router *gin.Engine
}

func (server *Server) addRouter() {
	routes := Routes{
		router: gin.Default(),
	}
	routes.addAccountsRoute(server)
	routes.addTransfersRoute(server)
	routes.addUserRoute(server)
	routes.addTokenRoute(server)
	server.router = routes.router
}
