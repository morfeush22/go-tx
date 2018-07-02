# ---------------------------
FROM elixir:1.5-alpine as builder

RUN mix local.hex --force && \
    mix local.rebar --force

WORKDIR /app
ENV MIX_ENV prod
COPY . .
RUN mix deps.get
RUN mix release --env=$MIX_ENV

# --------------------
FROM elixir:1.5-alpine

WORKDIR /app
COPY --from=builder /app/_build/prod/rel/scheduler .

RUN apk add --update bash && rm -rf /var/cache/apk/*

EXPOSE 8080

CMD ["./bin/scheduler", "foreground"]
