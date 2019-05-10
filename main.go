package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	//vendor packages
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const AuthFile string = "./token.json"
const LogJson string = "./log/log.json"
const Update string = "./log/update.txt"

type ID struct {
	Consumer_key        string `json:"ConsumerKey"`
	Consumer_secret     string `json:"ConsumerSecret"`
	Access_token        string `json:"AccessToken"`
	Access_token_secret string `json:"AccessTokenSecret"`
}

type Log struct {
	LogId     string `json:"id"`
	Text      string `json:"text"`
	TimeStamp string `json:"time-stamp"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var (
		id          ID
		l           []Log
		updateParam twitter.StatusUpdateParams
	)

	//read update log file
	file, err := ioutil.ReadFile(Update)
	check(err)
	msg := string(file)

	//get twitter access token from json file
	idBytes, err := ioutil.ReadFile(AuthFile)
	check(err)
	json.Unmarshal(idBytes, &id)

	//set update parameters
	logBytes, err := ioutil.ReadFile(LogJson)
	check(err)

	json.Unmarshal(logBytes, &l)
	logSize := len(l)
	if logSize > 0 {
		ToReply, err := strconv.Atoi(l[logSize-1].LogId)
		check(err)

		updateParam.InReplyToStatusID = int64(ToReply)
	}

	client, err := SetClient(id)
	check(err)
	tweet, _, err := client.Statuses.Update(msg, func() *twitter.StatusUpdateParams {
		if updateParam.InReplyToStatusID == 0 {
			return nil
		}
		return &updateParam
	}())
	check(err)
	newTweet := Log{
		LogId:     tweet.IDStr,
		Text:      tweet.Text,
		TimeStamp: tweet.CreatedAt,
	}
	l = append(l, newTweet)

	log, err := json.Marshal(l)
	check(err)

	err = ioutil.WriteFile(LogJson, log, 0644)
	check(err)
}

func SetClient(t ID) (*twitter.Client, error) {
	config := oauth1.NewConfig(t.Consumer_key, t.Consumer_secret)
	_token := oauth1.NewToken(t.Access_token, t.Access_token_secret)

	httpClient := config.Client(oauth1.NoContext, _token)
	client := twitter.NewClient(httpClient)

	return client, nil
}
