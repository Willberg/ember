package main

import (
	"bufio"
	"ember/datastruct/tree/sp"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// test: Bill 70,Peter 80,Jack 90,Carl 80,Steven 85,Tom 60
	reader := bufio.NewReader(os.Stdin)
	bs, _, _ := reader.ReadLine()
	imputStr := string(bs)
	names, vals := []string{}, []int{}
	for _, s := range strings.Split(imputStr, ",") {
		a := strings.Split(s, " ")
		names = append(names, a[0])
		v, _ := strconv.Atoi(a[1])
		vals = append(vals, v)
	}
	tr := sp.BuildTree(names, vals)
	sp.PrintTree(tr)
	sp.PrintTree2(tr, 0, 0)
	ts := sp.Output(tr, 4)
	for _, t := range ts {
		fmt.Println(t.Name, t.Val)
	}
}
