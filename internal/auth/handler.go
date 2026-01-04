package auth

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	service *Service
}

func NewHandler(service *Service) *Handler{
	return &Handler{service : service}
}

//Register Endpoint 
func (h *Handler) Register (c *gin.Context){
	var req struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err:=c.ShouldBindBodyWithJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Bad request body"})
		return
	}

	if req.Email=="" || req.Password==""{
		c.JSON(http.StatusBadRequest,gin.H{"error":"email and password are required."})
		return
	}

	if err:=h.service.Register(req.Email,req.Password);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"user registered succesfully"})


}

//LOGIN
func (h *Handler) Login (c *gin.Context){
	var req struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
	if err:=c.ShouldBindBodyWithJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Bad request body"})
		return
	}
	if req.Email=="" || req.Password==""{
		c.JSON(http.StatusBadRequest,gin.H{"error":"email and password are required."})
		return
	}
	token,err:=h.service.Login(req.Email,req.Password)
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":err.Error()})
	}

	c.JSON(http.StatusOK,gin.H{"token":token})


}