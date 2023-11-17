package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snapfood/domain"
	"snapfood/logic"
	"strconv"
	"time"
)

type ServicesInterface interface {
	DelayOrder(c *gin.Context)
	AssignDelayReport(c *gin.Context)
	ReportDelayReport(c *gin.Context)
}

type Services struct {
	DelayServicesLogic logic.DelayLogicInterface
}

func NewServices(DelayLogic logic.DelayLogicInterface) *Services {
	return &Services{DelayServicesLogic: DelayLogic}
}

// DelayOrder an api for report a Delay order
func (s *Services) DelayOrder(c *gin.Context) {
	orderIDStr, _ := c.GetPostForm("order_id")
	OrderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order id is not valid"})
		return
	}
	useIDStr, _ := c.GetPostForm("user_id")
	userID, err := strconv.Atoi(useIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is not valid"})
		return
	}
	newDeliveryTime, err := s.DelayServicesLogic.DelayOrder(userID, OrderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newDeliveryTime != 0 {
		c.JSON(http.StatusOK, domain.SubmitOrderDelayResp{
			Message:         "Courier on the way",
			NewDeliveryTime: newDeliveryTime,
		})
		return
	}
	c.JSON(http.StatusOK, domain.SubmitOrderDelayResp{
		Message: "order delay submitted",
	})
	return
}

// AssignDelayReport an api for assign a Delay report to an agent
func (s *Services) AssignDelayReport(c *gin.Context) {
	agentIDStr, _ := c.GetPostForm("agent_id")
	AgentID, err := strconv.Atoi(agentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent id is not valid"})
		return
	}

	delayReport, err := s.DelayServicesLogic.AssignDelayReport(AgentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, delayReport)
	return
}

// ReportDelayReport an api for get report of order delays by vendors
func (s *Services) ReportDelayReport(c *gin.Context) {
	vendorIDStr, _ := c.GetQuery("vendor_id")
	VendorID, err := strconv.Atoi(vendorIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "vendor id is not valid"})
		return
	}
	fromStr := c.Query("from")
	from, err := time.Parse(time.RFC3339, fromStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	toStr := c.Query("to")
	to, err := time.Parse(time.RFC3339, toStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	reports, err := s.DelayServicesLogic.GetVendorOrdersDelay(VendorID, from, to)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reports)
	return
}
