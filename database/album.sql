CREATE TABLE album (
    album_id SERIAL PRIMARY KEY,
    artist_id INTEGER NOT NULL,
    spotify_id VARCHAR(250) UNIQUE NOT NULL,
    "name" VARCHAR(250) NOT NULL,
    "type" VARCHAR(250) NOT NULL,
    release_date VARCHAR(250) NOT NULL,
    "url" VARCHAR(250) NOT NULL,
    "image" VARCHAR(250) NOT NULL,
    max_danceability DECIMAL(5,2) NOT NULL,
    max_danceability_track VARCHAR(250) NOT NULL,
    max_energy DECIMAL(5,2) NOT NULL,
    max_energy_track VARCHAR(250) NOT NULL,
    tracks_count INTEGER NOT NULL,
    duration_ms INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (artist_id) REFERENCES artist
);

CREATE INDEX ix_album_name ON album("name");
CREATE INDEX ix_album_id ON album(spotify_id);
