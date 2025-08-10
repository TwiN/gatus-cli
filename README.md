# gatus-cli
This is a CLI tool for interacting with deployed Gatus status pages as well as Gatus.io

## Usage
> **NOTE:** If you're looking to interact with Gatus programmatically, see [TwiN/gatus-sdk](https://github.com/TwiN/gatus-sdk)

### Status Pages (Gatus.io only)
For Gatus.io hosted service, you need to set your API key:
```console
export GATUS_CLI_API_KEY=your-api-key-here
```

```console
# Retrieve a status page by ID
gatus-cli status-page get --status-page-id 12345
```

### Endpoints
All endpoint commands require the `--url` flag to specify your Gatus server URL.

#### Get Endpoint Status
```console
# Get status of all endpoints
gatus-cli endpoint status all --url https://status.example.com

# Get status by endpoint key
gatus-cli endpoint status get --url https://status.example.com --key "group_endpoint"

# Get status by group and name
gatus-cli endpoint status get --url https://status.example.com --group "web" --name "frontend"
```

#### Get Uptime Information
```console
# Get uptime percentage for an endpoint
gatus-cli endpoint uptime --url https://status.example.com --key "group_endpoint" --duration "7d"
```

#### Get Response Times
```console
# Get response time statistics
gatus-cli endpoint response-times --url https://status.example.com --key "group_endpoint" --duration "24h"
```

#### Generate Badge URLs
```console
# Generate health badge URL
gatus-cli endpoint badge health --url https://status.example.com --key "group_endpoint"

# Generate uptime badge URL
gatus-cli endpoint badge uptime --url https://status.example.com --key "group_endpoint" --duration "7d"

# Generate response time badge URL
gatus-cli endpoint badge response-time --url https://status.example.com --key "group_endpoint" --duration "24h"
```

### Duration Format
Duration values must be one of the following Gatus API supported formats:
- `1h` - 1 hour
- `24h` - 24 hours (1 day)
- `7d` - 7 days (1 week)
- `30d` - 30 days (1 month)
