/*
Package provides access to Stretchr data services from within Go code.

To access resources, first create a Session object:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

The session object will hold information about your project and your credentials for the
Stretchr data services.

You can then use the following methods to interact with the data services:

 // creating a resource
 err := session.MakeResource("people").Set("name", "Mat").Create()

 // reading a resource by ID
 resource, err := session.Read("people", "123")

 // reading all resources
 resources, err := session.Many("people").Read()

 // update a resource
 err := session.MakeResource("people").SetID("123").Set("surname", "Ryer").Update()

 // replace a resource
 err := session.MakeResource("people").SetID("123").Set("fullname", "Mat Ryer").Replace()

 // to delete a resource by ID
 err := session.Delete("people", "123")

 // deleting all resources
 err := session.Many("people").Delete()

To read a resource when you know the ID:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 // read person with ID "mat"
 mat, readErr := session.Read("people", "mat")

 if readErr != nil {
   panic(fmt.Sprintf("Failed to read 'mat': %s", readErr))
 } else {
   log.Printf("Mat's full name is: %s %s.", resource.Get("first-name"), resource.Get("last-name"))
 }

To read a collection of resources:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 people, readErr := session.Many("people").Read()

 if readErr != nil {
   panic(fmt.Sprintf("Failed to read people: %s", readErr))
 } else {
   for _, person := range people.Resources {
     log.Printf("I have found %s", person.Get("name"))
   }
 }

To get pages one and two of a collection of resources, with 10 resources in a page:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 peoplePage1, readErr := session.Many("people").Page(1, 10).Read()
 peoplePage2, readErr := session.Many("people").Page(2, 10).Read()

To create a new resource:

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

To make some changes to (update) a resource:

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

To completely replace a resource:

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

To delete a resource when you know the ID:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 // delete person with ID "mat"
 deleteErr := session.Delete("people", "mat")

 if deleteErr != nil {
   panic(fmt.Sprintf("Failed to delete 'mat': %s", deleteErr))
 } else {
   log.Printf("Mat has been deleted!")
 }

If you have the resource object, you can directly delete it using the Delete method:

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
have the resource anyway it makes sense to use the Resource.Delete method.

To delete multiple resources based on specific criteria, you can use the Many.Delete method:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")
 
 // load mat
 deleteErr := session.Many("people").Where("age", "<18").Delete()

 if deleteErr != nil {
   panic(fmt.Sprintf("Failed to delete 'mat': %s", deleteErr))
 } else {
   log.Printf("Everyone under 18 has been deleted!")
 }

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
