package ex18

import (
	"fmt"
	"strconv"
	"strings"
)

const Path = "./solutions/ex18"

type expr string
type operator expr

func A(lines []string) {
	count := 0

	for _, l := range lines {
		count += expr(l).evaluate()
	}

	fmt.Println(count)
}

func (e expr) isSimple() bool {
	return !strings.HasPrefix(string(e), "(")
}

func (e expr) isOperator() bool {
	return e == "+" || e == "*"
}

func (e expr) evaluate() int {
	val := 0
	op := operator(expr("+"))
	spl := e.split()
	for _, ex := range spl {
		if ex.isOperator() {
			op = operator(ex)
			continue
		}
		if ex.isSimple() {
			val = op.execute(val, ex.parseInt())
		} else {
			s := string(ex)
			simplerEx := expr(strings.TrimPrefix(strings.TrimSuffix(s, ")"), "("))
			val = op.execute(val, simplerEx.evaluate())
		}
	}
	return val
}

func (e expr) split() []expr {
	s := string(e)
	split := strings.Split(s, " ")

	expressions := []expr{}

	i := 0

	for i < len(split) {
		ss := split[i]
		if strings.HasPrefix(ss, "(") {
			c := 0
			for _, cc := range ss {
				if cc != '(' {
					break
				}
				c++
			}
			sss := []string{ss}
			i++
			for c > 0 {
				x := split[i]
				sss = append(sss, x)
				if strings.HasPrefix(x, "(") {
					for _, cc := range x {
						if cc != '(' {
							break
						}
						c++
					}
				}
				if strings.HasSuffix(x, ")") {
					for _, cc := range x {
						if cc == ')' {
							c--
						}
					}
				}
				i++
			}
			i--
			expressions = append(expressions, expr(strings.Join(sss, " ")))
		} else {
			expressions = append(expressions, expr(ss))
		}
		i++
	}

	return expressions
}

func (op operator) execute(x, y int) int {
	if op == "+" {
		return x + y
	}
	if op == "*" {
		return x * y
	}
	fmt.Println("Warning, unknown operator ", op)
	return 0
}

func (e expr) parseInt() int {
	v, err := strconv.ParseInt(string(e), 10, 64)
	if err != nil {
		fmt.Errorf("Error parsing int from %v: %v", e, err)
	}
	return int(v)
}

func B(lines []string) {
	count := 0

	for _, l := range lines {
		count += expr(l).evaluate2()
	}

	fmt.Println(count)
}

func (e expr) evaluate2() int {
	val := 0
	op := operator(expr("+"))
	e = e.plusize()
	spl := e.split()
	for _, ex := range spl {
		if ex.isOperator() {
			op = operator(ex)
			continue
		}
		if ex.isSimple() {
			val = op.execute(val, ex.parseInt())
		} else {
			s := string(ex)
			simplerEx := expr(strings.TrimPrefix(strings.TrimSuffix(s, ")"), "("))
			val = op.execute(val, simplerEx.evaluate2())
		}
	}
	return val
}

func (e expr) plusize() expr {
	split := e.split()
	l := 0
	p := 1
	i := 1

	for i < len(split) {
		ex := split[i]
		if ex == expr("*") {
			if l != p-1 {
				split[l] = expr("(" + string(split[l]))
				split[p-1] = expr(string(split[p-1]) + ")")
			}
			l = i + 1
			p = i + 1
			i++
			continue
		}
		p++
		i++
	}
	if l != p-1 && l > 0 {
		split[l] = expr("(" + string(split[l]))
		split[p-1] = expr(string(split[p-1]) + ")")
	}

	s := []string{}

	for _, ex := range split {
		s = append(s, string(ex))
	}

	return expr(strings.Join(s, " "))
}
