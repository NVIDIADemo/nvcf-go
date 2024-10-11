# Shared Params Types

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedPartyDTOParam">AuthorizedPartyDTOParam</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#HealthDTOParam">HealthDTOParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedParties">AuthorizedParties</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedPartyDTO">AuthorizedPartyDTO</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#CreateFunctionResponse">CreateFunctionResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#Function">Function</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#FunctionDTO">FunctionDTO</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#FunctionsResponse">FunctionsResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#HealthDTO">HealthDTO</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#QueuesResponse">QueuesResponse</a>

# UserSecretManagement

## Functions

### Versions

Methods:

- <code title="put /v2/nvcf/secrets/functions/{functionId}/versions/{versionId}">client.UserSecretManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#UserSecretManagementFunctionVersionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#UserSecretManagementFunctionVersionUpdateParams">UserSecretManagementFunctionVersionUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# FunctionManagement

## Functions

### Versions

Methods:

- <code title="get /v2/nvcf/functions/{functionId}/versions/{functionVersionId}">client.FunctionManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionManagementFunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v2/nvcf/metadata/functions/{functionId}/versions/{functionVersionId}">client.FunctionManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionManagementFunctionVersionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionManagementFunctionVersionUpdateParams">FunctionManagementFunctionVersionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/functions/{functionId}/versions/{functionVersionId}">client.FunctionManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionManagementFunctionVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### IDs

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionManagementFunctionIDGetAllResponse">FunctionManagementFunctionIDGetAllResponse</a>

Methods:

- <code title="get /v2/nvcf/functions/ids">client.FunctionManagement.Functions.IDs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionManagementFunctionIDService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionManagementFunctionIDGetAllParams">FunctionManagementFunctionIDGetAllParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionManagementFunctionIDGetAllResponse">FunctionManagementFunctionIDGetAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FunctionDeployment

## Functions

### Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#DeploymentResponse">DeploymentResponse</a>

Methods:

- <code title="post /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionDeploymentFunctionVersionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionDeploymentFunctionVersionNewParams">FunctionDeploymentFunctionVersionNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionDeploymentFunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionDeploymentFunctionVersionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionDeploymentFunctionVersionUpdateParams">FunctionDeploymentFunctionVersionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionDeploymentFunctionVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionDeploymentFunctionVersionDeleteParams">FunctionDeploymentFunctionVersionDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FunctionInvocation

## Functions

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionInvocationFunctionInvokeResponse">FunctionInvocationFunctionInvokeResponse</a>

Methods:

- <code title="post /v2/nvcf/pexec/functions/{functionId}">client.FunctionInvocation.Functions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionInvocationFunctionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionInvocationFunctionInvokeParams">FunctionInvocationFunctionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionInvocationFunctionInvokeResponse">FunctionInvocationFunctionInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionInvocationFunctionVersionInvokeResponse">FunctionInvocationFunctionVersionInvokeResponse</a>

Methods:

- <code title="post /v2/nvcf/pexec/functions/{functionId}/versions/{versionId}">client.FunctionInvocation.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionInvocationFunctionVersionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionInvocationFunctionVersionInvokeParams">FunctionInvocationFunctionVersionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionInvocationFunctionVersionInvokeResponse">FunctionInvocationFunctionVersionInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# EnvelopeFunctionInvocation

## Functions

Methods:

- <code title="post /v2/nvcf/exec/functions/{functionId}">client.EnvelopeFunctionInvocation.Functions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#EnvelopeFunctionInvocationFunctionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#EnvelopeFunctionInvocationFunctionInvokeParams">EnvelopeFunctionInvocationFunctionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="post /v2/nvcf/exec/functions/{functionId}/versions/{versionId}">client.EnvelopeFunctionInvocation.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#EnvelopeFunctionInvocationFunctionVersionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#EnvelopeFunctionInvocationFunctionVersionInvokeParams">EnvelopeFunctionInvocationFunctionVersionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Exec

### Status

Methods:

- <code title="get /v2/nvcf/exec/status/{requestId}">client.EnvelopeFunctionInvocation.Exec.Status.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#EnvelopeFunctionInvocationExecStatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Functions

Methods:

