# media-db

🎥 🎶

`media-db` is a CLI that uses AWS S3 as a database to keep track of movies, music, and other media. It's a useful way to record media you've consumed over time and query it for reference in the future.

Currently supported media types are movies and music. Other types to potentially add in the future include books, visual art, and theater.

## Prerequisites

* A recent version of [Go](https://golang.org/). This project was primarily coded and tested against Go 1.16.
* An AWS account.
* An AWS IAM user with an AWS CLI access key, secret key, and S3 permissions. Refer to [Configuration and credential file settings](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) for further information.
* An AWS S3 bucket to use for the backend database. A bucket can easily be created using the AWS CLI (e.g. `aws s3 mb s3://<bucket_name>`).

## Installation

* Run `go get github.com/alexpcook/media-db` to download the source code.
* Inside the source directory, compile an executable using `go install`.
* Ensure that the executable is in your PATH environment variable (e.g. on Linux `export PATH=$PATH:$(go env GOPATH)/bin`).
* (_Optional_) Create a short alias for the `media-db` command (e.g. on Linux `alias mdb=media-db`).

## Setup

* Run `media-db setup` to configure the AWS profile, region, and S3 bucket name connection settings.
  * This saves a configuration file to $HOME/.mediadb/config. The default configuration path can be overridden by setting the environment variable `MEDIA_DB_CONFIG_FILE`.

## Usage

* There are four main commands for interacting with the database.
  * `create` - Creates entries in the database. The required flags for creating objects vary depending on the type of media entry being created (e.g. movie vs. music).
  * `read` - Reads entries from the database. `media-db read` reads all entries. It's also possible to filter by media type and id (e.g. `media-db read music` and `media-db read movie -id=<id>` respectively).
  * `update` - Updates entries in the database. In addition to all flags required to create the media type, the `id` of the entry to update is also a required flag.
  * `delete` - Deletes entries from the database. The `id` flag is required.

## Credits

This code makes use of APIs in:

* [AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2/) for interacting programmatically between Go and AWS S3.
* [uuid](https://github.com/google/uuid) for generating unique IDs for the S3 object storage keys.
