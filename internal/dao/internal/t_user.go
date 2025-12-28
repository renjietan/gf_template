// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserDao is the data access object for the table t_user.
type TUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TUserColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TUserColumns defines and stores column names for the table t_user.
type TUserColumns struct {
	Id       string //
	Name     string //
	FId      string //
	DeleteAt string //
	CreateAt string //
	UpdateAt string //
	InfoId   string //
}

// tUserColumns holds the columns for the table t_user.
var tUserColumns = TUserColumns{
	Id:       "id",
	Name:     "name",
	FId:      "f_id",
	DeleteAt: "delete_at",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	InfoId:   "infoId",
}

// NewTUserDao creates and returns a new DAO object for table data access.
func NewTUserDao(handlers ...gdb.ModelHandler) *TUserDao {
	return &TUserDao{
		group:    "default",
		table:    "t_user",
		columns:  tUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TUserDao) Columns() TUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TUserDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
