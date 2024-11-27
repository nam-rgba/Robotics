-- Backup and delete all rows from all tables

-- Backup data
CREATE TABLE candidate_backup AS TABLE candidate;
CREATE TABLE coach_backup AS TABLE coach;
CREATE TABLE team_backup AS TABLE team;
CREATE TABLE challenges_backup AS TABLE challenges;
CREATE TABLE competition_backup AS TABLE competition;
CREATE TABLE organization_backup AS TABLE organization;
CREATE TABLE team_candidate_backup AS TABLE team_candidate;

-- Delete all rows
DELETE FROM candidate;
DELETE FROM coach;
DELETE FROM team;
DELETE FROM challenges;
DELETE FROM competition;
DELETE FROM organization;
DELETE FROM team_candidate;
