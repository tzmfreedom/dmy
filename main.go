package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/icrowley/fake"
	"github.com/urfave/cli"
)

type config struct {
	Number     int
	DateFormat string
}

type record struct {
	Index int
}

var (
	Version  string
	Revision string
)

var now = time.Now()

var fakeMap = map[string](func() string){
	"brand": func() string {
		return fake.Brand()
	},
	"character": func() string {
		return fake.Character()
	},
	"characters": func() string {
		return fake.Characters()
	},
	//"charactersn": func() string {
	//        return fake.CharactersN()
	//},
	"city": func() string {
		return fake.City()
	},
	"color": func() string {
		return fake.Color()
	},
	"company": func() string {
		return fake.Company()
	},
	"continent": func() string {
		return fake.Continent()
	},
	"country": func() string {
		return fake.Country()
	},
	//"creditcardnum": func() string {
	//        return fake.CreditCardNum()
	//},
	"creditcardtype": func() string {
		return fake.CreditCardType()
	},
	"currency": func() string {
		return fake.Currency()
	},
	"currencycode": func() string {
		return fake.CurrencyCode()
	},
	"day": func() string {
		return strconv.Itoa(fake.Day())
	},
	"digits": func() string {
		return fake.Digits()
	},
	//"digitsn": func() string {
	//        return fake.DigitsN()
	//},
	"domainname": func() string {
		return fake.DomainName()
	},
	"domainzone": func() string {
		return fake.DomainZone()
	},
	"emailaddress": func() string {
		return fake.EmailAddress()
	},
	"emailbody": func() string {
		return fake.EmailBody()
	},
	"emailsubject": func() string {
		return fake.EmailSubject()
	},
	//"enfallback": func() string {
	//        return fake.EnFallback()
	//},
	//"fs": func() string {
	//        return fake.FS()
	//},
	"femalefirstname": func() string {
		return fake.FemaleFirstName()
	},
	"femalefullname": func() string {
		return fake.FemaleFullName()
	},
	"femalefullnamewithprefix": func() string {
		return fake.FemaleFullNameWithPrefix()
	},
	"femalefullnamewithsuffix": func() string {
		return fake.FemaleFullNameWithSuffix()
	},
	"femalelastname": func() string {
		return fake.FemaleLastName()
	},
	"femalepatronymic": func() string {
		return fake.FemalePatronymic()
	},
	"firstname": func() string {
		return fake.FirstName()
	},
	"fullname": func() string {
		return fake.FullName()
	},
	"fullnamewithprefix": func() string {
		return fake.FullNameWithPrefix()
	},
	"fullnamewithsuffix": func() string {
		return fake.FullNameWithSuffix()
	},
	"gender": func() string {
		return fake.Gender()
	},
	"genderabbrev": func() string {
		return fake.GenderAbbrev()
	},
	//"getlangs": func() string {
	//        return fake.GetLangs()
	//},
	"hexcolor": func() string {
		return fake.HexColor()
	},
	"hexcolorshort": func() string {
		return fake.HexColorShort()
	},
	"ipv4": func() string {
		return fake.IPv4()
	},
	"ipv6": func() string {
		return fake.IPv6()
	},
	"industry": func() string {
		return fake.Industry()
	},
	"jobtitle": func() string {
		return fake.JobTitle()
	},
	"language": func() string {
		return fake.Language()
	},
	"lastname": func() string {
		return fake.LastName()
	},
	"latitudedegreess": func() string {
		return strconv.Itoa(fake.LatitudeDegreess())
	},
	"latitudedirection": func() string {
		return fake.LatitudeDirection()
	},
	"latitudeminutes": func() string {
		return strconv.Itoa(fake.LatitudeMinutes())
	},
	"latitudeseconds": func() string {
		return strconv.Itoa(fake.LatitudeSeconds())
	},
	"latitute": func() string {
		return fmt.Sprint(fake.Latitute())
	},
	"longitude": func() string {
		return fmt.Sprint(fake.Longitude())
	},
	"longitudedegrees": func() string {
		return strconv.Itoa(fake.LongitudeDegrees())
	},
	"longitudedirection": func() string {
		return fake.LongitudeDirection()
	},
	"longitudeminutes": func() string {
		return strconv.Itoa(fake.LongitudeMinutes())
	},
	"longitudeseconds": func() string {
		return strconv.Itoa(fake.LongitudeSeconds())
	},
	"malefirstname": func() string {
		return fake.MaleFirstName()
	},
	"malefullname": func() string {
		return fake.MaleFullName()
	},
	"malefullnamewithprefix": func() string {
		return fake.MaleFullNameWithPrefix()
	},
	"malefullnamewithsuffix": func() string {
		return fake.MaleFullNameWithSuffix()
	},
	"malelastname": func() string {
		return fake.MaleLastName()
	},
	"malepatronymic": func() string {
		return fake.MalePatronymic()
	},
	"model": func() string {
		return fake.Model()
	},
	"month": func() string {
		return fake.Month()
	},
	"monthnum": func() string {
		return strconv.Itoa(fake.MonthNum())
	},
	"monthshort": func() string {
		return fake.MonthShort()
	},
	"paragraph": func() string {
		return fake.Paragraph()
	},
	"paragraphs": func() string {
		return fake.Paragraphs()
	},
	//"paragraphsn": func() string {
	//        return fake.ParagraphsN()
	//},
	//"password": func() string {
	//        return fake.Password()
	//},
	"patronymic": func() string {
		return fake.Patronymic()
	},
	"phone": func() string {
		return fake.Phone()
	},
	"product": func() string {
		return fake.Product()
	},
	"productname": func() string {
		return fake.ProductName()
	},
	//"seed": func() string {
	//        return fake.Seed()
	//},
	"sentence": func() string {
		return fake.Sentence()
	},
	"sentences": func() string {
		return fake.Sentences()
	},
	//"sentencesn": func() string {
	//        return fake.SentencesN()
	//},
	"simplepassword": func() string {
		return fake.SimplePassword()
	},
	"state": func() string {
		return fake.State()
	},
	"stateabbrev": func() string {
		return fake.StateAbbrev()
	},
	"street": func() string {
		return fake.Street()
	},
	"streetaddress": func() string {
		return fake.StreetAddress()
	},
	"title": func() string {
		return fake.Title()
	},
	"topleveldomain": func() string {
		return fake.TopLevelDomain()
	},
	//"useexternaldata": func() string {
	//        return fake.UseExternalData()
	//},
	"useragent": func() string {
		return fake.UserAgent()
	},
	"username": func() string {
		return fake.UserName()
	},
	"weekday": func() string {
		return fake.WeekDay()
	},
	"weekdayshort": func() string {
		return fake.WeekDayShort()
	},
	"weekdaynum": func() string {
		return strconv.Itoa(fake.WeekdayNum())
	},
	"word": func() string {
		return fake.Word()
	},
	"words": func() string {
		return fake.Words()
	},
	//"wordsn": func() string {
	//        return fake.WordsN()
	//},
	//"year": func() string {
	//        return strconv.Itoa(fake.Year())
	//},
	"zip": func() string {
		return fake.Zip()
	},
}

