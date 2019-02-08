FROM redis:4.0.11

ENV REDIS_PASSWORD yLRt14UsRSnV4gC5

CMD ["sh", "-c", "exec redis-server --requirepass \"$REDIS_PASSWORD\""]
