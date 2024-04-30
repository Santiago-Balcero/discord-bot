CREATE TABLE album (
    db_id SERIAL PRIMARY KEY,
    id VARCHAR(150) UNIQUE NOT NULL,
    "name" VARCHAR(150) NOT NULL,
    "type" VARCHAR(150) NOT NULL,
    release_date VARCHAR(150) NOT NULL,
    "url" VARCHAR(150) NOT NULL,
    "image" VARCHAR(150) NOT NULL,
    max_danceability DECIMAL(5,2) NOT NULL,
    max_danceability_track VARCHAR(150) NOT NULL,
    max_energy DECIMAL(5,2) NOT NULL,
    max_energy_track VARCHAR(150) NOT NULL,
    tracks_count INTEGER NOT NULL,
    duration_ms INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE INDEX ix_album_name ON album("name");
CREATE INDEX ix_album_id ON album(id);
