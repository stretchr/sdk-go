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
    changes, err := Stretchr.At(resource.ResourcePath()).Create(resource)
    
    if err != nil {
      // TODO: handle errors
    }
    
    if changes.Created() == 1 {
      
      // log the ID that Stretchr generated for us
      log.Printf("Person was created and given ID: %s", resource.ID())
    
    }
    
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

Stretchr allows you to modify the resources that are returned by further _filtering and ordering_ your request before you call `ReadMany`.  The following examples demonstrate how this works:

    // find people whose age is greater than 18
    collection, err := Stretchr.At("people").Where("age", ">18").ReadMany()

    // find the three oldest people
    collection, err := Stretchr.At("people").Order("-age").Limit(3).ReadMany()

    // find the next three oldest people (paging)
    collection, err := Stretchr.At("people").Order("-age").Page(2, 3).ReadMany()

    // find people called John
    collection, err := Stretchr.At("people").Where("name", "John").ReadMany()

    // find people NOT called John
    collection, err := Stretchr.At("people").Where("name", "!John").ReadMany()

### Updating a resource

To update the data in of existing resource, you can use the `Update` function.  Updating will change any fields present, and leave alone fields not mentioned in the resource.  If you do want to replace the resource wholesale, see `Replace`.

This example updates the name of person `123`:

    resource := stretchr.MakeResourceAt("people/123").
                         Set("name", "Mathew")
    
    changes, err := Stretchr.At(resource.ResourcePath()).Update(resource)
    
    if err != nil {
      // TODO: handle errors
    }
    
    if changes.Updated() == 1 {
    
      log.Printf("I just updated person 123's name to: %s", resource.Get("name"))
      
    }

### Replacing a resource

Since `Update` modifies a resource with the specified data, it can sometimes leave unwanted data in the resource.  The Go SDK provides a handy `Replace` function that is similar to deleting a resource and recreating it (with the same ID).

Assuming we have an existing resource with lots of data at `people/123`, the following code will leave the same resource with *only* a `name` field:

    resource := stretchr.MakeResourceAt("people/123").
                         Set("name", "Mathew")
    
    changes, err := Stretchr.At(resource.ResourcePath()).Replace(resource)
    
    if err != nil {
      // TODO: handle errors
    }
    
    if changes.Updated() == 1 {
    
      log.Printf("I just replaced person %s entirely", resource.ID())
      
    }
    
### Deleting resources
    
#### Deleting a single resource

Once you have finished with a resource and wish to delete it forever, you can use the `Delete` function.  To delete person `123`, we just need to do:

    changes, err := Stretchr.At("people/123").Delete()
    
    if err != nil {
      // handle errors
    }
    
    log.Printf("%d resource(s) were deleted.", changes.Deleted())
    
#### Deleting many resources

Deleting many resources uses the same function as when you are deleting a single resource, except that you provide a different path:

    changes, err := Stretchr.At("people").Delete()
    
    if err != nil {
      // handle errors
    }
    
    log.Print("ALL people were deleted")

The same filtering, limiting and skipping methods that work when using `ReadMany` also work when deleting resources.  For example, to delete all people called John you could do:

    changes, err := Stretchr.At("people").Where("name", "John").Delete()

## Advanced

The Stretchr Go SDK is built on top of a lower level `api` package that deals with requests and responses rather than Resources.  If you wish to perform advanced actions, or be in more control over how your code interacts with the Stretchr services you are free to use the advanced `api` package instead.

Much of the interactions are the same as using the higher level SDK, but you tend to have more control over the specifics, including access to the underlying `http.Request` object that ends up talking to the services.

For more information, check out the `api` documentation.