package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"web-aggregator/parser"
)

type Params struct {
	lang     *string
	country  *string
	category *string
	kwords   *string
}

/*
TODO: Need to integrate this struct into json.Unmarshal func
type User struct {
	Status       string                     `json:"status"`
	TotalResults int                        `json:"totalResults"`
	Results      map[string]json.RawMessage `json:"results"`
	NextPage     int                        `json:"nextPage"`
}
*/

func RunCLI() {
	PrintTitle()
	CLI := flag.NewFlagSet("get", flag.ExitOnError)

	P := Params{
		lang:     CLI.String("lang", "", "News language"),
		country:  CLI.String("country", "", "Select country"),
		category: CLI.String("category", "", "News category"),
		kwords:   CLI.String("kwords", "", "Search by keywords")}

	if len(os.Args) < 2 {
		fmt.Printf("Ecpected 'get' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		P.HandleGet(CLI)
	default:
		fmt.Println("could not recognize this command")
	}

}

func (i *Params) HandleGet(getCmd *flag.FlagSet) {
	if err := getCmd.Parse(os.Args[2:]); err != nil {
		fmt.Println(err)
	}

	if *i.category == "" && *i.kwords == "" {
		fmt.Print("-kw is required\nor specify category via -cat")
	}

	body, _ := parser.ParseWithParams(*i.lang, *i.country, *i.category, strings.Fields(*i.kwords))
	fmt.Print(string(body))
	if err := ioutil.WriteFile("parsed_data.json", body, 0644); err != nil {
		fmt.Println(err)
	}

	/*
		TODO: Add clean json output

		var users []User

		jsonFile, err := ioutil.ReadFile("parsed_data.json")
		if err != nil {
			fmt.Println(err)
		}


		json.Unmarshal([]byte(jsonFile), &users)
		fmt.Print(users)
		prettyJson, _ := json.MarshalIndent(&users, "", "\t🐱")
		fmt.Print(users)
		fmt.Print(string(prettyJson))
	*/
}
