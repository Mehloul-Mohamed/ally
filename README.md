<h1 align="center"> butler </h1>

<p align="center">Helper utility to make playing CTFs a little less messy.</p>

> [!Important]
> While the tool is somewhat functional currently, it is far from complete and the codebase is in need of a refactor.

## Usage  
```
usage: butler {start,list,attempt}
positional aguments:
  {start,list,attempt}
        start                   Start a CTF
        list                    Show challenge list
        attempt                 Attempt a challenge
```

`start`:
```
usage: butler start name url token

positional aguments:
  name                  CTF name
  url                   CTF url
  token                 Your API token
```
`list`:  
```
usage: butler list
```
`attempt`:
```
usage: butler attempt id

positional aguments:
  id                    Challenge id
```
