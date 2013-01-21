# Stretchr Go SDK

The Stretchr Go SDK allows you to interact with Stretchr services in your Go code.  You should become familiar with how Stretchr works first by visiting the [Stretchr website](http://www.stretchr.com/).

## Quick overview

This guide covers the primary functions of Stretchr, and how you execute such activities using the Go SDK.

### Create a shared `Session` object

To interact with Stretchr, you need a `Session` object.  A session object holds details about the Stretchr project you are interacting with, and the authentication that gives you access to do so.

    Stretchr := stretchr.NewSession("{project}", "{public-key}", "{private-key}")

  * `{project}` - The fully qualified name of the Stretchr project you wish to interact with.
  * `{public-key}` - The public key of an account with permission to access the project.
  * `{private-key}` - The private key of the same account.
  
For example:

    Stretchr := stretchr.NewSession("test.project.company", "ABCDEFGHJIKLMNOPQRSTU", "VWXYZABCDEFGHJIKLM")

  
It is good practice to name your variable `Stretchr` (as in the example above) because it makes the rest of your code very easy to read.  Be careful to remember the difference between uppercase `Stretchr`, the session variable, and lowercase `stretchr`, the package.  The `Stretchr` variable in the following code samples refers to a `stretchr.Session` object.

### Creating a resource

To persist something in Stretchr, you just need to create a resource and call the `Create` method.  In this example, we are going to persist a person resource at `/people`:

    // make a resource and set some data
    resource := stretchr.MakeResourceAt("people").
                         Set("name", "Mat").
                         Set("age", 30)
    
    // create the resource
    change, err := Stretchr.At(resource.ResourcePath()).Create(resource)
    
    if err != nil {
      // TODO: handle errors
    }
    
    // log the ID that Stretchr generated for us
    log.Printf("Person was created and given ID: %s", resource.ID())
    
### Reading resources

Once you have data in Stretchr, it would be nice to get it back.  Luckily, we thought of this.
    
#### Reading a single resource

If you know the path and ID of a resource you wish to read, you can do so using the `ReadOne` method as in the following example where we read the `people` resource with ID `123`:

    resource, err := Stretchr.At("people/123").ReadOne()
    
    if err != nil {
      // TODO: handle errors
    }
    
    log.Printf("Just loaded %s who is %g years old.", resource.Get("name"), resource.Get("age"))

#### Reading many resources

If you want to load a collection of resources in one go, you can do so using the `ReadMany` function:

    collection, err := Stretchr.At("people").ReadMany()
    
    if err != nil {
      // TODO: handle errors
    }
    
    // iterate over each person
    for _, resource := range collection.Resources {
      log.Printf("Did you know %s was %g years old?", resource.Get("name"), resource.Get("age"))
    }