# What is Gofish?

Gofish is an HTTP server **(In development)**

Heavily based on my friend's project "Go File Upload Service" [Gofus](https://github.com/mqnr/gofus).

## Upload images

You can upload images to the server using the `/upload` endpoint. The server will respond with a JSON object containing the URL of the uploaded image

## Download images

```bash
http://localhost:8080/download?name=[Image Name]
```

## Get All Images

```bash
http://localhost:8080/images
```