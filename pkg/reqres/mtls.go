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

import "github.com/zerotohero-dev/fizz-entity/pkg/method"

type FizzServiceName string

const IdmService FizzServiceName = "idm"
const CryptoService FizzServiceName = "crypto"
const MailerService FizzServiceName = "mailer"

type MtlsApiRequest struct {
	Service  FizzServiceName `json:"service"`
	Endpoint string          `json:"endpoint"`
	Method   method.Method   `json:"method"`
	Body     string          `json:"body"`
}