- <code title="post /v2/nvcf/functions">client.Functions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionNewParams">FunctionNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#FunctionDTO">FunctionDTO</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/functions">client.Functions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionGetAllParams">FunctionGetAllParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#FunctionsResponse">FunctionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Versions

Methods:

- <code title="post /v2/nvcf/functions/{functionId}/versions">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionVersionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionVersionNewParams">FunctionVersionNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#CreateFunctionResponse">CreateFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/functions/{functionId}/versions">client.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#FunctionVersionService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#FunctionsResponse">FunctionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Authorizations

## Functions

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#ListAuthorizedPartiesResponse">ListAuthorizedPartiesResponse</a>

Methods:

- <code title="delete /v2/nvcf/authorizations/functions/{functionId}">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedParties">AuthorizedParties</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/add">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionAddParams">AuthorizationFunctionAddParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedParties">AuthorizedParties</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/nvcf/authorizations/functions/{functionId}">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionAuthorizeParams">AuthorizationFunctionAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedParties">AuthorizedParties</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/remove">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionRemoveParams">AuthorizationFunctionRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedParties">AuthorizedParties</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/authorizations/functions/{functionId}">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#ListAuthorizedPartiesResponse">ListAuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="get /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedParties">AuthorizedParties</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedParties">AuthorizedParties</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}/add">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionVersionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionVersionAddParams">AuthorizationFunctionVersionAddParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedParties">AuthorizedParties</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionVersionService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionVersionAuthorizeParams">AuthorizationFunctionVersionAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedParties">AuthorizedParties</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}/remove">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionVersionService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AuthorizationFunctionVersionRemoveParams">AuthorizationFunctionVersionRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#AuthorizedParties">AuthorizedParties</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#CreateAssetResponse">CreateAssetResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#ListAssetsResponse">ListAssetsResponse</a>

Methods:

- <code title="post /v2/nvcf/assets">client.Assets.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AssetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AssetNewParams">AssetNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#CreateAssetResponse">CreateAssetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/assets">client.Assets.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AssetService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#ListAssetsResponse">ListAssetsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Queues

## Functions

Methods:

- <code title="get /v2/nvcf/queues/functions/{functionId}">client.Queues.Functions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#QueueFunctionService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#QueuesResponse">QueuesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="get /v2/nvcf/queues/functions/{functionId}/versions/{versionId}">client.Queues.Functions.Versions.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#QueueFunctionVersionService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go/shared#QueuesResponse">QueuesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Position

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#GetPositionInQueueResponse">GetPositionInQueueResponse</a>

Methods:

- <code title="get /v2/nvcf/queues/{requestId}/position">client.Queues.Position.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#QueuePositionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#GetPositionInQueueResponse">GetPositionInQueueResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Pexec

## Status

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#PexecStatusGetResponse">PexecStatusGetResponse</a>

Methods:

- <code title="get /v2/nvcf/pexec/status/{requestId}">client.Pexec.Status.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#PexecStatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#PexecStatusGetParams">PexecStatusGetParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#PexecStatusGetResponse">PexecStatusGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ClusterGroupsAndGPUs

## ClusterGroups

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#ClusterGroupsResponse">ClusterGroupsResponse</a>

Methods:

- <code title="get /v2/nvcf/clusterGroups">client.ClusterGroupsAndGPUs.ClusterGroups.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#ClusterGroupsAndGPUClusterGroupService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#ClusterGroupsResponse">ClusterGroupsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ClientManagementForNVIDIASuperAdmins

## Clients

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#ClientManagementForNVIDIASuperAdminClientGetResponse">ClientManagementForNVIDIASuperAdminClientGetResponse</a>

Methods:

- <code title="get /v2/nvcf/clients/{clientId}">client.ClientManagementForNVIDIASuperAdmins.Clients.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#ClientManagementForNVIDIASuperAdminClientService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#ClientManagementForNVIDIASuperAdminClientGetResponse">ClientManagementForNVIDIASuperAdminClientGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AssetManagement

## Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AssetResponse">AssetResponse</a>

Methods:

- <code title="get /v2/nvcf/assets/{assetId}">client.AssetManagement.Assets.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AssetManagementAssetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go">nvcf</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AssetResponse">AssetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/assets/{assetId}">client.AssetManagement.Assets.<a href="https://pkg.go.dev/github.com/NVIDIADemo/nvcf-go#AssetManagementAssetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
