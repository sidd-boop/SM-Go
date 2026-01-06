package thought

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	service *Service
}

func NewHandler(service *Service) *Handler{
	return &Handler{service: service}
}

//createthought
func (h *Handler) Create(c *gin.Context){
	var req struct{
		Content string `json:"content"`
		Tag string `json:"tag"`
	}
	if err:=c.ShouldBindBodyWithJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid request body"})
	}

	if req.Content==""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"content is req"})
	}

	userID:=c.GetUint("user_id")

	if err:=h.service.Create(req.Content, req.Tag, userID);err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message":"thought created"})
}

//listthoughts
func (h *Handler) List(c *gin.Context){
	thoughts,err :=h.service.List()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,thoughts)
} 

//getsinglethought
func (h *Handler) Get(c *gin.Context){
	idParam:=c.Param("id")
	id, err:=strconv.Atoi(idParam)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid id"})
		return
	}

	thought, err:=h.service.Get(uint(id))
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"couldn't find thought"})
		return
	}
	c.JSON(http.StatusOK,thought)

}