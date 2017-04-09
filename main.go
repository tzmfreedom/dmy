package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
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
	Language   string
	Delimiter  string
	Enclosure  string
	StartTime  int64
	Linebreak  string
	Header     string
}

type record struct {
	Index int
}

var (
	Version  string
	Revision string
)

var Out = os.Stdout

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
			Name:        "dateformat, F",
			Destination: &cfg.DateFormat,
			Value:       time.RFC3339,
			EnvVar:      "DATEFORMAT",
		},
		cli.StringFlag{
			Name:        "language, L",
			Destination: &cfg.Language,
			Value:       "en",
		},
		cli.StringFlag{
			Name:        "delimiter, D",
			Destination: &cfg.Delimiter,
			Value:       "\t",
		},
		cli.StringFlag{
			Name:        "enclosure, E",
			Destination: &cfg.Enclosure,
			Value:       "",
		},
		cli.Int64Flag{
			Name:        "starttime",
			Destination: &cfg.StartTime,
			Value:       time.Now().Unix(),
		},
		cli.StringFlag{
			Name:        "linebreak",
			Destination: &cfg.Linebreak,
			Value:       "\n",
		},
		cli.StringFlag{
			Name:        "header",
			Destination: &cfg.Header,
			Value:       "",
		},
	}

	app.Action = func(c *cli.Context) error {
		if cfg.Number <= 0 {
			return errors.New("You must specify --number option")
		}
		if len(c.Args()) == 0 {
			return errors.New("You mus specify template strings to output")
		}
		return outputDummyData(c.Args(), cfg)
	}
	app.Run(os.Args)
}

func outputDummyData(columns []string, cfg *config) error {
	fake.SetLang(cfg.Language)

	d := getDecorator(cfg)
	tpls := getTemplates(cfg, columns)

	if cfg.Header != "" {
		headers := strings.Split(cfg.Header, ",")
		Out.Write([]byte(strings.Join(headers, cfg.Delimiter) + cfg.Linebreak))
	}
	for i := 0; i < cfg.Number; i++ {
		cols := []string{}
		for j := 0; j < len(columns); j++ {
			buf := bytes.NewBuffer([]byte{})
			tpls[j].Execute(buf, record{Index: i})
			cols = append(cols, d.Decorate(buf.String()))
		}
		Out.Write([]byte(strings.Join(cols, cfg.Delimiter) + cfg.Linebreak))
	}
	return nil
}

func getTemplates(cfg *config, columns []string) []*template.Template {
	funcMap := getFuncMap(cfg)
	tpls := []*template.Template{}
	for i := 0; i < len(columns); i++ {
		tpls = append(tpls, template.Must(template.New("colmn_tpl_"+fmt.Sprint(i)).Funcs(funcMap).Parse(columns[i])))
	}
	return tpls
}

func getDecorator(cfg *config) decorator {
	if cfg.Enclosure != "" {
		return &encloseDecorator{enclosure: cfg.Enclosure}
	} else {
		return &nullDecorator{}
	}
}

func getFuncMap(cfg *config) template.FuncMap {
	rand.Seed(time.Now().UnixNano())

	now := time.Unix(cfg.StartTime, 0)
	return template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
		"div": func(a, b int) int { return a / b },
		"mod": func(a, b int) int { return a % b },
		"choice": func(choices ...string) string {
			i := rand.Intn(len(choices))
			return choices[i]
		},
		"date":     func(s int) string { return now.Add(time.Duration(s) * time.Second).Format(cfg.DateFormat) },
		"date_m":   func(m int) string { return now.Add(time.Duration(m) * time.Minute).Format(cfg.DateFormat) },
		"date_h":   func(h int) string { return now.Add(time.Duration(h) * time.Hour).Format(cfg.DateFormat) },
		"date_add": func(y int, m int, d int) string { return now.AddDate(y, m, d).Format(cfg.DateFormat) },
		"fake": func(t string) string {
			if f, ok := fakeMap[strings.ToLower(t)]; ok {
				return f()
			}
			return ""
		},
	}
}

type decorator interface {
	Decorate(string) string
}

type encloseDecorator struct {
	enclosure string
}

func (d *encloseDecorator) Decorate(src string) string {
	return d.enclosure + strings.Replace(src, d.enclosure, "\\"+d.enclosure, -1) + d.enclosure
}

type nullDecorator struct{}

func (d *nullDecorator) Decorate(src string) string {
	return src
}
