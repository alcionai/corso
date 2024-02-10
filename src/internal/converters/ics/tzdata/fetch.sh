#!/bin/sh

set -eo pipefail

if ! echo "$PWD" | grep -q '/tzdata$'; then
    echo "Please run this script from the tzdata dir"
    exit 1
fi

# TODO: Generate from https://www.iana.org/time-zones
if [ ! -d /tmp/corso-tzdata ]; then
    git clone --depth 1 https://github.com/add2cal/timezones-ical-library.git /tmp/corso-tzdata
else
    cd /tmp/corso-tzdata
    git pull
    cd -
fi

# Generate a huge go file with all the timezones
echo "package tzdata" >data.go
echo "" >>data.go

echo "var TZData = map[string]string{" >>data.go

find /tmp/corso-tzdata/ -name '*.ics' | while read -r f; do
    tz=$(echo "$f" | sed 's|/tmp/corso-tzdata/api/||;s|\.ics$||')
    echo "Processing $tz"
    printf "\t\"%s\": \`" "$tz" >>data.go
    cat "$f" | grep -Ev "(BEGIN:|END:|TZID:)" |
        sed 's|`|\\`|g;s|\r||;s|TZID:/timezones-ical-library/|TZID:|' |
        perl -pe 'chomp if eof' >>data.go
    echo "\`," >>data.go
done

echo "}" >>data.go
