
## vCloudLogs
This module is used to enhance logging information with metadata such as filename and line number of the log.
Initialization of the logger allows the user to set where the logs are written, use os.Stdout, or os.Stderr to
work with default google-fluentd in GKE. ioutil.Discard allows particular log levels to be ignored.

### What you get:
1. A tag, which all logs using will be prepended with (Use "" when you call init if you don't need a tag)
2. The ability to log Trace, Debug, Info, Warning, Error and Critical levels.
    -Caveat: These can't be filtered in the Google Cloud Log Viewer, only searched.
3. The filename which the log originated.
4. The line number which the log originated.
5. The ability to use formatted or non-formatted log calls.
6. An endpoint handler to enable/disable Trace/Debug log levels without redeploying your application

## Installation
```
$ go get github.com/vendasta/vCloudLogs
```

## Usage

### Import
```
import "github.com/vendasta/vCloudLogs"
```

### The logger must be initialized
This example sets trace and debug as disabled, all other levels as enabled:
```
vCloudLogs.InitLoggers("[TAG]",
		ioutil.Discard, //Trace
		ioutil.Discard, //Debug
		os.Stdout,      //Info
		os.Stdout,      //Warning
		os.Stderr,      //Error
		os.Stderr)      //Critical
```

### Then call the logger
```
msg := "Oh no, an error happened"
vCloudLogs.Errorf("What happened?: %s", msg)
```

### Using the Trace/Debug on/off handler
```
http.HandleFunc("/logging", func(w http.ResponseWriter, r *http.Request) { vCloudLogs.LoggingOnOffHandler(w, r) })
```
