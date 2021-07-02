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
	Email string `json:"email"`
	Name  string `json:"fullName"`
}

type Token struct {
	EmailVerificationToken string `json:"emailVerificationToken"`
	PasswordResetToken     string `json:"passwordResetToken"`
}

type User struct {
	Info
	Token
	Password                string    `json:"password"`
	Status                  string    `json:"status"`
	SubscribedToMailingList bool      `json:"optIn"`
	RecordCreated           time.Time `json:"recordCreated"`
	RecordUpdated           time.Time `json:"recordUpdated"`
}
