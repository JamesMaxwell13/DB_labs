INSERT INTO tour (date, time, duration, price)
SELECT
    CURRENT_DATE + (floor(random() * 30) || ' days')::interval AS date,
    ('08:00'::time + (floor(random() * 8) || ' hours')::interval) AS time,
    ('01:00'::time + (floor(random() * 2) || ' hours')::interval) AS duration,
    (20.00 + (random() * 30))::numeric(10, 2) AS price
FROM generate_series(1, 27);