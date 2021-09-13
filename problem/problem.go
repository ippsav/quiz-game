package problem

type Problem struct {
	Question string
	Answer   string
}

type Problems []Problem

func (p *Problems) ParseLines(lines [][]string) {
	for _, line := range lines {
		*p = append(*p, Problem{Answer: line[1], Question: line[0]})
	}
}
