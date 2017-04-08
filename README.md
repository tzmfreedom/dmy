# DMY

Command line interface to create dummy data.

## Install

```bash
$ go get github.com/tzmfreedom/dmy
```

## Usage

```
NAME:
   dmy - create dummy data

USAGE:
   dmy [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --number value, -N value      (default: 0)
   --dateformat value, -D value  (default: "2006-01-02T15:04:05Z07:00") [$DATEFORMAT]
   --help, -h                    show help
   --version, -v                 print the version

ex)
$ dmy -N 10 hoge "foo_{{.Index}}" "{{date .Index}}"
```

DMY use golang template.
`.Index` is record's index that starts at 0.

### Function

You can use some function to bind variables.

|function|feature|example|
| ------ | ----- | ----- |
|add|add two values|{{add 1 .Index}}|
|sub|subtract two values|{{add 100 .Index}}|
|mul|multiply two values|{{mul 2 .Index}}|
|div|divide two values|{{div .Index 2}}|
|mod|modify two values|{{mod .Index 10}}|
|date|date to add seconds from now|{{date .Index}}|
|fake|create fake data|fake "fullname"|

You can change datetime format to use `--dateformat` option.

### Fake function

You can use fake function to create dummy data.
You can set dummy data type to fake function's first argument.

Supported data type is following.

* Brand
* Character
* Characters
* City
* Color
* Company
* Continent
* Country
* CreditCardType
* Currency
* CurrencyCode
* Day
* Digits
* DomainName
* DomainZone
* EmailAddress
* EmailBody
* EmailSubject
* FemaleFirstName
* FemaleFullName
* FemaleFullNameWithPrefix
* FemaleFullNameWithSuffix
* FemaleLastName
* FemalePatronymic
* FirstName
* FullName
* FullNameWithPrefix
* FullNameWithSuffix
* Gender
* GenderAbbrev
* HexColor
* HexColorShort
* IPv4
* IPv6
* Industry
* JobTitle
* Language
* LastName
* LatitudeDegreess
* LatitudeDirection
* LatitudeMinutes
* LatitudeSeconds
* Latitute
* Longitude
* LongitudeDegrees
* LongitudeDirection
* LongitudeMinutes
* LongitudeSeconds
* MaleFirstName
* MaleFullName
* MaleFullNameWithPrefix
* MaleFullNameWithSuffix
* MaleLastName
* MalePatronymic
* Model
* Month
* MonthNum
* MonthShort
* Paragraph
* Paragraphs
* Patronymic
* Phone
* Product
* ProductName
* Sentence
* Sentences
* SimplePassword
* State
* StateAbbrev
* Street
* StreetAddress
* Title
* TopLevelDomain
* UserAgent
* UserName
* WeekDay
* WeekDayShort
* WeekdayNum
* Word
* Words
* Zip