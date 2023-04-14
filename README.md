# Simple Bank
<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#tech-stack">Tech Stack</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#run-project-locally">Run project locally</a></li>
      </ul>
    </li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Simple Bank is an API that enables the following features:

* Create a new user
* User can login to their account 
* User can perform CRUD (Create, Read, Update, Delete) actions to manage their account
* User can transfer money to another user
* Authenticate and authorize user



### Tech Stack

* Gin 
* Postgres + SQLC 
* PASETO 
* GitHub Action 
* Docker 
* AWS 



<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

Install tools:
* <a href="https://go.dev/">Golang</a>
* <a href="https://www.docker.com/products/docker-desktop/">Docker desktop</a>
  

### Run project locally

1. Clone the repo
   ```sh
   git clone https://github.com/iamjeremylim/simplebank.git
   ```
2. Change directory to project root folder
3. Run project 
   ```sh
   docker compose up
   ```   
4. End project 
   ```sh
   docker compose down
   ```

<p align="right">(<a href="#simple-bank">back to top</a>)</p>
