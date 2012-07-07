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

Various parameter and limiting methods exist that allow you to fine tune the request.  These methods
chain, that is they return the same object allowing you to make multiple calls on one line which makes it
very easy to read and write.  For example, the following code will read page 2 of the "people" resources,
with 10 resources in a page, ordered by age descending:

	rs, _ := session.ReadMany("people").Skip(9).Limit(10).Order("-age").Now()
*/
package stretchr
