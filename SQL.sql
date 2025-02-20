CREATE TABLE Author (
    Author_ID INT PRIMARY KEY,
    Name VARCHAR(255),
    Surname VARCHAR(255),
    Country VARCHAR(255),
    Date_Of_Birth DATE
);

CREATE TABLE Loan (
    Loan_ID INT PRIMARY KEY,
    Book INT,
    Reader INT,
    Loan_Date DATE,
    Return_Date DATE,
    Status VARCHAR(255),
    FOREIGN KEY (Book) REFERENCES Book(Book_ID),  
    FOREIGN KEY (Reader) REFERENCES Reader(Reader_ID)
);

CREATE TABLE Genre (
    Genre_ID INT PRIMARY KEY,
    Name VARCHAR(255),
    Style VARCHAR(255)
);

CREATE TABLE Fine (
    Fine_ID INT PRIMARY KEY,
    Time_Holded INTEGER,
    Amount_Money INTEGER
);

CREATE TABLE Book (
    Book_ID INT PRIMARY KEY,
    Title VARCHAR(255),
    Genre INT,
    Author INT,
    Fine INT,
    Year_Release INTEGER,
    Language VARCHAR(255),
    Style VARCHAR(255),
    FOREIGN KEY (Genre) REFERENCES Genre(Genre_ID),
    FOREIGN KEY (Author) REFERENCES Author(Author_ID),
    FOREIGN KEY (Fine) REFERENCES Fine(Fine_ID)
);

CREATE TABLE Reader (
    Reader_ID INT PRIMARY KEY,
    Name VARCHAR(255),
    Surname VARCHAR(255),
    Home_Address VARCHAR(255)
);

CREATE TABLE Library (
    Library_ID INT PRIMARY KEY,
    Book INT,
    Reader INT,
    FOREIGN KEY (Book) REFERENCES Book(Book_ID),
    FOREIGN KEY (Reader) REFERENCES Reader(Reader_ID)
);

CREATE TABLE Publisher (
    Publisher_ID INT PRIMARY KEY,
    Name VARCHAR(255),
    Address VARCHAR(255),
    Contact_Number VARCHAR(255)
);



