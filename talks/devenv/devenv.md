---
title: Development Environment
author: Bibek, Terence, Chai, Kevin
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
- should have been installed beforehand
- will only be using terminal, package manager, and git today
  - please install this now if you have not
  - if you need help installing Docker or Nodejs please seek help afterward
:::

# Development Tools

# Command line

## Command line

- Make sure that you have access to a command line.
- The command line is *the* way to interact with your computer as a developer.
  The suite of tools available to you on the command line is unparalleled.
- This is how you can interact with your environment over the network
  - ssh into production servers to debug issues
- bash is one shell that is extremely common
- zsh is another with more powerful features and a plugin system

## How to get help

- `man` pages
- TODO: include a line of history
- `--help` flag exists for most tools informing how they should be used

## Git

- The defacto version control system
- TODO: Link to linus torvalds talk
- do not use SVN or CVS

## SSH

- access remote machines
- necessary to setup production environments
- useful as a general point to point communication protocol
  - proxy local port from/to a remote machine
- TODO: include screenshot of an ssh session

## Learn how to use a text editor

- vim
  - if you need a gui: vscode/atom
- normally all you need in web development
- plugins make the development experience really powerful
  - code completion, windowing, etc.
  - plus everything can work when ssh'd

## Package manager

- comes installed by default on Linux
- brew on mac
- use the Linux subsystem for windows
- allows you to install/uninstall/update the programs on your computer

## Node

- Chrome's V8 Javascript engine ported onto the server
- Comes with `npm`, node package manager, to install and manage dependencies
  for node projects

## Browser Developer Tools

- Chrome / Firefox Dev Tools
- `Inspector`: interactive view of the DOM
- `Network`: shows all network requests, into and out of the browser
  - helpful for debugging caching
  - debugging loading of certain resources
- `Storage/Application`: shows cookies, local storage, service workers

## Prettier

- js/css code formatter
- do not dispute about how code should be formatted
- deterministic formatter

> Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
> - Rob Pike

## Docker

- equivalent development and production environments
- allows for multiple development environments, e.g. multiple databases
- let me know if you want to learn more

# Any Questions?
