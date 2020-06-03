# SHELL NOTIFY

Written in go. Easily send notification cross platform from shell command ( bash, cmd etc...)

It is very handy when you just need to show notification from your bash script or windows shel script

## Install

- golang user:

    `go get github.com/bayucandra/shell-notify` or `go install github.com/bayucandra/shell-notify/cmd/shell-notify`
    
- non golang user:
    Sorry, Not yet build the binary :)

## Usage

to get help: ``shell-notify -help``


Send message with title: ``shell-notify -title message title -message your message body``

You can omit whether ``-title`` or ``-message`` but not both.