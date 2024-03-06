# Go Microservice Project

Personal project for learning details about implementing a microservice project.
Doesn't currently do much.

## Status:

goweb: 
  - handles http requests and dispatches gRPC requests
  - login should be working, not much else yet
web:
  - original idea was a typescript web interface to explore multiple languages on the back end
  - wasn't thrilled with Hapi typing
  - grpc stuff wasn't as helped by typescript as I expected
  - abandoned for now
proto:
  - contains protobuff files and generated code for go and typescript
user:
  - user service, for authentication and user management
client:
  - react frontend
cli:
  - cli tooling to replace scripts I'm using
  - wip
revert / verify / deploy:
  - planning to use squitch for migrations, not working at this point

buf_update_protos.sh:
  - regenerate files based on protos in proto directory

dockerfiles:
  - docker config stuff goes here, currently only starts the user database