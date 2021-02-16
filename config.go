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
// limitations under the License

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var configFile string

const RPC_SERVER_VERSION string = "v1.0"

const DefaultConfigFile string = `{
	"webhookUrl": "",
	"botName": "",
	"channel":"",
	"emoji":""
}`

type ServerConfig struct {
	WebhookUrl string //Get url from https://api.slack.com/incoming-webhooks
	BotName    string //The name display of your bot
	Channel    string //The channel your bot want to push message
	Emoji      string //The emoji of your bot, such `:octocat:`. Refer http://www.emoji-cheat-sheet.com/ for more
}

func LoadConfig() ServerConfig {
	ex, error := os.Executable()
	if error != nil {
		log.Panic(error)
	}
	exPath := filepath.Dir(ex)
	configFile = fmt.Sprintf("%v/go-slack.json", exPath)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err := ioutil.WriteFile(configFile, []byte(DefaultConfigFile), 0644)
		if err != nil {
			log.Panic("Cannot create config file")
		} else {
			fmt.Printf("No configuration file, create a default in ", configFile)
		}
	}

	fmt.Println(configFile)
	file, _ := os.Open(configFile)

	defer file.Close()

	decoder := json.NewDecoder(file)
	configObj := ServerConfig{}

	err := decoder.Decode(&configObj)
	if err != nil {
		log.Panic("No config file")
	}

	return configObj
}
