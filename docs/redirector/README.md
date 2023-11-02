# Redirector Service

A simple web service that takes incoming request and redirects it to the respective website.

## Manual

### How does redirections work

The system listens for HTTP requests with the entry's `UUID` as the single element in the path (ex. `https://do.main/EntrysUUID`)

Validates that the `UUID` has the correct length, then searches it on the database, and finally it redirects the user to the Target Link stored on the database for the given `UUID`.

If an error occurs anywhere in this chain of actions it simply responds with this error screen.

<img src="../images/layout/RedirectorError.png" alt="RedirectorError" width="412" />

---
> ### Disclaimer
> I do not recommend to open any of the services to the internet. I didn't implement User Auth on purpose. I designed this system as an exercise to develop something simple with the ROM Stack and *maybe* use it as part of my Home Lab network. - [@NordGus](https://github.com/NordGus)

*Built with the ROM Stack*
