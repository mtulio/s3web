# s3web

s3web can serve objects from AWS S3 bucket.

NOTE: the most description and the idea of the solution are in the begging of the implementation and may not tested enough to be broadly used. Contributions are open! ;)

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

### Get objects from Query string

```bash
curl 'localhost:8080/?bucket=myBucker&object=docs/index.html'
```

### Get objects from URI

- Bucket name is the first path of the URI
- The object path are the leading values of path

The bucket name should have the 
`curl http://localhost/myBucket/docs/index.html`

where:
- `myBucker` is the bucket name
- `docs/index.html` is the object path (only supported objects - not paths)

### Get objects from Headers (TODO)

- Host header

~~~
curl http://myBucker/My/path/of/object
curl -H 'Host: myBucker' http://localhost/docs/index.html
~~~

- Custom Header `BUCKET_NAME` and `BUCKET_OBJECT`

~~~
curl -H 'BUCKET_NAME: myBucker' -H 'BUCKET_OBJECT: docs/index.html' http://localhost
~~~

## Contribute

Open an issue, or clone and open an merge request. =)
