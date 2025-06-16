# Gofish

Gofish is a simple HTTP server for uploading, listing, downloading, and serving images.  
Heavily inspired by [Gofus](https://github.com/mqnr/gofus).

## Features

- Upload images via HTTP POST
- List all uploaded images
- Download images by filename
- Serve images directly via URL
- Simple HTML frontend (optional)

## Getting Started

### Prerequisites

- Go 1.18 or newer

### Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/gofish.git
cd gofish
```

### Build & Run

```bash
go build -o gofish
./gofish
```

The server will start on `http://localhost:8080` by default.

## Usage

### Upload an Image

Send a POST request to `/upload` with a form field named `file`:

```bash
curl -F "file=@yourimage.png" http://localhost:8080/upload
```

### List Images

Visit `http://localhost:8080/list` or send a GET request to `/list` to get a JSON array of filenames.

### Download or View an Image

Access `http://localhost:8080/images/<filename>` to view or download an image.

### HTML Frontend

Visit `http://localhost:8080/` for a simple web interface (if enabled).

## API Endpoints

| Method | Endpoint              | Description                |
|--------|-----------------------|----------------------------|
| POST   | `/upload`             | Upload an image            |
| GET    | `/list`               | List all images            |
| GET    | `/images/<filename>`  | Download/view an image     |
| GET    | `/`                   | HTML frontend (optional)   |

## Configuration

- Images are stored in the `images/` directory by default.
- Static files (frontend) are served from `static/`.
- To change the port, set the `PORT` environment variable:

```bash
PORT=9090 ./gofish
```

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

MIT
