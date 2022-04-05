package pkg

import (
	"strings"
)

func ValidateHistoryParam(id int64, orderBy, sort *string, limit, offset *int64) error {

	err := ValidateId(id)
	if err != nil {
		return err
	}

	err = ValidateOrderBy(orderBy)
	if err != nil {
		return err
	}

	err = ValidateSort(sort)
	if err != nil {
		return err
	}

	err = ValidateLimit(limit)
	if err != nil {
		return err
	}

	err = ValidateOffset(offset)
	if err != nil {
		return err
	}

	return nil
}

func ValidateId(id int64) error {

	if id <= 0 {
		return ErrInvalidInput
	}

	return nil
}

func ValidateOrderBy(orderBy *string) error {
	if *orderBy == "" {
		*orderBy = "date"
		return nil
	}

	if *orderBy != "date" && *orderBy != "amount"{
		return ErrInvalidInput
	}

	return nil
}

func ValidateSort(sort *string) error {
	if *sort == "" {
		*sort = "ASC"
		return nil
	}

	*sort = strings.ToUpper(*sort)

	if *sort != "ASC" && *sort != "DESC" {
		return ErrInvalidInput
	}

	return nil
}

func ValidateLimit(limit *int64) error {
	if *limit < 0 {
		return ErrInvalidInput
	}

	if *limit == 0 {
		*limit = 10
		return nil
	}

	return nil
}

func ValidateOffset(offset *int64) error {
	if *offset < 0 {
		return ErrInvalidInput
	}

	return nil
}