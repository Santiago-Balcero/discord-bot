CREATE TABLE artist (
    artist_id SERIAL PRIMARY KEY,
    spotify_id VARCHAR(250) UNIQUE NOT NULL,
    "name" VARCHAR(250) NOT NULL,
    popularity INTEGER NOT NULL,
    genres TEXT ARRAY NOT NULL,
    "url" VARCHAR(250) NOT NULL,
    followers INTEGER NOT NULL,
    "image" VARCHAR(250) NOT NULL,
    max_danceability DECIMAL(5,2) NOT NULL,
    max_danceability_track VARCHAR(250) NOT NULL,
    max_energy DECIMAL(5,2) NOT NULL,
    max_energy_track VARCHAR(250) NOT NULL,
    albums_count INTEGER NOT NULL,
    singles_count INTEGER NOT NULL,
    compilations_count INTEGER NOT NULL,
    tracks_count INTEGER NOT NULL,
    duration_ms INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE INDEX ix_artist_name ON artist("name");
CREATE INDEX ix_artist_id ON artist(spotify_id);
