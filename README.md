# SSM Restarter

Simple agent to listen on a TCP port an accept pre-defined command to restart amazon-ssm-agent.

Intended for use within a private network due to lack of any authentication mechanism.

## Building binaries for distribution

```sh
docker build --output=dist --target=binaries .

ls dist
```

## Building and running in docker (for testing)

```sh
docker build -t klevo/ssm_restart_agent .

docker run --rm --name ssm_restart_agent -p 9009:9009 klevo/ssm_restart_agent

# Sending a message (locally)
echo "restart_ssm_agent" | nc localhost 9009

# Sending with timeout
echo "restart_ssm_agent" | nc -w 3 localhost 9009 &> /dev/null
```

## Installing on Linux instance

```sh
wget https://github.com/envirobly/ssm-restart/releases/download/v0.1/ssm_restart_agent
chmod +x ssm_restart_agent
./ssm_restart_agent
```
