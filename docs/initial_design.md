# Initial Design

Proof of concepts on how plugin should be built based on source and sink components.

## Basic Components
- source-any
- sink-any

## ETL

### Extract (Importer)
Data from outside warehouse to warehouse. Components: source-X, sink-Y

### Transform (Executor)
Data transformation within warehouse. It can also be used as simple one time execution.

### Load (Exporter)
Data from warehouse to outside warehouse. Components: source-X, sink-Y

## Connecting between components

**source components**:
- authenication to access data from source
- data size control
- data parser (something -> json)
- instrumentation

**sink components**:
- authentication to push data to destination
- data size control
- data parser (json -> something)
- instrumentation
