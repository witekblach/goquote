# GoQuote <sub><sup><sub><sup>Don't forget what you have learned

![goquote_vader.png](/assets/goquote_vader.png)

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

# add quotes
goquote add "Don't communicate by sharing memory, share memory by communicating."

# prints all quotes in db
goquote list

# search for "communicate" substring in all quotes
goquote list -s communicate

# search for "app" substring in all quotes
goquote remove app
```

FYI: we'll be storing the db of quotes in your `{homedir}/db.json`. Let's see what is inside:

![goquote_dragon.png](/assets/cat_db.png)

## My setup

Personally I added Go proverbs and chapter names of positive patterns from the book Sooner Safer Happier. My list:

```bash
goquote add "Don't communicate by sharing memory, share memory by communicating.";
goquote add "The bigger the interface, the weaker the abstraction.";
goquote add "A little copying is better than a little dependency.";
goquote add "Clear is better than clever. Reflection is never clear.";
goquote add "Documentation is for users.";
goquote add "Start with Why; Empower the How";
goquote add "Scale agility, not Agile, Vertically then sideways";
goquote add "Invite over Inflict";
goquote add "Communicate, communicate, communicate; Three times as much as you think";
goquote add "Measure for learning";
```

## Enjoy!

```
cowsay -l
apt bud-frogs bunny calvin cheese cock cower daemon default dragon dragon-and-cow duck elephant elephant-in-snake eyes 
flaming-sheep fox ghostbusters gnu hellokitty kangaroo kiss koala kosh luke-koala mech-and-cow milk moofasa moose pony 
pony-smaller ren sheep skeleton snowman stegosaurus stimpy suse three-eyes turkey turtle tux unipony unipony-smaller 
vader vader-koala www
```

![goquote_turkey.png](/assets/goquote_turkey.png)

![goquote_dragon.png](/assets/goquote_dragon.png)

Also, visit https://github.com/chubin/wttr.in

```bash
curl wttr.in
# output
Weather report: Warsaw

     \  /       Partly cloudy
   _ /"".-.     +8(6) °C
     \_(   ).   ↓ 17 km/h
     /(___(__)  10 km
                0.1 mm

```

### todo land

This isn't finished, I'm just finished with it (for now). I was thinking this when i left it:

1) storage has to return errors, this should give option to configure goquote on first run (no storage)
2) 2-3 e2e tests would save me time from manual testing
3) views don't do anything, maybe we can make the domain bigger and introduce priority for lower view numbers, maybe
   better way to record views
