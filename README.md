# library
This is a simple REST API designed to manage a list of books. [go-swagger](https://github.com/go-swagger/go-swagger) was used to generate code. 
And [go-cache](https://github.com/patrickmn/go-cache/) is used as a simple key-value storage.
For info on operations please refer to swagger spec [here](../master/internal/api/spec.yaml).

# Docker
To run the library in docker I use an image with `golang` installed and build my project inside (`docker build .`).
It's a simple way to build, but the image size is huge. I'll fix it in future.
Please don't forget to map container port on your host port. I run it on my local machine this way: `docker run -p 8090:8090 [image id]`

# Query examples
To query the library API I use Postman. To create a book record call `POST` on `http://192.168.99.100:8090/books` (I run Docker on Windows machine).
The body is like `{"title": "The Hobbit", "author": "J. R. R. Tolkien", "publisher": "HarperCollins", "publication_date": "2014-12-16"}`

To check the book in call `PATCH` on `http://192.168.99.100:8090/books/{id}` with body `{"status": "checked_in"}`

# TODOs
- add transactions
- build outside and put binaries into scratch image
- think if `publication_date` should be just year, not the full date