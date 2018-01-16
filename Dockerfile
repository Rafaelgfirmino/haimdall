FROM golang:1.9.2-stretch

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/app
RUN mkdir -p /go/src/app/pem
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