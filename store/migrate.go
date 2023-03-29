package store

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

func (p postgres) Migrate(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			migrationSQL, err := os.ReadFile(filepath.Join(path, file.Name()))
			if err != nil {
				return err
			}

			_, err = p.conn.Exec(context.Background(), string(migrationSQL))
			if err != nil {
				return fmt.Errorf("error in migration %s. %w", file.Name(), err)
			}
		}
	}
	return nil
}
