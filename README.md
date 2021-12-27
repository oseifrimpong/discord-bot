# Discussion-bot

## Introduction

This bot creates a new private thread for each new discussion. The bot has a life time of one hour.

## Installation

### Technical requirements

- Go version: 1.7

- Disgord package: <https://github.com/andersfylling/disgord>

- Docker / Docker-compose: <https://docs.docker.com/compose/install/>

### Steps for running locally

    1. Clone the repository
    2. make a copy of the env file
    3. Insert your discord bot token in the env file
    3. Run `make start` to start the bot
    4. Check the logs for connection status. If the bot is not connected, check the token and make sure it is correct.
    5. Run `make stop` to stop the bot


## How to use bot

Once bot is installed and running in Discord, you can use it in any channel of the Discord server.

- Use command `!help` to see the list of commands.
- To start a new private thread, type `!chat @username`.
- To leave the thread, type `!leave`.

## Improvements

- Don't open multiple threads with the same members. Either close the old thread or redirect them to the old one. \n
Or simply notifying them to close existing thread before opening a new one.

- Notify the members when the thread is about to expire (Does this feature exist in Discord? Investigation needed.).
