// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/op/go-logging"
	"github.com/spf13/cobra"
)

var cfg ServerConfig

type SlackObj struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
	Emoji    string `json:"emoji"`
}

func sendServerWarning(msg string) {
	url := cfg.WebhookUrl
	if url == "" {
		fmt.Println("Your configuration is empty, go to ", configFile, " to modifty it.")
		return
	}

	slackObj := SlackObj{}
	slackObj.Channel = cfg.Channel
	slackObj.Username = cfg.BotName
	slackObj.Text = msg
	slackObj.Emoji = cfg.Emoji

	jsonStr, _ := json.Marshal(slackObj)
	connectServer(url, "POST", jsonStr)
}

func connectServer(url string, requestMode string, jsonStr []byte) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(requestMode, url, bytes.NewBuffer(jsonStr))
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Server log failed. err:", err)
		return
	}

	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)
}

func main() {
	cfg = LoadConfig()

	var format = logging.MustStringFormatter("%{level} %{message}")
	logging.SetFormatter(format)
	logging.SetLevel(logging.INFO, "go-slack")
	var chatMsg string
	rootCmd := &cobra.Command{
		Use:   "go-slack",
		Short: "Post message to your slack channel",
		Run: func(cmd *cobra.Command, args []string) {
			sendServerWarning(chatMsg)
		},
	}
	rootCmd.Flags().StringVarP(&chatMsg, "msg", "m", "some text", "Text content")
	rootCmd.Execute()
}
