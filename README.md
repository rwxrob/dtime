# Easy Date and Time Formats with Duration Spans

[![GoDoc](https://godoc.org/gitlab.com/skilstak/go/htime?status.svg)](https://godoc.org/gitlab.com/skilstak/go/htime)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/skilstak/go/htime)](https://goreportcard.com/report/gitlab.com/skilstak/go/htime)
[![Coverage](https://gocover.io/_badge/gitlab.com/skilstak/go/htime)](https://gocover.io/gitlab.com/skilstak/go/htime)

Work in progress. Public for collaboration and education

See the [PEG](grammar.peg) and [test data](testdata/dtime.yaml) for more examples and specifics.

## Motivation

When using a mobile device the only characters available on the default keyboard are alpha-numeric and the comma (,) and period (.). While it is only a minor convenience to shift to the character keyboard why not create a set of formats that worth with the least amount of trouble. Therefore, these formats use the shortest, best format possible to convey the most common references to dates and times. 

This also makes these time formats particularly useful to add to applications with a terse command-line interface.

## See Also

### TJ Holowaychuk's `go-naturaldate` Package

TJ's [go-naturaldate](https://github.com/tj/go-naturaldate) package came out while I was developing this one. I noted his use of PEG and reworked the internals of my package to also use it. 

TJ's package is far better for conversational UIs. 

Mine started with emphasis on the least amount of typing possible and no spaces so that queries can easily be added as singular command-line arguments. 

Mine also comes with the `dtime` command for easy integration into shell scripts or while editing files with `vi/m` using "wand" syntax (`!!`,`!}`,`!G`}. 

Mine also assumes weekdays refer to that of current week whether it be before or after the current date. So, for example, on Tuesday `mon` refers to the previous day, not the next Monday. (For that use `nextmon` or `nmon`.)

I also focus mostly on time spans rather than specific dates.

### Andrew Snodgrass' PEG Golang Package

This [PEG package](https://github.com/pointlander/peg) is truly amazing. My days of writing ABNF are likely over.



