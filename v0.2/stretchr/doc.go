/*
Package provides access to Stretchr data services from within Go code.

To access resources, first create a Session object:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

The session object will hold information about your project and your credentials for the
Stretchr data services.

You can then use the following methods to interact with the data services:


Quick examples

Creating a resource

 err := session.MakeResource("people").Set("name", "Mat").Create()

Reading a resource by ID

 resource, err := session.Read("people", "123")

Reading all resources

 resources, err := session.Many("people").Read()

Update a resource

 err := session.MakeResource("people").SetID("123").Set("surname", "Ryer").Update()

Replace a resource

 err := session.MakeResource("people").SetID("123").Set("fullname", "Mat Ryer").Replace()

To delete a resource by ID

 err := session.Delete("people", "123")

Deleting all resources

 err := session.Many("people").Delete()



Reading a resource

To read a resource when you know the ID, use the `Read` method:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 // read person with ID "mat"
 mat, readErr := session.Read("people", "mat")

 if readErr != nil {
   panic(fmt.Sprintf("Failed to read 'mat': %s", readErr))
 } else {
   log.Printf("Mat's full name is: %s %s.", resource.Get("first-name"), resource.Get("last-name"))
 }


Reading multiple resources

To read a collection of resources, use the `Many` and `Read` methods:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 people, readErr := session.Many("people").Read()

 if readErr != nil {
   panic(fmt.Sprintf("Failed to read people: %s", readErr))
 } else {
   for _, person := range people.Resources {
     log.Printf("I have found %s", person.Get("name"))
   }
 }


Paging

To get pages one and two of a collection of resources, with 10 resources in a page, use
the `Many`, `Page` and `Read` methods:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 peoplePage1, readErr := session.Many("people").Page(1, 10).Read()
 peoplePage2, readErr := session.Many("people").Page(2, 10).Read()



Create a new resource

To create a new resource, use the `Create` method:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 // make a resource
 resource := session.MakeResource("people")

 // set some data
 resource.Set("name", "Mat").Set("age", 29).Set("when", time.Now())

 // call Create
 createErr := resource.Create()

 if createErr != nil {
   panic(fmt.Sprintf("Failed to create resource: %s", createErr))
 } else {
    log.Printf("New resource created with ID: %s", resource.GetID())
 }


Updating resources

To make some changes to (update) a resource, use the `Update` method:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 // find a resource
 resource := session.MakeResource("people").SetID("123")

 // set the fields you want to change
 resource.Set("age", 30)

 // update the resource
 updateErr := resource.Update()

 if updateErr != nil {
   panic(fmt.Sprintf("Failed to update resource: %s", updateErr))
 } else {
    log.Printf("Updated %s's age", resource.Get("name"))
 }

Notice that we never actually set the 'name' field, this is because when you do an Update,
the entire resource is returned by Stretchr and loaded into the resource.



Replace a resource

To completely replace a resource, use the `Replace` method:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 // find a resource
 resource := session.MakeResource("people", "mat")

 // set the fields you want to change
 resource.SetID("ABC123").Set("age", 30).Set("name", "Mat")

 // update the resource
 replaceErr := resource.Replace()

 if replaceErr != nil {
   panic(fmt.Sprintf("Failed to replace resource: %s", replaceErr))
 } else {
    log.Printf("Replaced resource ID: %s", resource.GetID())
 }


Deleting resources

To delete a resource when you know the ID, use the `Delete` method:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 // delete person with ID "mat"
 deleteErr := session.Delete("people", "mat")

 if deleteErr != nil {
   panic(fmt.Sprintf("Failed to delete 'mat': %s", deleteErr))
 } else {
   log.Printf("Mat has been deleted!")
 }

If you have the resource object, you can directly delete it using the `Resource.Delete` method:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 // load mat
 mat, _ := session.Find("people", "mat")

 // delete person with ID "mat"
 deleteErr := mat.Delete()

 if deleteErr != nil {
   panic(fmt.Sprintf("Failed to delete 'mat': %s", deleteErr))
 } else {
   log.Printf("Mat has been deleted!")
 }

There's little point in finding a resource just to delete it, but if you happen to
have the resource anyway it makes sense to use the `Resource.Delete` method.

To delete multiple resources based on specific criteria, you can use the `Many.Delete` method:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 // load mat
 deleteErr := session.Many("people").Where("age", "<18").Delete()

 if deleteErr != nil {
   panic(fmt.Sprintf("Failed to delete 'mat': %s", deleteErr))
 } else {
   log.Printf("Everyone under 18 has been deleted!")
 }


Filtering


To find resources of people between the ages of 18 to 30, who have signed up to receive 
email updates:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 people, readErr := session.Many("people").Where("age", "18..30").Where("get-updates", true).Read()

 if readErr != nil {
   panic(fmt.Sprintf("Failed to read people: %s", readErr))
 } else {
   for _, person := range people.Resources {
     log.Printf("I have found %s", person.Get("name"))
   }
 }

*/
package stretchr
