package helpers

import (
	"encoding/csv"
	"errors"
	"os"
)

func LoadStoreIds(filePath string, storeIdCache map[string]bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New(" Unable to Locate the CSV File ")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()

	if err != nil {
		return err
	}

	var store_id_col_index = -1

	for i, col := range data[0] {
		if col == "StoreID" {
			store_id_col_index = i
		}
	}

	if store_id_col_index == -1 {
		return errors.New(" StoreID Column Missing in the CSV File")
	}

	for i, row := range data {
		if i == 0 {
			continue
		}
		storeIdCache[row[store_id_col_index]] = true
	}

	return nil
}

func ValidateStoreId(storeId string, storeIdCache map[string]bool) bool {
	_, exists := storeIdCache[storeId]
	return exists
}
