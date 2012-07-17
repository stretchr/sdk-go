# Stretchr Command Line tool

This tool allows you to access and manipulate Stretchr services from the command line.

## Building the tool

Since the tool is written in [Google's Go language](http://www.golang.org/), you will need to install 
that first.

Then use `go build -o stretchr stretchr_command_line.go` to build it.

Once the `stretchr` executable file has been generated, you can use it directly or add the directory to your `PATH` variable for generic use everywhere.

## Specifying your public and private keys

You can provide your public and private keys to the `stretchr` command every time you make a request, or you can specify the values using environment 
variables, which makes your work a little less messy and more secure.

To specify the environment variables, open a Terminal and do:

    export StretchrPublicKey=PUBLICKEY
    export StretchrPrivateKey=PRIVATEKEY

You can also specify your project in a similar way (using `StretchrProjectName`) if you want to.

## Command line flags

The `stretchr` command makes heavy use of flags that allow you make requests.

For a full list of the supported flags, do:

    stretchr -help

Which will yield:

    Usage for ./stretchr:
      -help=false: Whether to print out this help or not.
      -host="stretchr.com": The host of the service to connect to.
      -json="": The JSON body to send in the request.
      -key="": The public key used to access the services.  Alternativly set the StretchrPublicKey environment variable.
      -method="GET": The HTTP method to use in this request (GET, POST, PUT or DELETE).
      -path="": The path of the request to make.
      -privatekey="": The private key used when signing requests for security.  Alternativly set the StretchrPrivateKey environment variable.
      -project="": The name of the Stretchr project to connect to.  Alternativly set the StretchrProjectName environment variable.
      -query="": The query part of the request to modify behaviour.
      -verbose=false: Whether to be verbose or not.

## Example usage

### Getting a list of resources

    stretchr -method="GET" -path="people"
    
### Creating a Person

    stretchr -method="POST" -path="people" -json='{"name":"Mat","age":29}'
    
### Updating a Person

    stretchr -method="PUT" -path="people/ABC123" -json='{"age":30}'

### Replacing a Person

    stretchr -method="POST" -path="people/ABC123" -json='{"fullname":"Mat Ryer","age":29}'

### Deleting a Person

    stretchr -method="DELETE" -path="people/ABC123"
    
### Deleting all people

    stretchr -method="DELETE" -path="people"
