FROM golang:1.9
MAINTAINER Olivier El Mekki <olivier@el-mekki.com>

RUN apt-get update && apt-get install -y \
	apt-transport-https \
	ca-certificates \
	curl \
  gnupg \
	--no-install-recommends \
	&& curl -sSL https://dl.google.com/linux/linux_signing_key.pub | apt-key add - \
	&& echo "deb [arch=amd64] https://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list \
	&& apt-get update && apt-get install -y \
	google-chrome-stable \
	--no-install-recommends \
	&& apt-get purge --auto-remove -y curl gnupg \
	&& rm -rf /var/lib/apt/lists/*

RUN useradd -m -G audio,video chromessr
RUN mkdir /cache && chown chromessr:chromessr /cache && chown -R chromessr:chromessr /go

ENV SSR_CACHE_PATH /cache

USER chromessr
RUN go get github.com/raff/godet
RUN mkdir -p /go/src/github.com/oelmekki/chromessr
ADD . /go/src/github.com/oelmekki/chromessr/

EXPOSE 3001

CMD [ "/bin/sh", "/go/src/github.com/oelmekki/chromessr/entrypoint.sh" ]
