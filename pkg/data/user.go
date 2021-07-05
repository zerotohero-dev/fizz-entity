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

type Info struct {
	Email string `json:"email" bson:"email"`
	Name  string `json:"name" bson:"name"`
}

type Token struct {
	EmailVerificationToken string `json:"emailVerificationToken" bson:"emailVerificationToken"`
	PasswordResetToken     string `json:"passwordResetToken" bson:"passwordResetToken"`
}

type User struct {
	Info
	Token
	Password                string    `json:"password" bson:"password"`
	Status                  string    `json:"status" bson:"status"`
	SubscribedToMailingList bool      `json:"optIn" bson:"optIn"`
	RecordCreated           time.Time `json:"recordCreated" bson:"recordCreated"`
	RecordUpdated           time.Time `json:"recordUpdated" bson:"recordUpdated"`
}
