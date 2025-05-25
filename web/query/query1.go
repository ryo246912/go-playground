package query

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

// 疎通確認
func Query1(db *bun.DB) {
	ctx := context.Background()

	var num int
	if err := db.NewSelect().ColumnExpr("1").Scan(ctx, &num); err != nil {
		fmt.Println(err)
	}

	fmt.Println(num)
}
