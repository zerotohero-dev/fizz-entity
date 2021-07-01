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

package collection

import (
	"github.com/zerotohero-dev/fizz-entity/pkg/connection"
	"go.mongodb.org/mongo-driver/mongo"
)

func Collection(dbName, tableName string) *mongo.Collection {
	return connection.Db(dbName).Collection(tableName)
}
