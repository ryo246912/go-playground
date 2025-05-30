package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// InventoryInStock calls the stored function 'sakila.inventory_in_stock(int) tinyint' on db.
func InventoryInStock(ctx context.Context, db DB, pInventoryID int) (bool, error) {
	// call sakila.inventory_in_stock
	const sqlstr = `SELECT sakila.inventory_in_stock(?)`
	// run
	var r0 bool
	logf(sqlstr, pInventoryID)
	if err := db.QueryRowContext(ctx, sqlstr, pInventoryID).Scan(&r0); err != nil {
		return false, logerror(err)
	}
	return r0, nil
}
