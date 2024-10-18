package api

import db "./sqlc"

type Server struct {
	store *db.Store
}
