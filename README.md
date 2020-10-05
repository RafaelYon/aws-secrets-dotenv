# aws-secrets-dotenv
A simple tool to obtain the secrets stored in AWS Secrets Manager in dotenv format

## Installing

### Requirements
 - Golang `1.11` or higher
 - AWS access key ID and secret access key stored in your machine (`ENV`s or `~/.aws/credentials`)

### Getting

```
$ go get github.com/RafaelYon/aws-secrets-dotenv
```

## Using

The entry point of the tool is the executable `aws-secrets-dotenv`. With this it is possible to view all commands and options by executing:

```
$ aws-secrets-dotenv
```

It is important to note that **these examples assume that the `$GOPATH/bin` directory has been added to your environment's `$PATH`.**

### Retrieving the secrets in dotenv format

```
$ aws-secrets-dotenv get SecretId [--version-id VersionId] [--version-stage VersionStage] [{-f --file} .env] [--aws-region us-east-2]
```

To retrieve a secret all you need to specify is the AWS SecretId which will generate a `.env` in the current folder:

```
$ ls -la
drwxr-xr-x 4 user user  4096 out  5 10:53 .
drwxr-xr-x 9 user user  4096 out  2 16:43 ..
drwxr-xr-x 8 user user  4096 out  5 10:53 .git
```

```
$ aws-secrets-dotenv get example-secret
```

Silence is golden.

```
$ ls -la
drwxr-xr-x 4 user user  4096 out  5 10:53 .
drwxr-xr-x 9 user user  4096 out  2 16:43 ..
-rw-r--r-- 1 user user    17 out  5 10:53 .env
drwxr-xr-x 8 user user  4096 out  5 10:53 .git
```