func main() {
	cfg := &config{}
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s revision=%s\n", c.App.Version, Revision)
	}

	app := cli.NewApp()
	app.Name = "dmy"
	app.Usage = "create dummy data"
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "number, N",
			Destination: &cfg.Number,
		},
		cli.StringFlag{
			Name:        "dateformat, D",
			Destination: &cfg.DateFormat,
			Value:       time.RFC3339,
			EnvVar:      "DATEFORMAT",
		},
	}

	app.Action = func(c *cli.Context) error {
		outputDummyData(c.Args(), cfg)
		return nil
	}
	app.Run(os.Args)
}

func outputDummyData(columns []string, cfg *config) {
	var funcMap = template.FuncMap{
		"add":  func(a, b int) int { return a + b },
		"sub":  func(a, b int) int { return a - b },
		"mul":  func(a, b int) int { return a * b },
		"div":  func(a, b int) int { return a / b },
		"mod":  func(a, b int) int { return a % b },
		"date": func(s int) string { return now.Add(time.Duration(s) * time.Second).Format(cfg.DateFormat) },
		"fake": func(t string) string {
			if f, ok := fakeMap[t]; ok {
				return f()
			}
			return ""
		},
	}
	tpl := template.Must(template.New("dummy_data").Funcs(funcMap).Parse(strings.Join(columns, "\t") + "\n"))
	for i := 0; i < cfg.Number; i++ {
		tpl.Execute(os.Stdout, record{Index: i})
	}
}
