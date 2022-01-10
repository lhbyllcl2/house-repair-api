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
	repairOrderFieldNames          = builder.RawFieldNames(&RepairOrder{})
	repairOrderRows                = strings.Join(repairOrderFieldNames, ",")
	repairOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(repairOrderFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	repairOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(repairOrderFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	RepairOrderModel interface {
		Insert(data RepairOrder) (sql.Result, error)
		FindOne(id int64) (*RepairOrder, error)
		FindOneByOrderNo(orderNo int64) (*RepairOrder, error)
		Update(data RepairOrder) error
		Delete(id int64) error
	}

	defaultRepairOrderModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RepairOrder struct {
		Id                 int64        `db:"id"`
		OrderNo            int64        `db:"order_no"`              // 订单号
		UserId             int64        `db:"user_id"`               // 用户id
		RepairStaffId      string       `db:"repair_staff_id"`       // 维修人员编号
		CreateTime         time.Time    `db:"create_time"`           // 创建时间
		PartId             int64        `db:"part_id"`               // 部品id
		PositionId         int64        `db:"position_id"`           // 部位id
		QuestionId         int64        `db:"question_id"`           // 问题id，在维修方案确定后才确定
		Describe           string       `db:"describe"`              // 报修时问题描述
		MajorDescription   string       `db:"major_description"`     // 父部品-子部品-问题
		HopeVisitStartTime sql.NullTime `db:"hope_visit_start_time"` // 期望上门时间-开始
		HopeVisitStartEnd  sql.NullTime `db:"hope_visit_start_end"`  // 期望上门时间-结束
		Phone              string       `db:"phone"`                 // 报修电话号码
		Name               string       `db:"name"`                  // 报修人
		Address            string       `db:"address"`               // 地址
		NodeId             int64        `db:"node_id"`               // 进度ID
		Status             int64        `db:"status"`                // 订单状态 1-正常 0-用户取消  2-关闭
		RepairFrom         int64        `db:"repair_from"`           // 报修入口,1微信报修,2android app报修,3ios app报修,4电话报修
		CheckStaffId       string       `db:"check_staff_id"`        // 审核人员编号
		Longitude          float64      `db:"longitude"`             // 报修时的经度
		Latitude           float64      `db:"latitude"`              // 报修时的纬度
		UpdateTime         time.Time    `db:"update_time"`           // 更新时间
	}
)

func NewRepairOrderModel(conn sqlx.SqlConn) RepairOrderModel {
	return &defaultRepairOrderModel{
		conn:  conn,
		table: "`repair_order`",
	}
}

func (m *defaultRepairOrderModel) Insert(data RepairOrder) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, repairOrderRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.OrderNo, data.UserId, data.RepairStaffId, data.PartId, data.PositionId, data.QuestionId, data.Describe, data.MajorDescription, data.HopeVisitStartTime, data.HopeVisitStartEnd, data.Phone, data.Name, data.Address, data.NodeId, data.Status, data.RepairFrom, data.CheckStaffId, data.Longitude, data.Latitude)
	return ret, err
}

func (m *defaultRepairOrderModel) FindOne(id int64) (*RepairOrder, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", repairOrderRows, m.table)
	var resp RepairOrder
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

func (m *defaultRepairOrderModel) FindOneByOrderNo(orderNo int64) (*RepairOrder, error) {
	var resp RepairOrder
	query := fmt.Sprintf("select %s from %s where `order_no` = ? limit 1", repairOrderRows, m.table)
	err := m.conn.QueryRow(&resp, query, orderNo)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRepairOrderModel) Update(data RepairOrder) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, repairOrderRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.OrderNo, data.UserId, data.RepairStaffId, data.PartId, data.PositionId, data.QuestionId, data.Describe, data.MajorDescription, data.HopeVisitStartTime, data.HopeVisitStartEnd, data.Phone, data.Name, data.Address, data.NodeId, data.Status, data.RepairFrom, data.CheckStaffId, data.Longitude, data.Latitude, data.Id)
	return err
}

func (m *defaultRepairOrderModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
