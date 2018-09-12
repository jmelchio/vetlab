## Vetlab

### Overview
A simple service API that provides services for keeping track of customers,
veterinary practices and their requests for laboratory diagnostics and resulting
reports that are delivered back to them.

### Go language
The project is also meant to help me figure out how to properly put together
a Go program that is testable and relatively easy to maintain and expand.

The structure is meant to group code into packages that define interfaces for
those things that they are dependent on so that they can be tested without too
much trouble and to allow the implementations to be easily swapped out if the
need for that arises.

#### Libraries used
* [Ginkgo](https://onsi.github.io/ginkgo/) for testing
* [Rata](https://github.com/tedsuo/rata) for routing
* [Gorm](http://gorm.io/) for persistence
