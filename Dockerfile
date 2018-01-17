FROM golang:1.9.2-stretch

WORKDIR /go/src/app

RUN mkdir -p /go/src/app/pem; \
	mkdir -p /go/src/app/log


ADD ./app /go/src/app

RUN go-wrapper download
RUN go-wrapper install 

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi

EXPOSE 8080