# Human-Friendly Date and Time Formats for Go (golang)

When using a mobile device the only characters available on the default keyboard are alpha-numeric and the comma (,) and period (.). While it is only a minor convenience to shift to the character keyboard why not create a set of formats that worth with the least amount of trouble. Therefore, these formats use the shortest, best format possible to convey the most common references to dates and times.

See the [ABNF specification](timefmt.abnf) for the specifics.

## Priority

Part of the format inference is a priority based on which formats are most likely to be used. For example, it is very rare to provide the specific minute when entering information into a mobile device or on the command-line quickly. Usually the hour will suffice. Therefore those formats with a higher likelihood of being used are resolved before others. This is why `mon3p` resolves very high on the priority list and `mon3:30p` rather low, even after the years and months. This resolution happens in nanoseconds and usually involves a human entering them by hand so these differences in priority are beyond negligible.
