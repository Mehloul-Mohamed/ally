<h1 align="center"> Ally </h1>

<p align="center">Helper utility to make playing CTFs a little less messy.</p>

> [!Important]
> While the tool is somewhat functional currently, it is far from complete and the codebase is in need of a refactor.

## Usage
```
usage: ally {start,list,attempt}
positional aguments:
  {start,list,attempt}
        start                   Start a CTF
        list                    Show challenge list
        attempt                 Attempt a challenge
```

`start`:  
This will create a directory inside `~/ctf/` with the name of the CTF.  
Inside the directory there is one file:
- `credentials.txt` : Containing the CTF url and api token
```
usage: ally start name url token

positional aguments:
  name                  CTF name
  url                   CTF url
  token                 Your API token
```
`list`:  
(must be used while in the directory of a CTF)  
```
usage: ally list
```
`attempt`:  
(must be used while in the directory of a CTF)  
This will create a directory for the challenge inside the directory of the CTF and will download the files for the challenge into it.
```
usage: ally attempt id

positional aguments:
  id                    Challenge id
```
