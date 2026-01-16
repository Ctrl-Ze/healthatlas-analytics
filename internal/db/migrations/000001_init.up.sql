-- Enable TimescaleDB
CREATE EXTENSION IF NOT EXISTS timescaledb;

-- Metrics table (hypertable)
CREATE TABLE IF NOT EXISTS blood_test_metrics (
    time            TIMESTAMPTZ NOT NULL,
    user_ref        UUID        NOT NULL,
    analyte_code    TEXT        NOT NULL,
    value           DOUBLE PRECISION,
    zscore          DOUBLE PRECISION,
    slope_30d       DOUBLE PRECISION,
    slope_90d       DOUBLE PRECISION,
    anomaly_score   DOUBLE PRECISION,
    computed_at     TIMESTAMPTZ NOT NULL DEFAULT now(),

    PRIMARY KEY (user_ref, analyte_code, time)
);

SELECT create_hypertable('blood_test_metrics', 'time', if_not_exists => TRUE);

-- Trends table
CREATE TABLE IF NOT EXISTS blood_test_trends (
    user_ref        UUID        NOT NULL,
    analyte_code    TEXT        NOT NULL,
    last_value      DOUBLE PRECISION,
    prev_value      DOUBLE PRECISION,
    delta           DOUBLE PRECISION,
    slope_30d       DOUBLE PRECISION,
    slope_90d       DOUBLE PRECISION,
    anomaly_score   DOUBLE PRECISION,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),

    PRIMARY KEY (user_ref, analyte_code)
);
