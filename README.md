# SSM Restarter

Simple agent to listen on a TCP port an accept pre-defined command to restart amazon-ssm-agent.

## Building

```sh
docker build -t klevo/ssm_restart_agent .
```

## Running

```sh
docker run --rm --name ssm_restart_agent -p 9009:9009 klevo/ssm_restart_agent

# Sending a message (locally)
echo "restart_ssm_agent" | nc localhost 9009
```
