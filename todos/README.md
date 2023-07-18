CREATE TABLE todos (
id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
todo VARCHAR(320) NOT NULL,
completed BOOLEAN NOT NULL,
created_date TIMESTAMPTZ DEFAULT current_timestamp
);
