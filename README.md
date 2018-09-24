# dierollcmd
A command-line utility for generating AD&amp;D-style die rolls.
## How does it work?
Once you have it installed somewhere in your path, open up your favorite shell...

    $ dierollcmd d20
    15
    $ dierollcmd 3d6
    18
    $ dierollcmd -d 3d6
    11 : 3 5 3
    $ dierollcmd -d 2d4-1
    4 : 1 4

The `-d` flag provides detailed information -- the number that was rolled for each die.

In addition to the existing, well-known format for die rolls, `dierollcmd` provides a couple of useful extensions to make your die-rolling life easier.

You can provide a multiplier number which asks for multiple sets of the same roll. For instance, if you were creating a standard six-attirubute character, you could do it in one go, like this:

    $ dierollcmd -d 3d6x6
    10 : 1 6 3
    10 : 1 4 5
    10 : 2 3 5
    18 : 6 6 6
    7 : 1 1 5
    8 : 1 5 2

Well, one bright spot there, but the future is looking a little uncertain for your brand new Cleric/Thief, so maybe you'd rather use the roll-four-pick-the-best-three strategy. You can do that easily like this:

    $ dierollcmd -d 3,4d6x6
    12 : 2 4 2 6
    15 : 4 5 6 4
    6 : 1 3 1 2
    15 : 1 4 6 5
    11 : 1 4 4 3
    10 : 1 3 4 3

A bit better! Who needs a high Charisma anyway?

## How can I get it?
It's written in Go, so to build it, you need to have Go installed on your computer.
* Clone the repo
* `go build`
* rename `dierollcmd` to `roll` or some other sensible name
* put it somewhere in your path

You're good. Let the dungeon crawling commence!