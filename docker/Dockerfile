FROM gcr.io/gcp-etigcp-nprd-12855/pipelines/agent_redhat_base:v5

ARG IMAGE_VERSION

LABEL name='Trace Sampling Manager'
LABEL vendor='Cisco'
LABEL version=${IMAGE_VERSION}
LABEL release=${IMAGE_VERSION}
LABEL summary='Trace Sampling Manager Microservice'

RUN mkdir /app
COPY ./artifacts/trace-sampling-manager /app/

ENTRYPOINT ["/app/trace-sampling-manager"]