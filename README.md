# s3web

s3web can serve files from AWS S3 objects.

It's common used when you wouldn't like to provide AWS Credentials to access S3 **and** the bucket is not exposed to the internet (Eg. internal web app to serve S3 data).

**WARNING**: This web application is useful to retrieve S3 data on an secure network that the users should not have AWS Credentials, please, understand the risks to expose this application to the internet. One authentication on front of the API could be provided, but it is not covered here.

## Authentication

The `s3web` can authenticate in S3 in three ways (proposal):

- the host application access the S3 through IAM Roles (instance roles) or environment variables - default behavior from SDK
- provide AWS credentials on bootstrap (TODO)
- send authentication through headers (TODO)
- send authentication through query strings (TODO, unsafe/needed?)

## Build

`make build`

## Run

### Local

```bash
$ ./bin/s3web-app
2019/11/03 04:14:26 Listening on port :8080...
```

<!-- ### Docker (TODO)-->

<!-- ### Kubernetes (TODO)-->

<!-- ### Lambda (TODO)-->

## Usage

### Config

You can specify the Bucket, Object and the Region (optional) in three ways:

- URI
- Query string
- Headers

### Get objects - Query string

```bash
curl 'localhost:8080/?bucket=myBucker&object=docs/index.html'
```

### Get objects - URI

- Bucket name is the first path of the URI
- The object path are the leading values of path

The bucket name should have the 
`curl http://localhost/myBucker/My/path/of/object`

where:
- `myBucker` is the bucket name
- `My/path/of/object` is the object path (only supported objects - not paths)

<!-- ### Get objects - Headers (TODO)-->

## Contribute

Open an issue, or clone and open an merge request. =)
