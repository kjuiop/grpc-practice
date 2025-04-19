package network

import (
	"github.com/gin-gonic/gin"
	"grpc-practice/types"
	"net/http"
)

func (n *Network) Login(c *gin.Context) {
	var req types.LoginReq

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := n.service.CreateAuth(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (n *Network) Verify(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}
