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

type HashCreateRequest struct {
	Value string `json:"value"`
	Err   string `json:"err,omitempty"`
}

type HashCreateResponse struct {
	Hash string `json:"hash"`
	Err  string `json:"err,omitempty"`
}

type HashVerifyRequest struct {
	Value string `json:"value"`
	Hash  string `json:"hash"`
	Err   string `json:"err,omitempty"`
}

type HashVerifyResponse struct {
	Verified bool   `json:"verified"`
	Err      string `json:"err,omitempty"`
}
