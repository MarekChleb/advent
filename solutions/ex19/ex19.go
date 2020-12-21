package ex19

import (
	"fmt"
	"regexp"
	"strings"
)

const Path = "./solutions/ex19"

type rule [][]string

func A(lines []string) {
	count := 0

	rules := map[string]rule{}
	queries := false
	var possibleWords map[string]struct{}
	for _, l := range lines {
		if l == "" {
			queries = true
			possibleWords = getPossibleWords(rules)
			continue
		}
		if queries {
			if _, ok := possibleWords[l]; ok {
				count++
			}
		} else {
			key, r := parseLine(l)
			rules[key] = r
		}
	}

	fmt.Println(count)
}

func parseLine(l string) (string, rule) {
	re := regexp.MustCompile(`(\d+): (.*)$`)
	info := re.FindStringSubmatch(l)
	key := info[1]

	constr := info[2]

	cRe := regexp.MustCompile(`\"(.*)\"`)
	if cRe.MatchString(constr) {
		cInfo := cRe.FindStringSubmatch(constr)
		c := cInfo[1]
		return key, rule{[]string{c}}
	}

	constrSplit := strings.Split(constr, " ")
	r := rule{}
	cc := []string{}
	for _, c := range constrSplit {
		if c == "|" {
			r = append(r, cc)
			cc = []string{}
			continue
		}
		cc = append(cc, c)
	}
	r = append(r, cc)

	return key, r
}

func (r rule) isSimple() bool {
	if len(r) == 0 {
		fmt.Errorf("Rule is empty")
		return false
	}

	fr := r[0]
	if len(fr) == 0 {
		fmt.Errorf("Rule is empty")
		return false
	}

	re := regexp.MustCompile(`(([a-zA-Z]+)|\[8\]|\[11\])`)
	return re.MatchString(fr[0])
}

func getPossibleWords(rules map[string]rule) map[string]struct{} {
	possibleWords := map[string]struct{}{}

	pwMap := map[string][]string{}

	for k := range rules {
		dfs(k, &pwMap, rules)
	}

	for _, w := range pwMap["0"] {
		possibleWords[w] = struct{}{}
	}

	return possibleWords
}

func dfs(key string, pwMap *map[string][]string, rules map[string]rule) {
	if _, ok := (*pwMap)[key]; ok {
		return
	}

	r := rules[key]
	if r.isSimple() {
		(*pwMap)[key] = []string{r[0][0]}
		return
	}

	for _, rs := range r {
		for _, kk := range rs {
			dfs(kk, pwMap, rules)
		}
	}

	pw := []string{}

	for _, rs := range r {
		s := combine(rs, *pwMap)
		pw = append(pw, s...)
	}

	(*pwMap)[key] = pw
}

func combine(r []string, pwMap map[string][]string) []string {
	s := []string{""}

	for _, k := range r {
		ss := []string{}
		for _, w := range pwMap[k] {
			for _, tw := range s {
				ss = append(ss, tw+w)
			}
		}
		s = ss
	}
	return s
}

func B(lines []string) {
	count := 0

	rules := map[string]rule{}
	queries := false
	var possiblePrefixes, possibleSuffixes map[string]struct{}
	for _, l := range lines {
		if l == "" {
			queries = true
			possiblePrefixes, possibleSuffixes = getPossibleSubwords(rules)
			continue
		}
		if queries {
			if isPossible(l, possiblePrefixes, possibleSuffixes) {
				count++
			}
		} else {
			key, r := parseLine(l)
			if key == "8" {
				r = rule{[]string{"[8]"}}
			} else if key == "11" {
				r = rule{[]string{"[11]"}}
			}
			rules[key] = r
		}
	}

	fmt.Println(count)
}

func getPossibleSubwords(rules map[string]rule) (map[string]struct{}, map[string]struct{}) {
	pwMap := map[string][]string{}

	for k := range rules {
		dfs(k, &pwMap, rules)
	}

	possiblePrefixes := map[string]struct{}{}

	for _, w := range pwMap["42"] {
		possiblePrefixes[w] = struct{}{}
	}

	possibleSuffixes := map[string]struct{}{}

	for _, w := range pwMap["31"] {
		possibleSuffixes[w] = struct{}{}
	}

	return possiblePrefixes, possibleSuffixes
}

func isPossible(l string, possiblePrefixes, possibleSuffixes map[string]struct{}) bool {
	w := l
	sc := 0
	ppp := 8
	for len(w) >= ppp {
		suff := w[len(w)-ppp:]
		w = w[:len(w)-ppp]
		if _, ok := possibleSuffixes[suff]; ok {
			sc += ppp
		} else {
			break
		}
	}

	w = l
	pc := 0
	for len(w) >= ppp {
		pref := w[:ppp]
		w = w[ppp:]
		if _, ok := possiblePrefixes[pref]; ok {
			pc += ppp
		} else {
			break
		}
	}
	rr := pc >= sc && pc+sc >= len(l) && sc >= ppp
	if rr {
		fmt.Println(l)
	}
	return rr
}
