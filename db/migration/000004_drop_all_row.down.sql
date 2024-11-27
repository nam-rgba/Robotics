-- Restore data from backups
INSERT INTO candidate SELECT * FROM candidate_backup;
INSERT INTO coach SELECT * FROM coach_backup;
INSERT INTO team SELECT * FROM team_backup;
INSERT INTO challenges SELECT * FROM challenges_backup;
INSERT INTO competition SELECT * FROM competition_backup;
INSERT INTO organization SELECT * FROM organization_backup;
INSERT INTO team_candidate SELECT * FROM team_candidate_backup;
