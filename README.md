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
This will create a directory inside `~/ctf/` with the name of the CTF.  
Inside the directory there are 2 files 
- `challenges.json` : Containing a list of challenges
- `credentials.txt` : Containing the CTF url and api token
```
usage: butler start name url token

positional aguments:
  name                  CTF name
  url                   CTF url
  token                 Your API token
```
`list`:  (must be used while in a CTF directory)  
```
usage: butler list
```
`attempt`:  (must be used while in a CTF directory)  
This will create a directory for the challenge inside the CTF directory and will download the challenges files into it.
```
usage: butler attempt id

positional aguments:
  id                    Challenge id
```
