package postgres

import (
	"database/sql/driver"
	"encoding/json"
)

// Если поле в БД имеет тип данных JSONB
// В таком случае для работы в коде с этим полем, нужно прописать
// Пользовательские преобразования JSON
// АКТУАЛЬНО для: native go, pgxpool или gorm

type JSONBData map[string]string

func (j JSONBData) Value() (driver.Value, error) {
	value, err := json.Marshal(&j)
	return value, err
}

func (j *JSONBData) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}
