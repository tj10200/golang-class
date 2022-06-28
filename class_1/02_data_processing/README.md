# Basic Directory Layout for New project
This file contains the basics for creating a hello world CLI application

## Directory Structure

Project root contains the [main.go](main.go) application entry point as well
as the [Dockerfile](Dockerfile) and project `README.md`

### Dockerfile


### Other directories
- cmd
  - This directory contains the cobra commands that run the application. There is
    always one root command containing 0+ child commands that comprise the application.
- config
  - This directory contains the config used to run the application in whatever 
    ways the application is intended to run. For this example, it contains the variables
    `hello` and `world`.
- pkg
  - This directory contains the application logic that would be presentable as an exported
    library.


## More info
More information about 
  