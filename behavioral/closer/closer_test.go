package closer

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestCloser(t *testing.T) {
	conn, err := connect_db.New(psqlInfo, "postgres")
	if err != nil {
		log.Fatal().Err(err).Msg("connect_db error")
	}

	closer := GetInstance()
	closer.Add(func() error {
		return conn.Close()
	})

	if err := closer.CloseAll(); err != nil {
		log.Error().Err(err).Msg("errors occurred during shutdown")
	} else {
		log.Info().Msg("successfully closer all")
	}
}
