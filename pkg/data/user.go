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

package data

import "time"

type User struct {
	// Info
	Email string `json:"email" bson:"email"`
	Name  string `json:"name" bson:"name"`
	// Token
	EmailVerificationToken string `json:"emailVerificationToken" bson:"emailVerificationToken"`
	PasswordResetToken     string `json:"passwordResetToken" bson:"passwordResetToken"`
	// Rest
	Password                string    `json:"password" bson:"password"`
	Status                  string    `json:"status" bson:"status"`
	SubscribedToMailingList bool      `json:"optIn" bson:"optIn"`
	RecordCreated           time.Time `json:"recordCreated" bson:"recordCreated"`
	RecordUpdated           time.Time `json:"recordUpdated" bson:"recordUpdated"`
}
