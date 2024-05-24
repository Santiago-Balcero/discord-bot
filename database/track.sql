CREATE TABLE track (
    track_id SERIAL PRIMARY KEY,
    album_id INTEGER NOT NULL,
    spotify_id VARCHAR(250) UNIQUE NOT NULL,
    "name" VARCHAR(250) NOT NULL,
    "url" VARCHAR(250) NOT NULL,
    danceability DECIMAL(5,2) NOT NULL,
    energy DECIMAL(5,2) NOT NULL,
    acousticness DECIMAL(5,2) NOT NULL,
    loudness DECIMAL(5,2) NOT NULL,
    liveness DECIMAL(5,2) NOT NULL,
    instrumentalness DECIMAL(5,2) NOT NULL,
    tempo DECIMAL(5,2) NOT NULL,
    duration_ms INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (album_id) REFERENCES album
);

CREATE INDEX ix_track_name ON track("name");
CREATE INDEX ix_track_id ON track(spotify_id);
