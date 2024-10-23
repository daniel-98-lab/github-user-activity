# GitHub User Activity CLI - Fetch GitHub User Events with Ease

This is a simple command-line interface (CLI) tool written in Go that retrieves user activity events from the GitHub API. It allows you to view the most recent events for a specific GitHub user, with the option to filter events by type.
Challenge from github-user-activity: https://roadmap.sh/projects/github-user-activity

## Features

- Fetch recent activity events for a GitHub user.
- Optionally filter events by type (e.g., `PushEvent`, `ForkEvent`, `IssuesEvent`, etc.).
- Display details of the events in a user-friendly format.

## Installation

To get started, you will need to have Go installed on your machine. Follow the instructions [here](https://golang.org/doc/install) if you don't have Go set up.

1. Clone this repository:

    ```bash
    git clone https://github.com/daniel-98-lab/github-user-activity.git
    ```

2. Navigate to the project directory:

    ```bash
    cd github-user-activity
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Build the project:

    ```bash
    go build -o user-activity ./cmd
    ```

## Usage

To use the tool, run the following command from the terminal:

```bash
./user-activity <username> [eventType]
```

- <username>: Github username to fetch events (required)
- [eventType]: Optional parameter to filter by event type. Available event types include:
    + PushEvent
    + CreateEvent
    + DeleteEvent
    + ForkEvent
    + IssueCommentEvent
    + IssuesEvent
    + PullRequestEvent
    + PullRequestReviewEvent
    + PullRequestReviewCommentEvent
    + WatchEvent
    + ReleaseEvent

## Example

1. Fetch all events for a user:

    ```bash
    ./user-activity daniel-98-lab
    ```

2. Fetch only PushEvent for a user:

    ```bash
    ./user-activity daniel-98-lab PushEvent
    ```

## Error Handling

If an error occurs (e.g., invalid user or event type), the application will display an error message in the terminal.

This project is licensed under the MIT License.