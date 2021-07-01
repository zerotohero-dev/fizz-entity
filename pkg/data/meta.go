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

var Keys = struct {
	Email                  string
	Status                 string
	Name                   string
	EmailVerificationToken string
	RecordUpdated          string
	RecordCreated          string
	Password               string
}{
	Email:                  "email",
	Status:                 "status",
	Name:                   "Name",
	EmailVerificationToken: "emailVerificationToken",
	RecordCreated:          "recordCreated",
	RecordUpdated:          "recordUpdated",
	Password:               "password",
}

// Status enum.
var Status = struct {
	Unverified string
	Verified   string
	Blocked    string
}{
	Unverified: "unverified",
	Verified:   "verified",
	Blocked:    "blocked",
}
