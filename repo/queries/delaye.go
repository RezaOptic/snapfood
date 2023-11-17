package queries

var (
	FindUnhandledDelayReport = `SELECT report_id,order_id,user_id,agent_id, delay_reason,report_time,state  FROM DelayReports WHERE order_id = $1 and state = 'unhandled'`
	SubmitNewDelayOrder      = `INSERT INTO DelayReports (order_id, user_id,vendor_id, delay_reason, report_time, state)
		VALUES ($1, $2,$3, '', NOW(), 'unhandled')`
	CountOfAgentReports           = `SELECT COUNT(*) FROM DelayReports WHERE agent_id = $1 and state = 'on_process'`
	GetFirstDelayReport           = `SELECT report_id,order_id,user_id,vendor_id,agent_id, delay_reason,report_time,state  FROM DelayReports WHERE state = 'unhandled' order by report_time desc limit 1`
	AssignReportToAgent           = `UPDATE DelayReports SET agent_id = $1, state = 'on_process' where report_id = $2`
	FindUnhandledAgentDelayReport = `SELECT report_id,order_id,user_id,agent_id, delay_reason,report_time,state  FROM DelayReports WHERE agent_id =$1 and state = 'on_process' `
	GetVendorOrdersDelay          = `SELECT report_id,order_id,user_id,vendor_id,agent_id, delay_reason,report_time,state  FROM DelayReports WHERE vendor_id =$1 and report_time >= $2 and report_time <= $3`
)
