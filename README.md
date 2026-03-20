# GitHub User Activity CLI

A command-line tool to fetch and display recent activity of a GitHub user. This is an implementation of the [GitHub User Activity](https://roadmap.sh/projects/github-user-activity) project idea from roadmap.sh.

## Prerequisites

- Go 1.19 or later installed on your system.

## Installation

Install the CLI using Go:

```bash
go install github.com/ne1ascorbinka/github-activity@latest
```

## Usage

Run the tool with a GitHub username:

```bash
github-activity <username>
```

### Example Output

```
- Pushed 3 commits to kamranahmedse/developer-roadmap
- Opened a new issue in kamranahmedse/developer-roadmap
- Starred kamranahmedse/developer-roadmap
- ...
```