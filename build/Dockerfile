FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/ssl-proxy-operator \
    USER_UID=1001 \
    USER_NAME=ssl-proxy-operator

# install operator binary
COPY build/_output/bin/ssl-proxy-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
