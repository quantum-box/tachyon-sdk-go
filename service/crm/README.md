go test ./service/crm/... -count=1

grpcurl -plaintext -d '{"aggregationName": "test"}' -proto ./proto/crm.proto -rpc-header 'authorization: Bearer some-auth-token' localhost:50051 crm.CrmApi.Create