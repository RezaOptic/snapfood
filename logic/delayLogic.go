package logic

import (
	"errors"
	"fmt"
	"snapfood/constants"
	"snapfood/constants/PrivateErrors"
	"snapfood/domain"
	"snapfood/repo"
	"snapfood/services"
	"time"
)

type DelayLogicInterface interface {
	DelayOrder(UserID, OrderID int) (int, error)
	AssignDelayReport(AgentID int) (*domain.DelayReport, error)
	GetVendorOrdersDelay(VendorID int, From, To time.Time) ([]domain.DelayReport, error)
}

type DelayLogic struct {
	OrderRepo         repo.OrdersRepoInterface
	TripRepo          repo.TripsRepoInterface
	DelayReportRepo   repo.DelayRepoInterface
	ThirdPartyService services.Interface
}

func NewDelayLogic(TripRepo repo.TripsRepoInterface, ThirdPartyService services.Interface, OrderRepo repo.OrdersRepoInterface, DelayReportRepo repo.DelayRepoInterface) *DelayLogic {
	return &DelayLogic{TripRepo: TripRepo, ThirdPartyService: ThirdPartyService, OrderRepo: OrderRepo, DelayReportRepo: DelayReportRepo}
}

func (d *DelayLogic) DelayOrder(UserID, OrderID int) (int, error) {
	order, err := d.CheckOrderDelayed(OrderID)
	if err != nil {
		return 0, fmt.Errorf("error while checking order delayed: %v", err)
	}
	err = d.CheckUnhandledOrderDelayReport(OrderID)
	if err != nil {
		return 0, fmt.Errorf("error while checking unhandled order delay report: %v", err)
	}

	orderTrip, err := d.TripRepo.TripsOrder(OrderID)
	if err != nil && !errors.Is(err, PrivateErrors.NotFound) {
		return 0, fmt.Errorf("error while getting order trip: %v", err)
	}
	if orderTrip != nil && (orderTrip.TripStatus == constants.AtVendor || orderTrip.TripStatus == constants.Assigned || orderTrip.TripStatus == constants.Picked) {
		newDeliveryTime, err := d.ThirdPartyService.GetDeliveryTime()
		if err != nil {
			return 0, fmt.Errorf("error while getting delivery time: %v", err)
		}
		return newDeliveryTime, nil
	}
	err = d.DelayReportRepo.SubmitNewDelayOrder(OrderID, UserID, order.VendorID)
	if err != nil {
		return 0, fmt.Errorf("error while submitting new delay order: %v", err)
	}
	return 0, nil
}

func (d *DelayLogic) CheckOrderDelayed(OrderID int) (*domain.Order, error) {
	order, err := d.OrderRepo.GetOrders(OrderID)
	if err != nil {
		return nil, fmt.Errorf("error while getting order: %v", err)
	}
	orderTime, err := time.Parse(constants.TimeLayout, order.OrderTime)
	if err != nil {
		return nil, fmt.Errorf("error while parsing order time: %v", err)
	}

	if orderTime.Add(time.Duration(order.TimeDelivery) * time.Minute).Before(time.Now()) {
		return nil, fmt.Errorf("order is not Delay")
	}
	return order, nil
}

func (d *DelayLogic) CheckUnhandledOrderDelayReport(OrderID int) error {
	report, err := d.DelayReportRepo.GetDelayReport(OrderID)
	if err != nil && !errors.Is(err, PrivateErrors.NotFound) {
		return fmt.Errorf("error while getting delay report: %v", err)
	}
	if report != nil {
		return fmt.Errorf("order is already delayed")
	}
	return nil
}

func (d *DelayLogic) AssignDelayReport(AgentID int) (*domain.DelayReport, error) {
	err := d.CheckAgentIsFree(AgentID)
	if err != nil {
		return nil, fmt.Errorf("error while checking agent is free: %v", err)
	}
	reportInfo, err := d.AssignReportToAgent(AgentID)
	if err != nil {
		return nil, fmt.Errorf("error while checking unhandled order delay report: %v", err)
	}

	return reportInfo, nil
}

func (d *DelayLogic) CheckAgentIsFree(AgentID int) error {
	count, err := d.DelayReportRepo.CountOfAgentReports(AgentID)
	if err != nil {
		return fmt.Errorf("error while getting delay report: %v", err)
	}
	if count > 0 {
		return fmt.Errorf("agent is not free")
	}

	return nil
}

func (d *DelayLogic) AssignReportToAgent(AgentID int) (*domain.DelayReport, error) {
	delayReport, err := d.DelayReportRepo.AssignDelayReport(AgentID)
	if err != nil {
		return nil, fmt.Errorf("error while getting delay report: %v", err)
	}
	return delayReport, nil
}

func (d *DelayLogic) GetVendorOrdersDelay(VendorID int, From, To time.Time) ([]domain.DelayReport, error) {
	reports, err := d.DelayReportRepo.GetVendorOrdersDelay(VendorID, From, To)
	if err != nil {
		return nil, fmt.Errorf("error while getting delay report: %v", err)
	}
	return reports, nil
}
