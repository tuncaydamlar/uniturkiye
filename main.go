package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gosimple/slug"
	"github.com/xuri/excelize/v2"
)

type Parser struct {
	FileName  string
	SheetName string
}

type University struct {
	Code string `json:"code"`
	Name string `json:"name"`
	City string `json:"city"`
	Type string `json:"type"`
}
type Faculty struct {
	Code           string `json:"code"`
	UniversityCode string `json:"universityCode"`
	Name           string `json:"name"`
}
type Department struct {
	UniversityCode string `json:"universityCode"`
	FacultyCode    string `json:"facultyCode"`
	Name           string `json:"name"`
}

func NewYokUniversityParser(fileName, sheetName string) *Parser {
	return &Parser{
		FileName:  fileName,
		SheetName: sheetName,
	}
}

func (p *Parser) parse() error {
	f, err := excelize.OpenFile(p.FileName)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	rows, err := f.GetRows(p.SheetName)
	if err != nil {
		return fmt.Errorf("failed to get rows: %w", err)
	}

	universitiesMap, facultyMap, departmentArr := p.processRows(rows)

	if err := writeJSON("university.json", universitiesMap); err != nil {
		return err
	}
	if err := writeJSON("faculties.json", facultyMap); err != nil {
		return err
	}
	if err := writeJSON("departments.json", departmentArr); err != nil {
		return err
	}
	if err := writeUniversitySql("universities.sql", universitiesMap); err != nil {
		return err
	}
	if err := writeFacultySql("faculties.sql", facultyMap); err != nil {
		return err
	}
	if err := writeDepartmentSql("departments.sql", departmentArr); err != nil {
		return err
	}

	return nil
}

func (p *Parser) processRows(rows [][]string) (map[string]University, map[string]Faculty, []Department) {
	universitiesMap := make(map[string]University)
	facultyMap := make(map[string]Faculty)
	var departmentArr []Department

	for i, row := range rows {
		if i == 0 {
			continue
		}

		universityName := row[0]
		facultyName := row[1]
		departmentName := row[2]
		universityType := row[4]
		universityCity := row[5]

		universityCode := slug.Make(universityName)
		facultyCode := slug.Make(facultyName)

		// Populate universities map

		if _, exists := universitiesMap[universityCode]; !exists {
			universitiesMap[universityCode] = University{
				Code: universityCode,
				Name: universityName,
				Type: universityType,
				City: universityCity,
			}
		}

		// Populate faculty map
		_, ok := facultyMap[facultyCode+"-"+universityCode]
		if !ok {
			facultyMap[facultyCode+"-"+universityCode] = Faculty{
				Code:           facultyCode + "-" + universityCode,
				Name:           facultyName,
				UniversityCode: universityCode,
			}
		}

		// Populate department array
		departmentArr = append(departmentArr, Department{
			UniversityCode: universityCode,
			FacultyCode:    facultyCode + "-" + universityCode,
			Name:           departmentName,
		})
	}

	return universitiesMap, facultyMap, departmentArr
}

func writeJSON(fileName string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := os.WriteFile(fileName, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", fileName, err)
	}

	return nil
}

func writeUniversitySql(fileName string, universities map[string]University) error {
	sqlHead := `
	CREATE TABLE universities(
		id SERIAL,
		code TEXT,
		name TEXT,
		city TEXT,
		type TEXT,
			UNIQUE(code)
	);
	INSERT INTO universities (code, name, city, type) VALUES`

	var values []string
	for _, u := range universities {
		values = append(values, fmt.Sprintf("('%s', '%s', '%s', '%s')", u.Code, u.Name, u.City, u.Type))
	}

	sql := sqlHead + "\n" + strings.Join(values, ",\n") + ";\n"

	if err := ioutil.WriteFile(fileName, []byte(sql), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", fileName, err)
	}

	return nil
}

func writeFacultySql(fileName string, faculties map[string]Faculty) error {
	sqlHead := `
	CREATE TABLE faculties (
		id SERIAL,
		code TEXT,
		name TEXT,
		university_code TEXT,
		UNIQUE(code,university_code),
		FOREIGN KEY (university_code) REFERENCES universities(code)
	);
	INSERT INTO faculties (code, name, university_code) VALUES`

	var values []string
	for _, f := range faculties {
		values = append(values, fmt.Sprintf("('%s', '%s', '%s')", f.Code, f.Name, f.UniversityCode))
	}

	sql := sqlHead + "\n" + strings.Join(values, ",\n") + ";\n"

	if err := ioutil.WriteFile(fileName, []byte(sql), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", fileName, err)
	}

	return nil
}

func writeDepartmentSql(fileName string, departments []Department) error {
	sqlHead := `
	CREATE TABLE departments (
		id SERIAL,
		name TEXT,
		university_code TEXT,
		faculty_code TEXT,
		FOREIGN KEY (university_code) REFERENCES universities(code),
		FOREIGN KEY (faculty_code) REFERENCES faculties(code)
	);
	INSERT INTO departments ( university_code, faculty_code, name) VALUES`

	var values []string
	for _, d := range departments {
		values = append(values, fmt.Sprintf("('%s', '%s', '%s')", d.UniversityCode, d.FacultyCode, d.Name))
	}

	sql := sqlHead + "\n" + strings.Join(values, ",\n") + ";\n"

	if err := ioutil.WriteFile(fileName, []byte(sql), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", fileName, err)
	}

	return nil
}

func main() {

	parser := NewYokUniversityParser("universiteler.xlsx", "Sheet1")

	if err := parser.parse(); err != nil {
		log.Fatalln(err)
	}
	return
}
