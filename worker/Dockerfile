FROM scratch as init
ADD . /app/

FROM alpine
COPY --from=init /app /app
WORKDIR /app
