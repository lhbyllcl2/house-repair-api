package repairmodel

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	repairOrderBillFieldNames          = builder.RawFieldNames(&RepairOrderBill{})
	repairOrderBillRows                = strings.Join(repairOrderBillFieldNames, ",")
	repairOrderBillRowsExpectAutoSet   = strings.Join(stringx.Remove(repairOrderBillFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	repairOrderBillRowsWithPlaceHolder = strings.Join(stringx.Remove(repairOrderBillFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	RepairOrderBillModel interface {
		Insert(data RepairOrderBill) (sql.Result, error)
		FindOne(id int64) (*RepairOrderBill, error)
		Update(data RepairOrderBill) error
		Delete(id int64) error
	}

	defaultRepairOrderBillModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RepairOrderBill struct {
		Id                       int64        `db:"id"`                         //
		OrderNo                  int64        `db:"order_no"`                   // 订单号
		OutTradeNo               string       `db:"out_trade_no"`               // 发起支付时的订单号
		TotalFee                 float64      `db:"total_fee"`                  // 总费用
		Status                   int64        `db:"status"`                     // 状态  1-预估状态  2-实际状态  3-账单已发送 4-支付完成
		PayType                  string       `db:"pay_type"`                   // 支付方式
		PayTime                  sql.NullTime `db:"pay_time"`                   // 支付时间
		AdjunctiveFee            float64      `db:"adjunctive_fee"`             // 附加费
		AdjunctiveFeeDescription string       `db:"adjunctive_fee_description"` // 附加费说明
		DiscountTicketSn         string       `db:"discount_ticket_sn"`         // 使用优惠券的sn码
		DiscountAmount           float64      `db:"discount_amount"`            // 抵扣金额
		CreateTime               time.Time    `db:"create_time"`                // 创建时间
		UpdateTime               time.Time    `db:"update_time"`                // 更新时间
	}
)

func NewRepairOrderBillModel(conn sqlx.SqlConn) RepairOrderBillModel {
	return &defaultRepairOrderBillModel{
		conn:  conn,
		table: "`repair_order_bill`",
	}
}

func (m *defaultRepairOrderBillModel) Insert(data RepairOrderBill) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, repairOrderBillRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.OrderNo, data.OutTradeNo, data.TotalFee, data.Status, data.PayType, data.PayTime, data.AdjunctiveFee, data.AdjunctiveFeeDescription, data.DiscountTicketSn, data.DiscountAmount)
	return ret, err
}

func (m *defaultRepairOrderBillModel) FindOne(id int64) (*RepairOrderBill, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", repairOrderBillRows, m.table)
	var resp RepairOrderBill
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRepairOrderBillModel) Update(data RepairOrderBill) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, repairOrderBillRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.OrderNo, data.OutTradeNo, data.TotalFee, data.Status, data.PayType, data.PayTime, data.AdjunctiveFee, data.AdjunctiveFeeDescription, data.DiscountTicketSn, data.DiscountAmount, data.Id)
	return err
}

func (m *defaultRepairOrderBillModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
