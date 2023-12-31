FROM busybox:1.36.1 AS builder

WORKDIR /sql

COPY sql /lms/sql

RUN test -d /lms/sql/migrations && cp -r /lms/sql/migrations/. . || true
RUN test -d /lms/sql/seeders && cp -r /lms/sql/seeders/. . || true

FROM migrate/migrate:4 AS migrator

COPY --from=builder /sql /sql

CMD ["-database", "postgres://postgres:postgres@postgres:5432/lms?sslmode=disable", "-path", "/sql", "up"]

ENTRYPOINT ["migrate"]