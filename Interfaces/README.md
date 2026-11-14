we are now learning interfaces in golang, 

Interfaces refers to methods/rules that a class/struct must follow. They are showing what a type must do, without 
specifying how it will do it. 

Interfaces ensure that objects can be used interchangeably since they use the same method name. It can also help in decoupling because you don't need to worry about the details all the time. 

An example is we can have animals make sound, the interface is speak since all of them make a sound. The interface doesn't need to specify how a dog, cow, or cat sounds but it can specify that animals must speak.

Another example, is a payment method. Like Mpesa, Paypal, and Credit card use different ways of paying eg: mobile money or bank. An interface method specifies that all these have to pay without specifying how.  

We can say: type Animal interface {
		speak()
		}
type Payment interface{
	pay()
}

Pay and speak are the methods that any types trying to implement the interfaces must use. 
type defines a new type is created, Animal and Payment are the names of the interfaces. 

Now we can create structs to satisfy the interfaces.
