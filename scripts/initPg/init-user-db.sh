#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
        CREATE DATABASE hezzl;
        
        CREATE TABLE IF NOT EXISTS campaigns
(
    id serial 	PRIMARY KEY,
    name 		VARCHAR(100) NOT NULL UNIQUE
    );
CREATE INDEX idx_id ON campaigns(id);
INSERT INTO campaigns (name) VALUES ('First campaign');

CREATE TABLE IF NOT EXISTS items
(
    id serial 		PRIMARY KEY,
    campaign_id 	INTEGER NOT NULL REFERENCES campaigns(id),
    name 			VARCHAR(100) NOT NULL,
    description	 	VARCHAR(100),
    priority 		INTEGER DEFAULT(1),
    removed boolean DEFAULT(FALSE),
    created_at 		TIMESTAMP DEFAULT(CURRENT_TIMESTAMP)
    );

CREATE INDEX id_idx ON items(id);
CREATE INDEX campaign_id_idx ON items(campaign_id);
CREATE INDEX name_idx ON items USING HASH (name);
        
    
EOSQL 
