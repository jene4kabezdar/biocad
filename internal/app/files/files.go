package files

import (
	"strings"

	"github.com/xuri/excelize/v2"
)

func ParseStartData() ([][]string, error) {
	f, err := excelize.OpenFile("indata/startdata.xlsx")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	rows, err := f.GetRows("Исходные данные")
	if err != nil {
		return nil, err
	}

	mRows := ProcessingData(rows)

	return mRows, nil
}

func ProcessingData(rows [][]string) [][]string {
	mRows := make([][]string, len(rows))
	for i, row := range rows {
		mRow := make([]string, 15)
		for j, v := range row {
			mRow[j] = strings.TrimSpace(v)
		}
		mRows[i] = mRow
	}
	return mRows
}
