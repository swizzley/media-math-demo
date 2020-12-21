# IMDBaaS

## Overview

The intent is that this product will be used by customers to discover interesting things about the
movies they and their friends watch and the actors they love. 

### Solution Statement

A solution would need to handle not only enormous amounts of data, but — since this WILL be the most popular 
entertainment reference in the universe — be infinitely scalable.

### Feature Set

  * Concurrency
  * High availability
  * Fault-tolerant 
  
### Architecture 

    The diagram below is an example of the HA backend, accorss 3 Availability zones in each of the two regions provisioned.
    Network and and security profiles were omitted for simplicity. 

    ![Architecture](/docs/arch.png)
    
### Testing

```DEBUG=true make build; make test```

### Deployment 

#### Local

Run `make build` or `DEBUG=true make build` 

#### Cloud

Choose an environment, choose a region, and then finally choose a version, then just `tf apply` **INCOMPLETE / TEMPLATE ONLY**

#### Design

By using a stateless REST service to administer the gRPC backend, the front end will have easy single point calls for
any functionality. No need to implement graph clients in front end code, no need for multiple middleware calls, and 
best of all, full utilization of go multi threading and concurrency native to the language. The router will handle all
concurrency, and the stateless nature will allow infinite scalability. 


# Not Implemented

Actually anything fully functioning. I spent too much time learning neo4j and then dgraph, and troubleshooting both. I've spent 
over 4 hours so far, and this documentation is the last thing I'm doing. I started down the path for terraform but realized
that the architecture I diagrammed would be a whole day project unto itself. I would have liked to implement Istio, consul, 
or other cool stuff but I was told this would take me 3-4 hours and there's just no room for that stuff. I think by the work
that I have done, you can get a rough idea of how I engineer, and maybe why/how I got stuck where I did. It's now to the point
where I can't justify putting in the time for something I'm not getting paid to do. I think I'm pretty close to actually 
interacting with the database, and from that point I think building out the REST CRUD would go pretty quickly. From that point
I would build out user function handlers and then might experiment with kubernetes and try running tests at scale, 
like 500,000 calls/s to see where things break and maybe implement kafka or something similar. I've never used a graphDB 
before so I have no idea where it's issues are or how to architect it for scale, the diagram given was simply based on the 
HA composer file I found in documentation. 