
# <img src="https://www.telebox.online/TeleBox/icon.png" width="24" height="24" /> TeleBox

TeleBox is [a private cloud drive](https://www.telebox.online/).

## Configuration

Here is an example of making a remote for TeleBox.

First run:

     rclone config

This will guide you through an interactive setup process:

```
No remotes found, make a new one?
n) New remote
s) Set configuration password
q) Quit config
n/s/q> n

Enter name for new remote.
name> remote

Option Storage.
Type of storage to configure.
Choose a number from below, or type in your own value.
XX / TeleBox
   \ (TeleBox)
Storage> XX

Option token.
Authorize on https://telebox.online, then go to developer console of the browser and paste here the output of `localStorage.t` command
Enter a value.
token> testFromCLToken

Configuration complete.
Options:
- type: telebox
- token: XXXXXXXXXXX
Keep this "telebox" remote?
y) Yes this is OK (default)
e) Edit this remote
d) Delete this remote
y/e/d> y

```

### Standard options

Here are the Standard options specific to telebox (TeleBox).

#### --telebox-token

Authorize on https://telebox.online, then go to developer console of the browser and paste here the output of `localStorage.t` command

Properties:

- Config:      token
- Env Var:     RCLONE_TELEBOX_TOKEN
- Type:        string
- Required:    true

### Advanced options

Here are the Advanced options specific to telebox (TeleBox).

#### --telebox-description

Description of the remote.

Properties:

- Config:      description
- Env Var:     RCLONE_TELEBOX_DESCRIPTION
- Type:        string
- Required:    false

## Limitations

Invalid UTF-8 bytes will also be [replaced](https://rclone.org/overview/#invalid-utf8),
as they can't be used in JSON strings.
