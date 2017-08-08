![](docs/assets/banner.jpg)
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors)

[![First-timers-only Friendly](https://img.shields.io/badge/first--timers--only-friendly-blue.svg)](http://www.firsttimersonly.com/)
[![Twitter](https://img.shields.io/twitter/follow/theheap_.svg?style=social&label=Follow)](https://twitter.com/intent/follow?screen_name=theheap_)
# Project Description

> Read this to get to know what we're building for this project!

Beacon wants to be a helpful tool that makes it easier to keep up to date with _breaking changes_ in your applications when working on a team. Here's an ideal use case of Beacon:

- Your team is working on an app that tells household cats which furniture they can scratch on a specific day.
- You write a migration for your database; this change is a _breaking change_ because your team won't be able to continue hacking until they too update their databases on their local dev machines.
- Rather than having your team fumble trying to figure out what changes were made and why the cats are now able to scratch all the furniture they like; you can use Beacon to keep your team up to date!
- After making your breaking change, you can run beacon in the command line to write a note that you made a breaking change. Beacon will store this note in the log of breaking changes.
- Whenever a dev finds something a bit troubling with their development environment, they can run Beacon from their command line to see if there were any recent changes.

Note, I intend for this project to be written in [Golang](); there are a lot of [boilerplate](https://github.com/urfave/cli) [tools](https://github.com/mkideal/cli) and [libraries](https://github.com/spf13/cobra) for building CLI tools more easily; but we'll try building this from scratch so we can get to know the language. I, at least, want to improve my understanding of Go at a more basic level. Also, as someone who started out as a web dev, I have less experience with simple stuff like reading / writing files / command line tooling. 

**What problem does this tool address?**

Good question. If something breaks in a team application, it can take a while to find out what; you may have to crawl through some git logs or start yelling until someone tells you to calm down and what you need to do. Now you can avoid that by having a specific changelog that only deals with logging important breaking changes.


# Welcome Notes

> Read this to learn how The Heap works collaboratively on projects. 

Welcome to one of THE HEAP's open source projects! Let's take a moment to identify the environment that we will be collaborating in. This project aims to be as accessible as possible to people who want to contribute to open source code! We've got a few things to go over, so if you're new to Github, or contributing to open source software, take a breather and then read on! You got this. 👌


Before we get into the project itself, let's get to know our collaborative environment:
- Most of the activity regarding the project's status happens _right here_ on github, especially on the [issues](https://github.com/the-heap/beacon/issues) page. Here you can see a [roadmap](https://github.com/the-heap/beacon/issues/1) for our project, pick out issues, and keep an eye on conversations.
- The Heap has a [Twitter account to tweet](https://twitter.com/theheap_) both about projects and programming (and life, the universe...). Follow us to stay in the loop.

The following steps will inform the code you write and help you make good pull requests.

1. Understanding how to [contribute](./docs/CONTRIBUTING.md).
2. Read Project Components below, for getting setup.

# Project Setup

> Read this to learn how to get the  project running on your local computer

Instructions forthcoming... 🤔

# Terminology

> Read this to get to know terminology that might help to clarify some aspects of the project / dispel jargon

```
CLI - (Command line tool). A tool used from your computer's terminal (aka; command line)

Breaking Change - This is a broad term. For the purpose of this project, a breaking change is any changes to your tech / dev environment that will halt other people's ability to develop and hack until they update / upgrade / resolve the change as the original author of the change did.
```

# Other

- Thanks for being you
- Follow The Heap on [twitter](https://twitter.com/theheap_).

## Contributors

Thanks goes to these wonderful people ([emoji key](https://github.com/kentcdodds/all-contributors#emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
| [<img src="https://avatars2.githubusercontent.com/u/4494382?v=4" width="80px;"/><br /><sub>Bronek Szulc</sub>](https://github.com/broneks)<br />[💻](https://github.com/teesloane/Beacon/commits?author=broneks "Code") [👀](#review-broneks "Reviewed Pull Requests") | [<img src="https://avatars3.githubusercontent.com/u/563301?v=4" width="80px;"/><br /><sub>Matthew Mihok</sub>](http://mihok.today)<br />[💬](#question-mihok "Answering Questions") [👀](#review-mihok "Reviewed Pull Requests") |
| :---: | :---: |
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/kentcdodds/all-contributors) specification. Contributions of any kind welcome!