package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

// Читает данные из cvs файла
func readDataFile(path string) []byte {
	bytesData, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return bytesData
}

// Обрабатывает данные из cvs файла
func scanDataFile(path string) [][]string {

	// Получаем контент из файла
	bytesData := readDataFile(path)
	bytesDataStr := string(bytesData)

	// Удаляем лишние пробелы из данных
	bytesDataWithoutSpaces := strings.ReplaceAll(bytesDataStr, " ", "")

	// Создаем новый CVS Reader
	r := csv.NewReader(strings.NewReader(bytesDataWithoutSpaces))

	// Читаем данные из файла
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

// Удаляет элемент из слайса по индексу
func removeValueForIndex(slice [][]string, i int) [][]string {
	return append(slice[:i], slice[i+1:]...)
}

// Выводит данные из cvs файла в табличном представлении
func printDataFile(path string) {

	// Получаем слайс данных из cvs файла
	data := scanDataFile(path)

	// Заполняем заголовки из первого элемента (слайса) data
	headers := []interface{}{}
	for _, v := range data[0] {
		headers = append(headers, v)
	}

	// Удаляем первый элемент слайса (заголовки)
	data = removeValueForIndex(data, 0)

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)

	// Определеяем автоматическое индексирование
	t.SetAutoIndex(true)

	// Задаем стиль результирующей таблице
	t.SetStyle(table.StyleColoredBright)

	// Добавляем заголовки
	t.AppendHeader(headers)

	// Добавляем строки
	for i := 0; i < len(data); i++ {
		row := table.Row{}
		for j := 0; j < len(data[i]); j++ {
			row = append(row, data[i][j])
		}
		t.AppendRow(row)
	}

	// Рендер таблицы
	t.Render()
}

func main() {
	var path_to_cvs_file string
	fmt.Println("Введите путь до cvs файла:")
	fmt.Scan(&path_to_cvs_file)

	printDataFile(path_to_cvs_file)
}
