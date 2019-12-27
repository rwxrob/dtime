package htime

// Code generated by peg htime.peg DO NOT EDIT.

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"time"
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	ruleString
	ruleExpr
	ruleDateTimeOffset
	ruleDateTime
	ruleOffset
	ruleOffDir
	ruleOffYear
	ruleOffWeek
	ruleOffDay
	ruleOffHour
	ruleOffMinute
	ruleOffSecond
	ruleNUM
	ruleWS
	ruleEND
	ruleAction0
	ruleAction1
	ruleAction2
	ruleAction3
	ruleAction4
	ruleAction5
	ruleAction6
	ruleAction7
	ruleAction8
	rulePegText
	ruleAction9
)

var rul3s = [...]string{
	"Unknown",
	"String",
	"Expr",
	"DateTimeOffset",
	"DateTime",
	"Offset",
	"OffDir",
	"OffYear",
	"OffWeek",
	"OffDay",
	"OffHour",
	"OffMinute",
	"OffSecond",
	"NUM",
	"WS",
	"END",
	"Action0",
	"Action1",
	"Action2",
	"Action3",
	"Action4",
	"Action5",
	"Action6",
	"Action7",
	"Action8",
	"PegText",
	"Action9",
}

type token32 struct {
	pegRule
	begin, end uint32
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v", rul3s[t.pegRule], t.begin, t.end)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(w io.Writer, pretty bool, buffer string) {
	var print func(node *node32, depth int)
	print = func(node *node32, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Fprintf(w, " ")
			}
			rule := rul3s[node.pegRule]
			quote := strconv.Quote(string(([]rune(buffer)[node.begin:node.end])))
			if !pretty {
				fmt.Fprintf(w, "%v %v\n", rule, quote)
			} else {
				fmt.Fprintf(w, "\x1B[34m%v\x1B[m %v\n", rule, quote)
			}
			if node.up != nil {
				print(node.up, depth+1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (node *node32) Print(w io.Writer, buffer string) {
	node.print(w, false, buffer)
}

func (node *node32) PrettyPrint(w io.Writer, buffer string) {
	node.print(w, true, buffer)
}

type tokens32 struct {
	tree []token32
}

func (t *tokens32) Trim(length uint32) {
	t.tree = t.tree[:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) AST() *node32 {
	type element struct {
		node *node32
		down *element
	}
	tokens := t.Tokens()
	var stack *element
	for _, token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		return stack.node
	}
	return nil
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	t.AST().Print(os.Stdout, buffer)
}

func (t *tokens32) WriteSyntaxTree(w io.Writer, buffer string) {
	t.AST().Print(w, buffer)
}

func (t *tokens32) PrettyPrintSyntaxTree(buffer string) {
	t.AST().PrettyPrint(os.Stdout, buffer)
}

func (t *tokens32) Add(rule pegRule, begin, end, index uint32) {
	tree, i := t.tree, int(index)
	if i >= len(tree) {
		t.tree = append(tree, token32{pegRule: rule, begin: begin, end: end})
		return
	}
	tree[i] = token32{pegRule: rule, begin: begin, end: end}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}

type parser struct {
	t      time.Time
	offset float64
	offdir float64
	num    float64
	e      error

	Buffer string
	buffer []rune
	rules  [27]func() bool
	parse  func(rule ...int) error
	reset  func()
	Pretty bool
	tokens32
}

func (p *parser) Parse(rule ...int) error {
	return p.parse(rule...)
}

func (p *parser) Reset() {
	p.reset()
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p   *parser
	max token32
}

func (e *parseError) Error() string {
	tokens, err := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "parse error near %v (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		err += fmt.Sprintf(format,
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return err
}

func (p *parser) PrintSyntaxTree() {
	if p.Pretty {
		p.tokens32.PrettyPrintSyntaxTree(p.Buffer)
	} else {
		p.tokens32.PrintSyntaxTree(p.Buffer)
	}
}

func (p *parser) WriteSyntaxTree(w io.Writer) {
	p.tokens32.WriteSyntaxTree(w, p.Buffer)
}

func (p *parser) Execute() {
	buffer, _buffer, text, begin, end := p.Buffer, p.buffer, "", 0, 0
	for _, token := range p.Tokens() {
		switch token.pegRule {

		case rulePegText:
			begin, end = int(token.begin), int(token.end)
			text = string(_buffer[begin:end])

		case ruleAction0:
			p.offset *= p.offdir
		case ruleAction1:
			p.offdir = 1
		case ruleAction2:
			p.offdir = -1
		case ruleAction3:
			p.offset += p.num * YEAR
		case ruleAction4:
			p.offset += p.num * WEEK
		case ruleAction5:
			p.offset += p.num * DAY
		case ruleAction6:
			p.offset += p.num * HOUR
		case ruleAction7:
			p.offset += p.num * MINUTE
		case ruleAction8:
			p.offset += p.num * SECOND
		case ruleAction9:

			f, _ := strconv.ParseFloat(text, 64)
			p.num = f

		}
	}
	_, _, _, _, _ = buffer, _buffer, text, begin, end
}

func Pretty(pretty bool) func(*parser) error {
	return func(p *parser) error {
		p.Pretty = pretty
		return nil
	}
}

func Size(size int) func(*parser) error {
	return func(p *parser) error {
		p.tokens32 = tokens32{tree: make([]token32, 0, size)}
		return nil
	}
}
func (p *parser) Init(options ...func(*parser) error) error {
	var (
		max                  token32
		position, tokenIndex uint32
		buffer               []rune
	)
	for _, option := range options {
		err := option(p)
		if err != nil {
			return err
		}
	}
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != endSymbol {
			p.buffer = append(p.buffer, endSymbol)
		}
		buffer = p.buffer
	}
	p.reset()

	_rules := p.rules
	tree := p.tokens32
	p.parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokens32 = tree
		if matches {
			p.Trim(tokenIndex)
			return nil
		}
		return &parseError{p, max}
	}

	add := func(rule pegRule, begin uint32) {
		tree.Add(rule, begin, position, tokenIndex)
		tokenIndex++
		if begin != position && position > max.end {
			max = token32{rule, begin, position}
		}
	}

	matchDot := func() bool {
		if buffer[position] != endSymbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 String <- <(Expr END)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				if !_rules[ruleExpr]() {
					goto l0
				}
				if !_rules[ruleEND]() {
					goto l0
				}
				add(ruleString, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 Expr <- <(DateTimeOffset / Offset)> */
		func() bool {
			position2, tokenIndex2 := position, tokenIndex
			{
				position3 := position
				{
					position4, tokenIndex4 := position, tokenIndex
					if !_rules[ruleDateTimeOffset]() {
						goto l5
					}
					goto l4
				l5:
					position, tokenIndex = position4, tokenIndex4
					if !_rules[ruleOffset]() {
						goto l2
					}
				}
			l4:
				add(ruleExpr, position3)
			}
			return true
		l2:
			position, tokenIndex = position2, tokenIndex2
			return false
		},
		/* 2 DateTimeOffset <- <(DateTime Offset?)> */
		func() bool {
			position6, tokenIndex6 := position, tokenIndex
			{
				position7 := position
				if !_rules[ruleDateTime]() {
					goto l6
				}
				{
					position8, tokenIndex8 := position, tokenIndex
					if !_rules[ruleOffset]() {
						goto l8
					}
					goto l9
				l8:
					position, tokenIndex = position8, tokenIndex8
				}
			l9:
				add(ruleDateTimeOffset, position7)
			}
			return true
		l6:
			position, tokenIndex = position6, tokenIndex6
			return false
		},
		/* 3 DateTime <- <('d' 'a' 't' 'e' 't' 'i' 'm' 'e')> */
		func() bool {
			position10, tokenIndex10 := position, tokenIndex
			{
				position11 := position
				if buffer[position] != rune('d') {
					goto l10
				}
				position++
				if buffer[position] != rune('a') {
					goto l10
				}
				position++
				if buffer[position] != rune('t') {
					goto l10
				}
				position++
				if buffer[position] != rune('e') {
					goto l10
				}
				position++
				if buffer[position] != rune('t') {
					goto l10
				}
				position++
				if buffer[position] != rune('i') {
					goto l10
				}
				position++
				if buffer[position] != rune('m') {
					goto l10
				}
				position++
				if buffer[position] != rune('e') {
					goto l10
				}
				position++
				add(ruleDateTime, position11)
			}
			return true
		l10:
			position, tokenIndex = position10, tokenIndex10
			return false
		},
		/* 4 Offset <- <(OffDir OffYear? OffWeek? OffDay? OffHour? OffMinute? OffSecond? WS? Action0)> */
		func() bool {
			position12, tokenIndex12 := position, tokenIndex
			{
				position13 := position
				if !_rules[ruleOffDir]() {
					goto l12
				}
				{
					position14, tokenIndex14 := position, tokenIndex
					if !_rules[ruleOffYear]() {
						goto l14
					}
					goto l15
				l14:
					position, tokenIndex = position14, tokenIndex14
				}
			l15:
				{
					position16, tokenIndex16 := position, tokenIndex
					if !_rules[ruleOffWeek]() {
						goto l16
					}
					goto l17
				l16:
					position, tokenIndex = position16, tokenIndex16
				}
			l17:
				{
					position18, tokenIndex18 := position, tokenIndex
					if !_rules[ruleOffDay]() {
						goto l18
					}
					goto l19
				l18:
					position, tokenIndex = position18, tokenIndex18
				}
			l19:
				{
					position20, tokenIndex20 := position, tokenIndex
					if !_rules[ruleOffHour]() {
						goto l20
					}
					goto l21
				l20:
					position, tokenIndex = position20, tokenIndex20
				}
			l21:
				{
					position22, tokenIndex22 := position, tokenIndex
					if !_rules[ruleOffMinute]() {
						goto l22
					}
					goto l23
				l22:
					position, tokenIndex = position22, tokenIndex22
				}
			l23:
				{
					position24, tokenIndex24 := position, tokenIndex
					if !_rules[ruleOffSecond]() {
						goto l24
					}
					goto l25
				l24:
					position, tokenIndex = position24, tokenIndex24
				}
			l25:
				{
					position26, tokenIndex26 := position, tokenIndex
					if !_rules[ruleWS]() {
						goto l26
					}
					goto l27
				l26:
					position, tokenIndex = position26, tokenIndex26
				}
			l27:
				if !_rules[ruleAction0]() {
					goto l12
				}
				add(ruleOffset, position13)
			}
			return true
		l12:
			position, tokenIndex = position12, tokenIndex12
			return false
		},
		/* 5 OffDir <- <(('+' Action1) / ('-' Action2))> */
		func() bool {
			position28, tokenIndex28 := position, tokenIndex
			{
				position29 := position
				{
					position30, tokenIndex30 := position, tokenIndex
					if buffer[position] != rune('+') {
						goto l31
					}
					position++
					if !_rules[ruleAction1]() {
						goto l31
					}
					goto l30
				l31:
					position, tokenIndex = position30, tokenIndex30
					if buffer[position] != rune('-') {
						goto l28
					}
					position++
					if !_rules[ruleAction2]() {
						goto l28
					}
				}
			l30:
				add(ruleOffDir, position29)
			}
			return true
		l28:
			position, tokenIndex = position28, tokenIndex28
			return false
		},
		/* 6 OffYear <- <(NUM 'y' Action3)> */
		func() bool {
			position32, tokenIndex32 := position, tokenIndex
			{
				position33 := position
				if !_rules[ruleNUM]() {
					goto l32
				}
				if buffer[position] != rune('y') {
					goto l32
				}
				position++
				if !_rules[ruleAction3]() {
					goto l32
				}
				add(ruleOffYear, position33)
			}
			return true
		l32:
			position, tokenIndex = position32, tokenIndex32
			return false
		},
		/* 7 OffWeek <- <(NUM 'w' Action4)> */
		func() bool {
			position34, tokenIndex34 := position, tokenIndex
			{
				position35 := position
				if !_rules[ruleNUM]() {
					goto l34
				}
				if buffer[position] != rune('w') {
					goto l34
				}
				position++
				if !_rules[ruleAction4]() {
					goto l34
				}
				add(ruleOffWeek, position35)
			}
			return true
		l34:
			position, tokenIndex = position34, tokenIndex34
			return false
		},
		/* 8 OffDay <- <(NUM 'd' Action5)> */
		func() bool {
			position36, tokenIndex36 := position, tokenIndex
			{
				position37 := position
				if !_rules[ruleNUM]() {
					goto l36
				}
				if buffer[position] != rune('d') {
					goto l36
				}
				position++
				if !_rules[ruleAction5]() {
					goto l36
				}
				add(ruleOffDay, position37)
			}
			return true
		l36:
			position, tokenIndex = position36, tokenIndex36
			return false
		},
		/* 9 OffHour <- <(NUM 'h' Action6)> */
		func() bool {
			position38, tokenIndex38 := position, tokenIndex
			{
				position39 := position
				if !_rules[ruleNUM]() {
					goto l38
				}
				if buffer[position] != rune('h') {
					goto l38
				}
				position++
				if !_rules[ruleAction6]() {
					goto l38
				}
				add(ruleOffHour, position39)
			}
			return true
		l38:
			position, tokenIndex = position38, tokenIndex38
			return false
		},
		/* 10 OffMinute <- <(NUM 'm' Action7)> */
		func() bool {
			position40, tokenIndex40 := position, tokenIndex
			{
				position41 := position
				if !_rules[ruleNUM]() {
					goto l40
				}
				if buffer[position] != rune('m') {
					goto l40
				}
				position++
				if !_rules[ruleAction7]() {
					goto l40
				}
				add(ruleOffMinute, position41)
			}
			return true
		l40:
			position, tokenIndex = position40, tokenIndex40
			return false
		},
		/* 11 OffSecond <- <(NUM 's' Action8)> */
		func() bool {
			position42, tokenIndex42 := position, tokenIndex
			{
				position43 := position
				if !_rules[ruleNUM]() {
					goto l42
				}
				if buffer[position] != rune('s') {
					goto l42
				}
				position++
				if !_rules[ruleAction8]() {
					goto l42
				}
				add(ruleOffSecond, position43)
			}
			return true
		l42:
			position, tokenIndex = position42, tokenIndex42
			return false
		},
		/* 12 NUM <- <(<([0-9] ('.' [0-9]+)?)> Action9)> */
		func() bool {
			position44, tokenIndex44 := position, tokenIndex
			{
				position45 := position
				{
					position46 := position
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l44
					}
					position++
					{
						position47, tokenIndex47 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l47
						}
						position++
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l47
						}
						position++
					l49:
						{
							position50, tokenIndex50 := position, tokenIndex
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l50
							}
							position++
							goto l49
						l50:
							position, tokenIndex = position50, tokenIndex50
						}
						goto l48
					l47:
						position, tokenIndex = position47, tokenIndex47
					}
				l48:
					add(rulePegText, position46)
				}
				if !_rules[ruleAction9]() {
					goto l44
				}
				add(ruleNUM, position45)
			}
			return true
		l44:
			position, tokenIndex = position44, tokenIndex44
			return false
		},
		/* 13 WS <- <' '> */
		func() bool {
			position51, tokenIndex51 := position, tokenIndex
			{
				position52 := position
				if buffer[position] != rune(' ') {
					goto l51
				}
				position++
				add(ruleWS, position52)
			}
			return true
		l51:
			position, tokenIndex = position51, tokenIndex51
			return false
		},
		/* 14 END <- <!.> */
		func() bool {
			position53, tokenIndex53 := position, tokenIndex
			{
				position54 := position
				{
					position55, tokenIndex55 := position, tokenIndex
					if !matchDot() {
						goto l55
					}
					goto l53
				l55:
					position, tokenIndex = position55, tokenIndex55
				}
				add(ruleEND, position54)
			}
			return true
		l53:
			position, tokenIndex = position53, tokenIndex53
			return false
		},
		/* 16 Action0 <- <{ p.offset *= p.offdir }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
		/* 17 Action1 <- <{p.offdir=1}> */
		func() bool {
			{
				add(ruleAction1, position)
			}
			return true
		},
		/* 18 Action2 <- <{p.offdir=-1}> */
		func() bool {
			{
				add(ruleAction2, position)
			}
			return true
		},
		/* 19 Action3 <- <{p.offset += p.num*YEAR}> */
		func() bool {
			{
				add(ruleAction3, position)
			}
			return true
		},
		/* 20 Action4 <- <{p.offset += p.num*WEEK}> */
		func() bool {
			{
				add(ruleAction4, position)
			}
			return true
		},
		/* 21 Action5 <- <{p.offset += p.num*DAY}> */
		func() bool {
			{
				add(ruleAction5, position)
			}
			return true
		},
		/* 22 Action6 <- <{p.offset += p.num*HOUR}> */
		func() bool {
			{
				add(ruleAction6, position)
			}
			return true
		},
		/* 23 Action7 <- <{p.offset += p.num*MINUTE}> */
		func() bool {
			{
				add(ruleAction7, position)
			}
			return true
		},
		/* 24 Action8 <- <{p.offset += p.num*SECOND}> */
		func() bool {
			{
				add(ruleAction8, position)
			}
			return true
		},
		nil,
		/* 26 Action9 <- <{
		   f, _ := strconv.ParseFloat(text,64);
		   p.num = f
		}> */
		func() bool {
			{
				add(ruleAction9, position)
			}
			return true
		},
	}
	p.rules = _rules
	return nil
}
