/*
 *  \
 *  \\,
 *   \\\,^,.,,.                    “Zero to Hero”
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package data

type LoginResult struct {
	Token string
}

type Info struct {
	Email string `json:"email"`
	Name  string `json:"fullName"`
}

type Status string

const Unverified Status = "unverified"
const Verified Status = "active"
const Blocked Status = "blocked"

type User struct {
	Info
	Password                string `json:"password"`
	SubscribedToMailingList bool   `json:"subscribedToMailingList"`
	Status                  Status `json:"status"`
	EmailVerificationToken  string `json:"accountActivationToken"`
	PasswordResetToken      string `json:"passwordResetToken"`
	RecordCreated           int64  `json:"recordCreated"`
	RecordUpdated           int64  `json:"recordUpdated"`
}
