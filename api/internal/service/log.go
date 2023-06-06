// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IRequest interface {
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, page int, limit int) (list gdb.Result, err error)
	}
)

var (
	localRequest IRequest
)

func Request() IRequest {
	if localRequest == nil {
		panic("implement not found for interface IRequest, forgot register?")
	}
	return localRequest
}

func RegisterRequest(i IRequest) {
	localRequest = i
}
