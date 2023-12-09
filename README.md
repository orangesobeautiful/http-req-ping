# HTTP Request Print

## Purpose

I simply want to use this to test the headers forwarded by an Nginx reverse proxy.

## Usage

### Start the server

```bash
go run .
```

### Connection Test

```bash
curl http://localhost:8080
```

### Sample Result

```json
{
  "Method": "GET",
  "URLPath": "/",
  "Proto": "HTTP/1.1",
  "ProtoMajor": 1,
  "ProtoMinor": 1,
  "Header": {
    "Accept": ["*/*"],
    "User-Agent": ["curl/8.5.0"]
  },
  "ContentLength": 0,
  "Host": "localhost:8080",
  "Form": null,
  "PostForm": null,
  "MultipartForm": null,
  "Trailer": null,
  "RemoteAddr": "[::1]:58336",
  "RequestURI": "/",
  "TLS": null
}
```
