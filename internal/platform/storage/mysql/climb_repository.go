package mysql

import (
	"database/sql"
	"fmt"
	"context"
	"github.com/huandu/go-sqlbuilder"
	"github.com/DanielQuerolBeltran/Climbing-notebook-api/internal/platform"
)


type repository struct {
	table string
	db *sql.DB
}

func NewClimbRepository(db *sql.DB) entities.ClimbRepository {
	return &repository{db: db, table: CLIMB_TABLE}
}

func (r *repository) Save(ctx context.Context, climb entities.Climb) error {
	climbSQLStruct := sqlbuilder.NewStruct(new(sqlClimb))

	query, args := climbSQLStruct.InsertInto(r.table, sqlClimb{
		 Id: climb.Id().String(),
		 Grade: climb.Grade().String(),
		 Area: climb.Area().String(),
		 Description: climb.Description().String(),
		 Date: climb.Date().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist climb on database: %v", err)
	}

	return nil
}

func (r *repository) Get(ctx context.Context, id entities.ClimbId) ([]entities.Climb, error) {
	climbSQLStruct := sqlbuilder.NewStruct(new(sqlClimb))

	selectBuilder := climbSQLStruct.SelectFrom(r.table)
	query, args := selectBuilder.Build()

	rows, err := r.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() } ()
	
	var climbs []entities.Climb
	for rows.Next() {
		sqlClimb := sqlClimb{}

		err := rows.Scan(climbSQLStruct.Addr(&sqlClimb)...)
		if err != nil {
			return nil, err
		}

		climb, err := entities.NewClimb(
			sqlClimb.Id,
			sqlClimb.Date,
			sqlClimb.Grade,
			sqlClimb.Description,
			sqlClimb.Area,
		)

		if err != nil {
			return nil, fmt.Errorf("error trying to fetch climbs from database: %v", err)
		}

		climbs = append(climbs, climb)
	}
	
	return climbs, nil
}
