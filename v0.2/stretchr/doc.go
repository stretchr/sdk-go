/*
Package provides access to Stretchr data services from within Go code.

To access resources, first create a Session object:

	session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")

The session object will hold information about your project and your credentials for the
Stretchr data services.

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
    log.Printf("New resource created with ID: %s", response.GetID())
	}

*/
package stretchr
