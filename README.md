# Email Search System

## Description

This is a **full-stack** email search and management system, designed to be fast, efficient, and user-friendly. The project consists of three main parts:

- **Backend**: Developed in **Go**. This server handles search requests and manages the system logic, using **ZincSearch** as the search engine to store and retrieve email data.
  
- **Frontend**: Developed in **Vue.js**. The frontend provides an intuitive user interface that allows users to perform searches, view results, and easily manage emails.

- **Database**: **ZincSearch** is used as the database and search engine to index emails and make searches fast and efficient.

This system allows fast email searches in large volumes of data, using a powerful search engine to filter, index, and retrieve information efficiently.

## Technologies Used

- **Backend**: Go (Golang)
- **Frontend**: Vue.js
- **Database**: ZincSearch

## Features

- **Fast email search**: Using **ZincSearch** to index and query emails efficiently.
- **User-friendly interface**: A **Vue.js** based interface that facilitates interaction with the system.
- **Scalability**: The backend is designed to handle a high volume of data and queries.

## Installation

Follow the steps below to get the project up and running on your local machine.

### 1. Clone the repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/lanorrron/ZincEmail.git
cd ZincEmail
```

### 2. Backend Setup

The backend consists of two main components:

- **Navigate backend directory**:
  ```bash
   cd backend   
- **Indexer Program**: This component is responsible for indexing the email files.
  ```bash
   go run ./cmd/indexer/main.go   
- **Server**: This is the part of the backend that listens for incoming requests and interacts with the frontend.
   ```bash
   go run ./cmd/server/main.go
   
### 3. Frontend Setup

The frontend is developed in Vue.js. To set it up and run the application, follow these steps:
1. **Navigate to the frontend directory**:
   ```bash
   cd frontend
2. **Install dependencies using npm**:
    ```bash
    npm install
3. **Run the frontend**
    ```bash
   npm run dev
### 4. ZincSearch Database Setup

1. **Visit ZincSearch**:  
   Go to the official page [https://zincsearch-docs.zinc.dev/](https://zincsearch-docs.zinc.dev/) and follow the installation instructions.

2. **Run ZincSearch**:  
   After installation, run ZincSearch.


3. **Download the Email Data for Indexing**  
   To index emails, you need the data. Download the sample email data to use in this project:

   - **Enron Corp Data**: [Download Email Data](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)


 ### 5. Indexing Optimization

The indexing process is optimized using **goroutines** for parallel processing and **batch processing** for sending data to ZincSearch. The process works as follows:

1. **Executing Goroutines for Parallel Document Processing**:
   - The first step is executing **goroutines**. These goroutines process documents concurrently, which speeds up the document reading and parsing process.
   - The number of **goroutines** is configurable, allowing you to adjust how many documents can be processed in parallel.

2. **Dividing Document Processing Tasks Among Goroutines**:
   - Based on the number of goroutines specified, the total number of documents is divided into equal tasks for each goroutine. Each goroutine handles a specific chunk of the documents.

3. **Indexing to ZincSearch (After Goroutines Complete)**:
   - Once the goroutines have finished processing all the documents, the next step is to send the processed data to **ZincSearch**.
   - The documents are grouped into **batches of 1000** (by default) to optimize the upload process.
   - This batch size is **configurable**, so you can adjust it to match your system's capacity.

By using goroutines for parallel document processing and dividing the indexed data into smaller batches for ZincSearch, the system ensures fast and efficient indexing, minimizing the load on your system.
