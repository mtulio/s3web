# s3web

s3web can serve files from AWS S3 objects.

It's common used when you wouldn't like to provide AWS Credentials to access S3 **and** the bucket is not exposed to the internet (Eg. internal web app to serve S3 data).

**WARNING**: This web application is useful to retrieve S3 data on an secure network that the users should not have AWS Credentials, please, understand the risks to expose this application to the internet.

## Authentication

The `s3web` can authenticate in S3 in three ways:
- the host application access the S3 through IAM Roles (instance roles)
- provide AWS credentials on bootstrap (TODO)
- send authentication through headers (TODO)
- send authentication through query strings (TODO, needed?)

## Usage

### Config

You can specify the Bucket, Object and the Region (optional) in three ways:
- URI
- Query string
- Headers

### Get objects - URI

- Bucket name is the first path of the URI
- The object path are the leading values of path

The bucket name should have the 
`curl http://localhost/myBucker/My/path/of/object`

where:
- `myBucker` is the bucket name
- `My/path/of/object` is the object path (only supported objects - not paths)

### Get objects - Query string (TODO)

### Get objects - Headers (TODO)

## Contribute

Open an issue, or clone and open an merge request. =)
