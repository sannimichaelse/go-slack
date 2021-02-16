go-slack : A simple slack message tool in console.
==================

A simple console app that lets you send message to slack specific channel.

Before use this tool, remember you get your Incoming Webhook number in this [page](https://api.slack.com/incoming-webhooks).

Install
---------------
`go get github.com/sannimichaelse/go-slack`


Change the default setting
---------------

Change the configuration file which will generate when your first time launch this application.

```
{
	"webhookUrl" : "",	
	"botName" : "",
	"channel" : "",
	"emoji" : ""
}
```
Detail explaination as follow:

- `webhookUrl`: Get url from https://api.slack.com/incoming-webhooks
- `botName`: The name display of your bot
- `channel`: The channel your bot want to push message
- `emoji`: The emoji of your bot, such `:octocat:`. Refer http://www.emoji-cheat-sheet.com/ for more

Fill all infor base on this [page](https://api.slack.com/incoming-webhooks)

Usage
---------------

```
go-slack -m "test msg"
```
