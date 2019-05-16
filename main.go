package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cmmarslender/mack"
	icon "github.com/cmmarslender/zoom-status/icons"
	"github.com/getlantern/systray"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"regexp"
	"strings"
	"time"
)

type status struct {
	StatusText string `json:"status_text"`
	StatusEmoji string `json:"status_emoji"`
}

type slackAccount struct {
	Name string `json:name`
	Token string `json:token`
	MeetingStatus *status `json:meetingStatus,omitempty`
	NoMeetingStatus *status `json:noMeetingStatus,omitempty`
}

type slackProfile struct {
	Profile status `json:"profile"`
}

var ignoreMatches []string
var exactMatches []string
var regexStrings []string
var regexMatches []*regexp.Regexp
var defaultMeetingStatus status
var defaultNoMeetingStatus status
var slackAccounts []slackAccount

var inMeeting = false
var menuStatus *systray.MenuItem

func main() {
	systray.Run( onReady, onExit )
}

func onReady() {
	systray.SetTooltip( "Zoom Status" )
	systray.SetIcon( icon.Data )

	menuStatus = systray.AddMenuItem( "Status: Not In Meeting", "Not In Meeting" )
	menuStatus.Disable()

	systray.AddSeparator()

	systray.AddSeparator()
	mQuit := systray.AddMenuItem( "Quit Zoom Status", "Quit Zoom Status" )
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	loadConfig()
	initMatches()
	initStatuses()

	for {
		inMeetingNow := checkForMeeting()

		if true == inMeetingNow {
			if false == inMeeting {
				setInMeeting()
			} else {
				fmt.Println( "Status already set to in meeting" )
			}
		} else {
			if true == inMeeting {
				setNoMeeting()
			} else {
				fmt.Printf( "Status already set to not in meeting \n" )
			}
		}

		time.Sleep( 60 * time.Second )
	}
}

func onExit() {
	setNoMeeting()
}

func loadConfig() {
	fmt.Println( "Loading Config..." )

	usr, err := user.Current() ; if err != nil {
		panic( err )
	}

	jsonFile, err := os.Open( usr.HomeDir + "/slack-status-config.json" ) ; if err != nil {
		panic( err )
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll( jsonFile )

	_ = json.Unmarshal(byteValue, &slackAccounts)
}

func initMatches() {
	ignoreMatches = append( ignoreMatches, "Zoom - Pro Account", "Zoom - Free Account" )
	exactMatches = append( exactMatches, "Zoom" )
	regexStrings = append( regexStrings, "^Zoom Meeting ID.*" )

	for _, regString := range regexStrings {
		expr, err := regexp.Compile( regString ) ; if err != nil {
			panic( err )
		}

		regexMatches = append( regexMatches, expr )
	}
}

func initStatuses() {
	defaultMeetingStatus.StatusText = "In A Meeting"
	defaultMeetingStatus.StatusEmoji = ":zoom:"

	defaultNoMeetingStatus.StatusText = ""
	defaultNoMeetingStatus.StatusEmoji = ""
}

func checkForMeeting() bool {
	fmt.Println( "Checking for active meetings..." )
	result, err := mack.Tell("System Events", "get the title of every window of every process" ) ; if err != nil {
		panic( err )
	}

	apps := strings.Split( result, "," )
	apps = delete_empty( apps )

loop1:
	for _, app := range apps {
		if strings.Index( app, "Zoom" ) == -1 {
			continue
		}

		for _, ignoreString := range ignoreMatches {
			if app == ignoreString {
				continue loop1
			}
		}

		for _, exactMatch := range exactMatches {
			if app == exactMatch {
				return true
			}
		}

		for _, regexMatch := range regexMatches {
			if regexMatch.Match( []byte( app ) ) {
				return true
			}
		}
	}

	return false
}

func setInMeeting() {
	fmt.Println( "Setting status to in meeting" )
	inMeeting = true

	// Set status for all accounts
	for _, account := range slackAccounts {
		fmt.Println( "Setting slack status for " + account.Name )

		var profile slackProfile
		var status status

		if nil != account.MeetingStatus {
			status = *account.MeetingStatus
		} else {
			status = defaultMeetingStatus
		}

		profile.Profile = status

		setSlackProfile( profile, account.Token )
	}

	menuStatus.SetTitle( "Status: In Meeting" )
}

func setNoMeeting() {
	fmt.Printf( "Setting status to not in meeting \n" )
	inMeeting = false

	// Set status for all accounts
	for _, account := range slackAccounts {
		fmt.Println( "Setting slack status for " + account.Name )

		var profile slackProfile
		var status status

		if nil != account.NoMeetingStatus {
			status = *account.NoMeetingStatus
		} else {
			status = defaultNoMeetingStatus
		}

		profile.Profile = status

		setSlackProfile( profile, account.Token )
	}

	menuStatus.SetTitle( "Status: Not In Meeting" )
}

func setSlackProfile( profile slackProfile, token string ) bool {
	statusBytes, err := json.Marshal( profile ) ; if err != nil {
		panic( err )
	}

	req, err := http.NewRequest( "POST", "https://slack.com/api/users.profile.set", bytes.NewBuffer( statusBytes ) )
	if err != nil {
		panic( err )
	}

	// Add proper headers
	req.Header.Add( "Content-Type", "application/json" )
	req.Header.Add( "Authorization", "Bearer " + token )

	resp, err := http.DefaultClient.Do( req )
	if err != nil {
		panic( err )
	}
	defer resp.Body.Close()

	return true
}

/*
Removes empty strings from a slice of strings
 */
func delete_empty (s []string) []string {
	var r []string
	for _, str := range s {
		str = strings.Trim( str, " " )
		if str != "" {
			r = append( r, str )
		}
	}
	return r
}
