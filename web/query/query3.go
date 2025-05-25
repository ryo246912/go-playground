package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ryo246912/playground-go-web/models"
	"github.com/uptrace/bun"
)

func Query3(db *bun.DB) {
	ctx := context.Background()
	var err error

	// WHERE
	var film models.Film
	err = db.NewSelect().
		// SELECT句
		Model(&film).
		// WHERE句
		Where("film_id = ?", 10).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("record not found")
		} else {
			panic(err)
		}
	}
	fmt.Println(film.FilmID, film.Title)

	limit := 10
	// like句
	films := make([]models.Film, limit)
	err = db.NewSelect().
		// SELECT句
		Model(&films).
		// WHERE句
		Where("title LIKE ?", "%EGG").
		// limit句
		Limit(limit).
		Scan(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range films {
		fmt.Println(f.FilmID, f.Title)
	}

	// 大小比較
	highRateFilms := make([]models.Film, limit)
	err = db.NewSelect().
		Model(&highRateFilms).
		Where("rental_rate > ?", 3).Limit(limit).
		Scan(ctx)

	if err != nil {
		fmt.Println(err)
	}

	for _, f := range highRateFilms {
		fmt.Println(f.FilmID, f.Title, f.RentalRate)
	}

	// IN句
	rateFilms := make([]models.Film, limit)
	err = db.NewSelect().
		Model(&rateFilms).
		Where("rating IN (?)", bun.In([]models.Rating{models.RatingNc17, models.RatingG})).
		Limit(limit).
		Scan(ctx)

	if err != nil {
		fmt.Println(err)
	}

	for _, f := range rateFilms {
		fmt.Println(f.FilmID, f.Title, f.Rating)
	}

}
