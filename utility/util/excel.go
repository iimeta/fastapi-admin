package util

import (
	"fmt"
	"reflect"

	"github.com/xuri/excelize/v2"
)

const COL_A = 65

func ExcelExport(sheetName string, titleCols []string, colFieldMap map[string]string, values []any, saveExcelFilePath string) error {

	f := excelize.NewFile()

	if err := f.SetSheetName(f.GetSheetName(0), sheetName); err != nil {
		return err
	}

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 14,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		return err
	}

	for i, title := range titleCols {

		if err = f.SetCellValue(sheetName, fmt.Sprintf("%c1", COL_A+i), title); err != nil {
			return err
		}

		if err = f.SetCellStyle(sheetName, fmt.Sprintf("%c1", COL_A+i), fmt.Sprintf("%c1", COL_A+i), style); err != nil {
			return err
		}
	}

	if err = f.SetColWidth(sheetName, fmt.Sprintf("%c", COL_A), fmt.Sprintf("%c", COL_A+len(titleCols)-1), 24); err != nil {
		return err
	}

	for i := 0; i < len(values); i++ {
		value := reflect.ValueOf(values[i]).Elem()
		for j := 0; j < len(titleCols); j++ {
			if err = f.SetCellValue(sheetName, fmt.Sprintf("%c%d", COL_A+j, 2+i), value.FieldByName(colFieldMap[titleCols[j]])); err != nil {
				return err
			}
		}
	}

	if err = f.SaveAs(saveExcelFilePath); err != nil {
		return err
	}

	return nil
}
