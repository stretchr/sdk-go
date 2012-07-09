/*
Package provides access to Stretchr data services from within Go code.

To access resources, first create a Session object:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

The session object will hold information about your project and your credentials for the
Stretchr data services.

To read a resource when you know the ID:

 session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

 // read person with ID "mat"
 mat, readErr := session.Read("people", "mat")

 if readErr != nil {
   panic(fmt.Sprintf("Failed to load 'mat': %s", readErr))
 } else {
   log.Printf("Mat's full name is: %s %s.", resource.Get("first-name"), resource.Get("last-name"))
 }

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
 resource := session.MakeResource("people", "mat")

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
 resource.Set("age", 30).Set("name", "Mat")

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

NOTE: There's little point in finding a resource just to delete it, but if you happen to
need the data of a resource to decide whether or not to delete it, this is the way you would
achieve that.

*/
package stretchr
