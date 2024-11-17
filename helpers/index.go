package helpers

import (
	"encoding/csv"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"

	"golang.org/x/image/webp"
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

func CalculatePerimeter(imageURL string) (int, error) {
	resp, errr := http.Get(imageURL)
	imageType := imageURL[(strings.LastIndex(imageURL, "."))+1:]
	if errr != nil {
		return 0, errors.New(" Failed to Fetch the Image")
	}
	defer resp.Body.Close()

	var img image.Image
	var err error

	if imageType == "png" {
		img, err = png.Decode(resp.Body)
	} else if imageType == "jpg" {
		img, err = jpeg.Decode(resp.Body)
	} else if imageType == "gif" {
		img, err = gif.Decode(resp.Body)
	} else if imageType == "webp" {
		img, err = webp.Decode(resp.Body)
	} else {
		return 0, errors.New(" Unkown Image Format, Failed to Decode this image ")
	}

	if err != nil {
		return 0, errors.New(" Failed to Decode the Image")
	}

	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	perimeter := 2 * (width + height)

	return perimeter, nil

}
