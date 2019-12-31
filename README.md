# Easy Date/Time Formats with Duration Spans

[![GoDoc](https://godoc.org/gitlab.com/robmuh/dtime?status.svg)](https://godoc.org/gitlab.com/robmuh/dtime)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/robmuh/dtime)](https://goreportcard.com/report/gitlab.com/robmuh/dtime)
[![Coverage](https://gocover.io/_badge/gitlab.com/robmuh/dtime)](https://gocover.io/gitlab.com/robmuh/dtime)
![License](https://img.shields.io/github/license/robmuh/dtime)

Returns one or two `*time.Time` pointers, one for the first second for given time, and a bounding second of the end of the duration (second after last second). This allows easy checks for give times within that duration.

```
+20m
+1h
-2.5d
today
yesterday
tomorrow
nextweek
nextmonth
lweek
lmon
lastyear+4w
```

See the [test data](testdata/dtime.yaml) for hundreds of examples and the [PEG grammar](grammar.peg) for specifics.

## Motivation

When using a mobile device the only characters available on the default keyboard are alpha-numeric and the comma (,) and period (.). While it is only a minor convenience to shift to the character keyboard why not create a set of formats that worth with the least amount of trouble. Therefore, these formats use the shortest, best format possible to convey the most common references to dates and times. 

This also makes these time formats particularly useful to add to applications with a terse command-line interface.

## TODO

* Add the `dtime` command (with tab completion) to go with the package.

## See Also

### TJ Holowaychuk's `go-naturaldate` Package

TJ's [go-naturaldate](https://github.com/tj/go-naturaldate) package came out while I was developing this one. I noted his use of PEG and reworked the internals of my package to also use it. 

TJ's package is far better for conversational UIs. 

Mine started with emphasis on the least amount of typing possible and no spaces so that queries can easily be added as singular command-line arguments. 

Mine also comes with the `dtime` command for easy integration into shell scripts or while editing files with `vi/m` using "wand" syntax (`!!`,`!}`,`!G`}. 

I also focus mostly on time spans rather than specific dates.

### Andrew Snodgrass' PEG Golang Package

This [PEG package](https://github.com/pointlander/peg) is truly amazing. My days of writing ABNF are likely over.



