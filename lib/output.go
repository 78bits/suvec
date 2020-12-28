package suvec

import (
	"fmt"
	"strings"
)

func (m *Mat) String() string {
	str := ""

	for i := 0; i < m.rows; i++ {
		row := ""
		for j := 0; j < m.cols; j++ {
			row = fmt.Sprintf("%s %0.4f", row, m.data[i*m.cols+j])
		}
		row = strings.TrimSpace(row)
		str = str + row
		if i < m.rows-1 {
			str = str + "\n"
		}
	}
	return str
}

/** Print - Output on screen **/
func (m *Mat) Print(label string) {
	fmt.Println(label, "=")
	fmt.Println(m.String())
}
