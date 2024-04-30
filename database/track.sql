CREATE TABLE track (
    db_id SERIAL PRIMARY KEY,
    id VARCHAR(150) UNIQUE NOT NULL,
    "name" VARCHAR(150) NOT NULL,
    "url" VARCHAR(150) NOT NULL,
    danceability DECIMAL(5,2) NOT NULL,
    energy DECIMAL(5,2) NOT NULL,
    acousticness DECIMAL(5,2) NOT NULL,
    loudness DECIMAL(5,2) NOT NULL,
    liveness DECIMAL(5,2) NOT NULL,
    instrumentalness DECIMAL(5,2) NOT NULL,
    tempo DECIMAL(5,2) NOT NULL,
    duration_ms INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE INDEX ix_track_name ON track("name");
CREATE INDEX ix_track_id ON track(id);
