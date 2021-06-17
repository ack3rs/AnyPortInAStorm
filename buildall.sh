#/bin/sh
docker build . -t portdomainserviceserver -f PortDomainServiceServer/Dockerfile
docker build . -t client -f Client/Dockerfile
