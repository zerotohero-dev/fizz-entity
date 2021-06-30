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

type Status string

const Unverified Status = "unverified"
const Verified Status = "active"
const Blocked Status = "blocked"

type Info struct {
	Email string `json:"email"`
	Name  string `json:"fullName"`
}

type Token struct {
	Email                  string `json:"email"`
	EmailVerificationToken string `json:"emailVerificationToken"`
	PasswordResetToken     string `json:"passwordResetToken"`
}

type User struct {
	Info
	Password                string `json:"password"`
	Status                  Status `json:"status"`
	SubscribedToMailingList bool   `json:"optIn"`
	RecordCreated           int64  `json:"recordCreated"`
	RecordUpdated           int64  `json:"recordUpdated"`
}
