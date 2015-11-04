package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"menteslibres.net/gosexy/rest"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	version string = "1.0"
)

func getEnvVar(varName string) (result string) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == varName {
			return pair[1]
		}
	}
	return ""
}

func openUrlByBrowser(url string) (result int) {
	result = 0
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	default:
		fmt.Println("Your PC is not supported.")
	}

	return result
}

func showVersion() {
	fmt.Println("mymine version", version)
}

type Options struct {
	Open    []int  `short:"o" long:"open"    description:"Open specified ticket on a web browser"`
	Version []bool `short:"v" long:"version" description:"Show version"`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)

	parser.Name = "mymine"
	parser.Usage = "[OPTIONS]"

	_, err := parser.Parse()
	if err != nil {
		os.Exit(0)
	}

	if opts.Version != nil {
		showVersion()
		os.Exit(0)
	}

	redmineUrl := getEnvVar("REDMINE_URL")
	if redmineUrl == "" {
		fmt.Println("REDMINE_URL is not specified.")
		os.Exit(0)
	}

	if opts.Open != nil {
		url := redmineUrl + "issues/" + strconv.Itoa(opts.Open[0])
		openUrlByBrowser(url)
		os.Exit(0)
	}

	redmineApiKey := getEnvVar("REDMINE_API_KEY")
	if redmineApiKey == "" {
		fmt.Println("REDMINE_API_KEY is not specified.")
		os.Exit(0)
	}

	project_id := 946 // TODO: should be an option
	request := redmineUrl + "issues.json?key=" + redmineApiKey + "&limit=100&project_id=" + strconv.Itoa(project_id) + "&status_id=open"
	fmt.Println("request =", request)
	fmt.Println("fetching information...")
	var buf map[string]interface{}
	rest.Get(&buf, request, nil)

	issues := buf["issues"].([]interface{})
	for _, v := range issues {
		issue := v.(map[string]interface{})

		var assigned_to string
		tracker := issue["tracker"].(map[string]interface{})
		if issue["assigned_to"] != nil {
			assigned_to_map := issue["assigned_to"].(map[string]interface{})
			assigned_to = assigned_to_map["name"].(string)
		} else {
			assigned_to = "not_assigned"
		}
		id := int(issue["id"].(float64))
		status := issue["status"].(map[string]interface{})
		fmt.Printf("%s #%d %s %s %s\n", tracker["name"], id, status["name"], issue["subject"], assigned_to)
	}
}
