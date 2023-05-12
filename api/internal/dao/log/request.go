// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/log/internal"
)

// internalRequestDao is internal type for wrapping internal DAO implements.
type internalRequestDao = *internal.RequestDao

// requestDao is the data access object for table log_request.
// You can define custom methods on it to extend its functionality as you wish.
type requestDao struct {
	internalRequestDao
}

var (
	// Request is globally public accessible object for table log_request operations.
	Request = requestDao{
		internal.NewRequestDao(),
	}
)

// Fill with you ideas below.
