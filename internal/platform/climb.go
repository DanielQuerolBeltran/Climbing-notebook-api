package entities

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type ClimbId struct {
	value string
}

func NewClimbId() ClimbId {
	return ClimbId{value: uuid.NewString()}
}

func (id ClimbId) String() string {
	return id.value
}

type ClimbDescription struct {
	value string
}

func newClimbDescription(value string) ClimbDescription {
	return ClimbDescription{value: value}
} 

func (description ClimbDescription) String() string {
	return description.value
}

var ErrInvalidDate = errors.New("invalid date format")

type ClimbDate struct {
	value time.Time
}

func newClimbDate(value string) (ClimbDate, error) {
	date, err := time.Parse(time.RFC3339, value)

	if err != nil {
		return ClimbDate{}, ErrInvalidDate
	}
	return ClimbDate{value: date}, nil
}

func (date ClimbDate) String() string {
	return date.value.Format(time.RFC3339)
}

var ErrInvalidGrade = errors.New("the field Grade can not be empty")

type ClimbGrade struct {
	value string
}

func newClimbGrade(value string) (ClimbGrade, error) {
	return ClimbGrade{value}, nil
}

func (grade ClimbGrade) String() string {
	return grade.value
}

type ClimbArea struct {
	value string
}

func newClimbArea(value string) (ClimbArea, error) {
	return ClimbArea{value}, nil
}

func (area ClimbArea) String() string {
	return area.value
}

type Climb struct {
	id          ClimbId
	date        ClimbDate
	grade       ClimbGrade
	description ClimbDescription
	area        ClimbArea
}

type ClimbRepository interface {
	Save(ctx context.Context, climb Climb) error
	Get(ctx context.Context, id ClimbId) ([]Climb, error)
}

func NewClimb(id string, date string, grade string, description string, area string) (Climb, error) {
	idV0 :=  ClimbId{value: id}
	
	dateV0, err := newClimbDate(date)
	if err != nil {
		return Climb{}, err
	}

	gradeV0, err := newClimbGrade(grade)
	if err != nil {
		return Climb{}, err
	}

	areaV0, err := newClimbArea(area)
	if err != nil {
		return Climb{}, err
	}

	return Climb{
		id: idV0,
		date: dateV0,
		grade: gradeV0,
		description: newClimbDescription(description),
		area: areaV0,
	}, nil
}

func (c Climb) Id() ClimbId {
	return c.id
}

func (c Climb) Date() ClimbDate {
	return c.date
}

func (c Climb) Grade() ClimbGrade {
	return c.grade
}

func (c Climb) Description() ClimbDescription {
	return c.description
}

func (c Climb) Area() ClimbArea {
	return c.area
}