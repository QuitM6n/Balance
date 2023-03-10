package methods

import (
	"context"
	"fmt"

	"github.com/balance/api/database"
)

func (p Postgres) DeleteBalance(ctx context.Context, userID int64) error {
	if userID <= 0 {
		return fmt.Errorf("Wrong enter user id")
	}
	
	p.conn = database.ConnectDB()
	_, err := p.conn.Exec(ctx, "DELETE FROM BALANCE WHERE user_id = $1", userID)
	if err != nil {
		lgzap.Error(err.Error() + "Failed to delete user balance")
		return err
	}

	return nil
}
