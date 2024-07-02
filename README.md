# rotmg-go
A go gameserver emulation of rotmg for v1.0.3.0

I decided to learn go language so i made this as a project to practice with

# Current Features
- XML Parsing
- Multiple Connections
- Message Handling and Sending
- Map Parsing
- Basic Entity System 
- Basic Update Logic (Update, NewTick and Move)

# Current Todo
- A System to Update Stats
- Remove hardcoded logic and make everything dynamic (Load World from name with XMLWorld)

# Possible Future Features
- Better Entity System
- Shooting
- Database
- Implement validation and tracking of Players client state

## Dependencies
This project uses the following third-party packages:

- [etree](https://github.com/beevik/etree) - A lightweight element tree XML parser

## licenses

This project includes third-party software governed by the following licenses:

- `etree` - A lightweight element tree XML parser: [etree License](licenses/etree_LICENSE)
