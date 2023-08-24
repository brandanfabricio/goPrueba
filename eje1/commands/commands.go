package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task/commands/exponses"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {

	fmt.Print("-> ")
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	str = strings.TrimSpace(strings.Replace(str, "\n", "", -1))
	return str, nil
}

func ShowConsole(expens []float32) {

	fmt.Println(contenString(expens))
}

func contenString(expens []float32) string {
	builder := strings.Builder{}
	total, max, min, avg := exponsesDetail(expens)

	fmt.Println("")

	for i, expen := range expens {
		builder.WriteString(fmt.Sprintf("Exponse: %6.2f\n", expen))
		if i == len(expens)-1 {
			builder.WriteString("")
			builder.WriteString("=========================\n")
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", total))
			builder.WriteString(fmt.Sprintf("Maximo: %6.2f\n", max))
			builder.WriteString(fmt.Sprintf("Minimo: %6.2f\n", min))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", avg))
			builder.WriteString("=========================\n")


		}

	}

	return builder.String()
}

func exponsesDetail(expens []float32) (total, max, min, avg float32) {

	if len(expens) == 0 {

		return
	}
	total = exponses.Sum(expens...)
	max = exponses.Max(expens...)
	min = exponses.Min(expens...)
	avg = exponses.Average(expens...)
	return

}

func Export(fileName string, list []float32) error {

	f, err := os.Create(fileName)
	if err != nil {

		return err
	}

	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(contenString(list))

	if err != nil {
		return err
	}

	return w.Flush()

}
