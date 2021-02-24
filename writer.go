package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func write(fileName string, deliveries []Delivery) {
	f, e := os.Create(fileName)
	check(e)
	writer := bufio.NewWriter(f)

	_, e = fmt.Fprintln(writer, len(deliveries))
	check(e)

	for _, d := range deliveries {
		_, e = fmt.Fprintf(writer, "%d%s\n", d.teamSize, toStr(d.pizzaIds))
	}

	check(writer.Flush())
	check(f.Close())
}

func toStr(ids []int) string {
	str := ""

	for _, id := range ids {
		str = str + " " + strconv.Itoa(id)
	}

	return str
}
