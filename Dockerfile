# FROM golang
# WORKDIR /app
# ADD . .
# CMD ["go","run","main.go"]
FROM mcr.microsoft.com/dotnet/aspnet:6.0.25-jammy
RUN apt-get update \
    && apt-get install -y --no-install-recommends libgdiplus \
        openslide-tools \
        mtr \
        tcpdump \
        vim-tiny \
        curl \
        telnet \
        net-tools \
    && rm -rf /etc/localtime \
    && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime 
RUN apt-get clean \
    && rm -rf /var/cache/apt/archives/* \
    && rm -rf /var/lib/apt/lists/* \
