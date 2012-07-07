/*
Package stretchr-go/v1/stretchr provides access to Stretchr services from Go code.

To access resources, first create a Session object:

	session := stretchr.InProject("test").WithKeys("PUBLICKEY", "PRIVATEKEY")
	
Read a resource with a specific ID:

	r, err := session.Read("people", "123").Now()
	
Read a collection of resources:

	rs, err := session.ReadMany("people").Now()
	
Create a resource:

	err := session.Create(resource).Now()
	
Update a resource:

	err := session.Update(resource).Now()
	
Replace a resource:

	err := session.Replace("people", "123", resource).Now()
	
Delete a resource:

	err := session.Delete("people", "123").Now()
	
Delete many resources:

	err := session.DeleteMany("people").Now()
*/
package stretchr