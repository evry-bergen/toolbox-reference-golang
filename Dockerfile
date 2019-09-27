FROM gcr.io/distroless/static
COPY ./main .
ENTRYPOINT ["/main"]