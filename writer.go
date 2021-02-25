package main

import (
	"bufio"
	"fmt"
	"os"
)

func write(fileName string, schedules []Schedule) {
	f, err := os.Create(fileName)
	check(err)
	writer := bufio.NewWriter(f)

	_, err = fmt.Fprintln(writer, len(schedules))
	check(err)

	for _, s := range schedules {
		_, err = fmt.Fprintln(writer, s.intersectionId)
		check(err)

		_, err = fmt.Fprintln(writer, len(s.lights))
		check(err)

		for _, l := range s.lights {
			_, err = fmt.Fprintf(writer, "%s %d\n", l.streetName, l.time)
			check(err)
		}
	}

	err = writer.Flush()
	check(err)
	err = f.Close()
	check(err)
}
