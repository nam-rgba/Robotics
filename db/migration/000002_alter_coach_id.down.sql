ALTER TABLE candidate
DROP CONSTRAINT candidate_coach_id_fkey;

ALTER TABLE candidate
ALTER COLUMN coach_id SET DATA TYPE int;

ALTER TABLE candidate
ADD CONSTRAINT candidate_coach_id_fkey 
FOREIGN KEY (coach_id) REFERENCES coach(coach_id) ON DELETE CASCADE