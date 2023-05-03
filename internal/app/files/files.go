package files

import (
	"strings"
	"os"

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

func WritePDF(path string, data string) error {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.Row(10, func() {
		m.Col(5, func() {
			m.Text(data, props.Text{
				Top:         12,
				Size:        20,
				Extrapolate: false,
			})
		})
	})

	m.SetBorder(false)

	err := m.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	return nil
}

func WriteErrorLog(errorText string) error {
	f, err := os.OpenFile("outdata/errors/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
		return err
    }
    if _, err := f.Write([]byte(errorText)); err != nil {
        return err
    }
    if err := f.Close(); err != nil {
        return err
    }
	return nil
}
