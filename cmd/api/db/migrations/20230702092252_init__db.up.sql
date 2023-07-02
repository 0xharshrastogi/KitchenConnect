CREATE TABLE [address]
(
    id INT IDENTITY(1,1) PRIMARY KEY,
    [street] NVARCHAR(255),
    [city] NVARCHAR(255),
    [state] NVARCHAR(255),
    [zip_code] NVARCHAR(10),
    [country_code] NVARCHAR(5)
);

CREATE TABLE [user]
(
    [id] INT IDENTITY(1,1) PRIMARY KEY,
    [first_name] NVARCHAR(255) NOT NULL,
    [last_name] NVARCHAR(255) NOT NULL,
    [email] NVARCHAR(255) NOT NULL UNIQUE,
    [password] NVARCHAR(255) NOT NULL,
    [password_salt] NVARCHAR(255) NOT NULL,
    [address_id] INT NOT NULL,
    CONSTRAINT [FK_address_entity] FOREIGN KEY (address_id) REFERENCES [address](id) ON DELETE CASCADE
);
