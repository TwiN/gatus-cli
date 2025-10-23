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

Duration values for badges must be one of the following:
- `1h` - 1 hour
- `24h` - 24 hours (1 day)
- `7d` - 7 days (1 week)
- `30d` - 30 days (1 month)

### Suites
Suites are collections of sequential endpoint checks. All suite commands require the `--url` flag to specify your Gatus server URL.

#### Get Suite Status
```console
# Get status of all suites
gatus-cli suite status all --url https://status.example.com

# Get status by suite key
gatus-cli suite status get --url https://status.example.com --key "_check-authentication"

# Get status by group and name (group is optional)
gatus-cli suite status get --url https://status.example.com --group "monitoring" --name "health-checks"

# Get status by name only
gatus-cli suite status get --url https://status.example.com --name "health-checks"
```

### External Endpoints
External endpoints allow push-based monitoring where external systems report their health status to Gatus.

#### Push Result
```console
# Push a successful health check
gatus-cli external-endpoint push --url https://status.example.com --key "group_endpoint" --token "secret-token" --success

# Push a failed health check with error message
gatus-cli external-endpoint push --url https://status.example.com --key "group_endpoint" --token "secret-token" --error "Connection timeout"

# Push with duration
gatus-cli external-endpoint push --url https://status.example.com --key "group_endpoint" --token "secret-token" --success --duration "2s"
```
