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

package reqres

type UserInfoRequest struct {
	AuthToken string `json:"token"`
	Err       string `json:"err,omitempty"`
}

type UserInfoResponse struct {
	Email        string `json:"email"`
	Subscription string `json:"subscription"`
	Name         string `json:"name"`
	Err          string `json:"err,omitempty"`
}
