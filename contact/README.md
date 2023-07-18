CREATE TABLE contacts (
id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
email VARCHAR(320) NOT NULL,
message TEXT,
created_date TIMESTAMPTZ DEFAULT current_timestamp
);
