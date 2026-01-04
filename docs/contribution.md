# Contribution Guidelines
---

This is still a very young project, and we are ironing out the contribution process. If any of these guidelines are actually ridiculous, create a discussion and let me know! I'd be happy to ammend them as this project grows.

## Features and Issues

For the time being, we're going to try to avoid having users submit issues directly. Create a discussion under the category of #potential-bug, or #feature-idea. Discussions will help us iron out what the bug fix or feature needs to look like before we dive in to implementing it.

If you are creating a discussion for a potential bug, please provide exports of your class and character, with details on how to reproduce the issue. Include details from the log file if applicable.

## Pull Requests

If you want to contribute with a PR, take a look issues tagged 'Help Wanted' or 'Good First Issue'.

This will all be pretty standard, just go about the following at your leisure

- Fork the repo
- Make your changes on your forked branch
- Make sure tests are passing before creating a PR `go test ./...`
- When you do create a PR, reference the issue for the feature/bug that you are addressing
- In the PR description, include details with how you tested your change
- If you modified code that impacts other features, try to regression test them to verify all is above board.

Any help is welcome, and if you're unsure of how to go about something, feel free to start a discussion or shoot me an email at orion.callaghan@gmail.com. Again, we'd love to have you help!
