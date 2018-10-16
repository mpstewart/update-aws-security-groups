# update-aws-security-groups

## Description

`update-aws-security-groups` is a script, meant to be run as a cron, that will update Amazon EC2 security rules to accept certain incoming connections based on a DDNS hostname. The script checks the hostname, compares it to a stored IP address, then updates the security rules if necessary.

## Getting started

There's a lot of setup for this project, most of it manual. Sorry about that.

It is assumed you've set up the credentials and permissions on the AWS side, and you have a `~/.aws/credentials` file.

1. Clone the repo
```bash
git clone https://www.github.com/mpstewart/update-aws-security-group
cd update-aws-security-groups
```
2. Build the binary and install the script
```bash
go install
sudo cp $GOPATH/bin/update-aws-security-groups /usr/local/bin/update-aws-security-groups
```
3. Create a file to hold the config
```bash
sudo mkdir /etc/update-aws-security-groups
sudo vim /etc/update-aws-security-groups/config
```
with something like the following contents:
```json
{
  "awsProfile": "default",
  "hostname": "myhostname.ddns.net",
  "homeIP": "192.168.1.1/32",
  "groupID": "sg-123456",
  "region": "us-west-1",
  "ports": [
    {
      "port": 22,
      "protocol": "tcp",
      "description": "SSH"
    }
  ]
}
```

4. That's about it. Simply save the config file, then update your crontab to run at the preferred frequency, setting the `PATH` variable to include `/usr/local/bin/` and you're all set.

### Explanation of `config` keys:

#### `awsProfile`
This is used as AWS_PROFILE

#### `hostname`
The DDNS hostname used to check the IP Address.

#### `homeIP`
Used to persist the IP. You can populate this once if you want, but it will be tracked automatically afterwards.

#### `groupID`
Name of the Group ID you wish to adjust the rules for.

#### `region`
Your EC2 region.

#### `ports`
The main point of extension of the script. For each port you wish to open, you will need to indicate the port number, via `port`, the `protocol`, and the `description` that will appear in the EC2 management console.

## Why?

Wanted to make something with Go. This got made.
