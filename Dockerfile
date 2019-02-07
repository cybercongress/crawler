FROM ipfs/go-ipfs:v0.4.18

RUN ipfs init --profile=badgerds && \
    ipfs config --json Datastore.NoSync true && \
    ipfs config --json Experimental.ShardingEnabled true && \
    ipfs config Reprovider.Interval "0"

ENTRYPOINT ["/sbin/tini", "--", "/usr/local/bin/start_ipfs"]
CMD ["daemon", "--routing=dhtclient"]
