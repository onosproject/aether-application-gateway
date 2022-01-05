# aether-application-gateway

The Aether Application Gateway (AG) provides a unified API for both control and metrics of the Aether platform. 
The AG is device and application focused, enabling the application developer to implement intelligent control 
of Aether 4G/5G connectivity services, while abstracting some low-level intricacies of 4G/5G.

### Capabilities

* Application Enablement
* Device Management
* Telemetry & Analytics
* Slice Management

### Project Structure
    .
    ├── api                     # OpenAPI specs
    ├── build                   # dockerfiles
    ├── cmd                     # application entrypoints
    ├── docs                    # design and developer documents
    ├── infra                   # kind configs & kustomize manifests
    ├── internal                # private application and library code
    ├── pkg                     # library code that may be used by external applications
    │
    └── makefile                 # commands for developing this project
