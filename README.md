# Blogger Tools

Some tools to help with Blogger related tasks.

Yeah! I have a v0.1.0 ready to go!
These tools are currently only tested on MacOS.

*Go version: 1.11.5*

## Configuration
The tools use a unified config file. The tools will look for a config file in `./conf.json` by default.

* **LogDirectory**: the directory where log files will be stored. this value must be an absolute path.
* **SecretsDirectory**: the directory where files with sensitive data are stored. this value must be an absolute path.
* **Blogger**: configurations for Blogger assets.
    * **Blogs**: configurations for individual blogs.
* **SendGrid**: configurations for SendGrid assets.
    * **APIKey**: the API key used to access SendGrid.
    * **DefaultFromEmail**: the default address to send email from.
    * **DefaultFromName**: the default name to send email from.
* **Logs**: filenames of log files for specific tools. filenames are relative to the `LogDirectory`.
    * **General**: filename of general log file.
    * **NotifyTool**: filename of *NotifyTool* log file.
    * **OAuthTool**: filename of *OAuthTool* log file.
* **NotifyTool**: configurations specific to the NotifyTool.
    * **BlogUpdatedRecipientsFile**: filename relative to `SecretsDirectory` of file containing a JSON array of emails.
    
    ```
    // BlogUpdatedRecipientsFile file example
    {
       "BLOG_ID": [
         {
           "Name": "Test User",
           "Email": "test_user@somewhere.tld"
         }
       ]
     }
     ``` 
     
## Notify Tool
This tool sends notifications for various events connected with Blogger.
SendGrid is used to send out email notifications.
 
```
// Build
go build -o ./bin/macos/notify ./cmd/notify/...
```

```
NAME:
   Blogger Notify Tool - send notification of blog events

USAGE:
   notify [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     diagnostic   run notify tool diagnostics
     blog-update  send notification for a blog update
     help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value  filename of the configuration file to use.
   --help, -h      show help
   --version, -v   print the version
```
 
### Diagnostic
This command will run diagnostics to ensure that the tool can connect successfully to the Blogger API.
```
// Usage
./bin/macos/notify diagnostic
```

### Blog Update
This command will send out notifications to inform users that the blog has been updated.
```
// Usage
./bin/macos/notify blog-update
```

## OAuth Tool
This tool authenticates the user with Google using OAuth and stores the generated access token.
This tool is used within the NotifyTool if a Google access token is not found, but can also be used separately to generate the token.

```
// Build
go build -o ./bin/macos/oauth ./cmd/oauth/...
```

```
NAME:
   OAuth Tool - perform oauth workflow to gain access to the Google Blogger API

USAGE:
   oauth [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     run      run the oauth workflow
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value  filename of the configuration file to use.
   --blogs value   ids of the blogs to authenticate.
   --help, -h      show help
   --version, -v   print the version
```

### Run
This command will run the oauth tool, providing an authentication URL to visit. 
The user will be provided an access key by Google once authentated via the URL. 
The user should then enter the access key into OAuthTool. 
If an *AccessTokenFile* is configured for the blog then the access token entered by the user will be stored for use with other tools.

```
 // Usage
 ./bin/macos/oauth run
```