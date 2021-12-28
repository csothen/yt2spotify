# yt2spotify <!-- omit in toc --> 

- [Description](#description)
- [Getting Started](#getting-started)
  - [System requirements](#system-requirements)
  - [Environemnt variables](#environemnt-variables)
  - [Running the application](#running-the-application)
  - [Accessing the Database](#accessing-the-database)

# Description

yt2spotify is a platform to convert your YouTube playlists into Spotify playlists

# Getting Started

## System requirements

In order to run the application the following are necessary:

- [Docker](https://docs.docker.com/engine/install/) / [Docker Desktop](https://docs.docker.com/desktop/#download-and-install)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)

## Environemnt variables

In order to run the project make sure the following environment variables are set:
- `MYSQL_DB_NAME`: Name of the database which you will connect to
- `MYSQL_DB_USER`: Name of the DB user
- `MYSQL_DB_PASSWORD`: Password of the DB user
- `MYSQL_DB_ROOT_PASSWORD`: Password for the root DB user

This can be done by creating a file with the name `.env` where you will define each one of these

## Running the application

Once the environment variables are properly set up you can simply run `make start`

## Accessing the Database

In the occurrence you want to check the database you can run `make db-cli` in order to gain access to the MySQL CLI and interact with it
