# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
ADD bin/manager .
ADD bin/webhook .
USER 65532:65532

ENTRYPOINT ["/manager"]
