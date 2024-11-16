CREATE TABLE IF NOT EXISTS taro_cards (
    id SERIAL PRIMARY KEY,
    subcategory int,
    direct_meaning TEXT,
    reverse_meaning TEXT
);