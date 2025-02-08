# YouRicktube

Dummy telegram bot that provides link to ricktube if someone provides youtube link.  
Useful in Russia where youtube is "not blocked", but unavailable because "google servers are old, man".  
Ricktube downloads youtube video to own CDN and then shows it. 

## Requirements
- telegram token. You have to discuss it with [Botfather](https://t.me/botfather).  

## Launch

execution files are inside `cmd` directory.  

Don't forget to set environment variable  
`TG_BOT_TOKEN` (hope you got it from Botfather already).  

```bash
go run cmd/bot/main.go
```

## Usage

It sniffs messages for youtube link. If there is one, it posts ricktube link.
