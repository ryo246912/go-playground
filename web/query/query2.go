package query

import (
	"context"
	"fmt"

	"github.com/ryo246912/playground-go-web/models"
	"github.com/uptrace/bun"
)

func Query2(db *bun.DB) {
	ctx := context.Background()

	limit := 100
	films := make([]models.Film, limit)
	if err := db.NewRaw("SELECT film_id FROM ? LIMIT 100", bun.Ident("film"), limit).Scan(ctx, &films); err != nil {
		fmt.Println(err)
	}

	for _, f := range films {
		fmt.Printf("file_id: %d", f.FilmID)
	}
}
