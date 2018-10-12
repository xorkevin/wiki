---
title: Development Environment
author: Bibek, Terrence, Chai, Kevin
date: 2018-10-12
---

# Setting up your dev env

::: notes

- Opening remarks
- Hi! First of all, welcome to the very first DevX Bootcamp.
- An introduction to who we are:
  - Bibek, Terrence, Chai, Kevin
- Given that this is our first time hosting this series, we want feedback
  - want to make this program more useful
  - this could be topics that you want covered, teaching style, anything
- I do want to emphasize that the point of this Bootcamp series is not to teach
  specific technologies
  - only serve to give context to the larger concepts of how to build large
    scalable applications
  - The motivation for this is the DevX projects that you will be joining all
    use different technologies, and later in the industry as well. So remember
    to focus on the concepts.

:::

## Installed dependencies

1. Get access to a command line / terminal emulator
2. Get a package manager (if you can)
3. Install Git
4. Install Docker + Docker Compose
5. Install Nodejs + npm

::: notes

- if you need help with installation, seek help either on the slack, or ask one
  of us in person

:::

# Development Tools

# Command line

## Command line

- Make sure that you have access to a command line.
- The command line is *the* way to interact with your computer as a developer.
  The suite of tools available to you on the command line is unparalleled.
  - CLI tooling is built by developers for developers
- This is how you can interact with your environment over the network
  - `ssh` into production servers to debug issues
- `bash` is one shell that is extremely common
- `zsh` is another with more powerful features and a plugin system

## Navigating your computer

- everything on your computer is organized into a file system
  - you can navigate this file system from the terminal
- `pwd`: prints out your current location
- `ls [directory]`: lists the contents of your current directory
- `cd [directory]`: change directory

::: notes

- this echos a lot of CS35L, but is important to learn, and familiarize
  yourself with
- be careful because these commands are powerful, so understand what the
  command you are executing does before you do it
  - this is no cause to be afraid to learn, however
  - ask questions, be curious

:::

## How to get help

- DevX!
  - we are your best resource
- `man` pages
  - Originally published as a physical book, *Unix Programmer's Manual*, in
    1971, by Dennis Ritchie and Ken Thompson
  - Second edition, first interactive version, 1983
  - e.g. `man sort`
    - Written by Mike Haertel and Paul Eggert
  - e.g. `man ascii`
  - awesome blog post: (The Lineage of Man)[
    https://twobithistory.org/2017/09/28/the-lineage-of-man.html]
- `--help` flag exists for most tools informing how they should be used
  - e.g. `git status --help`

## Git

- The de facto version control system
- (Linus Torvalds on git)[https://www.youtube.com/watch?v=4XpnKHJAok8]
  - please do not use SVN or CVS

## SSH

- access remote machines
- necessary to setup production environments
- also useful as a general point to point communication protocol
  - proxy local port from/to a remote machine

## Learn how to use a text editor

- vim/emacs (but really vim)
  - if you need a gui: vscode/atom, but should still know how to use a command
    line text editor
- normally all you need in web development, IDE's serve little purpose with
  a command line available
- plugins make the development experience really powerful
  - code completion, windowing, etc.
  - plus everything just works when ssh'd into another machine

# Other tools

## Package manager

- comes installed by default on Linux
  - e.g.: `apt`, `yum`, `pacman`
- `brew` on MacOS
- use the Linux subsystem for Windows
- allows you to install/uninstall/update the programs on your computer
  - built by developers for developers

## Node

- Chrome's V8 Javascript engine ported onto the server
- Comes with `npm`, node package manager, to install and manage dependencies
  for node projects
  - just as `apt`, `brew`, etc. manages dependencies for your computer

## Browser Developer Tools

- Chrome / Firefox Dev Tools
- `Inspector`: interactive view of the DOM
- `Network`: shows all network requests, into and out of the browser
  - helpful for debugging caching
  - debugging loading of certain resources
- `Storage/Application`: shows cookies, local storage, service workers
- `Console`: for debugging running Javascript
  - print line debugging is extremely powerful in any programming environment
    since it requires no actual tools or dependencies
- `Performance/Memory`: debug memory leaks, performance problems
  - user experience, especially low power devices is hindered
    - phones, watches, etc.

## Prettier

- `js`/`css` code formatter
- do not dispute about how code should be formatted
- deterministic formatter
  - makes it easier to read other people's code, and helps other people read
    your code

> Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
> - Rob Pike
>
> Prettier's style is no one's favorite, yet prettier is everyone's favorite.
> - Javascript developers

## Docker

- equivalent development and production environments
- allows for multiple development environments, e.g. multiple databases

::: notes

- will be covering the why and how in a later session

:::

# Any Questions?
