
## Commands
```zsh
bigary <Command> ...
```
----
### init
```zsh
bigary init <Root>
```
make initioal directories and files from template **clean**




----
### make

```zsh
bigary make <what>
```
what:
<!-- what:
- module
- interface -->

```zsh
bigary make module <>
```



****


----
### where
```zsh
bigary where
```
show you where are you






{
   registerRoute:{
    v1 -> moduleX_Routes

    v1 -> moduleY_Routes
   }

   module :{
    n*( controller -> service(interface) -> repository (interface))

    routes -> n* version -> controller


   }
    
}

bigary make 


entity:
    module

    routes
    controller
    service
    service interface
    repository 
    repository interface


do:
    module:
        routes
        controller
        service
        service interface
        repository 
        repository interface
    controller:
        controller
        service
        service interface
        repository 
        repository interface


![enter image description here](diagram.svg)