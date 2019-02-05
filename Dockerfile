FROM ipfs/go-ipfs:v0.4.18

RUN ipfs init --profile=badgerds,lowpower,local-discovery && \
    ipfs config --json Datastore.NoSync true && \
    ipfs config Reprovider.Interval "0" && \
    ipfs config --json Experimental.FilestoreEnabled true

ENTRYPOINT ["/sbin/tini", "--", "/usr/local/bin/start_ipfs"]
CMD ["daemon", "--offline"]
