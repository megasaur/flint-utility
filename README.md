# Welcome to flint-utility

A web app for generating and collating statistics and insights to the flint ult footy fantasy league.
This is a hacktastic project, made with go and react by someone who has never used either of them.  Its just as much about learning these techs as it is about making something awesome!

## Process for adding new round data

1) Create a file called something like "formatter_feeder_rd_x.txt" in the scripts folder (split this out into years eventually)
2) Copy and paste the right blobs from the ult footy website
3) Check to make sure that disposal efficieny is not tied for any match. If it is, look up who the winner of it is and bump their disposal efficieny by the smallest practical value.  (ult footy is not showing all the digits and awards points even when it looks like a tie)
4) manipulate the array 'formatter_feeder_files' in the script 'formatter.py' for the desired target(s)
5) check and make sure all team names are covered by the dict 'coach_map' in 'formatter.py'
6) run `python3 formatter.py` script to preview the SQL and ensure the script runs
7) run `python3 formatter.py commit` to run the generated SQL in the database

## Run through docker

So long as you have docker and docker-compose, simply run:

   `docker-compose up --build`

## Stack info

Backend: go + echo

Frontend: react

Database: mysql
