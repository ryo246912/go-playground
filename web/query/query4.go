package query

import (
	"context"
	"fmt"

	"github.com/ryo246912/playground-go-web/models"
	"github.com/uptrace/bun"
)

func Query4(db *bun.DB) {
	ctx := context.Background()
	var err error

	limit := 10

	// relationを設定
	// https://bun.uptrace.dev/guide/relations.html

	// N-N
	highRateFilms := make([]models.Film, limit)
	err = db.NewSelect().
		Model(&highRateFilms).
		Where("rental_rate > ?", 3).
		// Relationの設定
		// Relation("Categoris").
		// Relationをネストしたい場合はRelation("<Relation1>.<Relation2>")みたいな形でJOINする
		Relation("Categoris.Ct"). // FilmCategoryのCategoryもJOIN
		Limit(limit).
		Scan(ctx)

	if err != nil {
		fmt.Println(err)
	}

	for _, f := range highRateFilms {
		fmt.Println(f.FilmID, f.Title, f.RentalRate)
		for _, c := range f.Categoris {
			fmt.Println(f.FilmID, c.CategoryID, c.Ct.Name)
		}
	}
}
