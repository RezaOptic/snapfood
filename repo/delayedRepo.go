package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"snapfood/constants/PrivateErrors"
	"snapfood/domain"
	"snapfood/repo/queries"
	"time"
)

type DelayRepoInterface interface {
	GetDelayReport(OrderID int) (*domain.DelayReport, error)
	SubmitNewDelayOrder(OrderID, VendorID, UserID int) error
	CountOfAgentReports(AgentID int) (int, error)
	AssignDelayReport(AgentID int) (*domain.DelayReport, error)
	GetVendorOrdersDelay(VendorID int, From, To time.Time) ([]domain.DelayReport, error)
}

type DelayRepo struct {
	DB *sql.DB
}

func NewDelayRepo(DB *sql.DB) *DelayRepo {
	return &DelayRepo{DB: DB}
}

func (d *DelayRepo) GetDelayReport(OrderID int) (*domain.DelayReport, error) {
	var report domain.DelayReport
	err := d.DB.QueryRow(queries.FindUnhandledDelayReport, OrderID).Scan(&report.ReportID, &report.OrderID, &report.UserID, &report.AgentID, &report.DelayReason, &report.ReportTime, &report.State)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("error in find order: %v", err)
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, PrivateErrors.NotFound
	}
	return &report, nil
}

func (d *DelayRepo) SubmitNewDelayOrder(OrderID, VendorID, UserID int) error {
	_, err := d.DB.Exec(queries.SubmitNewDelayOrder, OrderID, UserID, VendorID)
	if err != nil {
		return fmt.Errorf("error in submit new delay order: %v", err)
	}
	return err
}

func (d *DelayRepo) CountOfAgentReports(AgentID int) (int, error) {
	var count int
	err := d.DB.QueryRow(queries.CountOfAgentReports, AgentID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error in find order: %v", err)
	}
	return count, nil
}

func (d *DelayRepo) AssignDelayReport(AgentID int) (*domain.DelayReport, error) {
	trx, err := d.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("error in begin transaction: %v", err)
	}
	var report domain.DelayReport
	err = trx.QueryRow(queries.GetFirstDelayReport).Scan(&report.ReportID, &report.OrderID, &report.UserID, &report.VendorID, &report.AgentID, &report.DelayReason, &report.ReportTime, &report.State)
	if err != nil {
		return nil, fmt.Errorf("error in find order: %v", err)
	}
	_, err = trx.Exec(queries.AssignReportToAgent, AgentID, report.ReportID)
	if err != nil {
		return nil, fmt.Errorf("error in submit new delay order: %v", err)
	}
	err = trx.Commit()
	if err != nil {
		trx.Rollback()
		return nil, err
	}
	return &report, nil
}

func (d *DelayRepo) GetVendorOrdersDelay(VendorID int, From, To time.Time) ([]domain.DelayReport, error) {
	var reports []domain.DelayReport
	rows, err := d.DB.Query(queries.GetVendorOrdersDelay, VendorID, From, To)
	if err != nil {
		return nil, fmt.Errorf("error in submit new delay order: %v", err)
	}
	for rows.Next() {
		var report domain.DelayReport
		err := rows.Scan(&report.ReportID, &report.OrderID, &report.UserID, &report.VendorID, &report.AgentID, &report.DelayReason, &report.ReportTime, &report.State)
		if err != nil {
			return nil, fmt.Errorf("error in submit new delay order: %v", err)
		}
		reports = append(reports, report)
	}

	return reports, nil
}
