package files

import (
	"strings"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
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

func WritePDF(path string, data [][]string) error {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	for _, str := range data {
		m.Row(10, func() {
			m.Col(5, func() {
				m.Text(str[0], props.Text{
					Top:         12,
					Size:        10,
					Extrapolate: false,
				})
			})
			m.ColSpace(2)
			m.Col(5, func() {
				m.Text(str[1], props.Text{
					Top:         12,
					Size:        10,
					Extrapolate: false,
				})
			})
		})
	}

	m.SetBorder(false)

	err := m.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	return nil
}
