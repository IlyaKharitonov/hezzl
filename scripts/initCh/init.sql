create table IF NOT EXISTS logs
(
    time    DateTime default now(),
    message String
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;
