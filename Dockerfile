FROM alpine
ADD empty /empty
ENTRYPOINT [ "/empty" ]
