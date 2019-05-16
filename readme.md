# Slack Status Manager
Sets status of slack workspaces based on whether or not a zoom call is open.

## Configuration & Installation

1. First, get a user token from slack [here](https://api.slack.com/custom-integrations/legacy-tokens)
1. Create a configuration file like the following and store in `~/slack-status-config.json`
    ```json
    [
      {
        "name": "Slack Workspace 1",
        "token": "xoxp-123456890abcdefghijklmnopqrstuvwxyz"
      },
      {
        "name": "Slack Workspace 2",
        "token": "xoxp-123456890abcdefghijklmnopqrstuvwxyz",
        "meetingStatus": {
          "status_text": "I'm in a meeting",
          "status_emoji": ":warning:"
        },
        "noMeetingStatus": {
            "status_text": "",
            "status_emoji": ":hp-hufflepuff:"
        }
      }
    ]
    ```
1. Copy the zoom-status.app file to your Applications directory
1. Run the app
1. You should be prompted to allow the app accessibility features. If not, open System Preferences->Security & Privacy -> Privacy and add the app to the "Accessibility" section
