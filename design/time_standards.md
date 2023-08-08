# Corso Standard Time Format

Since Corso primarily deals with with historical point-in-time events (created-at, delivered-at, started-on, etc), as opposed to scheduled future events, we can safely represent all time with point-in-time, timezoned values.

The standard string format uses [iso-8601](https://en.wikipedia.org/wiki/ISO_8601) date & time + [rfc-3339](https://datatracker.ietf.org/doc/html/rfc3339) compliant formats, with milliseconds.  All time values should be stored in the UTC timezone.

ex:
* `2022-07-11T20:07:59.00000Z`  
* `2022-07-11T20:07:59Z`

In golang implementation, use the time package format `time.RFC3339Nano`.

## Deviation From the Standard

The above standard helps ensure clean transformations when serializing to and from time.Time structs and strings.  In certain cases time values will need to be displayed in a non-standard format.  These variations are acceptable as long as the application has no intent to consume that format again in a future process.

Examples: the date-time suffix in restoration destination root folders or the human-readable format displayed in CLI outputs.

## Input Leniency

End users are not required to utilize the standard time format when calling inputs.  In practice, all formats recognized in `/internal/common/time.go` can be used as a valid input value.  Official format support is detailed in the public Corso Documentation.

## Maintenance

All supported time formats should appear as consts in  `/internal/common/time.go`.  Usage of `time.Parse()` and `time.Format()` should be kept to that package.