# Stretchr Go SDK

The Stretchr Go SDK allows you to interact with Stretchr services in your Go code.

## Quick overview

### Create a shared `Session` object

To interact with Stretchr, you need a `Session` object.  A session object holds details about the Stretchr project you are interacting with, and the authentication that gives you access to do so.

    Stretchr := stretchr.NewSession("{project}", "{public-key}", "{private-key}")

  * `{project}` - The fully qualified name of the Stretchr project you wish to interact with.
  * `{public-key}` - The public key of an account with permission to access the project.
  * `{private-key}` - The private key of the same account.
  
For example:

    Stretchr := stretchr.NewSession("test.project.company", "ABCDEFGHJIKLMNOPQRSTU", "VWXYZABCDEFGHJIKLM")

  
It is good practice to name your variable `Stretchr` (as in the example above) because it makes the rest of your code very easy to read.  The `Stretchr` variable in the following code samples refers to a `stretchr.Session` object.

### Creating a resource

To persist something in Stretchr, you just need to create a resource and call the `Create` method.  In this example, we are going to persist a person resource at `/people`:

    // make a resource and set some data
    resource := stretchr.MakeResourceAt("people").
                         Set("name", "Mat").
                         Set("age", 30)
    
    // create the resource
    change, err := Stretchr.At(resource.ResourcePath()).Create(resource)
    
    // check for errors
    if err != nil {
      // TODO: handle errors
    }
    
    // log the ID that Stretchr generated for us
    log.Printf("Person was created and given ID: %s", resource.ID())