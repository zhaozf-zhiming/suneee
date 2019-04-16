package handler

import (
	"../common/log"
	"../common/types"
	"../k8s_cli"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//查询合约列表信息
func HandlerGetDeployment(c *gin.Context) {
	var queryInfo types.QueryDeployment
	queryInfo.Namespace = c.PostForm("namespace")
	if queryInfo.Namespace == "" {
		err := errors.New("namespace can not empty")
		log.Logger.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{"status": FailCode, "reason": err.Error()})
		return
	}
	queryInfo.Name = c.PostForm("name")
	if queryInfo.Name == "" {
		err := errors.New("name can not empty")
		log.Logger.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{"status": FailCode, "reason": err.Error()})
		return
	}
	queryInfo.Start, _ = strconv.Atoi(c.PostForm("start"))
	queryInfo.Limit, _ = strconv.Atoi(c.PostForm("limit"))

	rtVal, err := QueryDeploymentList(queryInfo)
	if err != nil {
		log.Logger.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{"status": FailCode, "reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": SuccCode, "code": *rtVal})
}

func QueryDeploymentList(queryInfo types.QueryDeployment) (*types.QueryDeploymentOut, error) {
	k8s_cli.QueryK8sInfo()
	return nil, nil
}
