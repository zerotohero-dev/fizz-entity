/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package connection

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func CloseCursor(cur *mongo.Cursor, ctx context.Context) error {
	if cur == nil {
		return nil
	}

	return cur.Close(ctx)
}

func CreateDbContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), connectionTimeout)
}
