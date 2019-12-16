/*
Package timefmt provides a minimal format for entering date and time
information in a practical way. It's primary use case is when a user needs to
enter such data quickly and regularly from the command line or into mobile and
other devices where human input speed is limited to what can be tapped out on
the screen. The characters used in the formatting are deliberately lowercase
and use only characters that appear on most default keyboards without shifting
or switching to symbolic input.

For a precise specification of the format see the timefmt.abnf time included
with this package source code.

Package timefmt also includes a number of useful convenience functions for rounding to closest weekday, computing the same time on a different day, and such.

Note that pointers to time.Time are used throughout the package which are much easier to work with than zero values for the same.
*/
package timefmt
