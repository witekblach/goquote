# GoQuote <sub><sup><sub><sup>Don't forget what you have learned

![goquote_dragon.png](/assets/goquote_dragon.png)

## Motivation

`Aphorism` is a concise, terse, laconic, or memorable expression of a general truth or principle. Also. In Go land there
are well known [Go Proverbs](https://go-proverbs.github.io/) that should lead anybody in their Go journey.

GoQuote is here to serve you your knowledge. It works by creating a `db.json` file in your `$HOME` directory, storing a
JSON representation in there of your Quotes. Feel free to edit them, as long as its valid JSON.

### Install

```bash
go install github.com/witekblach/goquote
```

### Use

```bash
# returns a random quote to stdout
goquote

# add quotes, or in $HOME/db.json
goquote add "Don't communicate by sharing memory, share memory by communicating."

# prints all quotes in db
goquote list

# search for "app" substring in all quotes
goquote list -s app
```

## My setup

Personally I added Go proverbs and chapter names of positive patterns from the book Sooner Safer Happier. My list:

* go run main.go add "Don't communicate by sharing memory, share memory by communicating.";
* go run main.go add "The bigger the interface, the weaker the abstraction.";
* go run main.go add "A little copying is better than a little dependency.";
* go run main.go add "Clear is better than clever. Reflection is never clear.";
* go run main.go add "Documentation is for users.";
* go run main.go add "Start with Why; Empower the How";
* go run main.go add "Scale agility, not Agile, Vertically then sideways";
* go run main.go add "Invite over Inflict";
* go run main.go add "Communicate, communicate, communicate; Three times as much as you think";
* go run main.go add "Measure for learning";
