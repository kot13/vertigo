FROM scratch

ENV PORT 8000
ENV LOG_LEVEL "debug"

EXPOSE $PORT

COPY vertigo /
CMD ["/vertigo"]