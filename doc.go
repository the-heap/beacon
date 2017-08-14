/*
Beacon is a command line tool that helps developers stay in sync with breaking changes among their team.
Beacon was made as a THE HEAP project, for the month of August 2017.

Project Overview

This document details how a user interacts with Beacon as well as a developer map for navigating the project.
If you are confused or lost, consider viewing the Terms at the bottom of this document for possible clarification.

Usage

A user interacts with Beacon by either Reading or Writing from the Beacon log.
At the base level, Beacon allows a series of sub commands that you can pass into your usage of the application.
Let's take a brief look of the core API:

	beacon init

Beacon Init handles setting up beacon for the first time. It creates a `.beaconrc` at the root of your project, and uses git information to provide message log metadata.
Will also create a beacon_log.json file at your project root if it does not yet exist.

	beacon add

Provides functionality to add new breaking changes to the beacon log file.

	beacon show <INT>

Displays a variable number of breaking changes as per <INT>

	beacon all

Displays all breaking changes.


Development

This project heartily welcomes new developers. Please consider looking into contributing.
This application was intended to be sketched out in a month's time, with further refinement and polish beyond that deadline.
Beyond the time restrictions, the project maintainer (yours truly) is writing in Go for their first time.
If you should come across any Go code that could use a refactor to become both cleaner, faster, or in general, more idiomatic, please make a PR along with an explanation of your changes and why it better fits Go idioms, etc.

Now, let's consider the main layout and structure of the project:

Our project is currently broken into two main pieces of functionality (along with a few extraneous functions / files.)

One part of our project (messages.log) handles everything to do with the creation, reading and writing of Logs to the Beacon Log.
messages.log interacts closely with the `beacon_log.json` file.


Credits and Licence

Beacon is a project from THE HEAP (http://github.com/the-heap). This project is open source and you are free to do with it as you like. If you enjoy the project, please let us know.
Keep up to date with THE HEAP on twitter (http://twitter.com/theheap_) and consider becoming a contributor either on forthcoming projects or work on previous projects.

// TODO: put in contributors


Terms and Glossary

- Breaking Change: Any change to a codebase that requires all members working on said codebase to upgrade, update, migrate etc in order to get back to a working development environment.

- Log: A message in the Beacon Log, describing a breaking change.

- Beacon Log - a JSON file responsible for chronicling breaking changes for a team. This file should be checked into your repository

*/
package main
