# commands
[![Build Status](https://travis-ci.org/limetext/commands.svg?branch=master)](https://travis-ci.org/limetext/commands)
[![Coverage Status](https://img.shields.io/coveralls/limetext/commands.svg?branch=master)](https://coveralls.io/r/limetext/commands?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/limetext/commands)](https://goreportcard.com/report/github.com/limetext/commands)
[![GoDoc](https://godoc.org/github.com/limetext/commands?status.svg)](https://godoc.org/github.com/limetext/commands)

This package contains the commands which Lime will use. A command is anything that peforms some action in the editor(like move lines, save file, open file etc).

Lime is designed with the goal of having a clear frontend and backend separation to allow and hopefully simplify the creation of multiple frontend versions.

The Lime Architecture is simple:
1. `backend`: Contain all the code that defines the text editor. The Windows, Folders, Files etc.
2. `front end`: The UI part which will import the backend. When you click on any menu item, the _command_ for that item will be executed via the _Run_ method which you declare. The menubar, the Editor where we will type text into etc
3. `commands`: Utilities which will execute whenever you interact with the UI. 


##Brief overview of Commands

You first build a type of the command which you are writing. 

    type (
        // JoinLines removes every new line in the
        // selections and the first new line after
        JoinLines struct {
            backend.DefaultCommand
        }
    )

Each command has a `Run` method which is executed when the command is invoked. Please note that the command needs to be invoked by the UI for it to work.

    //Run executes the SwapLineUp command
    func (c *SwapLineUp) Run(v *backend.View, e *backend.Edit) error {
        //Do something here
    }

#Imlementing custom commands

If you are interested in implementing your own command, look into the [Implementing commands wiki page](https://github.com/limetext/lime/wiki/Implementing-commands).


# Other references

* [Lime Command interface Api Documentation](http://godoc.org/github.com/limetext/backend#Command)
* Sublime Text 3 unofficial (and [open source](https://github.com/guillermooo/sublime-undocs/)!) [Command documentation](http://docs.sublimetext.info/en/sublime-text-3/extensibility/commands.html)