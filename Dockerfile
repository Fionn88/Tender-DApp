FROM golang
WORKDIR /
COPY fabric-asset-transfer-basic-server fabric-asset-transfer-basic-server
EXPOSE 8500
CMD /fabric-asset-transfer-basic-server
