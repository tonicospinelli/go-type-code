go-type-code
==============

It is a repository to explain how to apply the pattern *"Replace Type
Code with ~Class~ Struct"* from book *Refactoring To Patterns* written
by *Joshua Kerievsky* over Go language

> **What's *type code*?** The constants that you have to specify number/
> strings with understandable names.

#### Problem

A STRUCT has a field that contains type code. The values of this type
**aren't used in operator conditions** and **don’t affect the behavior
of the program**.

```
|         Payment        |
|------------------------|
| status          string |
| StatusInitial   string |
| StatusWaiting   string |
| StatusPaid      string |
| StatusCancelled string |
```

#### Solution

Create a new STRUCT and use its instance instead of the type code
values.

##### Solution with STRING
```
|    Payment    |     |       Status     |     
|---------------|  -> |------------------|
| status Status |     | Initial   string |
                      | Cancelled string |
                      | Paid      string |
                      | Waiting   string |
```

##### Solution with INT (using IOTA)
```
|    Payment    |     |     Status    |     
|---------------|  -> |---------------|
| status Status |     | Initial   int |
                      | Cancelled int |
                      | Paid      int |
                      | Waiting   int |
```
